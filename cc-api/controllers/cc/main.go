package cc

import (
	"net/http"
	// "encoding/json"
	"github.com/bxcodec/faker/v3"
	"fmt"
	"cc-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Token string `faker:"uuid_hyphenated" json:"token" binding:"required"`
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

	err := faker.FakeData(&response)
	if err != nil {
		fmt.Println(err)
	}

	log.Info().
		Str("Credit card tokenized retrieved", "true").
		Msg("Token retrieved")

	c.JSON(http.StatusOK, response)
}
