package utils

import (
	"golang.org/x/net/publicsuffix"

	"github.com/asciimoo/privacyscore/penalty"
)

func CropSubdomains(domain string) string {
	host, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return host
}

func GetScoreName(s penalty.Score) string {
	switch {
	case s >= 80:
		return "good"
	case s >= 60:
		return "medium"
	case s >= 40:
		return "bad"
	}
	return "poor"
}
