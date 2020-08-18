package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=gorm-db-test user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create table for `User`
	db.Migrator().CreateTable(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// ここから
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
}
