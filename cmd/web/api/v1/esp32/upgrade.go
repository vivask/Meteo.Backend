package esp32

import (
	"bytes"
	"fmt"
	"io"
	"meteo/internal/kit"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const UPLOAD = "./firmware"

func (p esp32API) UpgradeEsp32(c *gin.Context) {

	const name = "firmware"

	file, err := c.FormFile(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	// Upload the file to specific dst.
	dst := fmt.Sprintf("%s/%s", UPLOAD, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	content, buf, err := createForm(name, dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	_, err = kit.PostFormInt("/esp32/upload", content, buf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	if err := p.repo.UpgradeEsp32(file.Filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p esp32API) GetUpgradeStatus(c *gin.Context) {
	status, err := p.repo.GetUpgradeStatus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": status})
}

func (p esp32API) TerminateUpgrade(c *gin.Context) {
	if err := p.repo.TerminateUpgrade(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func createForm(name, filename string) (string, *bytes.Buffer, error) {
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile(name, filename)
	if err != nil {
		return "", nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()
	_, err = io.Copy(fw, file)
	if err != nil {
		return "", nil, err
	}
	writer.Close()
	return writer.FormDataContentType(), body, nil
}
