// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// For Resolver list operations (ListResolverEndpoints
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverEndpoints.html),
// ListResolverRules
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),
// ListResolverRuleAssociations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html),
// ListResolverQueryLogConfigs
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigs.html),
// ListResolverQueryLogConfigAssociations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigAssociations.html)),
// and ListResolverDnssecConfigs
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverDnssecConfigs.html)),
// an optional specification to return a subset of objects. To filter objects, such
// as Resolver endpoints or Resolver rules, you specify Name and Values. For
// example, to list only inbound Resolver endpoints, specify Direction for Name and
// specify INBOUND for Values.
type Filter struct {

	// The name of the parameter that you want to use to filter objects. The valid
	// values for Name depend on the action that you're including the filter in,
	// ListResolverEndpoints
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverEndpoints.html),
	// ListResolverRules
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),
	// ListResolverRuleAssociations
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html),
	// ListResolverQueryLogConfigs
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigs.html),
	// or ListResolverQueryLogConfigAssociations
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigAssociations.html).
	// In early versions of Resolver, values for Name were listed as uppercase, with
	// underscore (_) delimiters. For example, CreatorRequestId was originally listed
	// as CREATOR_REQUEST_ID. Uppercase values for Name are still supported.
	// ListResolverEndpoints Valid values for Name include the following:
	//
	// *
	// CreatorRequestId: The value that you specified when you created the Resolver
	// endpoint.
	//
	// * Direction: Whether you want to return inbound or outbound Resolver
	// endpoints. If you specify DIRECTION for Name, specify INBOUND or OUTBOUND for
	// Values.
	//
	// * HostVpcId: The ID of the VPC that inbound DNS queries pass through on
	// the way from your network to your VPCs in a region, or the VPC that outbound
	// queries pass through on the way from your VPCs to your network. In a
	// CreateResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html)
	// request, SubnetId indirectly identifies the VPC. In a GetResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html)
	// request, the VPC ID for a Resolver endpoint is returned in the HostVPCId
	// element.
	//
	// * IpAddressCount: The number of IP addresses that you have associated
	// with the Resolver endpoint.
	//
	// * Name: The name of the Resolver endpoint.
	//
	// *
	// SecurityGroupIds: The IDs of the VPC security groups that you specified when you
	// created the Resolver endpoint.
	//
	// * Status: The status of the Resolver endpoint.
	// If you specify Status for Name, specify one of the following status codes for
	// Values: CREATING, OPERATIONAL, UPDATING, AUTO_RECOVERING, ACTION_NEEDED, or
	// DELETING. For more information, see Status in ResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverEndpoint.html).
	//
	// ListResolverRules
	// Valid values for Name include the following:
	//
	// * CreatorRequestId: The value that
	// you specified when you created the Resolver rule.
	//
	// * DomainName: The domain name
	// for which Resolver is forwarding DNS queries to your network. In the value that
	// you specify for Values, include a trailing dot (.) after the domain name. For
	// example, if the domain name is example.com, specify the following value. Note
	// the "." after com: example.com.
	//
	// * Name: The name of the Resolver rule.
	//
	// *
	// ResolverEndpointId: The ID of the Resolver endpoint that the Resolver rule is
	// associated with. You can filter on the Resolver endpoint only for rules that
	// have a value of FORWARD for RuleType.
	//
	// * Status: The status of the Resolver
	// rule. If you specify Status for Name, specify one of the following status codes
	// for Values: COMPLETE, DELETING, UPDATING, or FAILED.
	//
	// * Type: The type of the
	// Resolver rule. If you specify TYPE for Name, specify FORWARD or SYSTEM for
	// Values.
	//
	// ListResolverRuleAssociations Valid values for Name include the
	// following:
	//
	// * Name: The name of the Resolver rule association.
	//
	// *
	// ResolverRuleId: The ID of the Resolver rule that is associated with one or more
	// VPCs.
	//
	// * Status: The status of the Resolver rule association. If you specify
	// Status for Name, specify one of the following status codes for Values: CREATING,
	// COMPLETE, DELETING, or FAILED.
	//
	// * VPCId: The ID of the VPC that the Resolver
	// rule is associated with.
	//
	// ListResolverQueryLogConfigs Valid values for Name
	// include the following:
	//
	// * Arn: The ARN for the query logging configuration.
	//
	// *
	// AssociationCount: The number of VPCs that are associated with the query logging
	// configuration.
	//
	// * CreationTime: The date and time that the query logging
	// configuration was created, in Unix time format and Coordinated Universal Time
	// (UTC).
	//
	// * CreatorRequestId: A unique string that identifies the request that
	// created the query logging configuration.
	//
	// * Destination: The AWS service that
	// you want to forward query logs to. Valid values include the following:
	//
	// * S3
	//
	// *
	// CloudWatchLogs
	//
	// * KinesisFirehose
	//
	// * DestinationArn: The ARN of the location
	// that Resolver is sending query logs to. This value can be the ARN for an S3
	// bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery
	// stream.
	//
	// * Id: The ID of the query logging configuration
	//
	// * Name: The name of
	// the query logging configuration
	//
	// * OwnerId: The AWS account ID for the account
	// that created the query logging configuration.
	//
	// * ShareStatus: An indication of
	// whether the query logging configuration is shared with other AWS accounts, or
	// was shared with the current account by another AWS account. Valid values
	// include: NOT_SHARED, SHARED_WITH_ME, or SHARED_BY_ME.
	//
	// * Status: The status of
	// the query logging configuration. If you specify Status for Name, specify the
	// applicable status code for Values: CREATING, CREATED, DELETING, or FAILED. For
	// more information, see Status
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverQueryLogConfig.html#Route53Resolver-Type-route53resolver_ResolverQueryLogConfig-Status).
	//
	// ListResolverQueryLogConfigAssociations
	// Valid values for Name include the following:
	//
	// * CreationTime: The date and time
	// that the VPC was associated with the query logging configuration, in Unix time
	// format and Coordinated Universal Time (UTC).
	//
	// * Error: If the value of Status is
	// FAILED, specify the cause: DESTINATION_NOT_FOUND or ACCESS_DENIED.
	//
	// * Id: The ID
	// of the query logging association.
	//
	// * ResolverQueryLogConfigId: The ID of the
	// query logging configuration that a VPC is associated with.
	//
	// * ResourceId: The ID
	// of the Amazon VPC that is associated with the query logging configuration.
	//
	// *
	// Status: The status of the query logging association. If you specify Status for
	// Name, specify the applicable status code for Values: CREATING, CREATED,
	// DELETING, or FAILED. For more information, see Status
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverQueryLogConfigAssociation.html#Route53Resolver-Type-route53resolver_ResolverQueryLogConfigAssociation-Status).
	Name *string

	// When you're using a List operation and you want the operation to return a subset
	// of objects, such as Resolver endpoints or Resolver rules, the value of the
	// parameter that you want to use to filter objects. For example, to list only
	// inbound Resolver endpoints, specify Direction for Name and specify INBOUND for
	// Values.
	Values []string
}

