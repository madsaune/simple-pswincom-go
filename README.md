# simple-pswincom-go

## Example

```go
package main

import (
  "log"

  "github.com/madsaune/simple-pswincom-go"
)

func main() {

  // Set credentials manually
  // client := pswincom.NewClient("<username>", "<password>", "<senderName>", false)

  // Get credentials from environment variables
  client := pswincom.NewClientFromEnv(nil)
  
  if err := client.SendMessage("+4799999999", "This is a test! With norwegian letters: æøå", false); err != nil {
    log.Fatalf("Could not send SMS: %v", err)
  }
}
```