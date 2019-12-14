package modbusy

//PhyType data type identical for physical parameter &~&
type PhyType struct {
	Name  string `json:"name"`            //物理参数名称
	Group string `json:"group"`           //单位分组
	Parid uint16 `json:"id"`              //物理参数id
	Alias string `json:"alias,omitempty"` //助记符，或英文字母标识，可省略
	Unit  string `json:"unit,omitempty"`  //物理量单位，默认为1
}

var phytypelist = []PhyType{
	//能量
	PhyType{
		Name:  "正向有功电度",
		Group: "能量_电",
		Parid: 1000,
		Alias: "Ep+",
		Unit:  "kWh",
	},
	PhyType{
		Name:  "正向无功",
		Group: "能量_电",
		Parid: 1001,
		Alias: "EQ+",
		Unit:  "kVarh",
	},

	//功率单位
	PhyType{
		Name:  "正向有功功率",
		Group: "功率_电",
		Parid: 1100,
		Alias: "Σp",
		Unit:  "W",
	},
	PhyType{
		Name:  "A相有功功率",
		Group: "功率_电",
		Parid: 1101,
		Alias: "Pa",
		Unit:  "kW",
	},
	PhyType{
		Name:  "B相功率",
		Group: "功率_电",
		Parid: 1102,
		Alias: "Pb",
		Unit:  "kW",
	},
	PhyType{
		Name:  "C相功率",
		Group: "功率_电",
		Parid: 1103,
		Alias: "Pc",
		Unit:  "kW",
	},
	PhyType{
		Name:  "正向无功功率",
		Group: "功率_电",
		Parid: 1110,
		Alias: "EQ",
		Unit:  "kVar",
	},

	//用电状态

	//用电状态_电压:120x
	PhyType{
		Name:  "A相电压",
		Group: "用电状态_电压",
		Parid: 1200,
		Alias: "Ua",
		Unit:  "V",
	},
	PhyType{
		Name:  "B相电压",
		Group: "用电状态_电压",
		Parid: 1201,
		Alias: "Ub",
		Unit:  "V",
	},
	PhyType{
		Name:  "C相电压",
		Group: "用电状态_电压",
		Parid: 1202,
		Alias: "Uc",
		Unit:  "V",
	},
	PhyType{
		Name:  "AB相电压",
		Group: "用电状态_电压",
		Parid: 1203,
		Alias: "Uab",
		Unit:  "V",
	},
	PhyType{
		Name:  "AC相电压",
		Group: "用电状态_电压",
		Parid: 1204,
		Alias: "Uac",
		Unit:  "V",
	},
	PhyType{
		Name:  "BC相电压",
		Group: "用电状态_电压",
		Parid: 1205,
		Alias: "Ubc",
		Unit:  "V",
	},
	// 用电状态_电流:121x
	PhyType{
		Name:  "A相电流",
		Group: "用电状态_电流",
		Parid: 1210,
		Alias: "Ia",
		Unit:  "A",
	},
	PhyType{
		Name:  "B相电流",
		Group: "用电状态_电流",
		Parid: 1211,
		Alias: "Ib",
		Unit:  "A",
	},
	PhyType{
		Name:  "C相电流",
		Group: "用电状态_电流",
		Parid: 1212,
		Alias: "Ic",
		Unit:  "A",
	},
	//用电状态_功率因素:122x
	PhyType{
		Name:  "功率因素",
		Group: "用电状态_功率因素",
		Parid: 1220,
		Alias: "PF",
		Unit:  "",
	},
	PhyType{
		Name:  "A相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1221,
		Alias: "PFa",
		Unit:  "",
	},
	PhyType{
		Name:  "B相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1222,
		Alias: "PFb",
		Unit:  "",
	},
	PhyType{
		Name:  "C相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1223,
		Alias: "PFc",
		Unit:  "",
	},
	// 用电状态_谐波:1250+~1299

	//水 14xx
	PhyType{
		Name:  "累计用水量",
		Group: "水",
		Parid: 1400,
		Alias: "PFc",
		Unit:  "",
	},
	//气 15xx
	PhyType{
		Name:  "累计用气量",
		Group: "气",
		Parid: 1500,
		Alias: "PFc",
		Unit:  "",
	},
	//冷热 16xx
	PhyType{
		Name:  "累计冷热量",
		Group: "能量",
		Parid: 1600,
		Alias: "ΣHC+",
		Unit:  "MJ",
	},
	PhyType{
		Name:  "累计用冷量",
		Group: "能量_冷",
		Parid: 1610,
		Alias: "ΣC+",
		Unit:  "MJ",
	},
	PhyType{
		Name:  "累计用冷量",
		Group: "能量_热",
		Parid: 1640,
		Alias: "ΣH+",
		Unit:  "MJ",
	},
}
