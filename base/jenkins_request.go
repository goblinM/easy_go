package main

import (
	"encoding/json"
	"fmt"
)

const USERNAME = ""
const TOKEN = ""

var HOST = fmt.Sprintf("http://%s:%s@jenkins.igeeker.org", USERNAME, TOKEN)

// jenkins 调用定制项目
type (
	JenkinsInterface interface {
		Create(data map[string]interface{}) error // 创建jenkins任务
		GetQueueId(uuid string) error             // 获取队列中的queue id
		GetLastQueueId(uuid string) error         // 获取最后构建的queue id
		GetQueueStatus(queueId string) error      // 获取队列状态
		GetJobStatus(jobId string) error          // 获取job 状态
	}

	defaultJenkins struct {
		UserName string
		Token    string
		Host     string
	}

	JenkinsRequest struct {
		SoundPluginUrl   string `json:"SoundPluginUrl"`
		Package          string `json:"package"`
		Version          string `json:"version"`
		UserId           string `json:"userId"`
		WinApp           string `json:"win_app"`
		MacApp           string `json:"mac_app"`
		MacAppEn         string `json:"mac_app_en"`
		Forget           string `json:"forget"`
		WinIcon          string `json:"win_icon"`
		WinBack          string `json:"win_back"`
		MacIcon          string `json:"mac_icon"`
		MacBack          string `json:"mac_back"`
		LoginImg         string `json:"login_img"`
		Title            string `json:"title"`
		ProtocolName     string `json:"protocol_name"`
		ProtocolUrl      string `json:"protocol_url"`
		PrivacyName      string `json:"privacy_name"`
		PrivacyUrl       string `json:"privacy_url"`
		IsNotSign        bool   `json:"is_not_sign"`
		IsMac            bool   `json:"is_mac"`
		IsWin            bool   `json:"is_win"`
		IsCcb            bool   `json:"is_ccb"`
		IsHidePrivacy    bool   `json:"is_hide_privacy"`
		Bit              string `json:"bit"`
		ClientHost       string `json:"client_host"`
		DrawScene        bool   `json:"drawScene"`
		IsDisableProcess bool   `json:"is_disable_process"`
		Scheme           string `json:"scheme"`
		UpdateLog        string `json:"update_log"`
	}

	ParamterList struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}
)

func NewJenkinsJob() JenkinsInterface {
	return defaultJenkins{
		UserName: USERNAME,
		Token:    TOKEN,
		Host:     HOST,
	}
}

// 生成任务相关数据
func MakeJobData() {

}

// 创建jenkins任务
func (d defaultJenkins) Create(data map[string]interface{}) error {
	uri := fmt.Sprintf("%s/view/windows_client/job/custom_client/build/api/json", d.Host)
	fmt.Println(uri)
	parameterList := make([]*ParamterList, 0)
	for k, v := range data {
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
	pl, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	resp := PostController(uri, map[string]interface{}{
		"json": string(pl),
	}, "application/json;charset=utf-8")
	//输出内容
	fmt.Println(resp)
	return nil
}

// 获取队列中的queue id
func (d defaultJenkins) GetQueueId(uuid string) error {
	panic("implement me")
}

// 获取最后构建的queue id
func (d defaultJenkins) GetLastQueueId(uuid string) error {
	panic("implement me")
}

// 获取队列状态
func (d defaultJenkins) GetQueueStatus(queueId string) error {
	panic("implement me")
}

// 获取job 状态
func (d defaultJenkins) GetJobStatus(jobId string) error {
	panic("implement me")
}

func (d defaultJenkins) StructToMap() (key string, value interface{}) {
	return "time", ""
}

func main() {
	c := NewJenkinsJob()
	cd := JenkinsRequest{
		SoundPluginUrl:   "https://live-manager.u.ccb.com/html/soundflower.html",
		Package:          "CCBliveSetup",
		Version:          "4.4.1",
		MacAppEn:         "直播助手",
		MacApp:           "直播助手",
		UserId:           "b42eb05630",
		WinApp:           "直播助手",
		Forget:           "https://live-manager.u.ccb.com/html/newclientuse.html",
		WinIcon:          "https://demo.polyv.net/packages/custom_download/img_synch/b42eb05630/win_icon/直播助手icon-256.png",
		WinBack:          "https://demo.polyv.net/packages/custom_download/img_synch/b42eb05630/win_back/pic.png",
		MacIcon:          "https://demo.polyv.net/packages/custom_download/img_synch/b42eb05630/mac_icon/icon-logo.png",
		MacBack:          "https://demo.polyv.net/packages/custom_download/img_synch/b42eb05630/mac_back/dmg_bg@2x.png",
		LoginImg:         "https://demo.polyv.net/packages/custom_download/img_synch/b42eb05630/login_img/1608623582091_logo.png",
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
	}
	// 结构体转map
	data, _ := json.Marshal(&cd)
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err)
	}
	if err := c.Create(m); err != nil {
		fmt.Println(err)
	}
}