// Configuration of the firewall behavior provided by DNS Firewall for a single
// Amazon virtual private cloud (VPC).
type FirewallConfig struct {

	// Determines how DNS Firewall operates during failures, for example when all
	// traffic that is sent to DNS Firewall fails to receive a reply.
	//
	// * By default,
	// fail open is disabled, which means the failure mode is closed. This approach
	// favors security over availability. DNS Firewall returns a failure error when it
	// is unable to properly evaluate a query.
	//
	// * If you enable this option, the
	// failure mode is open. This approach favors availability over security. DNS
	// Firewall allows queries to proceed if it is unable to properly evaluate
	// them.
	//
	// This behavior is only enforced for VPCs that have at least one DNS
	// Firewall rule group association.
	FirewallFailOpen FirewallFailOpenStatus

	// The Id of the firewall configuration.
	Id *string

	// The AWS account ID of the owner of the VPC that this firewall configuration
	// applies to.
	OwnerId *string

	// The ID of the VPC that this firewall configuration applies to.
	ResourceId *string
}

// High level information about a list of firewall domains for use in a
// FirewallRule. This is returned by GetFirewallDomainList. To retrieve the domains
// that are defined for this domain list, call ListFirewallDomains.
type FirewallDomainList struct {

	// The Amazon Resource Name (ARN) of the firewall domain list.
	Arn *string

	// The date and time that the domain list was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The number of domain names that are specified in the domain list.
	DomainCount *int32

	// The ID of the domain list.
	Id *string

	// The owner of the list, used only for lists that are not managed by you. For
	// example, the managed domain list AWSManagedDomainsMalwareDomainList has the
	// managed owner name Route 53 Resolver DNS Firewall.
	ManagedOwnerName *string

	// The date and time that the domain list was last modified, in Unix time format
	// and Coordinated Universal Time (UTC).
	ModificationTime *string

	// The name of the domain list.
	Name *string

	// The status of the domain list.
	Status FirewallDomainListStatus

	// Additional information about the status of the list, if available.
	StatusMessage *string
}

