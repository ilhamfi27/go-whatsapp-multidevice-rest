package app

import (
	"log"

	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/app/database"
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/env"
)

var (
	AppWebhookURL       string
	AppWebhookBasicAuth string
	AppDatabase         *database.DatabaseContainer
)

func init() {
	var err error

	dbType, err := env.GetEnvString("WHATSAPP_DATASTORE_TYPE")
	if err != nil {
		log.Fatal("Error Parse Environment Variable for Application Datastore Type")
	}

	dbURI, err := env.GetEnvString("WHATSAPP_DATASTORE_URI")
	if err != nil {
		log.Fatal("Error Parse Environment Variable for Application Datastore URI")
	}

	// Initialize App Client Datastore
	appDb, err := database.New(dbType, dbURI)
	if err != nil {
		log.Fatal("Error Connect Application Datastore: ", err)
	}
	AppDatabase = appDb

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
}
