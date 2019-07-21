package orm

type WechatUser struct {
	OrmModel
	AvatarURL string `json:"avatarUrl"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    uint64 `json:"gender"`
	Language  string `json:"language"`
	NickName  string `json:"nickname"`
	Province  string `json:"province"`
	OpenID    string `json:"openid"`
	UnionID   string `json:"unionid"`
}

func init() {	
	db.AutoMigrate(&WechatUser{})
}

func (*WechatUser) UpdateOrCreateByOpenID(v *WechatUser) error {	
	return db.Where(WechatUser{OpenID: v.OpenID}).Assign(*v).FirstOrCreate(v).Error
}