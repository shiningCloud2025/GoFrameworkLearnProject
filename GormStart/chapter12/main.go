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

// User 拥有并属于多种 language，`user_languages` 是连接表
type User1 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}
/**
自定义表名:
1.我们自已定义表名是什么
2.统一的给所有的表名加一个前缀
 */
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

	//db.AutoMigrate(&User1{})

	//languages := []Language{}
	//languages = append(languages, Language{Name: "go"})
	//languages = append(languages, Language{Name: "python"})
	//user := User1{
	//	Languages: languages,
	//}
	//db.Create(&user)

	var user User1
	db.Preload("Languages").First(&user)
	for _, lang := range user.Languages {
		fmt.Println(lang.Name)
	}

	// 如果我已经取出一个用户来了，但是这个用户我们之前没有使用preload来加载对应的Languages
	// 不是说用户有languages就一定要取出
	var user1 User1
	db.First(&user1)
	var languages []Language
	db.Model(&user1).Association("Languages").Find(&languages)
	for _, lang := range user1.Languages {
		fmt.Println(lang.Name)
	}
}
