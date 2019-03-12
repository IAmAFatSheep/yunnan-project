package main

import "time"

//ZZ_种植户信息表
type Grower struct {
	// ID     bson.ObjectId `bson:"_id"`
	Qycode string `bson:"种植批次码"`
	Name   string `bson:"种植户户主姓名"`
	Userid string `bson:"种植户户主身份证"`
	Phone  string `bson:"联系电话"`
	Ispoor bool   `bson:"是否为平困户"`
}

//入库信息表
type StorageIn struct {
	ComponyName string    `bson:"企业名称"`
	EnterTime   time.Time `bson:"入库时间"`
	Object      string    `bson:"入库对象"`
	Taxcode     string    `bson:"商品编码"`
	Sourecode   string    `bson:"追溯码"`
}

//出库信息表
type StorageOut struct {
	ComponyName  string    `bson:"企业名称"`
	OutTime      time.Time `bson:"出库时间"`
	NewCorpName  string    `bson:"新货主企业名称"`
	WeightAll    int       `bson:"出货总量"`
	Amount       int       `bson:"合计金额"`
	Ptype        string    `bson:"出库对象"`
	Batchno      string    `bson:"药材/饮片批次号"`
	Spec         string    `bson:"规格"`
	AtchWeightAl int       `bson:"出库总重量小计"`
	UnitPrice    int       `bson:"单价"`
	SourceCode   int       `bson:"商品编码"`
	Code         string    `bson:"追溯码"`
}

//YP_饮片生产信息表
type MedSlices struct {
	Qycode          string    `bson:"饮片批次码"`
	Name            string    `bson:"饮片名称"`
	Spec            string    `bson:"生产规格"`
	Standard        string    `bson:"执行标准"`
	Processflow     string    `bson:"生产流程"`
	Operator        string    `bson:"工艺员"`
	Manager         string    `bson:"生产部门经理"`
	Qa              string    `bson:"QA"` //QA负责人名称
	Producedata     time.Time `bson:"生产日期"`
	Remark          string    `bson:"备注"`
	Weight          int       `bson:"待包装重量"`
	Medicinalcode   string    `bson:"药材批次码"`
	Code            string    `bson:"药材追溯码"`
	Medicinalweight int       `bson:"药材总重量"`
	Areacode        string    `bson:"地区编码"`
}

//YP_饮片检验信息表
type MedCheckout struct {
	Qycode         string    `bson:"饮片批次码"`
	Checkmethod    string    `bson:"检验方法"`
	Checkgrade     string    `bson:"质量级别"`
	Checkthickme   string    `bson:"粗加工方法"`
	Checkcondition string    `bson:"贮藏条件"`
	Checkuser      string    `bson:"检验人员"`
	Checktime      time.Time `bson:"检验时间"`
	Type           string    `bson:"检验报告类型"`
	Url            string    `bson:"检验报告地址"`
}

//YP_饮片赋码信息表
type MedTagging struct {
	Code           string    `bson:"饮片批次码"`
	Boxcode        string    `bson:"箱码"`
	Pcode          string    `bson:"商品编码"`
	Pname          string    `bson:"品名"`
	Areacode       string    `bson:"区域编码"`
	Spec           string    `bson:"规格"`
	Boxweight      int       `bson:"每箱重量"`
	Weight         int       `bson:"每包重量"`
	Boxtime        time.Time `bson:"包装时间"`
	ProductionCode string    `bson:"商品码"`
	SourceCode     string    `bson:"追溯码"`
}

//ZCY_中成药生产信息表
type CMedProduct struct {
	Qycode      string    `bson:"中成药批次码"`
	Name        string    `bson:"中成药名称"`
	Spec        string    `bson:"生产规格"`
	Standard    string    `bson:"执行标准"`
	Operator    string    `bson:"工艺员"`
	Manager     string    `bson:"生产部经理"`
	Qa          string    `bson:"QA"` //QA负责人名称
	Producedate time.Time `bson:"生产日期"`
	Remark      string    `bson:"备注"`
	Weight      int       `bson:"待包装重量"`
	Areacode    string    `bson:"地区编码，按国际处理"`
}

//ZCY_中成药药材信息表
type CMedMaterals struct {
	Qycode          string `bson:"中成药批次码"`
	Assistplantname string `bson:"原料名称"`
	Assistpcode     string `bson:"原料批次码"`
	Assistscode     string `bson:"原料追溯码"`
	Assistarea      string `bson:"产地编码"`
	Assistweight    int    `bson:"原料重量"`
	Assistusecount  string `bson:"原料使用比例"`
	Assistmaterial  string `bson:"原料类型"`
}

