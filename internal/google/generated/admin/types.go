// Admin SDK API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package admin

import "time"

// JSON template for Alias object in Directory API.
type Alias struct {
	Alias string `json:"alias,omitempty\"`

	Etag string `json:"etag,omitempty\"`

	Id string `json:"id,omitempty\"`

	Kind string `json:"kind,omitempty\"`

	PrimaryEmail string `json:"primaryEmail,omitempty\"`
}

// JSON response template to list aliases in Directory API.
type Aliases struct {
	Aliases []interface{} `json:"aliases,omitempty\"`

	Etag string `json:"etag,omitempty\"`

	Kind string `json:"kind,omitempty\"`
}

// An application-specific password (ASP) is used with applications that do not accept a verification code when logging into the application on certain devices. The ASP access code is used instead of the login and password you commonly use when accessing an application through a browser. For more information about ASPs and how to create one, see the [help center](https://support.google.com/a/answer/2537800#asp).
type Asp struct {
	CodeId int `json:"codeId,omitempty\"` // The unique ID of the ASP.

	CreationTime int64 `json:"creationTime,omitempty\"` // The time when the ASP was created. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format.

	Etag string `json:"etag,omitempty\"` // ETag of the ASP.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#asp`.

	LastTimeUsed int64 `json:"lastTimeUsed,omitempty\"` // The time when the ASP was last used. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format.

	Name string `json:"name,omitempty\"` // The name of the application that the user, represented by their `userId`, entered when the ASP was created.

	UserKey string `json:"userKey,omitempty\"` // The unique ID of the user who issued the ASP.

}

type Asps struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []Asp `json:"items,omitempty\"` // A list of ASP resources.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#aspList`.

}

// Auxiliary message about issues with printers or settings. Example: {message_type:AUXILIARY_MESSAGE_WARNING, field_mask:make_and_model, message:"Given printer is invalid or no longer supported."}
type AuxiliaryMessage struct {
	AuxiliaryMessage string `json:"auxiliaryMessage,omitempty\"` // Human readable message in English. Example: "Given printer is invalid or no longer supported."

	FieldMask string `json:"fieldMask,omitempty\"` // Field that this message concerns.

	Severity string `json:"severity,omitempty\"` // Message severity

}

// Information about the device's backlights.
type BacklightInfo struct {
	Brightness int `json:"brightness,omitempty\"` // Output only. Current brightness of the backlight, between 0 and max_brightness.

	MaxBrightness int `json:"maxBrightness,omitempty\"` // Output only. Maximum brightness for the backlight.

	Path string `json:"path,omitempty\"` // Output only. Path to this backlight on the system. Useful if the caller needs to correlate with other information.

}

// A request for changing the status of a batch of ChromeOS devices.
type BatchChangeChromeOsDeviceStatusRequest struct {
	ChangeChromeOsDeviceStatusAction string `json:"changeChromeOsDeviceStatusAction,omitempty\"` // Required. The action to take on the ChromeOS device in order to change its status.

	DeprovisionReason string `json:"deprovisionReason,omitempty\"` // Optional. The reason behind a device deprovision. Must be provided if 'changeChromeOsDeviceStatusAction' is set to 'CHANGE_CHROME_OS_DEVICE_STATUS_ACTION_DEPROVISION'. Otherwise, omit this field.

	DeviceIds []string `json:"deviceIds,omitempty\"` // Required. List of the IDs of the ChromeOS devices to change. Maximum 50.

}

// The response of changing the status of a batch of ChromeOS devices.
type BatchChangeChromeOsDeviceStatusResponse struct {
	ChangeChromeOsDeviceStatusResults []ChangeChromeOsDeviceStatusResult `json:"changeChromeOsDeviceStatusResults,omitempty\"` // The results for each of the ChromeOS devices provided in the request.

}

// Request to add multiple new print servers in a batch.
type BatchCreatePrintServersRequest struct {
	Requests []CreatePrintServerRequest `json:"requests,omitempty\"` // Required. A list of `PrintServer` resources to be created (max `50` per batch).

}

type BatchCreatePrintServersResponse struct {
	Failures []PrintServerFailureInfo `json:"failures,omitempty\"` // A list of create failures. `PrintServer` IDs are not populated, as print servers were not created.

	PrintServers []PrintServer `json:"printServers,omitempty\"` // A list of successfully created print servers with their IDs populated.

}

// Request for adding new printers in batch.
type BatchCreatePrintersRequest struct {
	Requests []CreatePrinterRequest `json:"requests,omitempty\"` // A list of Printers to be created. Max 50 at a time.

}

// Response for adding new printers in batch.
type BatchCreatePrintersResponse struct {
	Failures []FailureInfo `json:"failures,omitempty\"` // A list of create failures. Printer IDs are not populated, as printer were not created.

	Printers []Printer `json:"printers,omitempty\"` // A list of successfully created printers with their IDs populated.

}

// Request to delete multiple existing print servers in a batch.
type BatchDeletePrintServersRequest struct {
	PrintServerIds []string `json:"printServerIds,omitempty\"` // A list of print server IDs that should be deleted (max `100` per batch).

}

type BatchDeletePrintServersResponse struct {
	FailedPrintServers []PrintServerFailureInfo `json:"failedPrintServers,omitempty\"` // A list of update failures.

	PrintServerIds []string `json:"printServerIds,omitempty\"` // A list of print server IDs that were successfully deleted.

}

// Request for deleting existing printers in batch.
type BatchDeletePrintersRequest struct {
	PrinterIds []string `json:"printerIds,omitempty\"` // A list of Printer.id that should be deleted. Max 100 at a time.

}

// Response for deleting existing printers in batch.
type BatchDeletePrintersResponse struct {
	FailedPrinters []FailureInfo `json:"failedPrinters,omitempty\"` // A list of update failures.

	PrinterIds []string `json:"printerIds,omitempty\"` // A list of Printer.id that were successfully deleted.

}

// Information about a device's Bluetooth adapter.
type BluetoothAdapterInfo struct {
	Address string `json:"address,omitempty\"` // Output only. The MAC address of the adapter.

	NumConnectedDevices int `json:"numConnectedDevices,omitempty\"` // Output only. The number of devices connected to this adapter.

}

// Public API: Resources.buildings
type Building struct {
	Address BuildingAddress `json:"address,omitempty\"` // The postal address of the building. See [`PostalAddress`](/my-business/reference/rest/v4/PostalAddress) for details. Note that only a single address line and region code are required.

	BuildingId string `json:"buildingId,omitempty\"` // Unique identifier for the building. The maximum length is 100 characters.

	BuildingName string `json:"buildingName,omitempty\"` // The building name as seen by users in Calendar. Must be unique for the customer. For example, "NYC-CHEL". The maximum length is 100 characters.

	Coordinates BuildingCoordinates `json:"coordinates,omitempty\"` // The geographic coordinates of the center of the building, expressed as latitude and longitude in decimal degrees.

	Description string `json:"description,omitempty\"` // A brief description of the building. For example, "Chelsea Market".

	Etags string `json:"etags,omitempty\"` // ETag of the resource.

	FloorNames []string `json:"floorNames,omitempty\"` // The display names for all floors in this building. The floors are expected to be sorted in ascending order, from lowest floor to highest floor. For example, ["B2", "B1", "L", "1", "2", "2M", "3", "PH"] Must contain at least one entry.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

}

// Public API: Resources.buildings
type BuildingAddress struct {
	AddressLines []string `json:"addressLines,omitempty\"` // Unstructured address lines describing the lower levels of an address.

	AdministrativeArea string `json:"administrativeArea,omitempty\"` // Optional. Highest administrative subdivision which is used for postal addresses of a country or region.

	LanguageCode string `json:"languageCode,omitempty\"` // Optional. BCP-47 language code of the contents of this address (if known).

	Locality string `json:"locality,omitempty\"` // Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.

	PostalCode string `json:"postalCode,omitempty\"` // Optional. Postal code of the address.

	RegionCode string `json:"regionCode,omitempty\"` // Required. CLDR region code of the country/region of the address.

	Sublocality string `json:"sublocality,omitempty\"` // Optional. Sublocality of the address.

}

// Public API: Resources.buildings
type BuildingCoordinates struct {
	Latitude float64 `json:"latitude,omitempty\"` // Latitude in decimal degrees.

	Longitude float64 `json:"longitude,omitempty\"` // Longitude in decimal degrees.

}

// Public API: Resources.buildings
type Buildings struct {
	Buildings []Building `json:"buildings,omitempty\"` // The Buildings in this page of results.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.

}

// Represents a data capacity with some amount of current usage in bytes.
type ByteUsage struct {
	CapacityBytes int64 `json:"capacityBytes,omitempty\"` // Output only. The total capacity value, in bytes.

	UsedBytes int64 `json:"usedBytes,omitempty\"` // Output only. The current usage value, in bytes.

}

// Public API: Resources.calendars
type CalendarResource struct {
	BuildingId string `json:"buildingId,omitempty\"` // Unique ID for the building a resource is located in.

	Capacity int `json:"capacity,omitempty\"` // Capacity of a resource, number of seats in a room.

	Etags string `json:"etags,omitempty\"` // ETag of the resource.

	FeatureInstances interface{} `json:"featureInstances,omitempty\"` // Instances of features for the calendar resource.

	FloorName string `json:"floorName,omitempty\"` // Name of the floor a resource is located on.

	FloorSection string `json:"floorSection,omitempty\"` // Name of the section within a floor a resource is located in.

	GeneratedResourceName string `json:"generatedResourceName,omitempty\"` // The read-only auto-generated name of the calendar resource which includes metadata about the resource such as building name, floor, capacity, etc. For example, "NYC-2-Training Room 1A (16)".

	Kind string `json:"kind,omitempty\"` // The type of the resource. For calendar resources, the value is `admin#directory#resources#calendars#CalendarResource`.

	ResourceCategory string `json:"resourceCategory,omitempty\"` // The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.

	ResourceDescription string `json:"resourceDescription,omitempty\"` // Description of the resource, visible only to admins.

	ResourceEmail string `json:"resourceEmail,omitempty\"` // The read-only email for the calendar resource. Generated as part of creating a new calendar resource.

	ResourceId string `json:"resourceId,omitempty\"` // The unique ID for the calendar resource.

	ResourceName string `json:"resourceName,omitempty\"` // The name of the calendar resource. For example, "Training Room 1A".

	ResourceType string `json:"resourceType,omitempty\"` // The type of the calendar resource, intended for non-room resources.

	UserVisibleDescription string `json:"userVisibleDescription,omitempty\"` // Description of the resource, visible to users and admins.

}

