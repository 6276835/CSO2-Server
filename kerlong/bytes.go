package kerlong

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io/ioutil"
)

//Encode 结构体转数组
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//Decode 数组转结构体
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

//BytesCombine 连接字符串
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

//WriteUint8 写入1字节uint8
func WriteUint8(b *[]byte, i uint8, offset *int) {
	(*b)[*offset] = i
	(*offset)++
}

//WriteUint16 写入2字节uint16
func WriteUint16(b *[]byte, i uint16, offset *int) {
	(*b)[*offset] = uint8(i)
	(*b)[*offset+1] = uint8(i >> 8)
	(*offset) += 2
}

//WriteUint32 写入4字节uint32
func WriteUint32(b *[]byte, i uint32, offset *int) {
	(*b)[*offset] = uint8(i)
	(*b)[*offset+1] = uint8(i >> 8)
	(*b)[*offset+2] = uint8(i >> 16)
	(*b)[*offset+3] = uint8(i >> 24)
	(*offset) += 4
}

//WriteUint64 写入8字节uint64
func WriteUint64(b *[]byte, i uint64, offset *int) {
	(*b)[*offset] = uint8(i)
	(*b)[*offset+1] = uint8(i >> 8)
	(*b)[*offset+2] = uint8(i >> 16)
	(*b)[*offset+3] = uint8(i >> 24)
	(*b)[*offset+4] = uint8(i >> 32)
	(*b)[*offset+5] = uint8(i >> 40)
	(*b)[*offset+6] = uint8(i >> 48)
	(*b)[*offset+7] = uint8(i >> 56)
	(*offset) += 8
}

//WriteUint16BE 写入2字节uint16，大端模式
func WriteUint16BE(b *[]byte, i uint16, offset *int) {
	(*b)[*offset] = uint8(i >> 8)
	(*b)[*offset+1] = uint8(i)
	(*offset) += 2
}

//WriteUint32BE 写入4字节uint32，大端模式
func WriteUint32BE(b *[]byte, i uint32, offset *int) {
	(*b)[*offset] = uint8(i >> 24)
	(*b)[*offset+1] = uint8(i >> 16)
	(*b)[*offset+2] = uint8(i >> 8)
	(*b)[*offset+3] = uint8(i)
	(*offset) += 4
}

//WriteUint64BE 写入8字节uint64，大端模式
func WriteUint64BE(b *[]byte, i uint64, offset *int) {
	(*b)[*offset] = uint8(i >> 56)
	(*b)[*offset+1] = uint8(i >> 48)
	(*b)[*offset+2] = uint8(i >> 40)
	(*b)[*offset+3] = uint8(i >> 32)
	(*b)[*offset+4] = uint8(i >> 24)
	(*b)[*offset+5] = uint8(i >> 16)
	(*b)[*offset+6] = uint8(i >> 8)
	(*b)[*offset+7] = uint8(i)
	(*offset) += 8
}

//WriteString 写入字符串，包括长度
func WriteString(dest *[]byte, src []byte, offset *int) int {
	l := len(src)
	WriteUint8(dest, uint8(l), offset)
	if l == 0 {
		return 1
	}
	for i := 0; i < l; i++ {
		(*dest)[*offset] = src[i]
		(*offset)++
	}
	return l + 1
}

//WriteString 写入字符串，不包括长度
func WriteStringWithNull(dest *[]byte, src []byte, offset *int) int {
	l := len(src)
	for i := 0; i < l; i++ {
		(*dest)[*offset] = src[i]
		(*offset)++
	}
	WriteUint8(dest, 0x00, offset)
	return l + 1
}

//WriteLongString 写入字符串，包括2字节长度
func WriteLongString(dest *[]byte, src []byte, offset *int) int {
	l := len(src)
	WriteUint16(dest, uint16(l), offset)
	if l == 0 {
		return 2
	}
	for i := 0; i < l; i++ {
		(*dest)[*offset] = src[i]
		(*offset)++
	}
	return l + 2
}

//WriteUint32Array 写入uint32数组
func WriteUint32Array(b *[]byte, a []uint32, offset *int) {
	for i := 0; i < len(a); i++ {
		WriteUint32(b, a[i], offset)
	}
}

//ReadUint8 读取1字节到uint8
func ReadUint8(b []byte, offset *int) uint8 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint8
	binary.Read(buf, binary.LittleEndian, &i)
	(*offset)++
	return i
}

//ReadUint16 读取2字节到uint16
func ReadUint16(b []byte, offset *int) uint16 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint16
	binary.Read(buf, binary.LittleEndian, &i)
	(*offset) += 2
	return i
}

//ReadUint32 读取4字节到uint32
func ReadUint32(b []byte, offset *int) uint32 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint32
	binary.Read(buf, binary.LittleEndian, &i)
	(*offset) += 4
	return i
}

