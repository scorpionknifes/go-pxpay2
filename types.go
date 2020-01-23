package pxpay2

import (
	"net/http"
)

const (
	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = "https://demo.windcave.com"
	// APIBaseLive points to the live version of the API
	APIBaseLive = "https://sec.windcave.com"
)

type PaymentMethod string

type CurrencyType string

type ClientType string

type (
	// Client represents a Paypal REST API Client
	Client struct {
		Client      *http.Client
		PxPayUserId string
		PxPayKey    string
		APIBase     string
	}

	Purchase struct {
		MechantReference string
		Amount           Amount
		BillingId        string
		EmailAddress     string
		Txn              Txn
		UrlSuccess       string
		UrlFail          string
	}

	Auth struct {
	}

	Txn struct {
		TxnData1 string
		TxnData2 string
		TxnData3 string
		TxnId    string
	}

	Amount struct {
		AmountInput   string
		CurrencyInput CurrencyType
	}

	Transcation struct {
		Result string
	}

	ProcessResponse struct {
		Response string
	}
)
