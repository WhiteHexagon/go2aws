module github.com/whitehexagon/go2aws/login

require (
	github.com/aws/aws-sdk-go v1.16.35
	github.com/whitehexagon/go2aws/cognito v0.0.0
	github.com/whitehexagon/go2aws/lambda v0.0.0
)

replace github.com/whitehexagon/go2aws/cognito v0.0.0 => ../cognito

replace github.com/whitehexagon/go2aws/lambda v0.0.0 => ../lambda
