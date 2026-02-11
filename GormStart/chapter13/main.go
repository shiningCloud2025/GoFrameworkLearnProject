package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Language struct {
	//gorm.Model
	Name    string
	AddTime time.Time // 我希望每个记录创建的时候自动加入当前时间到AddTime中
}

// 在gorm中可以通过给某一个struct添加TableName方法来自定义表名
func (Language) TableName() string {
	return "my_languages"
}

func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return
}

/*
*
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
	// NamingStrategy这个配置和TableName不能同时配置 如果同时的话以TableName为主
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshoop",
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "Go",
	})

}
