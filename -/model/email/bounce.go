package email

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	config "github.com/kevinrizkhy/Internal/config"
	database "github.com/kevinrizkhy/Internal/db"
	"log"
	"strings"
)

func CheckBounce() {
	//log.Println("Connecting to server...")
	c, err := client.DialTLS(config.BOUNCE_IMAP+config.BOUNCE_PORT, nil)
	if err != nil {
		log.Println("Not connected...")
		log.Println(err)
	}
	//log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(config.BOUNCE_EMAIL, config.BOUNCE_PASSWORD); err != nil {
		log.Println(err)
	} else {
		mbox, err := c.Select("INBOX", false)
		if err != nil {
			log.Println("No folder INBOX")
		}
		// Get the last message
		if mbox.Messages == 0 {
			//log.Println("No message in mailbox")
		}
		seqset := new(imap.SeqSet)
		seqset.AddRange(1, mbox.Messages)
		// Get the whole message body
		attrs := []string{"BODY[]"}
		messages := make(chan *imap.Message, 1)
		go func() {
			if err := c.Fetch(seqset, attrs, messages); err != nil {
				log.Println(err)
			}
		}()
		count := 0

		for a := range messages {
			message := a.Format()[3]
			messageText := fmt.Sprint(message)
			if strings.Contains(messageText, "!@#$%^&*()_+") {
				id := strings.Split(strings.Split(messageText, "!@#$%^&*()_+")[1], "!@#$%^&*()_+")[0]
				updateCampaignEmail(id)
				fmt.Println("ID ke - ", count+1, " : ", id)
				//s := new(imap.SeqSet)
				//s.AddNum(a.SeqNum)
				//operation := "+FLAGS.SILENT"
				//flags := []interface{}{imap.DeletedFlag}

				// if err := c.Store(s, operation, flags, nil); err != nil {
				// 	log.Println(err)
				// }
				// fmt.Println("Delete")
				// if err := c.Expunge(nil); err != nil {
				// 	log.Println(err)
				// }
				// fmt.Println("Delete OKE")
				count++
			}
		}
		fmt.Println(count)
		c.Logout()
	}
}

func updateCampaignEmail(ids string) {
	rows, err := database.ExecuteQuery("SELECT reference FROM cron_email WHERE id=$1", ids)
	if len(rows) > 0 && err == nil {
		id := rows[0]["reference"]
		_, err := database.ExecuteQuery("update campaign_email set bounce=true where id=$1", id)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
