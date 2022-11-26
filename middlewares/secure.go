package middlewares

import(
	"net/http"
	"go.uber.org/zap"
	"time"
)

func SecureResponse(handler http.Handler, logger *zap.SugaredLogger) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add response headers for security checks
		// This could also be done through istio
		w.Header().Add("X-Frame-Options", "deny")
		w.Header().Add("Expires", time.Now().Local().Add(time.Second * time.Duration(3600)).Format(time.RFC1123))
		handler.ServeHTTP(w, r)
	})
}