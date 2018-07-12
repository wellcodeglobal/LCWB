package utility

import (
	"fmt"
	database "github.com/wellcode/LCWB/-/model/db"
	"time"
)

func AddNotification(sender, receiver, message, link string) error {
	token := GenerateToken(sender + receiver + message + link + fmt.Sprint(time.Now()))
	link = link + "?t=" + token + "#bottompage"
	_, err := database.ExecuteQuery("insert into notification (sender,receiver,content, link, token) values ($1,$2,$3,$4,$5);", sender, receiver, message, link, token)
	if err != nil {
		fmt.Println("ERR ADD-NOTIFICATION")
		fmt.Println(sender+" - "+receiver+" - "+message+" - "+link, err)
		return err
	} else {
		return nil
	}
}

func AddCronEmail(receiver, subject, message, types, reference string) error {
	_, err := database.ExecuteQuery("INSERT INTO cron_email(target,subject,message,type,reference) VALUES($1,$2,$3,$4,$5)", receiver, subject, message, types, reference)
	if err != nil {
		fmt.Println("ERR ADD-EMAIL")
		fmt.Println(receiver + " - " + subject + " - " + reference + " - " + types)
		return err
	} else {
		return nil
	}
}
