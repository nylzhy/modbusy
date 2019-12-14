package modbusy

// RawClient a modbus raw client read/write operation collection for modbus instruction with modbus protocol style
type RawClient interface {
	// 寄存器类型			功能				功能码		寄存器地址范围		读取/写入寄存器数量			操作数据
	// 线圈				读取线圈				0x01		0x0000-0xFFFF		1-2000(0x001~0x7D0)			bit
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

// Client a high api client but without flexible
type Client interface {
	// Bit access

	// ReadCoils reads from 1 to 2000 contiguous status of coils in a
	// remote device and returns coil status.
	ReadCoils(address, quantity uint16) (results []byte, err error)
	// ReadDiscreteInputs reads from 1 to 2000 contiguous status of
	// discrete inputs in a remote device and returns input status.
	ReadDiscreteInputs(address, quantity uint16) (results []byte, err error)
	// WriteSingleCoil write a single output to either ON or OFF in a
	// remote device and returns output value.
	WriteSingleCoil(address, value uint16) (results []byte, err error)
	// WriteMultipleCoils forces each coil in a sequence of coils to either
	// ON or OFF in a remote device and returns quantity of outputs.
	WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error)

	// 16-bit access

	// ReadInputRegisters reads from 1 to 125 contiguous input registers in
	// a remote device and returns input registers.
	ReadInputRegisters(address, quantity uint16) (results []byte, err error)
	// ReadHoldingRegisters reads the contents of a contiguous block of
	// holding registers in a remote device and returns register value.
	ReadHoldingRegisters(address, quantity uint16) (results []byte, err error)
	// WriteSingleRegister writes a single holding register in a remote
	// device and returns register value.
	WriteSingleRegister(address, value uint16) (results []byte, err error)
	// WriteMultipleRegisters writes a block of contiguous registers
	// (1 to 123 registers) in a remote device and returns quantity of
	// registers.
	WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error)
	// ReadWriteMultipleRegisters performs a combination of one read
	// operation and one write operation. It returns read registers value.
	ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error)
	// MaskWriteRegister modify the contents of a specified holding
	// register using a combination of an AND mask, an OR mask, and the
	// register's current contents. The function returns
	// AND-mask and OR-mask.
	MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error)
	//ReadFIFOQueue reads the contents of a First-In-First-Out (FIFO) queue
	// of register in a remote device and returns FIFO value register.
	ReadFIFOQueue(address uint16) (results []byte, err error)
}

// configure not conside right
type comchannel struct {
	highQueue *[]int
	lowQueue  *[]int
	//每个通道下面必须包含一个收发器，必须提供一个数据接入通道，一个数据处理通道，两个队列
	//
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
