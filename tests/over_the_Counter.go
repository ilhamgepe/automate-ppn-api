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

type OverTheCounter struct {
	ExpectedAmount int    `json:"expected_amount"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	PaymentChannel string `json:"payment_channel"`
	ExpiredAt      string `json:"expired_at"`
	ExternalId     string `json:"external_id"`
	MerchantTrxId  string `json:"merchant_trx_id"`
}

func OverTheCounterTest() ([]map[string]interface{}, error) {
	baseUrl := os.Getenv("BASE_URL")
	var OTCError []map[string]interface{}
	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(OverTheCounterChannels))

	fmt.Println("###  TESTING OVER THE COUNTER ###")

	bar := pb.StartNew(len(OverTheCounterChannels))
	for _, bank := range OverTheCounterChannels {
		var OTCResponse utils.ApiResponse
		var currentBankScope string = bank
		bar.Increment()
		time.Sleep(100 * time.Millisecond)
		go func(bank string) {
			defer wg.Done()
			currentBank := OverTheCounter{
				ExpectedAmount: 25000,
				Name:           fmt.Sprintf("Test Counter %s", bank),
				Description:    fmt.Sprintf("Description %s", bank),
				PaymentChannel: bank,
				ExpiredAt:      utils.CurrentTimeFormatted(),
				ExternalId:     fmt.Sprintf("EXT-ID-%v-%v-%v-%v-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
				MerchantTrxId:  fmt.Sprintf("MTX-ID-%v-%v-%v-%v-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute()),
			}
			reqBody, err := json.Marshal(currentBank)
			utils.PanicIfError(err)
			responseRaw, _, err := utils.SendHttpRequest(baseUrl+"/counter", utils.POST, reqBody)
			utils.PanicIfError(err)
			response, err := io.ReadAll(responseRaw.Body)
			utils.PanicIfError(err)
			err = json.Unmarshal(response, &OTCResponse)
			utils.PanicIfError(err)

			if OTCResponse.Status != 201 {
				OTCError = append(OTCError, map[string]interface{}{
					"Error":   OTCResponse.Error.Code,
					"Message": OTCResponse.Error.Message,
					"Channel": fmt.Sprintf("Counter %s", bank),
				})
			}
		}(currentBankScope)

	}
	bar.Finish()
	wg.Wait()
	fmt.Printf("### FINISHED OVER THE COUNTER TEST WITH %d ERRORS ###\n", len(OTCError))

	return OTCError, err
}
