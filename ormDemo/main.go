package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var db *gorm.DB
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);size:255;default:'' "`
	Age int `gorm:"type:int;default:0"`
	Gender string `gorm:"type:varchar(100);default:'male';comment:'test'"`
}


func main() {
	//user  :=  User{ID:"123",Name:"xiaoxiao",Age:12,Gender:"female"}
	//exists := db.NewRecord(user)
	//if exists{
	//	log.Println("xxxx")
	//}
	var err error
	db,err := gorm.Open("mysql","root:@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)

	if !db.HasTable(&User{}) {
		if err := db.Set("user", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	user := User{}
	//user  :=  User{Name:"xiaoxiao",Age:12,Gender:"female"}
	errs := db.Model(&user).Update(&User{Name:"xxx"}).Where("age","12")
	log.Println(errs)
}