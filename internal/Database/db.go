package database

import (
	"fmt"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go_API/sevices"
)
func ConnectToDatabase() (*gorm.DB,error){

	//dsn := "host=localhost user=postgres password=hung05112005 dbname=vocab_db port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	dsn:=sevices.String_connect()
	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("REFUSE CONNECT TO DATABASE")
	}
	fmt.Println("Kết nối thành công đến PostgreSQL!")
	return db,nil
}