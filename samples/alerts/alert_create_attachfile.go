package main

import (
	"fmt"
	"os"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

func main() {
	API_KEY := "YOUR API KEY HERE"
	PATH_TO_FILE := "/your/path/to/file/here"

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

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

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	file, err := os.OpenFile(PATH_TO_FILE, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	attachFileReq := alerts.AttachFileAlertRequest{Id: response.AlertId, Attachment: file}
	attachFileResp, attachFileErr := alertCli.AttachFile(attachFileReq)

	if attachFileErr != nil {
		panic(attachFileErr)
	}

	fmt.Println("Status:", attachFileResp.Status)
	fmt.Println("Code:", attachFileResp.Code)
}
