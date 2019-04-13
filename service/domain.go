package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

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

type ServiceSetup struct {
	ChaincodeID	string //链码名称
	Client	*channel.Client //创建通道，安装实例化链码的资源管理客户端对象
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