// Minimal high-level information for a firewall domain list. The action
// ListFirewallDomainLists returns an array of these objects. To retrieve full
// information for a firewall domain list, call GetFirewallDomainList and
// ListFirewallDomains.
type FirewallDomainListMetadata struct {

	// The Amazon Resource Name (ARN) of the firewall domain list metadata.
	Arn *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The ID of the domain list.
	Id *string

	// The owner of the list, used only for lists that are not managed by you. For
	// example, the managed domain list AWSManagedDomainsMalwareDomainList has the
	// managed owner name Route 53 Resolver DNS Firewall.
	ManagedOwnerName *string

	// The name of the domain list.
	Name *string
}

// A single firewall rule in a rule group.
type FirewallRule struct {

	// The action that DNS Firewall should take on a DNS query when it matches one of
	// the domains in the rule's domain list:
	//
	// * ALLOW - Permit the request to go
	// through.
	//
	// * ALERT - Permit the request to go through but send an alert to the
	// logs.
	//
	// * BLOCK - Disallow the request. If this is specified, additional handling
	// details are provided in the rule's BlockResponse setting.
	Action Action

	// The DNS record's type. This determines the format of the record value that you
	// provided in BlockOverrideDomain. Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE.
	BlockOverrideDnsType BlockOverrideDnsType

	// The custom DNS record to send back in response to the query. Used for the rule
	// action BLOCK with a BlockResponse setting of OVERRIDE.
	BlockOverrideDomain *string

	// The recommended amount of time, in seconds, for the DNS resolver or web browser
	// to cache the provided override record. Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE.
	BlockOverrideTtl *int32

	// The way that you want DNS Firewall to block the request. Used for the rule
	// action setting BLOCK.
	//
	// * NODATA - Respond indicating that the query was
	// successful, but no response is available for it.
	//
	// * NXDOMAIN - Respond
	// indicating that the domain name that's in the query doesn't exist.
	//
	// * OVERRIDE -
	// Provide a custom override in the response. This option requires custom handling
	// details in the rule's BlockOverride* settings.
	BlockResponse BlockResponse

	// The date and time that the rule was created, in Unix time format and Coordinated
	// Universal Time (UTC).
	CreationTime *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The ID of the domain list that's used in the rule.
	FirewallDomainListId *string

	// The unique identifier of the firewall rule group of the rule.
	FirewallRuleGroupId *string

	// The date and time that the rule was last modified, in Unix time format and
	// Coordinated Universal Time (UTC).
	ModificationTime *string

	// The name of the rule.
	Name *string

	// The priority of the rule in the rule group. This value must be unique within the
	// rule group. DNS Firewall processes the rules in a rule group by order of
	// priority, starting from the lowest setting.
	Priority *int32
}

// High-level information for a firewall rule group. A firewall rule group is a
// collection of rules that DNS Firewall uses to filter DNS network traffic for a
// VPC. To retrieve the rules for the rule group, call ListFirewallRules.
type FirewallRuleGroup struct {

	// The ARN (Amazon Resource Name) of the rule group.
	Arn *string

	// The date and time that the rule group was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The ID of the rule group.
	Id *string

	// The date and time that the rule group was last modified, in Unix time format and
	// Coordinated Universal Time (UTC).
	ModificationTime *string

	// The name of the rule group.
	Name *string

	// The AWS account ID for the account that created the rule group. When a rule
	// group is shared with your account, this is the account that has shared the rule
	// group with you.
	OwnerId *string

	// The number of rules in the rule group.
	RuleCount *int32

	// Whether the rule group is shared with other AWS accounts, or was shared with the
	// current account by another AWS account. Sharing is configured through AWS
	// Resource Access Manager (AWS RAM).
	ShareStatus ShareStatus

	// The status of the domain list.
	Status FirewallRuleGroupStatus

	// Additional information about the status of the rule group, if available.
	StatusMessage *string
}

