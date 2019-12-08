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

type comchannel struct {
	HiQueue   *[]int
	LoQueue   *[]int
	configure int
}

type deviceInfo struct {
	Name     string
	Addr     string //无论任何时候都是从站地址
	Enable   bool   //是否启用
	Channel  string //通道名称或编码
	Protocol string // 协议iD
	// ParseRule string //解析规则名称或编码
	// Cmds      string // 命令规则，是否和解析规则合并
}

type bytesptr *[]byte

type protocol struct {
	Name      string
	ID        string
	Cmds      []command
	ParseRule [][]IndexParse //第一层用于和cmds匹配，一个cmd需要多个tag
}

// 解析完没有错误，需要注册全局变量
// 按照dev_addr.tag_cmd 或者全局ID生成唯一索引

// Value a value and value status
type Value struct {
	vtype   string
	partype parametertype
	value   interface{}
	status  uint8
}

//RealTimeVarTable real time value string 应该是全局名称+ID
//一个全局名称包括信道_设备名称_参数名称
type RealTimeVarTable map[string]Value

//其他信息归上层服务器管理，例如设备安装位置，厂家、编号等等，与收发无关的信息，该scada的作用就是让采集信息有意义，但是全部意义

// IndexParse Indexarse
type IndexParse struct {
	DevCmdID   string  //"Dev.Cmd"
	Tag        string  `json:"tag,omitempty"`
	ParType    string  // 参数类型或参数编码
	StartIndex uint    `json:"start_index,omitempty"`
	EndIndex   uint    `json:"end_index,omitempty"`
	Offset     float64 `json:"offset,omitempty"`
	DataType   uint8   `json:"data_type,omitempty"`
	PT         uint    `json:"pt,omitempty"`
	CT         uint    `json:"ct,omitempty"`
	//Value      float64 `json:"value,omitempty"` // value = (d-offset)*pt/ct
}

type parametertype struct {
	Name  string
	parid uint16
	unit  string
}

// 由于各个字段不同情况范围不一样，那么由业务检查，代码里不做语言要求
type command struct {
	Name      string
	ID        string
	funcode   uint
	StartAddr uint //正常情况下uint16足够，但是对于PLC系统，则需要uint才能保证
	RegLength uint
	Buf       bytesptr
}
