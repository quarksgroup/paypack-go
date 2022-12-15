package paypack

//Merchant represent transaformed merchant profile information
type Merchant struct {
	ID            string  `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	InRate        float64 `json:"in_rate,omitempty"`
	OutRate       float64 `json:"out_rate,omitempty"`
	AirtelInRate  float64 `json:"airtel_in_rate,omitempty"`
	AirtelOutRate float64 `json:"airtel_out_rate,omitempty"`
	Balance       float64 `json:"balance,omitempty"`
	AirtelBalance float64 `json:"airtel_balance,omitempty"`
	MtnBalance    float64 `json:"mtn_balance,omitempty"`
}

//Checkout represent checkout information
type Checkout struct {
	ID           string `json:"id,omitempty"`
	Merchant     string `json:"merchant,omitempty"`
	Name         string `json:"name,omitempty"`
	Logo         string `json:"logo,omitempty"`
	Email        string `json:"email,omitempty"` //This will be support email that will be sent to customer email for reply
	SendEmail    bool   `json:"send_email"`      // This will represent if merchant need to send email when payments successed
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	CancelUrl    string `json:"cancel_url,omitempty"`
	SuccessUrl   string `json:"success_url,omitempty"`
}
