package bv5

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	kratos "github.com/ory/kratos-client-go"
	"io"
	"net/http"
	"server/bardlog"
)

func (b *BardView5) getKratosSession(c *gin.Context) (*kratos.Session, error) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", "http://proxy.local/sessions/whoami", nil)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	req.Header.Add("Cookie", c.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	var result kratos.Session
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err)
		return nil, err
	}
	return &result, nil
}
