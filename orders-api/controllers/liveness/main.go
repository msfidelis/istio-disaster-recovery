package liveness

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int `json:"status" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in liveness
// @Tags liveness
// @Produce json
// @Success 200 {object} Response
// @Router /liveness [get]
func Ok(c *gin.Context) {
	var response Response
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
