# CountryPolicyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitiesAllowedStatuses** | **[]string** | List of allowed statuses for cities in the country | 
**CreatedAt** | **time.Time** | creation date | 
**UpdatedAt** | **time.Time** | last update date | 

## Methods

### NewCountryPolicyAttributes

`func NewCountryPolicyAttributes(citiesAllowedStatuses []string, createdAt time.Time, updatedAt time.Time, ) *CountryPolicyAttributes`

NewCountryPolicyAttributes instantiates a new CountryPolicyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryPolicyAttributesWithDefaults

`func NewCountryPolicyAttributesWithDefaults() *CountryPolicyAttributes`

NewCountryPolicyAttributesWithDefaults instantiates a new CountryPolicyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitiesAllowedStatuses

`func (o *CountryPolicyAttributes) GetCitiesAllowedStatuses() []string`

GetCitiesAllowedStatuses returns the CitiesAllowedStatuses field if non-nil, zero value otherwise.

### GetCitiesAllowedStatusesOk

`func (o *CountryPolicyAttributes) GetCitiesAllowedStatusesOk() (*[]string, bool)`

GetCitiesAllowedStatusesOk returns a tuple with the CitiesAllowedStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitiesAllowedStatuses

`func (o *CountryPolicyAttributes) SetCitiesAllowedStatuses(v []string)`

SetCitiesAllowedStatuses sets CitiesAllowedStatuses field to given value.


### GetCreatedAt

`func (o *CountryPolicyAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CountryPolicyAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CountryPolicyAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CountryPolicyAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CountryPolicyAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CountryPolicyAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