// Public API: Resources.calendars
type CalendarResources struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []CalendarResource `json:"items,omitempty\"` // The CalendarResources in this page of results.

	Kind string `json:"kind,omitempty\"` // Identifies this as a collection of CalendarResources. This is always `admin#directory#resources#calendars#calendarResourcesList`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.

}

// The result of a single ChromeOS device for a Change state operation.
type ChangeChromeOsDeviceStatusResult struct {
	DeviceId string `json:"deviceId,omitempty\"` // The unique ID of the ChromeOS device.

	Error Status `json:"error,omitempty\"` // The error result of the operation in case of failure.

	Response ChangeChromeOsDeviceStatusSucceeded `json:"response,omitempty\"` // The device could change its status successfully.

}

// Response for a successful ChromeOS device status change.
type ChangeChromeOsDeviceStatusSucceeded struct {
}

// An notification channel used to watch for resource changes.
type Channel struct {
	Address string `json:"address,omitempty\"` // The address where notifications are delivered for this channel.

	Expiration int64 `json:"expiration,omitempty\"` // Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.

	Id string `json:"id,omitempty\"` // A UUID or similar unique string that identifies this channel.

	Kind string `json:"kind,omitempty\"` // Identifies this as a notification channel used to watch for changes to a resource, which is `api#channel`.

	Params map[string]interface{} `json:"params,omitempty\"` // Additional parameters controlling delivery channel behavior. Optional. For example, `params.ttl` specifies the time-to-live in seconds for the notification channel, where the default is 2 hours and the maximum TTL is 2 days.

	Payload bool `json:"payload,omitempty\"` // A Boolean value to indicate whether payload is wanted. Optional.

	ResourceId string `json:"resourceId,omitempty\"` // An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.

	ResourceUri string `json:"resourceUri,omitempty\"` // A version-specific identifier for the watched resource.

	Token string `json:"token,omitempty\"` // An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.

	TypeValue string `json:"type,omitempty\"` // The type of delivery mechanism used for this channel.

}

// Google Chrome devices run on the [Chrome OS](https://support.google.com/chromeos). For more information about common API tasks, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-chrome-devices).
type ChromeOsDevice struct {
	ActiveTimeRanges []map[string]interface{} `json:"activeTimeRanges,omitempty\"` // A list of active time ranges (Read-only).

	AnnotatedAssetId string `json:"annotatedAssetId,omitempty\"` // The asset identifier as noted by an administrator or specified during enrollment.

	AnnotatedLocation string `json:"annotatedLocation,omitempty\"` // The address or location of the device as noted by the administrator. Maximum length is `200` characters. Empty values are allowed.

	AnnotatedUser string `json:"annotatedUser,omitempty\"` // The user of the device as noted by the administrator. Maximum length is 100 characters. Empty values are allowed.

	AutoUpdateExpiration int64 `json:"autoUpdateExpiration,omitempty\"` // (Read-only) The timestamp after which the device will stop receiving Chrome updates or support. Please use "autoUpdateThrough" instead.

	AutoUpdateThrough string `json:"autoUpdateThrough,omitempty\"` // Output only. The timestamp after which the device will stop receiving Chrome updates or support.

	BacklightInfo []BacklightInfo `json:"backlightInfo,omitempty\"` // Output only. Contains backlight information for the device.

	BluetoothAdapterInfo []BluetoothAdapterInfo `json:"bluetoothAdapterInfo,omitempty\"` // Output only. Information about Bluetooth adapters of the device.

	BootMode string `json:"bootMode,omitempty\"` // The boot mode for the device. The possible values are: * `Verified`: The device is running a valid version of the Chrome OS. * `Dev`: The devices's developer hardware switch is enabled. When booted, the device has a command line shell. For an example of a developer switch, see the [Chromebook developer information](https://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices/samsung-series-5-chromebook#TOC-Developer-switch).

	ChromeOsType string `json:"chromeOsType,omitempty\"` // Output only. Chrome OS type of the device.

	CpuInfo []map[string]interface{} `json:"cpuInfo,omitempty\"` // Information regarding CPU specs in the device.

	CpuStatusReports []map[string]interface{} `json:"cpuStatusReports,omitempty\"` // Reports of CPU utilization and temperature (Read-only)

	DeprovisionReason string `json:"deprovisionReason,omitempty\"` // (Read-only) Deprovision reason.

	DeviceFiles []map[string]interface{} `json:"deviceFiles,omitempty\"` // A list of device files to download (Read-only)

	DeviceId string `json:"deviceId,omitempty\"` // The unique ID of the Chrome device.

	DeviceLicenseType string `json:"deviceLicenseType,omitempty\"` // Output only. Device license type.

	DiskSpaceUsage ByteUsage `json:"diskSpaceUsage,omitempty\"` // Output only. How much disk space the device has available and is currently using.

	DiskVolumeReports []map[string]interface{} `json:"diskVolumeReports,omitempty\"` // Reports of disk space and other info about mounted/connected volumes.

	DockMacAddress string `json:"dockMacAddress,omitempty\"` // (Read-only) Built-in MAC address for the docking station that the device connected to. Factory sets Media access control address (MAC address) assigned for use by a dock. It is reserved specifically for MAC pass through device policy. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty\"` // The device's MAC address on the ethernet network interface.

	EthernetMacAddress0 string `json:"ethernetMacAddress0,omitempty\"` // (Read-only) MAC address used by the Chromebook’s internal ethernet port, and for onboard network (ethernet) interface. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.

	ExtendedSupportEligible bool `json:"extendedSupportEligible,omitempty\"` // Output only. Whether or not the device requires the extended support opt in.

	ExtendedSupportEnabled bool `json:"extendedSupportEnabled,omitempty\"` // Output only. Whether extended support policy is enabled on the device.

	ExtendedSupportStart string `json:"extendedSupportStart,omitempty\"` // Output only. Date of the device when extended support policy for automatic updates starts.

	FanInfo []FanInfo `json:"fanInfo,omitempty\"` // Output only. Fan information for the device.

	FirmwareVersion string `json:"firmwareVersion,omitempty\"` // The Chrome device's firmware version.

	FirstEnrollmentTime string `json:"firstEnrollmentTime,omitempty\"` // Date and time for the first time the device was enrolled.

	Kind string `json:"kind,omitempty\"` // The type of resource. For the Chromeosdevices resource, the value is `admin#directory#chromeosdevice`.

	LastDeprovisionTimestamp string `json:"lastDeprovisionTimestamp,omitempty\"` // (Read-only) Date and time for the last deprovision of the device.

	LastEnrollmentTime time.Time `json:"lastEnrollmentTime,omitempty\"` // Date and time the device was last enrolled (Read-only)

	LastKnownNetwork []map[string]interface{} `json:"lastKnownNetwork,omitempty\"` // Contains last known network (Read-only)

	LastSync time.Time `json:"lastSync,omitempty\"` // Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel (Read-only)

	MacAddress string `json:"macAddress,omitempty\"` // The device's wireless MAC address. If the device does not have this information, it is not included in the response.

	ManufactureDate string `json:"manufactureDate,omitempty\"` // (Read-only) The date the device was manufactured in yyyy-mm-dd format.

	Meid string `json:"meid,omitempty\"` // The Mobile Equipment Identifier (MEID) or the International Mobile Equipment Identity (IMEI) for the 3G mobile card in a mobile device. A MEID/IMEI is typically used when adding a device to a wireless carrier's post-pay service plan. If the device does not have this information, this property is not included in the response. For more information on how to export a MEID/IMEI list, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-chrome-devices.html#export_meid).

	Model string `json:"model,omitempty\"` // The device's model information. If the device does not have this information, this property is not included in the response.

	Notes string `json:"notes,omitempty\"` // Notes about this device added by the administrator. This property can be [searched](https://support.google.com/chrome/a/answer/1698333) with the [list](https://developers.google.com/workspace/admin/directory/v1/reference/chromeosdevices/list) method's `query` parameter. Maximum length is 500 characters. Empty values are allowed.

	OrderNumber string `json:"orderNumber,omitempty\"` // The device's order number. Only devices directly purchased from Google have an order number.

	OrgUnitId string `json:"orgUnitId,omitempty\"` // The unique ID of the organizational unit. orgUnitPath is the human readable version of orgUnitId. While orgUnitPath may change by renaming an organizational unit within the path, orgUnitId is unchangeable for one organizational unit. This property can be [updated](https://developers.google.com/workspace/admin/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).

	OrgUnitPath string `json:"orgUnitPath,omitempty\"` // The full parent path with the organizational unit's name associated with the device. Path names are case insensitive. If the parent organizational unit is the top-level organization, it is represented as a forward slash, `/`. This property can be [updated](https://developers.google.com/workspace/admin/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).

	OsUpdateStatus OsUpdateStatus `json:"osUpdateStatus,omitempty\"` // The status of the OS updates for the device.

	OsVersion string `json:"osVersion,omitempty\"` // The Chrome device's operating system version.

	OsVersionCompliance string `json:"osVersionCompliance,omitempty\"` // Output only. Device policy compliance status of the OS version.

	PlatformVersion string `json:"platformVersion,omitempty\"` // The Chrome device's platform version.

	RecentUsers []map[string]interface{} `json:"recentUsers,omitempty\"` // A list of recent device users, in descending order, by last login time.

	ScreenshotFiles []map[string]interface{} `json:"screenshotFiles,omitempty\"` // A list of screenshot files to download. Type is always "SCREENSHOT_FILE". (Read-only)

	SerialNumber string `json:"serialNumber,omitempty\"` // The Chrome device serial number entered when the device was enabled. This value is the same as the Admin console's *Serial Number* in the *Chrome OS Devices* tab.

	Status string `json:"status,omitempty\"` // The status of the device.

	SupportEndDate time.Time `json:"supportEndDate,omitempty\"` // Final date the device will be supported (Read-only)

	SystemRamFreeReports []map[string]interface{} `json:"systemRamFreeReports,omitempty\"` // Reports of amounts of available RAM memory (Read-only)

	SystemRamTotal int64 `json:"systemRamTotal,omitempty\"` // Total RAM on the device [in bytes] (Read-only)

	TpmVersionInfo map[string]interface{} `json:"tpmVersionInfo,omitempty\"` // Trusted Platform Module (TPM) (Read-only)

	WillAutoRenew bool `json:"willAutoRenew,omitempty\"` // Determines if the device will auto renew its support after the support end date. This is a read-only property.

}

