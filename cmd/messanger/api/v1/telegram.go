package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p messangerAPI) SendTelegram(c *gin.Context) {
	var message string
	if err := c.ShouldBind(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "MESSANGERERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.sendTelegram(message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "MESSANGERERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p messangerAPI) ScheduleSendTelegram(c *gin.Context) {

	if !config.Default.Messanger.Active {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "MESSANGERERR",
				"message": "Messanger inactive!"})
		return
	}

	var params []entities.JobParams

	if err := c.ShouldBind(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "MESSANGERERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	if len(params) != 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "MESSANGERERR",
				"message": "Invalid number of parameters. Please check your inputs"})
		return
	}

	err := p.sendTelegram(params[0].Value)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "MESSANGERERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p messangerAPI) sendTelegram(message string) error {

	if !config.Default.Messanger.Telegram.Active {
		return nil
	}

	type sendMessageReqBody struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}

	reqBody := &sendMessageReqBody{
		ChatID: config.Default.Messanger.Telegram.ChatId,
		Text:   message,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	url := config.Default.Messanger.Telegram.Url + config.Default.Messanger.Telegram.Key + "/sendMessage"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("POST error: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("Unexpected status" + res.Status)
		return fmt.Errorf("unknown error: %w", err)
	}
	return nil
}
