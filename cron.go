package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
)

func format(title, content string) string {
	var res string = fmt.Sprintf("<h3>Titre:</h3> %s<br />", title)
	res += fmt.Sprintf("<h4>Contenu:</h4> %s<br /><br />", title)
	res += "********************************<br />"
	return res
}

func runCron(c **cron.Cron) {
	fmt.Println("start !")
	(*c) = cron.New()

	(*c).AddFunc("@every 5s", func() {
		fmt.Println("Send !")
		var resultRequest []*NofifResponse = getNotif()
		var message string
		for _, v := range resultRequest {
			message += format(v.Title, v.Content)
		}
		fmt.Println(message)
		if err := SendEmail(message); err != nil {
			fmt.Println(err)
		}
	})

	(*c).Start()
	log.Printf("%+v\n", (*c).Entries())
}
