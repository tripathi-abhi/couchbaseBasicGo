package main

import (
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

func main() {

	// Connect golang to couchbase (couchbase capella). Will do for local server later.

	// variables required for connection
	var (
		databaseEndpoint = "cb.4eyb9sg8da5ac5od.cloud.couchbase.com"
		bucketName       = "travel-sample"
		username         = "kira"
		password         = "Kira@123"
	)

	// Initialize connection with the db

	// 1. Connect to the cluster using username and password
	travelSampleCluster, connectionError := gocb.Connect("couchbases://"+databaseEndpoint, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})

	if connectionError != nil {
		log.Fatal(connectionError)
	}

	// 2. Get the bucket and wait for the bucket object to be ready.
	bucket := travelSampleCluster.Bucket(bucketName)
	bucketReadyError := bucket.WaitUntilReady(3*time.Second, nil)
	if bucketReadyError != nil {
		log.Fatal(bucketReadyError)
	}

	col := bucket.Scope("inventory").Collection("hotel")
	fmt.Println("Collection: ",col)

}
