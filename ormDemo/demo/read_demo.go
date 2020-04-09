package demo

import "ormDemo/model"

func GetAllLanguages() (languages []model.Language) {
	db := model.GetDB()
	if err := db.Model(model.Language{}).Find(&languages); err != nil {
		return
	}
	return
}

func GetUser(maps interface{}) (user model.User) {
	db := model.GetDB()
	db.Where(maps).First(&user)
	return
}

func RelatedLanguage2User(user *model.User) (languages []model.Language) {
	db := model.GetDB()
	db.Model(&user).Related(&languages, "Languages")
	return
}

func RelatedLanguage2UserFirst(user *model.User) (language model.Language) {
	db := model.GetDB()
	db.Model(&user).Related(&model.Language{}, "Languages").First(&language)
	return
}
