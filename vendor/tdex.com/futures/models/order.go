package models

// Order 订单
type Order struct {
	ID       uint    `xorm:"ID"`       // 编号
	CID      uint    `xorm:"CID"`      // 产品
	UID      uint    `xorm:"UID"`      // 账号
	PID      uint    `xorm:"PID"`      // 仓位
	RID      uint    `xorm:"RID"`      // 依赖订单(开仓订单中设置的止损止盈单)
	Attempt  Attempt `xorm:"Attempt"`  // 指令
	Shared   bool    `xorm:"Shared"`   // 全仓|逐仓 (开仓订单)
	Side     Side    `xorm:"Side"`     // 买|卖
	Scale    float64 `xorm:"Scale"`    // 杠杆(开仓订单)
	Kind     Kind    `xorm:"Kind"`     // 市价|限价
	Volume   float64 `xorm:"Volume"`   // 数量
	Distance bool    `xorm:"Distance"` // 价距(限价类型)(限价订单)
	Price    float64 `xorm:"Price"`    // 限价(限价订单)
	Locked   float64 `xorm:"Locked"`   // 限价触发后计算初始保证金使用的价格参数
	Limit    float64 `xorm:"Limit"`    // 开仓单挂单限价
	State    State   `xorm:"State"`    // 状态
	// ------------------------------
	Passive bool    `xorm:"Passive"` // 被动性(限价单)
	Visible float64 `xorm:"Visible"` // 隐藏性 <0可见(限价单)
	Better  bool    `xorm:"Better"`  // 以买一卖一价进入订单簿
	// ------------------------------
	Timely      Timely `xorm:"Timely"`      // 时效性(一直有效|触发后一段时间有效|截止时间)(限价单)
	TimelyParam int    `xorm:"TimelyParam"` // 时效参数(限价单)
	Deadline    int    `xorm:"Deadline"`    // 截止时间(限价单)
	//-------------------------------
	Strategy Strategy `xorm:"Strategy"` // 策略(无条件|界限|追踪)
	Variable Variable `xorm:"Variable"` // 条件变量(条件单)
	Constant float64  `xorm:"Constant"` // 条件常量(条件单)
	Relative bool     `xorm:"Relative"` // 条件常量是否位相对成交价的值(开仓订单设置的止损止盈)(条件单)
	Vertex   float64  `xorm:"Vertex"`   // 顶点(追踪订单)(条件单)
	//-------------------------------
	Filled   float64 `xorm:"Filled"`   // 已成交量
	SerialID int     `xorm:"SerialID"` // 序号
	//-------------------------------
	Activated int `xorm:"Activated"` // 激活时间
	Triggered int `xorm:"Triggered"` // 触发时间
	//-------------------------------
	Margin     float64 `xorm:"Margin"`     // 保证金（避免重复计算）
	InitMargin float64 `xorm:"InitMargin"` // 初始保证金
	Weight     float64 `xorm:"Weight"`     // 辅算金额：用于均价计算
	Notes      string  `xorm:"Notes"`      // 备注
	CreatedAt  int     `xorm:"CreatedAt1"` // 创建时间
	FID        uint    `xorm:"FID"`        // 源订单ID
	FVersion   int     `xorm:"FVersion"`   // 源订单版本号
	Version    int     `xorm:"Version"`    // 自身版本号
	UpdatedAt  int     `xorm:"UpdatedAt1"` // 修改时间
}

func (t *Order) TableName() string {
	return "orders"
}

// Validate 参数有效性验证
func (x Order) Validate() bool {
	return x.Side.Validate() && x.Kind.Validate() && x.Timely.Validate() && x.Strategy.Validate() && x.Variable.Validate()
}

// Change 直接修个数量
type Change struct {
	ID     uint
	Volume float64
}
