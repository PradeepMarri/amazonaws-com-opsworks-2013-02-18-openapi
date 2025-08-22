package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SelfUserProfile represents the SelfUserProfile schema from the OpenAPI specification
type SelfUserProfile struct {
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Sshpublickey interface{} `json:"SshPublicKey,omitempty"`
	Sshusername interface{} `json:"SshUsername,omitempty"`
}

// UpdateLayerRequest represents the UpdateLayerRequest schema from the OpenAPI specification
type UpdateLayerRequest struct {
	Autoassignelasticips interface{} `json:"AutoAssignElasticIps,omitempty"`
	Lifecycleeventconfiguration interface{} `json:"LifecycleEventConfiguration,omitempty"`
	Customsecuritygroupids interface{} `json:"CustomSecurityGroupIds,omitempty"`
	Enableautohealing interface{} `json:"EnableAutoHealing,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Customrecipes interface{} `json:"CustomRecipes,omitempty"`
	Packages interface{} `json:"Packages,omitempty"`
	Volumeconfigurations interface{} `json:"VolumeConfigurations,omitempty"`
	Useebsoptimizedinstances interface{} `json:"UseEbsOptimizedInstances,omitempty"`
	Autoassignpublicips interface{} `json:"AutoAssignPublicIps,omitempty"`
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Custominstanceprofilearn interface{} `json:"CustomInstanceProfileArn,omitempty"`
	Shortname interface{} `json:"Shortname,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Cloudwatchlogsconfiguration interface{} `json:"CloudWatchLogsConfiguration,omitempty"`
	Layerid interface{} `json:"LayerId"`
}

// DeregisterElasticIpRequest represents the DeregisterElasticIpRequest schema from the OpenAPI specification
type DeregisterElasticIpRequest struct {
	Elasticip interface{} `json:"ElasticIp"`
}

