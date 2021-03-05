package pswincom

import "testing"

func TestMessageRecipientFormatting(t *testing.T) {
	message1 := NewMessage("+4799999999", "Hello, World!", false)
	if message1.Recipient != "4799999999" {
		t.Errorf("Input: %s, Expected: %s, Got: %s", "+4799999999", "4799999999", message1.Recipient)
	}

	message2 := NewMessage("004799999999", "Hello, World!", false)
	if message2.Recipient != "4799999999" {
		t.Errorf("Input: %s, Expected: %s, Got: %s", "004799999999", "4799999999", message2.Recipient)
	}

	message3 := NewMessage("99999999", "Hello, World!", false)
	if message3.Recipient != "4799999999" {
		t.Errorf("Input: %s, Expected: %s, Got: %s", "99999999", "4799999999", message3.Recipient)
	}

	message4 := NewMessage("4799999999", "Hello, World!", false)
	if message4.Recipient != "4799999999" {
		t.Errorf("Input: %s, Expected: %s, Got: %s", "4799999999", "4799999999", message4.Recipient)
	}
}
