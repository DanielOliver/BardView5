package bv5

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	kratos "github.com/ory/kratos-client-go"
	"io"
	"net/http"
	"server/bardlog"
)

const (
	kratosRelativeWhoAmI = "/sessions/whoami"
)

type HttpKratos struct {
	KratosBaseUrl string
}

type KratosDep interface {
	GetKratosSessionM(c *gin.Context) (*kratos.Session, error)
	GetKratosSession(b *BardView5Http) (*kratos.Session, error)
}

func (h *HttpKratos) GetKratosSessionM(c *gin.Context) (*kratos.Session, error) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", h.KratosBaseUrl+kratosRelativeWhoAmI, nil)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepInit(kratosRelativeWhoAmI, DepHttpKratos)
	}

	req.Header.Add("Cookie", c.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, session401
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}

	var result kratos.Session
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}
	return &result, nil
}

func (h *HttpKratos) GetKratosSession(b *BardView5Http) (*kratos.Session, error) {
	logger := b.Logger
	req, err := http.NewRequest("GET", h.KratosBaseUrl+kratosRelativeWhoAmI, nil)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepInit(kratosRelativeWhoAmI, DepHttpKratos)
	}

	req.Header.Add("Cookie", b.Context.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, session401
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}

	var result kratos.Session
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err).Msg(kratosRelativeWhoAmI)
		return nil, ErrHttpDepUnknown(kratosRelativeWhoAmI, DepHttpKratos)
	}
	return &result, nil
}
