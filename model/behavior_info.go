package model

import "danci-api/global"

// page得点击记录
type BehaviorInfo struct {
	global.GVA_MODEL

	PageUrl      string `json:"page_url"`
	UserId       string `json:"user_id"`
	ApiKey       string `json:"api_key"`
	UploadType   string `json:"upload_type"`
	HappenTime   string `json:"happen_time"`
	BehaviorType string `json:"behavior_type"`
	ClassName    string `json:"class_name"`
	Placeholder  string `json:"placeholder"`
	InputValue   string `json:"Input_value"`
	TagNameint   string `json:"tag_name"`
	InnterText   string `json:"innter_text" gorm:"type:text"`

	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}