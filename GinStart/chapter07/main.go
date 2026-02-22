package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type SignUpForm struct {
	Age        uint8  ` json:"age" xml:"age"  binding:"gte=1,lte=130" `
	Name       string `json:"name" xml:"name" form:"name"  binding:"required,min=3"`
	Email      string ` json:"email" xml:"email" binding:"required,email"`
	Password   string ` json:"password" xml:"password"  binding:"required"`
	RePassword string ` json:"rePassword" xml:"rePassword"  binding:"required,eqfield=Password"` // 跨字段
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

var trans ut.Translator

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm

		if err := c.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})

			}
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": removeTopStruct(errs.Translate(trans))})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})

	router.POST("/signUp", func(c *gin.Context) {
		var signUpForm SignUpForm
		if err := c.ShouldBind(&signUpForm); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	router.Run(":8083")

}