// An association between a firewall rul group and a VPC, which enables DNS
// filtering for the VPC.
type FirewallRuleGroupAssociation struct {

	// The Amazon Resource Name (ARN) of the firewall rule group association.
	Arn *string

	// The date and time that the association was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The unique identifier of the firewall rule group.
	FirewallRuleGroupId *string

	// The identifier for the association.
	Id *string

	// The owner of the association, used only for associations that are not managed by
	// you. If you use AWS Firewall Manager to manage your DNS Firewalls, then this
	// reports Firewall Manager as the managed owner.
	ManagedOwnerName *string

	// The date and time that the association was last modified, in Unix time format
	// and Coordinated Universal Time (UTC).
	ModificationTime *string

	// If enabled, this setting disallows modification or removal of the association,
	// to help prevent against accidentally altering DNS firewall protections.
	MutationProtection MutationProtectionStatus

	// The name of the association.
	Name *string

	// The setting that determines the processing order of the rule group among the
	// rule groups that are associated with a single VPC. DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	Priority *int32

	// The current status of the association.
	Status FirewallRuleGroupAssociationStatus

	// Additional information about the status of the response, if available.
	StatusMessage *string

	// The unique identifier of the VPC that is associated with the rule group.
	VpcId *string
}

// Minimal high-level information for a firewall rule group. The action
// ListFirewallRuleGroups returns an array of these objects. To retrieve full
// information for a firewall rule group, call GetFirewallRuleGroup and
// ListFirewallRules.
type FirewallRuleGroupMetadata struct {

	// The ARN (Amazon Resource Name) of the rule group.
	Arn *string

	// A unique string defined by you to identify the request. This allows you to retry
	// failed requests without the risk of executing the operation twice. This can be
	// any unique string, for example, a timestamp.
	CreatorRequestId *string

	// The ID of the rule group.
	Id *string

	// The name of the rule group.
	Name *string

	// The AWS account ID for the account that created the rule group. When a rule
	// group is shared with your account, this is the account that has shared the rule
	// group with you.
	OwnerId *string

	// Whether the rule group is shared with other AWS accounts, or was shared with the
	// current account by another AWS account. Sharing is configured through AWS
	// Resource Access Manager (AWS RAM).
	ShareStatus ShareStatus
}

// In a CreateResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html)
// request, the IP address that DNS queries originate from (for outbound endpoints)
// or that you forward DNS queries to (for inbound endpoints). IpAddressRequest
// also includes the ID of the subnet that contains the IP address.
type IpAddressRequest struct {

	// The ID of the subnet that contains the IP address.
	//
	// This member is required.
	SubnetId *string

	// The IP address that you want to use for DNS queries.
	Ip *string
}

// In the response to a GetResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html)
// request, information about the IP addresses that the Resolver endpoint uses for
// DNS queries.
type IpAddressResponse struct {

	// The date and time that the IP address was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// One IP address that the Resolver endpoint uses for DNS queries.
	Ip *string

	// The ID of one IP address.
	IpId *string

	// The date and time that the IP address was last modified, in Unix time format and
	// Coordinated Universal Time (UTC).
	ModificationTime *string

	// A status code that gives the current status of the request.
	Status IpAddressStatus

	// A message that provides additional information about the status of the request.
	StatusMessage *string

	// The ID of one subnet.
	SubnetId *string
}

// In an UpdateResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverEndpoint.html)
// request, information about an IP address to update.
type IpAddressUpdate struct {

	// The new IP address.
	Ip *string

	// Only when removing an IP address from a Resolver endpoint: The ID of the IP
	// address that you want to remove. To get this ID, use GetResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html).
	IpId *string

	// The ID of the subnet that includes the IP address that you want to update. To
	// get this ID, use GetResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html).
	SubnetId *string
}

