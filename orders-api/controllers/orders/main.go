package orders

import (
	"net/http"
	"encoding/json"
	"fmt"
	"orders-api/pkg/logger"
	"orders-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

type Response struct {
	OrderId string `json: "id"`
	Client struct {
		FirstName  string `json:"first_name"`
		LastName string `json:"last_name"`
		PhoneNumber string `json:"phone_num er"`
	} `json:"client"`
	Payment struct {
		Amount    		float64  	`json:"amount" binding:"required"`
		Currency  		string 		`json:"currency" binding:"required"`
		PaymentMethod 	string 		`json:"payment_method" binding:"required"`
	} `json:"payment"`
}

type Client struct {
	FirstName          string  `json:"first_name" binding:"required"`
	LastName           string  `json:"last_name" binding:"required"`
	UserName           string  `json:"username" binding:"required"`
	PhoneNumber        string  `json:"phone_number" binding:"required"`
	Email              string  `json:"email" binding:"required"`
	Website            string  `json:"website" binding:"required"`
}

type Payment struct {
	Amount    		float64  	`json:"amount" binding:"required"`
	Currency  		string 		`json:"currency" binding:"required"`
	PaymentMethod 	string 		`json:"payment_method" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in healthcheck
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Response
// @Router /orders/{id} [get]
func Get(c *gin.Context) {

	log := logger.Instance()

	var response Response
	var client Client
	var payment Payment

	response.OrderId = c.Param("id")

	headers := make(map[string][]string)
	headers["Content-type"] = append(headers["Content-type"], "Application/json")

	// Get Clients - Mock
	responseClients, body := httpclient.Request("GET", "http://clients-api:8080", "/clients", headers, "{}")

	if responseClients.StatusCode != 200 {
		log.Error().
			Str("response for clients-api", body).
			Int("status_code", responseClients.StatusCode).
			Msg("Failed to retrieve a client")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"component": "clients-api",
		})
		return
	}

	json.Unmarshal([]byte(body ), &client)

	response.Client.FirstName = client.FirstName
	response.Client.LastName = client.LastName
	response.Client.PhoneNumber = client.PhoneNumber

	// Get Payment
	responsePayments, bodyPayment := httpclient.Request("GET", "http://payment-api:8080", fmt.Sprintf("/payments/%v", response.OrderId), headers, "{}")

	if responsePayments.StatusCode != 200 {
		log.Error().
			Str("response for payment-api", body).
			Int("status_code", responseClients.StatusCode).
			Msg("Failed to retrieve a payment information")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"component": "payment-api",
		})
		return
	}

	json.Unmarshal([]byte(bodyPayment), &payment)

	response.Payment.Amount = payment.Amount
	response.Payment.Currency = payment.Currency
	response.Payment.PaymentMethod = payment.PaymentMethod	

	log.Info().
		Str("Order retrieved", "1").
		Str("Infos", body).
		Msg("Client retrieved")


	c.JSON(http.StatusOK, response)
}
