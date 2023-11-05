package mailer

import (
	"fmt"
	"github.com/badoux/checkmail"
	"go-boss/config"
	"gopkg.in/gomail.v2"
)

func SendText(to string, subject string, body string) error {
	// 检查电子邮件地址是否可用
	err := checkmail.ValidateFormat(to)
	if err != nil {
		fmt.Printf("email address %s is not available: %s", to, err.Error())
		return err
	}
	c := config.NewConfig()
	host := c.Viper.Get("email.host").(string)
	port := c.Viper.Get("email.port").(int)
	username := c.Viper.Get("email.username").(string)
	password := c.Viper.Get("email.password").(string)

	// 创建消息
	msg := gomail.NewMessage()
	// 设置发件人
	msg.SetHeader("From", username)
	// 设置收件人
	msg.SetHeader("To", to)
	// 设置主题
	msg.SetHeader("Subject", subject)
	// 设置正文
	msg.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
