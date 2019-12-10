package modbusy

var phytypelist = []PhyType{
	//能量
	PhyType{
		Name:  "正向有功电度",
		Group: "能量_电",
		Parid: 1000,
		Alias: "Ep+",
		unit:  "kWh",
	},
	PhyType{
		Name:  "正向无功",
		Group: "能量_电",
		Parid: 1001,
		Alias: "EQ+",
		unit:  "kVarh",
	},

	//功率单位
	PhyType{
		Name:  "正向有功功率",
		Group: "功率_电",
		Parid: 1100,
		Alias: "Σp",
		unit:  "W",
	},
	PhyType{
		Name:  "A相有功功率",
		Group: "功率_电",
		Parid: 1101,
		Alias: "Pa",
		unit:  "kW",
	},
	PhyType{
		Name:  "B相功率",
		Group: "功率_电",
		Parid: 1102,
		Alias: "Pb",
		unit:  "kW",
	},
	PhyType{
		Name:  "C相功率",
		Group: "功率_电",
		Parid: 1103,
		Alias: "Pc",
		unit:  "kW",
	},
	PhyType{
		Name:  "正向无功功率",
		Group: "功率_电",
		Parid: 1110,
		Alias: "EQ",
		unit:  "kVar",
	},

	//用电状态

	//用电状态_电压:120x
	PhyType{
		Name:  "A相电压",
		Group: "用电状态_电压",
		Parid: 1200,
		Alias: "Ua",
		unit:  "V",
	},
	PhyType{
		Name:  "B相电压",
		Group: "用电状态_电压",
		Parid: 1201,
		Alias: "Ub",
		unit:  "V",
	},
	PhyType{
		Name:  "C相电压",
		Group: "用电状态_电压",
		Parid: 1202,
		Alias: "Uc",
		unit:  "V",
	},
	PhyType{
		Name:  "AB相电压",
		Group: "用电状态_电压",
		Parid: 1203,
		Alias: "Uab",
		unit:  "V",
	},
	PhyType{
		Name:  "AC相电压",
		Group: "用电状态_电压",
		Parid: 1204,
		Alias: "Uac",
		unit:  "V",
	},
	PhyType{
		Name:  "BC相电压",
		Group: "用电状态_电压",
		Parid: 1205,
		Alias: "Ubc",
		unit:  "V",
	},
	// 用电状态_电流:121x
	PhyType{
		Name:  "A相电流",
		Group: "用电状态_电流",
		Parid: 1210,
		Alias: "Ia",
		unit:  "A",
	},
	PhyType{
		Name:  "B相电流",
		Group: "用电状态_电流",
		Parid: 1211,
		Alias: "Ib",
		unit:  "A",
	},
	PhyType{
		Name:  "C相电流",
		Group: "用电状态_电流",
		Parid: 1212,
		Alias: "Ic",
		unit:  "A",
	},
	//用电状态_功率因素:122x
	PhyType{
		Name:  "功率因素",
		Group: "用电状态_功率因素",
		Parid: 1220,
		Alias: "PF",
		unit:  "",
	},
	PhyType{
		Name:  "A相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1221,
		Alias: "PFa",
		unit:  "",
	},
	PhyType{
		Name:  "B相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1222,
		Alias: "PFb",
		unit:  "",
	},
	PhyType{
		Name:  "C相功率因素",
		Group: "用电状态_功率因素",
		Parid: 1223,
		Alias: "PFc",
		unit:  "",
	},
	// 用电状态_谐波:1250+~1299

	//水 14xx
	PhyType{
		Name:  "累计用水量",
		Group: "水",
		Parid: 1400,
		Alias: "PFc",
		unit:  "",
	},
	//气 15xx
	PhyType{
		Name:  "累计用气量",
		Group: "气",
		Parid: 1500,
		Alias: "PFc",
		unit:  "",
	},
	//冷热 16xx
	PhyType{
		Name:  "累计冷热量",
		Group: "能量",
		Parid: 1600,
		Alias: "ΣHC+",
		unit:  "MJ",
	},
	PhyType{
		Name:  "累计用冷量",
		Group: "能量_冷",
		Parid: 1610,
		Alias: "ΣC+",
		unit:  "MJ",
	},
	PhyType{
		Name:  "累计用冷量",
		Group: "能量_热",
		Parid: 1640,
		Alias: "ΣH+",
		unit:  "MJ",
	},
}
