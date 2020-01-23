### Go client for Windcave

### New Client

```go
import "github.com/scorpionknifes/go-pxpay2"

// Create a client instance
c, err := pxpay2.NewClient("PxPayUserID", "PxPayKey", pxpay2.APIBaseLive)
```

### Create Purchase

```go
// Config Purchase
p := pxpay2.Purchase{
    MechantReference: "test",
    Amount:           pxpay2.Amount{pxpay2.AmountInput: "1.00", pxpay2.CurrencyInput: NZD},
    BillingId:        "test",
    EmailAddress:     "name@email.com",
    Txn:              pxpay2.Txn{"", "", "", ""},
    UrlSuccess:       "https://demo.windcave.com/SandboxSuccess.aspx",
    UrlFail:          "https://demo.windcave.com/SandboxSuccess.aspx",
}
// Redirect User to Hosted Payments Page using URL
redirecturl, err := c.CreatePurchase(p)
```

### Verify Purchase
```go
// Use result parameter from redirect
p := ProcessResponse{
	Response: "__result__", // __UrlSuccess__?result=__result__
}
//Return XML struct
res, err := c.ProcessResponse(p)
```
