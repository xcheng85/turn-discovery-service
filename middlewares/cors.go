package middlewares

import(
	"net/http"
	"go.uber.org/zap"
)

func Cors(handler http.Handler, logger *zap.SugaredLogger) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for the preflight request
		// This could also be done through istio
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		handler.ServeHTTP(w, r)
	})
}