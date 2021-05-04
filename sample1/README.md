# 概要
nginxをリバースプロキシにしてgoのアプリに接続、goのアプリはDB接続してDatabase一覧を返す。

```
front-nginx -> back-go -> db-mysql
```

# build
```
docker-compose build
```

# 起動
```
docker-compose up -d
```

# 確認
```
curl http://localhost
```

# 停止
```
docker-compose down
```
