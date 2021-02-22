package pswincom

import (
	"os"
	"testing"
)

func TestParseNumber(t *testing.T) {

	p1 := parseNumber("+4799999999")

	if p1 != "4799999999" {
		t.Errorf("Expected +4799999999 to be 4799999999, got: %s", p1)
	}

	p2 := parseNumber("004799999999")
	if p2 != "4799999999" {
		t.Errorf("Expected 004799999999 to be 4799999999, got: %s", p2)
	}

	p3 := parseNumber("99999999")
	if p3 != "4799999999" {
		t.Errorf("Expected 99999999 to be 4799999999, got: %s", p1)
	}

	p4 := parseNumber("4799999999")
	if p4 != "4799999999" {
		t.Errorf("Expected 4799999999 to be 4799999999, got: %s", p1)
	}
}

// TODO: Change phoneNumber to fetch from env
func TestSendMessage(t *testing.T) {
	client := NewClient(os.Getenv("PSWINCOM_USER"), os.Getenv("PSWINCOM_PASSWORD"), os.Getenv("PSWINCOM_SENDER"))
	m, err := client.SendMessage(os.Getenv("PSWINCOM_RECIPIENT_TEST"), "This is a test! With norwegian letters: æøå", false)

	if err != nil {
		t.Errorf("Request failed: %v", err)
	}

	if m.StatusCode != 200 {
		t.Errorf("Expected StatusCode 200, got: %d", m.StatusCode)
	}
}
