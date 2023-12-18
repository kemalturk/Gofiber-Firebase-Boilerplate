package app

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/GofiberFirebaseBoilerplate/src"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func init() {
	app = src.CreateServer()
}

// Start start Fiber app with normal interface
func Start(addr string) error {
	if -1 == strings.IndexByte(addr, ':') {
		addr = ":" + addr
	}

	return app.Listen(addr)
}

// MyCloudFunction Exported http.HandlerFunc to be deployed to as a Cloud Function
func MyCloudFunction(w http.ResponseWriter, r *http.Request) {
	err := CloudFunctionRouteToFiber(app, w, r)
	if err != nil {
		fmt.Fprintf(w, "err : %v", err)
		return
	}
}