// UpdateAppRequest represents the UpdateAppRequest schema from the OpenAPI specification
type UpdateAppRequest struct {
	Datasources interface{} `json:"DataSources,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Domains interface{} `json:"Domains,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Appid interface{} `json:"AppId"`
	Enablessl interface{} `json:"EnableSsl,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Appsource interface{} `json:"AppSource,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Sslconfiguration interface{} `json:"SslConfiguration,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// DescribeMyUserProfileResult represents the DescribeMyUserProfileResult schema from the OpenAPI specification
type DescribeMyUserProfileResult struct {
	Userprofile interface{} `json:"UserProfile,omitempty"`
}

// DescribeTimeBasedAutoScalingRequest represents the DescribeTimeBasedAutoScalingRequest schema from the OpenAPI specification
type DescribeTimeBasedAutoScalingRequest struct {
	Instanceids interface{} `json:"InstanceIds"`
}

// DeleteStackRequest represents the DeleteStackRequest schema from the OpenAPI specification
type DeleteStackRequest struct {
	Stackid interface{} `json:"StackId"`
}

// RdsDbInstance represents the RdsDbInstance schema from the OpenAPI specification
type RdsDbInstance struct {
	Address interface{} `json:"Address,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Dbinstanceidentifier interface{} `json:"DbInstanceIdentifier,omitempty"`
	Dbpassword interface{} `json:"DbPassword,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Dbuser interface{} `json:"DbUser,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Missingonrds interface{} `json:"MissingOnRds,omitempty"`
	Rdsdbinstancearn interface{} `json:"RdsDbInstanceArn,omitempty"`
}

// DescribeRdsDbInstancesRequest represents the DescribeRdsDbInstancesRequest schema from the OpenAPI specification
type DescribeRdsDbInstancesRequest struct {
	Rdsdbinstancearns interface{} `json:"RdsDbInstanceArns,omitempty"`
	Stackid interface{} `json:"StackId"`
}

// CreateDeploymentRequest represents the CreateDeploymentRequest schema from the OpenAPI specification
type CreateDeploymentRequest struct {
	Customjson interface{} `json:"CustomJson,omitempty"`
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Layerids interface{} `json:"LayerIds,omitempty"`
	Stackid interface{} `json:"StackId"`
	Appid interface{} `json:"AppId,omitempty"`
	Command interface{} `json:"Command"`
	Comment interface{} `json:"Comment,omitempty"`
}

// StartStackRequest represents the StartStackRequest schema from the OpenAPI specification
type StartStackRequest struct {
	Stackid interface{} `json:"StackId"`
}

// CreateLayerResult represents the CreateLayerResult schema from the OpenAPI specification
type CreateLayerResult struct {
	Layerid interface{} `json:"LayerId,omitempty"`
}

// WeeklyAutoScalingSchedule represents the WeeklyAutoScalingSchedule schema from the OpenAPI specification
type WeeklyAutoScalingSchedule struct {
	Friday interface{} `json:"Friday,omitempty"`
	Monday interface{} `json:"Monday,omitempty"`
	Saturday interface{} `json:"Saturday,omitempty"`
	Sunday interface{} `json:"Sunday,omitempty"`
	Thursday interface{} `json:"Thursday,omitempty"`
	Tuesday interface{} `json:"Tuesday,omitempty"`
	Wednesday interface{} `json:"Wednesday,omitempty"`
}

// InstanceIdentity represents the InstanceIdentity schema from the OpenAPI specification
type InstanceIdentity struct {
	Signature interface{} `json:"Signature,omitempty"`
	Document interface{} `json:"Document,omitempty"`
}

// DescribeStacksRequest represents the DescribeStacksRequest schema from the OpenAPI specification
type DescribeStacksRequest struct {
	Stackids interface{} `json:"StackIds,omitempty"`
}

// AttachElasticLoadBalancerRequest represents the AttachElasticLoadBalancerRequest schema from the OpenAPI specification
type AttachElasticLoadBalancerRequest struct {
	Elasticloadbalancername interface{} `json:"ElasticLoadBalancerName"`
	Layerid interface{} `json:"LayerId"`
}

// DescribeRdsDbInstancesResult represents the DescribeRdsDbInstancesResult schema from the OpenAPI specification
type DescribeRdsDbInstancesResult struct {
	Rdsdbinstances interface{} `json:"RdsDbInstances,omitempty"`
}

// SslConfiguration represents the SslConfiguration schema from the OpenAPI specification
type SslConfiguration struct {
	Privatekey interface{} `json:"PrivateKey"`
	Certificate interface{} `json:"Certificate"`
	Chain interface{} `json:"Chain,omitempty"`
}

// CreateDeploymentResult represents the CreateDeploymentResult schema from the OpenAPI specification
type CreateDeploymentResult struct {
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
}

// Volume represents the Volume schema from the OpenAPI specification
type Volume struct {
	Volumeid interface{} `json:"VolumeId,omitempty"`
	Ec2volumeid interface{} `json:"Ec2VolumeId,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Device interface{} `json:"Device,omitempty"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Mountpoint interface{} `json:"MountPoint,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Raidarrayid interface{} `json:"RaidArrayId,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeEcsClustersResult represents the DescribeEcsClustersResult schema from the OpenAPI specification
type DescribeEcsClustersResult struct {
	Ecsclusters interface{} `json:"EcsClusters,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Permission represents the Permission schema from the OpenAPI specification
type Permission struct {
	Level interface{} `json:"Level,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Allowssh interface{} `json:"AllowSsh,omitempty"`
	Allowsudo interface{} `json:"AllowSudo,omitempty"`
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
}

// GetHostnameSuggestionRequest represents the GetHostnameSuggestionRequest schema from the OpenAPI specification
type GetHostnameSuggestionRequest struct {
	Layerid interface{} `json:"LayerId"`
}

// Recipes represents the Recipes schema from the OpenAPI specification
type Recipes struct {
	Deploy interface{} `json:"Deploy,omitempty"`
	Setup interface{} `json:"Setup,omitempty"`
	Shutdown interface{} `json:"Shutdown,omitempty"`
	Undeploy interface{} `json:"Undeploy,omitempty"`
	Configure interface{} `json:"Configure,omitempty"`
}

// DeregisterEcsClusterRequest represents the DeregisterEcsClusterRequest schema from the OpenAPI specification
type DeregisterEcsClusterRequest struct {
	Ecsclusterarn interface{} `json:"EcsClusterArn"`
}

// RegisterElasticIpRequest represents the RegisterElasticIpRequest schema from the OpenAPI specification
type RegisterElasticIpRequest struct {
	Stackid interface{} `json:"StackId"`
	Elasticip interface{} `json:"ElasticIp"`
}

// RegisterVolumeRequest represents the RegisterVolumeRequest schema from the OpenAPI specification
type RegisterVolumeRequest struct {
	Ec2volumeid interface{} `json:"Ec2VolumeId,omitempty"`
	Stackid interface{} `json:"StackId"`
}

// CloneStackRequest represents the CloneStackRequest schema from the OpenAPI specification
type CloneStackRequest struct {
	Cloneappids interface{} `json:"CloneAppIds,omitempty"`
	Sourcestackid interface{} `json:"SourceStackId"`
	Chefconfiguration interface{} `json:"ChefConfiguration,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Hostnametheme interface{} `json:"HostnameTheme,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Useopsworkssecuritygroups interface{} `json:"UseOpsworksSecurityGroups,omitempty"`
	Customcookbookssource interface{} `json:"CustomCookbooksSource,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Defaultos interface{} `json:"DefaultOs,omitempty"`
	Defaultsshkeyname interface{} `json:"DefaultSshKeyName,omitempty"`
	Usecustomcookbooks interface{} `json:"UseCustomCookbooks,omitempty"`
	Defaultrootdevicetype interface{} `json:"DefaultRootDeviceType,omitempty"`
	Defaultinstanceprofilearn interface{} `json:"DefaultInstanceProfileArn,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn"`
	Clonepermissions interface{} `json:"ClonePermissions,omitempty"`
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Defaultsubnetid interface{} `json:"DefaultSubnetId,omitempty"`
	Defaultavailabilityzone interface{} `json:"DefaultAvailabilityZone,omitempty"`
}

// CreateStackResult represents the CreateStackResult schema from the OpenAPI specification
type CreateStackResult struct {
	Stackid interface{} `json:"StackId,omitempty"`
}

// DescribeAppsRequest represents the DescribeAppsRequest schema from the OpenAPI specification
type DescribeAppsRequest struct {
	Appids interface{} `json:"AppIds,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// GrantAccessRequest represents the GrantAccessRequest schema from the OpenAPI specification
type GrantAccessRequest struct {
	Validforinminutes interface{} `json:"ValidForInMinutes,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
}

// CreateStackRequest represents the CreateStackRequest schema from the OpenAPI specification
type CreateStackRequest struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn"`
	Defaultrootdevicetype interface{} `json:"DefaultRootDeviceType,omitempty"`
	Usecustomcookbooks interface{} `json:"UseCustomCookbooks,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Defaultos interface{} `json:"DefaultOs,omitempty"`
	Chefconfiguration interface{} `json:"ChefConfiguration,omitempty"`
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Defaultavailabilityzone interface{} `json:"DefaultAvailabilityZone,omitempty"`
	Name interface{} `json:"Name"`
	Defaultsubnetid interface{} `json:"DefaultSubnetId,omitempty"`
	Region interface{} `json:"Region"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Useopsworkssecuritygroups interface{} `json:"UseOpsworksSecurityGroups,omitempty"`
	Customcookbookssource interface{} `json:"CustomCookbooksSource,omitempty"`
	Defaultsshkeyname interface{} `json:"DefaultSshKeyName,omitempty"`
	Hostnametheme interface{} `json:"HostnameTheme,omitempty"`
	Defaultinstanceprofilearn interface{} `json:"DefaultInstanceProfileArn"`
}

// InstancesCount represents the InstancesCount schema from the OpenAPI specification
type InstancesCount struct {
	Setupfailed interface{} `json:"SetupFailed,omitempty"`
	Stopping interface{} `json:"Stopping,omitempty"`
	Booting interface{} `json:"Booting,omitempty"`
	Deregistering interface{} `json:"Deregistering,omitempty"`
	Stopfailed interface{} `json:"StopFailed,omitempty"`
	Connectionlost interface{} `json:"ConnectionLost,omitempty"`
	Runningsetup interface{} `json:"RunningSetup,omitempty"`
	Shuttingdown interface{} `json:"ShuttingDown,omitempty"`
	Stopped interface{} `json:"Stopped,omitempty"`
	Registered interface{} `json:"Registered,omitempty"`
	Online interface{} `json:"Online,omitempty"`
	Unassigning interface{} `json:"Unassigning,omitempty"`
	Assigning interface{} `json:"Assigning,omitempty"`
	Rebooting interface{} `json:"Rebooting,omitempty"`
	Requested interface{} `json:"Requested,omitempty"`
	Startfailed interface{} `json:"StartFailed,omitempty"`
	Terminating interface{} `json:"Terminating,omitempty"`
	Pending interface{} `json:"Pending,omitempty"`
	Terminated interface{} `json:"Terminated,omitempty"`
	Registering interface{} `json:"Registering,omitempty"`
}

// RegisterEcsClusterResult represents the RegisterEcsClusterResult schema from the OpenAPI specification
type RegisterEcsClusterResult struct {
	Ecsclusterarn interface{} `json:"EcsClusterArn,omitempty"`
}

// DescribeDeploymentsResult represents the DescribeDeploymentsResult schema from the OpenAPI specification
type DescribeDeploymentsResult struct {
	Deployments interface{} `json:"Deployments,omitempty"`
}

// AssociateElasticIpRequest represents the AssociateElasticIpRequest schema from the OpenAPI specification
type AssociateElasticIpRequest struct {
	Elasticip interface{} `json:"ElasticIp"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// UnassignInstanceRequest represents the UnassignInstanceRequest schema from the OpenAPI specification
type UnassignInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
}

// DeleteInstanceRequest represents the DeleteInstanceRequest schema from the OpenAPI specification
type DeleteInstanceRequest struct {
	Deleteelasticip interface{} `json:"DeleteElasticIp,omitempty"`
	Deletevolumes interface{} `json:"DeleteVolumes,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
}

// DescribeUserProfilesResult represents the DescribeUserProfilesResult schema from the OpenAPI specification
type DescribeUserProfilesResult struct {
	Userprofiles interface{} `json:"UserProfiles,omitempty"`
}

// DeleteAppRequest represents the DeleteAppRequest schema from the OpenAPI specification
type DeleteAppRequest struct {
	Appid interface{} `json:"AppId"`
}

// ElasticLoadBalancer represents the ElasticLoadBalancer schema from the OpenAPI specification
type ElasticLoadBalancer struct {
	Dnsname interface{} `json:"DnsName,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Layerid interface{} `json:"LayerId,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Elasticloadbalancername interface{} `json:"ElasticLoadBalancerName,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Ec2instanceids interface{} `json:"Ec2InstanceIds,omitempty"`
}

// CreateInstanceRequest represents the CreateInstanceRequest schema from the OpenAPI specification
type CreateInstanceRequest struct {
	Amiid interface{} `json:"AmiId,omitempty"`
	Autoscalingtype interface{} `json:"AutoScalingType,omitempty"`
	Architecture interface{} `json:"Architecture,omitempty"`
	Blockdevicemappings interface{} `json:"BlockDeviceMappings,omitempty"`
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Instancetype interface{} `json:"InstanceType"`
	Stackid interface{} `json:"StackId"`
	Sshkeyname interface{} `json:"SshKeyName,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Rootdevicetype interface{} `json:"RootDeviceType,omitempty"`
	Ebsoptimized interface{} `json:"EbsOptimized,omitempty"`
	Os interface{} `json:"Os,omitempty"`
	Tenancy interface{} `json:"Tenancy,omitempty"`
	Virtualizationtype interface{} `json:"VirtualizationType,omitempty"`
	Layerids interface{} `json:"LayerIds"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
}

// DescribeRaidArraysRequest represents the DescribeRaidArraysRequest schema from the OpenAPI specification
type DescribeRaidArraysRequest struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Raidarrayids interface{} `json:"RaidArrayIds,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// SetPermissionRequest represents the SetPermissionRequest schema from the OpenAPI specification
type SetPermissionRequest struct {
	Level interface{} `json:"Level,omitempty"`
	Stackid interface{} `json:"StackId"`
	Allowssh interface{} `json:"AllowSsh,omitempty"`
	Allowsudo interface{} `json:"AllowSudo,omitempty"`
	Iamuserarn interface{} `json:"IamUserArn"`
}

// DescribeRaidArraysResult represents the DescribeRaidArraysResult schema from the OpenAPI specification
type DescribeRaidArraysResult struct {
	Raidarrays interface{} `json:"RaidArrays,omitempty"`
}

// LoadBasedAutoScalingConfiguration represents the LoadBasedAutoScalingConfiguration schema from the OpenAPI specification
type LoadBasedAutoScalingConfiguration struct {
	Enable interface{} `json:"Enable,omitempty"`
	Layerid interface{} `json:"LayerId,omitempty"`
	Upscaling interface{} `json:"UpScaling,omitempty"`
	Downscaling interface{} `json:"DownScaling,omitempty"`
}

// Parameters represents the Parameters schema from the OpenAPI specification
type Parameters struct {
}

// ListTagsResult represents the ListTagsResult schema from the OpenAPI specification
type ListTagsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ServiceError represents the ServiceError schema from the OpenAPI specification
type ServiceError struct {
	Stackid interface{} `json:"StackId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Serviceerrorid interface{} `json:"ServiceErrorId,omitempty"`
}

// RegisterEcsClusterRequest represents the RegisterEcsClusterRequest schema from the OpenAPI specification
type RegisterEcsClusterRequest struct {
	Ecsclusterarn interface{} `json:"EcsClusterArn"`
	Stackid interface{} `json:"StackId"`
}

// CreateAppResult represents the CreateAppResult schema from the OpenAPI specification
type CreateAppResult struct {
	Appid interface{} `json:"AppId,omitempty"`
}

// RegisterElasticIpResult represents the RegisterElasticIpResult schema from the OpenAPI specification
type RegisterElasticIpResult struct {
	Elasticip interface{} `json:"ElasticIp,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// UpdateMyUserProfileRequest represents the UpdateMyUserProfileRequest schema from the OpenAPI specification
type UpdateMyUserProfileRequest struct {
	Sshpublickey interface{} `json:"SshPublicKey,omitempty"`
}

// EnvironmentVariable represents the EnvironmentVariable schema from the OpenAPI specification
type EnvironmentVariable struct {
	Key interface{} `json:"Key"`
	Secure interface{} `json:"Secure,omitempty"`
	Value interface{} `json:"Value"`
}

// CreateLayerRequest represents the CreateLayerRequest schema from the OpenAPI specification
type CreateLayerRequest struct {
	Useebsoptimizedinstances interface{} `json:"UseEbsOptimizedInstances,omitempty"`
	Customrecipes interface{} `json:"CustomRecipes,omitempty"`
	Shortname interface{} `json:"Shortname"`
	Autoassignelasticips interface{} `json:"AutoAssignElasticIps,omitempty"`
	Stackid interface{} `json:"StackId"`
	Customsecuritygroupids interface{} `json:"CustomSecurityGroupIds,omitempty"`
	Enableautohealing interface{} `json:"EnableAutoHealing,omitempty"`
	Lifecycleeventconfiguration interface{} `json:"LifecycleEventConfiguration,omitempty"`
	Cloudwatchlogsconfiguration interface{} `json:"CloudWatchLogsConfiguration,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	TypeField interface{} `json:"Type"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Custominstanceprofilearn interface{} `json:"CustomInstanceProfileArn,omitempty"`
	Packages interface{} `json:"Packages,omitempty"`
	Volumeconfigurations interface{} `json:"VolumeConfigurations,omitempty"`
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Name interface{} `json:"Name"`
	Autoassignpublicips interface{} `json:"AutoAssignPublicIps,omitempty"`
}

// DescribeAgentVersionsRequest represents the DescribeAgentVersionsRequest schema from the OpenAPI specification
type DescribeAgentVersionsRequest struct {
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// RaidArray represents the RaidArray schema from the OpenAPI specification
type RaidArray struct {
	Device interface{} `json:"Device,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Numberofdisks interface{} `json:"NumberOfDisks,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Mountpoint interface{} `json:"MountPoint,omitempty"`
	Raidlevel interface{} `json:"RaidLevel,omitempty"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Raidarrayid interface{} `json:"RaidArrayId,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// UpdateInstanceRequest represents the UpdateInstanceRequest schema from the OpenAPI specification
type UpdateInstanceRequest struct {
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Os interface{} `json:"Os,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
	Amiid interface{} `json:"AmiId,omitempty"`
	Ebsoptimized interface{} `json:"EbsOptimized,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Autoscalingtype interface{} `json:"AutoScalingType,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Layerids interface{} `json:"LayerIds,omitempty"`
	Sshkeyname interface{} `json:"SshKeyName,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Architecture interface{} `json:"Architecture,omitempty"`
}

// RegisterRdsDbInstanceRequest represents the RegisterRdsDbInstanceRequest schema from the OpenAPI specification
type RegisterRdsDbInstanceRequest struct {
	Rdsdbinstancearn interface{} `json:"RdsDbInstanceArn"`
	Stackid interface{} `json:"StackId"`
	Dbpassword interface{} `json:"DbPassword"`
	Dbuser interface{} `json:"DbUser"`
}

// DescribeUserProfilesRequest represents the DescribeUserProfilesRequest schema from the OpenAPI specification
type DescribeUserProfilesRequest struct {
	Iamuserarns interface{} `json:"IamUserArns,omitempty"`
}

// CloudWatchLogsLogStream represents the CloudWatchLogsLogStream schema from the OpenAPI specification
type CloudWatchLogsLogStream struct {
	Encoding interface{} `json:"Encoding,omitempty"`
	Loggroupname interface{} `json:"LogGroupName,omitempty"`
	Timezone interface{} `json:"TimeZone,omitempty"`
	Batchcount interface{} `json:"BatchCount,omitempty"`
	Bufferduration interface{} `json:"BufferDuration,omitempty"`
	File interface{} `json:"File,omitempty"`
	Filefingerprintlines interface{} `json:"FileFingerprintLines,omitempty"`
	Initialposition interface{} `json:"InitialPosition,omitempty"`
	Datetimeformat interface{} `json:"DatetimeFormat,omitempty"`
	Multilinestartpattern interface{} `json:"MultiLineStartPattern,omitempty"`
	Batchsize interface{} `json:"BatchSize,omitempty"`
}

// StackSummary represents the StackSummary schema from the OpenAPI specification
type StackSummary struct {
	Appscount interface{} `json:"AppsCount,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Instancescount interface{} `json:"InstancesCount,omitempty"`
	Layerscount interface{} `json:"LayersCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// DeregisterRdsDbInstanceRequest represents the DeregisterRdsDbInstanceRequest schema from the OpenAPI specification
type DeregisterRdsDbInstanceRequest struct {
	Rdsdbinstancearn interface{} `json:"RdsDbInstanceArn"`
}

// DescribeElasticIpsResult represents the DescribeElasticIpsResult schema from the OpenAPI specification
type DescribeElasticIpsResult struct {
	Elasticips interface{} `json:"ElasticIps,omitempty"`
}

// Stack represents the Stack schema from the OpenAPI specification
type Stack struct {
	Defaultsshkeyname interface{} `json:"DefaultSshKeyName,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Customcookbookssource interface{} `json:"CustomCookbooksSource,omitempty"`
	Defaultavailabilityzone interface{} `json:"DefaultAvailabilityZone,omitempty"`
	Hostnametheme interface{} `json:"HostnameTheme,omitempty"`
	Defaultinstanceprofilearn interface{} `json:"DefaultInstanceProfileArn,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Chefconfiguration interface{} `json:"ChefConfiguration,omitempty"`
	Usecustomcookbooks interface{} `json:"UseCustomCookbooks,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Useopsworkssecuritygroups interface{} `json:"UseOpsworksSecurityGroups,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Defaultos interface{} `json:"DefaultOs,omitempty"`
	Defaultsubnetid interface{} `json:"DefaultSubnetId,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Defaultrootdevicetype interface{} `json:"DefaultRootDeviceType,omitempty"`
}

// UnassignVolumeRequest represents the UnassignVolumeRequest schema from the OpenAPI specification
type UnassignVolumeRequest struct {
	Volumeid interface{} `json:"VolumeId"`
}

// GetHostnameSuggestionResult represents the GetHostnameSuggestionResult schema from the OpenAPI specification
type GetHostnameSuggestionResult struct {
	Hostname interface{} `json:"Hostname,omitempty"`
	Layerid interface{} `json:"LayerId,omitempty"`
}

// DescribePermissionsRequest represents the DescribePermissionsRequest schema from the OpenAPI specification
type DescribePermissionsRequest struct {
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// StartInstanceRequest represents the StartInstanceRequest schema from the OpenAPI specification
type StartInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
}

// UpdateRdsDbInstanceRequest represents the UpdateRdsDbInstanceRequest schema from the OpenAPI specification
type UpdateRdsDbInstanceRequest struct {
	Rdsdbinstancearn interface{} `json:"RdsDbInstanceArn"`
	Dbpassword interface{} `json:"DbPassword,omitempty"`
	Dbuser interface{} `json:"DbUser,omitempty"`
}

// DeregisterVolumeRequest represents the DeregisterVolumeRequest schema from the OpenAPI specification
type DeregisterVolumeRequest struct {
	Volumeid interface{} `json:"VolumeId"`
}

// UpdateVolumeRequest represents the UpdateVolumeRequest schema from the OpenAPI specification
type UpdateVolumeRequest struct {
	Mountpoint interface{} `json:"MountPoint,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Volumeid interface{} `json:"VolumeId"`
}

// RegisterInstanceRequest represents the RegisterInstanceRequest schema from the OpenAPI specification
type RegisterInstanceRequest struct {
	Instanceidentity interface{} `json:"InstanceIdentity,omitempty"`
	Privateip interface{} `json:"PrivateIp,omitempty"`
	Publicip interface{} `json:"PublicIp,omitempty"`
	Rsapublickey interface{} `json:"RsaPublicKey,omitempty"`
	Rsapublickeyfingerprint interface{} `json:"RsaPublicKeyFingerprint,omitempty"`
	Stackid interface{} `json:"StackId"`
	Hostname interface{} `json:"Hostname,omitempty"`
}

// DailyAutoScalingSchedule represents the DailyAutoScalingSchedule schema from the OpenAPI specification
type DailyAutoScalingSchedule struct {
}

// App represents the App schema from the OpenAPI specification
type App struct {
	Datasources interface{} `json:"DataSources,omitempty"`
	Shortname interface{} `json:"Shortname,omitempty"`
	Sslconfiguration interface{} `json:"SslConfiguration,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Appsource interface{} `json:"AppSource,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Enablessl interface{} `json:"EnableSsl,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Domains interface{} `json:"Domains,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
}

// DescribeOperatingSystemsResponse represents the DescribeOperatingSystemsResponse schema from the OpenAPI specification
type DescribeOperatingSystemsResponse struct {
	Operatingsystems interface{} `json:"OperatingSystems,omitempty"`
}

// TimeBasedAutoScalingConfiguration represents the TimeBasedAutoScalingConfiguration schema from the OpenAPI specification
type TimeBasedAutoScalingConfiguration struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Autoscalingschedule interface{} `json:"AutoScalingSchedule,omitempty"`
}

// StopStackRequest represents the StopStackRequest schema from the OpenAPI specification
type StopStackRequest struct {
	Stackid interface{} `json:"StackId"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Command interface{} `json:"Command,omitempty"`
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Completedat interface{} `json:"CompletedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
}

// DeploymentCommand represents the DeploymentCommand schema from the OpenAPI specification
type DeploymentCommand struct {
	Args interface{} `json:"Args,omitempty"`
	Name interface{} `json:"Name"`
}

// StackAttributes represents the StackAttributes schema from the OpenAPI specification
type StackAttributes struct {
}

// UpdateStackRequest represents the UpdateStackRequest schema from the OpenAPI specification
type UpdateStackRequest struct {
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Defaultrootdevicetype interface{} `json:"DefaultRootDeviceType,omitempty"`
	Defaultinstanceprofilearn interface{} `json:"DefaultInstanceProfileArn,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Customcookbookssource interface{} `json:"CustomCookbooksSource,omitempty"`
	Defaultsshkeyname interface{} `json:"DefaultSshKeyName,omitempty"`
	Defaultsubnetid interface{} `json:"DefaultSubnetId,omitempty"`
	Hostnametheme interface{} `json:"HostnameTheme,omitempty"`
	Chefconfiguration interface{} `json:"ChefConfiguration,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Stackid interface{} `json:"StackId"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Usecustomcookbooks interface{} `json:"UseCustomCookbooks,omitempty"`
	Defaultos interface{} `json:"DefaultOs,omitempty"`
	Defaultavailabilityzone interface{} `json:"DefaultAvailabilityZone,omitempty"`
	Useopsworkssecuritygroups interface{} `json:"UseOpsworksSecurityGroups,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// DescribeAgentVersionsResult represents the DescribeAgentVersionsResult schema from the OpenAPI specification
type DescribeAgentVersionsResult struct {
	Agentversions interface{} `json:"AgentVersions,omitempty"`
}

// DescribeEcsClustersRequest represents the DescribeEcsClustersRequest schema from the OpenAPI specification
type DescribeEcsClustersRequest struct {
	Ecsclusterarns interface{} `json:"EcsClusterArns,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// CloudWatchLogsConfiguration represents the CloudWatchLogsConfiguration schema from the OpenAPI specification
type CloudWatchLogsConfiguration struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Logstreams interface{} `json:"LogStreams,omitempty"`
}

// RegisterVolumeResult represents the RegisterVolumeResult schema from the OpenAPI specification
type RegisterVolumeResult struct {
	Volumeid interface{} `json:"VolumeId,omitempty"`
}

// DescribeLayersRequest represents the DescribeLayersRequest schema from the OpenAPI specification
type DescribeLayersRequest struct {
	Stackid interface{} `json:"StackId,omitempty"`
	Layerids interface{} `json:"LayerIds,omitempty"`
}

// UpdateElasticIpRequest represents the UpdateElasticIpRequest schema from the OpenAPI specification
type UpdateElasticIpRequest struct {
	Elasticip interface{} `json:"ElasticIp"`
	Name interface{} `json:"Name,omitempty"`
}

// DescribeElasticLoadBalancersRequest represents the DescribeElasticLoadBalancersRequest schema from the OpenAPI specification
type DescribeElasticLoadBalancersRequest struct {
	Layerids interface{} `json:"LayerIds,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// Source represents the Source schema from the OpenAPI specification
type Source struct {
	Url interface{} `json:"Url,omitempty"`
	Username interface{} `json:"Username,omitempty"`
	Password interface{} `json:"Password,omitempty"`
	Revision interface{} `json:"Revision,omitempty"`
	Sshkey interface{} `json:"SshKey,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// RegisterInstanceResult represents the RegisterInstanceResult schema from the OpenAPI specification
type RegisterInstanceResult struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// DescribeStacksResult represents the DescribeStacksResult schema from the OpenAPI specification
type DescribeStacksResult struct {
	Stacks interface{} `json:"Stacks,omitempty"`
}

// EbsBlockDevice represents the EbsBlockDevice schema from the OpenAPI specification
type EbsBlockDevice struct {
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
	Volumesize interface{} `json:"VolumeSize,omitempty"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
}

// DescribeStackProvisioningParametersRequest represents the DescribeStackProvisioningParametersRequest schema from the OpenAPI specification
type DescribeStackProvisioningParametersRequest struct {
	Stackid interface{} `json:"StackId"`
}

// DescribeAppsResult represents the DescribeAppsResult schema from the OpenAPI specification
type DescribeAppsResult struct {
	Apps interface{} `json:"Apps,omitempty"`
}

// DescribeServiceErrorsRequest represents the DescribeServiceErrorsRequest schema from the OpenAPI specification
type DescribeServiceErrorsRequest struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Serviceerrorids interface{} `json:"ServiceErrorIds,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// DeleteLayerRequest represents the DeleteLayerRequest schema from the OpenAPI specification
type DeleteLayerRequest struct {
	Layerid interface{} `json:"LayerId"`
}

// CreateUserProfileResult represents the CreateUserProfileResult schema from the OpenAPI specification
type CreateUserProfileResult struct {
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
}

// DeregisterInstanceRequest represents the DeregisterInstanceRequest schema from the OpenAPI specification
type DeregisterInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
}

// AgentVersion represents the AgentVersion schema from the OpenAPI specification
type AgentVersion struct {
	Configurationmanager interface{} `json:"ConfigurationManager,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// ChefConfiguration represents the ChefConfiguration schema from the OpenAPI specification
type ChefConfiguration struct {
	Manageberkshelf interface{} `json:"ManageBerkshelf,omitempty"`
	Berkshelfversion interface{} `json:"BerkshelfVersion,omitempty"`
}

// DescribeServiceErrorsResult represents the DescribeServiceErrorsResult schema from the OpenAPI specification
type DescribeServiceErrorsResult struct {
	Serviceerrors interface{} `json:"ServiceErrors,omitempty"`
}

// RebootInstanceRequest represents the RebootInstanceRequest schema from the OpenAPI specification
type RebootInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
}

// UpdateUserProfileRequest represents the UpdateUserProfileRequest schema from the OpenAPI specification
type UpdateUserProfileRequest struct {
	Iamuserarn interface{} `json:"IamUserArn"`
	Sshpublickey interface{} `json:"SshPublicKey,omitempty"`
	Sshusername interface{} `json:"SshUsername,omitempty"`
	Allowselfmanagement interface{} `json:"AllowSelfManagement,omitempty"`
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	Virtualizationtype interface{} `json:"VirtualizationType,omitempty"`
	Os interface{} `json:"Os,omitempty"`
	Registeredby interface{} `json:"RegisteredBy,omitempty"`
	Reportedagentversion interface{} `json:"ReportedAgentVersion,omitempty"`
	Ecscontainerinstancearn interface{} `json:"EcsContainerInstanceArn,omitempty"`
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Infrastructureclass interface{} `json:"InfrastructureClass,omitempty"`
	Amiid interface{} `json:"AmiId,omitempty"`
	Ecsclusterarn interface{} `json:"EcsClusterArn,omitempty"`
	Ebsoptimized interface{} `json:"EbsOptimized,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Privatedns interface{} `json:"PrivateDns,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Autoscalingtype interface{} `json:"AutoScalingType,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Reportedos interface{} `json:"ReportedOs,omitempty"`
	Ec2instanceid interface{} `json:"Ec2InstanceId,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Layerids interface{} `json:"LayerIds,omitempty"`
	Sshhostdsakeyfingerprint interface{} `json:"SshHostDsaKeyFingerprint,omitempty"`
	Sshhostrsakeyfingerprint interface{} `json:"SshHostRsaKeyFingerprint,omitempty"`
	Elasticip interface{} `json:"ElasticIp,omitempty"`
	Rootdevicetype interface{} `json:"RootDeviceType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Privateip interface{} `json:"PrivateIp,omitempty"`
	Rootdevicevolumeid interface{} `json:"RootDeviceVolumeId,omitempty"`
	Architecture interface{} `json:"Architecture,omitempty"`
	Blockdevicemappings interface{} `json:"BlockDeviceMappings,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Sshkeyname interface{} `json:"SshKeyName,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Publicip interface{} `json:"PublicIp,omitempty"`
	Lastserviceerrorid interface{} `json:"LastServiceErrorId,omitempty"`
	Tenancy interface{} `json:"Tenancy,omitempty"`
	Publicdns interface{} `json:"PublicDns,omitempty"`
	Platform interface{} `json:"Platform,omitempty"`
	Instanceprofilearn interface{} `json:"InstanceProfileArn,omitempty"`
}

// DescribeDeploymentsRequest represents the DescribeDeploymentsRequest schema from the OpenAPI specification
type DescribeDeploymentsRequest struct {
	Appid interface{} `json:"AppId,omitempty"`
	Deploymentids interface{} `json:"DeploymentIds,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// DeleteUserProfileRequest represents the DeleteUserProfileRequest schema from the OpenAPI specification
type DeleteUserProfileRequest struct {
	Iamuserarn interface{} `json:"IamUserArn"`
}

// ElasticIp represents the ElasticIp schema from the OpenAPI specification
type ElasticIp struct {
	Domain interface{} `json:"Domain,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Ip interface{} `json:"Ip,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Region interface{} `json:"Region,omitempty"`
}

// AssignVolumeRequest represents the AssignVolumeRequest schema from the OpenAPI specification
type AssignVolumeRequest struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Volumeid interface{} `json:"VolumeId"`
}

// EcsCluster represents the EcsCluster schema from the OpenAPI specification
type EcsCluster struct {
	Stackid interface{} `json:"StackId,omitempty"`
	Ecsclusterarn interface{} `json:"EcsClusterArn,omitempty"`
	Ecsclustername interface{} `json:"EcsClusterName,omitempty"`
	Registeredat interface{} `json:"RegisteredAt,omitempty"`
}

// DescribeLayersResult represents the DescribeLayersResult schema from the OpenAPI specification
type DescribeLayersResult struct {
	Layers interface{} `json:"Layers,omitempty"`
}

// DescribeCommandsResult represents the DescribeCommandsResult schema from the OpenAPI specification
type DescribeCommandsResult struct {
	Commands interface{} `json:"Commands,omitempty"`
}

// ReportedOs represents the ReportedOs schema from the OpenAPI specification
type ReportedOs struct {
	Family interface{} `json:"Family,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// AssignInstanceRequest represents the AssignInstanceRequest schema from the OpenAPI specification
type AssignInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Layerids interface{} `json:"LayerIds"`
}

// Command represents the Command schema from the OpenAPI specification
type Command struct {
	Acknowledgedat interface{} `json:"AcknowledgedAt,omitempty"`
	Exitcode interface{} `json:"ExitCode,omitempty"`
	Logurl interface{} `json:"LogUrl,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Completedat interface{} `json:"CompletedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
}

// DescribeStackSummaryRequest represents the DescribeStackSummaryRequest schema from the OpenAPI specification
type DescribeStackSummaryRequest struct {
	Stackid interface{} `json:"StackId"`
}

// DescribeCommandsRequest represents the DescribeCommandsRequest schema from the OpenAPI specification
type DescribeCommandsRequest struct {
	Commandids interface{} `json:"CommandIds,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// DescribeLoadBasedAutoScalingRequest represents the DescribeLoadBasedAutoScalingRequest schema from the OpenAPI specification
type DescribeLoadBasedAutoScalingRequest struct {
	Layerids interface{} `json:"LayerIds"`
}

// AutoScalingThresholds represents the AutoScalingThresholds schema from the OpenAPI specification
type AutoScalingThresholds struct {
	Cputhreshold interface{} `json:"CpuThreshold,omitempty"`
	Ignoremetricstime interface{} `json:"IgnoreMetricsTime,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Loadthreshold interface{} `json:"LoadThreshold,omitempty"`
	Memorythreshold interface{} `json:"MemoryThreshold,omitempty"`
	Thresholdswaittime interface{} `json:"ThresholdsWaitTime,omitempty"`
	Alarms interface{} `json:"Alarms,omitempty"`
}

// DataSource represents the DataSource schema from the OpenAPI specification
type DataSource struct {
	Arn interface{} `json:"Arn,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// CreateInstanceResult represents the CreateInstanceResult schema from the OpenAPI specification
type CreateInstanceResult struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// DescribeElasticIpsRequest represents the DescribeElasticIpsRequest schema from the OpenAPI specification
type DescribeElasticIpsRequest struct {
	Stackid interface{} `json:"StackId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Ips interface{} `json:"Ips,omitempty"`
}

// SetTimeBasedAutoScalingRequest represents the SetTimeBasedAutoScalingRequest schema from the OpenAPI specification
type SetTimeBasedAutoScalingRequest struct {
	Autoscalingschedule interface{} `json:"AutoScalingSchedule,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
}

// DescribeVolumesResult represents the DescribeVolumesResult schema from the OpenAPI specification
type DescribeVolumesResult struct {
	Volumes interface{} `json:"Volumes,omitempty"`
}

// CloneStackResult represents the CloneStackResult schema from the OpenAPI specification
type CloneStackResult struct {
	Stackid interface{} `json:"StackId,omitempty"`
}

// LifecycleEventConfiguration represents the LifecycleEventConfiguration schema from the OpenAPI specification
type LifecycleEventConfiguration struct {
	Shutdown interface{} `json:"Shutdown,omitempty"`
}

// DescribeTimeBasedAutoScalingResult represents the DescribeTimeBasedAutoScalingResult schema from the OpenAPI specification
type DescribeTimeBasedAutoScalingResult struct {
	Timebasedautoscalingconfigurations interface{} `json:"TimeBasedAutoScalingConfigurations,omitempty"`
}

// CreateAppRequest represents the CreateAppRequest schema from the OpenAPI specification
type CreateAppRequest struct {
	Sslconfiguration interface{} `json:"SslConfiguration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Stackid interface{} `json:"StackId"`
	TypeField interface{} `json:"Type"`
	Enablessl interface{} `json:"EnableSsl,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Name interface{} `json:"Name"`
	Shortname interface{} `json:"Shortname,omitempty"`
	Appsource interface{} `json:"AppSource,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Datasources interface{} `json:"DataSources,omitempty"`
	Domains interface{} `json:"Domains,omitempty"`
}

// OperatingSystemConfigurationManager represents the OperatingSystemConfigurationManager schema from the OpenAPI specification
type OperatingSystemConfigurationManager struct {
	Name interface{} `json:"Name,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// CreateUserProfileRequest represents the CreateUserProfileRequest schema from the OpenAPI specification
type CreateUserProfileRequest struct {
	Sshpublickey interface{} `json:"SshPublicKey,omitempty"`
	Sshusername interface{} `json:"SshUsername,omitempty"`
	Allowselfmanagement interface{} `json:"AllowSelfManagement,omitempty"`
	Iamuserarn interface{} `json:"IamUserArn"`
}

// LayerAttributes represents the LayerAttributes schema from the OpenAPI specification
type LayerAttributes struct {
}

// BlockDeviceMapping represents the BlockDeviceMapping schema from the OpenAPI specification
type BlockDeviceMapping struct {
	Devicename interface{} `json:"DeviceName,omitempty"`
	Ebs interface{} `json:"Ebs,omitempty"`
	Nodevice interface{} `json:"NoDevice,omitempty"`
	Virtualname interface{} `json:"VirtualName,omitempty"`
}

// Layer represents the Layer schema from the OpenAPI specification
type Layer struct {
	Autoassignelasticips interface{} `json:"AutoAssignElasticIps,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Layerid interface{} `json:"LayerId,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Customjson interface{} `json:"CustomJson,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Custominstanceprofilearn interface{} `json:"CustomInstanceProfileArn,omitempty"`
	Lifecycleeventconfiguration interface{} `json:"LifecycleEventConfiguration,omitempty"`
	Volumeconfigurations interface{} `json:"VolumeConfigurations,omitempty"`
	Customrecipes interface{} `json:"CustomRecipes,omitempty"`
	Enableautohealing interface{} `json:"EnableAutoHealing,omitempty"`
	Useebsoptimizedinstances interface{} `json:"UseEbsOptimizedInstances,omitempty"`
	Defaultsecuritygroupnames interface{} `json:"DefaultSecurityGroupNames,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Defaultrecipes interface{} `json:"DefaultRecipes,omitempty"`
	Shortname interface{} `json:"Shortname,omitempty"`
	Cloudwatchlogsconfiguration interface{} `json:"CloudWatchLogsConfiguration,omitempty"`
	Customsecuritygroupids interface{} `json:"CustomSecurityGroupIds,omitempty"`
	Packages interface{} `json:"Packages,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Autoassignpublicips interface{} `json:"AutoAssignPublicIps,omitempty"`
	Installupdatesonboot interface{} `json:"InstallUpdatesOnBoot,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DisassociateElasticIpRequest represents the DisassociateElasticIpRequest schema from the OpenAPI specification
type DisassociateElasticIpRequest struct {
	Elasticip interface{} `json:"ElasticIp"`
}

// DescribeInstancesRequest represents the DescribeInstancesRequest schema from the OpenAPI specification
type DescribeInstancesRequest struct {
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Layerid interface{} `json:"LayerId,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
}

// DescribePermissionsResult represents the DescribePermissionsResult schema from the OpenAPI specification
type DescribePermissionsResult struct {
	Permissions interface{} `json:"Permissions,omitempty"`
}

// StopInstanceRequest represents the StopInstanceRequest schema from the OpenAPI specification
type StopInstanceRequest struct {
	Force interface{} `json:"Force,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
}

// GrantAccessResult represents the GrantAccessResult schema from the OpenAPI specification
type GrantAccessResult struct {
	Temporarycredential interface{} `json:"TemporaryCredential,omitempty"`
}

// StackConfigurationManager represents the StackConfigurationManager schema from the OpenAPI specification
type StackConfigurationManager struct {
	Name interface{} `json:"Name,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// UserProfile represents the UserProfile schema from the OpenAPI specification
type UserProfile struct {
	Allowselfmanagement interface{} `json:"AllowSelfManagement,omitempty"`
	Iamuserarn interface{} `json:"IamUserArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Sshpublickey interface{} `json:"SshPublicKey,omitempty"`
	Sshusername interface{} `json:"SshUsername,omitempty"`
}

// SetLoadBasedAutoScalingRequest represents the SetLoadBasedAutoScalingRequest schema from the OpenAPI specification
type SetLoadBasedAutoScalingRequest struct {
	Layerid interface{} `json:"LayerId"`
	Upscaling interface{} `json:"UpScaling,omitempty"`
	Downscaling interface{} `json:"DownScaling,omitempty"`
	Enable interface{} `json:"Enable,omitempty"`
}

// DeploymentCommandArgs represents the DeploymentCommandArgs schema from the OpenAPI specification
type DeploymentCommandArgs struct {
}

// OperatingSystem represents the OperatingSystem schema from the OpenAPI specification
type OperatingSystem struct {
	TypeField interface{} `json:"Type,omitempty"`
	Configurationmanagers interface{} `json:"ConfigurationManagers,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Reportedname interface{} `json:"ReportedName,omitempty"`
	Reportedversion interface{} `json:"ReportedVersion,omitempty"`
	Supported interface{} `json:"Supported,omitempty"`
}

// ListTagsRequest represents the ListTagsRequest schema from the OpenAPI specification
type ListTagsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// DescribeVolumesRequest represents the DescribeVolumesRequest schema from the OpenAPI specification
type DescribeVolumesRequest struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Raidarrayid interface{} `json:"RaidArrayId,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Volumeids interface{} `json:"VolumeIds,omitempty"`
}

// AppAttributes represents the AppAttributes schema from the OpenAPI specification
type AppAttributes struct {
}

// DetachElasticLoadBalancerRequest represents the DetachElasticLoadBalancerRequest schema from the OpenAPI specification
type DetachElasticLoadBalancerRequest struct {
	Layerid interface{} `json:"LayerId"`
	Elasticloadbalancername interface{} `json:"ElasticLoadBalancerName"`
}

// DescribeStackProvisioningParametersResult represents the DescribeStackProvisioningParametersResult schema from the OpenAPI specification
type DescribeStackProvisioningParametersResult struct {
	Agentinstallerurl interface{} `json:"AgentInstallerUrl,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
}

// DescribeElasticLoadBalancersResult represents the DescribeElasticLoadBalancersResult schema from the OpenAPI specification
type DescribeElasticLoadBalancersResult struct {
	Elasticloadbalancers interface{} `json:"ElasticLoadBalancers,omitempty"`
}

// DescribeLoadBasedAutoScalingResult represents the DescribeLoadBasedAutoScalingResult schema from the OpenAPI specification
type DescribeLoadBasedAutoScalingResult struct {
	Loadbasedautoscalingconfigurations interface{} `json:"LoadBasedAutoScalingConfigurations,omitempty"`
}

// DescribeInstancesResult represents the DescribeInstancesResult schema from the OpenAPI specification
type DescribeInstancesResult struct {
	Instances interface{} `json:"Instances,omitempty"`
}

// TemporaryCredential represents the TemporaryCredential schema from the OpenAPI specification
type TemporaryCredential struct {
	Password interface{} `json:"Password,omitempty"`
	Username interface{} `json:"Username,omitempty"`
	Validforinminutes interface{} `json:"ValidForInMinutes,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// ShutdownEventConfiguration represents the ShutdownEventConfiguration schema from the OpenAPI specification
type ShutdownEventConfiguration struct {
	Delayuntilelbconnectionsdrained interface{} `json:"DelayUntilElbConnectionsDrained,omitempty"`
	Executiontimeout interface{} `json:"ExecutionTimeout,omitempty"`
}

// DescribeStackSummaryResult represents the DescribeStackSummaryResult schema from the OpenAPI specification
type DescribeStackSummaryResult struct {
	Stacksummary interface{} `json:"StackSummary,omitempty"`
}

// VolumeConfiguration represents the VolumeConfiguration schema from the OpenAPI specification
type VolumeConfiguration struct {
	Mountpoint interface{} `json:"MountPoint"`
	Numberofdisks interface{} `json:"NumberOfDisks"`
	Raidlevel interface{} `json:"RaidLevel,omitempty"`
	Size interface{} `json:"Size"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
}
