package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var cae CAECollection
	endpoint := "192.168.1.156:9000"
	accessKeyID := "caeadmin"
	secretAccessKey := "caeadmin"
	useSSL := false // 没有安装证书的填false

	// Initialize minio client object. 建立连接
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client1, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client1.Database("datebase").Collection("collection1")
	cae.minioC = minioClient
	cae.CollectionObj = collection

	cae.GetData("szsc006", "YL-test")
	// str := "trailId=002&fileName=Test01&status=false"
	// fmt.Println("上传文件信息: " + str)
	// http.Get("http://localhost:8080/websocket?" + str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
