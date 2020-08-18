# gorm検証

## TODO

マイグレーション機能の検証

[ ] drop tableでテーブルが削除されるか
[ ] drop collumで列が削除されるか

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
