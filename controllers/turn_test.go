package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xcheng85/turn-discovery-service/utils"
	"github.com/xcheng85/turn-discovery-service/webrtc"
)

func TestRetrieveRTCPeerConnection(t *testing.T) {
	c := NewTurnController()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.URL.Path = "turn-discovery-service"
	w := httptest.NewRecorder()
	logger := utils.NewLogger()
	config := utils.NewConfig(logger)
	handler := c.retrieveRTCPeerConnection(config)
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	rtc := webrtc.RTCPeerConnection{}
	err := json.NewDecoder(res.Body).Decode(&rtc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "3600s", rtc.LifetimeDuration)
}