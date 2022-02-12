package minecraft

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/jsii-runtime-go"
)

type ServiceProps struct {
	Cluster awsecs.Cluster
	Vpc     awsec2.IVpc
	Task    awsecs.TaskDefinition
}

func NewService(scope awscdk.Stack, id string, props *ServiceProps) awsecs.IService {

	service := awsecs.NewEc2Service(scope, jsii.String("service"), &awsecs.Ec2ServiceProps{
		Cluster:        props.Cluster,
		DesiredCount:   jsii.Number(1),
		TaskDefinition: props.Task,
	})

	// create loadbalancer to ecs service
	//service.AttachToNetworkTargetGroup()

	return service
}
