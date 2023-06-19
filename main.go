package main

import (
	"net/http"

	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
}

var student []Student

var err error
var db *gorm.DB

func init() {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

	db, err = gorm.Open(postgres.Open(os.Getenv("DB_DNS")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func main() {

	db.Raw("SELECT * FROM student").Scan(&student)
	router := gin.Default()
	router.GET("/student", getStudent)
	router.Run("localhost:3000")
}

func getStudent(s *gin.Context) {
	s.IndentedJSON(http.StatusOK, student)
}
