package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	gonderen := "golangahks@gmail.com"
	sifre := "Udemy444"

	alici := []string{
		"hamid.soyler@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//MIME -- Multipurpose Internet Mail Extensions

	baslik := "Subject:Merhaba Udemy\n"
	mime := "MIME-version:1.0\nContent-Type:text/html; charset=\"UTF-8\";\n\n"
	icerik := "<html><body><h1 style=\"color:blue;\">Merhaba Bu bir HTML göstergesidir.</h1></body></html>"

	mesaj := []byte(baslik + mime + icerik)

	auth := smtp.PlainAuth("", gonderen, sifre, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, gonderen, alici, mesaj)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mail gönderildi")

}
