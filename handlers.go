package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type API struct {
	host string
	server *echo.Echo
	pg *Postgres
}

func NewAPI(host string, pgHost string) (*API, error) {
	pg, err := NewPostgres(pgHost)
	if err != nil {
		return nil, err
	}
	
	return &API{
		host: host,
		server: echo.New(),
		pg: pg,
	}, nil
}

func (api *API) Start() error {
	api.server.POST("/transactions", api.CreateTransaction)	
	return api.server.Start(api.host)
}

func (api *API) CreateTransaction(ctx echo.Context) error {
	var txn Transaction
	if err := ctx.Bind(&txn); err != nil {
		fmt.Println(err)
		return err
	}

	reqctx := ctx.Request().Context()
	_, err := api.pg.CreateTransaction(reqctx, txn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
