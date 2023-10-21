package main

import (
	"automation-purwantara/tests"
	"automation-purwantara/utils"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// running test
	interval := 2 * time.Minute
	ticker := time.NewTicker(interval)

	RunTests()
	for {
		select {
		case <-ticker.C:
			RunTests()
		}
	}

}

func RunTests() {

	virtualAccountTest, err := tests.VirtualAccountTest()
	utils.PanicIfError(err)

	qrisTest, err := tests.QRISTest()
	utils.PanicIfError(err)

	ovoTest, err := tests.OvoTest()
	utils.PanicIfError(err)

	shopeeWalletTest, err := tests.ShopeeWalletTest()
	utils.PanicIfError(err)

	danaWalletTest, err := tests.DanaWalletTest()
	utils.PanicIfError(err)
	overTheCounterTest, err := tests.OverTheCounterTest()
	utils.PanicIfError(err)

	//Mengirim pesan ke Discord setelah goroutine selesai
	if len(virtualAccountTest) > 0 || len(qrisTest) > 0 || len(ovoTest) > 0 || len(shopeeWalletTest) > 0 || len(danaWalletTest) > 0 || len(overTheCounterTest) > 0 {
		utils.SendMessageToDiscord(utils.FormattedErrorMessage(virtualAccountTest, qrisTest, ovoTest, shopeeWalletTest, danaWalletTest, overTheCounterTest))
	} else {
		utils.SendMessageToDiscord("All test passed")
	}
}
