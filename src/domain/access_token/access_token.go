package access_token

import "time"

// Access Token struct
// // AccessToken generate when user and password is correct and successed
// // UserId Identify userafter authen is successed
// // ClientId for indentify client Ex. id: 1 = web, id 2 = mobile
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

const (
	//Expires 24 hours
	expirationTime = 24
)

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.unix(at.isExpired(), 0).Before(time.Now().UTC())
}
