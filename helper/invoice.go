package helper

import (
	"Mini-Project_Coaching-Clinic/models"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

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

	logoBytes, _ := ioutil.ReadFile("./images/company-logo.png")

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
	}

	// err = pdf.OutputFileAndClose("./upload/invoices/" + invoice + ".pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var buf bytes.Buffer

	err = pdf.Output(&buf)
	if err != nil {
		log.Fatal(err)
	}

	url := UploadFileToFirebase(buf, invoice+".pdf")
	return url

}
