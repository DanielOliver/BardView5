package bv5

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"server/bardlog"
)

func (b *BardView5) GetWhoAmI(c *gin.Context) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", "http://proxy.local/sessions/whoami", nil)
	if err != nil {
		logger.Err(err)
		c.Status(500)
		return
	}
	req.Header.Add("Cookie", c.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err)
		c.Status(500)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Err(err)
		c.Status(500)
		return
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err)
		c.Status(500)
		return
	}
	logger.Info().Interface("response", result).Msg("Response!")
	c.JSON(resp.StatusCode, result)
}
