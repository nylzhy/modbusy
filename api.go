package modbusy

// RawClient a modbus raw client read/write operation collection for modbus instruction with modbus protocol style
type RawClient interface {
	// 寄存器类型			功能				功能码		寄存器地址范围		读取/写入寄存器数量			操作数据
	// 线圈				读取线圈				0x01		0x0000-0xFFFF		0-2000(0x001~0x7D0)			bit
	// 离散输入			读取离散输入	 	 	0x02		0x0000-0xFFFF		1-2000(0x001~0x7D0)	 		bit
	// 保持寄存器		读取保持寄存器			0x03		0x0000-0xFFFF		1-125(0x01~0x7D)		 	2bytes
	// 输入寄存器		读取输入寄存器			0x04		0x0000-0xFFFF		1-125(0x01~0x7D)			2bytes
	// 线圈				写单个线圈				0x05		0x0000-0xFFFF		1							0xFF00/0x0000
	// 保持寄存器		写单个保持寄存器		0x06		0x0000-0xFFFF 		1							2bytes
	// 线圈				写多个线圈				0x0F 		0x0000-0XFFFF		1-2000(0x001~0x7D0)     	Bit
	// 保持寄存器		写多个保持寄存器	 	0X10		0x0000-0xFFFF		1-120(0x01-0x78)			2bytes

	//Read Coils/DiscreteInputs/InputRegisters/HoldingRegisters
	Read(slaveIDAddr uint8, funCode uint8, regAddr, length uint) (result []byte, err error)
	//Read/Write Coils/DiscreteInputs/InputRegisters/HoldingRegisters
	Exec(slaveIDAddr uint8, funCode uint8, regAddr, length uint, data []byte) (result []byte, err error)
}

// HoldRegClient a hold regigter client read/write operation collection for modbus instruction with PLC style
type HoldRegClient interface {
	//Read HoldingRegisters
	Read(slaveIDAddr uint8, regAddr, length uint) (result []byte, err error)
	//Read/Write HoldingRegisters
	Write(slaveIDAddr uint8, regAddr, length uint, data []byte) (result []byte, err error)
}

// PLCClient a client read/write operation collection for modbus instruction with PLC style
type PLCClient interface {
	// Exec execute one modbus instruction with plc mode
	// PLC Mode: false 表示读取，true 表写入
	//
	// PLC Modbus 数据地址、数据长度与功能码关系表
	// 数据读取
	// *******************************************************************
	// 数据起始地址, 	数据长度,	 		功能码
	// 0-9999,	  		1-2000,				01			// 读取离散输出
	// 10001-19999,	  	1-2000,          	02			// 读取离散输入
	// 40001-49999,	  	1-125,          	03			// 读取保持寄存器
	// 400001-465535,	1-125,				03			// 读取保持寄存器
	// 30001-39999,		1-125,				04			// 读取输入寄存器
	// 20001-29999,		1-125,				03/04		// 读取浮点寄存器
	// *******************************************************************
	//
	// 数据写入
	// *******************************************************************
	// 数据起始地址, 	数据长度,	 		功能码
	// 0-9999,	  		1,					05			// 写入单个离散输出
	// 40001-49999,	  	1,          		06			// 写入单个保持寄存器
	// 400001-465535,	1,					06			// 写入单个保持寄存器
	// 0-9999,	  		2-1968,				15			// 写入多个离散输出
	// 40001-49999,	  	2-123,          	16			// 写入多个保持寄存器
	// 400001-465535,	2-123,				16			// 写入单个保持寄存器
	// 20001-29999,		1-125,				06/16		// 写入多个浮点寄存器
	// *******************************************************************

	//Read/Write Coils/DiscreteInputs/InputRegisters/HoldingRegisters with PLC mode
	Exec(SlaveIDaddr uint8, mode bool, addr, length uint, databuf []byte) (result []byte, err error)
}

