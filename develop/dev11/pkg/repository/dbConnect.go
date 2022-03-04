package repository

import (
	"database/sql"
	"fmt"
)

type dBConnect struct {
	*sql.DB
}

func NewDbConnect() (*dBConnect, error) {
	db := &dBConnect{}
	err := db.openConnect()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (d *dBConnect) openConnect() error {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	d.DB, err = sql.Open("postgres", psqlconn)
	return err
}

func (d *dBConnect) CloseConnect() {
	d.DB.Close()
}
