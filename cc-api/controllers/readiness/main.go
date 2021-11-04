package readiness

import (
	"github.com/gin-gonic/gin"
	"orders-api/pkg/logger"
	"orders-api/pkg/memory_cache"
	"net/http"
)

type Response struct {
	Status string `json:"status" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in readiness
// @Tags readiness
// @Produce json
// @Success 200 {object} Response
// @Router /readiness [get]
func Ok(c *gin.Context) {
	m := memory_cache.GetInstance()
	log := logger.Instance()

	var response Response
	_, readiness_lock := m.Get("readiness.ok")

	if readiness_lock {
		response.Status = "NotReady"
		log.Warn().
			Str("status", response.Status).
			Str("user_agent", c.Request.Header.Get("User-Agent")).
			Msg("Liveness request probe failed")
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		response.Status = "Ready"
		log.Info().
			Str("status", response.Status).
			Str("user_agent", c.Request.Header.Get("User-Agent")).
			Msg("Liveness request successful")
		c.JSON(http.StatusOK, response)
	}
}
