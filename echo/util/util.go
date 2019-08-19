package util

import (
	"errors"
	httpUtil "github.com/jinlongchen/golang-utilities/http"
	"github.com/jinlongchen/golang-utilities/json"
	"github.com/labstack/echo/v4"
)

func ParseJSON(ctx echo.Context, v interface{}) error {
	data := httpUtil.GetRequestBody(ctx.Request())
	if data == nil {
		return errors.New("bad request")
	}
	return json.Unmarshal(data, v)
}

func WriteJSON(ctx echo.Context, code int, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ctx.Blob(code, echo.MIMEApplicationJSON, data)
}
