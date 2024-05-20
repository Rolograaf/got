package handler

import (
	"github.com/rolograaf/got/view/user"
	"github.com/rolograaf/got/handler/util"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
