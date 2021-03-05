package pswincom

import (
	"os"
	"testing"
)

func TestClientSendMessage(t *testing.T) {
	client := NewClient(os.Getenv("PSWINCOM_USER"), os.Getenv("PSWINCOM_PASSWORD"), os.Getenv("PSWINCOM_SENDER"), nil)

	message := NewMessage(os.Getenv("PSWINCOM_RECIPIENT_TEST"), "This is a test using NewClient! With norwegian letters: æøå", false)
	err := client.SendMessage(message)

	if err != nil {
		t.Errorf("Request failed: %v", err)
	}
}

func TestClientFromEnvSendMessage(t *testing.T) {
	client := NewClientFromEnv(nil)
	message := NewMessage(os.Getenv("PSWINCOM_RECIPIENT_TEST"), "This is a test using NewClientFromEnv! With norwegian letters: æøå", false)
	err := client.SendMessage(message)

	if err != nil {
		t.Errorf("Request failed: %v", err)
	}
}
