// Cloud Identity API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package cloudidentity

// LRO response metadata for InboundSamlSsoProfilesService.AddIdpCredential.
type AddIdpCredentialOperationMetadata struct {
	State string `json:"state,omitempty\"` // State of this Operation Will be "awaiting-multi-party-approval" when the operation is deferred due to the target customer having enabled [Multi-party approval for sensitive actions](https://support.google.com/a/answer/13790448).

}

// The request for creating an IdpCredential with its associated payload. An InboundSamlSsoProfile can own up to 2 credentials.
type AddIdpCredentialRequest struct {
	PemData string `json:"pemData,omitempty\"` // PEM encoded x509 certificate containing the public key for verifying IdP signatures.

}

// This resource object defines a domain that has been designated as allowlisted.
type AllowlistedDomain struct {
	Domain string `json:"domain,omitempty\"` // Required. Immutable. Name of the domain that is in the allowlist. e.g. "google.com"

	Name string `json:"name,omitempty\"` // Output only. Identifier. Resource name of the domain in the allowlist e.g. "allowlistedDomains/0184mhaj1smlusv"

}

// Request to cancel sent invitation for target email in UserInvitation.
type CancelUserInvitationRequest struct {
}

// The response message for MembershipsService.CheckTransitiveMembership.
type CheckTransitiveMembershipResponse struct {
	HasMembership bool `json:"hasMembership,omitempty\"` // Response does not include the possible roles of a member since the behavior of this rpc is not all-or-nothing unlike the other rpcs. So, it may not be possible to list all the roles definitively, due to possible lack of authorization in some of the paths.

}

// Metadata for CreateGroup LRO.
type CreateGroupMetadata struct {
}

// LRO response metadata for InboundOidcSsoProfilesService.CreateInboundOidcSsoProfile.
type CreateInboundOidcSsoProfileOperationMetadata struct {
	State string `json:"state,omitempty\"` // State of this Operation Will be "awaiting-multi-party-approval" when the operation is deferred due to the target customer having enabled [Multi-party approval for sensitive actions](https://support.google.com/a/answer/13790448).

}

// LRO response metadata for InboundSamlSsoProfilesService.CreateInboundSamlSsoProfile.
type CreateInboundSamlSsoProfileOperationMetadata struct {
	State string `json:"state,omitempty\"` // State of this Operation Will be "awaiting-multi-party-approval" when the operation is deferred due to the target customer having enabled [Multi-party approval for sensitive actions](https://support.google.com/a/answer/13790448).

}

// LRO response metadata for InboundSsoAssignmentsService.CreateInboundSsoAssignment.
type CreateInboundSsoAssignmentOperationMetadata struct {
}

// Metadata for CreateMembership LRO.
type CreateMembershipMetadata struct {
}

// Metadata for DeleteGroup LRO.
type DeleteGroupMetadata struct {
}

// LRO response metadata for InboundSamlSsoProfilesService.DeleteIdpCredential.
type DeleteIdpCredentialOperationMetadata struct {
}

// LRO response metadata for InboundOidcSsoProfilesService.DeleteInboundOidcSsoProfile.
type DeleteInboundOidcSsoProfileOperationMetadata struct {
}

// LRO response metadata for InboundSamlSsoProfilesService.DeleteInboundSamlSsoProfile.
type DeleteInboundSamlSsoProfileOperationMetadata struct {
}

// LRO response metadata for InboundSsoAssignmentsService.DeleteInboundSsoAssignment.
type DeleteInboundSsoAssignmentOperationMetadata struct {
}

// Metadata for DeleteMembership LRO.
type DeleteMembershipMetadata struct {
}

// Information of a DSA public key.
type DsaPublicKeyInfo struct {
	KeySize int `json:"keySize,omitempty\"` // Key size in bits (size of parameter P).

}

// Dynamic group metadata like queries and status.
type DynamicGroupMetadata struct {
	Queries []DynamicGroupQuery `json:"queries,omitempty\"` // Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Customers can create up to 500 dynamic groups.

	Status DynamicGroupStatus `json:"status,omitempty\"` // Output only. Status of the dynamic group.

}

// Defines a query on a resource.
type DynamicGroupQuery struct {
	Query string `json:"query,omitempty\"` // Query that determines the memberships of the dynamic group. Examples: All users with at least one `organizations.department` of engineering. `user.organizations.exists(org, org.department=='engineering')` All users with at least one location that has `area` of `foo` and `building_id` of `bar`. `user.locations.exists(loc, loc.area=='foo' && loc.building_id=='bar')` All users with any variation of the name John Doe (case-insensitive queries add `equalsIgnoreCase()` to the value being queried). `user.name.value.equalsIgnoreCase('jOhn DoE')`

	ResourceType string `json:"resourceType,omitempty\"` // Resource type for the Dynamic Group Query

}

// The current status of a dynamic group along with timestamp.
type DynamicGroupStatus struct {
	Status string `json:"status,omitempty\"` // Status of the dynamic group.

	StatusTime string `json:"statusTime,omitempty\"` // The latest time at which the dynamic group is guaranteed to be in the given status. If status is `UP_TO_DATE`, the latest time at which the dynamic group was confirmed to be up-to-date. If status is `UPDATING_MEMBERSHIPS`, the time at which dynamic group was created.

}

// A unique identifier for an entity in the Cloud Identity Groups API. An entity can represent either a group with an optional `namespace` or a user without a `namespace`. The combination of `id` and `namespace` must be unique; however, the same `id` can be used with different `namespace`s.
type EntityKey struct {
	Id string `json:"id,omitempty\"` // The ID of the entity. For Google-managed entities, the `id` should be the email address of an existing group or user. Email addresses need to adhere to [name guidelines for users and groups](https://support.google.com/a/answer/9193374). For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.

	Namespace string `json:"namespace,omitempty\"` // The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source}`.

}

// The `MembershipRole` expiry details.
type ExpiryDetail struct {
	ExpireTime string `json:"expireTime,omitempty\"` // The time at which the `MembershipRole` will expire.

}

// Metadata of GetMembershipGraphResponse LRO. This is currently empty to permit future extensibility.
type GetMembershipGraphMetadata struct {
}

// The response message for MembershipsService.GetMembershipGraph.
type GetMembershipGraphResponse struct {
	AdjacencyList []MembershipAdjacencyList `json:"adjacencyList,omitempty\"` // The membership graph's path information represented as an adjacency list.

	Groups []Group `json:"groups,omitempty\"` // The resources representing each group in the adjacency list. Each group in this list can be correlated to a 'group' of the MembershipAdjacencyList using the 'name' of the Group resource.

}

