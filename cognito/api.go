package cognito

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

//GetIdentityFor - If this is a first connect then we need to get an identity for this new client from the configured pool.  Once we have that we should cahce that identity for later use.
func GetIdentityFor(poolID string) (string, error) {
	//TODO check cache, if not provided then fetch and then cache, or do we let the caller manage this? they could pass a func to manage cache? read/writer
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
