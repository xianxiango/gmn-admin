package models

// Side 交易方向
type Side uint8

// Validate 有效性检查
func (x Side) Validate() bool {
	return x >= BUY && x <= SELL
}

const (
	// BUY 买
	BUY Side = iota
	// SELL 卖
	SELL
)

// Kind 市价|限价
type Kind uint8

// Validate 有效性检查
func (x Kind) Validate() bool {
	return x >= MKT && x <= LMT
}

const (
	// MKT 市价
	MKT Kind = iota
	// LMT 限价
	LMT
)

// Attempt 企图
type Attempt uint8

// Validate 有效性检查
func (x Attempt) Validate() bool {
	return x >= TP && x <= OPEN
}

const (
	// TP 止盈
	TP Attempt = iota
	// SL 止损
	SL
	// CLOSE 平仓
	CLOSE
	// OPEN 开仓
	OPEN
	// CLEAR 清算(爆仓处理)
	CLEAR
	// ADL 自动减仓
	ADL
	// Splited 拆分
	Splited
	// Merged 合并
	Merged
	// Income 盈亏
	Income
)

// Timely 时效性
type Timely uint8

// Validate 有效性检查
func (x Timely) Validate() bool {
	return x >= GTC && x <= DEADLINE
}

const (
	// GTC 一直有效
	GTC Timely = iota
	// LIFE 触发之后一定时间内
	LIFE
	// DEADLINE 截止时间
	DEADLINE
)

// State 状态
type State uint8

// 取消中应该用变量表示
const (
	// CREATED 创建
	CREATED State = iota
	// WAITING 等待中
	WAITING
	// TRIGGERED 已触发
	TRIGGERED
	// CANCELED 已撤销
	CANCELED
	// FILLED 已成交
	FILLED
	// CLOSED 已关闭
	CLOSED
	// QUEUING 排队
	QUEUING
)

// Variable 变量
type Variable uint8

// Validate 有效性检查
func (x Variable) Validate() bool {
	return x >= LastPrice && x <= MarkPrice
}

const (
	// LastPrice 市价
	LastPrice Variable = iota
	// MarkPrice 标记价
	MarkPrice
	// BidPrice 买一价
	BidPrice
	// AskPrice 卖一价
	AskPrice
)

// Strategy 策略
type Strategy uint8

// Validate 有效性检查
func (x Strategy) Validate() bool {
	return x >= Immediate && x <= Trail
}

const (
	// Immediate 直接
	Immediate Strategy = iota
	// Limit 限价
	Limit
	// Trail 追踪
	Trail
)

// Target 编辑选择对象
type Target uint8

// Validate 有效性检查
func (x Target) Validate() bool {
	return x >= ORDER && x <= POSITION
}

const (
	// ORDER 订单
	ORDER Target = iota
	// POSITION 仓位
	POSITION
)

// Rote 角色
type Rote uint8

const (
	// MAKER 挂单者
	MAKER Rote = iota
	// TAKER 吃单者
	TAKER
)

// Operate 操作
type Operate uint8

const (
	// Unknown 未知
	Unknown Operate = iota
	// MktClose 市价平仓
	MktClose
	// LmtClose 限价平仓
	LmtClose
	// TPClose 止盈平仓
	TPClose
	// SLClose 止损平仓
	SLClose
	// TrailSLClose 追踪止损平仓
	TrailSLClose
	// ForceClose 强制平仓
	ForceClose
	// WeakenClose 自动减仓
	WeakenClose
	// MergeClose 合并仓位
	MergeClose
	// SplitClose 拆分仓位
	SplitClose
)

const (
	// OffsetUpdatedAt 更新时间字段
	OffsetUpdatedAt uint8 = iota + 50
	// OffsetFee 佣金字段
	OffsetFee
	// OffsetRote 角色字段
	OffsetRote
	// OffsetFinalPrice 成交价字段
	OffsetFinalPrice
	// OffsetCostPrice 成本价字段
	OffsetCostPrice
	// OffsetXVolume 成交量字段
	OffsetXVolume
	// OffsetPositionStar 仓位危险系数
	OffsetPositionStar
	// OffsetMargin 保证金
	OffsetMargin
	// OffsetIncome 盈亏
	OffsetIncome
	// OffsetDeductSection 抵扣序号
	OffsetDeductSection

	// OffsetShared 全仓模式
	OffsetShared = 0
	// OffsetMerged 自动合仓
	OffsetMerged = 1
	// OffsetRobot 机器人
	OffsetRobot = 2
	// OffsetInfinite 无限钱包
	OffsetInfinite = 3
)
