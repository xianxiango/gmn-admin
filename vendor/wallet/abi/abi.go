package abi

import (
	"encoding/binary"
	"encoding/hex"
	"log"

	"strings"
)

const unkonw = "0x02e31512"
const single = "0xa9059cbb"
const multible = "0xad8733ca"

type AbiInfo struct {
	Address string
	Amount  string
}

type ApiDecoder struct {
	data    string
	index   int
	method  string
	address string
}

func NewAbi(d string) *ApiDecoder {
	return &ApiDecoder{data: d}
}

func (c *ApiDecoder) Decode() []*AbiInfo {
	if len(c.data) < 138 {
		return []*AbiInfo{}
	}
	header := 10
	bit := 64
	c.method = c.data[:header]
	if c.method == single {
		//单个发送
		address := "0x" + c.data[header : header+bit][24:]
		amount := "0x" + c.data[header+bit:]
		i := &AbiInfo{Address: address, Amount: amount}
		return []*AbiInfo{i}

	} else if c.method == multible {
		//批量发送
		//计算Token 地址
		c.address = "0x" + c.data[header : header+bit][24:] //token address
		offset := bit*3 + header
		//偏移计算共给多小个地址发送
		count := "0x" + c.data[offset : offset+bit][48:] // count of address
		length := ParseInt64(count)
		//计算资金地址块
		amountIndex := int64(offset) + int64(bit)*(length+1)
		amountDataBlock := c.data[amountIndex:]
		//循环构造信息
		var infos = []*AbiInfo{}
		addrOffset := offset + bit
		for i := int64(0); i < length; i++ {
			addr := "0x" + c.data[addrOffset : addrOffset+bit][24:]

			addrOffset += bit
			amountHex := "0x" + c.getAmount(amountDataBlock, i)

			abiInfo := &AbiInfo{Amount: amountHex, Address: addr}
			infos = append(infos, abiInfo)
		}
		return infos

	} //done
	return []*AbiInfo{}
}

func (c *ApiDecoder) getAmount(dataBlock string, index int64) string {
	//log.Print(dataBlock)
	length := int64(64)
	bit := int64(64)
	start := length + index*bit
	return dataBlock[start : start+bit]
}

func ParseInt64(hexStr string) int64 {
	trimStr := strings.TrimPrefix(hexStr, "0x")
	byt, _ := hex.DecodeString(trimStr)
	return int64(binary.BigEndian.Uint64(byt))

}

func BuildPrefix(n int) string {
	prefix := ""
	for i := 0; i < n; i++ {
		prefix += "0"
	}
	return prefix
}

func BuildSingleInputData(_to, _amount string) string {
	toAddr := strings.TrimPrefix(_to, "0x")
	amount := strings.TrimPrefix(_amount, "0x")
	padZero := 64 - len(amount)
	log.Print(_amount)
	log.Print(amount)
	return single + BuildPrefix(24) + toAddr + BuildPrefix(padZero) + amount
}

/*
func BuildMultibleInputData(_contractsAddress string, _abis []*AbiInfo) string {
	contractsAddress := strings.TrimPrefix(_contractsAddress, "0x")
	return multible + BuildPrefix(24) + contractsAddress
}
*/
