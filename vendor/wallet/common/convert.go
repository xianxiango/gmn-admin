package common

import (
	"encoding/binary"
	"encoding/hex"
	sysmath "math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/math"
)

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func HtoI(hexStr string) int64 {
	if hexStr == "" || hexStr == "0x" {
		return 0
	}
	if strings.Contains(hexStr, "0x") {
		return int64(math.MustParseUint64(hexStr))
	}
	return int64(math.MustParseUint64("0x" + hexStr))
}

func ItoH(i int64) string {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	r := strings.TrimLeft(hex.EncodeToString(b), "0")
	return "0x" + r
}

func ItoETH(i int64, pow int) string {
	a := big.NewInt(int64(1))
	b := big.NewInt(int64(sysmath.Pow10(pow)))
	c := big.NewInt(int64(i))
	return "0x" + hex.EncodeToString(a.Mul(b, c).Bytes())
}

func buildByte(n int) []byte {
	var bs = []byte{}
	for i := 0; i < n; i++ {
		bs = append(bs, 0)
	}
	return bs
}
