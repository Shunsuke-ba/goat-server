# goat-server

### ディレクトリ構造
- /cmd
/api https://github.com/Shunsuke-ba/goat-server/cmd/api/README.md(主に外部APIとの連携とサービス機能)
/cron coming soon ...

- /mod
/app - domain ドメイン
     - driver　インフラストラクチャ
     - handler　インターフェース
     - usecase ユースケース

/uil - config 諸config


### 概要
- 頑張る
- 頑張る


### migrate
- make .sql
ex. sql-migrate new (create_users)
→ 20201019153137-create_users.sql

- exe migration
sql-migrate up

- migrate status
sql-migrate status