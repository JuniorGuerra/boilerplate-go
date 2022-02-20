package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JuniorGuerra/boilerplate-go/microservices/base/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// apiName := "api"

	router := gin.New()

	// Aqui van las rutas

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	gGroup := router.Group("v1/")

	for _, group := range routes.AppRouting {

		for _, route := range group.Routes {
			path := group.Prefix + "/" + route.Path

			switch route.Method {
			case http.MethodGet:
				gGroup.GET(path, route.Handler)
			case http.MethodPost:
				gGroup.POST(path, route.Handler)
			case http.MethodDelete:
				gGroup.DELETE(path, route.Handler)
			case http.MethodPatch:
				gGroup.PATCH(path, route.Handler)
			case http.MethodPut:
				gGroup.PUT(path, route.Handler)
			case http.MethodHead:
				gGroup.HEAD(path, route.Handler)
			case http.MethodOptions:
				gGroup.OPTIONS(path, route.Handler)
			default:
				gGroup.Any(path, route.Handler)
			}
			fmt.Println("path: ", path)
		}
	}

	// err := routes.Routers(router)
	// if err != nil {
	// 	panic(err)
	// }

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
