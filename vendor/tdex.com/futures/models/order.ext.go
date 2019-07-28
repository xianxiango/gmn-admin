package models

import (
	"tdex.com/futures/common/binary"
)

var NameOfOrder = []string{
	"ID",
	"CID",
	"UID",
	"PID",
	"RID",
	"Attempt",
	"Shared",
	"Side",
	"Scale",
	"Kind",
	"Volume",
	"Distance",
	"Price",
	"Locked",
	"Limit",
	"State",
	"Passive",
	"Visible",
	"Better",
	"Timely",
	"TimelyParam",
	"Deadline",
	"Strategy",
	"Variable",
	"Constant",
	"Relative",
	"Vertex",
	"Filled",
	"SerialID",
	"Activated",
	"Triggered",
	"Margin",
	"InitMargin",
	"Weight",
	"Notes",
	"CreatedAt",
	"FID",
	"FVersion",
	"Version",
}

const (
	// IndexOfOrderID is ID index of Order
	IndexOfOrderID uint8 = 0
	// IndexOfOrderCID is CID index of Order
	IndexOfOrderCID uint8 = 1
	// IndexOfOrderUID is UID index of Order
	IndexOfOrderUID uint8 = 2
	// IndexOfOrderPID is PID index of Order
	IndexOfOrderPID uint8 = 3
	// IndexOfOrderRID is RID index of Order
	IndexOfOrderRID uint8 = 4
	// IndexOfOrderAttempt is Attempt index of Order
	IndexOfOrderAttempt uint8 = 5
	// IndexOfOrderShared is Shared index of Order
	IndexOfOrderShared uint8 = 6
	// IndexOfOrderSide is Side index of Order
	IndexOfOrderSide uint8 = 7
	// IndexOfOrderScale is Scale index of Order
	IndexOfOrderScale uint8 = 8
	// IndexOfOrderKind is Kind index of Order
	IndexOfOrderKind uint8 = 9
	// IndexOfOrderVolume is Volume index of Order
	IndexOfOrderVolume uint8 = 10
	// IndexOfOrderDistance is Distance index of Order
	IndexOfOrderDistance uint8 = 11
	// IndexOfOrderPrice is Price index of Order
	IndexOfOrderPrice uint8 = 12
	// IndexOfOrderLocked is Locked index of Order
	IndexOfOrderLocked uint8 = 13
	// IndexOfOrderLimit is Limit index of Order
	IndexOfOrderLimit uint8 = 14
	// IndexOfOrderState is State index of Order
	IndexOfOrderState uint8 = 15
	// IndexOfOrderPassive is Passive index of Order
	IndexOfOrderPassive uint8 = 16
	// IndexOfOrderVisible is Visible index of Order
	IndexOfOrderVisible uint8 = 17
	// IndexOfOrderBetter is Better index of Order
	IndexOfOrderBetter uint8 = 18
	// IndexOfOrderTimely is Timely index of Order
	IndexOfOrderTimely uint8 = 19
	// IndexOfOrderTimelyParam is TimelyParam index of Order
	IndexOfOrderTimelyParam uint8 = 20
	// IndexOfOrderDeadline is Deadline index of Order
	IndexOfOrderDeadline uint8 = 21
	// IndexOfOrderStrategy is Strategy index of Order
	IndexOfOrderStrategy uint8 = 22
	// IndexOfOrderVariable is Variable index of Order
	IndexOfOrderVariable uint8 = 23
	// IndexOfOrderConstant is Constant index of Order
	IndexOfOrderConstant uint8 = 24
	// IndexOfOrderRelative is Relative index of Order
	IndexOfOrderRelative uint8 = 25
	// IndexOfOrderVertex is Vertex index of Order
	IndexOfOrderVertex uint8 = 26
	// IndexOfOrderFilled is Filled index of Order
	IndexOfOrderFilled uint8 = 27
	// IndexOfOrderSerialID is SerialID index of Order
	IndexOfOrderSerialID uint8 = 28
	// IndexOfOrderActivated is Activated index of Order
	IndexOfOrderActivated uint8 = 29
	// IndexOfOrderTriggered is Triggered index of Order
	IndexOfOrderTriggered uint8 = 30
	// IndexOfOrderMargin is Margin index of Order
	IndexOfOrderMargin uint8 = 31
	// IndexOfOrderInitMargin is InitMargin index of Order
	IndexOfOrderInitMargin uint8 = 32
	// IndexOfOrderWeight is Weight index of Order
	IndexOfOrderWeight uint8 = 33
	// IndexOfOrderNotes is Notes index of Order
	IndexOfOrderNotes uint8 = 34
	// IndexOfOrderCreatedAt is CreatedAt index of Order
	IndexOfOrderCreatedAt uint8 = 35
	// IndexOfOrderFID is FID index of Order
	IndexOfOrderFID uint8 = 36
	// IndexOfOrderFVersion is FVersion index of Order
	IndexOfOrderFVersion uint8 = 37
	// IndexOfOrderVersion is Version index of Order
	IndexOfOrderVersion uint8 = 38
	// IndexOfUpdatedAt is Version index of Order
	IndexOfUpdatedAt uint8 = 39
)

