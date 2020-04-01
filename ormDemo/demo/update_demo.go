package demo

import "ormDemo/model"

func UpdateLanguage2User(user *model.User, languages []model.Language) {
	db := model.GetDB()
	user.Languages = languages
	err := db.Model(&user).Update(user).Error
	if err != nil {
		panic(err)
	}
}