// A complex type that contains information about a configuration for DNSSEC
// validation.
type ResolverDnssecConfig struct {

	// The ID for a configuration for DNSSEC validation.
	Id *string

	// The owner account ID of the virtual private cloud (VPC) for a configuration for
	// DNSSEC validation.
	OwnerId *string

	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC
	// validation status for.
	ResourceId *string

	// The validation status for a DNSSEC configuration. The status can be one of the
	// following:
	//
	// * ENABLING: DNSSEC validation is being enabled but is not
	// complete.
	//
	// * ENABLED: DNSSEC validation is enabled.
	//
	// * DISABLING: DNSSEC
	// validation is being disabled but is not complete.
	//
	// * DISABLED DNSSEC validation
	// is disabled.
	ValidationStatus ResolverDNSSECValidationStatus
}

// In the response to a CreateResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html),
// DeleteResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverEndpoint.html),
// GetResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html),
// ListResolverEndpoints
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverEndpoints.html),
// or UpdateResolverEndpoint
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverEndpoint.html)
// request, a complex type that contains settings for an existing inbound or
// outbound Resolver endpoint.
type ResolverEndpoint struct {

	// The ARN (Amazon Resource Name) for the Resolver endpoint.
	Arn *string

	// The date and time that the endpoint was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string that identifies the request that created the Resolver endpoint.
	// The CreatorRequestId allows failed requests to be retried without the risk of
	// executing the operation twice.
	CreatorRequestId *string

	// Indicates whether the Resolver endpoint allows inbound or outbound DNS
	// queries:
	//
	// * INBOUND: allows DNS queries to your VPC from your network
	//
	// *
	// OUTBOUND: allows DNS queries from your VPC to your network
	Direction ResolverEndpointDirection

	// The ID of the VPC that you want to create the Resolver endpoint in.
	HostVPCId *string

	// The ID of the Resolver endpoint.
	Id *string

	// The number of IP addresses that the Resolver endpoint can use for DNS queries.
	IpAddressCount *int32

	// The date and time that the endpoint was last modified, in Unix time format and
	// Coordinated Universal Time (UTC).
	ModificationTime *string

	// The name that you assigned to the Resolver endpoint when you submitted a
	// CreateResolverEndpoint
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html)
	// request.
	Name *string

	// The ID of one or more security groups that control access to this VPC. The
	// security group must include one or more inbound rules (for inbound endpoints) or
	// outbound rules (for outbound endpoints). Inbound and outbound rules must allow
	// TCP and UDP access. For inbound access, open port 53. For outbound access, open
	// the port that you're using for DNS queries on your network.
	SecurityGroupIds []string

	// A code that specifies the current status of the Resolver endpoint. Valid values
	// include the following:
	//
	// * CREATING: Resolver is creating and configuring one or
	// more Amazon VPC network interfaces for this endpoint.
	//
	// * OPERATIONAL: The Amazon
	// VPC network interfaces for this endpoint are correctly configured and able to
	// pass inbound or outbound DNS queries between your network and Resolver.
	//
	// *
	// UPDATING: Resolver is associating or disassociating one or more network
	// interfaces with this endpoint.
	//
	// * AUTO_RECOVERING: Resolver is trying to recover
	// one or more of the network interfaces that are associated with this endpoint.
	// During the recovery process, the endpoint functions with limited capacity
	// because of the limit on the number of DNS queries per IP address (per network
	// interface). For the current limit, see Limits on Route 53 Resolver
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html#limits-api-entities-resolver).
	//
	// *
	// ACTION_NEEDED: This endpoint is unhealthy, and Resolver can't automatically
	// recover it. To resolve the problem, we recommend that you check each IP address
	// that you associated with the endpoint. For each IP address that isn't available,
	// add another IP address and then delete the IP address that isn't available. (An
	// endpoint must always include at least two IP addresses.) A status of
	// ACTION_NEEDED can have a variety of causes. Here are two common causes:
	//
	// * One
	// or more of the network interfaces that are associated with the endpoint were
	// deleted using Amazon VPC.
	//
	// * The network interface couldn't be created for some
	// reason that's outside the control of Resolver.
	//
	// * DELETING: Resolver is deleting
	// this endpoint and the associated network interfaces.
	Status ResolverEndpointStatus

	// A detailed description of the status of the Resolver endpoint.
	StatusMessage *string
}