// Resource representing the Android specific attributes of a Device.
type GoogleAppsCloudidentityDevicesV1AndroidAttributes struct {
	CtsProfileMatch bool `json:"ctsProfileMatch,omitempty\"` // Whether the device passes Android CTS compliance.

	EnabledUnknownSources bool `json:"enabledUnknownSources,omitempty\"` // Whether applications from unknown sources can be installed on device.

	HasPotentiallyHarmfulApps bool `json:"hasPotentiallyHarmfulApps,omitempty\"` // Whether any potentially harmful apps were detected on the device.

	OwnerProfileAccount bool `json:"ownerProfileAccount,omitempty\"` // Whether this account is on an owner/primary profile. For phones, only true for owner profiles. Android 4+ devices can have secondary or restricted user profiles.

	OwnershipPrivilege string `json:"ownershipPrivilege,omitempty\"` // Ownership privileges on device.

	SupportsWorkProfile bool `json:"supportsWorkProfile,omitempty\"` // Whether device supports Android work profiles. If false, this service will not block access to corp data even if an administrator turns on the "Enforce Work Profile" policy.

	VerifiedBoot bool `json:"verifiedBoot,omitempty\"` // Whether Android verified boot status is GREEN.

	VerifyAppsEnabled bool `json:"verifyAppsEnabled,omitempty\"` // Whether Google Play Protect Verify Apps is enabled.

}

// Metadata for ApproveDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1ApproveDeviceUserMetadata struct {
}

// Request message for approving the device to access user data.
type GoogleAppsCloudidentityDevicesV1ApproveDeviceUserRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

}

// Response message for approving the device to access user data.
type GoogleAppsCloudidentityDevicesV1ApproveDeviceUserResponse struct {
	DeviceUser GoogleAppsCloudidentityDevicesV1DeviceUser `json:"deviceUser,omitempty\"` // Resultant DeviceUser object for the action.

}

// Metadata for BlockDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1BlockDeviceUserMetadata struct {
}

// Request message for blocking account on device.
type GoogleAppsCloudidentityDevicesV1BlockDeviceUserRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

}

// Response message for blocking the device from accessing user data.
type GoogleAppsCloudidentityDevicesV1BlockDeviceUserResponse struct {
	DeviceUser GoogleAppsCloudidentityDevicesV1DeviceUser `json:"deviceUser,omitempty\"` // Resultant DeviceUser object for the action.

}

// Contains information about browser profiles reported by the [Endpoint Verification extension](https://chromewebstore.google.com/detail/endpoint-verification/callobklhcbilhphinckomhgkigmfocg?pli=1).
type GoogleAppsCloudidentityDevicesV1BrowserAttributes struct {
	ChromeBrowserInfo GoogleAppsCloudidentityDevicesV1BrowserInfo `json:"chromeBrowserInfo,omitempty\"` // Represents the current state of the [Chrome browser attributes](https://cloud.google.com/access-context-manager/docs/browser-attributes) sent by the [Endpoint Verification extension](https://chromewebstore.google.com/detail/endpoint-verification/callobklhcbilhphinckomhgkigmfocg?pli=1).

	ChromeProfileId string `json:"chromeProfileId,omitempty\"` // Chrome profile ID that is exposed by the Chrome API. It is unique for each device.

	LastProfileSyncTime string `json:"lastProfileSyncTime,omitempty\"` // Timestamp in milliseconds since the Unix epoch when the profile/gcm id was last synced.

}

// Browser-specific fields reported by the [Endpoint Verification extension](https://chromewebstore.google.com/detail/endpoint-verification/callobklhcbilhphinckomhgkigmfocg?pli=1).
type GoogleAppsCloudidentityDevicesV1BrowserInfo struct {
	BrowserManagementState string `json:"browserManagementState,omitempty\"` // Output only. Browser's management state.

	BrowserVersion string `json:"browserVersion,omitempty\"` // Version of the request initiating browser. E.g. `91.0.4442.4`.

	IsBuiltInDnsClientEnabled bool `json:"isBuiltInDnsClientEnabled,omitempty\"` // Current state of [built-in DNS client](https://chromeenterprise.google/policies/#BuiltInDnsClientEnabled).

	IsBulkDataEntryAnalysisEnabled bool `json:"isBulkDataEntryAnalysisEnabled,omitempty\"` // Current state of [bulk data analysis](https://chromeenterprise.google/policies/#OnBulkDataEntryEnterpriseConnector). Set to true if provider list from Chrome is non-empty.

	IsChromeCleanupEnabled bool `json:"isChromeCleanupEnabled,omitempty\"` // Deprecated: This field is not used for Chrome version 118 and later. Current state of [Chrome Cleanup](https://chromeenterprise.google/policies/#ChromeCleanupEnabled).

	IsChromeRemoteDesktopAppBlocked bool `json:"isChromeRemoteDesktopAppBlocked,omitempty\"` // Current state of [Chrome Remote Desktop app](https://chromeenterprise.google/policies/#URLBlocklist).

	IsFileDownloadAnalysisEnabled bool `json:"isFileDownloadAnalysisEnabled,omitempty\"` // Current state of [file download analysis](https://chromeenterprise.google/policies/#OnFileDownloadedEnterpriseConnector). Set to true if provider list from Chrome is non-empty.

	IsFileUploadAnalysisEnabled bool `json:"isFileUploadAnalysisEnabled,omitempty\"` // Current state of [file upload analysis](https://chromeenterprise.google/policies/#OnFileAttachedEnterpriseConnector). Set to true if provider list from Chrome is non-empty.

	IsRealtimeUrlCheckEnabled bool `json:"isRealtimeUrlCheckEnabled,omitempty\"` // Current state of [real-time URL check](https://chromeenterprise.google/policies/#EnterpriseRealTimeUrlCheckMode). Set to true if provider list from Chrome is non-empty.

	IsSecurityEventAnalysisEnabled bool `json:"isSecurityEventAnalysisEnabled,omitempty\"` // Current state of [security event analysis](https://chromeenterprise.google/policies/#OnSecurityEventEnterpriseConnector). Set to true if provider list from Chrome is non-empty.

	IsSiteIsolationEnabled bool `json:"isSiteIsolationEnabled,omitempty\"` // Current state of [site isolation](https://chromeenterprise.google/policies/?policy=IsolateOrigins).

	IsThirdPartyBlockingEnabled bool `json:"isThirdPartyBlockingEnabled,omitempty\"` // Current state of [third-party blocking](https://chromeenterprise.google/policies/#ThirdPartyBlockingEnabled).

	PasswordProtectionWarningTrigger string `json:"passwordProtectionWarningTrigger,omitempty\"` // Current state of [password protection trigger](https://chromeenterprise.google/policies/#PasswordProtectionWarningTrigger).

	SafeBrowsingProtectionLevel string `json:"safeBrowsingProtectionLevel,omitempty\"` // Current state of [Safe Browsing protection level](https://chromeenterprise.google/policies/#SafeBrowsingProtectionLevel).

}

// Metadata for CancelWipeDevice LRO.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceMetadata struct {
}

