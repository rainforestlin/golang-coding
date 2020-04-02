package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ormDemo/model"
)

func main() {
	//var languages model.Language
	//languages = demo.GetAllLanguages()
	//maps := make(map[string]interface{})
	//maps["id"] = 1
	//user := demo.GetUser(maps)
	//demo.ClearAssociation(&user)
	//languages = demo.RelatedLanguage2User(&user)
	//languages = demo.RelatedLanguage2UserFirst(&user)
	//fmt.Println(languages)

	//demo.UpdateLanguage2User(&user, languages)
	var user model.User
	db := model.GetDB()
	db.First(&user)
	fmt.Println(user.CreatedAt)
}
