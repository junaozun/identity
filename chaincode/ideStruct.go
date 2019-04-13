package main

type Identity struct {
	ObjectType    string    `json:"docType"`//couchdb数据库富查询使用，相当于文档类型
	Name    string    `json:"Name"`        // 姓名
	Sex    string    `json:"Sex"`        // 性别
	Nation    string    `json:"Nation"`        // 民族
	IDNumber    string    `json:"IDNumber"`        // 身份证号
	NativePlace    string    `json:"NativePlace"`        // 籍贯
	BirthDay    string    `json:"BirthDay"`        // 出生日期

	EnrollDate    string    `json:"EnrollDate"`        // 入学日期
	GraduationDate    string    `json:"GraduationDate"`    // 毕（结）业日期
	SchoolName    string    `json:"SchoolName"`    // 学校名称
	Major    string    `json:"Major"`    // 专业
	QuaType    string    `json:"QuaType"`    // 学历类别
	Length    string    `json:"Length"`    // 学制
	Mode    string    `json:"Mode"`    // 学习形式
	Level    string    `json:"Level"`    // 层次
	Graduation    string    `json:"Graduation"`    // 毕（结）业
	CertNo    string    `json:"CertNo"`    // 证书编号

	Photo    string    `json:"Photo"`    // 照片

	Historys    []HistoryItem    // 当前身份的历史记录
}

type HistoryItem struct {
	TxId    string			//交易ID
	Identity    Identity	//一个人的身份信息对象
}
