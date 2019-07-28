package models

// Delegate 委托
type Delegate struct {
	UID       uint    // 账号
	ID        uint    // 编号 等于订单编号
	Side      Side    // 交易方向
	Price     float64 // 价格 =0:市价
	Volume    float64 // 数量
	Visible   float64 // 隐藏性 <0完全可见 =0 完全不可见。 >0 可见数量
	Deadline  int     // 截止时间 <=0: 一直有效
	SerialID  int     // 序号
	Activated int     // 激活时间
	Merged    bool    // 合仓模式|分仓模式
	Attempt   Attempt // 爆仓
	Scale     float64 // 杠杆
	Better    bool    // 买一卖一价入市
	Passive   bool    // 被动性
}

// Quote 报价
type Quote struct {
	Price  float64
	Volume float64
}
