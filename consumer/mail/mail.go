package mail

import (
	"net/smtp"
)


func Send(email string)(er error) {
	
	from := ""
	password := ""
	to := email	
	msg:="Subject: Transaction Successfull\n\n" +
	 	"Your transaction has been completed.";
	
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	return err;
	

}