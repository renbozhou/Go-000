package main

import (
	"database/sql"
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

func GetUserApi() (string, error) {
	_, err := GetUserService()
	if err != nil {
		return "fail", xerrors.Wrap(errors.New("servic not find user"), "user api")
	}
	return "ok", nil

}
func GetUserService() (string, error) {
	return "", xerrors.Wrap(errors.New("not find user"), "user service")
}

func GetUserDAO() error {
	return sql.ErrNoRows
}

func main() {
	_, err := GetUserApi()
	if err != nil {
		fmt.Printf("Original error: %T %v \n", xerrors.Cause(err), xerrors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
	}
}
