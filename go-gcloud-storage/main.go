package main

import (
	"context"
	"fmt"
	"go-gcloud-storage/env"
	"go-gcloud-storage/lib"
	"go-gcloud-storage/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//UploadImageToBucket()
	GetImageLink()
}

func UploadImageToBucket() {
	image, err := ioutil.ReadFile("assets/image.jpg")
	if err != nil {
		log.Panicln("unable to read image: ", err.Error())
	}

	filename := os.Getenv(env.DEFAULT_IMAGE)

	gcs := lib.New()
	defer gcs.Close()

	err = gcs.Create(context.Background(), filename, image)
	if err != nil {
		log.Panicln("unable to upload image to bucket: ", err.Error())
	}
	log.Printf("uploaded %s successfully!\n", filename)
}

func GetImageFromBucket() {
	gcs := lib.New()
	defer gcs.Close()

	filename := os.Getenv(env.DEFAULT_IMAGE)

	image, err := gcs.ReadFile(context.Background(), filename)
	if err != nil {
		log.Panicln("unable to get image from bucket: ", err.Error())
	}

	imagePath := filepath.Join("assets/", filename)

	file, err := os.Create(imagePath)
	if err != nil {
		log.Panicln("unable to create image on disk local: ", err.Error())
	}
	defer utils.HandleClose(file)
	_, err = file.Write(image)
	if err != nil {
		log.Panicln("unable to close image file: ", err.Error())
	}
	log.Printf("image %s created successfully!\n", file.Name())
}

func GetImageLink() {
	gcs := lib.New()
	defer gcs.Close()

	filename := os.Getenv(env.DEFAULT_IMAGE)

	link, err := gcs.ShareFileWithTimeout(context.Background(), filename, time.Minute*3)
	if err != nil {
		log.Panicln("unable to get link from image: ", err.Error())
	}

	fmt.Println(link)
}
