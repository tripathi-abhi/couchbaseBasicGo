package main

import (
	"fmt"
	"queryexamples/database"
	"queryexamples/utils"
	"time"

	"github.com/couchbase/gocb/v2"
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

	var (
		endpoint   = "cb.4eyb9sg8da5ac5od.cloud.couchbase.com"
		bucketName = "travel-sample"
		// scopeName  = "inventory"
		// collectionName = "airline"
		username       = "kira"
		password       = "Kira@123"
		airlineDetails []Airline
	)

	fmt.Println("Connecting to the cluster")
	clusterObj := database.ConnectWithCluster(endpoint, username, password)

	bucketObj := clusterObj.Bucket(bucketName)
	err := bucketObj.WaitUntilReady(3*time.Second, nil)
	utils.CheckErrorNil(err)

	// scopeObj := bucketObj.Scope(scopeName)

	// query using named parameters
	query := "SELECT airline.* from `travel-sample`.inventory.airline where country = $countryname AND id=$id"
	params := make(map[string]interface{})
	params["countryname"] = "United States"
	// params["id"] = 10765

	queryResults, err := clusterObj.Query(query, &gocb.QueryOptions{
		// Adhoc:           true,
		NamedParameters: params,
	})

	utils.CheckErrorNil(err)
	fmt.Println("result: ", queryResults)
	for queryResults.Next() {
		var airlineDetail Airline
		err := queryResults.Row(&airlineDetail)
		if err != nil {
			panic(err)
		}
		airlineDetails = append(airlineDetails, airlineDetail)
		fmt.Println(airlineDetail)
	}
	fmt.Println("No of records: ", len(airlineDetails))
}
