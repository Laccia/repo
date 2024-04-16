package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"repo/cmd/base"
	"repo/cmd/delivery"
	"repo/cmd/jwt"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func App() {
	key := os.Getenv("KEYJWT")
	log.Info(key)
	base := base.New()
	delivery := delivery.New(*base)
	jwtBase := jwt.New(base.UsersBase, key)
	jwt.JWTAutoMiddleware(key)
	//Инициализация объекта сервера
	server := echo.New()
	//Установка функций логирования, перехвата ошибок и меткир
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	//Установка уровня логирования
	server.Logger.SetLevel(log.DEBUG)
	//Установка обработки ендпоинтов
	server.POST("/new_user", delivery.NewUser) //Создание пользователя
	server.POST("/login", jwtBase.Login)       //Вход пользователя

	//Ендпоинты для обычного пользователя
	tokenGroup := server.Group("/token")
	tokenGroup.Use(jwt.JWTAutoMiddleware(key)) //Установка проверки токена подключения
	userGroup := tokenGroup.Group("/shop")
	userGroup.GET("/list", delivery.List) //Вывод всех товаров
	//Эндпоинты для администратора
	adminGroup := tokenGroup.Group("/admin")

	adminGroup.POST("/create", delivery.NewItems)          //Создание товара
	adminGroup.DELETE("/delete/:id", delivery.DeleteItems) //Удаление товара по ID

	Init(server)

}

func Init(server *echo.Echo) {
	go func() {
		if err := server.Start(":2000"); err != nil && errors.Is(err, http.ErrServerClosed) {
			server.Logger.Fatal(err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	server.Logger.Info("shutdown inited")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
