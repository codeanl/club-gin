package util

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"school-bbs/conf"
	"strconv"
	"time"
)

// GetRand
// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, authcode string) error {
	e := email.NewEmail()
	e.From = "校园论坛注册验证码 <" + conf.SmtpEmail + ">"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + authcode + "</b>" + "！ 欢迎您使用校园论坛")
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", conf.SmtpEmail, conf.SmtpPass, conf.SmtpHost),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}
