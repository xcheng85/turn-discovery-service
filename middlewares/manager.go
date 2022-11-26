package middlewares

import(
	"net/http"
	"go.uber.org/zap"
)

// function programming example
type Middleware func(http.Handler, *zap.SugaredLogger) http.Handler
func MiddlewareManager(h http.Handler, logger *zap.SugaredLogger, m ...Middleware) http.Handler {
	if(len(m) < 1){
		return h
	}
	wrapped := h
	// pay attention to the order
	for i := len(m) - 1; i > 0; i-- {
		wrapped = m[i](wrapped, logger)
	}
	return wrapped
}
