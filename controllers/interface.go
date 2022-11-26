package controllers

import (
	"github.com/xcheng85/turn-discovery-service/utils"
	"net/http"
)

// dependencies injection
type Controller interface {
	RegisterRoutes(mux *http.ServeMux, cfg *utils.AppConfig)
}
