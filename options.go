package pxpay2

const (
	Account2Account   PaymentMethod = "Account2Account"
	Alipay            PaymentMethod = "Alipay"
	AmexExpressWallet PaymentMethod = "AmexExpressWallet"
	Card              PaymentMethod = "Card"
	GECard            PaymentMethod = "GECard"
	MasterPass        PaymentMethod = "MasterPass"
	MonerisIOP        PaymentMethod = "MonerisIOP"
	Oxipay            PaymentMethod = "Oxipay"
	PayPal            PaymentMethod = "PayPal"
	UPOP5             PaymentMethod = "UPOP5"
	UPOP              PaymentMethod = "UPOP"
	VisaCheckout      PaymentMethod = "VisaCheckout"
	WeChat            PaymentMethod = "WeChat"
)

const (
	AUD CurrencyType = "AUD" //Australian Dollar
	HKD CurrencyType = "HKD" //Hong Kong Dollar
	SGD CurrencyType = "SGD" //Singapore Dollar
	BND CurrencyType = "BND" //Brunei Dollar
	INR CurrencyType = "INR" //Indian Rupee
	THB CurrencyType = "THB" //Thai Baht
	CAD CurrencyType = "CAD" //Canadian Dollar
	JPY CurrencyType = "JPY" //Japanese Yen
	TOP CurrencyType = "TOP" //Tongan Pa'anga
	CHF CurrencyType = "CHF" //Switzerland Franc
	KWD CurrencyType = "KWD" //Kuwait Dinar
	USD CurrencyType = "USD" //United States Dollar
	EUR CurrencyType = "EUR" //Euro
	MYR CurrencyType = "MYR" //Malaysian Ringgit
	VUV CurrencyType = "VUV" //Vanuatu Vatu
	FJD CurrencyType = "FJD" //Fiji Dollar
	NZD CurrencyType = "NZD" //New Zealand Dollar
	WST CurrencyType = "WST" //Samoan Tala
)

const (
	Internet ClientType = "I"
	//Recurring ClientType = "R" // NOT SUPPORTED
)
