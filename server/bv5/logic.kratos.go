package bv5

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
	"io"
	"net/http"
	"server/bardlog"
)

type SessionError struct {
	msg string
}

func (s *SessionError) Error() string {
	return s.msg
}

var (
	session401 = &SessionError{
		msg: "401",
	}
)

func (b *BardView5) getKratosSession(c *gin.Context) (*kratos.Session, error) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", b.conf.kratosBaseUrl+"/sessions/whoami", nil)
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

func (b *BardView5) AddSessionToContext(c *gin.Context) {
	session, err := b.getKratosSession(c)
	if err != nil {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	userUuid := uuid.MustParse(session.Identity.Id)
	users, err := b.querier.UserFindByUuid(c, userUuid)
	if err != nil || len(users) == 0 || len(users) > 1 {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	c.Set(Session, MakeSession(users[0].UserID))
}

func (b *BardView5) RequireValidSession(c *gin.Context) {
	s := SessionCriteria(c)
	if s.Anonymous {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
