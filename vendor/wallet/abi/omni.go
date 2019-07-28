package abi

import (
	"errors"
	"strings"
	"wallet/common"
)

var demo = "6a146f6d6e6900000000000000030000000000000035"

type OmniScript struct {
	hex     string
	amount  string
	tokenNo string
}

func NewOmniScript(scriptHex string) *OmniScript {
	return &OmniScript{
		hex: scriptHex,
	}
}

func (c *OmniScript) Decode() error {
	hex := c.hex
	if len(hex) >= len(demo) {
		if strings.ToUpper(hex[:12]) == strings.ToUpper(demo[:12]) {
			tokenNo := hex[16:28]
			amount := hex[28:]
			c.amount = amount
			c.tokenNo = tokenNo
			return nil
		} else {
			return errors.New("header is not an Omni protocol")
		}
	} // end length compare
	return errors.New("length is not an Omni protocol")
}

func (c *OmniScript) GetAmount() int64 {
	return common.HtoI(c.amount)
}

func (c *OmniScript) GetTokenNo() int {
	return int(common.HtoI(c.tokenNo))
}
