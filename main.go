

package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type Handler struct {
	db map[string]*User
}

func (h *Handler) CreateUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    validator := validator.New()
    if err := validator.Struct(u); err != nil {
		fmt.Println("error")
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    h.db[u.Email] = u

    return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	h := &Handler{
		db: make(map[string]*User),
	}
	e.POST("/users", h.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
