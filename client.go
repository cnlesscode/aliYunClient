package aliYunClient

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	_client "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AliYunClient struct {
	Client *_client.Client
	Config map[string]string
}

var aliYunClientMap = map[string]*AliYunClient{}

// 初始化阿里云客户端
func NewAliYunClient(config map[string]string) (*AliYunClient, error) {
	if aliYunClientMap[config["AccessKeyId"]] != nil {
		return aliYunClientMap[config["AccessKeyId"]], nil
	}

	configIN := &openapi.Config{
		AccessKeyId:     tea.String(config["AccessKeyId"]),
		AccessKeySecret: tea.String(config["AccessKeySecret"]),
	}
	configIN.Endpoint = tea.String(config["Endpoint"])
	_result := &_client.Client{}
	var _err error
	_result, _err = _client.NewClient(configIN)
	if _err != nil {
		return nil, _err
	}
	return &AliYunClient{Client: _result, Config: config}, nil
}
