package main

//SMTP -- Simple Mail Transfer Protocol

import (
	"fmt"
	"log"
	"net/smtp"
)

//SMTP mail göndermek
//POP IMAP mail almak

func main() {

	// Gönderici Bilgileri
	gonderen := "golangahks@gmail.com"
	sifre := "Udemy444"

	// Alıcı Bilgiler
	alici := []string{
		"hamid.soyler@gmail.com",
	}

	// STMP Ayarları
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Mesaj.
	baslik := "Subject:Bu bir başlıktır\n"
	icerik := "Bu mail icerigidir"

	mesaj := []byte(baslik + icerik)

	//Authentication İşlemleri

	auth := smtp.PlainAuth("", gonderen, sifre, smtpHost)

	// Email Gönderme

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, gonderen, alici, mesaj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("E-Posta başarıyla gönderildi.")

}
