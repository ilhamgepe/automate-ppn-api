package tests

import (
	"automation-purwantara/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type ShopeeWallet struct {
	Amount                 int    `json:"amount"`
	TransactionDescription string `json:"transaction_description"`
	CustomerEmail          string `json:"customer_email"`
	CustomerFirstName      string `json:"customer_first_name"`
	CustomerLastName       string `json:"customer_last_name"`
	CustomerPhone          string `json:"customer_phone"`
	PaymentChannel         string `json:"payment_channel"`
	ExternalId             string `json:"external_id"`
	ReturnUrl              string `json:"return_url"`
}

func ShopeeWalletTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var ShopeeWalletError []map[string]interface{}
	var err error

	fmt.Println("###  TESTING SHOPEE WALLET ###")

	var shopeepayResponse utils.ApiResponse
	CurrentOvo := ShopeeWallet{
		Amount:                 25000,
		TransactionDescription: "SHOPEE WALLET",
		CustomerEmail:          "ilhamganteng@gmail.com",
		CustomerFirstName:      "ilham",
		CustomerLastName:       "ganteng",
		CustomerPhone:          "089675544501",
		PaymentChannel:         "shopee_pay",
		ExternalId:             fmt.Sprintf("%v-%v-%v-%v", time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
		ReturnUrl:              "https://youtu.be/dQw4w9WgXcQ?si=1ZRP8oheBs5A35q5",
	}
	reqBody, err := json.Marshal(CurrentOvo)
	utils.PanicIfError(err)
	responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/ewallet/shopeepay", utils.POST, reqBody)
	utils.PanicIfError(err)
	response, err := io.ReadAll(responseRaw.Body)
	utils.PanicIfError(err)
	err = json.Unmarshal(response, &shopeepayResponse)
	utils.PanicIfError(err)

	if shopeepayResponse.Status != 201 {
		ShopeeWalletError = append(ShopeeWalletError, map[string]interface{}{
			"Error":   shopeepayResponse.Error.Code,
			"Message": shopeepayResponse.Error.Message,
			"Channel": fmt.Sprintf("SHOPEE WALLET"),
		})
	}

	fmt.Printf("### FINISHED SHOPEE WALLET TEST WITH %d ERRORS ###\n", len(ShopeeWalletError))

	return ShopeeWalletError, err
}