// Data about an update to the status of a Chrome OS device.
type ChromeOsDeviceAction struct {
	Action string `json:"action,omitempty\"` // Action to be taken on the Chrome OS device.

	DeprovisionReason string `json:"deprovisionReason,omitempty\"` // Only used when the action is `deprovision`. With the `deprovision` action, this field is required. *Note*: The deprovision reason is audited because it might have implications on licenses for perpetual subscription customers.

}

type ChromeOsDevices struct {
	Chromeosdevices []ChromeOsDevice `json:"chromeosdevices,omitempty\"` // A list of Chrome OS Device objects.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access the next page of this result. To access the next page, use this token's value in the `pageToken` query string of this request.

}

type ChromeOsMoveDevicesToOu struct {
	DeviceIds []string `json:"deviceIds,omitempty\"` // Chrome OS devices to be moved to OU

}

// A response for counting ChromeOS devices.
type CountChromeOsDevicesResponse struct {
	Count int64 `json:"count,omitempty\"` // The total number of devices matching the request.

}

// Request for adding a new print server.
type CreatePrintServerRequest struct {
	Parent string `json:"parent,omitempty\"` // Required. The [unique ID](https://developers.google.com/workspace/admin/directory/reference/rest/v1/customers) of the customer's Google Workspace account. Format: `customers/{id}`

	PrintServer PrintServer `json:"printServer,omitempty\"` // Required. A print server to create. If you want to place the print server under a specific organizational unit (OU), then populate the `org_unit_id`. Otherwise the print server is created under the root OU. The `org_unit_id` can be retrieved using the [Directory API](https://developers.google.com/workspace/admin/directory/v1/guides/manage-org-units).

}

// Request for adding a new printer.
type CreatePrinterRequest struct {
	Parent string `json:"parent,omitempty\"` // Required. The name of the customer. Format: customers/{customer_id}

	Printer Printer `json:"printer,omitempty\"` // Required. A printer to create. If you want to place the printer under particular OU then populate printer.org_unit_id filed. Otherwise the printer will be placed under root OU.

}

type Customer struct {
	AlternateEmail string `json:"alternateEmail,omitempty\"` // The customer's secondary contact email address. This email address cannot be on the same domain as the `customerDomain`

	CustomerCreationTime time.Time `json:"customerCreationTime,omitempty\"` // The customer's creation time (Readonly)

	CustomerDomain string `json:"customerDomain,omitempty\"` // The customer's primary domain name string. Do not include the `www` prefix when creating a new customer.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Id string `json:"id,omitempty\"` // The unique ID for the customer's Google Workspace account. (Readonly)

	Kind string `json:"kind,omitempty\"` // Identifies the resource as a customer. Value: `admin#directory#customer`

	Language string `json:"language,omitempty\"` // The customer's ISO 639-2 language code. See the [Language Codes](https://developers.google.com/workspace/admin/directory/v1/languages) page for the list of supported codes. Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior. The default value is `en`.

	PhoneNumber string `json:"phoneNumber,omitempty\"` // The customer's contact phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.

	PostalAddress CustomerPostalAddress `json:"postalAddress,omitempty\"` // The customer's postal address information.

}

type CustomerPostalAddress struct {
	AddressLine1 string `json:"addressLine1,omitempty\"` // A customer's physical address. The address can be composed of one to three lines.

	AddressLine2 string `json:"addressLine2,omitempty\"` // Address line 2 of the address.

	AddressLine3 string `json:"addressLine3,omitempty\"` // Address line 3 of the address.

	ContactName string `json:"contactName,omitempty\"` // The customer contact's name.

	CountryCode string `json:"countryCode,omitempty\"` // This is a required property. For `countryCode` information see the [ISO 3166 country code elements](https://www.iso.org/iso/country_codes.htm).

	Locality string `json:"locality,omitempty\"` // Name of the locality. An example of a locality value is the city of `San Francisco`.

	OrganizationName string `json:"organizationName,omitempty\"` // The company or company division name.

	PostalCode string `json:"postalCode,omitempty\"` // The postal code. A postalCode example is a postal zip code such as `10009`. This is in accordance with - http: //portablecontacts.net/draft-spec.html#address_element.

	Region string `json:"region,omitempty\"` // Name of the region. An example of a region value is `NY` for the state of New York.

}

// Information regarding a command that was issued to a device.
type DirectoryChromeosdevicesCommand struct {
	CommandExpireTime string `json:"commandExpireTime,omitempty\"` // The time at which the command will expire. If the device doesn't execute the command within this time the command will become expired.

	CommandId int64 `json:"commandId,omitempty\"` // Unique ID of a device command.

	CommandResult DirectoryChromeosdevicesCommandResult `json:"commandResult,omitempty\"` // The result of the command execution.

	IssueTime string `json:"issueTime,omitempty\"` // The timestamp when the command was issued by the admin.

	Payload string `json:"payload,omitempty\"` // The payload that the command specified, if any.

	State string `json:"state,omitempty\"` // Indicates the command state.

	TypeValue string `json:"type,omitempty\"` // The type of the command.

}

// The result of executing a command.
type DirectoryChromeosdevicesCommandResult struct {
	CommandResultPayload string `json:"commandResultPayload,omitempty\"` // The payload for the command result. The following commands respond with a payload: * `DEVICE_START_CRD_SESSION`: Payload is a stringified JSON object in the form: { "url": url }. The provided URL links to the Chrome Remote Desktop session and requires authentication using only the `email` associated with the command's issuance. * `FETCH_CRD_AVAILABILITY_INFO`: Payload is a stringified JSON object in the form: { "deviceIdleTimeInSeconds": number, "userSessionType": string, "remoteSupportAvailability": string, "remoteAccessAvailability": string }. The "remoteSupportAvailability" field is set to "AVAILABLE" if `shared` CRD session to the device is available. The "remoteAccessAvailability" field is set to "AVAILABLE" if `private` CRD session to the device is available.

	ErrorMessage string `json:"errorMessage,omitempty\"` // The error message with a short explanation as to why the command failed. Only present if the command failed.

	ExecuteTime string `json:"executeTime,omitempty\"` // The time at which the command was executed or failed to execute.

	Result string `json:"result,omitempty\"` // The result of the command.

}

// A request for issuing a command.
type DirectoryChromeosdevicesIssueCommandRequest struct {
	CommandType string `json:"commandType,omitempty\"` // The type of command.

	Payload string `json:"payload,omitempty\"` // The payload for the command, provide it only if command supports it. The following commands support adding payload: * `SET_VOLUME`: Payload is a stringified JSON object in the form: { "volume": 50 }. The volume has to be an integer in the range [0,100]. * `DEVICE_START_CRD_SESSION`: Payload is optionally a stringified JSON object in the form: { "ackedUserPresence": true, "crdSessionType": string }. `ackedUserPresence` is a boolean. By default, `ackedUserPresence` is set to `false`. To start a Chrome Remote Desktop session for an active device, set `ackedUserPresence` to `true`. `crdSessionType` can only select from values `private` (which grants the remote admin exclusive control of the ChromeOS device) or `shared` (which allows the admin and the local user to share control of the ChromeOS device). If not set, `crdSessionType` defaults to `shared`. The `FETCH_CRD_AVAILABILITY_INFO` command can be used to determine available session types on the device. * `REBOOT`: Payload is a stringified JSON object in the form: { "user_session_delay_seconds": 300 }. The `user_session_delay_seconds` is the amount of seconds to wait before rebooting the device if a user is logged in. It has to be an integer in the range [0,300]. When payload is not present for reboot, 0 delay is the default. Note: This only applies if an actual user is logged in, including a Guest. If the device is in the login screen or in Kiosk mode the value is not respected and the device immediately reboots. * `FETCH_SUPPORT_PACKET`: Payload is optionally a stringified JSON object in the form: {"supportPacketDetails":{ "issueCaseId": optional_support_case_id_string, "issueDescription": optional_issue_description_string, "requestedDataCollectors": []}} The list of available `data_collector_enums` are as following: Chrome System Information (1), Crash IDs (2), Memory Details (3), UI Hierarchy (4), Additional ChromeOS Platform Logs (5), Device Event (6), Intel WiFi NICs Debug Dump (7), Touch Events (8), Lacros (9), Lacros System Information (10), ChromeOS Flex Logs (11), DBus Details (12), ChromeOS Network Routes (13), ChromeOS Shill (Connection Manager) Logs (14), Policies (15), ChromeOS System State and Logs (16), ChromeOS System Logs (17), ChromeOS Chrome User Logs (18), ChromeOS Bluetooth (19), ChromeOS Connected Input Devices (20), ChromeOS Traffic Counters (21), ChromeOS Virtual Keyboard (22), ChromeOS Network Health (23). See more details in [help article](https://support.google.com/chrome/a?p=remote-log).

}