// Request message for cancelling an unfinished device wipe.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

}

// Response message for cancelling an unfinished device wipe.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceResponse struct {
	Device GoogleAppsCloudidentityDevicesV1Device `json:"device,omitempty\"` // Resultant Device object for the action. Note that asset tags will not be returned in the device object.

}

// Metadata for CancelWipeDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserMetadata struct {
}

// Request message for cancelling an unfinished user account wipe.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

}

// Response message for cancelling an unfinished user account wipe.
type GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserResponse struct {
	DeviceUser GoogleAppsCloudidentityDevicesV1DeviceUser `json:"deviceUser,omitempty\"` // Resultant DeviceUser object for the action.

}

// Stores information about a certificate.
type GoogleAppsCloudidentityDevicesV1CertificateAttributes struct {
	CertificateTemplate GoogleAppsCloudidentityDevicesV1CertificateTemplate `json:"certificateTemplate,omitempty\"` // The X.509 extension for CertificateTemplate.

	Fingerprint string `json:"fingerprint,omitempty\"` // The encoded certificate fingerprint.

	Issuer string `json:"issuer,omitempty\"` // The name of the issuer of this certificate.

	SerialNumber string `json:"serialNumber,omitempty\"` // Serial number of the certificate, Example: "123456789".

	Subject string `json:"subject,omitempty\"` // The subject name of this certificate.

	Thumbprint string `json:"thumbprint,omitempty\"` // The certificate thumbprint.

	ValidationState string `json:"validationState,omitempty\"` // Output only. Validation state of this certificate.

	ValidityExpirationTime string `json:"validityExpirationTime,omitempty\"` // Certificate not valid at or after this timestamp.

	ValidityStartTime string `json:"validityStartTime,omitempty\"` // Certificate not valid before this timestamp.

}

// CertificateTemplate (v3 Extension in X.509).
type GoogleAppsCloudidentityDevicesV1CertificateTemplate struct {
	Id string `json:"id,omitempty\"` // The template id of the template. Example: "1.3.6.1.4.1.311.21.8.15608621.11768144.5720724.16068415.6889630.81.2472537.7784047".

	MajorVersion int `json:"majorVersion,omitempty\"` // The Major version of the template. Example: 100.

	MinorVersion int `json:"minorVersion,omitempty\"` // The minor version of the template. Example: 12.

}

// Represents the state associated with an API client calling the Devices API. Resource representing ClientState and supports updates from API users
type GoogleAppsCloudidentityDevicesV1ClientState struct {
	AssetTags []string `json:"assetTags,omitempty\"` // The caller can specify asset tags for this resource

	ComplianceState string `json:"complianceState,omitempty\"` // The compliance state of the resource as specified by the API client.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time the client state data was created.

	CustomId string `json:"customId,omitempty\"` // This field may be used to store a unique identifier for the API resource within which these CustomAttributes are a field.

	Etag string `json:"etag,omitempty\"` // The token that needs to be passed back for concurrency control in updates. Token needs to be passed back in UpdateRequest

	HealthScore string `json:"healthScore,omitempty\"` // The Health score of the resource. The Health score is the callers specification of the condition of the device from a usability point of view. For example, a third-party device management provider may specify a health score based on its compliance with organizational policies.

	KeyValuePairs map[string]interface{} `json:"keyValuePairs,omitempty\"` // The map of key-value attributes stored by callers specific to a device. The total serialized length of this map may not exceed 10KB. No limit is placed on the number of attributes in a map.

	LastUpdateTime string `json:"lastUpdateTime,omitempty\"` // Output only. The time the client state data was last updated.

	Managed string `json:"managed,omitempty\"` // The management state of the resource as specified by the API client.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the ClientState in format: `devices/{device}/deviceUsers/{device_user}/clientState/{partner}`, where partner corresponds to the partner storing the data. For partners belonging to the "BeyondCorp Alliance", this is the partner ID specified to you by Google. For all other callers, this is a string of the form: `{customer}-suffix`, where `customer` is your customer ID. The *suffix* is any string the caller specifies. This string will be displayed verbatim in the administration console. This suffix is used in setting up Custom Access Levels in Context-Aware Access. Your organization's customer ID can be obtained from the URL: `GET https://www.googleapis.com/admin/directory/v1/customers/my_customer` The `id` field in the response contains the customer ID starting with the letter 'C'. The customer ID to be used in this API is the string after the letter 'C' (not including 'C')

	OwnerType string `json:"ownerType,omitempty\"` // Output only. The owner of the ClientState

	ScoreReason string `json:"scoreReason,omitempty\"` // A descriptive cause of the health score.

}

// Metadata for CreateDevice LRO.
type GoogleAppsCloudidentityDevicesV1CreateDeviceMetadata struct {
}

// Additional custom attribute values may be one of these types
type GoogleAppsCloudidentityDevicesV1CustomAttributeValue struct {
	BoolValue bool `json:"boolValue,omitempty\"` // Represents a boolean value.

	NumberValue float64 `json:"numberValue,omitempty\"` // Represents a double value.

	StringValue string `json:"stringValue,omitempty\"` // Represents a string value.

}

// Metadata for DeleteDevice LRO.
type GoogleAppsCloudidentityDevicesV1DeleteDeviceMetadata struct {
}

// Metadata for DeleteDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1DeleteDeviceUserMetadata struct {
}

