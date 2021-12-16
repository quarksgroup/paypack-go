package api

import "strings"

// Provider defines an upstream payment channel(telco)
type provider string

const (
	mtn    provider = `momo-mtn-rw`
	airtel provider = `momo-airtel-rw`
	uknown provider = "unknown"
)

func detectProvider(phone string) provider {
	if strings.HasPrefix(phone, "078") {
		return mtn
	}
	if strings.HasPrefix(phone, "079") {
		return mtn
	}
	if strings.HasPrefix(phone, "073") {
		return airtel
	}
	if strings.HasPrefix(phone, "072") {
		return airtel
	}
	if strings.HasPrefix(phone, "+25078") {
		return mtn
	}
	if strings.HasPrefix(phone, "+25079") {
		return mtn
	}
	if strings.HasPrefix(phone, "+25073") {
		return airtel
	}
	if strings.HasPrefix(phone, "+25072") {
		return airtel
	}
	return uknown
}