// A response for issuing a command.
type DirectoryChromeosdevicesIssueCommandResponse struct {
	CommandId int64 `json:"commandId,omitempty\"` // The unique ID of the issued command, used to retrieve the command status.

}

// Directory users guest creation request message.
type DirectoryUsersCreateGuestRequest struct {
	Customer string `json:"customer,omitempty\"` // Optional. Immutable ID of the Google Workspace account. Only required when request is created by a service account. Defaults to the authenticated user's customer ID otherwise.

	PrimaryGuestEmail string `json:"primaryGuestEmail,omitempty\"` // Required. External email of the guest user being created.

}

type DomainAlias struct {
	CreationTime int64 `json:"creationTime,omitempty\"` // The creation time of the domain alias. (Read-only).

	DomainAliasName string `json:"domainAliasName,omitempty\"` // The domain alias name.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	ParentDomainName string `json:"parentDomainName,omitempty\"` // The parent domain name that the domain alias is associated with. This can either be a primary or secondary domain name within a customer.

	Verified bool `json:"verified,omitempty\"` // Indicates the verification state of a domain alias. (Read-only)

}

type DomainAliases struct {
	DomainAliases []DomainAlias `json:"domainAliases,omitempty\"` // A list of domain alias objects.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

}

type Domains struct {
	CreationTime int64 `json:"creationTime,omitempty\"` // Creation time of the domain. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format. (Read-only).

	DomainAliases []DomainAlias `json:"domainAliases,omitempty\"` // A list of domain alias objects. (Read-only)

	DomainName string `json:"domainName,omitempty\"` // The domain name of the customer.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	IsPrimary bool `json:"isPrimary,omitempty\"` // Indicates if the domain is a primary domain (Read-only).

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	Verified bool `json:"verified,omitempty\"` // Indicates the verification state of a domain. (Read-only).

}

type Domains2 struct {
	Domains []Domains `json:"domains,omitempty\"` // A list of domain objects.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs. A typical example is to use it as the request or the response type of an API method. For instance: service Foo { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
}

// Details regarding the expiration of this role assignment. Used to automatically revoke access when the time limit is reached.
type ExpirationDetails struct {
	ExpireTime string `json:"expireTime,omitempty\"` // The specific timestamp when the role assignment expires.

}

// External identifier used to link and identify this group across external directory systems.
type ExternalId struct {
	Id string `json:"id,omitempty\"` // The unique identifier string assigned by the external provider.

	Namespace string `json:"namespace,omitempty\"` // The system or identity provider managing this ID.

}

// Info about failures
type FailureInfo struct {
	ErrorCode string `json:"errorCode,omitempty\"` // Canonical code for why the update failed to apply.

	ErrorMessage string `json:"errorMessage,omitempty\"` // Failure reason message.

	Printer Printer `json:"printer,omitempty\"` // Failed printer.

	PrinterId string `json:"printerId,omitempty\"` // Id of a failed printer.

}

// Information about the device's fan.
type FanInfo struct {
	SpeedRpm int `json:"speedRpm,omitempty\"` // Output only. Fan speed in RPM.

}

// JSON template for Feature object in Directory API.
type Feature struct {
	Etags string `json:"etags,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	Name string `json:"name,omitempty\"` // The name of the feature.

}

// JSON template for a feature instance.
type FeatureInstance struct {
	Feature Feature `json:"feature,omitempty\"` // The feature that this is an instance of. A calendar resource may have multiple instances of a feature.

}

type FeatureRename struct {
	NewName string `json:"newName,omitempty\"` // New name of the feature.

}

// Public API: Resources.features
type Features struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Features []Feature `json:"features,omitempty\"` // The Features in this page of results.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.

}

// Google Groups provide your users the ability to send messages to groups of people using the group's email address. For more information about common tasks, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-groups). For information about other types of groups, see the [Cloud Identity Groups API documentation](https://cloud.google.com/identity/docs/groups). Note: The user calling the API (or being impersonated by a service account) must have an assigned [role](https://developers.google.com/workspace/admin/directory/v1/guides/manage-roles) that includes Admin API Groups permissions, such as Super Admin or Groups Admin.
type Group struct {
	AdminCreated bool `json:"adminCreated,omitempty\"` // Read-only. Value is `true` if this group was created by an administrator rather than a user.

	Aliases []string `json:"aliases,omitempty\"` // Read-only. The list of a group's alias email addresses. To add, update, or remove a group's aliases, use the `groups.aliases` methods. If edited in a group's POST or PUT request, the edit is ignored.

	Description string `json:"description,omitempty\"` // An extended description to help users determine the purpose of a group. For example, you can include information about who should join the group, the types of messages to send to the group, links to FAQs about the group, or related groups. Maximum length is `4,096` characters.

	DirectMembersCount int64 `json:"directMembersCount,omitempty\"` // The number of users that are direct members of the group. If a group is a member (child) of this group (the parent), members of the child group are not counted in the `directMembersCount` property of the parent group.

	Email string `json:"email,omitempty\"` // The group's email address. If your account has multiple domains, select the appropriate domain for the email address. The `email` must be unique. This property is required when creating a group. Group email addresses are subject to the same character usage rules as usernames, see the [help center](https://support.google.com/a/answer/9193374) for details.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	ExternalIds []ExternalId `json:"externalIds,omitempty\"` // Optional. The list of external IDs for the group, such as an immutable identifier from an external identity provider or directory sync client. Each entry contains a namespace and an ID value.

	Id string `json:"id,omitempty\"` // Read-only. The unique ID of a group. A group `id` can be used as a group request URI's `groupKey`.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Groups resources, the value is `admin#directory#group`.

	Name string `json:"name,omitempty\"` // The group's display name.

	NonEditableAliases []string `json:"nonEditableAliases,omitempty\"` // Read-only. The list of the group's non-editable alias email addresses that are outside of the account's primary domain or subdomains. These are functioning email addresses used by the group. This is a read-only property returned in the API's response for a group. If edited in a group's POST or PUT request, the edit is ignored.

}

// The Directory API manages aliases, which are alternative email addresses.
type GroupAlias struct {
	Alias string `json:"alias,omitempty\"` // The alias email address.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Id string `json:"id,omitempty\"` // The unique ID of the group.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Alias resources, the value is `admin#directory#alias`.

	PrimaryEmail string `json:"primaryEmail,omitempty\"` // The primary email address of the group.

}

type Groups struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Groups []Group `json:"groups,omitempty\"` // A list of group objects.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access next page of this result.

}

// Account info specific to Guest users.
type GuestAccountInfo struct {
	PrimaryGuestEmail string `json:"primaryGuestEmail,omitempty\"` // Immutable. The guest's external email.

}

type ListPrintServersResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token that can be sent as `page_token` in a request to retrieve the next page. If this field is omitted, there are no subsequent pages.

	PrintServers []PrintServer `json:"printServers,omitempty\"` // List of print servers.

}

// Response for listing allowed printer models.
type ListPrinterModelsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	PrinterModels []PrinterModel `json:"printerModels,omitempty\"` // Printer models that are currently allowed to be configured for ChromeOs. Some printers may be added or removed over time.

}

// Response for listing printers.
type ListPrintersResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	Printers []Printer `json:"printers,omitempty\"` // List of printers. If `org_unit_id` was given in the request, then only printers visible for this OU will be returned. If `org_unit_id` was not given in the request, then all printers will be returned.

}

// A Google Groups member can be a user or another group. This member can be inside or outside of your account's domains. For more information about common group member tasks, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-group-members).
type Member struct {
	DeliverySettings string `json:"delivery_settings,omitempty\"` // Defines mail delivery preferences of member. This field is only supported by `insert`, `update`, and `get` methods.

	Email string `json:"email,omitempty\"` // The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The `email` must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Id string `json:"id,omitempty\"` // The unique ID of the group member. A member `id` can be used as a member request URI's `memberKey`.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Members resources, the value is `admin#directory#member`.

	Role string `json:"role,omitempty\"` // The member's role in a group. The API returns an error for cycles in group memberships. For example, if `group1` is a member of `group2`, `group2` cannot be a member of `group1`. For more information about a member's role, see the [administration help center](https://support.google.com/a/answer/167094).

	Status string `json:"status,omitempty\"` // Status of member (Immutable)

	TypeValue string `json:"type,omitempty\"` // The type of group member.

}

type Members struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	Members []Member `json:"members,omitempty\"` // A list of member objects.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access next page of this result.

}

// JSON template for Has Member response in Directory API.
type MembersHasMember struct {
	IsMember bool `json:"isMember,omitempty\"` // Output only. Identifies whether the given user is a member of the group. Membership can be direct or nested.

}

