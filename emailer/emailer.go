package emailer

import (
	"gopkg.in/gomail.v2"
)

func SendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "from@gmail.com")
	m.SetHeader("To", "to@gmail.com")
	m.SetHeader("Subject", "Latest Tech News")
	m.SetBody("text/html", "<p>Hello!</p><br/><p>Here is the latest Tech News for ya!!</p><br/><br/><p>Enjoy,</p><p>Name :)</p>")
	m.Attach("tectnews.csv")

	d := gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "from pwd")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
