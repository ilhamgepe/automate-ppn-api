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

type VirtualAccount struct {
	ExpectedAmount int    `json:"expected_amount"`
	Name           string `json:"name"`
	Bank           string `json:"bank"`
	Description    string `json:"description"`
	ExpiredAt      string `json:"expired_at"`
	ExternalId     string `json:"external_id"`
	MerchantTrxId  string `json:"merchant_trx_id"`
}

func VirtualAccountTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var VAErrors []map[string]interface{}
	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(VirtualAccountBanks))

	fmt.Println("###  TESTING VIRTUAL ACCOUNT  ###")

	bar := pb.StartNew(len(VirtualAccountBanks))
	for _, bank := range VirtualAccountBanks {
		var VAResponse utils.ApiResponse
		var currentBankScope string = bank
		bar.Increment()
		time.Sleep(100 * time.Millisecond)
		go func(bank string) {
			defer wg.Done()
			currentBank := VirtualAccount{
				ExpectedAmount: 25000,
				Name:           fmt.Sprintf("Test Virtual Account %s", bank),
				Bank:           bank,
				Description:    fmt.Sprintf("Description %s", bank),
				ExpiredAt:      utils.CurrentTimeFormatted(),
				ExternalId:     fmt.Sprintf("EXT-ID-%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
				MerchantTrxId:  fmt.Sprintf("MTX-ID-%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
			}
			reqBody, err := json.Marshal(currentBank)
			utils.PanicIfError(err)
			responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/virtual-account", utils.POST, reqBody)
			utils.PanicIfError(err)
			response, err := io.ReadAll(responseRaw.Body)
			utils.PanicIfError(err)
			err = json.Unmarshal(response, &VAResponse)
			utils.PanicIfError(err)

			if VAResponse.Status != 201 {
				VAErrors = append(VAErrors, map[string]interface{}{
					"Error":   VAResponse.Error.Code,
					"Message": VAResponse.Error.Message,
					"Channel": fmt.Sprintf("VA %s", bank),
				})
			}
		}(currentBankScope)

	}
	bar.Finish()
	wg.Wait()
	fmt.Printf("### FINISHED VIRTUAL ACCOUNT TEST WITH %d ERRORS ###\n", len(VAErrors))

	return VAErrors, err
}
