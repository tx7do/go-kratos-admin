# GoWindAdmin（GoWind 管理システム）

> GoWindAdmin — **企業向け管理システムを効率的に構築し、開発を風のようにスムーズにする。**

すぐに使える Golang フルスタック管理システム。バックエンドは GO マイクロサービスフレームワーク [go-kratos](https://go-kratos.dev/)、フロントエンドは Vue マイクロサービスフレームワーク [Vben Admin](https://doc.vben.pro/) をベースにしています。

どちらもマイクロサービスフレームワークを使用していますが、フロントエンドとバックエンドはモノリス構成で開発・デプロイすることも可能です。

導入が容易で機能が充実しており、企業向け管理システムの迅速な開発に適しています。

[English](./README.en-US.md) | **中文** | [日本語](./README.ja-JP.md)

## デモ

> フロントエンド：<http://124.221.26.30:8080/>
>
> バックエンド Swagger：<http://124.221.26.30:7788/docs/>
>
> デフォルトアカウント/パスワード：`admin` / `admin`

## 技術スタック

- バックエンド: [Golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/) + [wire](https://github.com/google/wire) + [ent](https://entgo.io/docs/getting-started/)
- フロントエンド: [Vue](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/) + [Ant Design Vue](https://antdv.com/) + [Vben Admin](https://doc.vben.pro/)

## クイックスタート

1. Docker と Go をインストール（参照：`backend/script/prepare_ubuntu.sh`、`backend/script/prepare_centos.sh`、`backend/script/prepare_rocky.sh`）
2. `backend` ディレクトリに移動し、以下のコマンドを実行してバックエンドサービス `kratos-admin` をコンパイルし、Docker イメージをビルドして依存サービスとともに起動します：
~~~bash
make init
make docker
make compose-up
~~~
3. npm と pnpm をインストール（インストール方法は問い合わせ可）
4. `frontend` ディレクトリに移動し、以下のコマンドを実行してフロントエンドをビルド・起動（開発モード）します：
~~~bash
pnpm install
pnpm dev
~~~
5. 動作確認

- フロントエンド：<http://localhost:5666>、ログイン：`admin` / `admin`
- バックエンド：<http://localhost:7788/docs/openapi.yaml>

## 機能一覧

| 機能 | 説明 |
|------|--------------------------------------------------------------------------|
| ユーザー管理 | ユーザーの管理と検索。高度な検索、部署連動によるユーザー選択、ユーザーの有効/無効化、上長設定/解除、パスワードリセット、複数ロール・複数部署・上長設定、一鍵で指定ユーザーとしてログイン等をサポート。 |
| テナント管理 | テナント管理。テナント追加時に自動でテナント部署、デフォルトロール、管理者を初期化。プラン設定、有効/無効、一鍵でテナント管理者としてログイン可能。 |
| ロール管理 | ロールとロールグループの管理。ロールによるユーザー連動、メニューとデータ権限の設定、社員の一括追加/削除をサポート。 |
| 組織管理 | 組織の管理、ツリー表示をサポート。 |
| 部署管理 | 部署の管理、ツリー表示をサポート。 |
| 権限管理 | 権限グループ、メニュー、権限ポイントの管理、ツリー表示をサポート。 |
| API管理 | API の管理、API 同期機能（主に権限ポイント追加時のインターフェース選択用）、ツリー表示、操作ログのリクエストパラメータとレスポンス設定をサポート。 |
| 辞書管理 | 辞書カテゴリとエントリの管理。カテゴリ連動、サーバー側の複数列ソート、データのインポート/エクスポートをサポート。 |
| タスクスケジューラ | タスクと実行ログの管理。作成、更新、削除、開始、停止、即時実行をサポート。 |
| ファイル管理 | ファイルアップロードの管理、検索、OSS またはローカルへのアップロード、ダウンロード、URL コピー、削除、画像の大きいプレビューをサポート。 |
| メッセージカテゴリ | メッセージカテゴリの管理（2階層のカスタムカテゴリ）をサポート。 |
| メッセージ管理 | 指定ユーザーへのメッセージ送信、既読状況と既読時間の確認をサポート。 |
| 站内信（社内メッセージ） | 社内メッセージ管理。詳細表示、削除、既読、全既読をサポート。 |
| 個人センター | 個人情報の表示・編集、最終ログイン情報の表示、パスワード変更など。 |
| キャッシュ管理 | キャッシュ一覧の照会とキー指定でのクリアをサポート。 |
| ログインログ | 成功・失敗のログインログを照会、IP 位置情報をサポート。 |
| 操作ログ | 正常／異常の操作ログを照会、IP 位置情報と操作詳細の確認をサポート。 |

## バックエンドのスクリーンショット

<table>
    <tr>
        <td><img src="./docs/images/admin_login_page.png" alt="管理者ログインページ"/></td>
        <td><img src="./docs/images/admin_dashboard.png" alt="管理ダッシュボード"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_user_list.png" alt="ユーザー一覧"/></td>
        <td><img src="./docs/images/admin_user_create.png" alt="ユーザー作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_tenant_list.png" alt="テナント一覧"/></td>
        <td><img src="./docs/images/admin_tenant_create.png" alt="テナント作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_organization_list.png" alt="組織一覧"/></td>
        <td><img src="./docs/images/admin_organization_create.png" alt="組織作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_department_list.png" alt="部署一覧"/></td>
        <td><img src="./docs/images/admin_department_create.png" alt="部署作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_position_list.png" alt="職位一覧"/></td>
        <td><img src="./docs/images/admin_position_create.png" alt="職位作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_role_list.png" alt="ロール一覧"/></td>
        <td><img src="./docs/images/admin_role_create.png" alt="ロール作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_menu_list.png" alt="メニュー一覧"/></td>
        <td><img src="./docs/images/admin_menu_create.png" alt="メニュー作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_task_list.png" alt="タスク一覧"/></td>
        <td><img src="./docs/images/admin_task_create.png" alt="タスク作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_dict_list.png" alt="辞書一覧"/></td>
        <td><img src="./docs/images/admin_dict_entry_create.png" alt="辞書エントリ作成"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_internal_message_list.png" alt="社内メッセージ一覧"/></td>
        <td><img src="./docs/images/admin_internal_message_publish.png" alt="社内メッセージ発行"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_login_restriction_list.png" alt="ログイン制限一覧"/></td>
        <td><img src="./docs/images/admin_api_resource_list.png" alt="API リソース一覧"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_operation_log_list.png" alt="操作ログ一覧"/></td>
        <td><img src="./docs/images/admin_login_log_list.png" alt="ログインログ一覧"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/api_swagger_ui.png" alt="バックエンド内蔵 Swagger UI"/></td>
    </tr>
</table>

## お問い合わせ

- WeChat: `yang_lin_bo`（備考：`go-wind-admin`）
- 掘金コラム: [go-wind-admin](https://juejin.cn/column/7541283508041826367)

## JetBrains に感謝（無料の GoLand & WebStorm 提供）

[![avatar](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://jb.gg/OpenSource)
