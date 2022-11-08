package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func sendMail(to string, data interface{}, templateFile string) error {
	// Set up authentication information.
	result, _ := ParseTemplate(templateFile, data)
	m := gomail.NewMessage()
	m.SetHeader("From", "j4ceucoachingclinic@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("Invoice Coaching Clinic - %s", data.(BodyLinkEmail).Firstname))
	m.SetBody("text/html", result)
	m.Attach(data.(BodyLinkEmail).Invoice.Name())

	d := gomail.NewDialer("smtp.gmail.com", 587, "j4ceucoachingclinic@gmail.com", "fssqilhrgqwrdzhl")
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	os.Remove(data.(BodyLinkEmail).Invoice.Name())
	return nil

}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmailVerification(to string, data interface{}) {
	var err error
	template := "html/invoice_email.html"
	err = sendMail(to, data, template)
	if err == nil {
		fmt.Println("send email success")
	} else {
		fmt.Println(err)
	}
}
