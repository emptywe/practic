package main

import (
	"FrameWorks/api/echoRouter"
	"FrameWorks/api/fastHttpRouter"
	"FrameWorks/api/fiberRouter"
	"FrameWorks/api/ginRouter"
	"FrameWorks/api/gorillaRouter"
	"FrameWorks/server/fastHttpServer"
	"FrameWorks/server/httpServer"
	"log"
	"os"
)

const (
	defaultHost = "localhost"
	defaultPort = "8080"
)

func main() {
	setup := os.Args
	if len(setup) != 2 {
		log.Println("choose framework(gin,echo,fiber,gorilla,fasthttp")
	}

	switch setup[1] {
	case "gin":
		handler := ginRouter.NewGinHandler()
		srv := new(httpServer.Server)
		log.Println("///*** Gin server start ***///")
		if err := srv.Run(defaultHost, defaultPort, handler.InitRoutes()); err != nil {
			log.Println(err)
			return
		}
	case "echo":
		handler := echoRouter.NewEchoHandler()
		srv := new(httpServer.Server)
		log.Println("///*** Echo server start ***///")
		if err := srv.Run(defaultHost, defaultPort, handler.InitRoutes()); err != nil {
			log.Println(err)
			return
		}
	case "fiber":
		handler := fiberRouter.NewFiberHandler()
		log.Println("///*** Fiber server start ***///")
		if err := handler.InitRoutesAndServer(defaultHost, defaultPort); err != nil {
			log.Println(err)
			return
		}
	case "gorilla":
		handler := gorillaRouter.NewGorillaHandler()
		srv := new(httpServer.Server)
		log.Println("///*** Gorilla server start ***///")
		if err := srv.Run(defaultHost, defaultPort, handler.InitRoutes()); err != nil {
			log.Println(err)
			return
		}
	case "fasthttp":
		handler := fastHttpRouter.NewFastHttpHandler()
		srv := new(fastHttpServer.Server)
		log.Println("///*** Fasthttp server start ***///")
		if err := srv.Run(defaultHost, defaultPort, handler.InitRoutes()); err != nil {
			log.Println(err)
			return
		}
	default:
		log.Println("unknown options")
	}
}
