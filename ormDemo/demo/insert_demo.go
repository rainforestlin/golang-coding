package demo

import (
	"fmt"
	"ormDemo/model"
)

func Insert2User(user *model.User) {
	err := user.Insert()
	if err != nil {
		panic(err)
	}
}

func Insert2Language(languages []model.Language) {
	db := model.GetDB()
	fmt.Println(languages)
	for _, language := range languages {
		err := db.Model(language).Create(&language).Error
		if err != nil {
			panic(err)
		}
	}
}

