# Module

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Relative URL for this Module resource | [optional] 
**Slug** | **string** | Unique textual identifier for this Module resource | [optional] 
**Name** | **string** | Base name of this module (without username/namespace) | [optional] 
**Downloads** | **int32** | Total download count for all released versions of this module | [optional] 
**CreatedAt** | **string** | Date and time this Module resource was created | [optional] 
**UpdatedAt** | **string** | Date and time this Module resource was last modified | [optional] 
**DeprecatedAt** | Pointer to **string** | Date and time this Module resource was marked as deprecated by the owner | [optional] 
**DeprecatedFor** | Pointer to **string** | Reason provided by the owner for marking this module as deprecated | [optional] 
**SupersededBy** | [**ModuleMinimal**](ModuleMinimal.md) |  | [optional] 
**Supported** | **bool** | Whether or not this module is a [Supported module](https://forge.puppet.com/supported), see &#x60;endorsement&#x60; field instead | [optional] 
**Endorsement** | Pointer to **string** | Indicates whether a module is endorsed through the Supported, Approved or Partner Supported programs, or &#x60;null&#x60; if not endorsed | [optional] 
**ModuleGroup** | **string** | Indicates whether or not a module is licensed for use by Puppet Enterprise customers only, indicated by a value of &#x60;pe_only&#x60;, for all other modules this value will be &#x60;base&#x60; | [optional] 
**Owner** | [**UserAbbreviated**](UserAbbreviated.md) |  | [optional] 
**CurrentRelease** | [**Release**](Release.md) |  | [optional] 
**Releases** | [**[]ReleaseAbbreviated**](ReleaseAbbreviated.md) | Array of abbreviated representations of all published Releases of this module | [optional] 
**FeedbackScore** | **int32** | Numeric score representing the average rating this module has received from Forge users | [optional] 
**HomepageUrl** | **string** | Author-provided URL for this module&#39;s homepage | [optional] 
**IssuesUrl** | **string** | Author-provided URL for reporting issues about this module | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


