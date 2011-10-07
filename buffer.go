// author: liuaifu
// laf163@gmail.com
// 2011-10-3

package buffer

import "fmt"

type Buffer struct{
	Data []byte
}

func New() *Buffer {
	p:=&Buffer{[]byte{}}
	return p
}

func (this *Buffer) Buffer() []byte {
	return this.Data
}

func (this *Buffer) Length() int{
	return len(this.Data)
}

func (this *Buffer) WriteUint8(n uint8) {
	this.Data=append(this.Data,n)
}

func (this *Buffer) WriteUint16(n uint16) {
	buf:=make([]byte,2)
	buf[0]=uint8(n)
	buf[1]=uint8(n>>8)
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) WriteUint32(n uint32) {
	buf:=make([]byte,4)
	buf[0]=uint8(n)
	buf[1]=uint8(n>>8)
	buf[2]=uint8(n>>16)
	buf[3]=uint8(n>>24)
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) WriteUint64(n uint64) {
	buf:=make([]byte,8)
	buf[0]=uint8(n)
	buf[1]=uint8(n>>8)
	buf[2]=uint8(n>>16)
	buf[3]=uint8(n>>24)
	buf[4]=uint8(n>>32)
	buf[5]=uint8(n>>40)
	buf[6]=uint8(n>>48)
	buf[7]=uint8(n>>56)
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) WriteInt16(n int16) {
	buf:=make([]byte,2)
	buf[0]=uint8(n)
	buf[1]=uint8(n>>8)
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) WriteInt32(n int32) {
	buf:=make([]byte,4)
	buf[0]=uint8(n)
	buf[1]=uint8(n>>8)
	buf[2]=uint8(n>>16)
	buf[3]=uint8(n>>24)
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) WriteStr(str string) {
	buf:=[]byte(str)
	this.Data=append(this.Data,buf...)
	this.Data=append(this.Data,0)
}

func (this *Buffer) WriteString(str string, Length int) {
	buf:=[]byte(str)
	if len(buf)>Length{
		buf=buf[:Length]
	}else if len(buf)<Length{
		zeroCount:=Length-len(buf)
		for i:=0;i<zeroCount;i++{
			buf=append(buf,0)
		}
	}
	this.Data=append(this.Data,buf...)
}


func (this *Buffer) ReadStr() string {
	for i:=0;i<len(this.Data);i++{
		if this.Data[i]==0{
			str := string(this.Data[:i])
			this.Data=this.Data[i+1:]
			return str
		}
	}
	return ""
}

func (this *Buffer) ReadString(Length int) string {
	buf:=this.Data[:Length]
	for i:=0;i<len(buf);i++{
		if buf[i]==0{
			buf=buf[:i]
			break
		}
	}
	this.Data=this.Data[Length:]
	return string(buf)
}

func (this *Buffer) ReadIp() string {
	ip:=fmt.Sprintf("%d.%d.%d.%d", this.Data[0], this.Data[1], this.Data[2], this.Data[3])
	this.Data=this.Data[4:]
	return ip
}

func (this *Buffer) ReadUint8() uint8 {
	n:=this.Data[0]
	this.Data=this.Data[1:]
	return n
}

func (this *Buffer) ReadUint16() uint16 {
	var n uint16
	n=uint16(this.Data[0])
	n|=uint16(this.Data[1])<<8
	this.Data=this.Data[2:]
	return n
}

func (this *Buffer) ReadUint32() uint32 {
	var n uint32
	n=uint32(this.Data[0])
	n|=uint32(this.Data[1])<<8
	n|=uint32(this.Data[2])<<16
	n|=uint32(this.Data[3])<<24
	this.Data=this.Data[4:]
	return n
}

func (this *Buffer) ReadInt32() int32 {
	var n int32
	n=int32(this.Data[0])
	n|=int32(this.Data[1])<<8
	n|=int32(this.Data[2])<<16
	n|=int32(this.Data[3])<<24
	this.Data=this.Data[4:]
	return n
}

func (this *Buffer) ReadUint64() uint64 {
	var n uint64
	n=uint64(this.Data[0])
	n|=uint64(this.Data[1])<<8
	n|=uint64(this.Data[2])<<16
	n|=uint64(this.Data[3])<<24
	n|=uint64(this.Data[4])<<32
	n|=uint64(this.Data[5])<<40
	n|=uint64(this.Data[6])<<48
	n|=uint64(this.Data[7])<<56
	this.Data=this.Data[8:]
	return n
}

func (this *Buffer) ReadInt64() int64 {
	var n int64
	n=int64(this.Data[0])
	n|=int64(this.Data[1])<<8
	n|=int64(this.Data[2])<<16
	n|=int64(this.Data[3])<<24
	n|=int64(this.Data[4])<<32
	n|=int64(this.Data[5])<<40
	n|=int64(this.Data[6])<<48
	n|=int64(this.Data[7])<<56
	this.Data=this.Data[8:]
	return n
}

func (this *Buffer) Append(buf []byte) {
	this.Data=append(this.Data,buf...)
}

func (this *Buffer) AppendBuffer(buf *Buffer) {
	this.Data=append(this.Data,buf.Data...)
}
