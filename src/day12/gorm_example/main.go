package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	UserId    int `gorm:"primary_key"` //指定主键
	UserName  string
	Age       uint8
	Sex       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"default:'null'"`
}

//指定表名
func (User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("user_id", 20)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.123.239:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql connect failed, err:", err)
	}
	db.LogMode(true)
	defer db.Close()
	fmt.Println("mysql connected success")

	//查询第一条数据
	//var user User
	//db.First(&user, 3)
	//fmt.Println("first data:", user.UserId)

	//新增数据
	//user := User{
	//	UserName:  "test_insert_data",
	//	Age:       10,
	//	Sex:       20,
	//}
	//db.Create(&user)
	//fmt.Println(user)

	//where查询
	var users []User
	db.Where("user_name = ?", "test_insert_data").Find(&users)
	fmt.Println(users)
	db.Where(&User{UserName: ""})
}
