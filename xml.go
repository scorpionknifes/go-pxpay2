package pxpay2

import "encoding/xml"

type (
	Payment struct {
		XMLName                xml.Name `xml:"GenerateRequest"`
		PxPayUserId            string   `xml:"PxPayUserId"`
		PxPayKey               string   `xml:"PxPayKLey"`
		AmountInput            string   `xml:"AmountInput"`         // Max 13 characters
		BillingId              string   `xml:"BillingId",omitempty` // Max 32 character field
		CurrencyInput          string   `xml:"CurrencyInput"`       // Max 4 characters
		EmailAddress           string   `xml:"EmailAddress",omitempty`
		EnableAddBillCard      int      `xml:"EnableAddBillCard",omitempty` //Max 1 character Setting this field to “1” will cause Windcave to store the credit card details for future billing purposes
		RecurringMode          string   `xml:"RecurringMode",omitempty`
		MerchantReference      string   `xml:"MerchantReference",omitempty` //max 64 bytes
		DpsBillingId           string   `xml:"DpsBillingId",omitempty`      //16 characters
		TxnData1               string   `xml:"TxnData1",omitempty`          //Max 255 bytes
		TxnData2               string   `xml:"TxnData2",omitempty`          //Max 255 bytes
		TxnData3               string   `xml:"TxnData3",omitempty`          //Max 255 bytes
		TxnType                string   `xml:"TxnType"`
		TxnId                  string   `xml:"TxnId"`
		UrlFail                string   `xml:"UrlFail"`
		UrlSuccess             string   `xml:"UrlSuccess"`
		UrlCallback            string   `xml:"UrlCallback",omitempty` // Max 255 bytes
		Opt                    string   `xml:"Opt",omitempty`         // Max 64 bytes
		ClientType             string   `xml:"ClientType",omitempty`
		ForcePaymentMethod     string   `xml:"ForcePaymentMethod",omitempty`
		DebtRepaymentIndicator int      `xml:"DebtRepaymentIndicator",omitempty`
		InstallmentNumber      int      `xml:"InstallmentNumber",omitempty`
		InstallmentCount       string   `xml:"InstallmentCount"`
	}
)
