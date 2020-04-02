package demo

import "ormDemo/model"

func ClearAssociation(user *model.User)  {
	db := model.GetDB()
	db.Model(&user).Association("Languages").Clear()
}
