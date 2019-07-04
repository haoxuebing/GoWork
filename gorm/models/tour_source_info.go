package models

import "time"

type TourSourceInfo struct {
	Id               int       `orm:"column(Id)" json:"Id"`                             // 官网ID
	ChineseName      string    `orm:"column(ChineseName)" json:"ChineseName"`           // 中文名
	EnglishName      string    `orm:"column(EnglishName)" json:"EnglishName"`           // 英文名
	Website          string    `orm:"column(Website)" json:"Website"`                   // 官网地址
	Languages        string    `orm:"column(Languages)" json:"Languages"`               // 语种，多个用|分割
	CategoryOne      string    `orm:"column(CategoryOne)" json:"CategoryOne"`           // 一级类别
	CategoryTwo      string    `orm:"column(CategoryTwo)" json:"CategoryTwo"`           // 二级类别
	Country          string    `orm:"column(Country)" json:"Country"`                   // 国家
	Province         string    `orm:"column(Province)" json:"Province"`                 // 省、州、大区
	City             string    `orm:"column(City)" json:"City"`                         // 城市
	Longitude        string    `orm:"column(Longitude)" json:"Longitude"`               // 经度
	Latitude         string    `orm:"column(Latitude)" json:"Latitude"`                 // 纬度
	Score            string    `orm:"column(Score)" json:"Score"`                       // 评分
	Rank             string    `orm:"column(Rank)" json:"Rank"`                         // 排序
	SupportBooking   string    `orm:"column(SupportBooking)" json:"SupportBooking"`     // 是否支持预定
	IsHighlight      string    `orm:"column(IsHighlight)" json:"IsHighlight"`           // 是否高亮
	IsRecommend      string    `orm:"column(IsRecommend)" json:"IsRecommend"`           // 是否推荐 0,首页普通;1,列表推荐;2,首页LOGO
	LogoUrl          string    `orm:"column(LogoUrl)" json:"LogoUrl"`                   // Logo
	CulturalHeritage string    `orm:"column(CulturalHeritage)" json:"CulturalHeritage"` // 是否文化遗产
	Introduce        string    `orm:"column(Introduce)" json:"Introduce"`               // 介绍
	WeChatURL        string    `orm:"column(WeChatURL)" json:"WeChatURL"`               // 微信地址
	SinaURL          string    `orm:"column(SinaURL)" json:"SinaURL"`                   // 新浪微博
	FaceBookURL      string    `orm:"column(FaceBookURL)" json:"FaceBookURL"`           // FaceBook
	TiwtterURL       string    `orm:"column(TiwtterURL)" json:"TiwtterURL"`             // tiwtter
	GoogleURL        string    `orm:"column(GoogleURL)" json:"GoogleURL"`               // google+
	AndroidAppURL    string    `orm:"column(AndroidAppURL)" json:"AndroidAppURL"`       // Android下载地址
	IosAppURL        string    `orm:"column(IosAppURL)" json:"IosAppURL"`               // IOS下载地址
	HotLine          string    `orm:"column(HotLine)" json:"HotLine"`                   // 热线
	WhoCreated       string    `orm:"column(WhoCreated)" json:"WhoCreated"`             // 创建人
	WhoUpdated       string    `orm:"column(WhoUpdated)" json:"WhoUpdated"`             // 最后修改人
	IsDelete         string    `orm:"column(IsDelete)" json:"IsDelete"`                 // 是否删除
	IsDisplay        string    `orm:"column(IsDisplay)" json:"IsDisplay"`               // 是否显示
	CreatedTime      time.Time `orm:"column(CreatedTime)" json:"CreatedTime"`           // 创建时间
	LastUpdatedTime  time.Time `orm:"column(LastUpdatedTime)" json:"LastUpdatedTime"`   // 最后更新时间
}

func (*TourSourceInfo) TableName() string {
	return "TourSourceInfo"
}
