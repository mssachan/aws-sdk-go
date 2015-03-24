package cognitosync

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteDatasetRequest generates a request for the DeleteDataset operation.
func (c *CognitoSync) DeleteDatasetRequest(input *DeleteDatasetInput) (req *aws.Request, output *DeleteDatasetOutput) {
	if opDeleteDataset == nil {
		opDeleteDataset = &aws.Operation{
			Name:       "DeleteDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteDataset, input, output)
	output = &DeleteDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DeleteDataset(input *DeleteDatasetInput) (output *DeleteDatasetOutput, err error) {
	req, out := c.DeleteDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteDataset *aws.Operation

// DescribeDatasetRequest generates a request for the DescribeDataset operation.
func (c *CognitoSync) DescribeDatasetRequest(input *DescribeDatasetInput) (req *aws.Request, output *DescribeDatasetOutput) {
	if opDescribeDataset == nil {
		opDescribeDataset = &aws.Operation{
			Name:       "DescribeDataset",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeDataset, input, output)
	output = &DescribeDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeDataset(input *DescribeDatasetInput) (output *DescribeDatasetOutput, err error) {
	req, out := c.DescribeDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDataset *aws.Operation

// DescribeIdentityPoolUsageRequest generates a request for the DescribeIdentityPoolUsage operation.
func (c *CognitoSync) DescribeIdentityPoolUsageRequest(input *DescribeIdentityPoolUsageInput) (req *aws.Request, output *DescribeIdentityPoolUsageOutput) {
	if opDescribeIdentityPoolUsage == nil {
		opDescribeIdentityPoolUsage = &aws.Operation{
			Name:       "DescribeIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityPoolUsage, input, output)
	output = &DescribeIdentityPoolUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeIdentityPoolUsage(input *DescribeIdentityPoolUsageInput) (output *DescribeIdentityPoolUsageOutput, err error) {
	req, out := c.DescribeIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityPoolUsage *aws.Operation

// DescribeIdentityUsageRequest generates a request for the DescribeIdentityUsage operation.
func (c *CognitoSync) DescribeIdentityUsageRequest(input *DescribeIdentityUsageInput) (req *aws.Request, output *DescribeIdentityUsageOutput) {
	if opDescribeIdentityUsage == nil {
		opDescribeIdentityUsage = &aws.Operation{
			Name:       "DescribeIdentityUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityUsage, input, output)
	output = &DescribeIdentityUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeIdentityUsage(input *DescribeIdentityUsageInput) (output *DescribeIdentityUsageOutput, err error) {
	req, out := c.DescribeIdentityUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityUsage *aws.Operation

// GetIdentityPoolConfigurationRequest generates a request for the GetIdentityPoolConfiguration operation.
func (c *CognitoSync) GetIdentityPoolConfigurationRequest(input *GetIdentityPoolConfigurationInput) (req *aws.Request, output *GetIdentityPoolConfigurationOutput) {
	if opGetIdentityPoolConfiguration == nil {
		opGetIdentityPoolConfiguration = &aws.Operation{
			Name:       "GetIdentityPoolConfiguration",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityPoolConfiguration, input, output)
	output = &GetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) GetIdentityPoolConfiguration(input *GetIdentityPoolConfigurationInput) (output *GetIdentityPoolConfigurationOutput, err error) {
	req, out := c.GetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityPoolConfiguration *aws.Operation

// ListDatasetsRequest generates a request for the ListDatasets operation.
func (c *CognitoSync) ListDatasetsRequest(input *ListDatasetsInput) (req *aws.Request, output *ListDatasetsOutput) {
	if opListDatasets == nil {
		opListDatasets = &aws.Operation{
			Name:       "ListDatasets",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets",
		}
	}

	req = aws.NewRequest(c.Service, opListDatasets, input, output)
	output = &ListDatasetsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListDatasets(input *ListDatasetsInput) (output *ListDatasetsOutput, err error) {
	req, out := c.ListDatasetsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDatasets *aws.Operation

// ListIdentityPoolUsageRequest generates a request for the ListIdentityPoolUsage operation.
func (c *CognitoSync) ListIdentityPoolUsageRequest(input *ListIdentityPoolUsageInput) (req *aws.Request, output *ListIdentityPoolUsageOutput) {
	if opListIdentityPoolUsage == nil {
		opListIdentityPoolUsage = &aws.Operation{
			Name:       "ListIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentityPoolUsage, input, output)
	output = &ListIdentityPoolUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListIdentityPoolUsage(input *ListIdentityPoolUsageInput) (output *ListIdentityPoolUsageOutput, err error) {
	req, out := c.ListIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentityPoolUsage *aws.Operation

// ListRecordsRequest generates a request for the ListRecords operation.
func (c *CognitoSync) ListRecordsRequest(input *ListRecordsInput) (req *aws.Request, output *ListRecordsOutput) {
	if opListRecords == nil {
		opListRecords = &aws.Operation{
			Name:       "ListRecords",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/records",
		}
	}

	req = aws.NewRequest(c.Service, opListRecords, input, output)
	output = &ListRecordsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListRecords(input *ListRecordsInput) (output *ListRecordsOutput, err error) {
	req, out := c.ListRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListRecords *aws.Operation

// RegisterDeviceRequest generates a request for the RegisterDevice operation.
func (c *CognitoSync) RegisterDeviceRequest(input *RegisterDeviceInput) (req *aws.Request, output *RegisterDeviceOutput) {
	if opRegisterDevice == nil {
		opRegisterDevice = &aws.Operation{
			Name:       "RegisterDevice",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device",
		}
	}

	req = aws.NewRequest(c.Service, opRegisterDevice, input, output)
	output = &RegisterDeviceOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) RegisterDevice(input *RegisterDeviceInput) (output *RegisterDeviceOutput, err error) {
	req, out := c.RegisterDeviceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterDevice *aws.Operation

// SetIdentityPoolConfigurationRequest generates a request for the SetIdentityPoolConfiguration operation.
func (c *CognitoSync) SetIdentityPoolConfigurationRequest(input *SetIdentityPoolConfigurationInput) (req *aws.Request, output *SetIdentityPoolConfigurationOutput) {
	if opSetIdentityPoolConfiguration == nil {
		opSetIdentityPoolConfiguration = &aws.Operation{
			Name:       "SetIdentityPoolConfiguration",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityPoolConfiguration, input, output)
	output = &SetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) SetIdentityPoolConfiguration(input *SetIdentityPoolConfigurationInput) (output *SetIdentityPoolConfigurationOutput, err error) {
	req, out := c.SetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityPoolConfiguration *aws.Operation

// SubscribeToDatasetRequest generates a request for the SubscribeToDataset operation.
func (c *CognitoSync) SubscribeToDatasetRequest(input *SubscribeToDatasetInput) (req *aws.Request, output *SubscribeToDatasetOutput) {
	if opSubscribeToDataset == nil {
		opSubscribeToDataset = &aws.Operation{
			Name:       "SubscribeToDataset",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opSubscribeToDataset, input, output)
	output = &SubscribeToDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) SubscribeToDataset(input *SubscribeToDatasetInput) (output *SubscribeToDatasetOutput, err error) {
	req, out := c.SubscribeToDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubscribeToDataset *aws.Operation

// UnsubscribeFromDatasetRequest generates a request for the UnsubscribeFromDataset operation.
func (c *CognitoSync) UnsubscribeFromDatasetRequest(input *UnsubscribeFromDatasetInput) (req *aws.Request, output *UnsubscribeFromDatasetOutput) {
	if opUnsubscribeFromDataset == nil {
		opUnsubscribeFromDataset = &aws.Operation{
			Name:       "UnsubscribeFromDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opUnsubscribeFromDataset, input, output)
	output = &UnsubscribeFromDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) UnsubscribeFromDataset(input *UnsubscribeFromDatasetInput) (output *UnsubscribeFromDatasetOutput, err error) {
	req, out := c.UnsubscribeFromDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnsubscribeFromDataset *aws.Operation

// UpdateRecordsRequest generates a request for the UpdateRecords operation.
func (c *CognitoSync) UpdateRecordsRequest(input *UpdateRecordsInput) (req *aws.Request, output *UpdateRecordsOutput) {
	if opUpdateRecords == nil {
		opUpdateRecords = &aws.Operation{
			Name:       "UpdateRecords",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateRecords, input, output)
	output = &UpdateRecordsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) UpdateRecords(input *UpdateRecordsInput) (output *UpdateRecordsOutput, err error) {
	req, out := c.UpdateRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateRecords *aws.Operation

type Dataset struct {
	CreationDate     *time.Time `type:"timestamp" timestampFormat:"unix"`
	DataStorage      *int64     `type:"long"`
	DatasetName      *string    `type:"string"`
	IdentityID       *string    `locationName:"IdentityId" type:"string"`
	LastModifiedBy   *string    `type:"string"`
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`
	NumRecords       *int64     `type:"long"`

	metadataDataset `json:"-", xml:"-"`
}

type metadataDataset struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDeleteDatasetInput `json:"-", xml:"-"`
}

type metadataDeleteDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteDatasetOutput struct {
	Dataset *Dataset `type:"structure"`

	metadataDeleteDatasetOutput `json:"-", xml:"-"`
}

type metadataDeleteDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeDatasetInput `json:"-", xml:"-"`
}

type metadataDescribeDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeDatasetOutput struct {
	Dataset *Dataset `type:"structure"`

	metadataDescribeDatasetOutput `json:"-", xml:"-"`
}

type metadataDescribeDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeIdentityPoolUsageInput struct {
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeIdentityPoolUsageOutput struct {
	IdentityPoolUsage *IdentityPoolUsage `type:"structure"`

	metadataDescribeIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeIdentityUsageInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataDescribeIdentityUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeIdentityUsageOutput struct {
	IdentityUsage *IdentityUsage `type:"structure"`

	metadataDescribeIdentityUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityPoolConfigurationInput struct {
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataGetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityPoolConfigurationOutput struct {
	IdentityPoolID *string   `locationName:"IdentityPoolId" type:"string"`
	PushSync       *PushSync `type:"structure"`

	metadataGetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type IdentityPoolUsage struct {
	DataStorage       *int64     `type:"long"`
	IdentityPoolID    *string    `locationName:"IdentityPoolId" type:"string"`
	LastModifiedDate  *time.Time `type:"timestamp" timestampFormat:"unix"`
	SyncSessionsCount *int64     `type:"long"`

	metadataIdentityPoolUsage `json:"-", xml:"-"`
}

type metadataIdentityPoolUsage struct {
	SDKShapeTraits bool `type:"structure"`
}

type IdentityUsage struct {
	DataStorage      *int64     `type:"long"`
	DatasetCount     *int64     `type:"integer"`
	IdentityID       *string    `locationName:"IdentityId" type:"string"`
	IdentityPoolID   *string    `locationName:"IdentityPoolId" type:"string"`
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	metadataIdentityUsage `json:"-", xml:"-"`
}

type metadataIdentityUsage struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDatasetsInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`
	MaxResults     *int64  `location:"querystring" locationName:"maxResults" type:"integer"`
	NextToken      *string `location:"querystring" locationName:"nextToken" type:"string"`

	metadataListDatasetsInput `json:"-", xml:"-"`
}

type metadataListDatasetsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDatasetsOutput struct {
	Count     *int64     `type:"integer"`
	Datasets  []*Dataset `type:"list"`
	NextToken *string    `type:"string"`

	metadataListDatasetsOutput `json:"-", xml:"-"`
}

type metadataListDatasetsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListIdentityPoolUsageInput struct {
	MaxResults *int64  `location:"querystring" locationName:"maxResults" type:"integer"`
	NextToken  *string `location:"querystring" locationName:"nextToken" type:"string"`

	metadataListIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListIdentityPoolUsageOutput struct {
	Count              *int64               `type:"integer"`
	IdentityPoolUsages []*IdentityPoolUsage `type:"list"`
	MaxResults         *int64               `type:"integer"`
	NextToken          *string              `type:"string"`

	metadataListIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListRecordsInput struct {
	DatasetName      *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	IdentityID       *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID   *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`
	LastSyncCount    *int64  `location:"querystring" locationName:"lastSyncCount" type:"long"`
	MaxResults       *int64  `location:"querystring" locationName:"maxResults" type:"integer"`
	NextToken        *string `location:"querystring" locationName:"nextToken" type:"string"`
	SyncSessionToken *string `location:"querystring" locationName:"syncSessionToken" type:"string"`

	metadataListRecordsInput `json:"-", xml:"-"`
}

type metadataListRecordsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListRecordsOutput struct {
	Count                                 *int64    `type:"integer"`
	DatasetDeletedAfterRequestedSyncCount *bool     `type:"boolean"`
	DatasetExists                         *bool     `type:"boolean"`
	DatasetSyncCount                      *int64    `type:"long"`
	LastModifiedBy                        *string   `type:"string"`
	MergedDatasetNames                    []*string `type:"list"`
	NextToken                             *string   `type:"string"`
	Records                               []*Record `type:"list"`
	SyncSessionToken                      *string   `type:"string"`

	metadataListRecordsOutput `json:"-", xml:"-"`
}

type metadataListRecordsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PushSync struct {
	ApplicationARNs []*string `locationName:"ApplicationArns" type:"list"`
	RoleARN         *string   `locationName:"RoleArn" type:"string"`

	metadataPushSync `json:"-", xml:"-"`
}

type metadataPushSync struct {
	SDKShapeTraits bool `type:"structure"`
}

type Record struct {
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`
	Key                    *string    `type:"string"`
	LastModifiedBy         *string    `type:"string"`
	LastModifiedDate       *time.Time `type:"timestamp" timestampFormat:"unix"`
	SyncCount              *int64     `type:"long"`
	Value                  *string    `type:"string"`

	metadataRecord `json:"-", xml:"-"`
}

type metadataRecord struct {
	SDKShapeTraits bool `type:"structure"`
}

type RecordPatch struct {
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`
	Key                    *string    `type:"string" required:"true"`
	Op                     *string    `type:"string" required:"true"`
	SyncCount              *int64     `type:"long" required:"true"`
	Value                  *string    `type:"string"`

	metadataRecordPatch `json:"-", xml:"-"`
}

type metadataRecordPatch struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterDeviceInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`
	Platform       *string `type:"string" required:"true"`
	Token          *string `type:"string" required:"true"`

	metadataRegisterDeviceInput `json:"-", xml:"-"`
}

type metadataRegisterDeviceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterDeviceOutput struct {
	DeviceID *string `locationName:"DeviceId" type:"string"`

	metadataRegisterDeviceOutput `json:"-", xml:"-"`
}

type metadataRegisterDeviceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityPoolConfigurationInput struct {
	IdentityPoolID *string   `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`
	PushSync       *PushSync `type:"structure"`

	metadataSetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityPoolConfigurationOutput struct {
	IdentityPoolID *string   `locationName:"IdentityPoolId" type:"string"`
	PushSync       *PushSync `type:"structure"`

	metadataSetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubscribeToDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	DeviceID       *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataSubscribeToDatasetInput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubscribeToDatasetOutput struct {
	metadataSubscribeToDatasetOutput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UnsubscribeFromDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	DeviceID       *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`

	metadataUnsubscribeFromDatasetInput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UnsubscribeFromDatasetOutput struct {
	metadataUnsubscribeFromDatasetOutput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateRecordsInput struct {
	ClientContext    *string        `location:"header" locationName:"x-amz-Client-Context" type:"string"`
	DatasetName      *string        `location:"uri" locationName:"DatasetName" type:"string" required:"true"`
	DeviceID         *string        `locationName:"DeviceId" type:"string"`
	IdentityID       *string        `location:"uri" locationName:"IdentityId" type:"string" required:"true"`
	IdentityPoolID   *string        `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"`
	RecordPatches    []*RecordPatch `type:"list"`
	SyncSessionToken *string        `type:"string" required:"true"`

	metadataUpdateRecordsInput `json:"-", xml:"-"`
}

type metadataUpdateRecordsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateRecordsOutput struct {
	Records []*Record `type:"list"`

	metadataUpdateRecordsOutput `json:"-", xml:"-"`
}

type metadataUpdateRecordsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}