package sender

import (
	"errors"
	"fmt"
	"watchAlert/pkg/ctx"
	"watchAlert/pkg/provider"
	"watchAlert/pkg/provider/aliyun"
)

// PhoneCallSender 邮件发送策略
type PhoneCallSender struct{}

func NewPhoneCallSender() SendInter {
	return &PhoneCallSender{}
}

func (e *PhoneCallSender) Send(params SendParams) error {
	setting, err := ctx.DB.Setting().Get()
	if err != nil {
		return errors.New("获取系统配置失败: " + err.Error())
	}

	var phoneCall provider.PhoneCall
	switch setting.PhoneCallConfig.Provider {
	case provider.PROVIDER_ALIYUN:
		aliyunPhoneCall := &aliyun.PhoneCall{
			Endpoint:        setting.PhoneCallConfig.Endpoint,
			AccessKeyId:     setting.PhoneCallConfig.AccessKeyId,
			AccessKeySecret: setting.PhoneCallConfig.AccessKeySecret,
		}
		err := aliyunPhoneCall.CreateClient()
		if err != nil {
			return fmt.Errorf("创建%s语音服务客户端失败: %v\n", setting.PhoneCallConfig.Provider, err)
		}
		phoneCall = aliyunPhoneCall
	}

	err = phoneCall.Call(params.Content, params.PhoneNumber)
	if err != nil {
		return errors.New("语音通知 类型报警发送失败" + err.Error())
	}

	return nil
}
