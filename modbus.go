package modbusy

// Modbus APU/PDU max length
const (
	MaxADULength = 256 //MaxADULength ADU 最大长度
	SerialMaxPDU = 253 //SerialMaxPDU 串行链路上最大的PDU长度
	TCPMaxPDU    = 249 //TCPMaxPDU TCP链路上最大的PDU长度
)

// FC Function Code Definition 功能码定义
const (
	//Register
	FCReadCoils          = 0x01 // FCReadCoils: Read coil status/DO /读取线圈状态 function code 1
	FCReadDiscreteInputs = 0x02 // FCReadCoils: Read input status/DI /读取输入离散量状态 function code 2
	FCReadHoldingReg     = 0x03 // FCReadHoldingReg: Read holding registers /读取保持寄存器 function code 3
	FCReadInputReg       = 0x04 // FCReadInputReg: Read input registers /读取输入寄存器 function code 4
	FCWriteSingleCoil    = 0x05 // FCWriteSingleCoil: Force single coil /强制写单个线圈  function code 5
	FCWriteSingleReg     = 0x06 // FCWriteSingleReg: Preset single register /预置单个寄存器 function code 6
	FCWriteMultipleCoils = 0x0F // FCWriteMultipleCoils: Write Multiple Coils /预置多个线圈 function code 15
	FCWriteMultipleReg   = 0x10 // FCWriteMultipleReg: Write Multiple registers /写多个寄存器 function code 16

	//File Operation
	FCReadFileRecord       = 0x14 // FCReadFileRecord: Read File Record /读文件记录 function code 20
	FCWriteFileRecord      = 0x15 // FCWriteFileRecord: Write File Record /写文件记录 function code 21
	FCMaskWriteReg         = 0x16 // FCMaskWriteReg:  Mask Write Register /屏蔽写寄存器 function code 22
	FCReadWriteMultipleReg = 0x17 // FCReadWriteMultipleReg: Read/Write Multiple registers /读写多个寄存器 function code 23
	FCReadFIFOQueue        = 0x18 // FCReadFIFOQueue:  Read FIFO Queue /读取文件队列 function code 24

	// Serial Line only
	FCReadException   = 0x07 // FCReadExcepStatus: Read Exception Status (Serial Line only) function code 7
	FCDiagnostics     = 0x08 // FCDiagnostics: Diagnostics (Serial Line only) function code 8
	FCGetEventCounter = 0x0B // FCGetEventCounter: Get Comm Event Counter (Serial Line only) function code 11
	FCGetEventLog     = 0x0C // FCGetEventLogr:  Get Comm Event Log (Serial Line only) function code 12
	FCReportServerID  = 0x11 // FCReportServerID: Report Server ID (Serial Line only) function code 17
)

/*
代码	名称	含义
01	非法功能	对于服务器（或从站）来说，询问中接收到的功能码是不可允许的操作，可能是因为功能码仅适用于新设备而被选单元中不可实现同时，还指出服务器（或从站）在错误状态中处理这种请求，例如：它是未配置的，且要求返回寄存器值。
02	非法数据地址	对于服务器（或从站）来说，询问中接收的数据地址是不可允许的地址，特别是参考号和传输长度的组合是无效的。对于带有100个寄存器的控制器来说，偏移量96和长度4的请求会成功，而偏移量96和长度5的请求将产生异常码02。
03	非法数据值	对于服务器（或从站）来说，询问中包括的值是不可允许的值。该值指示了组合请求剩余结构中的故障。例如：隐含长度是不正确的。modbus协议不知道任何特殊寄存器的任何特殊值的重要意义，寄存器中被提交存储的数据项有一个应用程序期望之外的值。
04	从站设备故障	当服务器（或从站）正在设法执行请求的操作时，产生不可重新获得的差错。
05	确认收到	与编程命令一起使用，服务器（或从站）已经接受请求，并且正在处理这个请求，但是需要长持续时间进行这些操作，返回这个响应防止在客户机（或主站）中发生超时错误，客户机（或主机）可以继续发送轮询程序完成报文来确认是否完成处理。
07	从属设备忙	与编程命令一起使用，服务器（或从站）正在处理长持续时间的程序命令，当服务器（或从站）空闲时，客户机（或主站）应该稍后重新传输报文。
08	存储奇偶性差错	与功能码20和21以及参考类型6一起使用，指示扩展文件区不能通过一致性校验。服务器（或从站）设备读取记录文件，但在存储器中发现一个奇偶校验错误。客户机（或主机）可重新发送请求，但可以在服务器（或从站）设备上要求服务。
0A	不可用网关路径	与网关一起使用，指示网关不能为处理请求分配输入端口值输出端口的内部通信路径，通常意味着网关是错误配置的或过载的。
0B	网关目标设备响应失败	与网关一起使用，指示没有从目标设备中获得响应，通常意味着设备未在网络中
*/

// Exception Code(EC) Definition 错误功能码定义
const (
	ECIllegalFunction                    = 0x01 // ILLEGAL FUNCTION 非法功能
	ECIllegalDataAddress                 = 0x02 // ILLEGAL DATA ADDRESS 非法数据地址
	ECIllegalDataValue                   = 0x03 // ILLEGAL DATA VALUE 非法数据值
	ECServerDeviceFailure                = 0x04 // SERVER DEVICE FAILURE 从站设备故障
	ECAcknowledge                        = 0x05 // ACKNOWLEDGE 确认
	ECServerDeviceBusy                   = 0x06 // SERVER DEVICE BUSY 从属设备忙
	ECMemoryParityError                  = 0x08 // MEMORY PARITY ERROR  存储奇偶性差错
	ECGatewayPathUnavailable             = 0x0A // GATEWAY PATH UNAVAILABLE 不可用网关路径
	ECGatewayTargetDeviceFailedToRespond = 0x0B // GATEWAY TARGET DEVICE FAILED TO RESPOND 网关目标设备响应失败
)

// PDU protocol data unit
type PDU struct {
	FunctionCode byte
	Data         []byte
}

//RequestPDU Modbus request PDU
type RequestPDU struct {
	FunctionCode byte
	RequestData  []byte
}

//ResponsePDU Modbus response PDU
type ResponsePDU struct {
	FunctionCode byte
	RequestData  []byte
}

//ExceptionResponsePDU Modbus exception response PDU
type ExceptionResponsePDU struct {
	FunctionCode  byte
	ExceptionCode byte
}

// ADU application data unit
type ADU struct {
	AddrField []byte
	PDU
	ErrCheck []byte
}
