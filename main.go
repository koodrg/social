package main

import (
	"log"
	"net/http"
	"os"

	"social/component/appctx"
	_ "social/docs"
	"social/middleware"

	// "social/middlewares"
	"social/module/post/transport/ginrestaurant"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//----------------------------------------------------------------
	// DB Connection
	dsn := os.Getenv("CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected success!  ", db)
	//----------------------------------------------------------------

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	router := gin.Default()

	//----------------------------------------------------------------
	// Middleware
	router.Use(middleware.Recover(appContext))
	//----------------------------------------------------------------

	v1 := router.Group("api/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Register successful",
		})
	})
	auth.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	})

	user := v1.Group("/user")
	user.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get profile successful",
		})
	})
	user.PUT("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update profile successful",
		})
	})

	// post
	posts := v1.Group("/posts")
	posts.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get successful",
		})
	})
	posts.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get successful",
		})
	})
	posts.POST("", ginrestaurant.CreateRestaurant(db))
	posts.PUT("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update successful",
		})
	})
	posts.DELETE("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update successful",
		})
	})

	//image
	image := v1.Group("/images")
	image.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "create image successful",
		})
	})
	image.DELETE("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete image successful",
		})
	})

	router.Run()
}
