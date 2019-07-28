package binary

import (
	"io"
	"unsafe"
)

// Reader 读
type Reader struct {
	io.Reader
	buf []byte
}

// Reset 重置
func (x *Reader) Reset(f io.Reader, cache []byte) (fold io.Reader, old []byte) {
	fold, old, x.Reader, x.buf = x.Reader, x.buf, f, cache
	return
}

// Skip 跳过
func (x *Reader) Skip(n int) (err error) {
	m := len(x.buf)
	if m > n {
		m = n
	}
	// 零头
	k := n % m
	if k > 0 {
		if _, err = x.Reader.Read(x.buf[:k]); err != nil {
			return
		}
	}
	n -= k
	for i := 0; i < n; i += m {
		if _, err = x.Reader.Read(x.buf[:m]); err != nil {
			return
		}
	}
	return
}

// Boolean 布尔
func (x *Reader) Boolean() (v bool, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Uint 无符号整数
func (x *Reader) Uint() (v uint, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Uint64 无符号64位整数
func (x *Reader) Uint64() (v uint64, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Uint32 无符号32位整数
func (x *Reader) Uint32() (v uint32, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Uint16 无符号16位整数
func (x *Reader) Uint16() (v uint16, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Uint8 无符号8位整数
func (x *Reader) Uint8() (v uint8, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Int 整数
func (x *Reader) Int() (v int, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Int64 64位整数
func (x *Reader) Int64() (v int64, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Int32 32位整数
func (x *Reader) Int32() (v int32, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Int16 16位整数
func (x *Reader) Int16() (v int16, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Int8 8位整数
func (x *Reader) Int8() (v int8, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Float32 32位浮点数
func (x *Reader) Float32() (v float32, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// Float64 64位浮点数
func (x *Reader) Float64() (v float64, err error) {
	_, err = x.Reader.Read((*(*[int(unsafe.Sizeof(v))]byte)(unsafe.Pointer(&v)))[:])
	return
}

// String 字符串
func (x *Reader) String() (v string, err error) {
	var n int
	if _, err = x.Reader.Read((*(*[int(unsafe.Sizeof(n))]byte)(unsafe.Pointer(&n)))[:]); err != nil {
		return
	}
	if _, err = x.Reader.Read(x.buf[:n]); err != nil {
		return
	}
	v = string(x.buf[:n])
	return
}

// Bytes 字节数组
func (x *Reader) Bytes() (v []byte, err error) {
	var n int
	if _, err = x.Reader.Read((*(*[int(unsafe.Sizeof(n))]byte)(unsafe.Pointer(&n)))[:]); err != nil {
		return
	}
	if _, err = x.Reader.Read(x.buf[:n]); err != nil {
		return
	}
	v = make([]byte, n, n)
	copy(v, x.buf[:n])
	return
}
