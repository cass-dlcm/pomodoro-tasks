package secrets

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var projectId = "pomodoro-tasks-327213"

func GetSecret(secret string) string {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Panicf("failed to setup client: %v", err)
	}
	defer func(client *secretmanager.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectId, secret),
	}
	response, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Panicf("failed to retrieve string: %v", err)
	}
	return string(response.GetPayload().Data)
}