// A Device within the Cloud Identity Devices API. Represents a Device known to Google Cloud, independent of the device ownership, type, and whether it is assigned or in use by a user.
type GoogleAppsCloudidentityDevicesV1Device struct {
	AndroidSpecificAttributes GoogleAppsCloudidentityDevicesV1AndroidAttributes `json:"androidSpecificAttributes,omitempty\"` // Output only. Attributes specific to Android devices.

	AssetTag string `json:"assetTag,omitempty\"` // Asset tag of the device.

	BasebandVersion string `json:"basebandVersion,omitempty\"` // Output only. Baseband version of the device.

	BootloaderVersion string `json:"bootloaderVersion,omitempty\"` // Output only. Device bootloader version. Example: 0.6.7.

	Brand string `json:"brand,omitempty\"` // Output only. Device brand. Example: Samsung.

	BuildNumber string `json:"buildNumber,omitempty\"` // Output only. Build number of the device.

	CompromisedState string `json:"compromisedState,omitempty\"` // Output only. Represents whether the Device is compromised.

	CreateTime string `json:"createTime,omitempty\"` // Output only. When the Company-Owned device was imported. This field is empty for BYOD devices.

	DeviceId string `json:"deviceId,omitempty\"` // Unique identifier for the device.

	DeviceType string `json:"deviceType,omitempty\"` // Output only. Type of device.

	EnabledDeveloperOptions bool `json:"enabledDeveloperOptions,omitempty\"` // Output only. Whether developer options is enabled on device.

	EnabledUsbDebugging bool `json:"enabledUsbDebugging,omitempty\"` // Output only. Whether USB debugging is enabled on device.

	EncryptionState string `json:"encryptionState,omitempty\"` // Output only. Device encryption state.

	EndpointVerificationSpecificAttributes GoogleAppsCloudidentityDevicesV1EndpointVerificationSpecificAttributes `json:"endpointVerificationSpecificAttributes,omitempty\"` // Output only. Attributes specific to [Endpoint Verification](https://cloud.google.com/endpoint-verification/docs/overview) devices.

	Hostname string `json:"hostname,omitempty\"` // Host name of the device.

	Imei string `json:"imei,omitempty\"` // Output only. IMEI number of device if GSM device; empty otherwise.

	KernelVersion string `json:"kernelVersion,omitempty\"` // Output only. Kernel version of the device.

	LastSyncTime string `json:"lastSyncTime,omitempty\"` // Most recent time when device synced with this service.

	ManagementState string `json:"managementState,omitempty\"` // Output only. Management state of the device

	Manufacturer string `json:"manufacturer,omitempty\"` // Output only. Device manufacturer. Example: Motorola.

	Meid string `json:"meid,omitempty\"` // Output only. MEID number of device if CDMA device; empty otherwise.

	Model string `json:"model,omitempty\"` // Output only. Model name of device. Example: Pixel 3.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device}`, where device is the unique id assigned to the Device. Important: Device API scopes require that you use domain-wide delegation to access the API. For more information, see [Set up the Devices API](https://cloud.google.com/identity/docs/how-to/setup-devices).

	NetworkOperator string `json:"networkOperator,omitempty\"` // Output only. Mobile or network operator of device, if available.

	OsVersion string `json:"osVersion,omitempty\"` // Output only. OS version of the device. Example: Android 8.1.0.

	OtherAccounts []string `json:"otherAccounts,omitempty\"` // Output only. Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".

	OwnerType string `json:"ownerType,omitempty\"` // Output only. Whether the device is owned by the company or an individual

	ReleaseVersion string `json:"releaseVersion,omitempty\"` // Output only. OS release version. Example: 6.0.

	SecurityPatchTime string `json:"securityPatchTime,omitempty\"` // Output only. OS security patch update time on device.

	SerialNumber string `json:"serialNumber,omitempty\"` // Serial Number of device. Example: HT82V1A01076.

	UnifiedDeviceId string `json:"unifiedDeviceId,omitempty\"` // Output only. Unified device id of the device.

	WifiMacAddresses []string `json:"wifiMacAddresses,omitempty\"` // WiFi MAC addresses of device.

}

// Represents a user's use of a Device in the Cloud Identity Devices API. A DeviceUser is a resource representing a user's use of a Device
type GoogleAppsCloudidentityDevicesV1DeviceUser struct {
	CompromisedState string `json:"compromisedState,omitempty\"` // Compromised State of the DeviceUser object

	CreateTime string `json:"createTime,omitempty\"` // When the user first signed in to the device

	FirstSyncTime string `json:"firstSyncTime,omitempty\"` // Output only. Most recent time when user registered with this service.

	LanguageCode string `json:"languageCode,omitempty\"` // Output only. Default locale used on device, in IETF BCP-47 format.

	LastSyncTime string `json:"lastSyncTime,omitempty\"` // Output only. Last time when user synced with policies.

	ManagementState string `json:"managementState,omitempty\"` // Output only. Management state of the user on the device.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the DeviceUser in format: `devices/{device}/deviceUsers/{device_user}`, where `device_user` uniquely identifies a user's use of a device.

	PasswordState string `json:"passwordState,omitempty\"` // Password state of the DeviceUser object

	UserAgent string `json:"userAgent,omitempty\"` // Output only. User agent on the device for this specific user

	UserEmail string `json:"userEmail,omitempty\"` // Email address of the user registered on the device.

}

// Resource representing the [Endpoint Verification-specific attributes](https://cloud.google.com/endpoint-verification/docs/device-information) of a device.
type GoogleAppsCloudidentityDevicesV1EndpointVerificationSpecificAttributes struct {
	AdditionalSignals map[string]interface{} `json:"additionalSignals,omitempty\"` // [Additional signals](https://cloud.google.com/endpoint-verification/docs/device-information) reported by Endpoint Verification. It includes the following attributes: * Non-configurable attributes: hotfixes, av_installed, av_enabled, windows_domain_name, is_os_native_firewall_enabled, and is_secure_boot_enabled. * [Configurable attributes](https://cloud.google.com/endpoint-verification/docs/collect-config-attributes): file, folder, and binary attributes; registry entries; and properties in a plist.

	BrowserAttributes []GoogleAppsCloudidentityDevicesV1BrowserAttributes `json:"browserAttributes,omitempty\"` // Details of browser profiles reported by Endpoint Verification.

	CertificateAttributes []GoogleAppsCloudidentityDevicesV1CertificateAttributes `json:"certificateAttributes,omitempty\"` // Details of certificates.

}

// Response message that is returned in ListClientStates.
type GoogleAppsCloudidentityDevicesV1ListClientStatesResponse struct {
	ClientStates []GoogleAppsCloudidentityDevicesV1ClientState `json:"clientStates,omitempty\"` // Client states meeting the list restrictions.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results. Empty if there are no more results.

}

// Response message that is returned from the ListDeviceUsers method.
type GoogleAppsCloudidentityDevicesV1ListDeviceUsersResponse struct {
	DeviceUsers []GoogleAppsCloudidentityDevicesV1DeviceUser `json:"deviceUsers,omitempty\"` // Devices meeting the list restrictions.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results. Empty if there are no more results.

}

// Response message that is returned from the ListDevices method.
type GoogleAppsCloudidentityDevicesV1ListDevicesResponse struct {
	Devices []GoogleAppsCloudidentityDevicesV1Device `json:"devices,omitempty\"` // Devices meeting the list restrictions.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results. Empty if there are no more results.

}

// Metadata for ListEndpointApps LRO.
type GoogleAppsCloudidentityDevicesV1ListEndpointAppsMetadata struct {
}

// Response containing resource names of the DeviceUsers associated with the caller's credentials.
type GoogleAppsCloudidentityDevicesV1LookupSelfDeviceUsersResponse struct {
	Customer string `json:"customer,omitempty\"` // The customer resource name that may be passed back to other Devices API methods such as List, Get, etc.

	Names []string `json:"names,omitempty\"` // [Resource names](https://cloud.google.com/apis/design/resource_names) of the DeviceUsers in the format: `devices/{device}/deviceUsers/{user_resource}`, where device is the unique ID assigned to a Device and user_resource is the unique user ID

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results. Empty if there are no more results.

}

// Metadata for SignoutDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1SignoutDeviceUserMetadata struct {
}

// Metadata for UpdateClientState LRO.
type GoogleAppsCloudidentityDevicesV1UpdateClientStateMetadata struct {
}