// Google Workspace Mobile Management includes Android, [Google Sync](https://support.google.com/a/answer/135937), and iOS devices. For more information about common group mobile device API tasks, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-mobile-devices.html).
type MobileDevice struct {
	AdbStatus bool `json:"adbStatus,omitempty\"` // Adb (USB debugging) enabled or disabled on device (Read-only)

	Applications []map[string]interface{} `json:"applications,omitempty\"` // The list of applications installed on an Android mobile device. It is not applicable to Google Sync and iOS devices. The list includes any Android applications that access Google Workspace data. When updating an applications list, it is important to note that updates replace the existing list. If the Android device has two existing applications and the API updates the list with five applications, the is now the updated list of five applications.

	BasebandVersion string `json:"basebandVersion,omitempty\"` // The device's baseband version.

	BootloaderVersion string `json:"bootloaderVersion,omitempty\"` // Mobile Device Bootloader version (Read-only)

	Brand string `json:"brand,omitempty\"` // Mobile Device Brand (Read-only)

	BuildNumber string `json:"buildNumber,omitempty\"` // The device's operating system build number.

	DefaultLanguage string `json:"defaultLanguage,omitempty\"` // The default locale used on the device.

	DeveloperOptionsStatus bool `json:"developerOptionsStatus,omitempty\"` // Developer options enabled or disabled on device (Read-only)

	DeviceCompromisedStatus string `json:"deviceCompromisedStatus,omitempty\"` // The compromised device status.

	DeviceId string `json:"deviceId,omitempty\"` // The serial number for a Google Sync mobile device. For Android and iOS devices, this is a software generated unique identifier.

	DevicePasswordStatus string `json:"devicePasswordStatus,omitempty\"` // DevicePasswordStatus (Read-only)

	Email []string `json:"email,omitempty\"` // The list of the owner's email addresses. If your application needs the current list of user emails, use the [get](https://developers.google.com/workspace/admin/directory/v1/reference/mobiledevices/get.html) method. For additional information, see the [retrieve a user](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users#get_user) method.

	EncryptionStatus string `json:"encryptionStatus,omitempty\"` // Mobile Device Encryption Status (Read-only)

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	FirstSync time.Time `json:"firstSync,omitempty\"` // Date and time the device was first synchronized with the policy settings in the G Suite administrator control panel (Read-only)

	Hardware string `json:"hardware,omitempty\"` // Mobile Device Hardware (Read-only)

	HardwareId string `json:"hardwareId,omitempty\"` // The IMEI/MEID unique identifier for Android hardware. It is not applicable to Google Sync devices. When adding an Android mobile device, this is an optional property. When updating one of these devices, this is a read-only property.

	Imei string `json:"imei,omitempty\"` // The device's IMEI number.

	KernelVersion string `json:"kernelVersion,omitempty\"` // The device's kernel version.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Mobiledevices resources, the value is `admin#directory#mobiledevice`.

	LastSync time.Time `json:"lastSync,omitempty\"` // Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel (Read-only)

	ManagedAccountIsOnOwnerProfile bool `json:"managedAccountIsOnOwnerProfile,omitempty\"` // Boolean indicating if this account is on owner/primary profile or not.

	Manufacturer string `json:"manufacturer,omitempty\"` // Mobile Device manufacturer (Read-only)

	Meid string `json:"meid,omitempty\"` // The device's MEID number.

	Model string `json:"model,omitempty\"` // The mobile device's model name, for example Nexus S. This property can be [updated](https://developers.google.com/workspace/admin/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-mobile=devices#update_mobile_device).

	Name []string `json:"name,omitempty\"` // The list of the owner's user names. If your application needs the current list of device owner names, use the [get](https://developers.google.com/workspace/admin/directory/v1/reference/mobiledevices/get.html) method. For more information about retrieving mobile device user information, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users#get_user).

	NetworkOperator string `json:"networkOperator,omitempty\"` // Mobile Device mobile or network operator (if available) (Read-only)

	Os string `json:"os,omitempty\"` // The mobile device's operating system, for example IOS 4.3 or Android 2.3.5. This property can be [updated](https://developers.google.com/workspace/admin/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-mobile-devices#update_mobile_device).

	OtherAccountsInfo []string `json:"otherAccountsInfo,omitempty\"` // The list of accounts added on device (Read-only)

	Privilege string `json:"privilege,omitempty\"` // DMAgentPermission (Read-only)

	ReleaseVersion string `json:"releaseVersion,omitempty\"` // Mobile Device release version version (Read-only)

	ResourceId string `json:"resourceId,omitempty\"` // The unique ID the API service uses to identify the mobile device.

	SecurityPatchLevel int64 `json:"securityPatchLevel,omitempty\"` // Mobile Device Security patch level (Read-only)

	SerialNumber string `json:"serialNumber,omitempty\"` // The device's serial number.

	Status string `json:"status,omitempty\"` // The device's status.

	SupportsWorkProfile bool `json:"supportsWorkProfile,omitempty\"` // Work profile supported on device (Read-only)

	TypeValue string `json:"type,omitempty\"` // The type of mobile device.

	UnknownSourcesStatus bool `json:"unknownSourcesStatus,omitempty\"` // Unknown sources enabled or disabled on device (Read-only)

	UserAgent string `json:"userAgent,omitempty\"` // Gives information about the device such as `os` version. This property can be [updated](https://developers.google.com/workspace/admin/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-mobile-devices#update_mobile_device).

	WifiMacAddress string `json:"wifiMacAddress,omitempty\"` // The device's MAC address on Wi-Fi networks.

}

type MobileDeviceAction struct {
	Action string `json:"action,omitempty\"` // The action to be performed on the device.

}

type MobileDevices struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	Mobiledevices []MobileDevice `json:"mobiledevices,omitempty\"` // A list of Mobile Device objects.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access next page of this result.

}

// Managing your account's organizational units allows you to configure your users' access to services and custom settings. For more information about common organizational unit tasks, see the [Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-org-units.html). The customer's organizational unit hierarchy is limited to 35 levels of depth.
type OrgUnit struct {
	BlockInheritance bool `json:"blockInheritance,omitempty\"` // This field is deprecated and setting its value has no effect.

	Description string `json:"description,omitempty\"` // Description of the organizational unit.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Orgunits resources, the value is `admin#directory#orgUnit`.

	Name string `json:"name,omitempty\"` // The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales_support parent path is sales_support. Required.

	OrgUnitId string `json:"orgUnitId,omitempty\"` // The unique ID of the organizational unit.

	OrgUnitPath string `json:"orgUnitPath,omitempty\"` // The full path to the organizational unit. The `orgUnitPath` is a derived property. When listed, it is derived from `parentOrgunitPath` and organizational unit's `name`. For example, for an organizational unit named 'apps' under parent organization '/engineering', the orgUnitPath is '/engineering/apps'. In order to edit an `orgUnitPath`, either update the name of the organization or the `parentOrgunitPath`. A user's organizational unit determines which Google Workspace services the user has access to. If the user is moved to a new organization, the user's access changes. For more information about organization structures, see the [administration help center](https://support.google.com/a/answer/4352075). For more information about moving a user to a different organization, see [Update a user](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users.html#update_user).

	ParentOrgUnitId string `json:"parentOrgUnitId,omitempty\"` // The unique ID of the parent organizational unit. Required, unless `parentOrgUnitPath` is set.

	ParentOrgUnitPath string `json:"parentOrgUnitPath,omitempty\"` // The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit. Required, unless `parentOrgUnitId` is set.

}

type OrgUnits struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Org Unit resources, the type is `admin#directory#orgUnits`.

	OrganizationUnits []OrgUnit `json:"organizationUnits,omitempty\"` // A list of organizational unit objects.

}

// Contains information regarding the current OS update status.
type OsUpdateStatus struct {
	RebootTime string `json:"rebootTime,omitempty\"` // Date and time of the last reboot.

	State string `json:"state,omitempty\"` // The update state of an OS update.

	TargetKioskAppVersion string `json:"targetKioskAppVersion,omitempty\"` // New required platform version from the pending updated kiosk app.

	TargetOsVersion string `json:"targetOsVersion,omitempty\"` // New platform version of the OS image being downloaded and applied. It is only set when update status is UPDATE_STATUS_DOWNLOAD_IN_PROGRESS or UPDATE_STATUS_NEED_REBOOT. Note this could be a dummy "0.0.0.0" for UPDATE_STATUS_NEED_REBOOT for some edge cases, e.g. update engine is restarted without a reboot.

	UpdateCheckTime string `json:"updateCheckTime,omitempty\"` // Date and time of the last update check.

	UpdateTime string `json:"updateTime,omitempty\"` // Date and time of the last successful OS update.

}

// Configuration for a print server.
type PrintServer struct {
	CreateTime string `json:"createTime,omitempty\"` // Output only. Time when the print server was created.

	Description string `json:"description,omitempty\"` // Editable. Description of the print server (as shown in the Admin console).

	DisplayName string `json:"displayName,omitempty\"` // Editable. Display name of the print server (as shown in the Admin console).

	Id string `json:"id,omitempty\"` // Immutable. ID of the print server. Leave empty when creating.

	Name string `json:"name,omitempty\"` // Identifier. Resource name of the print server. Leave empty when creating. Format: `customers/{customer.id}/printServers/{print_server.id}`

	OrgUnitId string `json:"orgUnitId,omitempty\"` // ID of the organization unit (OU) that owns this print server. This value can only be set when the print server is initially created. If it's not populated, the print server is placed under the root OU. The `org_unit_id` can be retrieved using the [Directory API](https://developers.google.com/workspace/admin/directory/reference/rest/v1/orgunits).

	Uri string `json:"uri,omitempty\"` // Editable. Print server URI.

}

// Info about failures
type PrintServerFailureInfo struct {
	ErrorCode string `json:"errorCode,omitempty\"` // Canonical code for why the update failed to apply.

	ErrorMessage string `json:"errorMessage,omitempty\"` // Failure reason message.

	PrintServer PrintServer `json:"printServer,omitempty\"` // Failed print server.

	PrintServerId string `json:"printServerId,omitempty\"` // ID of a failed print server.

}

// Printer configuration.
type Printer struct {
	AuxiliaryMessages []AuxiliaryMessage `json:"auxiliaryMessages,omitempty\"` // Output only. Auxiliary messages about issues with the printer configuration if any.

	CreateTime string `json:"createTime,omitempty\"` // Output only. Time when printer was created.

	Description string `json:"description,omitempty\"` // Editable. Description of printer.

	DisplayName string `json:"displayName,omitempty\"` // Editable. Name of printer.

	Id string `json:"id,omitempty\"` // Id of the printer. (During printer creation leave empty)

	MakeAndModel string `json:"makeAndModel,omitempty\"` // Editable. Make and model of printer. e.g. Lexmark MS610de Value must be in format as seen in ListPrinterModels response.

	Name string `json:"name,omitempty\"` // Identifier. The resource name of the Printer object, in the format customers/{customer-id}/printers/{printer-id} (During printer creation leave empty)

	OrgUnitId string `json:"orgUnitId,omitempty\"` // Organization Unit that owns this printer (Only can be set during Printer creation)

	Uri string `json:"uri,omitempty\"` // Editable. Printer URI.

	UseDriverlessConfig bool `json:"useDriverlessConfig,omitempty\"` // Editable. flag to use driverless configuration or not. If it's set to be true, make_and_model can be ignored

}

