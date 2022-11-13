package helper

import (
	"Mini-Project_Coaching-Clinic/constant"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

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
	fmt.Println("Init firebase app")
	privateKey := strings.Replace(os.Getenv("PRIVATE_KEY"), `\n`, "\n", -1)

	serviceAccountKey := serviceAccountKey{
		Type:                    os.Getenv("TYPE"),
		ProjectID:               os.Getenv("PROJECT_ID"),
		PrivateKeyID:            os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey:              string(privateKey),
		ClientEmail:             os.Getenv("CLIENT_EMAIL"),
		ClientID:                os.Getenv("CLIENT_ID"),
		AuthURI:                 constant.AuthURI,
		TokenURI:                constant.TokenURI,
		AuthProviderX509CertURL: constant.AuthProviderX509CertURL,
		ClientX509CertURL:       constant.ClientX509CertURL,
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