// Metadata for UpdateDevice LRO.
type GoogleAppsCloudidentityDevicesV1UpdateDeviceMetadata struct {
}

// Metadata for WipeDevice LRO.
type GoogleAppsCloudidentityDevicesV1WipeDeviceMetadata struct {
}

// Request message for wiping all data on the device.
type GoogleAppsCloudidentityDevicesV1WipeDeviceRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

	RemoveResetLock bool `json:"removeResetLock,omitempty\"` // Optional. Specifies if a user is able to factory reset a device after a Device Wipe. On iOS, this is called "Activation Lock", while on Android, this is known as "Factory Reset Protection". If true, this protection will be removed from the device, so that a user can successfully factory reset. If false, the setting is untouched on the device.

}

// Response message for wiping all data on the device.
type GoogleAppsCloudidentityDevicesV1WipeDeviceResponse struct {
	Device GoogleAppsCloudidentityDevicesV1Device `json:"device,omitempty\"` // Resultant Device object for the action. Note that asset tags will not be returned in the device object.

}

// Metadata for WipeDeviceUser LRO.
type GoogleAppsCloudidentityDevicesV1WipeDeviceUserMetadata struct {
}

// Request message for starting an account wipe on device.
type GoogleAppsCloudidentityDevicesV1WipeDeviceUserRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.

}

// Response message for wiping the user's account from the device.
type GoogleAppsCloudidentityDevicesV1WipeDeviceUserResponse struct {
	DeviceUser GoogleAppsCloudidentityDevicesV1DeviceUser `json:"deviceUser,omitempty\"` // Resultant DeviceUser object for the action.

}

// A group within the Cloud Identity Groups API. A `Group` is a collection of entities, where each entity is either a user, another group, or a service account.
type Group struct {
	AdditionalGroupKeys []EntityKey `json:"additionalGroupKeys,omitempty\"` // Output only. Additional group keys associated with the Group.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time when the `Group` was created.

	Description string `json:"description,omitempty\"` // An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.

	DisplayName string `json:"displayName,omitempty\"` // The display name of the `Group`.

	DynamicGroupMetadata DynamicGroupMetadata `json:"dynamicGroupMetadata,omitempty\"` // Optional. Dynamic group metadata like queries and status.

	GroupKey EntityKey `json:"groupKey,omitempty\"` // Required. The `EntityKey` of the `Group`.

	Labels map[string]interface{} `json:"labels,omitempty\"` // Required. One or more label entries that apply to the Group. Labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Google Groups can be [locked](https://support.google.com/a?p=locked-groups). To lock a group, add a label with a key of `cloudidentity.googleapis.com/groups.locked` and an empty value. Doing so locks the group. To unlock the group, remove this label.

	Name string `json:"name,omitempty\"` // Output only. The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group}`.

	Parent string `json:"parent,omitempty\"` // Required. Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external [identity-mapped groups](https://support.google.com/a/answer/9039510) or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn'). [Find your customer ID.] (https://support.google.com/cloudidentity/answer/10070793)

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. The time when the `Group` was last updated.

}

// Message representing a transitive group of a user or a group.
type GroupRelation struct {
	DisplayName string `json:"displayName,omitempty\"` // Display name for this group.

	Group string `json:"group,omitempty\"` // Resource name for this group.

	GroupKey EntityKey `json:"groupKey,omitempty\"` // Entity key has an id and a namespace. In case of discussion forums, the id will be an email address without a namespace.

	Labels map[string]interface{} `json:"labels,omitempty\"` // Labels for Group resource.

	RelationType string `json:"relationType,omitempty\"` // The relation between the member and the transitive group.

	Roles []TransitiveMembershipRole `json:"roles,omitempty\"` // Membership roles of the member for the group.

}

// Credential for verifying signatures produced by the Identity Provider.
type IdpCredential struct {
	DsaKeyInfo DsaPublicKeyInfo `json:"dsaKeyInfo,omitempty\"` // Output only. Information of a DSA public key.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the credential.

	RsaKeyInfo RsaPublicKeyInfo `json:"rsaKeyInfo,omitempty\"` // Output only. Information of a RSA public key.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. Time when the `IdpCredential` was last updated.

}

// An [OIDC](https://openid.net/developers/how-connect-works/) federation between a Google enterprise customer and an OIDC identity provider.
type InboundOidcSsoProfile struct {
	Customer string `json:"customer,omitempty\"` // Immutable. The customer. For example: `customers/C0123abc`.

	DisplayName string `json:"displayName,omitempty\"` // Human-readable name of the OIDC SSO profile.

	IdpConfig OidcIdpConfig `json:"idpConfig,omitempty\"` // OIDC identity provider configuration.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the OIDC SSO profile.

	RpConfig OidcRpConfig `json:"rpConfig,omitempty\"` // OIDC relying party (RP) configuration for this OIDC SSO profile. These are the RP details provided by Google that should be configured on the corresponding identity provider.

}

// A [SAML 2.0](https://www.oasis-open.org/standards#samlv2.0) federation between a Google enterprise customer and a SAML identity provider.
type InboundSamlSsoProfile struct {
	Customer string `json:"customer,omitempty\"` // Immutable. The customer. For example: `customers/C0123abc`.

	DisplayName string `json:"displayName,omitempty\"` // Human-readable name of the SAML SSO profile.

	IdpConfig SamlIdpConfig `json:"idpConfig,omitempty\"` // SAML identity provider configuration.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the SAML SSO profile.

	SpConfig SamlSpConfig `json:"spConfig,omitempty\"` // SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.

}

// Targets with "set" SSO assignments and their respective assignments.
type InboundSsoAssignment struct {
	Customer string `json:"customer,omitempty\"` // Immutable. The customer. For example: `customers/C0123abc`.

	Name string `json:"name,omitempty\"` // Output only. [Resource name](https://cloud.google.com/apis/design/resource_names) of the Inbound SSO Assignment.

	OidcSsoInfo OidcSsoInfo `json:"oidcSsoInfo,omitempty\"` // OpenID Connect SSO details. Must be set if and only if `sso_mode` is set to `OIDC_SSO`.

	Rank int `json:"rank,omitempty\"` // Must be zero (which is the default value so it can be omitted) for assignments with `target_org_unit` set and must be greater-than-or-equal-to one for assignments with `target_group` set.

	SamlSsoInfo SamlSsoInfo `json:"samlSsoInfo,omitempty\"` // SAML SSO details. Must be set if and only if `sso_mode` is set to `SAML_SSO`.

	SignInBehavior SignInBehavior `json:"signInBehavior,omitempty\"` // Assertions about users assigned to an IdP will always be accepted from that IdP. This controls whether/when Google should redirect a user to the IdP. Unset (defaults) is the recommended configuration.

	SsoMode string `json:"ssoMode,omitempty\"` // Inbound SSO behavior.

	TargetGroup string `json:"targetGroup,omitempty\"` // Immutable. Must be of the form `groups/{group}`.

	TargetOrgUnit string `json:"targetOrgUnit,omitempty\"` // Immutable. Must be of the form `orgUnits/{org_unit}`.

}