//ZCY_中成药检验信息表
type CMedCheckout struct {
	Qycode           string    `bson:"中成药批次码"`
	Checkmethod      string    `bson:"检验方法"`
	Checkgrade       string    `bson:"质量级别"`
	Checkthickmethod string    `bson:"粗加工方法"`
	Checkcondition   string    `bson:"贮藏条件"`
	Checkuser        string    `bson:"检验人员"`
	Checktime        time.Time `bson:"检验时间"`
	Type             string    `bson:"检验报告类型"`
	Url              string    `bson:"检验报告地址"`
}

//ZCY_中成药赋码信息表
type CMedTagging struct {
	Qycode    string    `bson:"饮片批次码"`
	Boxcode   string    `bson:"箱码"`
	Pcode     string    `bson:"商品编码"`
	Pname     string    `bson:"品名"`
	Areacode  string    `bson:"区域编码"`
	Spec      string    `bson:"规格"`
	Boxweight int       `bson:"每箱重量单位"`
	Weight    int       `bson:"每包重量单位"`
	Boxtime   time.Time `bson:"包装时间"`
	Code      string    `bson:"追溯码"`
}

//ZYC-上传企业基本信息表
type CompanyInfo struct {
	CorpName            string    `bson:"企业名称(中文)"`
	CorpEnName          string    `bson:"企业名称(英文)"`
	CountryID           string    `bson:"通信地址(国家)"`
	ProvinceID          string    `bson:"通信地址(省、直辖市、自治区)"`
	CityID              string    `bson:"通信地址(市、地区)"`
	RegisterAddress     string    `bson:"注册详细地址(中文)"`
	RegisterPostcode    string    `bson:"注册地邮编"`
	Orgcode             string    `bson:"组织机构代码"`
	GdsCode             string    `bson:"GDS厂商识别码"`
	BusinessCertificate string    `bson:"工商营业许可证号"`
	TaxCertificate      string    `bson:"税务登记证号"`
	OrgCertificate      string    `bson:"组织机构证书号"`
	RegisterDate        time.Time `bson:"注册日期"`
	ChanneluserID       string    `bson:"所属渠道用户ID"`
	Phone               string    `bson:"办公电话"`
	Fax                 string    `bson:"传真"`
	Mail                string    `bson:"邮箱"`
	Corptype            string    `bson:"企业类型"`
	Master              string    `bson:"企业法人"`
	Mobilephpne         string    `bson:"手机"`
	Linkman             string    `bson:"联系人"`
	Areacode            string    `bson:"所属区域代码"`
	Masterid            string    `bson:"法人身份证号码"`
	Istype              string    `bson:"是否为最下级企业"`
	Code                string    `bson:"商户所属专业市场代码"`
	Longitude           string    `bson:"地理位置经度"`
	Latitude            string    `bson:"地理位置纬度"`
}

//ZZ_种植基本信息表
type PlantInfo struct {
	Qycode        string    `bson:"种植批次码"`
	Plantname     string    `bson:"植物名称"`
	Plantsource   string    `bson:"种子幼体来源"`
	Weight        int       `bson:"种子重量"`
	Plantarea     int       `bson:"种植区域"`
	Plantcycle    int       `bson:"种植周期"`
	Areacode      string    `bson:"所属区域代码"`
	Planttime     time.Time `bson:"种植时间"`
	Plantuser     string    `bson:"负责人"`
	Address       string    `bson:"详细地址"`
	Name          string    `bson:"产出物名称"`
	Foreseeweight int       `bson:"预产重量"`
}

//ZZ_田间管理信息表
type FieldManager struct {
	Qycode  string    `bson:"种植批次码"`
	Name    string    `bson:"管理名称"`
	Type    string    `bson:"管理类型"`
	Weight  int       `bson:"用量"`
	Addtime time.Time `bson:"操作时间"`
	Remark  string    `bson:"备注"`
}

//YC_药材收货信息表
type MedMaterialReceive struct {
	Qycode      string    `bson:"药材批次码"`
	Name        string    `bson:"药材名称"`
	Weight      int       `bson:"药材重量"`
	Areacode    string    `bson:"所属区域代码"`
	Plantcode   string    `bson:"种植批次"`
	Harvesttime time.Time `bson:"收货时间"`
	Plantsource string    `bson:"种子来源"`
	Fertilizer  string    `bson:"施肥情况"`
	Pesticide   string    `bson:"喷药情况"`
	Code        string    `bson:"追溯码"`
}

//YC_药材赋码信息表
type MedMaterialTagging struct {
	Qycode   string    `bson:"药材批次码"`
	Pcode    string    `bson:"商品编码"`
	Areacode string    `bson:"区域编码"`
	Spec     int       `bson:"规格"`
	Weight   int       `bson:"重量"`
	Time     time.Time `bson:"赋码时间"`
}
