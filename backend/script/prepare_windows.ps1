<#
.SYNOPSIS
Windows 开发环境自动配置脚本（Scoop/Docker/Node/Go/PM2）
.NOTES
保存编码：UTF-8 with BOM | 运行权限：普通/管理员均可 | 兼容：PowerShell 5.1+、所有 Scoop 版本
#>
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Continue'  # 改为Continue，避免非致命错误直接退出

# 日志函数（纯英文避免编码问题）
function Log { param($m) Write-Host "==> $m" -ForegroundColor Cyan }
function Warn { param($m) Write-Host "⚠️ $m" -ForegroundColor Yellow }
function ErrorLog { param($m) Write-Host "❌ $m" -ForegroundColor Red }

# 错误处理：仅严重错误提示，不强制退出
function ErrTrap {
    param($Line, $Exception)
    ErrorLog "Error at line $Line : $($Exception.Message)"
    # 仅记录错误，不退出（避免整个脚本中断）
}
trap {
    # 过滤已知非致命错误
    if ($_.Exception.Message -match '不支持所指定的方法|not supported|permission denied|权限|Option -q not recognized') {
        Warn "Non-critical error: $($_.Exception.Message)"
        continue
    } elseif ($_.Exception.Message -notmatch 'exists|已存在|skip|跳过') {
        ErrTrap $($_.InvocationInfo.ScriptLineNumber) $_.Exception
        continue
    } else {
        Warn $_.Exception.Message
        continue
    }
}

# 检测管理员权限
$currentPrincipal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
$IsAdmin = $currentPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $IsAdmin) {
    Warn "Not running as administrator! PM2 service install/Docker auto-start will be skipped."
}

