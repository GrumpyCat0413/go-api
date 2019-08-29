package auth

import "golang.org/x/crypto/bcrypt"

// 使用bcrypt 加密明文
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

//将 加密文本与纯文本进行比较 如果相同 返回nil
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
