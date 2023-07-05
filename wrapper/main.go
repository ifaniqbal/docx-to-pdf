package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		fileContent := bytes.Buffer{}
		_, err = io.Copy(&fileContent, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
			return
		}

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form file"})
			return
		}

		_, err = io.Copy(part, &fileContent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file content"})
			return
		}

		type DataPayload struct {
			Nama string `json:"nama"`
		}
		data := DataPayload{
			Nama: "Iqbal",
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data payload"})
			return
		}

		dataField, err := writer.CreateFormField("data")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form field"})
			return
		}

		_, err = io.Copy(dataField, bytes.NewReader(dataBytes))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write data payload"})
			return
		}

		err = writer.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close multipart writer"})
			return
		}

		req, err := http.NewRequest("POST", "http://docx-renderer/render", body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create HTTP request"})
			return
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send HTTP request"})
			return
		}
		defer resp.Body.Close()

		// Handle the response from the server

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded and processed successfully"})
	})

	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
