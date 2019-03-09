package main

import (
	"fmt"

	"github.com/whitehexagon/go2aws/cognito"
	"github.com/whitehexagon/go2aws/lambda"
)

//Request - ...
type Request struct {
	Question string `json:"Meaning"`
}

//Response - ...
type Response struct {
	Answer string `json:"Answer"`
}

var region = "eu-west-1"                                      // <- your region
var poolID = "eu-west-1:99999999-999999999999999999999999999" // <- add your pool id here

func sampleWorkflow() {

	// register the local AWS Cognito service region we are using
	cognito.SetRegionURL("https://cognito-identity.eu-west-1.amazonaws.com") // <- your region

	// get a new identity - NOTE: should be cached for future use
	identity, err := cognito.GetIdentityFor(poolID)
	if err != nil {
		fmt.Println("GetIdentityFor:: failed because: " + err.Error())
		return
	}

	// now get 'unauthenticated' credentials for this identity
	creds, err := cognito.GetCredentialsFor(identity)
	if err != nil {
		fmt.Println("GetCredentialsFor:: failed because: " + err.Error())
		return
	}

	// register the local AWSLambda region we are using
	if err := lambda.InitFor(region, creds); err != nil {
		fmt.Println("InitFor:: failed because: " + err.Error())
		return
	}

	// now call our lambda
	result := Response{}
	err = lambda.Call("lamda-function-name", Request{"Life?"}, &result) // <- add your lambda function name here
	if err != nil {
		fmt.Println("Call:: failed because: " + err.Error())
		return
	}
	fmt.Printf("Result: %+v\nDone\n", result.Answer)
}

func main() {
	sampleWorkflow()
}
