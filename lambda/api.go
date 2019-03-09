package lambda

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var configured bool
var lambdaService *lambda.Lambda

//InitFor - Identitify the AWS region where our Lambdas are stored, and provide credentials for further lambda calls.
//TODO: we ignore credential expiration currently and leave it for the caller to renew them.  This avoids an inter-package dependency.
func InitFor(region string, creds *credentials.Credentials) error {
	session, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return errors.New("InitFor:: failed because: " + err.Error())
	}
	lambdaService = lambda.New(session, &aws.Config{Region: &region, Credentials: creds})
	configured = true
	return nil
}

//Call - assumes simple lambda that takes and respons with a JSON message
func Call(name string, request interface{}, result interface{}) error {
	if !configured {
		return errors.New("not configured, you must call InitFor first")
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return errors.New("failed to marshall request: " + err.Error())
	}

	resp, err := lambdaService.Invoke(&lambda.InvokeInput{FunctionName: &name, Payload: payload})
	if err != nil {
		return errors.New("failed to call lambda: " + name + "Reason: " + err.Error() + " Payload " + string(payload))
	}

	if *resp.StatusCode != 200 {
		return errors.New("unknown status: " + string(*resp.StatusCode))
	}

	err = json.Unmarshal(resp.Payload, &result)
	if err != nil {
		return errors.New("failed to unmarshall response: " + string(resp.Payload) + " Reason: " + err.Error())
	}
	return nil
}
