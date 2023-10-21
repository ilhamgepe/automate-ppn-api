package tests

import (
	"automation-purwantara/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type DanaWallet struct {
	Amount         int    `json:"amount"`
	CustomerEmail  string `json:"customer_email"`
	Description    string `json:"description"`
	CustomerPhone  string `json:"customer_phone"`
	PaymentChannel string `json:"payment_channel"`
	ExternalId     string `json:"external_id"`
	MerchantTrxId  string `json:"merchant_trx_id"`
	PayReturn      string `json:"pay_return"`
}

func DanaWalletTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var DanaWalletErr []map[string]interface{}
	var err error

	fmt.Println("###  TESTING DANA WALLET ###")

	var danaWalletResponse utils.ApiResponse
	CurrentOvo := DanaWallet{
		Amount:         25000,
		CustomerEmail:  "ilhamganteng@gmail.com",
		Description:    "DANA WALLET",
		CustomerPhone:  "089675544501",
		PaymentChannel: "wallet_dana",
		ExternalId:     fmt.Sprintf("EXT-ID-%v-%v-%v-%v", time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
		MerchantTrxId:  fmt.Sprintf("MTX-ID-%v-%v-%v-%v", time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
		PayReturn:      "https://youtu.be/dQw4w9WgXcQ?si=1ZRP8oheBs5A35q5",
	}
	reqBody, err := json.Marshal(CurrentOvo)
	utils.PanicIfError(err)
	responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/ewallet/dana", utils.POST, reqBody)
	utils.PanicIfError(err)
	response, err := io.ReadAll(responseRaw.Body)
	utils.PanicIfError(err)
	err = json.Unmarshal(response, &danaWalletResponse)
	utils.PanicIfError(err)

	if danaWalletResponse.Status != 201 {
		DanaWalletErr = append(DanaWalletErr, map[string]interface{}{
			"Error":   danaWalletResponse.Error.Code,
			"Message": danaWalletResponse.Error.Message,
			"Channel": fmt.Sprintf("DANA WALLET"),
		})
	}

	fmt.Printf("### FINISHED DANA WALLET TEST WITH %d ERRORS ###\n", len(DanaWalletErr))

	return DanaWalletErr, err
}
