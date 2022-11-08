package helper

import (
	"bytes"
	"context"
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

func InitAppFirebase() {
	config := &firebase.Config{
		StorageBucket: "coaching-clinic.appspot.com",
	}

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
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

	url, _ := url.Parse(fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", wc.Attrs().Name, fileName, wc.Metadata["firebaseStorageDownloadTokens"]))
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
