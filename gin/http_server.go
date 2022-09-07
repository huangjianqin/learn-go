package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type Person struct {
	Name string `form:"name" binding:"required,nameValidator"`
	Age  int    `form:"age" binding:"required"`
}

func main() {
	router := gin.Default()

	//通过tag注册validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("nameValidator", newNameValidator())
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//:name -> path参数匹配
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//*action模糊匹配, 即可有可无
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Printf("%v\n", c.FullPath())
		c.String(http.StatusOK, message)
	})

	//优先匹配固定path, 模糊path后匹配
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	//query
	router.GET("/query", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		//相当于 c.Request.URL.Query().Get("lastname")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	//表单
	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		names := c.PostFormMap("names")
		message := c.PostForm("message")

		fmt.Printf("name: %v; names: %v; message: %v\n", name, names, message)
	})

	//bind
	router.GET("/bind", func(c *gin.Context) {
		var person Person
		err := c.ShouldBindQuery(&person)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			fmt.Printf("person: %v\n", person)
		}
	})

	var middlewareFunc gin.HandlerFunc = func(c *gin.Context) {
		fmt.Printf("custom middleware `%v`\n", c.FullPath())
	}

	//定义path group
	v1Router := router.Group("/v1")
	//打印logger的middleware
	v1Router.Use(gin.Logger())
	v1Router.Use(middlewareFunc)

	v1Router.GET("/query", func(c *gin.Context) {
		c.String(http.StatusOK, "response")
	})

	err := router.Run(":8888")
	if err != nil {
		log.Fatalf("fatal error %v\n", err)
	}
}

func newNameValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().Interface().(string)
		if s == "abc" {
			return true
		}
		return false
	}
}
