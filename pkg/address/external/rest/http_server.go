package rest

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type httpConfig struct {
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func newHTTPConfig() (*httpConfig, error) {
	host, _ := os.LookupEnv("HTTP_HOST")

	port, b := os.LookupEnv("HTTP_PORT")
	if !b {
		port = "8081"
	}

	rt, b := os.LookupEnv("HTTP_READ_TIMEOUT")
	if !b {
		rt = "5000"
	}

	readTimeout, err := strconv.Atoi(rt)
	if err != nil {
		return nil, errors.Wrap(err, "env of HTTP_READ_TIMEOUT is not numeric.")
	}

	wt, b := os.LookupEnv("HTTP_WRITE_TIMEOUT")
	if !b {
		wt = "5000"
	}

	writeTimeout, err := strconv.Atoi(wt)
	if err != nil {
		return nil, errors.Wrap(err, "env of HTTP_WRITE_TIMEOUT is not numeric.")
	}

	return &httpConfig{
		addr:         host + ":" + port,
		readTimeout:  time.Duration(readTimeout) * time.Second,
		writeTimeout: time.Duration(writeTimeout) * time.Second,
	}, nil
}

func initRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}

func NewHttpServer() *http.Server {
	config, err := newHTTPConfig()
	if err != nil {
		panic(err)
	}

	router := initRouter()

	return &http.Server{
		Addr:         config.addr,
		Handler:      router,
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
	}
}
