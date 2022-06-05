package server

import (
	"context"
	"net/http"
	"time"
	"{{ .AppName }}/lib/log"
)

func Start(ctx context.Context, addr string) {

	r := InitRouter()
	svr := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   25 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("server closed")
			} else {
				log.Fatal("%v", err)
			}
		}
	}()

	log.Info("Started listening on %s\n", addr)

	<-ctx.Done()
	log.Info("shutting down  server")

	if err := svr.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown: %v", err)
	}

}