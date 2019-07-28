package binary

import (
	"os"
	"unsafe"
)

// IEncoder 自主打包
type IEncoder interface {
	Encode(*Encoder)
}

// Encoder 序列化
type Encoder struct {
	i   int
	buf []byte
}

// Init 初始化
func (x *Encoder) Init(data []byte) *Encoder {
	x.buf, x.i = data, 0
	return x
}

// Reset 重置
func (x *Encoder) Reset(data []byte) (old []byte) {
	old, x.buf, x.i = x.buf, data, 0
	return
}

// Seek 设置偏移量
func (x *Encoder) Seek(offset, whence int) {
	switch whence {
	case os.SEEK_SET:
		if offset >= 0 && offset < len(x.buf) {
			x.i = offset
		}
	case os.SEEK_END:
		n := len(x.buf)
		if offset < 0 && n+offset >= 1 {
			x.i = n - offset - 1
		}
	default: // os.SEEK_CUR
		if offset >= 0 {
			if x.i+offset < len(x.buf) {
				x.i += offset
			}
		} else if x.i+offset >= 0 {
			x.i += offset
		}
	}
}

// Len 数据长度
func (x *Encoder) Len() int {
	return x.i
}

// Cap 容积
func (x *Encoder) Cap() int {
	return len(x.buf)
}

// Boolean 布尔
func (x *Encoder) Boolean(v bool) {
	*(*bool)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Uint 无符号整数
func (x *Encoder) Uint(v uint) {
	*(*uint)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Uint8 8bits无符号整数
func (x *Encoder) Uint8(v uint8) {
	*(*uint8)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Uint16 16bits无符号整数
func (x *Encoder) Uint16(v uint16) {
	*(*uint16)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Uint32 32bits无符号整数
func (x *Encoder) Uint32(v uint32) {
	*(*uint32)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Uint64 64bits无符号整数
func (x *Encoder) Uint64(v uint64) {
	*(*uint64)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Int 整数
func (x *Encoder) Int(v int) {
	*(*int)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Int8 8bits整数
func (x *Encoder) Int8(v int8) {
	*(*int8)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Int16 16bits整数
func (x *Encoder) Int16(v int16) {
	*(*int16)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Int32 32bits整数
func (x *Encoder) Int32(v int32) {
	*(*int32)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Int64 64bits整数
func (x *Encoder) Int64(v int64) {
	*(*int64)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Float64 64bits浮点数
func (x *Encoder) Float64(v float64) {
	*(*float64)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Float32 32bits浮点数
func (x *Encoder) Float32(v float32) {
	*(*float32)(unsafe.Pointer(&x.buf[x.i])), x.i = v, x.i+int(unsafe.Sizeof(v))
}

// Bytes bytes
func (x *Encoder) Bytes(v []byte) {
	n := len(v)
	*(*int)(unsafe.Pointer(&x.buf[x.i])), x.i = n, x.i+int(unsafe.Sizeof(n))
	copy(x.buf[x.i:], v)
	x.i += n
}

// Join 拼接
func (x *Encoder) Join(data []byte) {
	copy(x.buf[x.i:], data)
	x.i += len(data)
}

// Use 直接使用
func (x *Encoder) Use(n int) []byte {
	r := x.buf[x.i : x.i+n]
	x.i += n
	return r
}

// String 字符串
func (x *Encoder) String(v string) {
	data := string2slice(v)
	n := len(data)
	*(*int)(unsafe.Pointer(&x.buf[x.i])), x.i = n, x.i+int(unsafe.Sizeof(n))
	// if err := x.assert(n); err != nil {
	// 	return err
	// }
	copy(x.buf[x.i:], data)
	x.i += n
}

// Data 数据
func (x *Encoder) Data() []byte {
	return x.buf[:x.i]
}

// Pointer 地址
func (x *Encoder) Pointer() []byte {
	return x.buf
}

// Write 写入
func (x *Encoder) Write(data []byte) (int, error) {
	if len(x.buf)-x.i >= len(data) {
		x.Join(data)
		return len(data), nil
	}
	return 0, os.ErrPermission
}

func string2slice(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
