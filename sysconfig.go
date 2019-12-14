package modbusy

import (
	"fmt"
	"strings"
)

// 本文件处理系统输入配置信息相关代码
// 主要负责导入通道信息、仪表信息、协议信息

// MDevice modbus device
// Addr is modbus slave/server addr: 0-255
// Emable means the device is enable
// Channel: modbus rs485: rs485_comxx for windows or rs485_ttySUBS0, or tcp_192.168.1.1:502 or udp_192.168.3.1:1052
// ProtocalID is protocal id that reg in system
type MDevice struct {
	Name       string `json:"name"`
	Addr       uint8  `json:"addr"`
	Enable     bool   `json:"enable"`
	Channel    string `json:"channel"`
	ProtocolID string `json:"protocolid"`
}

// Protocol a modbus protocol for some device include modbus commands and suitable data parse rule
// name and id is necessary for identify, id generate rule M_Ex/Wx/Hx/Ax/Ox_YYYYMMDD_XXX
// var suppchannel means the protocol support channel type, for std modbus support all channel.
// Some Device class: EL for ele meter, LW for life water meter, HC for heat/cool meter,
// AQ for air quality index, VF for pump/fan vfd, VR for VRV unit
type Protocol struct {
	Name        string           `json:"name"`                   // 协议名称
	Alias       string           `json:"company_proseries"`      // 设备厂家_产品系列
	ID          string           `json:"id"`                     // 协议编码
	SuppChannel byte             `json:"supp_channel,omitempty"` // 标记支持的通道
	Cmds        []Command        `json:"cmds"`                   // 命令列表
	ParseRules  [][]TagParseRule `json:"parse_rules"`            // 解析规则，第一层用于和cmds匹配，一个cmd需要多个tag
}

//其他信息归上层服务器管理，例如设备安装位置，厂家、编号等等，与收发无关的信息，该scada的作用就是让采集信息有意义，但非全部意义
//暂时不确定 ParseRules 是否合适用json配置，如果过于复杂的话，需要生成json或简化为 []TagParseRule

// func (p Protocol)cmds()

// TagParseRule tag value calc rule real_v = ((type)measure_value - offset)*pc/ct
type TagParseRule struct {
	ID         string  `json:"cmdid_tagid"`      // CMDID_TAGID
	Tag        string  `json:"tag"`              // 参数名
	PhyType    string  `json:"phy_type"`         // 参数编码
	StartIndex uint    `json:"start_index"`      // PDU数据区起始地址
	EndIndex   uint    `json:"end_index"`        // PDU数据区终止地址
	Offset     float64 `json:"offset,omitempty"` // 数据修正中心点偏移量，默认为0
	VType      string  `json:"vtype"`            // 数据类型，用来标记是解析 modbus 大小端，整数类型/浮点数类型/布尔型 &~&
	PT         uint    `json:"pt,omitempty"`     // PT变比,默认为1
	CT         uint    `json:"ct,omitempty"`     // CT变比，默认为1
}

//Command one command is modbus or rs485 protocol command for get a/some value
type Command struct {
	Alias     string `json:"name"`          //仅用于助记
	Style     string `json:"style"`         // 用于标记命令风格
	ID        string `json:"id"`            //用于索引一条命令，在同一个协议内不允许重复
	Funcode   uint   `json:"funcode"`       //Modbus/PLC 功能码
	StartAddr uint   `json:"startaddr"`     //Modbus寄存器地址，针对PLCStyle/0xFFFF
	RegLength uint   `json:"reglength"`     //Modbus寄存器长度
	Buf       []byte `json:"buf,omitempty"` //Buf 配置情况下，一般为空，仅用于写入情况
}

func (c *Command) encode() (pdu []byte, err error) {
	styleStd := strings.ToUpper(c.Style)
	switch styleStd {
	case "PLC":
		return nil, nil
	case "RAW":
		return nil, nil
	default:
		return nil, fmt.Errorf("unrecognized command sytle/无法识别的命令风格:%s", c.Style)
	}
}

func (c *Command) genCmdWPLC(funcode, startaddr, reglenth uint, buf *[]byte) {

}

// 由于各个字段不同情况范围不一样，那么由业务检查，代码里不做语言要求

//根据各个字段索引关系，规定如下
//1. Parameter 参数由软件自带，给出常用参数向及内容(配置界面1)
//2. 建立协议,意味着命令与解析规则同时建立并导入（配置界面2）
//3. 建立信道及设备编号表，导入(建立)
