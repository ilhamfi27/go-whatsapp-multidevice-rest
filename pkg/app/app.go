package app

import (
	"log"
	"time"

	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/app/http"
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/env"
)

var (
	AppWebhookURL       string
	AppWebhookBasicAuth string
	AppRequest          *http.HttpClient
)

func init() {
	var err error

	appWebhookUrl, err := env.GetEnvString("APP_WEBHOOK_URL_TARGET")
	if err != nil {
		log.Fatal("Error Parse Environment Variable for App Webhook URL Target")
	}
	AppWebhookURL = appWebhookUrl

	appWebhookBasicAuth, err := env.GetEnvString("APP_WEBHOOK_BASIC_AUTH")
	if err != nil {
		AppWebhookBasicAuth = ""
	}
	AppWebhookBasicAuth = appWebhookBasicAuth

	// Initialize App HTTP Request
	initHttpRequest()
}

func initHttpRequest() {
	// Initialize App HTTP Request
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if AppWebhookBasicAuth != "" {
		headers["Authorization"] = "Basic " + AppWebhookBasicAuth
	}

	client := http.NewHttpClient(http.HttpClientOptions{
		Timeout: 30 * time.Second,
		Headers: headers,
	})

	AppRequest = client
}
