package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"net/smtp"
	"strings"
)

type SMTPConfig struct {
	userName string `toml:"userName"`
	password string `toml:"password"`
	smtpHost string `toml:"smtpHost"`
	smtpPort string `toml:"smtpPort"`
}

type MailAddressConfig struct {
	fromAddress string `toml:"fromAddress"`
	toAddress []string `toml:"toAddress"`
}


type Config struct {
	SMTP        SMTPConfig `toml:"smtp"`
	MailAddress MailAddressConfig `toml:"mailAddress"`
}

func main() {

	// 設定の読み込み
	var config Config
	if _, tomlError := toml.DecodeFile("config.toml", &config); tomlError != nil {
		fmt.Println(tomlError)
		return
	}

	// メール内容 TODO:text/templateに変更
	subject := "タイトル"
	message := "テストメール"
	mailContent := []byte(
		"To: " + strings.Join(config.MailAddress.toAddress, ";") + "\r\n" +
			"Subject:" + subject + "\r\n" +
			"\r\n" +
			message)

	// 認証情報の設定
	auth := smtp.PlainAuth(
		"",
		config.SMTP.userName,
		config.SMTP.password,
		config.SMTP.smtpHost,
	)

	// メール送信
	mailError := smtp.SendMail(
		config.SMTP.smtpHost+":"+config.SMTP.smtpPort,
		auth,
		config.MailAddress.fromAddress,
		config.MailAddress.toAddress,
		mailContent,
	)

	if mailError != nil {
		fmt.Println(mailError)
		return
	}

}
