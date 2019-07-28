package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func Sha3(acc string) string {
	hash := sha3.NewKeccak256()
	hash.Write(decodeHex(acc))

	var buf []byte
	buf = hash.Sum(buf)

	s := hex.EncodeToString(buf)

	return "0x" + s
}

func GetPosition(addr string) string {
	address := strings.TrimPrefix(addr, "0x")
	address = "000000000000000000000000" + address + "0000000000000000000000000000000000000000000000000000000000000000"
	return Sha3(address)
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
