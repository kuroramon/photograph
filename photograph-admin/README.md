# photograph-admin 写真管理アプリ

## 環境構築

### 1.envファイル作成

- ルートディレクトリ直下に.envファイルを作成
- 新規で作成して.env.orgの内容コピーするでもよし、.env.orgをコピーして.envにリネームするでもよし

### 2.docker起動

```
docker-compose up
```

### 3 migrate実施

- goのコンテナに入る
```
docker exec -it webdrawer_server sh
```
- 以下のコマンドでマイグレーションを実行
```
go run db/migrate/migrate.go
```