package main

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"server/api"
	"server/db"
)

type Generators struct {
	userNode *snowflake.Node
}

type BardView5 struct {
	db         *sql.DB
	querier    db.Querier
	generators *Generators
}

func NewBardView5() (bv5 *BardView5, err error) {
	connectionString := viper.GetString("CONNECTION")
	if connectionString == "" {
		return nil, fmt.Errorf("expected bardview5 sql connection string")
	}
	pgConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open bardview5 connection string")
	}

	userNode, err := snowflake.NewNode(1)

	return &BardView5{
		db:      pgConnection,
		querier: db.New(pgConnection),
		generators: &Generators{
			userNode: userNode,
		},
	}, nil
}

func (b *BardView5) DB() *sql.DB {
	return b.db
}

func (b *BardView5) Querier() db.Querier {
	return b.querier
}

func (b *BardView5) CreateNewUser(context *gin.Context) {
	body := api.UserPost{}
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rowsInserted, err := b.querier.UserInsert(context, db.UserInsertParams{
		UserID:    b.generators.userNode.Generate().Int64(),
		Uuid:      uuid.New(),
		Name:      body.Name,
		Email:     string(body.Email),
		Tags:      []string{},
		CreatedBy: sql.NullInt64{},
	})
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	if rowsInserted == 0 {
		context.AbortWithStatus(http.StatusNotModified)
		return
	}
	context.Status(http.StatusOK)
}

type TGetUserAcl struct {
	UserId  int64  `json:"user_id" uri:"user_id"`
	Subject string `json:"subject" uri:"subject"`
}

func (b *BardView5) GetUserAcl(context *gin.Context) {
	uri := &TGetUserAcl{}
	if err := context.BindUri(&uri); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	acls, err := b.Querier().GetAcl(context, db.GetAclParams{
		Subject:   uri.Subject,
		SubjectID: sql.NullInt64{},
		UserID:    uri.UserId,
	})
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, acls)
}