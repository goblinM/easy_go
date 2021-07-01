package jenkins

import (
	"encoding/json"
	"fmt"
)

// 创建构建一个job任务
func CreateBuildJob() {
	cd := JenkinsRequest{
		SoundPluginUrl:   "https://live-d.u.ccb.com/html/soundflower.html",
		Package:          "CCBliveSetup",
		Version:          "4.4.1",
		MacAppEn:         "直播助手",
		MacApp:           "直播助手",
		UserId:           "b42ebd056dds30",
		WinApp:           "直播助手",
		Forget:           "https://live-manager.u.ccb.com/html/newclientuse.html",
		WinIcon:          "https://dd.polyv.net/packages/custom_download/img_synch/ddd/win_icon/直播助手icon-256.png",
		WinBack:          "https://dd.polyv.net/packages/custom_download/img_synch/ddd/win_back/pic.png",
		MacIcon:          "https://dd.polyv.net/packages/custom_download/img_synch/ddd/mac_icon/icon-logo.png",
		MacBack:          "https://dd.polyv.net/packages/custom_download/img_synch/ddd/mac_back/dmg_bg@2x.png",
		LoginImg:         "https://dd.polyv.net/packages/custom_download/img_synch/ddd/login_img/1608623582091_logo.png",
		Title:            "个人信息保护指引",
		ProtocolName:     "用户服务协议",
		ProtocolUrl:      "https://u.ccb.com/trabase/#/homelink/register/agreeservicereceivV2",
		PrivacyName:      "隐私政策",
		PrivacyUrl:       "https://u.ccb.com/trabase/#/homelink/register/agreeprivacyreceivV3",
		IsNotSign:        false,
		IsMac:            true,
		IsWin:            false,
		IsCcb:            true,
		Scheme:           "ccblive",
		IsDisableProcess: false,
		IsHidePrivacy:    true,
		DrawScene:        false,
		UpdateLog:        "",
		ClientHost:       "",
		HideNickname:     "N",
	}
	// 结构体转map
	data, _ := json.Marshal(&cd)
	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &dataMap); err != nil {
		fmt.Println(err)
	}
	parameterList := make([]*ParamterList, 0)
	for k, v := range dataMap {
		add := ParamterList{
			Name:  k,
			Value: v,
		}
		parameterList = append(parameterList, &add)
	}
	payload := map[string][]*ParamterList{
		"parameter": parameterList,
	}
	fmt.Println(payload)
	if pl, err := json.Marshal(payload); err != nil {
		fmt.Println(err)
	} else {
		suffixPath := "view/windows_client/job/custom_client"
		parameters := fmt.Sprintf("json=%s", string(pl))
		res := NewJenkinsJob().BuildJob(suffixPath, parameters)
		fmt.Println(res)
	}
}
