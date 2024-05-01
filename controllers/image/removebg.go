package controllers

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"os"

	"github.com/IsaelVVI/goremovebg.git/config"
	"github.com/gin-gonic/gin"
)

func HandleRemovebg(ctx *gin.Context) {
	logger := config.GetLogger("images")
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Algo deu errado!",
		})
		logger.Errorf("Validation Error: %v", err.Error())
		return
	}

	// Abrir o arquivo
	uploadFile, err := file.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Opa aconteceu algo inesperado... tente novamente!"})
		logger.Errorf("Error in Open File: %v", err.Error())
		return
	}

	defer uploadFile.Close()

	// Ler o conteúdo do arquivo
	fileBytes, err := io.ReadAll(uploadFile)

	if err != nil {
		logger.Errorf("Error in Read File: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Opa aconteceu algo inesperado... tente novamente!"})
		return
	}

	logger.Debug("Success in Read File")

	newfile, err := removebg(fileBytes)

	logger.Debug("passou do newfile")

	if err != nil {
		logger.Errorf("Error in Read File: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert file to base64

	filebase64 := base64.StdEncoding.EncodeToString([]byte(newfile))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sucesso!",
		"file":    filebase64,
	})

	// ctx.Data(http.StatusOK, http.DetectContentType([]byte(newfile)), []byte(newfile))
}

func removebg(file []byte) (string, error) {

	// url request
	url := os.Getenv("URL") + "computervision/imageanalysis:segment?api-version=2023-02-01-preview&mode=backgroundRemoval"

	// body request
	body := file

	// Create a new request POST with request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		// error
		return "", err
	}

	apikey := os.Getenv("API_KEY")

	// Adiciona o cabeçalho de chave da API
	req.Header.Set("Ocp-Apim-Subscription-Key", apikey)

	// Define o cabeçalho Content-Type
	req.Header.Set("Content-Type", "application/octet-stream")

	// client http for call api request
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		// error
		return "", err
	}

	defer res.Body.Close()

	// read response
	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		// Tratar erro
		return "", err
	}

	return string(responseBody), nil

}