// Encode 编码
func (x Order) Encode(w *binary.Encoder) {
	w.Uint(x.ID)
	w.Uint(x.CID)
	w.Uint(x.UID)
	w.Uint(x.PID)
	w.Uint(x.RID)
	w.Uint8(uint8(x.Attempt))
	w.Boolean(x.Shared)
	w.Uint8(uint8(x.Side))
	w.Float64(x.Scale)
	w.Uint8(uint8(x.Kind))
	w.Float64(x.Volume)
	w.Boolean(x.Distance)
	w.Float64(x.Price)
	w.Float64(x.Locked)
	w.Float64(x.Limit)
	w.Uint8(uint8(x.State))
	w.Boolean(x.Passive)
	w.Float64(x.Visible)
	w.Boolean(x.Better)
	w.Uint8(uint8(x.Timely))
	w.Int(x.TimelyParam)
	w.Int(x.Deadline)
	w.Uint8(uint8(x.Strategy))
	w.Uint8(uint8(x.Variable))
	w.Float64(x.Constant)
	w.Boolean(x.Relative)
	w.Float64(x.Vertex)
	w.Float64(x.Filled)
	w.Int(x.SerialID)
	w.Int(x.Activated)
	w.Int(x.Triggered)
	w.Float64(x.Margin)
	w.Float64(x.InitMargin)
	w.Float64(x.Weight)
	w.String(x.Notes)
	w.Int(x.CreatedAt)
	w.Uint(x.FID)
	w.Int(x.FVersion)
	w.Int(x.Version)
}

// Decode 解码
func (x *Order) Decode(r *binary.Decoder) {
	x.ID = r.Uint()
	x.CID = r.Uint()
	x.UID = r.Uint()
	x.PID = r.Uint()
	x.RID = r.Uint()
	x.Attempt = Attempt(r.Uint8())
	x.Shared = r.Boolean()
	x.Side = Side(r.Uint8())
	x.Scale = r.Float64()
	x.Kind = Kind(r.Uint8())
	x.Volume = r.Float64()
	x.Distance = r.Boolean()
	x.Price = r.Float64()
	x.Locked = r.Float64()
	x.Limit = r.Float64()
	x.State = State(r.Uint8())
	x.Passive = r.Boolean()
	x.Visible = r.Float64()
	x.Better = r.Boolean()
	x.Timely = Timely(r.Uint8())
	x.TimelyParam = r.Int()
	x.Deadline = r.Int()
	x.Strategy = Strategy(r.Uint8())
	x.Variable = Variable(r.Uint8())
	x.Constant = r.Float64()
	x.Relative = r.Boolean()
	x.Vertex = r.Float64()
	x.Filled = r.Float64()
	x.SerialID = r.Int()
	x.Activated = r.Int()
	x.Triggered = r.Int()
	x.Margin = r.Float64()
	x.InitMargin = r.Float64()
	x.Weight = r.Float64()
	x.Notes = r.String()
	x.CreatedAt = r.Int()
	x.FID = r.Uint()
	x.FVersion = r.Int()
	x.Version = r.Int()
}

var NameOfChange = []string{
	"ID",
	"Volume",
}

const (
	// IndexOfChangeID is ID index of Change
	IndexOfChangeID uint8 = 0
	// IndexOfChangeVolume is Volume index of Change
	IndexOfChangeVolume uint8 = 1
)

// Encode 编码
func (x Change) Encode(w *binary.Encoder) {
	w.Uint(x.ID)
	w.Float64(x.Volume)
}

// Decode 解码
func (x *Change) Decode(r *binary.Decoder) {
	x.ID = r.Uint()
	x.Volume = r.Float64()
}
