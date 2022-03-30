package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileUpload interface {
	SingleFile(*gin.Context)
	MultipleFile(*gin.Context)
}

func SingleFile(c *gin.Context) {
	file, err := c.FormFile("profile")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "files/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func MultipleFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["profile"]

	for _, file := range files {
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, "files/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", len(files)))
	}
}
