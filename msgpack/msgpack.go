package main

import (
	"fmt"
	"github.com/ugorji/go/codec"
)

type LuasRequest struct {
	Cmdid uint32 `protobuf:"varint,1,opt,name=cmdid" json:"cmdid,omitempty"`
	Appid string `protobuf:"bytes,2,opt,name=appid" json:"appid,omitempty"`
	Udid  string `protobuf:"bytes,3,opt,name=udid" json:"udid,omitempty"`
	Ucid  string `protobuf:"bytes,4,opt,name=ucid" json:"ucid,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func main() {
	resp := LuasRequest{
		Cmdid : 0x0308,
		Appid : "aa",
		Udid : "",
		Ucid : "cc",
		Data : nil,
	}
	var out []byte
	var mh codec.MsgpackHandle
	//编码
	err := codec.NewEncoderBytes(&out, &mh).Encode(resp)
	if err != nil {
		fmt.Println("err : ", err)
	} else {
		fmt.Println(out)
	}
	//解码
	req := new(LuasRequest)
	codec.NewDecoderBytes(out, &mh).Decode(req)
	fmt.Println(req)
}
