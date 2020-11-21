package main

import (
	"link-note/backend/pkg"
	"time"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("FRONT_URL") == "" {
		err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			fmt.Println(err)
		}

		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
				"PUT",
				"DELETE",
			},
			AllowHeaders: []string{
				"Content-Type",
				"Content-Length",
				"Authorization",
				"Uid",
			},
			AllowOrigins: []string{
				"http://localhost:3000",
			},
			MaxAge: 24 * time.Hour,
		}))

		pkg.Serve(r, ":8080")
	} else {
		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
				"PUT",
				"DELETE",
			},
			AllowHeaders: []string{
				"Content-Type",
				"Content-Length",
				"cache-control",
				"user_id",
			},
			AllowOrigins: []string{
				os.Getenv("FRONT_URL"),
			},
			MaxAge: 24 * time.Hour,
		}))

		pkg.Serve(r, ":"+os.Getenv("PORT"))
	}

}