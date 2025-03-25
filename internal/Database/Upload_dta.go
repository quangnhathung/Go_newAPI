package database

import (
	"fmt"
	"go_API/internal/model"
	"strings"

	"gorm.io/gorm"
)

func AddData_toDb(db *gorm.DB) {
	categories := List_category()
	vocabularies := List_vocab()

	//write to table categories
	for _,cate := range categories{
		var cat model.Category
		db.Where("name = ?", cate.Name).FirstOrCreate(&cat, model.Category{Name: cate.Name})
	}

	default_path :="https://storage.googleapis.com/static.saoladigital.com/public/npvu1510/"
	
	//write to table vocab
	for _,voc := range vocabularies{
		var vocab model.Vocab
		db.Where("vocabulary = ?", voc.Vocab).FirstOrCreate(&vocab, model.Vocab{
			Vocabulary: voc.Vocab,
			Image: fmt.Sprintf("%simages/%s.svg",default_path,voc.Vocab),
			AudioUS: fmt.Sprintf("%saudios/%s_us.mp3",default_path,voc.Vocab),
			AudioGB: fmt.Sprintf("%saudios/%s_gb.mp3",default_path,voc.Vocab),
		})

		var matchedCategories []model.Category
		for _, cate := range categories {
			if strings.Contains(voc.identify, cate.Id) {
				var existingCategory model.Category
				if err := db.Where("name = ?", cate.Name).First(&existingCategory).Error; err == nil {
					matchedCategories = append(matchedCategories, existingCategory)
				}
			}
		}
		db.Model(&vocab).Association("Categories").Append(matchedCategories)
	}

	fmt.Println("Sucess")
}