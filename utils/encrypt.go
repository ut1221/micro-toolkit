// Package utils
// @Author PTJ 2024-05-14 17:54:53
// @Description:
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt
//
//	@Description: 给密码加密
//	@Author PTJ 2024-05-14 17:59:04
//	@param pwd
//	@return string
func Encrypt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// Verify
//
//	@Description: 校验密码
//	@Author PTJ 2024-05-14 17:59:09
//	@param pwd1
//	@param pwd2
//	@return bool
func Verify(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
