package common

import (
	"log"
	"testing"
)

func TestSha3(t *testing.T) {
	acc := "000000000000000000000000391694e7e0b0cce554cb130d723a9d27458f92980000000000000000000000000000000000000000000000000000000000000001"
	log.Print(Sha3(acc))

	log.Print(GetPosition("0x391694e7e0b0cce554cb130d723a9d27458f9298"))
}

func TestH2I(t *testing.T) {
	log.Print(HtoI("7f5dd"))
}

func TestI2H(t *testing.T) {
	log.Print(ItoH(5489301))
	log.Print(HtoI("0x000000000053c295"))
}
