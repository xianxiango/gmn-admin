package abi

import (
	"log"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

func TestAbiDecodeMulti(t *testing.T) {
	//96 -64=32 448/64=7 10
	hex := "0xad8733ca00000000000000000000000066c36272a1609c582b4211becadf9dba5b913893000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000619046d59f45e15fb4ee3b6e1b671ae0d5d9ac8c000000000000000000000000f42fd7d53c62311b2e9d00f148f395b593336ec700000000000000000000000020907917c862ecf893110e2216c725c332dae5d900000000000000000000000003c26d5e22c0d9508281cdfcd8594ddb508ec3370000000000000000000000007ba44ae622b800d6c0c4f04fdc3478d9e2224e1a000000000000000000000000c6fc6835ce16679c89567c0b33604d1e9a71bf99000000000000000000000000f82366646c061e4a472b6632e344d787bcbb9c5a000000000000000000000000501dcc1f9fedb2eadc7a43c65d99ecb27fa3b1f20000000000000000000000008dd6fb6886711611655b358e3ff0c47c65ae198b0000000000000000000000001318c6835895381a7f7d9f23b5ee6c06297941f1000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000023000000000000000000000000000000000000000000000000000000000000002300000000000000000000000000000000000000000000000000000000000000230000000000000000000000000000000000000000000000000000000000000023000000000000000000000000000000000000000000000000000000000000002300000000000000000000000000000000000000000000000000000000000000230000000000000000000000000000000000000000000000000000000000000023000000000000000000000000000000000000000000000000000000000000002300000000000000000000000000000000000000000000000000000000000000230000000000000000000000000000000000000000000000000000000000000023"
	abiDecoder := NewAbi(hex)
	abiInfos := abiDecoder.Decode()

	for _, v := range abiInfos {
		log.Print(v.Address, "  amount : ", v.Amount)
	}

}

func TestAbiDecodeSingle(t *testing.T) {
	hex := "0xa9059cbb0000000000000000000000008a816e35ba9dd138bc4f17fd249b0397fbd9e88b0000000000000000000000000000000000000000000000008ac7230489e80000"
	abiDecoder := NewAbi(hex)
	abiInfos := abiDecoder.Decode()

	for _, v := range abiInfos {
		log.Print(v.Address, "   amount : ", v.Amount)
	}

	encode := BuildSingleInputData(abiInfos[0].Address, abiInfos[0].Amount)
	log.Print(encode)

}

func TestCompilerFuture(t *testing.T) {
	cp := NewCompiler()
	cp.SolidityCompile("./cgt-wallet.sol")
	log.Print(cp.GetContractsNames())
	log.Print(cp.GetDeployData("Wallet", "",
		[]common.Address{common.HexToAddress("0x05eD09Dbe6ccD01Ab5832CB49b216e8fd91279aA")},
		big.NewInt(1),
		big.NewInt(10000000000000),
		big.NewInt(10000000000000),
		"CRYPTO GAME TOKEN",
		"CCGGTT",
		uint8(6)))
}

func TestCompilerSimpleCGT(t *testing.T) {
	cp := NewCompiler()
	cp.SolidityCompile("./cgt.sol")
	log.Print(cp.GetContractsNames())

	log.Print(cp.GetDeployData("CGT", ""))
	_, abiData, _ := cp.GetComipleData("CGT")
	log.Print(abiData)
}

func TestConvert(t *testing.T) {
	bigInt, ok := math.ParseBig256("0x2ca84e219b25000")

	if ok {
		log.Print(bigInt.Int64())
		n := float64(bigInt.Int64() / 1000000000)
		log.Print(float64(n) / 1000000000)
	}
}

func TestBuildRandomAddress(t *testing.T) {

	log.Print("0xB2e5679b19A2c7139d1F70dac6FbAF39e875d625")
	fd, _ := os.Create("address.txt")
	for i := 0; i < 200; i++ {
		time.Sleep(5 * time.Microsecond)
		str := fmt.Sprintf(`"0x%s",`, buildRandom()[:40])
		log.Print(str)
		fd.WriteString(str)
	}
	fd.Close()
}

func buildRandom() string {
	var letterRunes = []rune("1234567890abcdef")
	rand.Seed(time.Now().UnixNano())

	str := ""
	for i := 0; i < 50; i++ {

		str += fmt.Sprintf("%v", letterRunes[rand.Intn(len(letterRunes))])
	}
	return str
}

func TestOmniDecode(t *testing.T) {
	script := NewOmniScript("6a146f6d6e69000000000000001f00000006e1eb630e")
	log.Print(script.Decode())
	log.Print(script)
	log.Print(script.GetAmount())
	log.Print(script.GetTokenNo())
}
