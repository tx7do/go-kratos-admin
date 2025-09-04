# Go Kratos Admin

すぐに使えるGolangフルスタック管理システム。

バックエンドはGOマイクロサービスフレームワーク[go-kratos](https://go-kratos.dev/)に基づき、フロントエンドはVueマイクロサービスフレームワーク[Vben Admin](https://doc.vben.pro/)に基づいています。

両方ともマイクロサービスフレームワークを使用していますが、フロントエンドとバックエンドはモノリシックアーキテクチャで開発およびデプロイすることも可能です。

簡単に始められ、機能が豊富で、エンタープライズレベルの管理システムの迅速な開発に適しています。

[English](./README.en-US.md) | [中文](./README.md) | **日本語**

## デモアドレス

> フロントエンド: <http://124.221.26.30:8080/>
>
> バックエンドSwagger: <http://124.221.26.30:7788/docs/>
>
> デフォルトのユーザー名とパスワード: `admin` / `admin`

## 技術スタック

- バックエンド: [Golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/) + [wire](https://github.com/google/wire) + [ent](https://entgo.io/docs/getting-started/)
- フロントエンド: [Vue](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/) + [Ant Design Vue](https://antdv.com/) + [Vben Admin](https://doc.vben.pro/)

## クイックスタートガイド

1. DockerとGolangをインストールします（`backend/script/prepare_ubuntu.sh`、`backend/script/prepare_centos.sh`、`backend/script/prepare_rocky.sh`を参照）。
2. `backend`ディレクトリに移動し、以下のコマンドを実行してバックエンドサービス`kratos-admin`をコンパイルし、Dockerイメージをビルドして、依存するDockerサービスと一緒に起動します：
    ```bash
    make init
    make docker
    make compose-up
    ```
3. npmとpnpmをインストールします（インストール方法はAIに問い合わせ可能）。
4. `frontend`ディレクトリに移動し、以下のコマンドを実行してフロントエンドをコンパイルして起動します（開発モード）：
    ```bash
    pnpm install
    pnpm dev
    ```
5. テスト環境にアクセスします：

- フロントエンド: <http://localhost:5666>、ログイン情報: `admin` / `admin`
- バックエンド: <http://localhost:7788/docs/openapi.yaml>

## 機能一覧

| 機能             | 説明                                                                                     |
|------------------|------------------------------------------------------------------------------------------|
| ユーザー管理       | ユーザーの管理とクエリをサポート、高度なクエリと部門にリンクされたユーザーをサポート。ユーザーの有効化/無効化、監督者の設定/解除、パスワードのリセット、複数の役割、部門、監督者の設定、特定のユーザーとしてワンクリックでログイン可能。 |
| テナント管理       | テナントの管理。テナントを追加すると、テナントの部門、デフォルトの役割、管理者が自動的に初期化されます。パッケージ構成、有効化/無効化、テナント管理者へのワンクリックログインをサポート。 |
| 役割管理          | 役割と役割グループの管理、役割によるユーザーのリンクをサポート、メニューとデータ権限の設定、従業員の一括追加/削除をサポート。 |
| 組織管理          | 組織の管理、ツリーリスト表示をサポート。 |
| 部門管理          | 部門の管理、ツリーリスト表示をサポート。 |
| 権限管理          | 権限グループ、メニュー、権限ポイントの管理。ツリーリスト表示をサポート。 |
| API管理          | APIの管理、権限ポイントを追加する際のAPI選択のためのAPI同期をサポート。ツリーリスト表示、操作ログでのリクエストパラメータとレスポンス結果の構成をサポート。 |
| 辞書管理          | データ辞書カテゴリとそのサブカテゴリの管理。辞書カテゴリのリンク、サーバーでの複数列のソート、データのインポート/エクスポートをサポート。 |
| タスクスケジューリング | タスクとその実行ログの管理と表示。タスクの追加、変更、削除、開始、一時停止、即時実行をサポート。 |
| ファイル管理       | ファイルのアップロード管理。ファイルクエリ、OSSまたはローカルストレージへのアップロード、ダウンロード、ファイルアドレスのコピー、ファイルの削除、画像のフルサイズ表示をサポート。 |
| メッセージカテゴリ   | メッセージカテゴリの管理、メッセージ管理のための2レベルのカスタムメッセージカテゴリをサポート。 |
| メッセージ管理      | メッセージの管理、特定のユーザーへのメッセージ送信をサポート、ユーザーがメッセージを読んだかどうかと読んだ時間を表示。 |
| 受信トレイ         | 内部メッセージの管理、詳細メッセージの表示、削除、既読としてマーク、すべて既読としてマークをサポート。 |
| 個人センター       | 個人情報の表示と変更、最後のログイン情報の表示、パスワードの変更など。 |
| キャッシュ管理      | キャッシュリストのクエリ、キャッシュキーによるキャッシュのクリアをサポート。 |
| ログインログ       | ログインログのクエリ、ユーザーのログイン成功と失敗ログの記録、IP位置情報の記録をサポート。 |
| 操作ログ          | 操作ログのクエリ、ユーザーの操作成功と失敗ログの記録、IP位置情報の記録、操作ログの詳細表示をサポート。 |

## バックエンドスクリーンショット

<table>
    <tr>
        <td><img src="./docs/images/admin_login_page.png" alt="管理者ログインページ"/></td>
        <td><img src="./docs/images/admin_dashboard.png" alt="管理者ダッシュボード"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_user_management.png" alt="ユーザー管理ページ"/></td>
        <td><img src="./docs/images/admin_create_user.png" alt="ユーザー作成ページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_menu_management.png" alt="メニュー管理ページ"/></td>
        <td><img src="./docs/images/admin_organization_management.png" alt="組織管理ページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_department_management.png" alt="部門管理ページ"/></td>
        <td><img src="./docs/images/admin_role_management.png" alt="役割管理ページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_position_management.png" alt="ポジション管理ページ"/></td>
        <td><img src="./docs/images/admin_tenant_management.png" alt="テナント管理ページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_task_management.png" alt="タスク管理ページ"/></td>
        <td><img src="./docs/images/admin_dict_management.png" alt="辞書管理ページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_operation_log.png" alt="操作ログページ"/></td>
        <td><img src="./docs/images/admin_login_log.png" alt="ログインログページ"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/api_swagger_ui.png" alt="Swagger UIページ"/></td>
    </tr>
</table>

## お問い合わせ

- WeChat: `yang_lin_bo`（備考: `go-kratos-admin`）
- Juejinコラム: [go-kratos-admin](https://juejin.cn/column/7541283508041826367)

## [JetBrainsによる無料のGoLand提供に感謝](https://jb.gg/OpenSource)

[![avatar](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://jb.gg/OpenSource)