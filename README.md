# gorm検証

## 結果

マイグレーション機能の検証

- [x] drop tableでテーブルが削除されるか = OK
- [x] drop collumで列が削除されるか     = OK

```go
	// カラム名の変更
	db.Debug().Migrator().RenameColumn(&Product{}, "Code", "ChangeName")
	/* 出力
	[0.238ms] [rows:0] ALTER TABLE "products" RENAME COLUMN "code" TO "ChangeName"
	*/

	// カラムの削除
	db.Debug().Migrator().DropColumn(&Product{}, "price")
	/* 出力
	[0.206ms] [rows:0] ALTER TABLE "products" DROP COLUMN "price"
	*/

	// テーブルの削除
	db.Debug().Migrator().DropTable(&Product{})
	/* 出力
	[24.685ms] [rows:0] DROP TABLE IF EXISTS "products" CASCADE
	*/
```

## DB確認方法

`make up`

`make db-in`

`psql -U test`