# ========== Scoop 安装 ==========
if (-not (Get-Command scoop -ErrorAction SilentlyContinue)) {
    Log "Installing Scoop (non-interactive)"
    try {
        Set-ExecutionPolicy RemoteSigned -Scope CurrentUser -Force -ErrorAction Stop
        [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
        Invoke-RestMethod -Uri https://get.scoop.sh -UseBasicParsing | Invoke-Expression
    } catch {
        ErrorLog "Scoop install failed: $($_.Exception.Message)"
    }
}

# 配置 Scoop Buckets
$neededBuckets = @('main','extras')
$bucketList = @(& scoop bucket list 2>$null)
foreach ($b in $neededBuckets) {
    if ($bucketList -notcontains $b) {
        Log "Adding Scoop bucket: $b"
        try { & scoop bucket add $b --no-update 2>$null }
        catch { Warn "Failed to add bucket $b : $($_.Exception.Message)" }
    } else {
        Log "Bucket $b already exists, skip"
    }
}

# 安装基础工具（移除 -q 参数，兼容旧版 Scoop）
$pkgs = @('wget','unzip','git','jq','make','grep','gawk','sed','touch','mingw','nodejs','go')
foreach ($p in $pkgs) {
    if (-not (Get-Command $p -ErrorAction SilentlyContinue)) {
        Log "Installing tool: $p"
        try { & scoop install $p 2>$null }  # 移除 -q 参数
        catch { Warn "Failed to install $p : $($_.Exception.Message)" }
    } else {
        Log "$p already installed, skip"
    }
}

# ========== Docker 安装 ==========
function Install-DockerDesktop {
    Log "Installing Docker Desktop"
    if (Get-Command winget -ErrorAction SilentlyContinue) {
        try {
            & winget install --id Docker.DockerDesktop `
                -e --accept-package-agreements --accept-source-agreements `
                --silent --disable-interactivity 2>$null
            Log "Docker Desktop install submitted via Winget (wait for background completion)"
        } catch {
            Warn "Winget install Docker failed: $($_.Exception.Message)"
        }
    } else {
        Warn "Winget not found, try install Docker CLI via Scoop"
        try { & scoop install docker 2>$null }  # 移除 -q 参数
        catch { Warn "Scoop install Docker CLI failed: $($_.Exception.Message)" }
    }
}

# 执行 Docker 安装
try { Install-DockerDesktop }
catch { Warn "Docker install process error: $($_.Exception.Message)" }

# 配置 Docker 服务（仅管理员）
$dockerServiceName = $null
if (Get-Service -Name com.docker.service -ErrorAction SilentlyContinue) {
    $dockerServiceName = "com.docker.service"
} elseif (Get-Service -Name docker -ErrorAction SilentlyContinue) {
    $dockerServiceName = "docker"
}

if ($dockerServiceName -and $IsAdmin) {
    Log "Configuring Docker service: $dockerServiceName"
    try {
        Set-Service -Name $dockerServiceName -StartupType Automatic -ErrorAction Stop
        Start-Service -Name $dockerServiceName -ErrorAction Stop
        Log "Docker service set to auto-start and started successfully"
    } catch {
        Warn "Failed to operate Docker service: $($_.Exception.Message)"
    }
} elseif ($dockerServiceName) {
    Warn "Non-admin: skip Docker service configuration (run as admin to auto-start)"
} else {
    Warn "Docker service not found (Docker Desktop may not be installed)"
}

# ========== NPM 全局工具（仅当前会话生效，避免权限错误） ==========
function Npm-GlobalInstall {
    param([string[]]$packages)

    if (-not (Get-Command npm -ErrorAction SilentlyContinue)) {
        Warn "NPM not installed, skip global packages: $($packages -join ', ')"
        return
    }

    # 强制使用用户目录安装（不修改全局配置，避免权限错误）
    Log "Installing NPM packages to user directory (no global config change)"
    $userNpmDir = Join-Path $env:USERPROFILE ".npm-global"
    New-Item -Path $userNpmDir -ItemType Directory -Force | Out-Null

    try {
        # 直接指定--prefix安装，不修改npm全局配置
        & npm install --prefix "$userNpmDir" -g $packages --silent 2>$null
        # 仅当前会话添加PATH
        $npmBinPath = Join-Path $userNpmDir "bin"
        if (-not ($env:PATH -like "*$npmBinPath*")) {
            $env:PATH += ";$npmBinPath"
            Log "Added NPM bin path to current session: $npmBinPath"
        }
    } catch {
        Warn "NPM global install failed: $($_.Exception.Message)"
        Warn "Manual install: npm install -g --prefix ""$userNpmDir"" pm2 pm2-windows-service"
    }
}

# 安装 PM2（仅用户目录）
Npm-GlobalInstall @('pm2', 'pm2-windows-service')

# ========== PM2 服务安装（仅管理员，且简化逻辑） ==========
if ($IsAdmin) {
    Log "Configuring PM2 Windows service (unattended)"
    try {
        # 查找PM2路径（用户目录）
        $userNpmDir = Join-Path $env:USERPROFILE ".npm-global"
        $pm2Path = Join-Path $userNpmDir "bin\pm2.cmd"

        if (Test-Path $pm2Path) {
            & $pm2Path service-install --name pm2 --unattended 2>$null
            # 检查服务是否存在再启动
            if (Get-Service -Name pm2 -ErrorAction SilentlyContinue) {
                Start-Service -Name pm2 -ErrorAction SilentlyContinue
                Log "PM2 service installed (check status with Get-Service pm2)"
            } else {
                Warn "PM2 service install completed but service not found (manual check required)"
            }
        } else {
            Warn "PM2 executable not found at $pm2Path (install may have failed)"
        }
    } catch {
        Warn "PM2 service install failed: $($_.Exception.Message)"
        Warn "Manual install: cd ""$userNpmDir\bin"" && .\pm2.cmd service-install --unattended"
    }
} else {
    Warn "Non-admin: skip PM2 service install (run as admin to install)"
    Warn "Manual PM2 install: npm install -g --prefix ""$env:USERPROFILE\.npm-global"" pm2 pm2-windows-service"
}

# ========== Go 环境配置（仅当前会话生效，避免写入错误） ==========
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Log "Installing Go environment"
    try { & scoop install go 2>$null }  # 移除 -q 参数
    catch { Warn "Scoop install Go failed: $($_.Exception.Message)" }
}

# 配置GOPATH（仅当前会话）
$gopath = Join-Path $env:USERPROFILE "go"
New-Item -Path $gopath -ItemType Directory -Force | Out-Null
if (-not (Test-Path env:GOPATH)) {
    $env:GOPATH = $gopath
    Log "Set GOPATH for current session: $gopath"
}
$goBinPath = Join-Path $gopath "bin"
if (-not ($env:PATH -like "*$goBinPath*")) {
    $env:PATH += ";$goBinPath"
    Log "Added GOPATH/bin to current session: $goBinPath"
}

# ========== 手动配置提示（避免自动写入Profile触发错误） ==========
Log "✅ Environment setup completed (current session only)!"
Write-Host @"

==================== MANUAL CONFIG TIPS ====================
1. To make NPM path permanent:
   Add this line to your PowerShell Profile:
   `$env:PATH += ";$($env:USERPROFILE)\.npm-global\bin"

2. To make GOPATH permanent:
   Add these lines to your PowerShell Profile:
   `$env:GOPATH = "$gopath"
   `$env:PATH += ";$gopath\bin"

3. To find PowerShell Profile path:
   echo `$PROFILE

4. PM2 service install (run as admin):
   cd "$($env:USERPROFILE)\.npm-global\bin" && .\pm2.cmd service-install --unattended

5. Check service status (admin):
   Get-Service com.docker.service, pm2
"@ -ForegroundColor Green