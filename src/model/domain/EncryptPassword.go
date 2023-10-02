package domain

import (
	"crypto/md5"
	"encoding/hex"
)

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.GetPassword()))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
