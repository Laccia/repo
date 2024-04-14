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
	"github.com/prometheus/client_golang/prometheus/promhttp"

	echoPrometheus "github.com/globocom/echo-prometheus"
)

func App() {
	key := os.Getenv("KEYJWT")
	log.Info(key)
	base := base.New()
	delivery := delivery.New(*base)
	jwtBase := jwt.New(*&base.UsersBase, key)
	jwt.JWTAutoMiddleware(key)
	//Инициализация объекта сервера
	server := echo.New()
	//Установка функций логирования, перехвата ошибок и меткир
	server.Use(echoPrometheus.MetricsMiddleware())
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	//Установка уровня логирования
	server.Logger.SetLevel(log.DEBUG)
	//Установка обработки ендпоинтов
	server.POST("/new_user", delivery.NewUser) //Создание пользователя
	server.POST("/login", nil)                 //Вход пользователя
	//Эндпоинты для обычного пользователя
	userGroup := server.Group("/shop")
	userGroup.GET("/list", delivery.List) //Вывод всех товаров
	userGroup.GET("/serach/:name", nil)   //Вывод с фильтром по ключевому значению
	//Эндпоинты для администратора
	adminGroup := server.Group("/admin")
	adminGroup.Use(nil) //Установка проверки токена подключения

	tokenGroup := server.Group("/token")
	tokenGroup.Use(nil)                                              //Установка проверки токена подключения
	adminGroup.POST("/create", delivery.NewItems)                    //Создание товара
	adminGroup.DELETE("/delete/:id", delivery.DeleteItems)           //Удаление товара по ID
	adminGroup.GET("/metrics", echo.WrapHandler(promhttp.Handler())) //Метрики сервера

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
