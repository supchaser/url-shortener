package logging

import (
	"net/http"
	"time"
	"url-shortener/internal/pkg/utils/logging"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func NewLogrusLogger() func(next http.Handler) http.Handler {
	logger := logging.Logger.WithField("component", "middleware/logger")

	logger.Info("Logger middleware initialized")

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := logger.WithFields(logrus.Fields{
				"method":      r.Method,
				"path":        r.URL.Path,
				"remote_addr": r.RemoteAddr,
				"user_agent":  r.UserAgent(),
				"request_id":  middleware.GetReqID(r.Context()),
			})
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.WithFields(logrus.Fields{
					"status":   ww.Status(),
					"bytes":    ww.BytesWritten(),
					"duration": time.Since(t1).String(),
				}).Info("Request completed")
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