// Response for IsInvitableUser RPC.
type IsInvitableUserResponse struct {
	IsInvitableUser bool `json:"isInvitableUser,omitempty\"` // Returns true if the email address is invitable.

}

// Response message for AllowlistedDomainsService.ListAllowlistedDomains.
type ListAllowlistedDomainsResponse struct {
	AllowlistedDomains []AllowlistedDomain `json:"allowlistedDomains,omitempty\"` // Contains the list of domains in the allowlist. There is no defined ordering of domains within a result.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Contains the next page token if the result is not exhaustive. If there are no more results, this token is empty.

}

// Response message for ListGroups operation.
type ListGroupsResponse struct {
	Groups []Group `json:"groups,omitempty\"` // Groups returned in response to list request. The results are not sorted.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results, or empty if there are no more results available for listing.

}

// Response of the InboundSamlSsoProfilesService.ListIdpCredentials method.
type ListIdpCredentialsResponse struct {
	IdpCredentials []IdpCredential `json:"idpCredentials,omitempty\"` // The IdpCredentials from the specified InboundSamlSsoProfile.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

}

// Response of the InboundOidcSsoProfilesService.ListInboundOidcSsoProfiles method.
type ListInboundOidcSsoProfilesResponse struct {
	InboundOidcSsoProfiles []InboundOidcSsoProfile `json:"inboundOidcSsoProfiles,omitempty\"` // List of InboundOidcSsoProfiles.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

}

// Response of the InboundSamlSsoProfilesService.ListInboundSamlSsoProfiles method.
type ListInboundSamlSsoProfilesResponse struct {
	InboundSamlSsoProfiles []InboundSamlSsoProfile `json:"inboundSamlSsoProfiles,omitempty\"` // List of InboundSamlSsoProfiles.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

}

// Response of the InboundSsoAssignmentsService.ListInboundSsoAssignments method.
type ListInboundSsoAssignmentsResponse struct {
	InboundSsoAssignments []InboundSsoAssignment `json:"inboundSsoAssignments,omitempty\"` // The assignments.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

}

// The response message for MembershipsService.ListMemberships.
type ListMembershipsResponse struct {
	Memberships []Membership `json:"memberships,omitempty\"` // The `Membership`s under the specified `parent`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A continuation token to retrieve the next page of results, or empty if there are no more results available.

}

// The response message for PoliciesService.ListPolicies.
type ListPoliciesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // The pagination token to retrieve the next page of results. If this field is empty, there are no subsequent pages.

	Policies []Policy `json:"policies,omitempty\"` // The results

}

// Response message for UserInvitation listing request.
type ListUserInvitationsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // The token for the next page. If not empty, indicates that there may be more `UserInvitation` resources that match the listing request; this value can be used in a subsequent ListUserInvitationsRequest to get continued results with the current list call.

	UserInvitations []UserInvitation `json:"userInvitations,omitempty\"` // The list of UserInvitation resources.

}

// The response message for GroupsService.LookupGroupName.
type LookupGroupNameResponse struct {
	Name string `json:"name,omitempty\"` // The [resource name](https://cloud.google.com/apis/design/resource_names) of the looked-up `Group`.

}

// The response message for MembershipsService.LookupMembershipName.
type LookupMembershipNameResponse struct {
	Name string `json:"name,omitempty\"` // The [resource name](https://cloud.google.com/apis/design/resource_names) of the looked-up `Membership`. Must be of the form `groups/{group}/memberships/{membership}`.

}

// Message representing a transitive membership of a group.
type MemberRelation struct {
	Member string `json:"member,omitempty\"` // Resource name for this member.

	PreferredMemberKey []EntityKey `json:"preferredMemberKey,omitempty\"` // Entity key has an id and a namespace. In case of discussion forums, the id will be an email address without a namespace.

	RelationType string `json:"relationType,omitempty\"` // The relation between the group and the transitive member.

	Roles []TransitiveMembershipRole `json:"roles,omitempty\"` // The membership role details (i.e name of role and expiry time).

}

// The definition of MemberRestriction
type MemberRestriction struct {
	Evaluation RestrictionEvaluation `json:"evaluation,omitempty\"` // The evaluated state of this restriction on a group.

	Query string `json:"query,omitempty\"` // Member Restriction as defined by CEL expression. Supported restrictions are: `member.customer_id` and `member.type`. Valid values for `member.type` are `1`, `2` and `3`. They correspond to USER, SERVICE_ACCOUNT, and GROUP respectively. The value for `member.customer_id` only supports `groupCustomerId()` currently which means the customer id of the group will be used for restriction. Supported operators are `&&`, `||` and `==`, corresponding to AND, OR, and EQUAL. Examples: Allow only service accounts of given customer to be members. `member.type == 2 && member.customer_id == groupCustomerId()` Allow only users or groups to be members. `member.type == 1 || member.type == 3`

}

// A membership within the Cloud Identity Groups API. A `Membership` defines a relationship between a `Group` and an entity belonging to that `Group`, referred to as a "member".
type Membership struct {
	CreateTime string `json:"createTime,omitempty\"` // Output only. The time when the `Membership` was created.

	DeliverySetting string `json:"deliverySetting,omitempty\"` // Output only. Delivery setting associated with the membership.

	Name string `json:"name,omitempty\"` // Output only. The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group}/memberships/{membership}`.

	PreferredMemberKey EntityKey `json:"preferredMemberKey,omitempty\"` // Required. Immutable. The `EntityKey` of the member.

	Roles []MembershipRole `json:"roles,omitempty\"` // The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.

	TypeValue string `json:"type,omitempty\"` // Output only. The type of the membership.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. The time when the `Membership` was last updated.

}

// Membership graph's path information as an adjacency list.
type MembershipAdjacencyList struct {
	Edges []Membership `json:"edges,omitempty\"` // Each edge contains information about the member that belongs to this group. Note: Fields returned here will help identify the specific Membership resource (e.g `name`, `preferred_member_key` and `role`), but may not be a comprehensive list of all fields.

	Group string `json:"group,omitempty\"` // Resource name of the group that the members belong to.

}

// Message containing membership relation.
type MembershipRelation struct {
	Description string `json:"description,omitempty\"` // An extended description to help users determine the purpose of a `Group`.

	DisplayName string `json:"displayName,omitempty\"` // The display name of the `Group`.

	Group string `json:"group,omitempty\"` // The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.

	GroupKey EntityKey `json:"groupKey,omitempty\"` // The `EntityKey` of the `Group`.

	Labels map[string]interface{} `json:"labels,omitempty\"` // One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.

	Membership string `json:"membership,omitempty\"` // The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group_id}/memberships/{membership_id}`.

	Roles []MembershipRole `json:"roles,omitempty\"` // The `MembershipRole`s that apply to the `Membership`.

}

