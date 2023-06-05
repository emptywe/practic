package fastHttpServer

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	fastHttpServer *fasthttp.Server
}

func (s *Server) Run(host, port string, handler fasthttp.RequestHandler) error {
	s.fastHttpServer = &fasthttp.Server{
		Handler:                            handler,
		ErrorHandler:                       nil,
		HeaderReceived:                     nil,
		ContinueHandler:                    nil,
		Name:                               "",
		Concurrency:                        0,
		ReadBufferSize:                     0,
		WriteBufferSize:                    0,
		ReadTimeout:                        10 * time.Second,
		WriteTimeout:                       10 * time.Second,
		IdleTimeout:                        0,
		MaxConnsPerIP:                      0,
		MaxRequestsPerConn:                 0,
		MaxKeepaliveDuration:               0,
		MaxIdleWorkerDuration:              0,
		TCPKeepalivePeriod:                 0,
		MaxRequestBodySize:                 0,
		DisableKeepalive:                   false,
		TCPKeepalive:                       false,
		ReduceMemoryUsage:                  false,
		GetOnly:                            false,
		DisablePreParseMultipartForm:       false,
		LogAllErrors:                       false,
		SecureErrorLogMessage:              false,
		DisableHeaderNamesNormalizing:      false,
		SleepWhenConcurrencyLimitsExceeded: 0,
		NoDefaultServerHeader:              false,
		NoDefaultDate:                      false,
		NoDefaultContentType:               false,
		KeepHijackedConns:                  false,
		CloseOnShutdown:                    false,
		StreamRequestBody:                  false,
		ConnState:                          nil,
		Logger:                             nil,
		TLSConfig:                          nil,
		FormValueFunc:                      nil,
	}
	return s.fastHttpServer.ListenAndServe(host + ":" + port)
}
