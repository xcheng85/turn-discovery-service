package controllers

import (
	"github.com/xcheng85/turn-discovery-service/utils"
	"net/http"
)

type K8sLivenessProbeController struct{}

func NewK8sLivenessProbeController() *K8sLivenessProbeController {
	return &K8sLivenessProbeController{}
}

func (c *K8sLivenessProbeController) RegisterRoutes(mux *http.ServeMux, cfg *utils.AppConfig) {
	mux.HandleFunc("/livenessProbe", c.livenessProbe())
}

func (c *K8sLivenessProbeController) livenessProbe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteStatusResponse(w, http.StatusOK, "livenessProbe passes")
	}
}

type K8sReadinessProbeController struct{}

func NewK8sReadinessProbeController() *K8sReadinessProbeController {
	return &K8sReadinessProbeController{}
}

func (c *K8sReadinessProbeController) RegisterRoutes(mux *http.ServeMux, cfg *utils.AppConfig) {
	mux.HandleFunc("/readinessProbe", c.readinessProbe())
}

func (c *K8sReadinessProbeController) readinessProbe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteStatusResponse(w, http.StatusOK, "readinessProbe passes")
	}
}
