package models
import (
    "tdex.com/futures/common/binary"
)

var NameOfPosition = []string{
"ID",
"UID",
"CID",
"Shared",
"Side",
"Price",
"Volume",
"Scale",
"Margin",
"InitMargin",
"Force",
"Fee",
"OpenFee",
"Completed",
"Notes",
"Bankrupt",
"SerialID",
"Weight",
"CloseWeight",
"RealisedPNL",
"CreatedAt",
"FundFee",
"OriginFee",
"CoinFee",
"Coin1Fee",
"Coin2Fee",
"ValueFee",
"ValueOriginFee",
}
const (
    // IndexOfPositionID is ID index of Position
    IndexOfPositionID uint8 = 0
    // IndexOfPositionUID is UID index of Position
    IndexOfPositionUID uint8 = 1
    // IndexOfPositionCID is CID index of Position
    IndexOfPositionCID uint8 = 2
    // IndexOfPositionShared is Shared index of Position
    IndexOfPositionShared uint8 = 3
    // IndexOfPositionSide is Side index of Position
    IndexOfPositionSide uint8 = 4
    // IndexOfPositionPrice is Price index of Position
    IndexOfPositionPrice uint8 = 5
    // IndexOfPositionVolume is Volume index of Position
    IndexOfPositionVolume uint8 = 6
    // IndexOfPositionScale is Scale index of Position
    IndexOfPositionScale uint8 = 7
    // IndexOfPositionMargin is Margin index of Position
    IndexOfPositionMargin uint8 = 8
    // IndexOfPositionInitMargin is InitMargin index of Position
    IndexOfPositionInitMargin uint8 = 9
    // IndexOfPositionForce is Force index of Position
    IndexOfPositionForce uint8 = 10
    // IndexOfPositionFee is Fee index of Position
    IndexOfPositionFee uint8 = 11
    // IndexOfPositionOpenFee is OpenFee index of Position
    IndexOfPositionOpenFee uint8 = 12
    // IndexOfPositionCompleted is Completed index of Position
    IndexOfPositionCompleted uint8 = 13
    // IndexOfPositionNotes is Notes index of Position
    IndexOfPositionNotes uint8 = 14
    // IndexOfPositionBankrupt is Bankrupt index of Position
    IndexOfPositionBankrupt uint8 = 15
    // IndexOfPositionSerialID is SerialID index of Position
    IndexOfPositionSerialID uint8 = 16
    // IndexOfPositionWeight is Weight index of Position
    IndexOfPositionWeight uint8 = 17
    // IndexOfPositionCloseWeight is CloseWeight index of Position
    IndexOfPositionCloseWeight uint8 = 18
    // IndexOfPositionRealisedPNL is RealisedPNL index of Position
    IndexOfPositionRealisedPNL uint8 = 19
    // IndexOfPositionCreatedAt is CreatedAt index of Position
    IndexOfPositionCreatedAt uint8 = 20
    // IndexOfPositionFundFee is FundFee index of Position
    IndexOfPositionFundFee uint8 = 21
    // IndexOfPositionOriginFee is OriginFee index of Position
    IndexOfPositionOriginFee uint8 = 22
    // IndexOfPositionCoinFee is CoinFee index of Position
    IndexOfPositionCoinFee uint8 = 23
    // IndexOfPositionCoin1Fee is Coin1Fee index of Position
    IndexOfPositionCoin1Fee uint8 = 24
    // IndexOfPositionCoin2Fee is Coin2Fee index of Position
    IndexOfPositionCoin2Fee uint8 = 25
    // IndexOfPositionValueFee is ValueFee index of Position
    IndexOfPositionValueFee uint8 = 26
    // IndexOfPositionValueOriginFee is ValueOriginFee index of Position
    IndexOfPositionValueOriginFee uint8 = 27
)
// Encode 编码
func (x Position) Encode(w *binary.Encoder) {
	w.Uint(x.ID) 
	w.Uint(x.UID) 
	w.Uint(x.CID) 
	w.Boolean(x.Shared) 
	w.Uint8(uint8(x.Side)) 
	w.Float64(x.Price) 
	w.Float64(x.Volume) 
	w.Float64(x.Scale) 
	w.Float64(x.Margin) 
	w.Float64(x.InitMargin) 
	w.Float64(x.Force) 
	w.Float64(x.Fee) 
	w.Float64(x.OpenFee) 
	w.Boolean(x.Completed) 
	w.String(x.Notes) 
	w.Boolean(x.Bankrupt) 
	w.Int(x.SerialID) 
	w.Float64(x.Weight) 
	w.Float64(x.CloseWeight) 
	w.Float64(x.RealisedPNL) 
	w.Int(x.CreatedAt) 
	w.Float64(x.FundFee) 
	w.Float64(x.OriginFee) 
	w.Float64(x.CoinFee) 
	w.Float64(x.Coin1Fee) 
	w.Float64(x.Coin2Fee) 
	w.Float64(x.ValueFee) 
	w.Float64(x.ValueOriginFee) 
}
// Decode 解码
func (x *Position) Decode(r *binary.Decoder) {
	x.ID = r.Uint()
	x.UID = r.Uint()
	x.CID = r.Uint()
	x.Shared = r.Boolean()
	x.Side = Side(r.Uint8())
	x.Price = r.Float64() 
	x.Volume = r.Float64() 
	x.Scale = r.Float64() 
	x.Margin = r.Float64() 
	x.InitMargin = r.Float64() 
	x.Force = r.Float64() 
	x.Fee = r.Float64() 
	x.OpenFee = r.Float64() 
	x.Completed = r.Boolean()
	x.Notes = r.String()
	x.Bankrupt = r.Boolean()
	x.SerialID = r.Int() 
	x.Weight = r.Float64() 
	x.CloseWeight = r.Float64() 
	x.RealisedPNL = r.Float64() 
	x.CreatedAt = r.Int() 
	x.FundFee = r.Float64() 
	x.OriginFee = r.Float64() 
	x.CoinFee = r.Float64() 
	x.Coin1Fee = r.Float64() 
	x.Coin2Fee = r.Float64() 
	x.ValueFee = r.Float64() 
	x.ValueOriginFee = r.Float64() 
}

