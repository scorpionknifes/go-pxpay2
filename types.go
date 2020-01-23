package pxpay2

import (
	"net/http"
)

const (
	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = "https://demo.windcave.com"
	// APIBaseLive points to the live version of the API
	APIBaseLive = " https://sec.windcave.com"
)

type PaymentMethod string

type CurrencyType string

type (
	// Client represents a Paypal REST API Client
	Client struct {
		Client      *http.Client
		PxPayUserId string
		PxPayKey    string
	}

	Purchase struct {
		MechantReference
		AmountInput
		CurrencyInput
	}

	Auth struct {
	}

	Txn struct {
		TxnData1
		TxnData2
		TxnData3
		TxnId
	}

	Amount struct {
		AmountInput
		CurrencyInput
	}
)