// A membership role within the Cloud Identity Groups API. A `MembershipRole` defines the privileges granted to a `Membership`.
type MembershipRole struct {
	ExpiryDetail ExpiryDetail `json:"expiryDetail,omitempty\"` // The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.

	Name string `json:"name,omitempty\"` // The name of the `MembershipRole`. Must be one of `OWNER`, `MANAGER`, `MEMBER`.

	RestrictionEvaluations RestrictionEvaluations `json:"restrictionEvaluations,omitempty\"` // Evaluations of restrictions applied to parent group on this membership.

}

// The evaluated state of this restriction.
type MembershipRoleRestrictionEvaluation struct {
	State string `json:"state,omitempty\"` // Output only. The current state of the restriction

}

// The request message for MembershipsService.ModifyMembershipRoles.
type ModifyMembershipRolesRequest struct {
	AddRoles []MembershipRole `json:"addRoles,omitempty\"` // The `MembershipRole`s to be added. Adding or removing roles in the same request as updating roles is not supported. Must not be set if `update_roles_params` is set.

	RemoveRoles []string `json:"removeRoles,omitempty\"` // The `name`s of the `MembershipRole`s to be removed. Adding or removing roles in the same request as updating roles is not supported. It is not possible to remove the `MEMBER` `MembershipRole`. If you wish to delete a `Membership`, call MembershipsService.DeleteMembership instead. Must not contain `MEMBER`. Must not be set if `update_roles_params` is set.

	UpdateRolesParams []UpdateMembershipRolesParams `json:"updateRolesParams,omitempty\"` // The `MembershipRole`s to be updated. Updating roles in the same request as adding or removing roles is not supported. Must not be set if either `add_roles` or `remove_roles` is set.

}

// The response message for MembershipsService.ModifyMembershipRoles.
type ModifyMembershipRolesResponse struct {
	Membership Membership `json:"membership,omitempty\"` // The `Membership` resource after modifying its `MembershipRole`s.

}

// OIDC IDP (identity provider) configuration.
type OidcIdpConfig struct {
	ChangePasswordUri string `json:"changePasswordUri,omitempty\"` // The **Change Password URL** of the identity provider. Users will be sent to this URL when changing their passwords at `myaccount.google.com`. This takes precedence over the change password URL configured at customer-level. Must use `HTTPS`.

	IssuerUri string `json:"issuerUri,omitempty\"` // Required. The Issuer identifier for the IdP. Must be a URL. The discovery URL will be derived from this as described in Section 4 of [the OIDC specification](https://openid.net/specs/openid-connect-discovery-1_0.html).

}

// OIDC RP (relying party) configuration.
type OidcRpConfig struct {
	ClientId string `json:"clientId,omitempty\"` // OAuth2 client ID for OIDC.

	ClientSecret string `json:"clientSecret,omitempty\"` // Input only. OAuth2 client secret for OIDC.

	RedirectUris []string `json:"redirectUris,omitempty\"` // Output only. The URL(s) that this client may use in authentication requests.

}

// Details that are applicable when `sso_mode` is set to `OIDC_SSO`.
type OidcSsoInfo struct {
	InboundOidcSsoProfile string `json:"inboundOidcSsoProfile,omitempty\"` // Required. Name of the `InboundOidcSsoProfile` to use. Must be of the form `inboundOidcSsoProfiles/{inbound_oidc_sso_profile}`.

}

// This resource represents a long-running operation that is the result of a network API call.
type Operation struct {
	Done bool `json:"done,omitempty\"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.

	Error Status `json:"error,omitempty\"` // The error result of the operation in case of failure or cancellation.

	Metadata map[string]interface{} `json:"metadata,omitempty\"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.

	Name string `json:"name,omitempty\"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.

	Response map[string]interface{} `json:"response,omitempty\"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.

}

// A Policy resource binds an instance of a single Setting with the scope of a PolicyQuery. The Setting instance will be applied to all entities that satisfy the query.
type Policy struct {
	Customer string `json:"customer,omitempty\"` // Immutable. Customer that the Policy belongs to. The value is in the format 'customers/{customerId}'. The `customerId` must begin with "C" To find your customer ID in Admin Console see https://support.google.com/a/answer/10070793.

	Name string `json:"name,omitempty\"` // Output only. Identifier. The [resource name](https://cloud.google.com/apis/design/resource_names) of the Policy. Format: policies/{policy}.

	PolicyQuery PolicyQuery `json:"policyQuery,omitempty\"` // Required. The PolicyQuery the Setting applies to.

	Setting Setting `json:"setting,omitempty\"` // Required. The Setting configured by this Policy.

	TypeValue string `json:"type,omitempty\"` // Output only. The type of the policy.

}

// PolicyQuery
type PolicyQuery struct {
	Group string `json:"group,omitempty\"` // Immutable. The group that the query applies to. This field is only set if there is a single value for group that satisfies all clauses of the query. If no group applies, this will be the empty string.

	OrgUnit string `json:"orgUnit,omitempty\"` // Required. Immutable. Non-empty default. The OrgUnit the query applies to. This field is only set if there is a single value for org_unit that satisfies all clauses of the query.

	Query string `json:"query,omitempty\"` // Immutable. The CEL query that defines which entities the Policy applies to (ex. a User entity). For details about CEL see https://opensource.google.com/projects/cel. The OrgUnits the Policy applies to are represented by a clause like so: entity.org_units.exists(org_unit, org_unit.org_unit_id == orgUnitId('{orgUnitId}')) The Group the Policy applies to are represented by a clause like so: entity.groups.exists(group, group.group_id == groupId('{groupId}')) The Licenses the Policy applies to are represented by a clause like so: entity.licenses.exists(license, license in ['/product/{productId}/sku/{skuId}']) **Note:** The licenses clause is not supported in mutate endpoints. The above clauses can be present in any combination, and used in conjunction with the &&, || and ! operators. The org_unit and group fields below are helper fields that contain the corresponding value(s) as the query to make the query easier to use.

	SortOrder float64 `json:"sortOrder,omitempty\"` // Output only. The decimal sort order of this PolicyQuery. The value is relative to all other policies with the same setting type for the customer. (There are no duplicates within this set).

}

// The evaluated state of this restriction.
type RestrictionEvaluation struct {
	State string `json:"state,omitempty\"` // Output only. The current state of the restriction

}

// Evaluations of restrictions applied to parent group on this membership.
type RestrictionEvaluations struct {
	MemberRestrictionEvaluation MembershipRoleRestrictionEvaluation `json:"memberRestrictionEvaluation,omitempty\"` // Evaluation of the member restriction applied to this membership. Empty if the user lacks permission to view the restriction evaluation.

}

