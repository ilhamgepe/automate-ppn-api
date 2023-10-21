package tests

import (
	"automation-purwantara/utils"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"sync"
	"time"
)

type Qris struct {
	Amount                 int    `json:"amount"`
	TransactionDescription string `json:"transaction_description"`
	CustomerEmail          string `json:"customer_email"`
	CustomerFirstName      string `json:"customer_first_name"`
	CustomerLastName       string `json:"customer_last_name"`
	CustomerPhone          string `json:"customer_phone"`
	PaymentChannel         string `json:"payment_channel"`
	PaymentMethod          string `json:"payment_method"`
	OrderId                string `json:"order_id"`
	MerchantTrxId          string `json:"merchant_trx_id"`
}

func QRISTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var QRISError []map[string]interface{}
	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(QRISChannels))

	fmt.Println("###  TESTING QRIS  ###")

	bar := pb.StartNew(len(QRISChannels))
	for _, bank := range QRISChannels {
		var QRISResponse utils.ApiResponse
		var currenQRISScope string = bank
		bar.Increment()
		time.Sleep(100 * time.Millisecond)
		go func(bank string) {
			defer wg.Done()
			currentQRIS := Qris{
				Amount:                 25000,
				TransactionDescription: fmt.Sprintf("Description %s", bank),
				CustomerEmail:          "ilhamganteng@gmail.com",
				CustomerFirstName:      "ilham",
				CustomerLastName:       "gilang",
				CustomerPhone:          "089675544503",
				PaymentChannel:         bank,
				PaymentMethod:          "wallet",
				OrderId:                fmt.Sprintf("%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
				MerchantTrxId:          fmt.Sprintf("%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
			}
			reqBody, err := json.Marshal(currentQRIS)
			utils.PanicIfError(err)
			responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/qris", utils.POST, reqBody)
			utils.PanicIfError(err)
			response, err := io.ReadAll(responseRaw.Body)
			utils.PanicIfError(err)
			err = json.Unmarshal(response, &QRISResponse)
			utils.PanicIfError(err)

			if QRISResponse.Status != 201 {
				QRISError = append(QRISError, map[string]interface{}{
					"Error":   QRISResponse.Error.Code,
					"Message": QRISResponse.Error.Message,
					"Channel": fmt.Sprintf("QRIS %s", bank),
				})
			}
		}(currenQRISScope)

	}
	bar.Finish()
	wg.Wait()
	fmt.Printf("### FINISHED QRIS  TEST WITH %d ERRORS ###\n", len(QRISError))

	return QRISError, err
}
