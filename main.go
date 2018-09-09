package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port     = flag.Int("port", 8080, "port number")
	logfile  = flag.String("log", "", "log file path")
	paramKey = flag.String("param", "i", "param key")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime)
	if fileName := *logfile; fileName != "" {
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("failed to open logfile: %v", err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	e := echo.New()
	e.HideBanner = true
	e.Debug = false
	e.Logger.SetLevel(99)
	e.Use(middleware.Recover())

	e.POST("/collect", collect)
	e.Static("/", "static")

	if err := e.Start(fmt.Sprintf(":%d", *port)); err != nil {
		log.Printf("stop server: %v", err)
	}
}

func collect(c echo.Context) error {
	if c.Response().Size > 2048 {
		return c.String(http.StatusBadRequest, "")
	}
	clientInfo := map[string]interface{}{}
	clientJSON := c.FormValue(*paramKey)
	if err := json.Unmarshal([]byte(clientJSON), &clientInfo); err != nil {
		log.Printf("failed to unmalshal client: %v", err)
		return err
	}
	req := c.Request()
	logdata := map[string]interface{}{
		"clinet":     clientInfo,
		"header":     convertHeader(c.Request().Header),
		"host":       req.Host,
		"referer":    req.Referer(),
		"remoteAddr": req.RemoteAddr,
	}
	b, err := json.Marshal(logdata)
	if err != nil {
		log.Printf("failed to marshal logdata: %v", err)
		return err
	}
	log.Print(string(b))
	return c.String(http.StatusOK, "")
}

func convertHeader(header http.Header) map[string]string {
	h := map[string]string{}
	for key, value := range header {
		if len(value) > 0 {
			h[key] = value[0]
		}
	}
	return h
}