// In the response to a CreateResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverQueryLogConfig.html),
// DeleteResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverQueryLogConfig.html),
// GetResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverQueryLogConfig.html),
// or ListResolverQueryLogConfigs
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigs.html)
// request, a complex type that contains settings for one query logging
// configuration.
type ResolverQueryLogConfig struct {

	// The ARN for the query logging configuration.
	Arn *string

	// The number of VPCs that are associated with the query logging configuration.
	AssociationCount int32

	// The date and time that the query logging configuration was created, in Unix time
	// format and Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string that identifies the request that created the query logging
	// configuration. The CreatorRequestId allows failed requests to be retried without
	// the risk of executing the operation twice.
	CreatorRequestId *string

	// The ARN of the resource that you want Resolver to send query logs: an Amazon S3
	// bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn *string

	// The ID for the query logging configuration.
	Id *string

	// The name of the query logging configuration.
	Name *string

	// The AWS account ID for the account that created the query logging configuration.
	OwnerId *string

	// An indication of whether the query logging configuration is shared with other
	// AWS accounts, or was shared with the current account by another AWS account.
	// Sharing is configured through AWS Resource Access Manager (AWS RAM).
	ShareStatus ShareStatus

	// The status of the specified query logging configuration. Valid values include
	// the following:
	//
	// * CREATING: Resolver is creating the query logging
	// configuration.
	//
	// * CREATED: The query logging configuration was successfully
	// created. Resolver is logging queries that originate in the specified VPC.
	//
	// *
	// DELETING: Resolver is deleting this query logging configuration.
	//
	// * FAILED:
	// Resolver can't deliver logs to the location that is specified in the query
	// logging configuration. Here are two common causes:
	//
	// * The specified destination
	// (for example, an Amazon S3 bucket) was deleted.
	//
	// * Permissions don't allow
	// sending logs to the destination.
	Status ResolverQueryLogConfigStatus
}

// In the response to an AssociateResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverQueryLogConfig.html),
// DisassociateResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html),
// GetResolverQueryLogConfigAssociation
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverQueryLogConfigAssociation.html),
// or ListResolverQueryLogConfigAssociations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverQueryLogConfigAssociations.html),
// request, a complex type that contains settings for a specified association
// between an Amazon VPC and a query logging configuration.
type ResolverQueryLogConfigAssociation struct {

	// The date and time that the VPC was associated with the query logging
	// configuration, in Unix time format and Coordinated Universal Time (UTC).
	CreationTime *string

	// If the value of Status is FAILED, the value of Error indicates the cause:
	//
	// *
	// DESTINATION_NOT_FOUND: The specified destination (for example, an Amazon S3
	// bucket) was deleted.
	//
	// * ACCESS_DENIED: Permissions don't allow sending logs to
	// the destination.
	//
	// If the value of Status is a value other than FAILED, Error is
	// null.
	Error ResolverQueryLogConfigAssociationError

	// Contains additional information about the error. If the value or Error is null,
	// the value of ErrorMessage also is null.
	ErrorMessage *string

	// The ID of the query logging association.
	Id *string

	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId *string

	// The ID of the Amazon VPC that is associated with the query logging
	// configuration.
	ResourceId *string

	// The status of the specified query logging association. Valid values include the
	// following:
	//
	// * CREATING: Resolver is creating an association between an Amazon
	// VPC and a query logging configuration.
	//
	// * CREATED: The association between an
	// Amazon VPC and a query logging configuration was successfully created. Resolver
	// is logging queries that originate in the specified VPC.
	//
	// * DELETING: Resolver is
	// deleting this query logging association.
	//
	// * FAILED: Resolver either couldn't
	// create or couldn't delete the query logging association.
	Status ResolverQueryLogConfigAssociationStatus
}

