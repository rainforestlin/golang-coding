package demo

import "ormDemo/model"

func GetAllLanguages() (languages []model.Language) {
	db := model.GetDB()
	if err := db.Model(model.Language{}).Find(&languages); err != nil {
		return
	}
	return
}

func GetUser(name string) (user model.User) {
	db := model.GetDB()
	db.Model(user).Where("name = ?", name).First(&user)
	return
}

func RelatedLanguage2User(user *model.User) (languages []model.Language) {
	db := model.GetDB()
	db.Model(&user).Related(&languages, "Languages")
	return
}
