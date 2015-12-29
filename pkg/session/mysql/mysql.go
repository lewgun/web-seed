package mysql

import (
	. "github.com/gin-gonic/contrib/sessions"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
)

type MySQLStore interface {
	Store
}

const (
	tableName   = "session_store"
	defaultPath = "/"
)

//endPoint: the mysql's dsn.
//maxAge: session's max age.
//key: session's security key.
//"testuser:testpw@tcp(localhost:3306)/testdb?parseTime=true&loc=Local"
func NewMySQLStore(dsn string, maxAge int, key string) (MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStore(dsn, tableName, defaultPath, maxAge, []byte(key))
	if err != nil {
		return nil, err
	}
	return &mysqlStore{store}, nil
}

type mysqlStore struct {
	*mysqlstore.MySQLStore
}

func (c *mysqlStore) Options(options Options) {
	c.MySQLStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
