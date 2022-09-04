package main

import (
	"bytes"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"math/rand"
	"time"
)

//默认创建表明为{struct名小写}+s
type User struct {
	//可不加, 那就要自己定义主键了
	gorm.Model
	Name  string
	Age   uint
	Email string
	//单词以_分割
	LoginTime time.Time
	//外键定义
	CompanyId int
	Company   Company `gorm:"foreignKey:CompanyId"`
	//1对1
	CreditCard CreditCard
	//多对多
	Languages []Language `gorm:"many2many:user_languages"`
}

type Company struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type CreditCard struct {
	ID   int `gorm:"primarykey"`
	Name string
	//1对1, 以实体名+Id则会自动创建外键, 不然得自己定义外键
	UserId uint
}

type Language struct {
	ID   int `gorm:"primarykey"`
	Name string
}

func (u *User) BeforeSave(db *gorm.DB) error {
	println(u.Name + " before save")
	return nil
}

func (u *User) BeforeDelete(db *gorm.DB) error {
	println(u.Name + " before delete")
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	println(u.Name + " before update")
	return nil
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	println(u.Name + " before create")
	return nil
}

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/kin_demo?charset=utf8&parseTime=True"), &gorm.Config{
		//打印sql
		Logger: logger.Default.LogMode(logger.Info),
		//生成sql, 但不执行, 用于测试
		//DryRun: true,
		//...
	})
	if err != nil {
		panic(any(errors.New("failed to connect database")))
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Company{}, &CreditCard{}, &Language{})

	//真随机
	rand.Seed(time.Now().UnixNano())

	// Create
	charBs := []byte("qwertyuiopasdfghjklzxcvbnm1234567890")
	nameBuf := bytes.Buffer{}
	for i := 1; i < 10; i++ {
		nameBuf.WriteByte(charBs[rand.Intn(cap(charBs))])
	}
	emailBuf := bytes.Buffer{}
	for i := 1; i < 20; i++ {
		if i == 10 {
			emailBuf.WriteString("@")
		} else {
			emailBuf.WriteByte(charBs[rand.Intn(cap(charBs))])
		}
	}
	emailBuf.WriteString(".com")

	cNameBuf := bytes.Buffer{}
	for i := 1; i < 10; i++ {
		cNameBuf.WriteByte(charBs[rand.Intn(cap(charBs))])
	}

	db.Create(&User{Name: nameBuf.String(), Age: uint(rand.Intn(100)), Email: emailBuf.String(), LoginTime: time.Now(),
		Company:    Company{Name: cNameBuf.String()},
		CreditCard: CreditCard{Name: nameBuf.String() + "-card"},
		Languages:  []Language{{Name: "C"}, {Name: "Java"}, {Name: "Go"}},
	})

	// Read
	var user User
	////按主键升序
	//不使用Preload, 那么实体中关联的其他实体并不会初始化, 只有外键字段值可以读取
	//Preload("{struct定义的字段名}")
	//db.Preload("Company").Preload("CreditCard").Preload("Languages").First(&user)
	//预加载全部
	db.Preload(clause.Associations).First(&user)
	fmt.Println(user)
	//db.First(&user, "Age < ?", "50")
	//fmt.Println(user)

	//表关联计数
	//println(db.Model(&user).Association("Languages").Count())

	//var users []User
	////select id int(1, 2, 3)
	//db.Find(&users, []int{1, 2, 3})
	//for _, user := range users {
	//	fmt.Printf("%v\n", user)
	//}
	//println("----")
	//
	//db.Find(&users)
	//for _, user := range users {
	//	fmt.Printf("%v\n", user)
	//}
	//println("----")
	//
	//var ages []int
	//db.Model(&User{}).Select("age").Find(&ages)
	//for _, age := range ages {
	//	fmt.Printf("%v\n", age)
	//}
	//println("----")

	//update
	//Save和Update方法执行sql同时会更新updated_at字段, 如果不想更新这个字段, 则需要调用db.UpdateColumn
	//user.Name = "first"
	//user.Age = 2
	////根据主键update set所有字段
	//db.Save(user)

	//拿出指定user, 并更新age
	//db.Model(&user).Where("age = ?", 30).Update("Age", 40)

	//delete
	//非真正删除, 只是设置了deleted_at字段, 想要真正删除, 则需db.Exec
	//满足条件则删除user
	//db.Delete(&user, "name = ?", "abs")
	//批量删除
	//db.Delete(&User{}, "age = ?", 41)

	//原生sql
	//var results []User
	//db.Raw("SELECT id, name, age FROM users").Scan(&results)
	//for _, r := range results {
	//	fmt.Printf("%v\n", r)
	//}

	//db.Exec("UPDATE users set name = 'second' WHERE id = 3")

	//rows, _ := db.Table("users").Where("age < ?", 50).Select("name").Rows()
	//for rows.Next() {
	//	var name string
	//	err := rows.Scan(&name)
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	} else{
	//		fmt.Printf("%v\n", name)
	//	}
	//}

	//带select的delete, 会把相应字段(外键)的关联(many2many2表)或者数据也delete
	//db.Select("Company", "CreditCard", "Languages").Delete(&user)

	//事务
	//1
	//tx := db.Begin()
	//err = tx.Create(&user).Error
	//if err != nil{
	//	tx.Rollback()
	//} else{
	//	tx.Commit()
	//}

	//2
	//db.Transaction(func(tx *gorm.DB) error {
	//	err = tx.Create(&user).Error
	//	if err != nil{
	//		return err
	//	}
	//	return nil
	//})
}