// Information of a RSA public key.
type RsaPublicKeyInfo struct {
	KeySize int `json:"keySize,omitempty\"` // Key size in bits (size of the modulus).

}

// SAML IDP (identity provider) configuration.
type SamlIdpConfig struct {
	ChangePasswordUri string `json:"changePasswordUri,omitempty\"` // The **Change Password URL** of the identity provider. Users will be sent to this URL when changing their passwords at `myaccount.google.com`. This takes precedence over the change password URL configured at customer-level. Must use `HTTPS`.

	EntityId string `json:"entityId,omitempty\"` // Required. The SAML **Entity ID** of the identity provider.

	LogoutRedirectUri string `json:"logoutRedirectUri,omitempty\"` // The **Logout Redirect URL** (sign-out page URL) of the identity provider. When a user clicks the sign-out link on a Google page, they will be redirected to this URL. This is a pure redirect with no attached SAML `LogoutRequest` i.e. SAML single logout is not supported. Must use `HTTPS`.

	SingleSignOnServiceUri string `json:"singleSignOnServiceUri,omitempty\"` // Required. The `SingleSignOnService` endpoint location (sign-in page URL) of the identity provider. This is the URL where the `AuthnRequest` will be sent. Must use `HTTPS`. Assumed to accept the `HTTP-Redirect` binding.

}

// SAML SP (service provider) configuration.
type SamlSpConfig struct {
	AssertionConsumerServiceUri string `json:"assertionConsumerServiceUri,omitempty\"` // Output only. The SAML **Assertion Consumer Service (ACS) URL** to be used for the IDP-initiated login. Assumed to accept response messages via the `HTTP-POST` binding.

	EntityId string `json:"entityId,omitempty\"` // Output only. The SAML **Entity ID** for this service provider.

}

// Details that are applicable when `sso_mode` == `SAML_SSO`.
type SamlSsoInfo struct {
	InboundSamlSsoProfile string `json:"inboundSamlSsoProfile,omitempty\"` // Required. Name of the `InboundSamlSsoProfile` to use. Must be of the form `inboundSamlSsoProfiles/{inbound_saml_sso_profile}`.

}

// The response message for MembershipsService.SearchDirectGroups.
type SearchDirectGroupsResponse struct {
	Memberships []MembershipRelation `json:"memberships,omitempty\"` // List of direct groups satisfying the query.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results, or empty if there are no more results available for listing.

}

// The response message for GroupsService.SearchGroups.
type SearchGroupsResponse struct {
	Groups []Group `json:"groups,omitempty\"` // The `Group` resources that match the search query.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A continuation token to retrieve the next page of results, or empty if there are no more results available.

}

// The response message for MembershipsService.SearchTransitiveGroups.
type SearchTransitiveGroupsResponse struct {
	Memberships []GroupRelation `json:"memberships,omitempty\"` // List of transitive groups satisfying the query.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results, or empty if there are no more results available for listing.

}

// The response message for MembershipsService.SearchTransitiveMemberships.
type SearchTransitiveMembershipsResponse struct {
	Memberships []MemberRelation `json:"memberships,omitempty\"` // List of transitive members satisfying the query.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results, or empty if there are no more results.

}

// The definition of security settings.
type SecuritySettings struct {
	MemberRestriction MemberRestriction `json:"memberRestriction,omitempty\"` // The Member Restriction value

	Name string `json:"name,omitempty\"` // Output only. The resource name of the security settings. Shall be of the form `groups/{group_id}/securitySettings`.

}

// A request to send email for inviting target user corresponding to the UserInvitation.
type SendUserInvitationRequest struct {
}

// Setting
type Setting struct {
	TypeValue string `json:"type,omitempty\"` // Required. Immutable. The type of the Setting.

	Value map[string]interface{} `json:"value,omitempty\"` // Required. The value of the Setting.

}

// Controls sign-in behavior.
type SignInBehavior struct {
	RedirectCondition string `json:"redirectCondition,omitempty\"` // When to redirect sign-ins to the IdP.

}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	Code int `json:"code,omitempty\"` // The status code, which should be an enum value of google.rpc.Code.

	Details []map[string]interface{} `json:"details,omitempty\"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.

	Message string `json:"message,omitempty\"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.

}

// Message representing the role of a TransitiveMembership.
type TransitiveMembershipRole struct {
	Role string `json:"role,omitempty\"` // TransitiveMembershipRole in string format. Currently supported TransitiveMembershipRoles: `"MEMBER"`, `"OWNER"`, and `"MANAGER"`.

}

// Metadata for UpdateGroup LRO.
type UpdateGroupMetadata struct {
}

// LRO response metadata for InboundOidcSsoProfilesService.UpdateInboundOidcSsoProfile.
type UpdateInboundOidcSsoProfileOperationMetadata struct {
	State string `json:"state,omitempty\"` // State of this Operation Will be "awaiting-multi-party-approval" when the operation is deferred due to the target customer having enabled [Multi-party approval for sensitive actions](https://support.google.com/a/answer/13790448).

}

// LRO response metadata for InboundSamlSsoProfilesService.UpdateInboundSamlSsoProfile.
type UpdateInboundSamlSsoProfileOperationMetadata struct {
	State string `json:"state,omitempty\"` // State of this Operation Will be "awaiting-multi-party-approval" when the operation is deferred due to the target customer having enabled [Multi-party approval for sensitive actions](https://support.google.com/a/answer/13790448).

}

// LRO response metadata for InboundSsoAssignmentsService.UpdateInboundSsoAssignment.
type UpdateInboundSsoAssignmentOperationMetadata struct {
}

// Metadata for UpdateMembership LRO.
type UpdateMembershipMetadata struct {
}

// The details of an update to a `MembershipRole`.
type UpdateMembershipRolesParams struct {
	FieldMask string `json:"fieldMask,omitempty\"` // The fully-qualified names of fields to update. May only contain the field `expiry_detail.expire_time`.

	MembershipRole MembershipRole `json:"membershipRole,omitempty\"` // The `MembershipRole`s to be updated. Only `MEMBER` `MembershipRole` can currently be updated.

}

// The `UserInvitation` resource represents an email that can be sent to an unmanaged user account inviting them to join the customer's Google Workspace or Cloud Identity account. An unmanaged account shares an email address domain with the Google Workspace or Cloud Identity account but is not managed by it yet. If the user accepts the `UserInvitation`, the user account will become managed.
type UserInvitation struct {
	MailsSentCount int64 `json:"mailsSentCount,omitempty\"` // Number of invitation emails sent to the user.

	Name string `json:"name,omitempty\"` // Shall be of the form `customers/{customer}/userinvitations/{user_email_address}`.

	State string `json:"state,omitempty\"` // State of the `UserInvitation`.

	UpdateTime string `json:"updateTime,omitempty\"` // Time when the `UserInvitation` was last updated.

}
