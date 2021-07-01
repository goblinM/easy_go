package jenkins

type (
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
		HideNickname     string `json:"hide_nickname"`
	}

	ParamterList struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}
)
