/*
 * Puppet Forge v3 API
 *
 * ## Introduction The Puppet Forge API (hereafter referred to as the Forge API) provides quick access to all the data on the Puppet Forge via a RESTful interface. Using the Forge API, you can write scripts and tools that interact with the Puppet Forge website.  The Forge API's current version is `v3`. It is considered regression-stable, meaning that the returned data is guaranteed to include the fields described in the schemas on this page; however, additional data might be added in the future and clients must ignore any properties they do not recognize.  ## OpenAPI Specification The Puppet Forge v3 API is described by an [OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md) formatted specification file. The most up-to-date version of this specification file can be accessed at [https://forgeapi.puppet.com/v3/openapi.json](/v3/openapi.json).  ## Features * The API is accessed over HTTPS via the `forgeapi.puppet.com` domain. All data is returned in JSON   format. * Blank fields are included as `null`. * Nested resources may use an abbreviated representation. A link to the full representation for the   resource is always included. * All timestamps in JSON responses are returned in ISO 8601 format: `YYYY-MM-DD HH:MM:SS ±HHMM`. * The HTTP response headers include caching hints for conditional requests.  ## Concepts and Terminology * **Module**: Modules are self-contained bundles of code and data with a specific directory structure. Modules are identified by a combination of the author's username and the module's name, separated by a hyphen. For example: `puppetlabs-apache` * **Release**: A single, specific version of a module is called a Release. Releases are identified by a combination of the module identifier (see above) and the Release version, separated by a hyphen. For example: `puppetlabs-apache-4.0.0`  ## Errors The Puppet Forge API follows [RFC 2616](https://tools.ietf.org/html/rfc2616) and [RFC 6585](https://tools.ietf.org/html/rfc6585).  Error responses are served with a `4xx` or `5xx` status code, and are sent as a JSON document with a content type of `application/json`. The error document contains the following top-level keys and values:    * `message`: a string value that summarizes the problem   * `errors`: a list (array) of strings containing additional details describing the underlying cause(s) of the     failure  An example error response is shown below:  ```json {   \"message\": \"400 Bad Request\",   \"errors\": [     \"Cannot parse request body as JSON\"   ] } ```  ## User-Agent Required All API requests must include a valid `User-Agent` header. Requests with no `User-Agent` header will be rejected. The `User-Agent` header helps identify your application or library, so we can communicate with you if necessary. If your use of the API is informal or personal, we recommend using your username as the value for the `User-Agent` header.  User-Agent headers are a list of one or more product descriptions, generally taking this form:  ``` <name-without-spaces>/<version> (comments) ```  For example, the following are all useful User-Agent values:  ``` MyApplication/0.0.0 Her/0.6.8 Faraday/0.8.8 Ruby/1.9.3-p194 (i386-linux) My-Library-Name/1.2.4 myusername ``` 
 *
 * API version: 18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package puppetforgeclient

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ReleaseOperationsApiService ReleaseOperationsApi service
type ReleaseOperationsApiService service

/*
AddRelease Create module release
Publish a new module or new release of an existing module 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param inlineObject
@return ReleaseMinimal
*/
func (a *ReleaseOperationsApiService) AddRelease(ctx _context.Context, inlineObject InlineObject) (ReleaseMinimal, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReleaseMinimal
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &inlineObject
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization: Bearer &lt;api_key&gt;"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 201 {
			var v ReleaseMinimal
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v InlineResponse409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DeleteRelease Delete module release
Perform a soft delete of a module release, identified by the module release&#39;s &#x60;slug&#x60; value. The release will still be available for direct download from it&#39;s &#x60;/v3/files&#x60; endpoint, but will not be readily available in the Forge web interface. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param releaseSlug Unique textual identifier (slug) of the module release to delete
 * @param reason Reason for deletion
*/
func (a *ReleaseOperationsApiService) DeleteRelease(ctx _context.Context, releaseSlug string, reason string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases/{release_slug}"
	localVarPath = strings.Replace(localVarPath, "{"+"release_slug"+"}", _neturl.QueryEscape(parameterToString(releaseSlug, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("reason", parameterToString(reason, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization: Bearer &lt;api_key&gt;"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
GetFile Download module release
Download a module release tarball 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param filename Module release filename to be downloaded (e.g. \"puppetlabs-apache-2.0.0.tar.gz\")
@return *os.File
*/
func (a *ReleaseOperationsApiService) GetFile(ctx _context.Context, filename string) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/files/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", _neturl.QueryEscape(parameterToString(filename, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v *os.File
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetReleaseOpts Optional parameters for the method 'GetRelease'
type GetReleaseOpts struct {
    WithHtml optional.Bool
    IncludeFields optional.Interface
    ExcludeFields optional.Interface
    IfModifiedSince optional.String
}

/*
GetRelease Fetch module release
Returns data for a single module Release resource identified by the module release&#39;s &#x60;slug&#x60; value. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param releaseSlug Unique textual identifier (slug) of the module Release resource to retrieve
 * @param optional nil or *GetReleaseOpts - Optional Parameters:
 * @param "WithHtml" (optional.Bool) -  Render markdown files (README, REFERENCE, etc.) to HTML before returning results
 * @param "IncludeFields" (optional.Interface of []string) -  List of top level keys to include in response object, only applies to fields marked 'optional'
 * @param "ExcludeFields" (optional.Interface of []string) -  List of top level keys to exclude from response object
 * @param "IfModifiedSince" (optional.String) - 
@return Release
*/
func (a *ReleaseOperationsApiService) GetRelease(ctx _context.Context, releaseSlug string, localVarOptionals *GetReleaseOpts) (Release, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Release
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases/{release_slug}"
	localVarPath = strings.Replace(localVarPath, "{"+"release_slug"+"}", _neturl.QueryEscape(parameterToString(releaseSlug, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.WithHtml.IsSet() {
		localVarQueryParams.Add("with_html", parameterToString(localVarOptionals.WithHtml.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeFields.IsSet() {
		localVarQueryParams.Add("include_fields", parameterToString(localVarOptionals.IncludeFields.Value(), "space"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "space"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Release
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetReleasePlan Fetch module release plan
Returns a summary of the given plan from the module release identified by the &#x60;release_slug&#x60; name. The &#x60;release_slug&#x60; is composed of the hyphenated module author, name, and version number (example: &#x60;puppetlabs-lvm-1.4.0&#x60;). The &#x60;plan_name&#x60; should be the full name including the module name (example: &#x60;lvm::expand&#x60;). 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param releaseSlug Unique textual identifier (slug) of the module Release resource to retrieve
 * @param planName Full name of desired plan, including module name and namespaces
@return ReleasePlan
*/
func (a *ReleaseOperationsApiService) GetReleasePlan(ctx _context.Context, releaseSlug string, planName string) (ReleasePlan, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReleasePlan
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases/{release_slug}/plans/{plan_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"release_slug"+"}", _neturl.QueryEscape(parameterToString(releaseSlug, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"plan_name"+"}", _neturl.QueryEscape(parameterToString(planName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ReleasePlan
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetReleasePlans List module release plans
Returns a paginated list of all plans from the module release identified by the &#x60;release_slug&#x60; name. The &#x60;release_slug&#x60; is composed of the hyphenated module author, name, and version number (example: &#x60;puppetlabs-lvm-1.4.0&#x60;). 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param releaseSlug Unique textual identifier (slug) of the module Release resource to retrieve
@return InlineResponse2002
*/
func (a *ReleaseOperationsApiService) GetReleasePlans(ctx _context.Context, releaseSlug string) (InlineResponse2002, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2002
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases/{release_slug}/plans"
	localVarPath = strings.Replace(localVarPath, "{"+"release_slug"+"}", _neturl.QueryEscape(parameterToString(releaseSlug, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v InlineResponse2002
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetReleasesOpts Optional parameters for the method 'GetReleases'
type GetReleasesOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    SortBy optional.String
    Module optional.String
    Owner optional.String
    WithPdk optional.Bool
    Operatingsystem optional.String
    PeRequirement optional.String
    PuppetRequirement optional.String
    ModuleGroups optional.Interface
    ShowDeleted optional.Bool
    HideDeprecated optional.Bool
    WithHtml optional.Bool
    IncludeFields optional.Interface
    ExcludeFields optional.Interface
    IfModifiedSince optional.String
    Supported optional.Bool
}

/*
GetReleases List module releases
Returns a list of module releases meeting the specified search criteria and filters. Results are paginated. All of the parameters are optional. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetReleasesOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  The numbers of items to return per page
 * @param "Offset" (optional.Int32) -  The number of items to skip before starting to collect the result set
 * @param "SortBy" (optional.String) -  Desired order in which to return results 
 * @param "Module" (optional.String) -  Constrain results to releases for a given module
 * @param "Owner" (optional.String) -  Constrain results to releases for any modules owned by the given user
 * @param "WithPdk" (optional.Bool) -  Constrain results to module releases created with [Puppet Development Kit](https://puppet.com/docs/pdk/latest/pdk.html) 
 * @param "Operatingsystem" (optional.String) -  Constrain results to module releases that explicitly support the given operating system, see [metadata.json documentation](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-operating-system-compatibility) for more information about possible values 
 * @param "PeRequirement" (optional.String) -  Constrain results to module releases that explicitly list a Puppet Enterprise version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions) 
 * @param "PuppetRequirement" (optional.String) -  Constrain results to module releases that explicitly list a Puppet version requirement in the given [semantic versioning range](https://puppet.com/docs/puppet/latest/modules_metadata.html#specifying-versions) 
 * @param "ModuleGroups" (optional.Interface of []string) -  Constrain results to releases for modules in any of the given module groups
 * @param "ShowDeleted" (optional.Bool) -  Include releases marked as \"deleted\" in results
 * @param "HideDeprecated" (optional.Bool) -  Exclude releases of deprecated modules from the results. Treats any value as true.
 * @param "WithHtml" (optional.Bool) -  Render markdown files (README, REFERENCE, etc.) to HTML before returning results
 * @param "IncludeFields" (optional.Interface of []string) -  List of top level keys to include in response object, only applies to fields marked 'optional'
 * @param "ExcludeFields" (optional.Interface of []string) -  List of top level keys to exclude from response object
 * @param "IfModifiedSince" (optional.String) - 
 * @param "Supported" (optional.Bool) -  Constrain results to releases for [Supported modules](https://forge.puppet.com/supported) 
@return InlineResponse2001
*/
func (a *ReleaseOperationsApiService) GetReleases(ctx _context.Context, localVarOptionals *GetReleasesOpts) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/releases"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Module.IsSet() {
		localVarQueryParams.Add("module", parameterToString(localVarOptionals.Module.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Owner.IsSet() {
		localVarQueryParams.Add("owner", parameterToString(localVarOptionals.Owner.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WithPdk.IsSet() {
		localVarQueryParams.Add("with_pdk", parameterToString(localVarOptionals.WithPdk.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Operatingsystem.IsSet() {
		localVarQueryParams.Add("operatingsystem", parameterToString(localVarOptionals.Operatingsystem.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PeRequirement.IsSet() {
		localVarQueryParams.Add("pe_requirement", parameterToString(localVarOptionals.PeRequirement.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PuppetRequirement.IsSet() {
		localVarQueryParams.Add("puppet_requirement", parameterToString(localVarOptionals.PuppetRequirement.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ModuleGroups.IsSet() {
		localVarQueryParams.Add("module_groups", parameterToString(localVarOptionals.ModuleGroups.Value(), "space"))
	}
	if localVarOptionals != nil && localVarOptionals.ShowDeleted.IsSet() {
		localVarQueryParams.Add("show_deleted", parameterToString(localVarOptionals.ShowDeleted.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HideDeprecated.IsSet() {
		localVarQueryParams.Add("hide_deprecated", parameterToString(localVarOptionals.HideDeprecated.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WithHtml.IsSet() {
		localVarQueryParams.Add("with_html", parameterToString(localVarOptionals.WithHtml.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeFields.IsSet() {
		localVarQueryParams.Add("include_fields", parameterToString(localVarOptionals.IncludeFields.Value(), "space"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "space"))
	}
	if localVarOptionals != nil && localVarOptionals.Supported.IsSet() {
		localVarQueryParams.Add("supported", parameterToString(localVarOptionals.Supported.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		localVarHeaderParams["If-Modified-Since"] = parameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v InlineResponse2001
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