// Printer manufacturer and model
type PrinterModel struct {
	DisplayName string `json:"displayName,omitempty\"` // Display name. eq. "Brother MFC-8840D"

	MakeAndModel string `json:"makeAndModel,omitempty\"` // Make and model as represented in "make_and_model" field in Printer object. eq. "brother mfc-8840d"

	Manufacturer string `json:"manufacturer,omitempty\"` // Manufacturer. eq. "Brother"

}

type Privilege struct {
	ChildPrivileges []Privilege `json:"childPrivileges,omitempty\"` // A list of child privileges. Privileges for a service form a tree. Each privilege can have a list of child privileges; this list is empty for a leaf privilege.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	IsOuScopable bool `json:"isOuScopable,omitempty\"` // If the privilege can be restricted to an organization unit.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#privilege`.

	PrivilegeName string `json:"privilegeName,omitempty\"` // The name of the privilege.

	ServiceId string `json:"serviceId,omitempty\"` // The obfuscated ID of the service this privilege is for. This value is returned with [`Privileges.list()`](https://developers.google.com/workspace/admin/directory/v1/reference/privileges/list).

	ServiceName string `json:"serviceName,omitempty\"` // The name of the service this privilege is for.

}

type Privileges struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []Privilege `json:"items,omitempty\"` // A list of Privilege resources.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#privileges`.

}

type Role struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	IsSuperAdminRole bool `json:"isSuperAdminRole,omitempty\"` // Returns `true` if the role is a super admin role.

	IsSystemRole bool `json:"isSystemRole,omitempty\"` // Returns `true` if this is a pre-defined system role.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#role`.

	RoleDescription string `json:"roleDescription,omitempty\"` // A short description of the role.

	RoleId int64 `json:"roleId,omitempty\"` // ID of the role.

	RoleName string `json:"roleName,omitempty\"` // Name of the role.

	RolePrivileges []map[string]interface{} `json:"rolePrivileges,omitempty\"` // The set of privileges that are granted to this role.

}

// Defines an assignment of a role.
type RoleAssignment struct {
	AssignedTo string `json:"assignedTo,omitempty\"` // The unique ID of the entity this role is assigned to—either the `user_id` of a user, the `group_id` of a group, or the `uniqueId` of a service account as defined in [Identity and Access Management (IAM)](https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts).

	AssigneeType string `json:"assigneeType,omitempty\"` // Output only. The type of the assignee (`USER` or `GROUP`).

	Condition string `json:"condition,omitempty\"` // Optional. The condition associated with this role assignment. Note: Feature is available to Enterprise Standard, Enterprise Plus, Google Workspace for Education Plus and Cloud Identity Premium customers. A `RoleAssignment` with the `condition` field set will only take effect when the resource being accessed meets the condition. If `condition` is empty, the role (`role_id`) is applied to the actor (`assigned_to`) at the scope (`scope_type`) unconditionally. Currently, the following conditions are supported: - To make the `RoleAssignment` only applicable to [Security Groups](https://cloud.google.com/identity/docs/groups#group_types): `api.getAttribute('cloudidentity.googleapis.com/groups.labels', []).hasAny(['groups.security']) && resource.type == 'cloudidentity.googleapis.com/Group'` - To make the `RoleAssignment` not applicable to [Security Groups](https://cloud.google.com/identity/docs/groups#group_types): `!api.getAttribute('cloudidentity.googleapis.com/groups.labels', []).hasAny(['groups.security']) && resource.type == 'cloudidentity.googleapis.com/Group'` Currently, the condition strings have to be verbatim and they only work with the following [pre-built administrator roles](https://support.google.com/a/answer/2405986): - Groups Editor - Groups Reader The condition follows [Cloud IAM condition syntax](https://cloud.google.com/iam/docs/conditions-overview). - To make the `RoleAssignment` not applicable to [Locked Groups](https://cloud.google.com/identity/docs/groups#group_types): `!api.getAttribute('cloudidentity.googleapis.com/groups.labels', []).hasAny(['groups.locked']) && resource.type == 'cloudidentity.googleapis.com/Group'` This condition can also be used in conjunction with a Security-related condition.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	ExpirationDetails ExpirationDetails `json:"expirationDetails,omitempty\"` // Optional. Details regarding the expiration of this role assignment.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#roleAssignment`.

	OrgUnitId string `json:"orgUnitId,omitempty\"` // If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of this role is restricted to.

	RoleAssignmentId int64 `json:"roleAssignmentId,omitempty\"` // ID of this roleAssignment.

	RoleId int64 `json:"roleId,omitempty\"` // The ID of the role that is assigned.

	ScopeType string `json:"scopeType,omitempty\"` // The scope in which this role is assigned.

}

type RoleAssignments struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []RoleAssignment `json:"items,omitempty\"` // A list of RoleAssignment resources.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#roleAssignments`.

	NextPageToken string `json:"nextPageToken,omitempty\"`
}

type Roles struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []Role `json:"items,omitempty\"` // A list of Role resources.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#roles`.

	NextPageToken string `json:"nextPageToken,omitempty\"`
}

// The type of API resource. For Schema resources, this is always `admin#directory#schema`.
type Schema struct {
	DisplayName string `json:"displayName,omitempty\"` // Display name for the schema.

	Etag string `json:"etag,omitempty\"` // The ETag of the resource.

	Fields []SchemaFieldSpec `json:"fields,omitempty\"` // A list of fields in the schema.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	SchemaId string `json:"schemaId,omitempty\"` // The unique identifier of the schema (Read-only)

	SchemaName string `json:"schemaName,omitempty\"` // The schema's name. Each `schema_name` must be unique within a customer. Reusing a name results in a `409: Entity already exists` error.

}

// You can use schemas to add custom fields to user profiles. You can use these fields to store information such as the projects your users work on, their physical locations, their hire dates, or whatever else fits your business needs. For more information, see [Custom User Fields](https://developers.google.com/workspace/admin/directory/v1/guides/manage-schemas).
type SchemaFieldSpec struct {
	DisplayName string `json:"displayName,omitempty\"` // Display Name of the field.

	Etag string `json:"etag,omitempty\"` // The ETag of the field.

	FieldId string `json:"fieldId,omitempty\"` // The unique identifier of the field (Read-only)

	FieldName string `json:"fieldName,omitempty\"` // The name of the field.

	FieldType string `json:"fieldType,omitempty\"` // The type of the field.

	Indexed bool `json:"indexed,omitempty\"` // Boolean specifying whether the field is indexed or not. Default: `true`.

	Kind string `json:"kind,omitempty\"` // The kind of resource this is. For schema fields this is always `admin#directory#schema#fieldspec`.

	MultiValued bool `json:"multiValued,omitempty\"` // A boolean specifying whether this is a multi-valued field or not. Default: `false`.

	NumericIndexingSpec map[string]interface{} `json:"numericIndexingSpec,omitempty\"` // Indexing spec for a numeric field. By default, only exact match queries will be supported for numeric fields. Setting the `numericIndexingSpec` allows range queries to be supported.

	ReadAccessType string `json:"readAccessType,omitempty\"` // Specifies who can view values of this field. See [Retrieve users as a non-administrator](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users#retrieve_users_non_admin) for more information. Note: It may take up to 24 hours for changes to this field to be reflected.

}

// JSON response template for List Schema operation in Directory API.
type Schemas struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	Schemas []Schema `json:"schemas,omitempty\"` // A list of UserSchema objects.

}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	Code int `json:"code,omitempty\"` // The status code, which should be an enum value of google.rpc.Code.

	Details []map[string]interface{} `json:"details,omitempty\"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.

	Message string `json:"message,omitempty\"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.

}

// JSON template for token resource in Directory API.
type Token struct {
	Anonymous bool `json:"anonymous,omitempty\"` // Whether the application is registered with Google. The value is `true` if the application has an anonymous Client ID.

	ClientId string `json:"clientId,omitempty\"` // The Client ID of the application the token is issued to.

	DisplayText string `json:"displayText,omitempty\"` // The displayable name of the application the token is issued to.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#token`.

	NativeApp bool `json:"nativeApp,omitempty\"` // Whether the token is issued to an installed application. The value is `true` if the application is installed to a desktop or mobile device.

	Scopes []string `json:"scopes,omitempty\"` // A list of authorization scopes the application is granted.

	UserKey string `json:"userKey,omitempty\"` // The unique ID of the user that issued the token.

}

// JSON response template for List tokens operation in Directory API.
type Tokens struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []Token `json:"items,omitempty\"` // A list of Token resources.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. This is always `admin#directory#tokenList`.

}

