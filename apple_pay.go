package applepay

type DataType string

const (
	SecureDataType DataType = "3DSecure"
	EMVDataType    DataType = "EMV"
)

type PaymentData3DSecure struct {
	PaymentCryptogram string
	EciIndicator      string
}

type PaymentDataEMV struct {
	Data             string
	EncryptedPINData string
}

type ApplePay[T PaymentData3DSecure | PaymentDataEMV] struct {
	AccountNumber     string
	ExpiryMonth       int
	ExpiryYear        int
	CurrencyCode      string
	TransactionAmount int64
	CardholderName    string
	Identifier        string
	PaymentDataType   DataType
	PaymentData       T
}
