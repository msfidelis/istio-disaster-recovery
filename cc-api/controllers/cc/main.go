package orders

import (
	"net/http"
	"encoding/json"
	"orders-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Response struct {
	OrderId string `json: "id" faker:"uuid_hyphenated"`
}

// Ok godoc
// @Summary Return 200 status Ok in healthcheck
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Response
// @Router /cc [get]
func Get(c *gin.Context) {

	log := logger.Instance()

	var response Response

	headers := make(map[string][]string)
	headers["Content-type"] = append(headers["Content-type"], "Application/json")

	log.Info().
		Str("Credit card tokenized retrieved", "true").
		Msg("Token retrieved")

	c.JSON(http.StatusOK, response)
}
