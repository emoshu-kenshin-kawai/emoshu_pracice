# emoshu 技術課題

## セットアップ

```
docker-compose build
docker-compose up

<!-- 初期データ挿入 -->
docker cp ./initdata/init.sql db:tmp
docker cp ./initdata/sample.sql db:tmp
docker-compose exec db bash
cd tmp/
mysql -u root -p  # password
source init.sql  # マスタデータ挿入 ex:) role,status
source sample.sql  # トランザクションデータ挿入 ex:) member
```

### Swagger初期化&更新

```
swag init
```