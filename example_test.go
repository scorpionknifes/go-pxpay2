package pxpay2

import (
	"testing"
)

func TestPurchase(t *testing.T) {
	PxPayUserID := "Sample User"
	PxPayKey := "Sample Key"
	if PxPayUserID == "Sample User" {
		t.Error("Please change PxPayUserID in example_test.go")
	}

	c, err := NewClient(PxPayUserID, PxPayKey, APIBaseLive)
	if err != nil {
		t.Errorf("Not expected error for NewClient(), got %s", err.Error())
	}

	p := Purchase{
		MechantReference: "test",
		Amount:           Amount{AmountInput: "1.00", CurrencyInput: NZD},
		BillingId:        "test",
		EmailAddress:     "name@email.com",
		Txn:              Txn{"", "", "", ""},
		UrlSuccess:       "https://demo.windcave.com/SandboxSuccess.aspx",
		UrlFail:          "https://demo.windcave.com/SandboxSuccess.aspx",
	}
	test, err := c.CreatePurchase(p)
	if err != nil {
		t.Errorf("Not expected error for CreatePurchase(), got %s", err.Error())
	}
	t.Errorf("Click Link: %s", test)
}

func TestTranscation(t *testing.T) {
	PxPayUserID := "Sample User"
	PxPayKey := "Sample Key"
	if PxPayUserID == "Sample User" {
		t.Error("Please change PxPayUserID in example_test.go")
	}

	c, err := NewClient(PxPayUserID, PxPayKey, APIBaseLive)
	if err != nil {
		t.Errorf("Not expected error for NewClient(), got %s", err.Error())
	}

	p := ProcessResponse{
		Response: "ResponseKEY", // UrlSuccess?result=__Response_Token__
	}

	test, err := c.ProcessResponse(p)
	if err != nil {
		t.Errorf("Not expected error for ProcessResponse(), got %s", err.Error())
	}
	t.Error(test)
}
