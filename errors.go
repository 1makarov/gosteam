package gosteam

import (
	"errors"
	"fmt"
)

var (
	errorRSA           = errors.New("error generate rsakey")
	errorNeedTwoFactor = errors.New("invalid twofactor code")
	errorApiKey        = errors.New("invalid apikey")
)

func errorStatusCode(functionname string, statuscode int) error {
	return errors.New(fmt.Sprintf("%s | %d", functionname, statuscode))
}

func errorText(text string) error {
	return errors.New(text)
}
