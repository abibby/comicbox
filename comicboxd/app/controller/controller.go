package controller

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/zwzn/comicbox/comicboxd/app/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	vars           map[string]string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	DB             *sqlx.DB
	User           *model.User
	Session        *sessions.Session
}

func (c Context) Var(key string) string {
	if c.vars == nil {
		c.vars = mux.Vars(c.Request)
	}
	return c.vars[key]
}

func (c Context) DecodeBody(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

func (c Context) SSet(key, value interface{}) {
	c.Session.Values[key] = value
}

func (c Context) SGet(key interface{}) (interface{}, bool) {
	value, ok := c.Session.Values[key]
	return value, ok
}

func (c Context) SGetString(key interface{}) (string, bool) {
	value, ok := c.SGet(key)
	strVal, ok := value.(string)
	return strVal, ok
}

func (c Context) SGetInt64(key interface{}) (int64, bool) {
	value, ok := c.SGet(key)
	strVal, ok := value.(int64)
	return strVal, ok
}