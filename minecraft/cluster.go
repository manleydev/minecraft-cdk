package minecraft

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/jsii-runtime-go"
)

type ClusterProps struct {
	Vpc awsec2.IVpc
}

func NewCluster(scope awscdk.Stack, id string, props *ClusterProps) awsecs.Cluster {

	cluster := awsecs.NewCluster(scope, jsii.String("cluster"), &awsecs.ClusterProps{
		Vpc:         props.Vpc,
		ClusterName: jsii.String("minecraft"),
	})

	cluster.AddCapacity(jsii.String("default"), &awsecs.AddCapacityOptions{
		DesiredCapacity: jsii.Number(1),
		MaxCapacity:     jsii.Number(1),
		InstanceType:    awsec2.NewInstanceType(jsii.String("t3a.large")),
	})

	return cluster
}
