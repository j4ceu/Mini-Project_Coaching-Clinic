package helper

import (
	"Mini-Project_Coaching-Clinic/models"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

type BodyLinkEmail struct {
	URL       string
	Firstname string
	Invoice   *os.File
}

func GenerateInvoice(userPayment models.UserPayment, invoice string) string {
	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: "INVOICE",
		GreyBgColor:     []int{151, 170, 183},
		DarkBgColor:     []int{244, 226, 219},
		CurrencySymbol:  "Rp.",
	})

	doc.SetDescription("Coaching Clinic - " + userPayment.User.FirstName)
	doc.SetRef(invoice)
	doc.SetDate(time.Now().Format("02 January 2006"))

	//Check Work Directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println(path)

	logoBytes, err := ioutil.ReadFile(path + "/images/company-logo.png")
	if err != nil {
		log.Fatal("Read file error")
		log.Fatal(err)
	}

	doc.SetCompany(&generator.Contact{
		Name: "Coaching Clinic",
		Address: &generator.Address{
			Address:    "Jl. Manggaraya B32",
			City:       "Bekasi",
			PostalCode: "17121",
			Country:    "Indonesia",
		},
		Logo: &logoBytes,
	})

	doc.SetCustomer(&generator.Contact{
		Name: userPayment.User.FirstName + " " + userPayment.User.LastName,
		Address: &generator.Address{
			Address: userPayment.User.Email,
		},
	})

	for _, userBook := range userPayment.UserBook {
		doc.AppendItem(&generator.Item{
			Name:     userBook.Title,
			Quantity: "1",
			UnitCost: strconv.Itoa(userBook.CoachAvailability.Coach.Price),
		})
	}

	doc.SetDefaultTax(&generator.Tax{
		Percent: "5",
	})

	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
		log.Fatal("error doc build")
	}

	var buf bytes.Buffer

	err = pdf.Output(&buf)
	if err != nil {
		log.Fatal(err)
		log.Fatal("error pdf output")
	}

	file := createTempFile(buf, invoice)

	url := UploadFileToFirebase(buf, invoice+".pdf")

	bodyLinkEmail := BodyLinkEmail{
		URL:       url,
		Firstname: userPayment.User.FirstName,
		Invoice:   file,
	}

	go SendEmailVerification(userPayment.User.Email, bodyLinkEmail)

	return url

}

func createTempFile(buf bytes.Buffer, invoice string) *os.File {
	invoiceName := invoice + ".pdf"
	f, err := os.Create(invoiceName)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Create file error")
	}

	defer f.Close()

	if _, err := f.Write(buf.Bytes()); err != nil {
		log.Fatal(err)
		log.Fatal("Write file error")
	}

	return f
}
