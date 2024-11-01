package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Note    string `json:"note"`
}

func me(c *gin.Context) {
	response := gin.H{
		"name":    "Arjun",
		"company": "Techgentsia",
		"role":    "Backend Engineer",
	}
	c.IndentedJSON(http.StatusOK, response)
}

func notes(c *gin.Context) {

	url := "http://localhost:3000/teacher/notes"
	response := callingExternalAPI(url)

	var notes *[]Note
	err := json.Unmarshal([]byte(response), &notes)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, notes)
}

func notesBySubject(c *gin.Context) {
	subject := c.Param("subject")
	fmt.Println(subject)
	url := "http://localhost:3000/teacher/note/" + subject
	fmt.Println(url)
	response := callingExternalAPI(url)

	var notes *Note
	err := json.Unmarshal([]byte(response), &notes)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, notes)
}

func callingExternalAPI(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	return string(body)
}
