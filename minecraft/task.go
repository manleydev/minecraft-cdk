package minecraft

import (
	"strconv"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

type TaskProps struct {
	Vpc awsec2.IVpc
}

const TASK_FAMILY_NAME = "minecraft"
const CPU_UNITS = 1024
const MEMORY_MIB = 4096
const CONTAINER_IMAGE = "hchasens/papermc:1.18.1"

func NewTask(scope awscdk.Stack, id string, props *TaskProps) awsecs.TaskDefinition {

	task := awsecs.NewTaskDefinition(scope, aws.String("Task"), &awsecs.TaskDefinitionProps{
		Family:        jsii.String(TASK_FAMILY_NAME),
		Compatibility: awsecs.Compatibility_EC2,
		Cpu:           jsii.String(strconv.Itoa(CPU_UNITS)),
		MemoryMiB:     jsii.String(strconv.Itoa(MEMORY_MIB)),
		NetworkMode:   awsecs.NetworkMode_HOST,
	})

	efs_drive := awsefs.NewFileSystem(scope, jsii.String("files"), &awsefs.FileSystemProps{
		Vpc: props.Vpc,
	})

	group := awsec2.NewSecurityGroup(efs_drive, jsii.String("efs-security-group"), &awsec2.SecurityGroupProps{
		Vpc:              props.Vpc,
		AllowAllOutbound: jsii.Bool(true),
	})
	group.AddIngressRule(awsec2.Peer_AnyIpv4(), awsec2.Port_AllTraffic(), jsii.String("Allow all traffic"), jsii.Bool(false))

	efs_drive.Connections().AddSecurityGroup(group)

	volume := &awsecs.Volume{
		Name: aws.String("minecraft-data"),
		EfsVolumeConfiguration: &awsecs.EfsVolumeConfiguration{
			FileSystemId:  efs_drive.FileSystemId(),
			RootDirectory: jsii.String("/"),
		},
	}

	task.AddVolume(volume)

	container := awsecs.NewContainerDefinition(scope, aws.String("minecraft"), &awsecs.ContainerDefinitionProps{
		Image:                awsecs.ContainerImage_FromRegistry(aws.String(CONTAINER_IMAGE), &awsecs.RepositoryImageProps{}),
		TaskDefinition:       task,
		MemoryReservationMiB: jsii.Number(MEMORY_MIB),
		Cpu:                  jsii.Number(CPU_UNITS),
		PortMappings: &[]*awsecs.PortMapping{
			{
				HostPort:      jsii.Number(25565),
				ContainerPort: jsii.Number(25565),
				Protocol:      awsecs.Protocol_TCP,
			},
		},
	})

	container.AddEnvironment(jsii.String("MC_RAM"), jsii.String(strconv.Itoa((MEMORY_MIB/1024)-1)+"G"))

	container.AddMountPoints(&awsecs.MountPoint{
		ContainerPath: jsii.String("/papermc/"),
		SourceVolume:  volume.Name,
		ReadOnly:      jsii.Bool(false),
	})

	return task
}
