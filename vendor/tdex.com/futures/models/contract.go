package models

import (
	"encoding/json"

	"tdex.com/common/scan"
)

// Risk 危险
type Risk struct {
	Level float64 `yaml:"Level"`
	Value float64 `yaml:"Value"`
}

// Contract 合同
type Contract struct {
	ID               uint    `yaml:"ID"`               // 编号
	Code             string  `yaml:"Code"`             // 代码
	Coin             uint    `yaml:"Coin"`             // 货币
	MPower           int     `yaml:"MPower"`           // 货币精度
	PPower           int     `yaml:"PPower"`           // 价格精度
	SPower           int     `yaml:"SPower"`           // 杠杆精度
	IPower           int     `yaml:"IPower"`           // 标记价|指数价精度
	JDigits          int     `yaml:"JDigits"`          // 价格精度
	Multiple         float64 `yaml:"Multiple"`         // 最高杠杆
	Risks            []Risk  `yaml:"Risks"`            // 风险系数 设置提交爆仓
	MakerFee         float64 `yaml:"MakerFee"`         // 挂单人佣金 <0 返佣
	TakerFee         float64 `yaml:"TakerFee"`         // 接盘侠佣金
	Scale            float64 `yaml:"Scale"`            // 合约规模
	UnitPrice        float64 `yaml:"UnitPrice"`        // 价格单位
	LastPrice        float64 `yaml:"-"`                // 最新价格
	MarkPrice        float64 `yaml:"-"`                // 标记价格
	AskPrice         float64 `yaml:"-"`                // 卖一价
	BidPrice         float64 `yaml:"-"`                // 买一价
	Fund             float64 `yaml:"Fund"`             // 保险基金
	AvailableFund    float64 `yaml:"AvailableFund"`    // 可用保险基金
	TaxRate          float64 `yaml:"TaxRate"`          // 资产费率
	SharedRisk       float64 `yaml:"SharedRisk"`       // 全仓风险系数(可用资金/已占用保证金<10%爆仓)
	AllowShared      bool    `yaml:"AllowShared"`      // 产品支持的全仓|逐仓模式
	AllowMerge       bool    `yaml:"AllowMerge"`       // 允许用户修改自动合仓设置
	MaxVolume        float64 `yaml:"MaxVolume"`        // 单次开仓数量上限
	MinVolume        float64 `yaml:"MinVolume"`        // 单次开仓数量下限
	Cushion          float64 `yaml:"Cushion"`          // 全仓有仓位或订单时，系统提现或调整杠杆需要预留以防暴仓
	Shared           bool    `yaml:"Shared"`           // 默认全仓模式
	Merged           bool    `yaml:"Merged"`           // 默认自动合仓
	Closed           bool    `yaml:"Closed"`           // 休市
	MaxPrice         float64 `yaml:"MaxPrice"`         // 最高价
	Waterline        float64 `yaml:"Waterline"`        // 爆仓后的平仓价按当前一价百分率
	ForceKeepAddtion float64 `yaml:"ForceKeepAddtion"` // 调整强平价时需要保留维持保证金率的增量
	LimitSecurity    float64 `yaml:"LimitSecurity"`    // 限价单安全系数
	Coin1            uint    `yaml:"Coin1"`            // 抵扣第一币种ID
	Coin2            uint    `yaml:"Coin2"`            // 抵扣第二币种ID
}

// Decode 解码
func (x *Contract) Decode(data map[string]string) {
	if v, ok := data[`MPower`]; ok {
		scan.Scan(scan.Pack(v), &x.MPower)
	}
	if v, ok := data[`PPower`]; ok {
		scan.Scan(scan.Pack(v), &x.PPower)
	}
	if v, ok := data[`SPower`]; ok {
		scan.Scan(scan.Pack(v), &x.SPower)
	}
	if v, ok := data[`IPower`]; ok {
		scan.Scan(scan.Pack(v), &x.IPower)
	}
	for k, v := range data {
		switch k {
		case `Code`:
			x.Code = v
		case `ID`:
			scan.Scan(scan.Pack(v), &x.ID)
		case `Coin`:
			scan.Scan(scan.Pack(v), &x.Coin)
		case `Multiple`:
			scan.Scan(scan.Pack(v), &x.Multiple)
		case `MakerFee`:
			scan.Scan(scan.Pack(v), &x.MakerFee)
		case `TakerFee`:
			scan.Scan(scan.Pack(v), &x.TakerFee)
		case `Scale`:
			scan.Scan(scan.Pack(v), &x.Scale)
		case `SharedRisk`:
			scan.Scan(scan.Pack(v), &x.SharedRisk)
		case `Risks`:
			if err := json.Unmarshal([]byte(v), &x.Risks); err == nil {
			}
		case `UnitPrice`:
			scan.Scan(scan.Pack(v), &x.UnitPrice)
		// case `LastPrice`:
		// 	if scan.Scan(scan.Pack(v), &double) > 0 {
		// 		x.LastPrice = Encode(double, x.PPower)
		// 	}
		// case `MarkPrice`:
		// 	if scan.Scan(scan.Pack(v), &double) > 0 {
		// 		x.MarkPrice = Encode(double, x.PPower)
		// 	}
		// case `BidPrice`:
		// 	if scan.Scan(scan.Pack(v), &double) > 0 {
		// 		x.BidPrice = Encode(double, x.PPower)
		// 	}
		// case `AskPrice`:
		// 	if scan.Scan(scan.Pack(v), &double) > 0 {
		// 		x.AskPrice = Encode(double, x.PPower)
		// 	}
		case `LimitSecurity`:
			scan.Scan(scan.Pack(v), &x.LimitSecurity)
		case `MaxPrice`:
			scan.Scan(scan.Pack(v), &x.MaxPrice)
		case `AllowShared`:
			scan.Scan(scan.Pack(v), &x.AllowShared)
		case `MaxVolume`:
			scan.Scan(scan.Pack(v), &x.MaxVolume)
		case `MinVolume`:
			scan.Scan(scan.Pack(v), &x.MinVolume)
		case `Cushion`:
			scan.Scan(scan.Pack(v), &x.Cushion)
		case `Waterline`:
			scan.Scan(scan.Pack(v), &x.Waterline)
		case `ForceKeepAddtion`:
			scan.Scan(scan.Pack(v), &x.ForceKeepAddtion)
		case `Closed`:
			scan.Scan(scan.Pack(v), &x.Closed)
		case `AllowMerge`:
			scan.Scan(scan.Pack(v), &x.AllowMerge)
		case `Shared`:
			scan.Scan(scan.Pack(v), &x.Shared)
		case `Merged`:
			scan.Scan(scan.Pack(v), &x.Merged)
		case `Coin1`:
			scan.Scan(scan.Pack(v), &x.Coin1)
		case `Coin2`:
			scan.Scan(scan.Pack(v), &x.Coin2)
		}
	}
}