var NameOfPositionV1 = []string{
"ID",
"UID",
"CID",
"Shared",
"Side",
"Price",
"Volume",
"Scale",
"Margin",
"InitMargin",
"Force",
"Fee",
"OpenFee",
"Completed",
"Notes",
"Bankrupt",
"SerialID",
"Weight",
"CloseWeight",
"RealisedPNL",
"CreatedAt",
}
const (
    // IndexOfPositionV1ID is ID index of PositionV1
    IndexOfPositionV1ID uint8 = 0
    // IndexOfPositionV1UID is UID index of PositionV1
    IndexOfPositionV1UID uint8 = 1
    // IndexOfPositionV1CID is CID index of PositionV1
    IndexOfPositionV1CID uint8 = 2
    // IndexOfPositionV1Shared is Shared index of PositionV1
    IndexOfPositionV1Shared uint8 = 3
    // IndexOfPositionV1Side is Side index of PositionV1
    IndexOfPositionV1Side uint8 = 4
    // IndexOfPositionV1Price is Price index of PositionV1
    IndexOfPositionV1Price uint8 = 5
    // IndexOfPositionV1Volume is Volume index of PositionV1
    IndexOfPositionV1Volume uint8 = 6
    // IndexOfPositionV1Scale is Scale index of PositionV1
    IndexOfPositionV1Scale uint8 = 7
    // IndexOfPositionV1Margin is Margin index of PositionV1
    IndexOfPositionV1Margin uint8 = 8
    // IndexOfPositionV1InitMargin is InitMargin index of PositionV1
    IndexOfPositionV1InitMargin uint8 = 9
    // IndexOfPositionV1Force is Force index of PositionV1
    IndexOfPositionV1Force uint8 = 10
    // IndexOfPositionV1Fee is Fee index of PositionV1
    IndexOfPositionV1Fee uint8 = 11
    // IndexOfPositionV1OpenFee is OpenFee index of PositionV1
    IndexOfPositionV1OpenFee uint8 = 12
    // IndexOfPositionV1Completed is Completed index of PositionV1
    IndexOfPositionV1Completed uint8 = 13
    // IndexOfPositionV1Notes is Notes index of PositionV1
    IndexOfPositionV1Notes uint8 = 14
    // IndexOfPositionV1Bankrupt is Bankrupt index of PositionV1
    IndexOfPositionV1Bankrupt uint8 = 15
    // IndexOfPositionV1SerialID is SerialID index of PositionV1
    IndexOfPositionV1SerialID uint8 = 16
    // IndexOfPositionV1Weight is Weight index of PositionV1
    IndexOfPositionV1Weight uint8 = 17
    // IndexOfPositionV1CloseWeight is CloseWeight index of PositionV1
    IndexOfPositionV1CloseWeight uint8 = 18
    // IndexOfPositionV1RealisedPNL is RealisedPNL index of PositionV1
    IndexOfPositionV1RealisedPNL uint8 = 19
    // IndexOfPositionV1CreatedAt is CreatedAt index of PositionV1
    IndexOfPositionV1CreatedAt uint8 = 20
)
// Encode 编码
func (x PositionV1) Encode(w *binary.Encoder) {
	w.Uint(x.ID) 
	w.Uint(x.UID) 
	w.Uint(x.CID) 
	w.Boolean(x.Shared) 
	w.Uint8(uint8(x.Side)) 
	w.Float64(x.Price) 
	w.Float64(x.Volume) 
	w.Float64(x.Scale) 
	w.Float64(x.Margin) 
	w.Float64(x.InitMargin) 
	w.Float64(x.Force) 
	w.Float64(x.Fee) 
	w.Float64(x.OpenFee) 
	w.Boolean(x.Completed) 
	w.String(x.Notes) 
	w.Boolean(x.Bankrupt) 
	w.Int(x.SerialID) 
	w.Float64(x.Weight) 
	w.Float64(x.CloseWeight) 
	w.Float64(x.RealisedPNL) 
	w.Int(x.CreatedAt) 
}
// Decode 解码
func (x *PositionV1) Decode(r *binary.Decoder) {
	x.ID = r.Uint()
	x.UID = r.Uint()
	x.CID = r.Uint()
	x.Shared = r.Boolean()
	x.Side = Side(r.Uint8())
	x.Price = r.Float64() 
	x.Volume = r.Float64() 
	x.Scale = r.Float64() 
	x.Margin = r.Float64() 
	x.InitMargin = r.Float64() 
	x.Force = r.Float64() 
	x.Fee = r.Float64() 
	x.OpenFee = r.Float64() 
	x.Completed = r.Boolean()
	x.Notes = r.String()
	x.Bankrupt = r.Boolean()
	x.SerialID = r.Int() 
	x.Weight = r.Float64() 
	x.CloseWeight = r.Float64() 
	x.RealisedPNL = r.Float64() 
	x.CreatedAt = r.Int() 
}