// The Directory API allows you to create and manage your account's users, user aliases, and user Google profile photos. For more information about common tasks, see the [User Accounts Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users.html) and the [User Aliases Developer's Guide](https://developers.google.com/workspace/admin/directory/v1/guides/manage-user-aliases.html).
type User struct {
	Addresses interface{} `json:"addresses,omitempty\"` // The list of the user's addresses. The maximum allowed data size for this field is 10KB.

	AgreedToTerms bool `json:"agreedToTerms,omitempty\"` // Output only. This property is `true` if the user has completed an initial login and accepted the Terms of Service agreement.

	Aliases []string `json:"aliases,omitempty\"` // Output only. The list of the user's alias email addresses.

	ArchivalTime string `json:"archivalTime,omitempty\"` // Output only. User's account archival time. (Read-only)

	Archived bool `json:"archived,omitempty\"` // Indicates if user is archived.

	ChangePasswordAtNextLogin bool `json:"changePasswordAtNextLogin,omitempty\"` // Indicates if the user is forced to change their password at next login. This setting doesn't apply when [the user signs in via a third-party identity provider](https://support.google.com/a/answer/60224).

	CreationTime time.Time `json:"creationTime,omitempty\"` // User's G Suite account creation time. (Read-only)

	CustomSchemas map[string]interface{} `json:"customSchemas,omitempty\"` // Custom fields of the user. The key is a `schema_name` and its values are `'field_name': 'field_value'`.

	CustomerId string `json:"customerId,omitempty\"` // Output only. The customer ID to [retrieve all account users](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users.html#get_all_users). You can use the alias `my_customer` to represent your account's `customerId`. As a reseller administrator, you can use the resold customer account's `customerId`. To get a `customerId`, use the account's primary domain in the `domain` parameter of a [users.list](https://developers.google.com/workspace/admin/directory/v1/reference/users/list) request.

	DeletionTime time.Time `json:"deletionTime,omitempty\"`

	Emails interface{} `json:"emails,omitempty\"` // The list of the user's email addresses. The maximum allowed data size for this field is 10KB. This excludes `publicKeyEncryptionCertificates`.

	Etag string `json:"etag,omitempty\"` // Output only. ETag of the resource.

	ExternalIds interface{} `json:"externalIds,omitempty\"` // The list of external IDs for the user, such as an employee or network ID. The maximum allowed data size for this field is 2KB.

	Gender interface{} `json:"gender,omitempty\"` // The user's gender. The maximum allowed data size for this field is 1KB.

	GuestAccountInfo GuestAccountInfo `json:"guestAccountInfo,omitempty\"` // Immutable. Additional guest-related metadata fields

	HashFunction string `json:"hashFunction,omitempty\"` // Stores the hash format of the `password` property. The following `hashFunction` values are allowed: * `MD5` - Accepts simple hex-encoded values. * `SHA-1` - Accepts simple hex-encoded values. * `crypt` - Compliant with the [C crypt library](https://en.wikipedia.org/wiki/Crypt_%28C%29). Supports the DES, MD5 (hash prefix `$1$`), SHA-256 (hash prefix `$5$`), and SHA-512 (hash prefix `$6$`) hash algorithms. If rounds are specified as part of the prefix, they must be 10,000 or fewer.

	Id string `json:"id,omitempty\"` // The unique ID for the user. A user `id` can be used as a user request URI's `userKey`.

	Ims interface{} `json:"ims,omitempty\"` // The list of the user's Instant Messenger (IM) accounts. A user account can have multiple ims properties. But, only one of these ims properties can be the primary IM contact. The maximum allowed data size for this field is 2KB.

	IncludeInGlobalAddressList bool `json:"includeInGlobalAddressList,omitempty\"` // Indicates if the user's profile is visible in the Google Workspace global address list when the contact sharing feature is enabled for the domain. For more information about excluding user profiles, see the [administration help center](https://support.google.com/a/answer/1285988).

	IpWhitelisted bool `json:"ipWhitelisted,omitempty\"` // If `true`, the user's IP address is subject to a deprecated IP address [`allowlist`](https://support.google.com/a/answer/60752) configuration.

	IsAdmin bool `json:"isAdmin,omitempty\"` // Output only. Indicates a user with super administrator privileges. The `isAdmin` property can only be edited in the [Make a user an administrator](https://developers.google.com/workspace/admin/directory/v1/guides/manage-users.html#make_admin) operation ( [makeAdmin](https://developers.google.com/workspace/admin/directory/v1/reference/users/makeAdmin.html) method). If edited in the user [insert](https://developers.google.com/workspace/admin/directory/v1/reference/users/insert.html) or [update](https://developers.google.com/workspace/admin/directory/v1/reference/users/update.html) methods, the edit is ignored by the API service.

	IsDelegatedAdmin bool `json:"isDelegatedAdmin,omitempty\"` // Output only. Indicates if the user is a delegated administrator. Delegated administrators are supported by the API but cannot create or undelete users, or make users administrators. These requests are ignored by the API service. Roles and privileges for administrators are assigned using the [Admin console](https://support.google.com/a/answer/33325).

	IsEnforcedIn2Sv bool `json:"isEnforcedIn2Sv,omitempty\"` // Output only. Is 2-step verification enforced (Read-only)

	IsEnrolledIn2Sv bool `json:"isEnrolledIn2Sv,omitempty\"` // Output only. Is enrolled in 2-step verification (Read-only)

	IsGuestUser bool `json:"isGuestUser,omitempty\"` // Immutable. Indicates if the inserted user is a guest.

	IsMailboxSetup bool `json:"isMailboxSetup,omitempty\"` // Output only. Indicates if the user's Google mailbox is created. This property is only applicable if the user has been assigned a Gmail license.

	Keywords interface{} `json:"keywords,omitempty\"` // The list of the user's keywords. The maximum allowed data size for this field is 1KB.

	Kind string `json:"kind,omitempty\"` // Output only. The type of the API resource. For Users resources, the value is `admin#directory#user`.

	Languages interface{} `json:"languages,omitempty\"` // The user's languages. The maximum allowed data size for this field is 1KB.

	LastLoginTime time.Time `json:"lastLoginTime,omitempty\"` // User's last login time. (Read-only)

	Locations interface{} `json:"locations,omitempty\"` // The user's locations. The maximum allowed data size for this field is 10KB.

	Name UserName `json:"name,omitempty\"` // Holds the given and family names of the user, and the read-only `fullName` value. The maximum number of characters in the `givenName` and in the `familyName` values is 60. In addition, name values support unicode/UTF-8 characters, and can contain spaces, letters (a-z), numbers (0-9), dashes (-), forward slashes (/), and periods (.). For more information about character usage rules, see the [administration help center](https://support.google.com/a/answer/9193374). Maximum allowed data size for this field is 1KB.

	NonEditableAliases []string `json:"nonEditableAliases,omitempty\"` // Output only. The list of the user's non-editable alias email addresses. These are typically outside the account's primary domain or sub-domain.

	Notes interface{} `json:"notes,omitempty\"` // Notes for the user.

	OrgUnitPath string `json:"orgUnitPath,omitempty\"` // The full path of the parent organization associated with the user. If the parent organization is the top-level, it is represented as a forward slash (`/`).

	Organizations interface{} `json:"organizations,omitempty\"` // The list of organizations the user belongs to. The maximum allowed data size for this field is 10KB.

	Password string `json:"password,omitempty\"` // User's password

	Phones interface{} `json:"phones,omitempty\"` // The list of the user's phone numbers. The maximum allowed data size for this field is 1KB.

	PosixAccounts interface{} `json:"posixAccounts,omitempty\"` // The list of [POSIX](https://www.opengroup.org/austin/papers/posix_faq.html) account information for the user.

	PrimaryEmail string `json:"primaryEmail,omitempty\"` // The user's primary email address. This property is required in a request to create a user account. The `primaryEmail` must be unique and cannot be an alias of another user.

	RecoveryEmail string `json:"recoveryEmail,omitempty\"` // Recovery email of the user.

	RecoveryPhone string `json:"recoveryPhone,omitempty\"` // Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: *+16506661212*.

	Relations interface{} `json:"relations,omitempty\"` // The list of the user's relationships to other users. The maximum allowed data size for this field is 2KB.

	SshPublicKeys interface{} `json:"sshPublicKeys,omitempty\"` // A list of SSH public keys.

	Suspended bool `json:"suspended,omitempty\"` // Indicates if user is suspended.

	SuspensionReason string `json:"suspensionReason,omitempty\"` // Output only. Has the reason a user account is suspended either by the administrator or by Google at the time of suspension. The property is returned only if the `suspended` property is `true`.

	SuspensionTime string `json:"suspensionTime,omitempty\"` // Output only. User's account suspension time. (Read-only)

	ThumbnailPhotoEtag string `json:"thumbnailPhotoEtag,omitempty\"` // Output only. ETag of the user's photo (Read-only)

	ThumbnailPhotoUrl string `json:"thumbnailPhotoUrl,omitempty\"` // Output only. The URL of the user's profile photo. The URL might be temporary or private.

	Websites interface{} `json:"websites,omitempty\"` // The user's websites. The maximum allowed data size for this field is 2KB.

}

// JSON template for About (notes) of a user in Directory API.
type UserAbout struct {
	ContentType string `json:"contentType,omitempty\"` // About entry can have a type which indicates the content type. It can either be plain or html. By default, notes contents are assumed to contain plain text.

	Value string `json:"value,omitempty\"` // Actual value of notes.

}

// JSON template for address.
type UserAddress struct {
	Country string `json:"country,omitempty\"` // Country.

	CountryCode string `json:"countryCode,omitempty\"` // Country code.

	CustomType string `json:"customType,omitempty\"` // Custom type.

	ExtendedAddress string `json:"extendedAddress,omitempty\"` // Extended Address.

	Formatted string `json:"formatted,omitempty\"` // Formatted address.

	Locality string `json:"locality,omitempty\"` // Locality.

	PoBox string `json:"poBox,omitempty\"` // Other parts of address.

	PostalCode string `json:"postalCode,omitempty\"` // Postal code.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary address. Only one entry could be marked as primary.

	Region string `json:"region,omitempty\"` // Region.

	SourceIsStructured bool `json:"sourceIsStructured,omitempty\"` // User supplied address was structured. Structured addresses are NOT supported at this time. You might be able to write structured addresses but any values will eventually be clobbered.

	StreetAddress string `json:"streetAddress,omitempty\"` // Street.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard values of that entry. For example address could be of home work etc. In addition to the standard type an entry can have a custom type and can take any value. Such type should have the CUSTOM value as type and also have a customType value.

}

// The Directory API manages aliases, which are alternative email addresses.
type UserAlias struct {
	Alias string `json:"alias,omitempty\"` // The alias email address.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Id string `json:"id,omitempty\"` // The unique ID for the user.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Alias resources, the value is `admin#directory#alias`.

	PrimaryEmail string `json:"primaryEmail,omitempty\"` // The user's primary email address.

}

// JSON template for a set of custom properties (i.e. all fields in a particular schema)
type UserCustomProperties struct {
}

