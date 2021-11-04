package payments

import (
	"net/http"
	"fmt"
	"payment-api/pkg/logger"
	"payment-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
	"github.com/bxcodec/faker/v3"
)

type Response struct {
	OrderId			string 		`json:"id" binding:"required"`
	Amount    		float64  	`faker:"amount" json:"amount" binding:"required"`
	Currency  		string 		`faker:"currency" json:"currency" binding:"required"`
	PaymentMethod 	string 		`faker:"oneof: cc, paypal, check, money order" json:"payment_method" binding:"required"`
}

// Ok godoc
// @Summary Mock for Payment
// @Tags Payment
// @Produce json
// @Success 200 {object} Response
// @Router /payments/{id} [get]
func Get(c *gin.Context) {

	log := logger.Instance()

	var response Response

	err := faker.FakeData(&response)
	if err != nil {
		fmt.Println(err)
	}

	response.OrderId = c.Param("id")

	if (response.PaymentMethod == "cc") {
		headers := make(map[string][]string)
		headers["Content-type"] = append(headers["Content-type"], "Application/json")
	
		// Get Clients - Mock
		responseCC, body := httpclient.Request("POST", "http://cc-api:8080", "/cc", headers, "{}")

		if responseCC.StatusCode != 200 {
			log.Error().
				Str("response for cc-api", body).
				Int("status_code", responseCC.StatusCode).
				Msg("Failed to retrieve a credit card information")
	
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
				"component": "cc-api",
			})
			return
		}
	}

	log.Info().
		Str("OrderId", response.OrderId).
		Msg("Payment retrieved")


	c.JSON(http.StatusOK, response)
}
