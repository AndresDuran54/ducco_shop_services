package lib

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

type Bind struct{}

func (o Bind) Bind(c echo.Context, i interface{}) error {
	data, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, i)

		if err != nil {
			return err
		}
	}

	defaultBuilder := echo.DefaultBinder{}
	err = defaultBuilder.BindPathParams(c, i)

	if err != nil {
		return nil
	}

	err = defaultBuilder.BindQueryParams(c, i)

	if err != nil {
		return err
	}

	err = defaultBuilder.BindHeaders(c, i)

	if err != nil {
		return err
	}

	return nil
}
