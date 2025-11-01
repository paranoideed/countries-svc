# CountryPoliciesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CountryPolicyData**](CountryPolicyData.md) |  | 
**Links** | [**PaginationData**](PaginationData.md) |  | 

## Methods

### NewCountryPoliciesCollection

`func NewCountryPoliciesCollection(data []CountryPolicyData, links PaginationData, ) *CountryPoliciesCollection`

NewCountryPoliciesCollection instantiates a new CountryPoliciesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryPoliciesCollectionWithDefaults

`func NewCountryPoliciesCollectionWithDefaults() *CountryPoliciesCollection`

NewCountryPoliciesCollectionWithDefaults instantiates a new CountryPoliciesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CountryPoliciesCollection) GetData() []CountryPolicyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CountryPoliciesCollection) GetDataOk() (*[]CountryPolicyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CountryPoliciesCollection) SetData(v []CountryPolicyData)`

SetData sets Data field to given value.


### GetLinks

`func (o *CountryPoliciesCollection) GetLinks() PaginationData`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CountryPoliciesCollection) GetLinksOk() (*PaginationData, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CountryPoliciesCollection) SetLinks(v PaginationData)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


