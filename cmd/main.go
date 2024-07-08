package main

import (
	"context"
	"log"

	miniostorage "github.com/thirteenths/go-minio-service/internal/storage/minioStorage"
)

func main() {
	endpoint := "localhost:9000"
	user := "user"
	password := "password"
	useSSL := false

	storage, err := miniostorage.NewMinioStorage(endpoint, user, password, useSSL)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// file, err := os.Open("/home/verendaya/Pictures/MANTRA.png")

	// if err != nil {
	//	log.Fatalf(err.Error())
	// }

	// imageLoad := models.ImageUnit{
	//	BucketName: "try",
	//	ObjectName: "",
	//	Path:       "/home/verendaya/Pictures/MANTRA.png",
	//	ObjectSize: -1,
	// }

	// imageName, err := storage.UploadFile(context.Background(), imageLoad)

	const name = "553994cb-571f-4109-a084-db727974c251"

	obj, err := storage.DownloadFile(context.Background(), name, "try")

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Print(obj.ObjectBytes)
}
