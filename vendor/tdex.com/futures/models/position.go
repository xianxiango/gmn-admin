package models

// Position 仓位
type Position struct {
	ID             uint    `xorm:"ID"`           // 编号
	UID            uint    `xorm:"UID"`          // 用户
	CID            uint    `xorm:"CID"`          // 产品
	Shared         bool    `xorm:"Shared"`       // 全仓|逐仓
	Side           Side    `xorm:"Side"`         // 持仓方向
	Price          float64 `xorm:"Price"`        // 均价
	Volume         float64 `xorm:"Volume"`       // 数量
	Scale          float64 `xorm:"Scale"`        // 杠杆
	Margin         float64 `xorm:"Margin"`       // 保证金
	InitMargin     float64 `xorm:"InitMargin"`   // 初始保证金
	Force          float64 `xorm:"Force"`        // 强平价格
	Fee            float64 `xorm:"Fee"`          // 内存费用
	OpenFee        float64 `xorm:"OpenFee"`      // 开仓费用
	Completed      bool    `xorm:"Completed"`    // 完整性（没有关联的开仓单）
	Notes          string  `xorm:"Notes"`        // 备注
	Bankrupt       bool    `xorm:"Bankrupt"`     // 破产
	SerialID       int     `xorm:"SerialID"`     // 序号（变化序号）
	Weight         float64 `xorm:"Weight"`       // 辅算金额：用于均价计算
	CloseWeight    float64 `xorm:"CloseWeight"`  // 平仓权重：用于计算平仓均价
	RealisedPNL    float64 `xorm:"RealisedPNL"`  // 已实现盈亏
	CreatedAt      int     `xorm:"CreatedAt1"`   // 创建时间
	FundFee        float64 `xorm:FundFee`        // 资金费用
	OriginFee      float64 `xorm:OriginFee`      // 成交 产品币种原始费用
	CoinFee        float64 `xorm:CoinFee`        // 成交 产品币种费用
	Coin1Fee       float64 `xorm:Coin1Fee`       // 成交 抵扣币种1费用
	Coin2Fee       float64 `xorm:Coin2Fee`       // 成交 抵扣币种2费用
	ValueFee       float64 `xorm:ValueFee`       // 成交等价 费用
	ValueOriginFee float64 `xorm:ValueOriginFee` // 成交等价 原始费用
}

func (t *Position) TableName() string {
	return "position"
}

// Upgrade 升级
func (x *Position) Upgrade(e interface{}) {
	switch a := e.(type) {
	case PositionV1:
		x.ID = a.ID
		x.UID = a.UID
		x.CID = a.CID
		x.Shared = a.Shared
		x.Side = a.Side
		x.Price = a.Price
		x.Volume = a.Volume
		x.Scale = a.Scale
		x.Margin = a.Margin
		x.InitMargin = a.InitMargin
		x.Force = a.Force
		x.Fee = a.Fee
		x.OriginFee = a.Fee
		x.CoinFee = a.Fee
		x.Coin1Fee = 0.0
		x.Coin2Fee = 0.0
		x.ValueFee = 0.0
		x.ValueOriginFee = 0.0
		x.FundFee = 0.0
		x.OpenFee = a.OpenFee
		x.Completed = a.Completed
		x.Notes = a.Notes
		x.Bankrupt = a.Bankrupt
		x.SerialID = a.SerialID
		x.Weight = a.Weight
		x.CloseWeight = a.CloseWeight
		x.RealisedPNL = a.RealisedPNL
		x.CreatedAt = a.CreatedAt
	}
}

// PositionV1 仓位
type PositionV1 struct {
	ID          uint    // 编号
	UID         uint    // 用户
	CID         uint    // 产品
	Shared      bool    // 全仓|逐仓
	Side        Side    // 持仓方向
	Price       float64 // 均价
	Volume      float64 // 数量
	Scale       float64 // 杠杆
	Margin      float64 // 保证金
	InitMargin  float64 // 初始保证金
	Force       float64 // 强平价格
	Fee         float64 // 费用
	OpenFee     float64 // 开仓费用
	Completed   bool    // 完整性（没有关联的开仓单）
	Notes       string  // 备注
	Bankrupt    bool    // 破产
	SerialID    int     // 序号（变化序号）
	Weight      float64 // 辅算金额：用于均价计算
	CloseWeight float64 // 平仓权重：用于计算平仓均价
	RealisedPNL float64 // 已实现盈亏
	CreatedAt   int     // 创建时间
}
