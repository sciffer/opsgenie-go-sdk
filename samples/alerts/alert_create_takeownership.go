package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
	"github.com/opsgenie/opsgenie-go-sdk/samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(constants.API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertId)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// take ownership of the alert
	takeOwnershipReq := alerts.TakeOwnershipAlertRequest{Id: response.AlertId}
	takeOwnershipResponse, alertErr := alertCli.TakeOwnership(takeOwnershipReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("status: %s\n", takeOwnershipResponse.Status)
	fmt.Printf("code: %d\n", takeOwnershipResponse.Code)
}