// JSON template for an email.
type UserEmail struct {
	Address string `json:"address,omitempty\"` // Email id of the user.

	CustomType string `json:"customType,omitempty\"` // Custom Type.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary email. Only one entry could be marked as primary.

	PublicKeyEncryptionCertificates map[string]interface{} `json:"public_key_encryption_certificates,omitempty\"` // Public Key Encryption Certificates. Current limit: 1 per email address, and 5 per user.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example email could be of home, work etc. In addition to the standard type, an entry can have a custom type and can take any value Such types should have the CUSTOM value as type and also have a customType value.

}

// JSON template for an externalId entry.
type UserExternalId struct {
	CustomType string `json:"customType,omitempty\"` // Custom type.

	TypeValue string `json:"type,omitempty\"` // The type of the Id.

	Value string `json:"value,omitempty\"` // The value of the id.

}

type UserGender struct {
	AddressMeAs string `json:"addressMeAs,omitempty\"` // AddressMeAs. A human-readable string containing the proper way to refer to the profile owner by humans for example he/him/his or they/them/their.

	CustomGender string `json:"customGender,omitempty\"` // Custom gender.

	TypeValue string `json:"type,omitempty\"` // Gender.

}

// JSON template for instant messenger of an user.
type UserIm struct {
	CustomProtocol string `json:"customProtocol,omitempty\"` // Custom protocol.

	CustomType string `json:"customType,omitempty\"` // Custom type.

	Im string `json:"im,omitempty\"` // Instant messenger id.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary im. Only one entry could be marked as primary.

	Protocol string `json:"protocol,omitempty\"` // Protocol used in the instant messenger. It should be one of the values from ImProtocolTypes map. Similar to type it can take a CUSTOM value and specify the custom name in customProtocol field.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example instant messengers could be of home work etc. In addition to the standard type an entry can have a custom type and can take any value. Such types should have the CUSTOM value as type and also have a customType value.

}

// JSON template for a keyword entry.
type UserKeyword struct {
	CustomType string `json:"customType,omitempty\"` // Custom Type.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard type of that entry. For example keyword could be of type occupation or outlook. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.

	Value string `json:"value,omitempty\"` // Keyword.

}

// JSON template for a language entry.
type UserLanguage struct {
	CustomLanguage string `json:"customLanguage,omitempty\"` // Other language. User can provide their own language name if there is no corresponding ISO 639 language code. If this is set, `languageCode` can't be set.

	LanguageCode string `json:"languageCode,omitempty\"` // ISO 639 string representation of a language. See [Language Codes](/admin-sdk/directory/v1/languages) for the list of supported codes. Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior. Illegal values cause `SchemaException`. If this is set, `customLanguage` can't be set.

	Preference string `json:"preference,omitempty\"` // Optional. If present, controls whether the specified `languageCode` is the user's preferred language. If `customLanguage` is set, this can't be set. Allowed values are `preferred` and `not_preferred`.

}

// JSON template for a location entry.
type UserLocation struct {
	Area string `json:"area,omitempty\"` // Required. Textual location. This is most useful for display purposes to concisely describe the location. For example 'Mountain View, CA', 'Near Seattle', 'US-NYC-9TH 9A209A.''

	BuildingId string `json:"buildingId,omitempty\"` // Building Identifier.

	CustomType string `json:"customType,omitempty\"` // Custom Type.

	DeskCode string `json:"deskCode,omitempty\"` // Most specific textual code of individual desk location.

	FloorName string `json:"floorName,omitempty\"` // Floor name/number.

	FloorSection string `json:"floorSection,omitempty\"` // Floor section. More specific location within the floor. For example if a floor is divided into sections 'A', 'B' and 'C' this field would identify one of those values.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example location could be of types default and desk. In addition to standard type an entry can have a custom type and can give it any name. Such types should have 'custom' as type and also have a customType value.

}

type UserMakeAdmin struct {
	Status bool `json:"status,omitempty\"` // Indicates the administrator status of the user.

}

type UserName struct {
	DisplayName string `json:"displayName,omitempty\"` // The user's display name. Limit: 256 characters.

	FamilyName string `json:"familyName,omitempty\"` // The user's last name. Required when creating a user account.

	FullName string `json:"fullName,omitempty\"` // The user's full name formed by concatenating the first and last name values.

	GivenName string `json:"givenName,omitempty\"` // The user's first name. Required when creating a user account.

}

// JSON template for an organization entry.
type UserOrganization struct {
	CostCenter string `json:"costCenter,omitempty\"` // The cost center of the users department.

	CustomType string `json:"customType,omitempty\"` // Custom type.

	Department string `json:"department,omitempty\"` // Department within the organization.

	Description string `json:"description,omitempty\"` // Description of the organization.

	Domain string `json:"domain,omitempty\"` // The domain to which the organization belongs to.

	FullTimeEquivalent int `json:"fullTimeEquivalent,omitempty\"` // The full-time equivalent millipercent within the organization (100000 = 100%).

	Location string `json:"location,omitempty\"` // Location of the organization. This need not be fully qualified address.

	Name string `json:"name,omitempty\"` // Name of the organization

	Primary bool `json:"primary,omitempty\"` // If it user's primary organization.

	Symbol string `json:"symbol,omitempty\"` // Symbol of the organization.

	Title string `json:"title,omitempty\"` // Title (designation) of the user in the organization.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example organization could be of school work etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a CustomType value.

}

// JSON template for a phone entry.
type UserPhone struct {
	CustomType string `json:"customType,omitempty\"` // Custom Type.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary phone or not.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example phone could be of home_fax work mobile etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.

	Value string `json:"value,omitempty\"` // Phone number.

}

type UserPhoto struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Height int `json:"height,omitempty\"` // Height of the photo in pixels.

	Id string `json:"id,omitempty\"` // The ID the API uses to uniquely identify the user.

	Kind string `json:"kind,omitempty\"` // The type of the API resource. For Photo resources, this is `admin#directory#user#photo`.

	MimeType string `json:"mimeType,omitempty\"` // The MIME type of the photo. Allowed values are `JPEG`, `PNG`, `GIF`, `BMP`, `TIFF`, and web-safe base64 encoding.

	PhotoData string `json:"photoData,omitempty\"` // The user photo's upload data in [web-safe Base64](https://en.wikipedia.org/wiki/Base64#URL_applications) format in bytes. This means: * The slash (/) character is replaced with the underscore (_) character. * The plus sign (+) character is replaced with the hyphen (-) character. * The equals sign (=) character is replaced with the asterisk (*). * For padding, the period (.) character is used instead of the RFC-4648 baseURL definition which uses the equals sign (=) for padding. This is done to simplify URL-parsing. * Whatever the size of the photo being uploaded, the API downsizes it to 96x96 pixels.

	PrimaryEmail string `json:"primaryEmail,omitempty\"` // The user's primary email address.

	Width int `json:"width,omitempty\"` // Width of the photo in pixels.

}

// JSON template for a POSIX account entry.
type UserPosixAccount struct {
	AccountId string `json:"accountId,omitempty\"` // A POSIX account field identifier.

	Gecos string `json:"gecos,omitempty\"` // The GECOS (user information) for this account.

	Gid uint64 `json:"gid,omitempty\"` // The default group ID.

	HomeDirectory string `json:"homeDirectory,omitempty\"` // The path to the home directory for this account.

	OperatingSystemType string `json:"operatingSystemType,omitempty\"` // The operating system type for this account.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary account within the SystemId.

	Shell string `json:"shell,omitempty\"` // The path to the login shell for this account.

	SystemId string `json:"systemId,omitempty\"` // System identifier for which account Username or Uid apply to.

	Uid uint64 `json:"uid,omitempty\"` // The POSIX compliant user ID.

	Username string `json:"username,omitempty\"` // The username of the account.

}

// JSON template for a relation entry.
type UserRelation struct {
	CustomType string `json:"customType,omitempty\"` // Custom Type.

	TypeValue string `json:"type,omitempty\"` // The relation of the user. Some of the possible values are mother father sister brother manager assistant partner.

	Value string `json:"value,omitempty\"` // The name of the relation.

}

// JSON template for a POSIX account entry.
type UserSshPublicKey struct {
	ExpirationTimeUsec int64 `json:"expirationTimeUsec,omitempty\"` // An expiration time in microseconds since epoch.

	Fingerprint string `json:"fingerprint,omitempty\"` // A SHA-256 fingerprint of the SSH public key. (Read-only)

	Key string `json:"key,omitempty\"` // An SSH public key.

}

type UserUndelete struct {
	OrgUnitPath string `json:"orgUnitPath,omitempty\"` // OrgUnit of User

}

// JSON template for a website entry.
type UserWebsite struct {
	CustomType string `json:"customType,omitempty\"` // Custom Type.

	Primary bool `json:"primary,omitempty\"` // If this is user's primary website or not.

	TypeValue string `json:"type,omitempty\"` // Each entry can have a type which indicates standard types of that entry. For example website could be of home work blog etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.

	Value string `json:"value,omitempty\"` // Website.

}

type Users struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // Kind of resource this is.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access next page of this result. The page token is only valid for three days.

	TriggerEvent string `json:"trigger_event,omitempty\"` // Event that triggered this response (only used in case of Push Response)

	Users []User `json:"users,omitempty\"` // A list of user objects.

}

// The Directory API allows you to view, generate, and invalidate backup verification codes for a user.
type VerificationCode struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Kind string `json:"kind,omitempty\"` // The type of the resource. This is always `admin#directory#verificationCode`.

	UserId string `json:"userId,omitempty\"` // The obfuscated unique ID of the user.

	VerificationCode string `json:"verificationCode,omitempty\"` // A current verification code for the user. Invalidated or used verification codes are not returned as part of the result.

}

// JSON response template for list verification codes operation in Directory API.
type VerificationCodes struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []VerificationCode `json:"items,omitempty\"` // A list of verification code resources.

	Kind string `json:"kind,omitempty\"` // The type of the resource. This is always `admin#directory#verificationCodesList`.

}
