package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func EmailUtil() {
	em := email.NewEmail()

	em.From = "WePlay <xxx@163.com>"

	em.To = []string{"ooo@qq.com"}

	em.Subject = "WePlay 验证码"

	em.Text = []byte("邮件来自 WePlay\n本次验证码是: 2389\n请勿回复本邮件!\n")

	err := em.Send("smtp.163.com:25", smtp.PlainAuth("", "xxx@163.com", "xxxooo", "smtp.163.com"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	EmailUtil()
}
