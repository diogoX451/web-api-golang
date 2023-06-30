package main

// https://github.com/irahardianto/service-pattern-go/blob/master/controllers/PlayerController_test.go

import (
	"io/ioutil"
	"log"
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

var filepath string
var url string
var err error
var db *gorm.DB

func init() {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

	connectDB()

}

func connectDB() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_DNS")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func main() {
	filepath = "embargo"
	url = "https://ppc-library.s3.amazonaws.com/docs/embargoes/16080673365050692_rel_areas_embargadas_0-65000_2020-12-10_080019.csv"
	router := gin.Default()
	router.GET("/student", getStudent)
	router.Run("localhost:3000")
}

func getStudent(s *gin.Context) {
	var student []Student
	err := db.Raw("SELECT * FROM product").Scan(&student).Error
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.IndentedJSON(http.StatusOK, student)
}

func DownloadFile() {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	file, err := ioutil.TempFile("", filepath)
	os.Mkdir(file.Name(), 0777)

	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(file.Name())

	filepath = file.Name()

}
