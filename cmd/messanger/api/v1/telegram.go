package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p messangerAPI) SendTelegram(c *gin.Context) {
	var message string
	if err := c.ShouldBind(&message); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.sendTelegram(message)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p messangerAPI) ScheduleSendTelegram(c *gin.Context) {

	if !config.Default.Messanger.Active {
		c.Error(errors.NewError(http.StatusInternalServerError, "Messanger inactive!"))
		return
	}

	var params []entities.JobParams

	if err := c.ShouldBind(&params); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	if len(params) != 1 {
		c.Error(errors.NewError(http.StatusBadRequest, "Invalid number of parameters. Please check your inputs"))
		return
	}

	err := p.sendTelegram(params[0].Value)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
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
		return fmt.Errorf("unknown status code: %s", res.Status)
	}
	return nil
}
