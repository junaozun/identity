/**
  定义在命令行中所需要的参数
 */

package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID     string//通道名称（mychannel）
	ChannelConfig string//通道交易配置文件（./channel-artifacts/channel.tx）
	OrgAdmin      string//已什么身份创建的
	OrgName       string//组织名称（org1）
	OrdererOrgName	string//Orderer地址及端口号
	OrgResMgmt *resmgmt.Client//org组织的资源管理客户端对象（用于创建通道，加入节点，安装实例化链码）

	ChaincodeID	string// 链码名称
	ChaincodeGoPath	string//链码的gopath路径
	ChaincodePath	string//链码所在路径
	UserName	string//用户
}
