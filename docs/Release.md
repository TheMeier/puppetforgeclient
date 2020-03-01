# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Relative URL for this Release resource | [optional] 
**Slug** | **string** | Unique textual identifier for this Release resource | [optional] 
**Module** | [**ModuleAbbreviated**](ModuleAbbreviated.md) |  | [optional] 
**Version** | **string** |  | [optional] 
**Metadata** | [**map[string]interface{}**](.md) | Verbatim contents of release&#39;s &#x60;metadata.json&#x60; file | [optional] 
**Tags** | **[]string** | List of tags applied to this module release (derived from &#x60;metadata.json&#x60;) | [optional] 
**Supported** | **bool** | Indicates whether or not this specific module release is supported, this value is no longer relevant, please see the &#x60;endorsement&#x60; property of the parent module instead  | [optional] 
**Pdk** | **bool** | Indicates whether or not this module release&#39;s metadata contains [Puppet Development Kit](https://puppet.com/docs/pdk/latest/pdk.html) related keys/values  | [optional] 
**ValidationScore** | **int32** | Numeric value representing module release&#39;s \&quot;Quality score\&quot;, see [module scoring](https://forge.puppet.com/about/scoring) to learn more  | [optional] 
**FileUri** | **string** | Relative or absolute URL to download this release&#39;s tarball | [optional] 
**FileSize** | **int32** | Size of this release&#39;s tarball in bytes | [optional] 
**FileMd5** | **string** | MD5 checksum for this release&#39;s tarball, can be used to verify successful download | [optional] 
**FileSha256** | **string** | SHA-256 checksum for this release&#39;s tarball, can be used to verify successful download | [optional] 
**Downloads** | **int32** | Number of times this release has been downloaded through the Forge | [optional] 
**Readme** | **string** | Contents of this release&#39;s README file, optionally converted to HTML format if possible (see &#x60;with_html&#x60; param) | [optional] 
**Changelog** | **string** | Contents of this release&#39;s CHANGELOG file, optionally converted to HTML format if possible (see &#x60;with_html&#x60; param) | [optional] 
**License** | **string** | Contents of this release&#39;s LICENSE file, optionally converted to HTML format if possible (see &#x60;with_html&#x60; param) | [optional] 
**Reference** | **string** | Contents of this release&#39;s REFERENCE file, optionally converted to HTML format if possible (see &#x60;with_html&#x60; param) | [optional] 
**Docs** | [**map[string]interface{}**](.md) | (OPTIONAL) Automatically generated documentation for custom Types included in this module release, see &#x60;reference&#x60; key instead (This field is not included in response by default, see &#x60;include_fields&#x60; param for more info)  | [optional] 
**Tasks** | [**[]ReleaseTask**](ReleaseTask.md) | Metadata for any Bolt tasks included in this module release | [optional] 
**Plans** | [**[]ReleasePlanAbbreviated**](ReleasePlanAbbreviated.md) | Links to metadata for Bolt plans included in this module release | [optional] 
**CreatedAt** | **string** | Date and time this Release resource was created | [optional] 
**UpdatedAt** | **string** | Date and time this Release resource was last updated | [optional] 
**DeletedAt** | Pointer to **string** | Date and time this Release resource was marked as deleted | [optional] 
**DeletedFor** | Pointer to **string** | Author-provided reason that this module release was marked as deleted | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


