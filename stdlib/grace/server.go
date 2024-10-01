package grace

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup, log zerolog.Logger, httpServer *http.Server) {
	defer wg.Done()

	// Start the HTTP server in a separate goroutine
	go func() {
		log.Debug().Msg("Starting HTTP server...")
		log.Debug().Msg(fmt.Sprintf("HTTP server start on %s", httpServer.Addr))

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Debug().AnErr("HTTP server error", err)
		}
	}()

	// Wait for the context to be canceled
	select {
	case <-ctx.Done():
		log.Debug().Msg("HTTP server started...")

		// Shutdown the server gracefully
		log.Debug().Msg("Shutting down HTTP server gracefully...")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Debug().AnErr("HTTP server shutdown error", err)
		}

		log.Debug().Msg("HTTP server stopped.")
	}
}
