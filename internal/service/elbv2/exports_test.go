// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elbv2

// Exports for use in tests only.
var (
	ResourceListener            = resourceListener
	ResourceListenerCertificate = resourceListenerCertificate
	ResourceListenerRule        = resourceListenerRule
	ResourceLoadBalancer        = resourceLoadBalancer
	ResourceTrustStore          = resourceTrustStore

	FindListenerByARN                   = findListenerByARN
	FindListenerCertificateByTwoPartKey = findListenerCertificateByTwoPartKey
	FindListenerRuleByARN               = findListenerRuleByARN
	FindLoadBalancerAttributesByARN     = findLoadBalancerAttributesByARN
	FindLoadBalancerByARN               = findLoadBalancerByARN
	FindTrustStoreByARN                 = findTrustStoreByARN
	HealthCheckProtocolEnumValues       = healthCheckProtocolEnumValues
	HostedZoneIDPerRegionALBMap         = hostedZoneIDPerRegionALBMap
	HostedZoneIDPerRegionNLBMap         = hostedZoneIDPerRegionNLBMap
	ListenerARNFromRuleARN              = listenerARNFromRuleARN
	ProtocolVersionEnumValues           = protocolVersionEnumValues
	SuffixFromARN                       = suffixFromARN
)

const (
	AlpnPolicyHTTP1Only      = alpnPolicyHTTP1Only
	AlpnPolicyHTTP2Only      = alpnPolicyHTTP2Only
	AlpnPolicyHTTP2Optional  = alpnPolicyHTTP2Optional
	AlpnPolicyHTTP2Preferred = alpnPolicyHTTP2Preferred
	AlpnPolicyNone           = alpnPolicyNone

	LoadBalancerAttributeClientKeepAliveSeconds = loadBalancerAttributeClientKeepAliveSeconds

	MutualAuthenticationOff         = mutualAuthenticationOff
	MutualAuthenticationVerify      = mutualAuthenticationVerify
	MutualAuthenticationPassthrough = mutualAuthenticationPassthrough
)
