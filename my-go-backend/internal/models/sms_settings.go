package models

type SmsProvider string

const (
	ProviderFarazSMS    SmsProvider = "farazsms"
	ProviderMeliPayamak SmsProvider = "melipayamak"
)

type SmsSettings struct {
	Provider     SmsProvider `json:"provider"`
	FarazApiKey  string      `json:"faraz_api_key"`
	FarazPattern string      `json:"faraz_pattern"`
	MeliUsername string      `json:"meli_username"`
	MeliPassword string      `json:"meli_password"`
	MeliFrom     string      `json:"meli_from"`
	MeliPattern  string      `json:"meli_pattern"`
}
