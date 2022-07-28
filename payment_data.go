package applepay

type PaymentDataHeader struct {
	ApplicationData    string `json:"applicationData"`
	EphemeralPublicKey []byte `json:"ephemeralPublicKey"`
	WrappedKey         []byte `json:"wrappedKey"`
	PublicKeyHash      []byte `json:"publicKeyHash"`
	TransactionID      string `json:"transactionId"`
}

// PaymentData is the payment information returned by Apple Pay with all data
// See https://developer.apple.com/documentation/passkit/apple_pay/payment_token_format_reference
type PaymentData struct {
	Header    PaymentDataHeader `json:"header"`
	Data      []byte            `json:"data"`
	Signature []byte            `json:"signature"`
	Version   string            `json:"version"`
}
