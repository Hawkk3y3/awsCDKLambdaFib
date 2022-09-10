package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	awscdkapigw "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	awsapigwintegrations "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaStackProps struct {
	awscdk.StackProps
}

func NewLambdaStack(scope constructs.Construct, id string, props *LambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	fibLambdaFunc := awscdklambdago.NewGoFunction(stack, jsii.String("FibonacciFunc"), &awscdklambdago.GoFunctionProps{
		FunctionName: jsii.String("FibonacciFunc"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Description:  jsii.String("an API GW handler that returns Fibonacci Number at Place N as JSON"),
		Entry:        jsii.String("../lambdaApp"),
	})

	fibApi := awscdkapigw.NewHttpApi(stack, jsii.String("FibonacciAPI"), nil)

	fibApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path:        jsii.String("/"),
		Methods:     &[]awscdkapigw.HttpMethod{awscdkapigw.HttpMethod_GET},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(jsii.String("FibonacciApiIntegration"), fibLambdaFunc, nil),
	})
	awscdk.NewCfnOutput(stack, jsii.String("FibonacciApiURL"), &awscdk.CfnOutputProps{
		Value:       fibApi.ApiEndpoint(),
		Description: jsii.String("the URL to the Fibonacci API"),
		ExportName:  jsii.String("FibonacciApiURL"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaStack(app, "LambdaStack", &LambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