// For queries that originate in your VPC, detailed information about a Resolver
// rule, which specifies how to route DNS queries out of the VPC. The ResolverRule
// parameter appears in the response to a CreateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html),
// DeleteResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverRule.html),
// GetResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverRule.html),
// ListResolverRules
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),
// or UpdateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverRule.html)
// request.
type ResolverRule struct {

	// The ARN (Amazon Resource Name) for the Resolver rule specified by Id.
	Arn *string

	// The date and time that the Resolver rule was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string

	// A unique string that you specified when you created the Resolver rule.
	// CreatorRequestId identifies the request and allows failed requests to be retried
	// without the risk of executing the operation twice.
	CreatorRequestId *string

	// DNS queries for this domain name are forwarded to the IP addresses that are
	// specified in TargetIps. If a query matches multiple Resolver rules (example.com
	// and www.example.com), the query is routed using the Resolver rule that contains
	// the most specific domain name (www.example.com).
	DomainName *string

	// The ID that Resolver assigned to the Resolver rule when you created it.
	Id *string

	// The date and time that the Resolver rule was last updated, in Unix time format
	// and Coordinated Universal Time (UTC).
	ModificationTime *string

	// The name for the Resolver rule, which you specified when you created the
	// Resolver rule.
	Name *string

	// When a rule is shared with another AWS account, the account ID of the account
	// that the rule is shared with.
	OwnerId *string

	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string

	// When you want to forward DNS queries for specified domain name to resolvers on
	// your network, specify FORWARD. When you have a forwarding rule to forward DNS
	// queries for a domain to your network and you want Resolver to process queries
	// for a subdomain of that domain, specify SYSTEM. For example, to forward DNS
	// queries for example.com to resolvers on your network, you create a rule and
	// specify FORWARD for RuleType. To then have Resolver process queries for
	// apex.example.com, you create a rule and specify SYSTEM for RuleType. Currently,
	// only Resolver can create rules that have a value of RECURSIVE for RuleType.
	RuleType RuleTypeOption

	// Whether the rules is shared and, if so, whether the current account is sharing
	// the rule with another account, or another account is sharing the rule with the
	// current account.
	ShareStatus ShareStatus

	// A code that specifies the current status of the Resolver rule.
	Status ResolverRuleStatus

	// A detailed description of the status of a Resolver rule.
	StatusMessage *string

	// An array that contains the IP addresses and ports that an outbound endpoint
	// forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers
	// on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps []TargetAddress
}

// In the response to an AssociateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html),
// DisassociateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html),
// or ListResolverRuleAssociations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html)
// request, provides information about an association between a Resolver rule and a
// VPC. The association determines which DNS queries that originate in the VPC are
// forwarded to your network.
type ResolverRuleAssociation struct {

	// The ID of the association between a Resolver rule and a VPC. Resolver assigns
	// this value when you submit an AssociateResolverRule
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html)
	// request.
	Id *string

	// The name of an association between a Resolver rule and a VPC.
	Name *string

	// The ID of the Resolver rule that you associated with the VPC that is specified
	// by VPCId.
	ResolverRuleId *string

	// A code that specifies the current status of the association between a Resolver
	// rule and a VPC.
	Status ResolverRuleAssociationStatus

	// A detailed description of the status of the association between a Resolver rule
	// and a VPC.
	StatusMessage *string

	// The ID of the VPC that you associated the Resolver rule with.
	VPCId *string
}

// In an UpdateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverRule.html)
// request, information about the changes that you want to make.
type ResolverRuleConfig struct {

	// The new name for the Resolver rule. The name that you specify appears in the
	// Resolver dashboard in the Route 53 console.
	Name *string

	// The ID of the new outbound Resolver endpoint that you want to use to route DNS
	// queries to the IP addresses that you specify in TargetIps.
	ResolverEndpointId *string

	// For DNS queries that originate in your VPC, the new IP addresses that you want
	// to route outbound DNS queries to.
	TargetIps []TargetAddress
}

// One tag that you want to add to the specified resource. A tag consists of a Key
// (a name for the tag) and a Value.
type Tag struct {

	// The name for the tag. For example, if you want to associate Resolver resources
	// with the account IDs of your customers for billing purposes, the value of Key
	// might be account-id.
	//
	// This member is required.
	Key *string

	// The value for the tag. For example, if Key is account-id, then Value might be
	// the ID of the customer account that you're creating the resource for.
	//
	// This member is required.
	Value *string
}

// In a CreateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html)
// request, an array of the IPs that you want to forward DNS queries to.
type TargetAddress struct {

	// One IP address that you want to forward DNS queries to. You can specify only
	// IPv4 addresses.
	//
	// This member is required.
	Ip *string

	// The port at Ip that you want to forward DNS queries to.
	Port *int32
}
