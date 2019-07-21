package biz

import (	
  "fmt"
  "encoding/json"
	"booksapi/src/models/orm"
  "booksapi/src/models/vo"
  "booksapi/src/models/enc"
)

func GetOrCreateUser(user *orm.User, wechatSessionKey, wechatOpenID, wechatUnionID string) error {	
	if ok := wechatSessionKey != "" && wechatOpenID != ""; !ok {
		return fmt.Errorf("wechat sessionkey [%s] and openid [%s] should both have value", wechatSessionKey, wechatOpenID)
	}

	var entity orm.User
	if err := entity.FindByWechatOpenID(user, wechatOpenID); err != nil {
		user.WechatSessionKey = wechatSessionKey
		user.WechatOpenID = wechatOpenID
		user.WechatUnionID = wechatUnionID

		return entity.Create(user)
	}
	
	user.WechatSessionKey = wechatSessionKey
	user.WechatUnionID = wechatUnionID
	return entity.Update(user)
}

func SaveWXUserInfo(rawData vo.WXData, userID uint64, openID string) (*orm.User, error) {
	var user orm.User
	if err := user.Get(&user, userID); err != nil {		
		if err := user.FindByWechatOpenID(&user, openID); err != nil {
			return nil, err
		}
	}

	if jsonStr, err := enc.DecryptWXData(user.WechatSessionKey, rawData.IV, rawData.EncryptedData); err != nil {
		return nil, err
	} else {
		fmt.Println(jsonStr)		
		var wechatUser orm.WechatUser
		if err := json.Unmarshal([]byte(jsonStr), &wechatUser); err != nil {
			return nil, err
		}

		// set openid
		wechatUser.OpenID = user.WechatOpenID
		wechatUser.UnionID = user.WechatUnionID
		if err := wechatUser.UpdateOrCreateByOpenID(&wechatUser); err != nil {
			return nil, err
		}

		// set nickname and avatar
		user.NickName = wechatUser.NickName
		user.AvatarURL = wechatUser.AvatarURL

		if err := user.Update(&user); err != nil {
			return nil, err
		}

		return &user, nil
		// }
	}
}

func GetUserByID(id uint64) (*orm.User, error) {
  var user orm.User
  if err := user.Get(&user, id); err != nil {
    return nil, err
  }

  return &user, nil
}