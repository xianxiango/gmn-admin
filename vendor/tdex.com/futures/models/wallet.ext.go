package models
import (
    "tdex.com/futures/common/binary"
)

var NameOfWallet = []string{
"UID",
"Coin",
"Balance",
"Margin",
}
const (
    // IndexOfWalletUID is UID index of Wallet
    IndexOfWalletUID uint8 = 0
    // IndexOfWalletCoin is Coin index of Wallet
    IndexOfWalletCoin uint8 = 1
    // IndexOfWalletBalance is Balance index of Wallet
    IndexOfWalletBalance uint8 = 2
    // IndexOfWalletMargin is Margin index of Wallet
    IndexOfWalletMargin uint8 = 3
)
// Encode 编码
func (x Wallet) Encode(w *binary.Encoder) {
	w.Uint(x.UID) 
	w.Uint(x.Coin) 
	w.Float64(x.Balance) 
	w.Float64(x.Margin) 
}
// Decode 解码
func (x *Wallet) Decode(r *binary.Decoder) {
	x.UID = r.Uint()
	x.Coin = r.Uint()
	x.Balance = r.Float64() 
	x.Margin = r.Float64() 
}
