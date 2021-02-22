# simple-pswincom-go

## Example

```go
package main

import (
  "log"

  "github.com/madsaune/simple-pswincom-go"
)

func main() {
  client := pswincom.NewClient("<username>", "<password>", "<senderName>")
	m, err := client.SendMessage("+4799999999", "This is a test! With norwegian letters: æøå", false)

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	if m.StatusCode != 200 {
		log.Fatalf("Something went wrong, got: %d", m.StatusCode)
	}
}
```