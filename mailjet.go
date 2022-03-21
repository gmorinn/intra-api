package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mailjet/mailjet-apiv3-go"
)

func SendEmail(msg string) error {
	var firstname string = strings.ToUpper(strings.Split(os.Getenv("EMAIL"), ".")[0])
	m := mailjet.NewMailjetClient("c76fcb2ae5762249d3747e08fe8f11ee", "a47b434df1700a5469a7fb549b5f81b9")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "guillaume.morin974@gmail.com",
				Name:  "GM API",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: os.Getenv("EMAIL"),
					Name:  firstname + "!",
				},
			},
			TemplateID:       3749194,
			TemplateLanguage: true,
			Subject:          fmt.Sprintf("Inscrit toi bordel %s", firstname),
			Variables:        map[string]interface{}{"msg": msg, "firstname": firstname},
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := m.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
