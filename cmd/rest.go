package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/simple-web-app/rest"
	"github.com/simple-web-app/service"
)

func serveRest() error {
	svc := service.NewService()

	hlthHndlr := rest.NewHealthHandler()
	usrHndlr := rest.NewUserHandler(svc)
	shopHndlr := rest.NewShopHandler(svc)
	productHndlr := rest.NewProductHandler(svc)
	brandHndlr := rest.NewBrandHandler(svc)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Mount("/api/health", hlthHndlr.Router())
	r.Mount("/api/v1/users", usrHndlr.Router())
	r.Mount("/api/v1/shops", shopHndlr.Router())
	r.Mount("/api/v1/products", productHndlr.Router())
	r.Mount("/api/v1/brands", brandHndlr.Router())

	timeout := 30 * time.Second
	srvr := http.Server{
		Addr:         "0.0.0.0:5000",
		Handler:      r,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
	}

	errCh := make(chan error)

	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt}

	graceful := func() error {
		log.Println("Shutting down server gracefully with in", timeout)
		log.Println("To shutdown immediately press again")

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		return srvr.Shutdown(ctx)
	}

	forced := func() error {
		log.Println("Shutting down server forcefully")
		return srvr.Close()
	}

	go func() {
		log.Println("Starting server on", srvr.Addr)
		if err := srvr.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	go func() {
		errCh <- HandleSignals(sigs, graceful, forced)
	}()

	return <-errCh
}

// HandleSignals listen on the registered signals and fires the gracefulHandler for the
// first signal and the forceHandler (if any) for the next this function blocks and
// return any error that returned by any of the handlers first
func HandleSignals(sigs []os.Signal, gracefulHandler, forceHandler func() error) error {
	sigCh := make(chan os.Signal)
	errCh := make(chan error, 1)

	signal.Notify(sigCh, sigs...)
	defer signal.Stop(sigCh)

	grace := true

	select {
	case err := <-errCh:
		return err
	case <-sigCh:
		if grace {
			grace = false
			go func() {
				errCh <- gracefulHandler()
			}()
		} else if forceHandler != nil {
			errCh <- forceHandler()
		}
	}

	return nil
}
