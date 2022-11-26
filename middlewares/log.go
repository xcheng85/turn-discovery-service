package middlewares

import (
	"go.uber.org/zap"
	"net/http"
)

func LogRequest(handler http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("logRequest", "RemoteAddr", r.RemoteAddr, "Method", r.Method, "URL", r.URL, "CorrelationId", r.Header.Get("correlation-id"))
		handler.ServeHTTP(w, r)
	})
}
