package modbusy

//负责app核心数据结构及方法

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

// ModDevice Modbus Device Model
type ModDevice struct {
}
