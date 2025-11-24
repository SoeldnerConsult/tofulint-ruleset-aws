package rules

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/arsiba/tofulint-ruleset-aws/rules/api"
	"github.com/arsiba/tofulint-ruleset-aws/rules/ephemeral"
	"github.com/arsiba/tofulint-ruleset-aws/rules/models"
)

var manualRules = []tflint.Rule{
	NewAwsDBInstanceDefaultParameterGroupRule(),
	NewAwsDBInstanceInvalidEngineRule(),
	NewAwsDBInstanceInvalidTypeRule(),
	NewAwsDBInstancePreviousTypeRule(),
	NewAwsDynamoDBTableInvalidStreamViewTypeRule(),
	NewAwsElastiCacheClusterDefaultParameterGroupRule(),
	NewAwsElastiCacheClusterInvalidTypeRule(),
	NewAwsElastiCacheClusterPreviousTypeRule(),
	NewAwsIAMPolicyDocumentGovFriendlyArnsRule(),
	NewAwsIAMPolicyGovFriendlyArnsRule(),
	NewAwsIAMRolePolicyGovFriendlyArnsRule(),
	NewAwsInstancePreviousTypeRule(),
	NewAwsMqBrokerInvalidEngineTypeRule(),
	NewAwsMqConfigurationInvalidEngineTypeRule(),
	NewAwsResourceMissingTagsRule(),
	NewAwsRouteNotSpecifiedTargetRule(),
	NewAwsRouteSpecifiedMultipleTargetsRule(),
	NewAwsS3BucketInvalidACLRule(),
	NewAwsS3BucketNameRule(),
	NewAwsSpotFleetRequestInvalidExcessCapacityTerminationPolicyRule(),
	NewAwsAPIGatewayModelInvalidNameRule(),
	NewAwsElastiCacheReplicationGroupDefaultParameterGroupRule(),
	NewAwsElastiCacheReplicationGroupInvalidTypeRule(),
	NewAwsElastiCacheReplicationGroupPreviousTypeRule(),
	NewAwsIAMPolicyAttachmentExclusiveAttachmentRule(),
	NewAwsIAMPolicySidInvalidCharactersRule(),
	NewAwsIAMPolicyTooLongPolicyRule(),
	NewAwsLambdaFunctionDeprecatedRuntimeRule(),
	NewAwsIAMGroupPolicyTooLongRule(),
	NewAwsAcmCertificateLifecycleRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidNameFormatRule(),
	NewAwsSecurityGroupInvalidProtocolRule(),
	NewAwsSecurityGroupRuleInvalidProtocolRule(),
	NewAwsProviderMissingDefaultTagsRule(),
	NewAwsSecurityGroupInlineRulesRule(),
	NewAwsSecurityGroupRuleDeprecatedRule(),
	NewAwsIAMRoleDeprecatedPolicyAttributesRule(),
	ephemeral.NewAwsWriteOnlyArgumentsRule(),
	ephemeral.NewAwsEphemeralResourcesRule(),
}

// Rules is a list of all rules
var Rules []tflint.Rule

func init() {
	Rules = append(Rules, manualRules...)
	Rules = append(Rules, models.Rules...)
	Rules = append(Rules, api.Rules...)
}
