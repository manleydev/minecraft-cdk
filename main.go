package main

import (
	"github.com/josephbmanley/minecraft-cdk/minecraft"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MinecraftStackProps struct {
	awscdk.StackProps
	VpcID *string
}

func NewMinecraftStack(scope constructs.Construct, id string, props *MinecraftStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := awsec2.Vpc_FromLookup(stack, aws.String("VPC"), &awsec2.VpcLookupOptions{
		VpcId: props.VpcID,
	})

	minecraft.NewTask(stack, "MinecraftTask", &minecraft.TaskProps{
		Vpc: vpc,
	})

	minecraft.NewCluster(stack, "MinecraftCluster", &minecraft.ClusterProps{
		Vpc: vpc,
	})

	// minecraft.NewService(stack, "MinecraftService", &minecraft.ServiceProps{
	// 	Vpc:     vpc,
	// 	Cluster: cluster,
	// 	Task:    task,
	// })

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewMinecraftStack(app, "MinecraftStack", &MinecraftStackProps{
		awscdk.StackProps{
			Env: env(),
		},
		jsii.String("VPC ID HERE"),
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
