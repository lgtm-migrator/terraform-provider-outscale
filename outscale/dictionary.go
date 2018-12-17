package outscale

import "github.com/hashicorp/terraform/helper/schema"

//Dictionary for the Outscale APIs maps the apis to their respective functions
type Dictionary map[string]ResourceMap

//ResourceMap maps a schema to their resource or datasource implementation
type ResourceMap map[string]SchemaFunc

//SchemaFunc maps a function that returns a schema
type SchemaFunc func() *schema.Resource

var resources Dictionary
var datasources Dictionary

func init() {
	resources = Dictionary{
		"fcu": ResourceMap{
			"outscale_vm":                            resourceOutscaleVM,
			"outscale_image":                         resourceOutscaleImage,
			"outscale_firewall_rules_set":            resourceOutscaleFirewallRulesSet,
			"outscale_outbound_rule":                 resourceOutscaleOutboundRule,
			"outscale_inbound_rule":                  resourceOutscaleInboundRule,
			"outscale_tag":                           resourceOutscaleTags,
			"outscale_keypair":                       resourceOutscaleKeyPair,
			"outscale_public_ip":                     resourceOutscalePublicIP,
			"outscale_public_ip_link":                resourceOutscalePublicIPLink,
			"outscale_volume":                        resourceOutscaleVolume,
			"outscale_volumes_link":                  resourceOutscaleVolumeLink,
			"outscale_lin":                           resourceOutscaleLin,
			"outscale_lin_attributes":                resourceOutscaleLinAttributes,
			"outscale_lin_internet_gateway":          resourceOutscaleLinInternetGateway,
			"outscale_lin_internet_gateway_link":     resourceOutscaleLinInternetGatewayLink,
			"outscale_vm_attributes":                 resourceOutscaleVMAttributes,
			"outscale_nat_service":                   resourceOutscaleNatService,
			"outscale_subnet":                        resourceOutscaleSubNet,
			"outscale_keypair_importation":           resourceOutscaleKeyPairImportation,
			"outscale_client_endpoint":               resourceOutscaleCustomerGateway,
			"outscale_route":                         resourceOutscaleRoute,
			"outscale_route_table":                   resourceOutscaleRouteTable,
			"outscale_route_table_link":              resourceOutscaleRouteTableAssociation,
			"outscale_dhcp_option":                   resourceOutscaleDHCPOption,
			"outscale_dhcp_option_link":              resourceOutscaleDHCPOptionLink,
			"outscale_image_copy":                    resourceOutscaleImageCopy,
			"outscale_vpn_connection":                resourceOutscaleVpnConnection,
			"outscale_vpn_gateway":                   resourceOutscaleVpnGateway,
			"outscale_image_tasks":                   resourceOutscaleImageTasks,
			"outscale_vpn_connection_route":          resourceOutscaleVpnConnectionRoute,
			"outscale_vpn_gateway_route_propagation": resourceOutscaleVpnGatewayRoutePropagation,
			"outscale_vpn_gateway_link":              resourceOutscaleVpnGatewayLink,
			"outscale_nic":                           resourceOutscaleNic,
			"outscale_snapshot_export_tasks":         resourceOutscaleImageExportTasks,
			"outscale_snapshot":                      resourceOutscaleSnapshot,
			"outscale_image_register":                resourceOutscaleImageRegister,
			"outscale_image_launch_permission":       resourceOutscaleImageLaunchPermission,
			"outscale_lin_peering":                   resourceOutscaleLinPeeringConnection,
			"outscale_lin_peering_acceptation":       resourceOutscaleLinPeeringConnectionAccepter,
			"outscale_nic_link":                      resourceOutscaleNetworkInterfaceAttachment,
			"outscale_nic_private_ip":                resourceOutscaleNetworkInterfacePrivateIP,
			"outscale_reserved_vms_offer_purchase":   resourceOutscaleReservedVmsOfferPurchase,
			"outscale_snapshot_attributes":           resourcedOutscaleSnapshotAttributes,
			"outscale_lin_api_access":                resourceOutscaleVpcEndpoint,
			"outscale_snapshot_import":               resourcedOutscaleSnapshotImport,
			"outscale_snapshot_copy":                 resourcedOutscaleSnapshotCopy,
		},
		"oapi": ResourceMap{
			"outscale_vm":                            resourceOutscaleOApiVM,
			"outscale_firewall_rules_set":            resourceOutscaleOAPIFirewallRulesSet,
			"outscale_image":                         resourceOutscaleOAPIImage,
			"outscale_keypair":                       resourceOutscaleOAPIKeyPair,
			"outscale_public_ip":                     resourceOutscaleOAPIPublicIP,
			"outscale_public_ip_link":                resourceOutscaleOAPIPublicIPLink,
			"outscale_volume":                        resourceOutscaleOAPIVolume,
			"outscale_volumes_link":                  resourceOutscaleOAPIVolumeLink,
			"outscale_vm_attributes":                 resourceOutscaleOAPIVMAttributes,
			"outscale_inbound_rule":                  resourceOutscaleOAPIInboundRule,
			"outscale_outbound_rule":                 resourceOutscaleOAPIOutboundRule,
			"outscale_tag":                           resourceOutscaleOAPITags,
			"outscale_lin_attributes":                resourceOutscaleOAPILinAttributes,
			"outscale_lin":                           resourceOutscaleOAPINet,
			"outscale_net_attributes":                resourceOutscaleOAPILinAttributes,
			"outscale_net":                           resourceOutscaleOAPINet,
			"outscale_lin_internet_gateway":          resourceOutscaleOAPIInternetService,
			"outscale_lin_internet_gateway_link":     resourceOutscaleOAPIInternetServiceLink,
			"outscale_internet_service":              resourceOutscaleOAPIInternetService,
			"outscale_internet_service_link":         resourceOutscaleOAPIInternetServiceLink,
			"outscale_nat_service":                   resourceOutscaleOAPINatService,
			"outscale_subnet":                        resourceOutscaleOAPISubNet,
			"outscale_client_endpoint":               resourceOutscaleCustomerGateway,
			"outscale_route":                         resourceOutscaleOAPIRoute,
			"outscale_route_table":                   resourceOutscaleOAPIRouteTable,
			"outscale_route_table_link":              resourceOutscaleOAPIRouteTableAssociation,
			"outscale_snapshot":                      resourceOutscaleOAPISnapshot,
			"outscale_api_key":                       resourceOutscaleOAPIIamAccessKey,
			"outscale_keypair_importation":           resourceOutscaleOAPIKeyPairImportation,
			"outscale_image_launch_permission":       resourceOutscaleOAPIImageLaunchPermission,
			"outscale_lin_peering":                   resourceOutscaleOAPILinPeeringConnection,
			"outscale_net_peering":                   resourceOutscaleOAPILinPeeringConnection,
			"outscale_load_balancer":                 resourceOutscaleOAPILoadBalancer,
			"outscale_nic_private_ip":                resourceOutscaleOAPINetworkInterfacePrivateIP,
			"outscale_nic_link":                      resourceOutscaleOAPINetworkInterfaceAttachment,
			"outscale_nic":                           resourceOutscaleOAPINic,
			"outscale_load_balancer_cookiepolicy":    resourceOutscaleAppCookieStickinessPolicy,
			"outscale_load_balancer_listeners":       resourceOutscaleOAPILoadBalancerListeners,
			"outscale_load_balancer_attributes":      resourceOutscaleOAPILoadBalancerAttributes,
			"outscale_load_balancer_vms":             resourceOutscaleOAPILBUAttachment,
			"outscale_load_balancer_tags":            resourceOutscaleOAPILBUTags,
			"outscale_image_tasks":                   resourceOutscaleOAPIImageTasks,
			"outscale_reserved_vms_offer_purchase":   resourceOutscaleOAPIReservedVmsOfferPurchase,
			"outscale_vpn_gateway":                   resourceOutscaleOAPIVpnGateway,
			"outscale_vpn_gateway_route_propagation": resourceOutscaleOAPIVpnGatewayRoutePropagation,
			"outscale_snapshot_export_tasks":         resourceOutscaleOAPIImageExportTasks,
			"outscale_vpn_connection_route":          resourceOutscaleOAPIVpnConnectionRoute,
			"outscale_vpn_connection":                resourceOutscaleOAPIVpnConnection,
			"outscale_image_copy":                    resourceOutscaleOAPIImageCopy,
			"outscale_vpn_gateway_link":              resourceOutscaleOAPIVpnGatewayLink,
			"outscale_snapshot_attributes":           resourcedOutscaleOAPISnapshotAttributes,
			"outscale_lin_api_access":                resourceOutscaleOAPIVpcEndpoint,
			"outscale_net_api_access":                resourceOutscaleOAPIVpcEndpoint,
			"outscale_image_register":                resourceOutscaleOAPIImageRegister,
			"outscale_snapshot_copy":                 resourcedOutscaleOAPISnapshotCopy,
			"outscale_policy":                        resourceOutscaleOAPIPolicy,
			"outscale_directlink":                    resourceOutscaleOAPIDirectLink,
			"outscale_group":                         resourceOutscaleOAPIGroup,
			"outscale_group_user":                    resourceOutscaleOAPIGroupUser,
			"outscale_user":                          resourceOutscaleOAPIUser,
			"outscale_policy_default_version":        resourceOutscaleOAPIPolicyDefaultVersion,
			"outscale_policy_user_link":              resourceOutscaleOAPIPolicyUserLink,
			"outscale_server_certificate":            resourceOutscaleOAPServerCertificate,
			"outscale_policy_user":                   resourceOutscaleOAPIUserPolicy,
			"outscale_policy_group":                  resourceOutscaleOAPIPolicyGroup,
			"outscale_policy_group_link":             resourceOutscaleOAPIPolicyGroupLink,
			"outscale_policy_version":                resourceOutscaleOAPIPolicyVersion,
			"outscale_user_api_keys":                 resourceOutscaleOAPIUserAPIKey,
			"outscale_directlink_interface":          resourceOutscaleOAPIDirectLinkInterface,
			"outscale_dhcp_option":                   resourceOutscaleDHCPOption,     //TODO: OAPI
			"outscale_dhcp_option_link":              resourceOutscaleDHCPOptionLink, //TODO: OAPI
			"outscale_lin_peering_acceptation":       resourceOutscaleOAPILinPeeringConnectionAccepter,
			"outscale_net_peering_acceptation":       resourceOutscaleOAPILinPeeringConnectionAccepter,
			"outscale_snapshot_import":               resourcedOutscaleOAPISnapshotImport,
		},
		"icu": ResourceMap{
			"outscale_api_key": resourceOutscaleIamAccessKey,
		},
		"dl": ResourceMap{
			"outscale_directlink":           resourceOutscaleDirectLink,
			"outscale_directlink_interface": resourceOutscaleDirectLinkInterface,
		},
		"lbu": ResourceMap{
			"outscale_load_balancer":                 resourceOutscaleLoadBalancer,
			"outscale_load_balancer_cookiepolicy":    resourceOutscaleAppCookieStickinessPolicy,
			"outscale_load_balancer_vms":             resourceOutscaleLBUAttachment,
			"outscale_load_balancer_listeners":       resourceOutscaleLoadBalancerListeners,
			"outscale_load_balancer_attributes":      resourceOutscaleLoadBalancerAttributes,
			"outscale_load_balancer_tags":            resourceOutscaleLBUTags,
			"outscale_load_balancer_health_check":    resourceOutscaleLoadBalancerHealthCheck,
			"outscale_load_balancer_policy":          resourceOutscaleLoadBalancerPolicy,
			"outscale_load_balancer_ssl_certificate": resourceOutscaleLoadBalancerSSLCertificate,
		},
		"eim": ResourceMap{
			"outscale_policy":                 resourceOutscalePolicy,
			"outscale_group":                  resourceOutscaleGroup,
			"outscale_group_user":             resourceOutscaleGroupUser,
			"outscale_user":                   resourceOutscaleUser,
			"outscale_policy_default_version": resourceOutscalePolicyDefaultVersion,
			"outscale_policy_user_link":       resourceOutscalePolicyUserLink,
			"outscale_server_certificate":     resourceOutscaleEIMServerCertificate,
			"outscale_policy_user":            resourceOutscaleUserPolicy,
			"outscale_policy_group":           resourceOutscalePolicyGroup,
			"outscale_policy_group_link":      resourceOutscalePolicyGroupLink,
			"outscale_policy_version":         resourceOutscalePolicyVersion,
			"outscale_user_api_keys":          resourceOutscaleUserAPIKeys,
		},
	}
	datasources = Dictionary{
		"fcu": ResourceMap{
			"outscale_vm":                      dataSourceOutscaleVM,
			"outscale_vms":                     dataSourceOutscaleVMS,
			"outscale_firewall_rules_set":      dataSourceOutscaleFirewallRuleSet,
			"outscale_firewall_rules_sets":     dataSourceOutscaleFirewallRulesSets,
			"outscale_image":                   dataSourceOutscaleImage,
			"outscale_images":                  dataSourceOutscaleImages,
			"outscale_tag":                     dataSourceOutscaleTag,
			"outscale_tags":                    dataSourceOutscaleTags,
			"outscale_public_ip":               dataSourceOutscalePublicIP,
			"outscale_public_ips":              dataSourceOutscalePublicIPS,
			"outscale_volume":                  datasourceOutscaleVolume,
			"outscale_volumes":                 datasourceOutscaleVolumes,
			"outscale_nat_service":             dataSourceOutscaleNatService,
			"outscale_nat_services":            dataSourceOutscaleNatServices,
			"outscale_keypair":                 datasourceOutscaleKeyPair,
			"outscale_keypairs":                datasourceOutscaleKeyPairs,
			"outscale_vm_state":                dataSourceOutscaleVMState,
			"outscale_vms_state":               dataSourceOutscaleVMSState,
			"outscale_lin_internet_gateway":    datasourceOutscaleLinInternetGateway,
			"outscale_lin_internet_gateways":   datasourceOutscaleLinInternetGateways,
			"outscale_subnet":                  dataSourceOutscaleSubnet,
			"outscale_subnets":                 dataSourceOutscaleSubnets,
			"outscale_lin":                     dataSourceOutscaleVpc,
			"outscale_lins":                    dataSourceOutscaleVpcs,
			"outscale_lin_attributes":          dataSourceOutscaleVpcAttr,
			"outscale_client_endpoint":         dataSourceOutscaleCustomerGateway,
			"outscale_client_endpoints":        dataSourceOutscaleCustomerGateways,
			"outscale_route_table":             dataSourceOutscaleRouteTable,
			"outscale_route_tables":            dataSourceOutscaleRouteTables,
			"outscale_vpn_gateway":             dataSourceOutscaleVpnGateway,
			"outscale_api_key":                 dataSourceOutscaleIamAccessKey,
			"outscale_vpn_gateways":            dataSourceOutscaleVpnGateways,
			"outscale_vpn_connection":          dataSourceOutscaleVpnConnection,
			"outscale_sub_region":              dataSourceOutscaleAvailabilityZone,
			"outscale_prefix_list":             dataSourceOutscalePrefixList,
			"outscale_quota":                   dataSourceOutscaleQuota,
			"outscale_quotas":                  dataSourceOutscaleQuotas,
			"outscale_prefix_lists":            dataSourceOutscalePrefixLists,
			"outscale_region":                  dataSourceOutscaleRegion,
			"outscale_sub_regions":             dataSourceOutscaleAvailabilityZones,
			"outscale_regions":                 dataSourceOutscaleRegions,
			"outscale_vpn_connections":         dataSourceOutscaleVpnConnections,
			"outscale_product_types":           dataSourceOutscaleProductTypes,
			"outscale_reserved_vms":            dataSourceOutscaleReservedVMS,
			"outscale_vm_type":                 dataSourceOutscaleVMType,
			"outscale_vm_types":                dataSourceOutscaleVMTypes,
			"outscale_reserved_vms_offers":     dataSourceOutscaleReservedVMOffers,
			"outscale_reserved_vms_offer":      dataSourceOutscaleReservedVMOffer,
			"outscale_snapshot":                dataSourceOutscaleSnapshot,
			"outscale_snapshots":               dataSourceOutscaleSnapshots,
			"outscale_lin_peering":             dataSourceOutscaleLinPeeringConnection,
			"outscale_lin_peerings":            dataSourceOutscaleLinPeeringsConnection,
			"outscale_nics":                    dataSourceOutscaleNics,
			"outscale_nic":                     dataSourceOutscaleNic,
			"outscale_lin_api_access":          dataSourceOutscaleVpcEndpoint,
			"outscale_lin_api_accesses":        dataSourceOutscaleVpcEndpoints,
			"outscale_lin_api_access_services": dataSourceOutscaleVpcEndpointServices,
		},
		"oapi": ResourceMap{
			"outscale_vm":                                  dataSourceOutscaleOAPIVM,
			"outscale_vms":                                 datasourceOutscaleOApiVMS,
			"outscale_firewall_rules_sets":                 dataSourceOutscaleOAPIFirewallRulesSets,
			"outscale_images":                              dataSourceOutscaleOAPIImages,
			"outscale_firewall_rules_set":                  dataSourceOutscaleOAPIFirewallRuleSet,
			"outscale_tag":                                 dataSourceOutscaleOAPITag,
			"outscale_tags":                                dataSourceOutscaleOAPITags,
			"outscale_volume":                              datasourceOutscaleOAPIVolume,
			"outscale_volumes":                             datasourceOutscaleOAPIVolumes,
			"outscale_keypair":                             datasourceOutscaleOAPIKeyPair,
			"outscale_keypairs":                            datasourceOutscaleOAPIKeyPairs,
			"outscale_lin_internet_gateway":                datasourceOutscaleOAPIInternetService,
			"outscale_lin_internet_gateways":               datasourceOutscaleOAPIInternetServices,
			"outscale_internet_service":                    datasourceOutscaleOAPIInternetService,
			"outscale_internet_services":                   datasourceOutscaleOAPIInternetServices,
			"outscale_subnet":                              dataSourceOutscaleOAPISubnet,
			"outscale_subnets":                             dataSourceOutscaleOAPISubnets,
			"outscale_vm_state":                            dataSourceOutscaleOAPIVMState,
			"outscale_vms_state":                           dataSourceOutscaleOAPIVMSState,
			"outscale_lin":                                 dataSourceOutscaleOAPIVpc,
			"outscale_lins":                                dataSourceOutscaleOAPIVpcs,
			"outscale_lin_attributes":                      dataSourceOutscaleOAPIVpcAttr,
			"outscale_net":                                 dataSourceOutscaleOAPIVpc,
			"outscale_nets":                                dataSourceOutscaleOAPIVpcs,
			"outscale_net_attributes":                      dataSourceOutscaleOAPIVpcAttr,
			"outscale_client_endpoint":                     dataSourceOutscaleOAPICustomerGateway,
			"outscale_client_endpoints":                    dataSourceOutscaleOAPICustomerGateways,
			"outscale_route_table":                         dataSourceOutscaleOAPIRouteTable,
			"outscale_route_tables":                        dataSourceOutscaleOAPIRouteTables,
			"outscale_snapshot":                            dataSourceOutscaleOAPISnapshot,
			"outscale_snapshots":                           dataSourceOutscaleOAPISnapshots,
			"outscale_lin_peering":                         dataSourceOutscaleOAPILinPeeringConnection,
			"outscale_lin_peerings":                        dataSourceOutscaleOAPILinPeeringsConnection,
			"outscale_net_peering":                         dataSourceOutscaleOAPILinPeeringConnection,
			"outscale_net_peerings":                        dataSourceOutscaleOAPILinPeeringsConnection,
			"outscale_load_balancer":                       dataSourceOutscaleOAPILoadBalancer,
			"outscale_load_balancers":                      dataSourceOutscaleOAPILoadBalancers,
			"outscale_nic":                                 dataSourceOutscaleOAPINic,
			"outscale_nics":                                dataSourceOutscaleOAPINics,
			"outscale_load_balancer_tags":                  dataSourceOutscaleOAPILBUTags,
			"outscale_vm_health":                           dataSourceOutscaleOAPIVMHealth,
			"outscale_load_balancer_health_check":          dataSourceOutscaleOAPILoadBalancerHealthCheck,
			"outscale_load_balancer_listener_description":  dataSourceOutscaleOAPILoadBalancerLD,
			"outscale_load_balancer_listener_descriptions": dataSourceOutscaleOAPILoadBalancerLDs,
			"outscale_load_balancer_attributes":            dataSourceOutscaleOAPILoadBalancerAttr,
			"outscale_regions":                             dataSourceOutscaleOAPIRegions,
			"outscale_region":                              dataSourceOutscaleOAPIRegion,
			"outscale_reserved_vms_offer":                  dataSourceOutscaleOAPIReservedVMOffer,
			"outscale_reserved_vms_offers":                 dataSourceOutscaleOAPIReservedVMOffers,
			"outscale_reserved_vms":                        dataSourceOutscaleOAPIReservedVMS,
			"outscale_vpn_gateways":                        dataSourceOutscaleOAPIVpnGateways,
			"outscale_vpn_gateway":                         dataSourceOutscaleOAPIVpnGateway,
			"outscale_vm_type":                             dataSourceOutscaleOAPIVMType,
			"outscale_vm_types":                            dataSourceOutscaleOAPIVMTypes,
			"outscale_quotas":                              dataSourceOutscaleOAPIQuotas,
			"outscale_quota":                               dataSourceOutscaleOAPIQuota,
			"outscale_prefix_lists":                        dataSourceOutscaleOAPIPrefixLists,
			"outscale_prefix_list":                         dataSourceOutscaleOAPIPrefixList,
			"outscale_vpn_connections":                     dataSourceOutscaleOAPIVpnConnections,
			"outscale_sub_region":                          dataSourceOutscaleOAPIAvailabilityZone,
			"outscale_product_types":                       dataSourceOutscaleOAPIProductTypes,
			"outscale_image":                               dataSourceOutscaleOAPIImage,
			"outscale_lin_api_access_services":             dataSourceOutscaleOAPIVpcEndpointServices,
			"outscale_net_api_access_services":             dataSourceOutscaleOAPIVpcEndpointServices,
			"outscale_group":                               dataSourceOutscaleOAPIGroup,
			"outscale_user":                                dataSourceOutscaleOAPIUser,
			"outscale_users":                               dataSourceOutscaleOAPIUsers,
			"outscale_policy_user_link":                    dataSourceOutscaleOAPIPolicyUserLink,
			"outscale_groups":                              dataSourceOutscaleOAPIGroups,
			"outscale_groups_for_user":                     dataSourceOutscaleOAPIGroupUser,
			"outscale_policy":                              dataSourceOutscaleOAPIPolicy,
			"outscale_server_certificate":                  datasourceOutscaleOAPIEIMServerCertificate,
			"outscale_policy_group_link":                   dataSourceOutscaleOAPIPolicyGroupLink,
			"outscale_catalog":                             dataSourceOutscaleOAPICatalog,
			"outscale_public_catalog":                      dataSourceOutscaleOAPIPublicCatalog,
			"outscale_account_consumption":                 dataSourceOutscaleOAPIAccountConsumption,
			"outscale_policy_version":                      dataSourceOutscaleOAPIPolicyVersion,
			"outscale_policy_versions":                     dataSourceOutscaleOAPIPolicyVersions,
			"outscale_account":                             dataSourceOutscaleOAPIAccount,
			"outscale_directlink_vpn_gateways":             dataSourceOutscaleOAPIDLVPNGateways,
			"outscale_directlink_interface":                dataSourceOutscaleOAPIDirectLinkInterface,
			"outscale_directlink_interfaces":               dataSourceOutscaleOAPIDirectLinkInterfaces,
			"outscale_public_ip":                           dataSourceOutscaleOAPIPublicIP,
			"outscale_public_ips":                          dataSourceOutscaleOAPIPublicIPS,
			"outscale_nat_service":                         dataSourceOutscaleOAPINatService,
			"outscale_nat_services":                        dataSourceOutscaleOAPINatServices,
			"outscale_api_key":                             dataSourceOutscaleOAPIIamAccessKey,
			"outscale_vpn_connection":                      dataSourceOutscaleVpnConnection,     //TODO: OAPI
			"outscale_sub_regions":                         dataSourceOutscaleAvailabilityZones, //TODO: OAPI
			"outscale_lin_api_access":                      dataSourceOutscaleOAPIVpcEndpoint,
			"outscale_lin_api_accesses":                    dataSourceOutscaleOAPIVpcEndpoints,
			"outscale_net_api_access":                      dataSourceOutscaleOAPIVpcEndpoint,
			"outscale_net_api_accesses":                    dataSourceOutscaleOAPIVpcEndpoints,
		},
		"lbu": ResourceMap{
			"outscale_load_balancer":                       dataSourceOutscaleLoadBalancer,
			"outscale_load_balancers":                      dataSourceOutscaleLoadBalancers,
			"outscale_load_balancer_access_logs":           dataSourceOutscaleLoadBalancerAccessLogs,
			"outscale_load_balancer_tags":                  dataSourceOutscaleLBUTags,
			"outscale_vm_health":                           dataSourceOutscaleVMHealth,
			"outscale_load_balancer_health_check":          dataSourceOutscaleLoadBalancerHealthCheck,
			"outscale_load_balancer_listener_description":  dataSourceOutscaleLoadBalancerLD,
			"outscale_load_balancer_listener_descriptions": dataSourceOutscaleLoadBalancerLDs,
			"outscale_load_balancer_attributes":            dataSourceOutscaleLoadBalancerAttr,
		},
		"eim": ResourceMap{
			"outscale_group":               dataSourceOutscaleGroup,
			"outscale_user":                dataSourceOutscaleUser,
			"outscale_users":               dataSourceOutscaleUsers,
			"outscale_policy_user_link":    dataSourceOutscalePolicyUserLink,
			"outscale_groups":              dataSourceOutscaleGroups,
			"outscale_groups_for_user":     dataSourceOutscaleGroupUser,
			"outscale_policy":              dataSourceOutscalePolicy,
			"outscale_server_certificate":  datasourceOutscaleEIMServerCertificate,
			"outscale_server_certificates": datasourceOutscaleEIMServerCertificates,
			"outscale_policy_group_link":   dataSourceOutscalePolicyGroupLink,
			"outscale_user_api_keys":       dataSourceOutscaleUserAPIKeys,
			"outscale_policy_version":      dataSourceOutscalePolicyVersion,
			"outscale_policy_versions":     dataSourceOutscalePolicyVersions,
		},
		"icu": ResourceMap{
			"outscale_catalog":             dataSourceOutscaleCatalog,
			"outscale_public_catalog":      dataSourceOutscalePublicCatalog,
			"outscale_account_consumption": dataSourceOutscaleAccountConsumption,
			"outscale_account":             dataSourceOutscaleAccount,
		},
		"dl": ResourceMap{
			"outscale_sites":                   dataSourceOutscaleSites,
			"outscale_directlink_vpn_gateways": dataSourceOutscaleDLVPNGateways,
			"outscale_directlink":              dataSourceOutscaleDirectLink,
			"outscale_directlink_interface":    dataSourceOutscaleDirectLinkInterface,
			"outscale_directlink_interfaces":   dataSourceOutscaleDirectLinkInterfaces,
		},
	}
}

// GetResource ...
func GetResource(api, resource string) SchemaFunc {
	var a ResourceMap

	if _, ok := resources[api]; !ok {
		return nil
	}

	a = resources[api]

	if _, ok := a[resource]; !ok {
		return nil
	}
	return a[resource]
}

//GetDatasource receives the apu and the name of the datasource
//and returns the corrresponding
func GetDatasource(api, datasource string) SchemaFunc {
	var a ResourceMap
	if _, ok := datasources[api]; !ok {
		return nil
	}

	a = datasources[api]

	if _, ok := a[datasource]; !ok {
		return nil
	}
	return a[datasource]
}
