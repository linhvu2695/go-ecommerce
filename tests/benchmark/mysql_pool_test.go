package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "testUser"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func benchmarkMaxOpenConns(b *testing.B, connsCount int) {
	dsnTemplate := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnTemplate, "root", "root1234", "127.0.0.1", "3306", "db")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Check if table exists
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
	}

	db.AutoMigrate(&User{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get SQL db: %v", err)
	}

	sqlDb.SetMaxOpenConns(connsCount)
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConn1(b *testing.B) {
	benchmarkMaxOpenConns(b, 1)
}

func BenchmarkMaxOpenConn10(b *testing.B) {
	benchmarkMaxOpenConns(b, 10)
}
