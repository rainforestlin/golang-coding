package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ormDemo/demo"
	"ormDemo/model"
)

func main() {
	var languages []model.Language
	//languages = demo.GetAllLanguages()
	user := demo.GetUser("LJL")
	languages = demo.RelatedLanguage2User(&user)
	fmt.Println(languages)
	//demo.UpdateLanguage2User(&user, languages)
}
