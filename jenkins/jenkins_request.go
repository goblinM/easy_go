package jenkins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const USERNAME = ""
const TOKEN = ""
const HOST = "http://jenkins.igeeker.org"

// jenkins 调用定制项目
type (
	JenkinsInterface interface {
		BuildJob(suffixPath, data string) string // 创建jenkins任务
		GetQueueId(uuid string) error            // 获取队列中的queue id
		GetLastQueueId(uuid string) error        // 获取最后构建的queue id
		GetQueueStatus(queueId string) error     // 获取队列状态
		GetJobStatus(jobId string) error         // 获取job 状态
	}

	defaultJenkins struct {
		UserName string
		Token    string
		Host     string
	}
)

func NewJenkinsJob() JenkinsInterface {
	return defaultJenkins{
		UserName: USERNAME,
		Token:    TOKEN,
		Host:     HOST,
	}
}

// 创建jenkins任务
func (d defaultJenkins) BuildJob(suffixPath string, data string) string {
	uri := fmt.Sprintf("%s/%s/build/api/json", d.Host, suffixPath)
	contentType := "application/x-www-form-urlencoded"
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", uri, strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(d.UserName, d.Token)
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 201 {
		fmt.Println("构建jenkins job 失败")
		return "构建jenkins job 失败"
	} else {
		fmt.Println("构建jenkins job 成功")
		return string(result)
	}
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
