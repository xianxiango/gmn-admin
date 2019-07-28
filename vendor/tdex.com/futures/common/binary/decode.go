package binary

import (
	"errors"
	"unsafe"
)

const Magic = 8192

// IDecoder 自主解包
type IDecoder interface {
	Decode(*Decoder)
}

// Decoder 读取器
type Decoder struct {
	buf []byte
	i   int
}

// Init 初始化
func (x *Decoder) Init(data []byte) *Decoder {
	x.buf, x.i = data, 0
	return x
}

// Reset 重置
func (x *Decoder) Reset(data []byte) (old []byte) {
	old, x.buf, x.i = x.buf, data, 0
	return
}

// Seek 设置偏移量
func (x *Decoder) Seek(offset int) {
	if offset >= 0 && offset < len(x.buf) {
		x.i = offset
	}
}

// Skip 跳过
func (x *Decoder) Skip(n int) {
	x.i += n
}

// Offset 当前位置
func (x *Decoder) Offset() int {
	return x.i
}

// Boolean 布尔
func (x *Decoder) Boolean() (v bool) {
	v, x.i = *(*bool)(unsafe.Pointer(&x.buf[x.i])), x.i+int(unsafe.Sizeof(v))
	return
}

// Uint 无符号整数
func (x *Decoder) Uint() (v uint) {
	v, x.i = *(*uint)(unsafe.Pointer(&x.buf[x.i])), x.i+int(unsafe.Sizeof(v))
	return
}

// Uint64 无符号64位整数
func (x *Decoder) Uint64() uint64 {
	v := *(*uint64)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Uint32 无符号32位整数
func (x *Decoder) Uint32() uint32 {
	v := *(*uint32)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Uint16 无符号16位整数
func (x *Decoder) Uint16() uint16 {
	v := *(*uint16)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Uint8 无符号8位整数
func (x *Decoder) Uint8() uint8 {
	v := *(*uint8)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Int 整数
func (x *Decoder) Int() int {
	v := *(*int)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Int64 64位整数
func (x *Decoder) Int64() int64 {
	v := *(*int64)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Int32 32位整数
func (x *Decoder) Int32() int32 {
	v := *(*int32)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Int16 16位整数
func (x *Decoder) Int16() int16 {
	v := *(*int16)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Int8 8位整数
func (x *Decoder) Int8() int8 {
	v := *(*int8)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Float32 32位浮点数
func (x *Decoder) Float32() float32 {
	v := *(*float32)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(v))
	return v
}

// Float64 64位浮点数
func (x *Decoder) Float64() (v float64) {
	v, x.i = *(*float64)(unsafe.Pointer(&x.buf[x.i])), x.i+int(unsafe.Sizeof(v))
	return
}

// String 字符串
func (x *Decoder) String() string {
	n := *(*int)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(n))
	o := string(x.buf[x.i : x.i+n])
	x.i += n
	return o
}

// Bytes 字节数组
func (x *Decoder) Bytes() []byte {
	n := *(*int)(unsafe.Pointer(&x.buf[x.i]))
	x.i += int(unsafe.Sizeof(n))
	o := make([]byte, n, n)
	copy(o, x.buf[x.i:])
	x.i += n
	return o[:n]
}

// Bytes 字节数组
func (x *Decoder) BytesSafe() ([]byte, error) {
	if x.Len() <= 0 {
		return []byte{}, errors.New("out of range")
	}
	n := *(*int)(unsafe.Pointer(&x.buf[x.i]))
	if n <= 0 || n > Magic || x.i+int(unsafe.Sizeof(n)) >= x.Len() {
		return []byte{}, errors.New("out of range")
	}
	x.i += int(unsafe.Sizeof(n))
	o := make([]byte, n, n)
	copy(o, x.buf[x.i:])
	x.i += n
	return o[:n], nil
}

// Data 剩余
func (x *Decoder) Data() []byte {
	return x.buf[x.i:]
}

// Pointer 内存
func (x *Decoder) Pointer() []byte {
	return x.buf
}

// Len 剩余数据
func (x *Decoder) Len() int {
	return len(x.buf) - x.i
}
