package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// `User` 属于 `Company`，`CompanyID` 是外键
// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mysql~!@#$%^&*()_+@tcp(117.50.184.138:37210)/zyh_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&CreditCard{})

	// 在大型的系统中 不建议使用外键约束 外键约束也有很大的优点:数据的完整性
	/**
	外键约束会让数据很完整，即便是业务代码有人考虑的不严谨
	在大型的系统，高并发系统一般不使用外键约束，自已在业务层面保持数据的一致性
	*/

	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number:    "12",
	//	UserRefer: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number:    "34",
	//	UserRefer: user.ID,
	//})

	var user User
	db.Preload("CreditCards").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card.Number)
	}

}
