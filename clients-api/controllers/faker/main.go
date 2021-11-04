package faker

import (
	"net/http"
	"fmt"
	"clients-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/bxcodec/faker/v3"
)

type Response struct {
	FirstName          string  `faker:"first_name" json:"first_name" binding:"required"`
	LastName           string  `faker:"last_name" json:"last_name" binding:"required"`
	UserName           string  `faker:"username" json:"username" binding:"required"`
	PhoneNumber        string  `faker:"phone_number" json:"phone_number" binding:"required"`
	Email              string  `faker:"email" json:"email" binding:"required"`
	Website            string  `faker:"url" json:"website" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in healthcheck
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Response
// @Router /faker [get]
func Faker(c *gin.Context) {

	log := logger.Instance()

	var response Response

	err := faker.FakeData(&response)
	if err != nil {
		fmt.Println(err)
	}

	log.Info().
		Str("FirstName", response.FirstName).
		Msg("Client retrieved")


	c.JSON(http.StatusOK, response)
}
