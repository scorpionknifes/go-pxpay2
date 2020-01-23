### Go client for Windcave

### New Client

```go
import "github.com/scorpionknifes/go-pxpay2"

// Create a client instance
c, err := pxpay.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
c.SetLog(os.Stdout) // Set log to terminal stdout

accessToken, err := c.GetAccessToken()
```

### Get authorization by ID

```go
auth, err := c.GetAuthorization("2DC87612EK520411B")
```