//ReadUint64 读取8字节到uint64
func ReadUint64(b []byte, offset *int) uint64 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint64
	binary.Read(buf, binary.LittleEndian, &i)
	(*offset) += 8
	return i
}

//ReadUint16BE 大端读取2字节到uint16
func ReadUint16BE(b []byte, offset *int) uint16 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint16
	binary.Read(buf, binary.BigEndian, &i)
	(*offset) += 2
	return i
}

//ReadUint32BE 大端读取4字节到uint32
func ReadUint32BE(b []byte, offset *int) uint32 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint32
	binary.Read(buf, binary.BigEndian, &i)
	(*offset) += 4
	return i
}

//ReadUint64BE 大端读取8字节到uint64
func ReadUint64BE(b []byte, offset *int) uint64 {
	buf := bytes.NewBuffer(b[(*offset):])
	var i uint64
	binary.Read(buf, binary.BigEndian, &i)
	(*offset) += 8
	return i
}

//ReadString 大端不会读取长度，需要单独先读取长度
func ReadString(b []byte, offset *int, len int) []byte {
	(*offset) += len
	return b[(*offset)-len : (*offset)]
}

//ReadUint32Array 读取数据到uint32数组
func ReadUint32Array(b []byte, offset *int, len int) []uint32 {
	var buf []uint32
	for i := 0; i < len; i++ {
		buf = append(buf, ReadUint32(b, offset))
	}
	return buf
}

//GetUint16 获取b数组最前面的2字节数据
func GetUint16(b []byte) uint16 {
	buf := bytes.NewBuffer(b)
	var i uint16
	binary.Read(buf, binary.LittleEndian, &i)
	return i
}

//GetUint32 获取b数组最前面的4字节数据
func GetUint32(b []byte) uint32 {
	buf := bytes.NewBuffer(b)
	var i uint32
	binary.Read(buf, binary.LittleEndian, &i)
	return i
}

//GetUint64 获取b数组最前面的8字节数据
func GetUint64(b []byte) uint64 {
	buf := bytes.NewBuffer(b)
	var i uint64
	binary.Read(buf, binary.LittleEndian, &i)
	return i
}

//GetUint16BE 获取b数组最前面的2字节数据，大端模式
func GetUint16BE(b []byte) uint16 {
	buf := bytes.NewBuffer(b)
	var i uint16
	binary.Read(buf, binary.BigEndian, &i)
	return i
}

//GetUint32BE 获取b数组最前面的4字节数据，大端模式
func GetUint32BE(b []byte) uint32 {
	buf := bytes.NewBuffer(b)
	var i uint32
	binary.Read(buf, binary.BigEndian, &i)
	return i
}

//GetUint64BE 获取b数组最前面的8字节数据，大端模式
func GetUint64BE(b []byte) uint64 {
	buf := bytes.NewBuffer(b)
	var i uint64
	binary.Read(buf, binary.BigEndian, &i)
	return i
}

//CompareBytes 比较两个数组是否内容相等,相等返回true
func CompareBytes(a []byte, b []byte) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//BuildLongString 构建长字符串消息，包含2字节的字符串大小
func BuildLongString(src []byte) []byte {
	buf := make([]byte, 2)
	offset := 0
	WriteUint16(&buf, uint16(len(src)), &offset)
	return BytesCombine(buf, src)
}

//BuildString 构建长字符串消息，包含2字节的字符串大小
func BuildString(src []byte) []byte {
	buf := make([]byte, 1)
	offset := 0
	WriteUint8(&buf, uint8(len(src)), &offset)
	return BytesCombine(buf, src)
}

//IsSameName 检查dest是否与src字符串像差不多
func IsSameName(dest, src []byte) bool {
	nextval := make([]int, len(src)+1)
	nextval = getnextval(src, nextval)
	i, j, ls, lt := 0, 0, len(dest), len(src)
	for i < ls && j < lt {
		if j == -1 || dest[i] == src[j] {
			i++
			j++
		} else {
			j = nextval[j]
		}
		if j == lt {
			return true
		}
	}
	return false
}

func getnext(str []byte, next []int) []int {
	j, k := 0, -1
	next[0] = -1
	for j < len(str)-1 {
		if k == -1 || str[j] == str[k] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func getnextval(str []byte, nextval []int) []int {
	j, k := 0, -1
	nextval[0] = -1
	for j < len(str)-1 {
		if k == -1 || str[j] == str[k] {
			j++
			k++
			nextval[j] = nextval[k]
		} else {
			k = nextval[k]
		}
	}
	return nextval
}

func ReadStringFromFile(path string) string {
	Bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(Bytes)
}
