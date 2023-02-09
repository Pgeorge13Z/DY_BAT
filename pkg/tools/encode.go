package tools

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
)

func Md5Util(password string, salt string) string {
	var bt bytes.Buffer
	bt.WriteString(salt)
	bt.WriteString(password)

	str := bt.String()

	data := []byte(str)
	//has := md5.Sum(data)
	////md5可以改成bcrypt或者scrypt,bcrypt和scrypt的慢速hash算法可以有效防止暴力破解密码
	//md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制

	//用bcrypt加密
	hash, _ := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	//用scrypt加密
	//hash, _ := scrypt.Key(data, []byte(salt), 16384, 8, 1, 32)
	return string(hash)
}
