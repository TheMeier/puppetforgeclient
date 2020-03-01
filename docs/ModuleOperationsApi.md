# \ModuleOperationsApi

All URIs are relative to *https://forgeapi.puppet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteModule**](ModuleOperationsApi.md#DeleteModule) | **Delete** /v3/modules/{module_slug} | Delete module
[**DeprecateModule**](ModuleOperationsApi.md#DeprecateModule) | **Patch** /v3/modules/{module_slug} | Deprecate module
[**GetModule**](ModuleOperationsApi.md#GetModule) | **Get** /v3/modules/{module_slug} | Fetch module
[**GetModules**](ModuleOperationsApi.md#GetModules) | **Get** /v3/modules | List modules



## DeleteModule

> DeleteModule(ctx, moduleSlug, reason)

Delete module

Perform a soft delete of a module, identified by the module's `slug` value. The module's releases will still be available for direct download via their associated `/v3/files` endpoints, but the module will no longer be readily available through the Forge web interface. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleSlug** | **string**| Unique textual identifier (slug) of the Module resource to retrieve | 
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


## DeprecateModule

> DeprecateModule(ctx, moduleSlug, deprecationRequest)

Deprecate module

Mark a module, identified by the module's `slug` value, as \"deprecated\". Deprecated modules are still visible on the Forge website, but users are directed to strongly consider alternate modules. **Because the deprecate action is intended to be one-way, there is no operation for undeprecating a module.** 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleSlug** | **string**| Unique textual identifier (slug) of the module to deprecate | 
**deprecationRequest** | [**DeprecationRequest**](DeprecationRequest.md)| Action and params for patch operation. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModule

> Module GetModule(ctx, moduleSlug, optional)

Fetch module

Returns data for a single Module resource identified by the module's `slug` value. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleSlug** | **string**| Unique textual identifier (slug) of the Module resource to retrieve | 
 **optional** | ***GetModuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetModuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 

### Return type

[**Module**](Module.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModules

> InlineResponse2003 GetModules(ctx, optional)

List modules

Returns a list of modules meeting the specified search criteria and filters. Results are paginated. All of the parameters are optional. To publish or delete a Release resource, see [Release operations](#tag/Release-Operations). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetModulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetModulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The numbers of items to return per page | [default to 20]
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **sortBy** | **optional.String**| Desired order in which to return results, note that &#x60;rank&#x60; is only available (and becomes the default) when a value has been provided for the &#x60;query&#x60; param  | 
 **query** | **optional.String**| Freeform textual search query | 
 **tag** | **optional.String**| Constrain results to modules that have the given tag | 
 **owner** | **optional.String**| Constrain results to modules owned by the given user | 
 **withTasks** | **optional.Bool**| Constrain results to modules that include [Bolt tasks](https://puppet.com/docs/bolt/latest/writing_tasks_and_plans.html)  | 
 **withPlans** | **optional.Bool**| Constrain results to modules that include [Bolt plans](https://puppet.com/docs/bolt/latest/writing_tasks_and_plans.html)  | 
 **withPdk** | **optional.Bool**| Constrain results to modules created with [Puppet Development Kit](https://puppet.com/docs/pdk/latest/pdk.html)  | 
 **endorsements** | [**optional.Interface of []string**](string.md)| Constrain results to modules with at least one of the given endorsements | 
 **operatingsystem** | **optional.String**| Constrain results to modules that explicitly support the given operating system, see [metadata.json documentation](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-operating-system-compatibility) for more information about possible values, see also the &#x60;only_latest&#x60; param  | 
 **peRequirement** | **optional.String**| Constrain results to modules that explicitly list a Puppet Enterprise version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions), see also the &#x60;only_latest&#x60; param  | 
 **puppetRequirement** | **optional.String**| Constrain results to modules that explicitly list a Puppet version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions), see also the &#x60;only_latest&#x60; param  | 
 **withMinimumScore** | **optional.Int32**| Constrain results to modules with at least one release that has a minimum validation_score greater than or equal to the provided integer, see also the &#x60;only_latest&#x60; param  | 
 **moduleGroups** | [**optional.Interface of []string**](string.md)| Constrain results to modules in any of the given module groups | 
 **showDeleted** | **optional.Bool**| Include modules where all released versions have been marked as deleted | 
 **hideDeprecated** | **optional.Bool**| Exclude deprecated modules from the results. Treats any value as true. | 
 **onlyLatest** | **optional.Bool**| When checking modules against the &#x60;operatingsystem&#x60;, &#x60;pe_requirement&#x60;, or &#x60;puppet_requirement&#x60; filters, only consider the compatibility information for the latest release of the module, where \&quot;latest\&quot; means: the module release with the highest precedence, non-pre-release version (see also: [semver.org](https://semver.org/#spec-item-11))  | [default to false]
 **slugs** | [**optional.Interface of []string**](string.md)| A comma-separated list of specific modules to retrieve. Module names should be specified in the hyphenated \&quot;slug\&quot; format (&#x60;author_name-module_name&#x60;). If at least one of the requested modules is not found, this endpoint returns a 404 error, with the slugs of the missing module(s) in the body. This parameter is incompatible with the &#x60;endorsements&#x60;, &#x60;module_groups&#x60;, &#x60;only_latest&#x60;, &#x60;operatingsystem&#x60;, &#x60;owner&#x60;, &#x60;pe_requirement&#x60;, &#x60;puppet_requirement&#x60;, &#x60;query&#x60;, &#x60;show_deleted&#x60;, &#x60;hide_deprecated&#x60;, &#x60;tag&#x60;, &#x60;with_minimum_score&#x60;, &#x60;with_pdk&#x60;, &#x60;with_tasks&#x60;, and &#x60;with_plans&#x60; parameters. These will produce 400 errors when included. Note that any matching soft-deleted modules (modules with no associated releases) will always be included in the results alongside modules with releases.  | 
 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 
 **supported** | **optional.Bool**| Constrain results to [Supported modules](https://forge.puppet.com/supported), see \&quot;endorsements\&quot; param instead  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

