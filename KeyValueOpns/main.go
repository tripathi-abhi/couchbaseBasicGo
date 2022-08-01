package main

import (
	"fmt"
	"kvOperations/database"
	"kvOperations/utils"
	"time"
)

type Airline struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	CallSign string `json:"callsign"`
	Country  string `json:"country"`
}

func main() {

	// All the variables needed for connection
	var (
		endpoint       = "cb.4eyb9sg8da5ac5od.cloud.couchbase.com"
		bucketName     = "travel-sample"
		scopeName      = "inventory"
		collectionName = "airline"
		username       = "kira"
		password       = "Kira@123"
	)

	fmt.Println("Connecting to cluster...")
	clusterObj := database.ConnectWithCluster(endpoint, username, password)
	fmt.Println("Connection to cluster is established successfully")

	bucketObj := clusterObj.Bucket(bucketName)
	bucketOpenErr := bucketObj.WaitUntilReady(3*time.Second, nil)
	utils.CheckErrorNil(bucketOpenErr)

	collectionObj := bucketObj.Scope(scopeName).Collection(collectionName)

	// Get data from the collection
	var airlineDetails Airline
	documentFetched, getError := collectionObj.Get("airline_10123", nil)
	utils.CheckErrorNil(getError)

	getError = documentFetched.Content(&airlineDetails)
	utils.CheckErrorNil(getError)
	fmt.Printf("AirlineDetails: %#v\n", airlineDetails)

	// insert data using INSERT (UPSERT can also be used)
	// newAirLineDetails := Airline{Id: 002, Type: "airline", Iata: "PAPPU", Icao: "TWW", CallSign: "TXW", Country: "India"}
	// data, err := collectionObj.Insert(fmt.Sprintf("airline_%d", newAirLineDetails.Id), newAirLineDetails, nil)
	// utils.CheckErrorNil(err)
	// fmt.Printf("data:%#v\n", data)

	// remove data using REMOVE
	removedData, err := collectionObj.Remove(fmt.Sprintf("airline_%d", 002), nil)
	utils.CheckErrorNil(err)
	fmt.Printf("data:%#v\n", removedData)

}
