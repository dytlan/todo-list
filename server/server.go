package server

import (
	"fmt"
	"github.com/dytlan/moonlay-todo-list/config"
	"github.com/dytlan/moonlay-todo-list/engine/file"
	"github.com/dytlan/moonlay-todo-list/engine/sql"
	"github.com/dytlan/moonlay-todo-list/engine/validation"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/todo"
	"github.com/dytlan/moonlay-todo-list/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Run() error {
	var configMap config.Map
	if err := viper.Unmarshal(&configMap); err != nil {
		return err
	}

	db, err := sql.NewClient(configMap.Database)
	if err != nil {
		return err
	}

	if err := sql.Migrate(db); err != nil {
		return err
	}

	accessor := accessor.NewAccessor(db)
	fileSvc := file.NewLocalStorageService(configMap.File)
	svc := todo.NewService(accessor, fileSvc)

	e := echo.New()
	e.Validator = &validation.Validator{Validator: validator.New()}
	e.Use(middleware.Recover())
	routes.Routing(e, svc)

	return e.Start(fmt.Sprintf(`:%s`, configMap.Server.Port))
}
