package cognito

import (
	"fmt"
)

//TODO should pass region down
var cognitoURL = "https://cognito-identity.eu-west-1.amazonaws.com"

//FetchNewIdentityFor - see https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html#API_GetId_Examples
func FetchNewIdentityFor(poolID string) (string, error) {
	target := "com.amazonaws.cognito.identity.model.AWSCognitoIdentityService.GetId"
	type reqT struct {
		ID string `json:"IdentityPoolId"`
	}
	type resT struct {
		ID string `json:"IdentityId"`
	}
	var res resT
	err := DoPost(cognitoURL, target, reqT{poolID}, &res)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}

// IdentityCredentials - ...
type IdentityCredentials struct {
	Key     string  `json:"AccessKeyId"`
	Expires float64 `json:"Expiration"`
	Secret  string  `json:"SecretKey"`
	Token   string  `json:"SessionToken"`
}

// FetchIdentityCredentials - https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetCredentialsForIdentity.html
func FetchCredentialsFor(identity string) (*IdentityCredentials, error) {
	target := "com.amazonaws.cognito.identity.model.AWSCognitoIdentityService.GetCredentialsForIdentity"
	type reqT struct {
		ID string `json:"IdentityId"`
	}
	type resT struct {
		Creds IdentityCredentials `json:"Credentials"`
		ID    string              `json:"IdentityId"`
	}
	var res resT
	err := DoPost(cognitoURL, target, reqT{identity}, &res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &res.Creds, nil
}
