package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/htnk128/go-ddd-sample/pkg/address/adapter/controller"
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

const (
	addressAPIRoot = "/addresses"
)

func initRouter() (*echo.Echo, error) {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	addressController, err := controller.NewAddressController()
	if err != nil {
		return nil, err
	}

	addressGroup := e.Group(addressAPIRoot)

	path := fmt.Sprintf("/:%s", controller.AddressIDParam)
	addressGroup.GET(path, addressController.Find())

	path = ""
	addressGroup.GET(path, addressController.FindAll())

	path = ""
	addressGroup.POST(path, addressController.Create())

	path = fmt.Sprintf("/:%s", controller.AddressIDParam)
	addressGroup.PUT(path, addressController.Update())

	path = fmt.Sprintf("/:%s", controller.AddressIDParam)
	addressGroup.DELETE(path, addressController.Delete())

	return e, nil
}

func NewHttpServer() *http.Server {
	config, err := newHTTPConfig()
	if err != nil {
		panic(err)
	}

	router, err := initRouter()
	if err != nil {
		panic(err)
	}

	return &http.Server{
		Addr:         config.addr,
		Handler:      router,
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
	}
}
