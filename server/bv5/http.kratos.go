package bv5

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	kratos "github.com/ory/kratos-client-go"
	"io"
	"net/http"
	"server/bardlog"
)

func GetKratosSessionM(b *BardView5, c *gin.Context) (*kratos.Session, error) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", b.Conf.KratosBaseUrl+"/sessions/whoami", nil)
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

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, session401
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

func GetKratosSession(b *BardView5Http) (*kratos.Session, error) {
	logger := b.Logger
	req, err := http.NewRequest("GET", b.BardView5.Conf.KratosBaseUrl+"/sessions/whoami", nil)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	req.Header.Add("Cookie", b.Context.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, session401
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
