package pxpay2

import "encoding/xml"

type (
	XmlPurchase struct {
		XMLName            xml.Name      `xml:"GenerateRequest"`
		PxPayUserId        string        `xml:"PxPayUserId"`
		PxPayKey           string        `xml:"PxPayKey"`
		AmountInput        string        `xml:"AmountInput"`         // Max 13 characters
		BillingId          string        `xml:"BillingId,omitempty"` // Max 32 character field
		CurrencyInput      CurrencyType  `xml:"CurrencyInput"`       // Max 4 characters
		EmailAddress       string        `xml:"EmailAddress,omitempty"`
		EnableAddBillCard  int           `xml:"EnableAddBillCard,omitempty"` //Max 1 character Setting this field to “1” will cause Windcave to store the credit card details for future billing purposes
		RecurringMode      string        `xml:"RecurringMode,omitempty"`
		MerchantReference  string        `xml:"MerchantReference,omitempty"` //max 64 bytes
		DpsBillingId       string        `xml:"DpsBillingId,omitempty"`      //16 characters
		TxnData1           string        `xml:"TxnData1,omitempty"`          //Max 255 bytes
		TxnData2           string        `xml:"TxnData2,omitempty"`          //Max 255 bytes
		TxnData3           string        `xml:"TxnData3,omitempty"`          //Max 255 bytes
		TxnType            string        `xml:"TxnType"`
		TxnId              string        `xml:"TxnId"`
		UrlFail            string        `xml:"UrlFail"`
		UrlSuccess         string        `xml:"UrlSuccess"`
		UrlCallback        string        `xml:"UrlCallback,omitempty"` // Max 255 bytes
		Opt                string        `xml:"Opt,omitempty"`         //  default timeout of the created hosted payment page session is 72 hours. Universal Time (UTC). The value must be in the format "TO=yymmddHHmm"
		ForcePaymentMethod PaymentMethod `xml:"ForcePaymentMethod,omitempty"`
		//ClientType ClientType   `xml:"ClientType,omitempty"` // Not Supported
		//DebtRepaymentIndicator int           `xml:"DebtRepaymentIndicator,omitempty"` // Only send this field and set it to 1 to indicate that a debt repayment transaction is to be processed.
		//InstallmentNumber      int    `xml:"InstallmentNumber,omitempty"` // Not Supported
		//InstallmentCount       string `xml:"InstallmentCount"` // Not Supported
	}

	XmlRequest struct {
		XMLName xml.Name `xml:"Request"`
		Valid   int      `xml:"valid,attr"`
		URI     string   `xml:"URI"`
	}

	XmlProcessResponse struct {
		XMLName     xml.Name `xml:"ProcessResponse"`
		PxPayUserId string   `xml:"PxPayUserId"`
		PxPayKey    string   `xml:"PxPayKey"`
		Response    string   `xml:"Response"`
	}

	XmlResponse struct {
		XMLName             xml.Name `xml:"Response"`
		Valid               int      `xml:"valid,attr"`
		AmountSettlement    string   `xml:"AmountSettlement"`
		TotalAmount         string   `xml:"TotalAmount"`
		AmountSurcharge     string   `xml:"AmountSurcharge"`
		AuthCode            string   `xml:"AuthCode"`
		CardName            string   `xml:"CardName"`
		CardNumber          string   `xml:"CardNumber"`
		DateExpiry          string   `xml:"DateExpiry"`
		DpsTxnRef           string   `xml:"DpsTxnRef"`
		SurchargeDpsTxnRef  string   `xml:"SurchargeDpsTxnRef"`
		Success             int      `xml:"Success"`
		ResponseText        string   `xml:"ResponseText"`
		DpsBillingId        string   `xml:"DpsBillingId"`
		CardHolderName      string   `xml:"CardHolderName"`
		CurrencySettlement  string   `xml:"CurrencySettlement"`
		TxnData1            string   `xml:"TxnData1"`
		TxnData2            string   `xml:"TxnData2"`
		TxnData3            string   `xml:"TxnData3"`
		TxnType             string   `xml:"TxnType"`
		CurrencyInput       string   `xml:"CurrencyInput"`
		MerchantReference   string   `xml:"MerchantReference"`
		ClientInfo          string   `xml:"ClientInfo"`
		TxnId               string   `xml:"TxnId"`
		EmailAddress        string   `xml:"EmailAddress"`
		BillingId           string   `xml:"BillingId"`
		TxnMac              string   `xml:"TxnMac"`
		CardNumber2         string   `xml:"CardNumber2"`
		DateSettlement      string   `xml:"DateSettlement"`
		IssuerCountryId     string   `xml:"IssuerCountryId"`
		IssuerCountryCode   string   `xml:"IssuerCountryCode"`
		Cvc2ResultCode      string   `xml:"Cvc2ResultCode"`
		ReCo                string   `xml:"ReCo"`
		ProductSku          string   `xml:"ProductSku"`
		ShippingName        string   `xml:"ShippingName"`
		ShippingAddress     string   `xml:"ShippingAddress"`
		ShippingPostalCode  string   `xml:"ShippingPostalCode"`
		ShippingPhoneNumber string   `xml:"ShippingPhoneNumber"`
		ShippingMethod      string   `xml:"ShippingMethod"`
		BillingName         string   `xml:"BillingName"`
		BillingPostalCode   string   `xml:"BillingPostalCode"`
		BillingAddress      string   `xml:"BillingAddress"`
		BillingPhoneNumber  string   `xml:"BillingPhoneNumber"`
		PhoneNumber         string   `xml:"PhoneNumber"`
		AccountInfo         string   `xml:"AccountInfo"`
	}
)
