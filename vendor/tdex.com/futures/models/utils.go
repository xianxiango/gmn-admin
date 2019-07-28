package models

import (
	"math"
)

var indexes = make([]float64, 22)

func init() {
	for i, n := 0, 21; i <= n; i++ {
		indexes[i] = math.Pow10(i)
	}
}

// Encode 编码
func Encode(v float64, power int, ceil bool) int {
	v = indexes[power] * v
	if ceil {
		return int(math.Ceil(v))
	}
	return int(math.Floor(v))
}

// Decode 解码
func Decode(v int, power int) float64 {
	return float64(v) / indexes[power]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Mode 转换模式
type Mode uint

const (
	// Trunc 裁剪
	Trunc Mode = iota
	// Ceil 向上
	Ceil
	// Floor 向下
	Floor
	// Round 四舍五入
	Round
)

// Decimal 标准化
func Decimal(v float64, precision int, mode Mode) float64 {
	scale := indexes[precision]
	convert := converts[mode]
	return convert(v*scale) / scale
}

var converts = func() []func(float64) float64 {
	slice := make([]func(float64) float64, 0, 3)
	slice = append(slice, math.Trunc, math.Ceil, math.Floor, math.Round)
	return slice
}()
