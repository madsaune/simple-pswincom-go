package pswincom

import (
	"os"
	"testing"
)

// TODO: Change phoneNumber to fetch from env
func TestSendMessage(t *testing.T) {
	client := NewClient(os.Getenv("PSWINCOM_USER"), os.Getenv("PSWINCOM_PASSWORD"), os.Getenv("PSWINCOM_SENDER"), nil)

	message := &Message{
		Recipient: os.Getenv("PSWINCOM_RECIPIENT_TEST"),
		Body:      "This is a test! With norwegian letters: æøå",
	}

	err := client.SendMessage(message)

	if err != nil {
		t.Errorf("Request failed: %v", err)
	}
}
