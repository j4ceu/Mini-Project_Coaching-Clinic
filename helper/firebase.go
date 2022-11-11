package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/url"

	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

var App *firebase.App

type serviceAccountKey struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func InitAppFirebase() {
	serviceAccountKey := serviceAccountKey{
		Type:                    "service_account",
		ProjectID:               "coaching-clinic",
		PrivateKeyID:            "fe58e7329298bb92a15a5fe0098166cd3c91458f",
		PrivateKey:              "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDRLkN8pJ2G4OwS\npebkHlRwelBSW+uo8rVpAB/M6wmCzBNMOAIoSTwKObP0zdnX5H8mmBH4FWE1pQEV\nbp3Iz+o0ZIixam1Zlw+NOu2NNrVdDF0ZRxufWv9RooYKReAo2wAKaO4vRzM4bvFU\nOdoH2Rd/IuQde3KwNoTNkEHzxeSmiJPON6pGwAHJ18mj9scf910/wn2YX2QHo5TU\nd/pPJF5i2C5KmgXIGgVaI/OLqTU4+fpncM3bKt1IoKSVOaikIyDMkk2BweB8onhI\nV2HDkR3xGxcmMYxVTnk10BLIQkLZJZHLCaat8whHtwNIHYIzej0zrR2+FCtohmP3\nqxSZSLvzAgMBAAECggEAIbzhTqJVnpJQERE9TwVN9Nbiyz3JzGCE/i/ycRW8Ms02\naZXOUfAUj8j5wN34qOJQ2F6GO2nmBAXJHyCjBedJgh9wOpPh73q2m+k+tNwiM/qR\nWMkyvSjV7YqV2Dn9QFRTdS+yXDa9c03VfOetxj6Z6bwCEyKSKgtJBRyjaS1oRlbi\n5G1mzkRJBQsS5xV7UXErodITbDZrQn6s6JSUR9nZkDlhCq4JtQdUuDhVCacdFERy\nkDbaJ87QbgYCxDtjdK/DfQgBwnJSF0TwrZr0GcUnYnWp3i+BS6LizQX3Pp/B5oCR\n+CVFYbvBkQc51UcjxTIuipZ6D1UG1XYJiBW8QNz4mQKBgQD/VUXtAaMI6OqdhYA1\nPapO+4EFLLp6WYNEctWkp9LUyYhKU1hlXOmPH1kDzpaWqXrlhA42fURCM5O/Cei/\nliL1UJUyyH7MZ8koFVWiycS+KElfs+IoZUYsBItu2TbMvHp/0FA2ITjRDNc+GOz8\nqGrJAWEFKLF+5uTY7oJq1uTMGwKBgQDRuiGHtJE6pEV/J4gwBkVLtfF6E5httGR5\nEkS1/q6GWQH4n5Gzp7dJogqjeyUISrvSu92ZZ8huXI4AbwMkbhe9J7XZSOdSzX6e\nN0lQAKBBW8R5ENfo7/3gcVp+43KjYiEdZsvbjLQxrlLTlBZFYO+5o2NPfSTg+jR2\n5xec3UOdCQKBgAhWZH4ku4oi8OZL4a/bX6BMnh3pI/2yxpKJnWhPApdoLUcgCZbl\nvcqqn2F8cXZh+l1cPoqQ9JWk0YI/dJYs9N9FzllmKp3KFct1RuKV7BK5hgvV9+CR\nzgTJ8TOhbCSrUuSxjKz30L8iyDSC49osNbBylxRwC7u1Fmvu/ds3QSlZAoGAIivH\nhsre1sUpJZyVTe3XoIxWeeNzdHxt2mQlmdmTKgSak528KZ9r961VOmm4EO/MRnuh\nkXsdZw3hfSSZSHg/mew8bti4B/+/X5v2b/iKI9wF2QvrgeKIZOdTLVV1ujUk3BuB\nn5X/ThDvIdYBAvDWXlLGvY7QUW+y2KSybjaG67ECgYA7rI/2QNZn2VPv+uBDl0vH\nvStbnaPMQwoKf4JVQ27mPYvysKrFHREUD9VVTctfFnpVs1EgjiYsqLVsF6n0E8l9\n5gOWR9M8W+E9ccsBLX7lJ9sQ5T2tOUGkpT0p9WqfGImgbdzqnK8TCptAkkk/pc2Q\n+w6f53rFgjCg3kKjbdkJpg==\n-----END PRIVATE KEY-----\n",
		ClientEmail:             "firebase-adminsdk-b6rek@coaching-clinic.iam.gserviceaccount.com",
		ClientID:                "102529450614053381959",
		AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
		TokenURI:                "https://oauth2.googleapis.com/token",
		AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
		ClientX509CertURL:       "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-b6rek%40coaching-clinic.iam.gserviceaccount.com",
	}

	// Struct to JSON
	jsonKey, _ := json.Marshal(serviceAccountKey)

	config := &firebase.Config{
		StorageBucket: "coaching-clinic.appspot.com",
	}

	opt := option.WithCredentialsJSON(jsonKey)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	App = app

}

func UploadFileToFirebase(buf bytes.Buffer, fileName string) string {
	fmt.Println("Uploading file to firebase")
	client, err := App.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	wc := bucket.Object(fileName).NewWriter(context.Background())
	wc.ChunkSize = 0
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid.New().String(),
	}

	if _, err = io.Copy(wc, &buf); err != nil {
		fmt.Println(err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(wc, "%v uploaded to %v.\n", fileName, bucket)

	url, _ := url.Parse(fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", wc.Attrs().Bucket, fileName, wc.Metadata["firebaseStorageDownloadTokens"]))
	return url.String()

}

func OpenFileFromMultipartForm(file *multipart.FileHeader) (string, *bytes.Buffer, error) {
	fileUpload, err := file.Open()

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return file.Filename, &bytes.Buffer{}, err
	}

	defer fileUpload.Close()

	byteContainer, err := ioutil.ReadAll(fileUpload)
	if err != nil {
		return file.Filename, &bytes.Buffer{}, err
	}

	buf := bytes.NewBuffer(byteContainer)

	return file.Filename, buf, nil
}
