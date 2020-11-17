package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func main() {

	// 標準出力テストコード
	fmt.Print("Hello World")

	// SMTP設定 TODO:外部ファイル化・構造体化
	userName := "ユーザー名@gmail.com"
	password := "パスワード" // gmailの場合は2段階認証設定+2段階設定後にアプリパスワードで設定したパスワード
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// 送信元・送信先メールアドレス
	fromAddress := userName
	toAddress := []string{"送信先メールアドレス１", "送信先メールアドレス２"}

	// メール内容 TODO:text/templateに変更
	subject := "タイトル"
	message := "テストメール"
	mailContent := []byte(
		"To: " + strings.Join(toAddress, ";") + "\r\n" +
			"Subject:" + subject + "\r\n" +
			"\r\n" +
			message)

	// 認証情報の設定
	auth := smtp.PlainAuth(
		"",
		userName,
		password,
		smtpHost,
	)

	// メール送信
	error := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		fromAddress,
		toAddress,
		mailContent,
	)

	if error != nil {
		log.Fatal(error)
	}

}
