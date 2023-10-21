package tests

import (
	"automation-purwantara/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type OVO struct {
	CustomerPhone  string `json:"customer_phone"`
	CustomerEmail  string `json:"customer_email"`
	PaymentChannel string `json:"payment_channel"`
	ExternalId     string `json:"external_id"`
	Description    string `json:"description"`
	Amount         int    `json:"amount"`
}

func OvoTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var OvoError []map[string]interface{}
	var err error

	fmt.Println("###  TESTING OVO ###")

	var OvoResponse utils.ApiResponse
	CurrentOvo := OVO{
		CustomerPhone:  "089675544501",
		CustomerEmail:  "ilhamganteng@gmail.com",
		PaymentChannel: "OVO",
		ExternalId:     fmt.Sprintf("EXT-ID-%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
		Description:    fmt.Sprintf("Description OVO"),
		Amount:         25000,
	}
	reqBody, err := json.Marshal(CurrentOvo)
	utils.PanicIfError(err)
	responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/ewallet/ovo", utils.POST, reqBody)
	utils.PanicIfError(err)
	response, err := io.ReadAll(responseRaw.Body)
	utils.PanicIfError(err)
	err = json.Unmarshal(response, &OvoResponse)
	utils.PanicIfError(err)

	if OvoResponse.Status != 201 {
		OvoError = append(OvoError, map[string]interface{}{
			"Error":   OvoResponse.Error.Code,
			"Message": OvoResponse.Error.Message,
			"Channel": fmt.Sprintf("OVO"),
		})
	}

	fmt.Printf("### FINISHED OVO TEST WITH %d ERRORS ###\n", len(OvoError))

	return OvoError, err
}
