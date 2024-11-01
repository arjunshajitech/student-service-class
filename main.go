package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.GET("/student/me", me)
	router.GET("/student/notes", notes)
	router.GET("/student/note/:subject", notesBySubject)

	fmt.Println("Teacher Service Started on port 4000")
	http.ListenAndServe(":4000", router)
}
