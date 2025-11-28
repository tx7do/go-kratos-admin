Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

function Log { param($m) Write-Host "==> $m" }
function ErrTrap { param($Line) Write-Host "Error at line $Line" -ForegroundColor Red ; exit 1 }
trap { ErrTrap $($_.InvocationInfo.ScriptLineNumber) }

# 检测是否以管理员运行
$IsAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $IsAdmin) {
    Log "未以管理员身份运行，某些操作（安装系统服务、Docker Desktop）可能会被跳过或提示权限错误。"
}

# 确保 Scoop 存在
if (-not (Get-Command scoop -ErrorAction SilentlyContinue)) {
    Log "安装 Scoop（非交互）"
    Try {
        Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned -Force
        Invoke-RestMethod -Uri https://get.scoop.sh -UseBasicParsing | Invoke-Expression
    } Catch {
        Write-Warning "Scoop 安装失败：$($_.Exception.Message)"
        throw
    }
}

# 确保常用 buckets 存在
$neededBuckets = @('main','extras')
foreach ($b in $neededBuckets) {
    $exists = (& scoop bucket list) -match "^$b`$"
    if (-not $exists) {
        Log "添加 scoop bucket: $b"
        & scoop bucket add $b
    }
}

# 基础工具清单并安装（使用 scoop）
$pkgs = @('wget','unzip','git','jq','make','grep','gawk','sed','touch','mingw','nodejs','go')
foreach ($p in $pkgs) {
    if (-not (Get-Command $p -ErrorAction SilentlyContinue)) {
        Log "安装 $p"
        & scoop install $p
    } else {
        Log "$p 已存在，跳过"
    }
}

# 优先使用 winget 安装 Docker Desktop（非交互），若不存在再尝试 scoop
function Install-DockerDesktop {
    if (Get-Command winget -ErrorAction SilentlyContinue) {
        Log "使用 winget 安装 Docker Desktop（可能需要管理员）"
        & winget install --id Docker.DockerDesktop -e --accept-package-agreements --accept-source-agreements
    } else {
        Log "winget 不可用，尝试使用 scoop 安装 docker（可能仅为 CLI）"
        & scoop install docker
    }
}

# 安装或确保 Docker
try {
    Install-DockerDesktop
} catch {
    Write-Warning "安装 Docker 失败：$($_.Exception.Message)"
}

# 启动并设置 Docker 服务自动启动（如果存在）
try {
    if (Get-Service -Name docker -ErrorAction SilentlyContinue) {
        Log "启动 docker 服务并设为自动"
        if ($IsAdmin) {
            Start-Service docker -ErrorAction SilentlyContinue
            Set-Service -Name docker -StartupType Automatic
        } else {
            Write-Warning "非管理员：无法修改服务启动类型，请以管理员身份运行以启用 docker 服务开机自启。"
        }
    } else {
        Log "docker 服务不存在（可能使用 Docker Desktop），请手动确认。"
    }
} catch {
    Write-Warning "操作 Docker 服务出错：$($_.Exception.Message)"
}

# 安装 npm 全局工具 (pm2, pm2-windows-service) 并处理权限问题
function Npm-GlobalInstall {
    param($packages)

    # 检查是否有权限写入全局 npm 前缀
    try {
        $globalPrefix = npm config get prefix 2>$null
    } catch {
        $globalPrefix = $null
    }

    $needsUserPrefix = $false
    if ($globalPrefix) {
        try {
            $testFile = Join-Path $globalPrefix "npm_install_test.txt"
            New-Item -Path $testFile -ItemType File -Force | Out-Null
            Remove-Item $testFile -Force
        } catch {
            $needsUserPrefix = $true
        }
    } else {
        $needsUserPrefix = $true
    }

    if ($needsUserPrefix) {
        Log "设置 npm 全局前缀到用户目录以避免权限问题"
        $userPrefix = Join-Path $env:USERPROFILE ".npm-global"
        New-Item -ItemType Directory -Path $userPrefix -Force | Out-Null
        & npm config set prefix $userPrefix
        $profilePath = "$env:USERPROFILE\Documents\WindowsPowerShell\profile.ps1"
        $pathLine = "if (-not ($env:PATH -like '*$userPrefix*')) { [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ';$userPrefix\bin','User') }"
        if (-not (Test-Path $profilePath)) { New-Item -ItemType File -Path $profilePath -Force | Out-Null }
        if (-not (Select-String -Path $profilePath -Pattern [regex]::Escape($userPrefix) -Quiet)) {
            Add-Content -Path $profilePath -Value $pathLine
            Log "已在 $profilePath 写入 PATH 更新，请重新打开终端以生效"
        }
    }

    Log "通过 npm 全局安装: $packages"
    & npm install -g $packages
}

Npm-GlobalInstall "pm2 pm2-windows-service"

# 尝试自动为 pm2 安装 Windows 服务（仅在管理员下）
if ($IsAdmin) {
    try {
        Log "尝试以非交互方式安装 pm2 Windows 服务（若脚本提示交互则会失败）"
        # pm2-windows-service 的交互选项可能有限，先尝试直接调用安装命令
        & pm2-service-install --name "pm2" --pm2-path (Join-Path (Split-Path (Get-Command npm).Path) '..\node_modules\pm2\bin\pm2') 2>$null
        Log "pm2 Windows 服务安装命令已执行（检查服务状态以确认）"
    } catch {
        Write-Warning "自动安装 pm2 Windows 服务失败，请以管理员手动运行 `pm2-service-install` 或使用 pm2-windows-service 文档中的非交互方式。"
    }
} else {
    Write-Warning "非管理员：跳过 pm2 Windows 服务自动安装。请以管理员运行 `pm2-service-install` 来完成服务注册。"
}

# 安装 Go（若未通过 scoop 安装）
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Log "安装 Go"
    & scoop install go
}

# 设置 GOPATH（写入用户 profile，如未存在）
$gopath = Join-Path $env:USERPROFILE "go"
if (-not (Test-Path $gopath)) { New-Item -ItemType Directory -Path $gopath | Out-Null }
$profileFile = "$env:USERPROFILE\Documents\WindowsPowerShell\profile.ps1"
$gopathLine = @"
if (-not (Test-Path env:GOPATH)) {
    [Environment]::SetEnvironmentVariable('GOPATH', '$gopath', 'User')
    [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ';' + '$gopath\bin', 'User')
}
"@
if (-not (Select-String -Path $profileFile -Pattern 'GOPATH' -Quiet)) {
    Add-Content -Path $profileFile -Value $gopathLine
    Log "已在 $profileFile 写入 GOPATH 设置，请重新打开终端以生效"
}

# 清理与状态输出
Log "安装完成。请注意："
if (-not $IsAdmin) {
    Write-Host "- 当前非管理员，部分步骤（Docker Desktop 安装、服务注册）需要管理员权限完成。" -ForegroundColor Yellow
}
Write-Host "- 如更改了用户级 PATH 或 npm prefix，请重新启动终端。" -ForegroundColor Green
Write-Host "- 检查 pm2 服务：Get-Service -Name pm2 (或在服务管理器中查找)" -ForegroundColor Green
