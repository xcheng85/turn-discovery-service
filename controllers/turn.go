package controllers

import (
	"fmt"
	"net/http"
	"github.com/xcheng85/turn-discovery-service/utils"
	"github.com/xcheng85/turn-discovery-service/webrtc"
)

type TurnController struct{}

func NewTurnController() *TurnController {
	return &TurnController{}
}

func (c *TurnController) RegisterRoutes(mux *http.ServeMux, cfg *utils.AppConfig) {
	mux.HandleFunc("/turn-web-api", c.retrieveRTCPeerConnection(cfg))
}

func (c *TurnController) retrieveRTCPeerConnection(cfg *utils.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Expose-Headers", "*")
			return
		}
		externalIP := cfg.TurnConfig.ExternalIp
		turnPort := cfg.TurnConfig.TurnPort
		sharedSecret := cfg.TurnSecret.Data.TurnSharedSecret
		ttlSeconds := cfg.TurnConfig.TTLSeconds
		// for lt cred match
		useLtCredMech := cfg.TurnConfig.UseLtCredMech
		username := cfg.TurnConfig.UserName
		password := cfg.TurnSecret.Data.Password

		rtcPeerConnection, err := webrtc.MakeRTCPeerConnection([]string{externalIP},
			turnPort, "fake", sharedSecret, ttlSeconds, useLtCredMech, username, password)
		if err != nil {
			utils.WriteJSONResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to retrieve RTC PeerConnection: %v", err))
			return
		}
		utils.WriteJSONResponse(w, http.StatusOK, rtcPeerConnection)
	}
}
