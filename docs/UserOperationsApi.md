# \UserOperationsApi

All URIs are relative to *https://forgeapi.puppet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](UserOperationsApi.md#GetUser) | **Get** /v3/users/{user_slug} | Fetch user
[**GetUsers**](UserOperationsApi.md#GetUsers) | **Get** /v3/users | List users



## GetUser

> User GetUser(ctx, userSlug, optional)

Fetch user

Returns data for a single User resource identified by the user's `slug` value. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string**| Unique textual identifier (slug) of the User resource to retrieve | 
 **optional** | ***GetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse200 GetUsers(ctx, optional)

List users

Provides information about Puppet Forge user accounts. By default, results are returned in alphabetical order by username and paginated with 20 users per page. It's also possible to sort by number of published releases, total download counts for all the user's modules, or by the date of the user's latest release. All parameters are optional. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The numbers of items to return per page | [default to 20]
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **sortBy** | **optional.String**| Desired order in which to return results | 
 **withHtml** | **optional.Bool**| Render markdown files (README, REFERENCE, etc.) to HTML before returning results | 
 **includeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to include in response object, only applies to fields marked &#39;optional&#39; | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| List of top level keys to exclude from response object | 
 **ifModifiedSince** | **optional.String**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

