package main

import (
	"fmt"
	DBS "go_API/internal/Database"
	_ "go_API/internal/model"
	_"github.com/gofiber/fiber/v2"
	"go_API/sevices"
)

func main(){
	//connect to postgres database
	db,_:=DBS.ConnectToDatabase()
	fmt.Print(db)
	sevices.Root(db)
}




//db.AutoMigrate(&model.Vocab{}, &model.Category{})
	//fmt.Print(db)
	//fmt.Print(DBS.List_category())
	//DBS.List_vocab()
	//DBS.AddData_toDb(db)