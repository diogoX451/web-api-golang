package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID       string `json:"id"`
	NAME     string `json:"NAME"`
	LASTNAME string `json:"LASTNAME"`
}

var student = []Student{
	{ID: "1", NAME: "Diogo", LASTNAME: "Almeida"},
	{ID: "2", NAME: "Guiller", LASTNAME: "jug"},
	{ID: "3", NAME: "NAME3", LASTNAME: "LASTNAME3"},
	{ID: "4", NAME: "NAME4", LASTNAME: "LASTNAME4"},
	{ID: "5", NAME: "NAME5", LASTNAME: "LASTNAME5"},
	{ID: "6", NAME: "NAME6", LASTNAME: "LASTNAME6"},
	{ID: "7", NAME: "NAME7", LASTNAME: "LASTNAME7"},
	{ID: "8", NAME: "NAME8", LASTNAME: "LASTNAME8"},
	{ID: "9", NAME: "NAME9", LASTNAME: "LASTNAME9"},
	{ID: "10", NAME: "NAME10", LASTNAME: "LASTNAME10"},
}

func main() {
	router := gin.Default()
	router.GET("/", getStudent)
	router.Run("localhost:3000")
}

func getStudent(s *gin.Context) {
	s.IndentedJSON(http.StatusOK, student)
}
