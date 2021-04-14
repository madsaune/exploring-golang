package main

import (
	"log"

	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	applicationID = ""
	tenantID      = ""
)

func main() {
	deviceConfig := auth.NewDeviceFlowConfig(applicationID, tenantID)
	authorizer, err := deviceConfig.Authorizer()

	if err != nil {
		log.Fatalf("Could not authenticate: %v", err)
	}
}
