package lib

import "encoding/json"

type Cookie struct {
	ID   uint `json:"id"`
	Type int  `json:"type"`
}

// cookie加密
func CookieEncrypt(cookie Cookie) (string, error) {
	str, err := json.Marshal(cookie)
	if err != nil {
		return "", err
	}
	s := Encode(string(str))
	return s, err
}

func CookieDecrypt(s string) (*Cookie, error) {
	str := Decode(s)
	cookie := &Cookie{}
	err := json.Unmarshal([]byte(str), cookie)
	return cookie, err
}
