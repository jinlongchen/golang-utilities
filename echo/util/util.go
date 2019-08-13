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