// configure not conside right
type comchannel struct {
	highQueue *[]int
	lowQueue  *[]int
}

// 设计时需要考虑，在等待串口返回过程中检查队列
type cmdptr *Command

//BQueue Base Queue
type BQueue struct {
	MDList []MDevice
	Cmds   [][]cmdptr
}

// 将BQueue 基于新的MD改造为MDN模式，就是需要的工作的

type mdnCycTable struct {
	mdnHead    MDNode
	length     int
	cmdlength  int
	currentmdn *MDevice
	currentcmd *Command
}

//MDivce 是属于配置系统中的文件，并非系统核心工作文件，需要创新新的核心文件，device.cmd 结构才正常

// MDNode Modbus Device Node
type MDNode struct {
	mid    MDevice
	mdnptr *MDNode
}

// channel 应该是client的一部分，自动化的那一部分，client.start 就可以运行，client.stop 停止 client.addtohi client.addlowqueue
// 队列应该是绑定要cmd比较合适，那就必须改造cmd，让其具有设备属性，负责更新值就会产生问题
// 当基于cmd时，临时队列就不会有设备

//不好……client 检查设备大队列，从大队列中挑选设备用于准备通讯，再检查每大队列中的每个小队列，都必须检查最有队列是否有数据，如果没有，则执行当前队列；
// 更好的思路是，记录每个设备cmd长度，记录设备，生成虚拟cmd队列，[MD][CMD]*CMD

type channeltype struct {
	rs485         bool //bit0
	modbustcp     bool //bit1
	modbusudp     bool //bit2
	modbusovertcp bool //bit3
	modbusoverudp bool //bit4
}

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

// MeasureValue a measure value with deal, it inclues data and value status
// cType is data pull from interface{} and convert rule, include int16/int32/int64/float32/float64/byte/bool
// phyTypeID is physical type id
// value is value
// status means the value trusted state 0 is normal, 1 is ... &~&
type MeasureValue struct {
	cType     string      // 数据传递解析类型，是TagParse解析的子集，用于表达数据，语言层面的数据表达解析
	phyTypeID PhyType     // 物理属性，用于上传标记，从协议中获取
	value     interface{} // 数据，根据TagParse生成原始数据
	state     uint8       // 数据传输或获取状态码
}

// RealTimeVarTable real time varible tag value table
// a map for key:value
// key:devName_Addr_channel.tag
type RealTimeVarTable map[string]MeasureValue

// 解析完没有错误，需要注册全局变量,一个全局名称包括 devName_Addr_channel.tag

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

//PhyType data type identical for physical parameter &~&
type PhyType struct {
	Name  string `json:"name"`            //物理参数名称
	Group string `json:"group"`           //单位分组
	Parid uint16 `json:"id"`              //物理参数id
	Alias string `json:"alias,omitempty"` //助记符，或英文字母标识，可省略
	Unit  string `json:"unit,omitempty"`  //物理量单位，默认为1
}

//Command one command is modbus or rs485 protocol command for get a/some value
type Command struct {
	Alias     string `json:"name"`          //仅用于助记
	ID        string `json:"id"`            //用于索引一条命令，在同一个协议内不允许重复
	Funcode   uint   `json:"funcode"`       //Modbus/PLC 功能码
	StartAddr uint   `json:"startaddr"`     //Modbus寄存器地址，针对PLCStyle/0xFFFF
	RegLength uint   `json:"reglength"`     //Modbus寄存器长度
	Buf       []byte `json:"buf,omitempty"` //Buf 配置情况下，一般为空，仅用于写入情况
}

// 由于各个字段不同情况范围不一样，那么由业务检查，代码里不做语言要求

//根据各个字段索引关系，规定如下
//1. Parameter 参数由软件自带，给出常用参数向及内容(配置界面1)
//2. 建立协议,意味着命令与解析规则同时建立并导入（配置界面2）
//3. 建立信道及设备编号表，导入(建立)
