package utils

import "fmt"

func FormattedErrorMessage(vaErr, qrisErr, ovoErr, shopeeWalletErr, danaWalletErr, overTheCounterErr []map[string]interface{}) string {
	var message string

	addErrorToMessage(&message, vaErr)
	addErrorToMessage(&message, qrisErr)
	addErrorToMessage(&message, ovoErr)
	addErrorToMessage(&message, shopeeWalletErr)
	addErrorToMessage(&message, danaWalletErr)
	addErrorToMessage(&message, overTheCounterErr)
	//
	//for _, value := range VaErr {
	//	message += fmt.Sprintf("error %v saat create %v, message: %v\n\n", value["Error"], value["Channel"], value["Message"])
	//}
	//
	//for _, value := range QRISErr {
	//	message += fmt.Sprintf("error %v saat create %v, message: %v\n\n", value["Error"], value["Channel"], value["Message"])
	//}

	return message
}

func addErrorToMessage(message *string, error []map[string]interface{}) {
	for _, value := range error {
		*message += fmt.Sprintf("error %v saat create %v, message: %v\n\n", value["Error"], value["Channel"], value["Message"])
	}
}
