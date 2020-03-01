# \ReleaseOperationsApi

All URIs are relative to *https://forgeapi.puppet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRelease**](ReleaseOperationsApi.md#AddRelease) | **Post** /v3/releases | Create module release
[**DeleteRelease**](ReleaseOperationsApi.md#DeleteRelease) | **Delete** /v3/releases/{release_slug} | Delete module release
[**GetFile**](ReleaseOperationsApi.md#GetFile) | **Get** /v3/files/{filename} | Download module release
[**GetRelease**](ReleaseOperationsApi.md#GetRelease) | **Get** /v3/releases/{release_slug} | Fetch module release
[**GetReleasePlan**](ReleaseOperationsApi.md#GetReleasePlan) | **Get** /v3/releases/{release_slug}/plans/{plan_name} | Fetch module release plan
[**GetReleasePlans**](ReleaseOperationsApi.md#GetReleasePlans) | **Get** /v3/releases/{release_slug}/plans | List module release plans
[**GetReleases**](ReleaseOperationsApi.md#GetReleases) | **Get** /v3/releases | List module releases



## AddRelease

> ReleaseMinimal AddRelease(ctx, inlineObject)

Create module release

Publish a new module or new release of an existing module 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**ReleaseMinimal**](ReleaseMinimal.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelease

> DeleteRelease(ctx, releaseSlug, reason)

Delete module release

Perform a soft delete of a module release, identified by the module release's `slug` value. The release will still be available for direct download from it's `/v3/files` endpoint, but will not be readily available in the Forge web interface. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseSlug** | **string**| Unique textual identifier (slug) of the module release to delete | 
**reason** | **string**| Reason for deletion | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> *os.File GetFile(ctx, filename)

Download module release

Download a module release tarball 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string**| Module release filename to be downloaded (e.g. \&quot;puppetlabs-apache-2.0.0.tar.gz\&quot;) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> Release GetRelease(ctx, releaseSlug, optional)

Fetch module release

Returns data for a single module Release resource identified by the module release's `slug` value. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseSlug** | **string**| Unique textual identifier (slug) of the module Release resource to retrieve | 
 **optional** | ***GetReleaseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReleaseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePlan

> ReleasePlan GetReleasePlan(ctx, releaseSlug, planName)

Fetch module release plan

Returns a summary of the given plan from the module release identified by the `release_slug` name. The `release_slug` is composed of the hyphenated module author, name, and version number (example: `puppetlabs-lvm-1.4.0`). The `plan_name` should be the full name including the module name (example: `lvm::expand`). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseSlug** | **string**| Unique textual identifier (slug) of the module Release resource to retrieve | 
**planName** | **string**| Full name of desired plan, including module name and namespaces | 

### Return type

[**ReleasePlan**](ReleasePlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePlans

> InlineResponse2002 GetReleasePlans(ctx, releaseSlug)

List module release plans

Returns a paginated list of all plans from the module release identified by the `release_slug` name. The `release_slug` is composed of the hyphenated module author, name, and version number (example: `puppetlabs-lvm-1.4.0`). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseSlug** | **string**| Unique textual identifier (slug) of the module Release resource to retrieve | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleases

> InlineResponse2001 GetReleases(ctx, optional)

List module releases

Returns a list of module releases meeting the specified search criteria and filters. Results are paginated. All of the parameters are optional. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReleasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The numbers of items to return per page | [default to 20]
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **sortBy** | **optional.String**| Desired order in which to return results  | 
 **module** | **optional.String**| Constrain results to releases for a given module | 
 **owner** | **optional.String**| Constrain results to releases for any modules owned by the given user | 
 **withPdk** | **optional.Bool**| Constrain results to module releases created with [Puppet Development Kit](https://puppet.com/docs/pdk/latest/pdk.html)  | 
 **operatingsystem** | **optional.String**| Constrain results to module releases that explicitly support the given operating system, see [metadata.json documentation](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-operating-system-compatibility) for more information about possible values  | 
 **peRequirement** | **optional.String**| Constrain results to module releases that explicitly list a Puppet Enterprise version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions)  | 
 **puppetRequirement** | **optional.String**| Constrain results to module releases that explicitly list a Puppet version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions)  | 
 **moduleGroups** | [**optional.Interface of []string**](string.md)| Constrain results to releases for modules in any of the given module groups | 
 **showDeleted** | **optional.Bool**| Include releases marked as \&quot;deleted\&quot; in results | 
 **hideDeprecated** | **optional.Bool**| Exclude releases of deprecated modules from the results. Treats any value as true. | 
 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 
 **supported** | **optional.Bool**| Constrain results to releases for [Supported modules](https://forge.puppet.com/supported)  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

