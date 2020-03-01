/*
 * Puppet Forge v3 API
 *
 * ## Introduction The Puppet Forge API (hereafter referred to as the Forge API) provides quick access to all the data on the Puppet Forge via a RESTful interface. Using the Forge API, you can write scripts and tools that interact with the Puppet Forge website.  The Forge API's current version is `v3`. It is considered regression-stable, meaning that the returned data is guaranteed to include the fields described in the schemas on this page; however, additional data might be added in the future and clients must ignore any properties they do not recognize.  ## OpenAPI Specification The Puppet Forge v3 API is described by an [OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md) formatted specification file. The most up-to-date version of this specification file can be accessed at [https://forgeapi.puppet.com/v3/openapi.json](/v3/openapi.json).  ## Features * The API is accessed over HTTPS via the `forgeapi.puppet.com` domain. All data is returned in JSON   format. * Blank fields are included as `null`. * Nested resources may use an abbreviated representation. A link to the full representation for the   resource is always included. * All timestamps in JSON responses are returned in ISO 8601 format: `YYYY-MM-DD HH:MM:SS ±HHMM`. * The HTTP response headers include caching hints for conditional requests.  ## Concepts and Terminology * **Module**: Modules are self-contained bundles of code and data with a specific directory structure. Modules are identified by a combination of the author's username and the module's name, separated by a hyphen. For example: `puppetlabs-apache` * **Release**: A single, specific version of a module is called a Release. Releases are identified by a combination of the module identifier (see above) and the Release version, separated by a hyphen. For example: `puppetlabs-apache-4.0.0`  ## Errors The Puppet Forge API follows [RFC 2616](https://tools.ietf.org/html/rfc2616) and [RFC 6585](https://tools.ietf.org/html/rfc6585).  Error responses are served with a `4xx` or `5xx` status code, and are sent as a JSON document with a content type of `application/json`. The error document contains the following top-level keys and values:    * `message`: a string value that summarizes the problem   * `errors`: a list (array) of strings containing additional details describing the underlying cause(s) of the     failure  An example error response is shown below:  ```json {   \"message\": \"400 Bad Request\",   \"errors\": [     \"Cannot parse request body as JSON\"   ] } ```  ## User-Agent Required All API requests must include a valid `User-Agent` header. Requests with no `User-Agent` header will be rejected. The `User-Agent` header helps identify your application or library, so we can communicate with you if necessary. If your use of the API is informal or personal, we recommend using your username as the value for the `User-Agent` header.  User-Agent headers are a list of one or more product descriptions, generally taking this form:  ``` <name-without-spaces>/<version> (comments) ```  For example, the following are all useful User-Agent values:  ``` MyApplication/0.0.0 Her/0.6.8 Faraday/0.8.8 Ruby/1.9.3-p194 (i386-linux) My-Library-Name/1.2.4 myusername ``` 
 *
 * API version: 18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package puppetforgeclient
// Release struct for Release
type Release struct {
	// Relative URL for this Release resource
	Uri string `json:"uri,omitempty"`
	// Unique textual identifier for this Release resource
	Slug string `json:"slug,omitempty"`
	Module ModuleAbbreviated `json:"module,omitempty"`
	Version string `json:"version,omitempty"`
	// Verbatim contents of release's `metadata.json` file
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// List of tags applied to this module release (derived from `metadata.json`)
	Tags []string `json:"tags,omitempty"`
	// Indicates whether or not this specific module release is supported, this value is no longer relevant, please see the `endorsement` property of the parent module instead 
	Supported bool `json:"supported,omitempty"`
	// Indicates whether or not this module release's metadata contains [Puppet Development Kit](https://puppet.com/docs/pdk/latest/pdk.html) related keys/values 
	Pdk bool `json:"pdk,omitempty"`
	// Numeric value representing module release's \"Quality score\", see [module scoring](https://forge.puppet.com/about/scoring) to learn more 
	ValidationScore int32 `json:"validation_score,omitempty"`
	// Relative or absolute URL to download this release's tarball
	FileUri string `json:"file_uri,omitempty"`
	// Size of this release's tarball in bytes
	FileSize int32 `json:"file_size,omitempty"`
	// MD5 checksum for this release's tarball, can be used to verify successful download
	FileMd5 string `json:"file_md5,omitempty"`
	// SHA-256 checksum for this release's tarball, can be used to verify successful download
	FileSha256 string `json:"file_sha256,omitempty"`
	// Number of times this release has been downloaded through the Forge
	Downloads int32 `json:"downloads,omitempty"`
	// Contents of this release's README file, optionally converted to HTML format if possible (see `with_html` param)
	Readme string `json:"readme,omitempty"`
	// Contents of this release's CHANGELOG file, optionally converted to HTML format if possible (see `with_html` param)
	Changelog string `json:"changelog,omitempty"`
	// Contents of this release's LICENSE file, optionally converted to HTML format if possible (see `with_html` param)
	License string `json:"license,omitempty"`
	// Contents of this release's REFERENCE file, optionally converted to HTML format if possible (see `with_html` param)
	Reference string `json:"reference,omitempty"`
	// (OPTIONAL) Automatically generated documentation for custom Types included in this module release, see `reference` key instead (This field is not included in response by default, see `include_fields` param for more info) 
	Docs map[string]interface{} `json:"docs,omitempty"`
	// Metadata for any Bolt tasks included in this module release
	Tasks []ReleaseTask `json:"tasks,omitempty"`
	// Links to metadata for Bolt plans included in this module release
	Plans []ReleasePlanAbbreviated `json:"plans,omitempty"`
	// Date and time this Release resource was created
	CreatedAt string `json:"created_at,omitempty"`
	// Date and time this Release resource was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Date and time this Release resource was marked as deleted
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Author-provided reason that this module release was marked as deleted
	DeletedFor *string `json:"deleted_for,omitempty"`
}