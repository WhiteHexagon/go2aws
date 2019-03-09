package cognito

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

var gotValidCreds bool
var cachedCredentials credentials.Value
var expires time.Time

//CustomCognitoCredentialsProvider a small provider that supports Cognito identity pools
type CustomCognitoCredentialsProvider struct {
	identity string
}

// Retrieve - ...
func (m *CustomCognitoCredentialsProvider) Retrieve() (credentials.Value, error) {
	if !gotValidCreds {
		fmt.Println("cache miss, so fetch")
		creds, err := FetchCredentialsFor(m.identity)
		if err != nil {
			return credentials.Value{}, errors.New("fetchIdentityCredentials failed: " + err.Error())
		}
		cachedCredentials = credentials.Value{}
		cachedCredentials.AccessKeyID = creds.Key
		cachedCredentials.ProviderName = "CustomCognitoCredentialsProvider"
		cachedCredentials.SecretAccessKey = creds.Secret
		cachedCredentials.SessionToken = creds.Token

		//store expiry
		ts := time.Millisecond * time.Duration(creds.Expires)
		expires = time.Now().Add(ts)
		gotValidCreds = true
		fmt.Println("Expires: " + expires.Format(time.RFC3339))
	}
	fmt.Println("cache hit")
	return cachedCredentials, nil
}

// IsExpired - ...
func (m *CustomCognitoCredentialsProvider) IsExpired() bool {
	if gotValidCreds {
		now := time.Now()
		return now.After(expires)
	}
	return true
}
