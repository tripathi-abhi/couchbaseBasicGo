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
		// databaseEndpoint = "cb.localhost.svc"
		bucketName = "travel-sample"
		username   = "Kira"
		password   = "Kira@123"
	)

	// gocb.SetLogger(gocb.VerboseStdioLogger())

	// Initialize connection with the db

	// 1. Connect to the cluster using username and password.
	/*
		 	* 1. The url can be couchbases://endpoint if tls connection is to be made
		 	* 2. Else couchbase://endpoint should be used
			* 3. "couchbase://locahost" is equivalent to just "localhost"
	*/
	travelSampleCluster, connectionError := gocb.Connect("couchbase://localhost", gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})

	if connectionError != nil {
		log.Fatal(connectionError.Error())
	}
	fmt.Println("Connected with cluster: ", travelSampleCluster)

	// 2. Get the bucket and wait for the bucket object to be ready.
	bucket := travelSampleCluster.Bucket(bucketName)
	bucketReadyError := bucket.WaitUntilReady(3*time.Second, nil)
	if bucketReadyError != nil {
		log.Fatal(bucketReadyError)
	}

	col := bucket.Scope("inventory").Collection("hotel")
	fmt.Println("Collection: ", col)

}
