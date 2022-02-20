package utils

import "fmt"

type UtilToken struct {
}

type DemoTokenInfo struct {
	IsOk       bool
	Prefix     string
	CurrencyId string
	Uui        string
	Error      string
}

type TokenInfo struct {
	IsOk       bool
	Prefix     string
	CurrencyId string
	PlayerId   string
	Error      string
}

func (o UtilToken) BuildDemo(prefix string, currencyId string, separator string) string {
	token := fmt.Sprintf("%s%s%s%s%s", prefix, separator, currencyId, separator, Util{}.Tracev4())

	return token
}

func (o UtilToken) ParseDemo(token string, separator string) DemoTokenInfo {
	var strUtil UtilString

	if token == "" {
		return DemoTokenInfo{IsOk: false, Error: "Invalid Token"}
	}

	tokenitems := strUtil.Split(token, separator)

	if len(tokenitems) != 3 {
		return DemoTokenInfo{IsOk: false, Error: "Mismatch token length"}
	}

	return DemoTokenInfo{
		IsOk:       true,
		Prefix:     tokenitems[0],
		CurrencyId: tokenitems[1],
		Uui:        tokenitems[2],
	}
}

func (o UtilToken) Build(prefix string, currencyId string, playerId string, separator string) string {
	token := fmt.Sprintf("%s%s%s%s%s", prefix, separator, currencyId, separator, playerId)

	return token
}

func (o UtilToken) Parse(token string, separator string) TokenInfo {
	var strUtil UtilString

	if token == "" {
		return TokenInfo{IsOk: false, Error: "Invalid Token"}
	}

	tokenitems := strUtil.Split(token, separator)

	if len(tokenitems) != 3 {
		return TokenInfo{IsOk: false, Error: "Mismatch token length"}
	}

	return TokenInfo{
		IsOk:       true,
		Prefix:     tokenitems[0],
		CurrencyId: tokenitems[1],
		PlayerId:   tokenitems[2],
	}
}
