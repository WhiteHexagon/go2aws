package cognito

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

var cognitoURL string

//SetRegionURL - We need to target the correct AWS Cognito region.  example:  "https://cognito-identity.eu-west-1.amazonaws.com"
func SetRegionURL(url string) {
	cognitoURL = url
}

//GetIdentityFor - If this is a new client, then we need to get a new identity from the configured pool.  Once we have that we should cache that identity for later use.
func GetIdentityFor(poolID string) (string, error) {
	//TODO check local cache, if not provided then fetch and then cache, or do we let the caller manage this? they could pass a func to manage cache? read/writer
	return FetchNewIdentityFor(poolID)
}

// GetCredentialsFor - Create a provider, and trigger the provider to fetch the credentials
func GetCredentialsFor(identity string) (*credentials.Credentials, error) {
	provider := CustomCognitoCredentialsProvider{identity}
	creds := credentials.NewCredentials(&provider)
	//only now do we actually force the retrieval of the credentials
	if _, err := creds.Get(); err != nil {
		return nil, errors.New("failed to get credentials: " + err.Error())
	}
	return creds, nil
}
