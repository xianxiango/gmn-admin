package models
import (
    "tdex.com/futures/common/binary"
)

var NameOfDelegate = []string{
"UID",
"ID",
"Side",
"Price",
"Volume",
"Visible",
"Deadline",
"SerialID",
"Activated",
"Merged",
"Attempt",
"Scale",
"Better",
"Passive",
}
const (
    // IndexOfDelegateUID is UID index of Delegate
    IndexOfDelegateUID uint8 = 0
    // IndexOfDelegateID is ID index of Delegate
    IndexOfDelegateID uint8 = 1
    // IndexOfDelegateSide is Side index of Delegate
    IndexOfDelegateSide uint8 = 2
    // IndexOfDelegatePrice is Price index of Delegate
    IndexOfDelegatePrice uint8 = 3
    // IndexOfDelegateVolume is Volume index of Delegate
    IndexOfDelegateVolume uint8 = 4
    // IndexOfDelegateVisible is Visible index of Delegate
    IndexOfDelegateVisible uint8 = 5
    // IndexOfDelegateDeadline is Deadline index of Delegate
    IndexOfDelegateDeadline uint8 = 6
    // IndexOfDelegateSerialID is SerialID index of Delegate
    IndexOfDelegateSerialID uint8 = 7
    // IndexOfDelegateActivated is Activated index of Delegate
    IndexOfDelegateActivated uint8 = 8
    // IndexOfDelegateMerged is Merged index of Delegate
    IndexOfDelegateMerged uint8 = 9
    // IndexOfDelegateAttempt is Attempt index of Delegate
    IndexOfDelegateAttempt uint8 = 10
    // IndexOfDelegateScale is Scale index of Delegate
    IndexOfDelegateScale uint8 = 11
    // IndexOfDelegateBetter is Better index of Delegate
    IndexOfDelegateBetter uint8 = 12
    // IndexOfDelegatePassive is Passive index of Delegate
    IndexOfDelegatePassive uint8 = 13
)
// Encode 编码
func (x Delegate) Encode(w *binary.Encoder) {
	w.Uint(x.UID) 
	w.Uint(x.ID) 
	w.Uint8(uint8(x.Side)) 
	w.Float64(x.Price) 
	w.Float64(x.Volume) 
	w.Float64(x.Visible) 
	w.Int(x.Deadline) 
	w.Int(x.SerialID) 
	w.Int(x.Activated) 
	w.Boolean(x.Merged) 
	w.Uint8(uint8(x.Attempt)) 
	w.Float64(x.Scale) 
	w.Boolean(x.Better) 
	w.Boolean(x.Passive) 
}
// Decode 解码
func (x *Delegate) Decode(r *binary.Decoder) {
	x.UID = r.Uint()
	x.ID = r.Uint()
	x.Side = Side(r.Uint8())
	x.Price = r.Float64() 
	x.Volume = r.Float64() 
	x.Visible = r.Float64() 
	x.Deadline = r.Int() 
	x.SerialID = r.Int() 
	x.Activated = r.Int() 
	x.Merged = r.Boolean()
	x.Attempt = Attempt(r.Uint8())
	x.Scale = r.Float64() 
	x.Better = r.Boolean()
	x.Passive = r.Boolean()
}

var NameOfQuote = []string{
"Price",
"Volume",
}
const (
    // IndexOfQuotePrice is Price index of Quote
    IndexOfQuotePrice uint8 = 0
    // IndexOfQuoteVolume is Volume index of Quote
    IndexOfQuoteVolume uint8 = 1
)
// Encode 编码
func (x Quote) Encode(w *binary.Encoder) {
	w.Float64(x.Price) 
	w.Float64(x.Volume) 
}
// Decode 解码
func (x *Quote) Decode(r *binary.Decoder) {
	x.Price = r.Float64() 
	x.Volume = r.Float64() 
}
