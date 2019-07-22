// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package gallery

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("69d21c00-f135-441b-b5ce-3626378e0819")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API]
// ctx
// extensionId (required)
// accountName (required)
func (client Client) ShareExtensionById(ctx context.Context, extensionId *uuid.UUID, accountName *string) error {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()
    if accountName == nil || *accountName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "accountName"} 
    }
    routeValues["accountName"] = *accountName

    locationId, _ := uuid.Parse("1f19631b-a0b4-4a03-89c2-d79785d24360")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// extensionId (required)
// accountName (required)
func (client Client) UnshareExtensionById(ctx context.Context, extensionId *uuid.UUID, accountName *string) error {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()
    if accountName == nil || *accountName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "accountName"} 
    }
    routeValues["accountName"] = *accountName

    locationId, _ := uuid.Parse("1f19631b-a0b4-4a03-89c2-d79785d24360")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// accountName (required)
func (client Client) ShareExtension(ctx context.Context, publisherName *string, extensionName *string, accountName *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if accountName == nil || *accountName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "accountName"} 
    }
    routeValues["accountName"] = *accountName

    locationId, _ := uuid.Parse("a1e66d8f-f5de-4d16-8309-91a4e015ee46")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// accountName (required)
func (client Client) UnshareExtension(ctx context.Context, publisherName *string, extensionName *string, accountName *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if accountName == nil || *accountName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "accountName"} 
    }
    routeValues["accountName"] = *accountName

    locationId, _ := uuid.Parse("a1e66d8f-f5de-4d16-8309-91a4e015ee46")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// itemId (required)
// installationTarget (required)
// testCommerce (optional)
// isFreeOrTrialInstall (optional)
func (client Client) GetAcquisitionOptions(ctx context.Context, itemId *string, installationTarget *string, testCommerce *bool, isFreeOrTrialInstall *bool) (*AcquisitionOptions, error) {
    routeValues := make(map[string]string)
    if itemId == nil || *itemId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "itemId"} 
    }
    routeValues["itemId"] = *itemId

    queryParams := url.Values{}
    if installationTarget == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "installationTarget"}
    }
    queryParams.Add("installationTarget", *installationTarget)
    if testCommerce != nil {
        queryParams.Add("testCommerce", strconv.FormatBool(*testCommerce))
    }
    if isFreeOrTrialInstall != nil {
        queryParams.Add("isFreeOrTrialInstall", strconv.FormatBool(*isFreeOrTrialInstall))
    }
    locationId, _ := uuid.Parse("9d0a0105-075e-4760-aa15-8bcf54d1bd7d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AcquisitionOptions
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// acquisitionRequest (required)
func (client Client) RequestAcquisition(ctx context.Context, acquisitionRequest *ExtensionAcquisitionRequest) (*ExtensionAcquisitionRequest, error) {
    if acquisitionRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "acquisitionRequest"}
    }
    body, marshalErr := json.Marshal(*acquisitionRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3adb1f2d-e328-446e-be73-9f6d98071c45")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionAcquisitionRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (required)
// assetType (required)
// accountToken (optional)
// acceptDefault (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetAssetByName(ctx context.Context, publisherName *string, extensionName *string, version *string, assetType *string, accountToken *string, acceptDefault *bool, accountTokenHeader *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    if acceptDefault != nil {
        queryParams.Add("acceptDefault", strconv.FormatBool(*acceptDefault))
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("7529171f-a002-4180-93ba-685f358a0482")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// extensionId (required)
// version (required)
// assetType (required)
// accountToken (optional)
// acceptDefault (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetAsset(ctx context.Context, extensionId *uuid.UUID, version *string, assetType *string, accountToken *string, acceptDefault *bool, accountTokenHeader *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    if acceptDefault != nil {
        queryParams.Add("acceptDefault", strconv.FormatBool(*acceptDefault))
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("5d545f3d-ef47-488b-8be3-f5ee1517856c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (required)
// assetType (required)
// accountToken (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetAssetAuthenticated(ctx context.Context, publisherName *string, extensionName *string, version *string, assetType *string, accountToken *string, accountTokenHeader *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("506aff36-2622-4f70-8063-77cce6366d20")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// azurePublisherId (required)
func (client Client) AssociateAzurePublisher(ctx context.Context, publisherName *string, azurePublisherId *string) (*AzurePublisher, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if azurePublisherId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "azurePublisherId"}
    }
    queryParams.Add("azurePublisherId", *azurePublisherId)
    locationId, _ := uuid.Parse("efd202a6-9d87-4ebc-9229-d2b8ae2fdb6d")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AzurePublisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
func (client Client) QueryAssociatedAzurePublisher(ctx context.Context, publisherName *string) (*AzurePublisher, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    locationId, _ := uuid.Parse("efd202a6-9d87-4ebc-9229-d2b8ae2fdb6d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AzurePublisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// languages (optional)
func (client Client) GetCategories(ctx context.Context, languages *string) (*[]string, error) {
    queryParams := url.Values{}
    if languages != nil {
        queryParams.Add("languages", *languages)
    }
    locationId, _ := uuid.Parse("e0a5a71e-3ac3-43a0-ae7d-0bb5c3046a2a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// categoryName (required)
// languages (optional)
// product (optional)
func (client Client) GetCategoryDetails(ctx context.Context, categoryName *string, languages *string, product *string) (*CategoriesResult, error) {
    routeValues := make(map[string]string)
    if categoryName == nil || *categoryName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "categoryName"} 
    }
    routeValues["categoryName"] = *categoryName

    queryParams := url.Values{}
    if languages != nil {
        queryParams.Add("languages", *languages)
    }
    if product != nil {
        queryParams.Add("product", *product)
    }
    locationId, _ := uuid.Parse("75d3c04d-84d2-4973-acd2-22627587dabc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CategoriesResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// product (required)
// categoryId (required)
// lcid (optional)
// source (optional)
// productVersion (optional)
// skus (optional)
// subSkus (optional)
func (client Client) GetCategoryTree(ctx context.Context, product *string, categoryId *string, lcid *int, source *string, productVersion *string, skus *string, subSkus *string) (*ProductCategory, error) {
    routeValues := make(map[string]string)
    if product == nil || *product == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "product"} 
    }
    routeValues["product"] = *product
    if categoryId == nil || *categoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "categoryId"} 
    }
    routeValues["categoryId"] = *categoryId

    queryParams := url.Values{}
    if lcid != nil {
        queryParams.Add("lcid", strconv.Itoa(*lcid))
    }
    if source != nil {
        queryParams.Add("source", *source)
    }
    if productVersion != nil {
        queryParams.Add("productVersion", *productVersion)
    }
    if skus != nil {
        queryParams.Add("skus", *skus)
    }
    if subSkus != nil {
        queryParams.Add("subSkus", *subSkus)
    }
    locationId, _ := uuid.Parse("1102bb42-82b0-4955-8d8a-435d6b4cedd3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProductCategory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// product (required)
// lcid (optional)
// source (optional)
// productVersion (optional)
// skus (optional)
// subSkus (optional)
func (client Client) GetRootCategories(ctx context.Context, product *string, lcid *int, source *string, productVersion *string, skus *string, subSkus *string) (*ProductCategoriesResult, error) {
    routeValues := make(map[string]string)
    if product == nil || *product == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "product"} 
    }
    routeValues["product"] = *product

    queryParams := url.Values{}
    if lcid != nil {
        queryParams.Add("lcid", strconv.Itoa(*lcid))
    }
    if source != nil {
        queryParams.Add("source", *source)
    }
    if productVersion != nil {
        queryParams.Add("productVersion", *productVersion)
    }
    if skus != nil {
        queryParams.Add("skus", *skus)
    }
    if subSkus != nil {
        queryParams.Add("subSkus", *subSkus)
    }
    locationId, _ := uuid.Parse("31fba831-35b2-46f6-a641-d05de5a877d8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProductCategoriesResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (optional)
func (client Client) GetCertificate(ctx context.Context, publisherName *string, extensionName *string, version *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version != nil && *version != "" {
        routeValues["version"] = *version
    }

    locationId, _ := uuid.Parse("e905ad6a-3f1f-4d08-9f6d-7d357ff8b7d0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
func (client Client) GetContentVerificationLog(ctx context.Context, publisherName *string, extensionName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    locationId, _ := uuid.Parse("c0f1c7c4-3557-4ffb-b774-1e48c4865e99")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
func (client Client) CreateDraftForEditExtension(ctx context.Context, publisherName *string, extensionName *string) (*ExtensionDraft, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    locationId, _ := uuid.Parse("02b33873-4e61-496e-83a2-59d1df46b7d8")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// draftPatch (required)
// publisherName (required)
// extensionName (required)
// draftId (required)
func (client Client) PerformEditExtensionDraftOperation(ctx context.Context, draftPatch *ExtensionDraftPatch, publisherName *string, extensionName *string, draftId *uuid.UUID) (*ExtensionDraft, error) {
    if draftPatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftPatch"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()

    body, marshalErr := json.Marshal(*draftPatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02b33873-4e61-496e-83a2-59d1df46b7d8")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
// extensionName (required)
// draftId (required)
// fileName (optional): Header to pass the filename of the uploaded data
func (client Client) UpdatePayloadInDraftForEditExtension(ctx context.Context, uploadStream *interface{}, publisherName *string, extensionName *string, draftId *uuid.UUID, fileName *string) (*ExtensionDraft, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()

    additionalHeaders := make(map[string]string)
    if fileName != nil {
        additionalHeaders["X-Market-UploadFileName"] = *fileName
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02b33873-4e61-496e-83a2-59d1df46b7d8")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
// extensionName (required)
// draftId (required)
// assetType (required)
func (client Client) AddAssetForEditExtensionDraft(ctx context.Context, uploadStream *interface{}, publisherName *string, extensionName *string, draftId *uuid.UUID, assetType *string) (*ExtensionDraftAsset, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f1db9c47-6619-4998-a7e5-d7f9f41a4617")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraftAsset
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
// product (required): Header to pass the product type of the payload file
// fileName (optional): Header to pass the filename of the uploaded data
func (client Client) CreateDraftForNewExtension(ctx context.Context, uploadStream *interface{}, publisherName *string, product *string, fileName *string) (*ExtensionDraft, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    additionalHeaders := make(map[string]string)
    if product != nil {
        additionalHeaders["X-Market-UploadFileProduct"] = *product
    }
    if fileName != nil {
        additionalHeaders["X-Market-UploadFileName"] = *fileName
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b3ab127d-ebb9-4d22-b611-4e09593c8d79")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// draftPatch (required)
// publisherName (required)
// draftId (required)
func (client Client) PerformNewExtensionDraftOperation(ctx context.Context, draftPatch *ExtensionDraftPatch, publisherName *string, draftId *uuid.UUID) (*ExtensionDraft, error) {
    if draftPatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftPatch"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()

    body, marshalErr := json.Marshal(*draftPatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b3ab127d-ebb9-4d22-b611-4e09593c8d79")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
// draftId (required)
// fileName (optional): Header to pass the filename of the uploaded data
func (client Client) UpdatePayloadInDraftForNewExtension(ctx context.Context, uploadStream *interface{}, publisherName *string, draftId *uuid.UUID, fileName *string) (*ExtensionDraft, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()

    additionalHeaders := make(map[string]string)
    if fileName != nil {
        additionalHeaders["X-Market-UploadFileName"] = *fileName
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b3ab127d-ebb9-4d22-b611-4e09593c8d79")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraft
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
// draftId (required)
// assetType (required)
func (client Client) AddAssetForNewExtensionDraft(ctx context.Context, uploadStream *interface{}, publisherName *string, draftId *uuid.UUID, assetType *string) (*ExtensionDraftAsset, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("88c0b1c8-b4f1-498a-9b2a-8446ef9f32e7")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDraftAsset
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// draftId (required)
// assetType (required)
// extensionName (required)
func (client Client) GetAssetFromEditExtensionDraft(ctx context.Context, publisherName *string, draftId *uuid.UUID, assetType *string, extensionName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    queryParams := url.Values{}
    if extensionName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "extensionName"}
    }
    queryParams.Add("extensionName", *extensionName)
    locationId, _ := uuid.Parse("88c0b1c8-b4f1-498a-9b2a-8446ef9f32e7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// draftId (required)
// assetType (required)
func (client Client) GetAssetFromNewExtensionDraft(ctx context.Context, publisherName *string, draftId *uuid.UUID, assetType *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if draftId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "draftId"} 
    }
    routeValues["draftId"] = (*draftId).String()
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType

    locationId, _ := uuid.Parse("88c0b1c8-b4f1-498a-9b2a-8446ef9f32e7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get install/uninstall events of an extension. If both count and afterDate parameters are specified, count takes precedence.
// ctx
// publisherName (required): Name of the publisher
// extensionName (required): Name of the extension
// count (optional): Count of events to fetch, applies to each event type.
// afterDate (optional): Fetch events that occurred on or after this date
// include (optional): Filter options. Supported values: install, uninstall, review, acquisition, sales. Default is to fetch all types of events
// includeProperty (optional): Event properties to include. Currently only 'lastContactDetails' is supported for uninstall events
func (client Client) GetExtensionEvents(ctx context.Context, publisherName *string, extensionName *string, count *int, afterDate *time.Time, include *string, includeProperty *string) (*ExtensionEvents, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if count != nil {
        queryParams.Add("count", strconv.Itoa(*count))
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    if include != nil {
        queryParams.Add("include", *include)
    }
    if includeProperty != nil {
        queryParams.Add("includeProperty", *includeProperty)
    }
    locationId, _ := uuid.Parse("3d13c499-2168-4d06-bef4-14aba185dcd5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionEvents
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] API endpoint to publish extension install/uninstall events. This is meant to be invoked by EMS only for sending us data related to install/uninstall of an extension.
// ctx
// extensionEvents (required)
func (client Client) PublishExtensionEvents(ctx context.Context, extensionEvents *[]ExtensionEvents) error {
    if extensionEvents == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "extensionEvents"}
    }
    body, marshalErr := json.Marshal(*extensionEvents)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("0bf2bd3a-70e0-4d5d-8bf7-bd4a9c2ab6e7")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// extensionQuery (required)
// accountToken (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) QueryExtensions(ctx context.Context, extensionQuery *ExtensionQuery, accountToken *string, accountTokenHeader *string) (*ExtensionQueryResult, error) {
    if extensionQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "extensionQuery"}
    }
    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    body, marshalErr := json.Marshal(*extensionQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("eb9d5ee1-6d43-456b-b80e-8a96fbc014b6")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
func (client Client) CreateExtension(ctx context.Context, uploadStream *interface{}) (*PublishedExtension, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a41192c8-9525-4b58-bc86-179fa549d80d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// extensionId (required)
// version (optional)
func (client Client) DeleteExtensionById(ctx context.Context, extensionId *uuid.UUID, version *string) error {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()

    queryParams := url.Values{}
    if version != nil {
        queryParams.Add("version", *version)
    }
    locationId, _ := uuid.Parse("a41192c8-9525-4b58-bc86-179fa549d80d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// extensionId (required)
// version (optional)
// flags (optional)
func (client Client) GetExtensionById(ctx context.Context, extensionId *uuid.UUID, version *string, flags *ExtensionQueryFlags) (*PublishedExtension, error) {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()

    queryParams := url.Values{}
    if version != nil {
        queryParams.Add("version", *version)
    }
    if flags != nil {
        queryParams.Add("flags", string(*flags))
    }
    locationId, _ := uuid.Parse("a41192c8-9525-4b58-bc86-179fa549d80d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// extensionId (required)
func (client Client) UpdateExtensionById(ctx context.Context, extensionId *uuid.UUID) (*PublishedExtension, error) {
    routeValues := make(map[string]string)
    if extensionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "extensionId"} 
    }
    routeValues["extensionId"] = (*extensionId).String()

    locationId, _ := uuid.Parse("a41192c8-9525-4b58-bc86-179fa549d80d")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// publisherName (required)
func (client Client) CreateExtensionWithPublisher(ctx context.Context, uploadStream *interface{}, publisherName *string) (*PublishedExtension, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e11ea35a-16fe-4b80-ab11-c4cab88a0966")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (optional)
func (client Client) DeleteExtension(ctx context.Context, publisherName *string, extensionName *string, version *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if version != nil {
        queryParams.Add("version", *version)
    }
    locationId, _ := uuid.Parse("e11ea35a-16fe-4b80-ab11-c4cab88a0966")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (optional)
// flags (optional)
// accountToken (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetExtension(ctx context.Context, publisherName *string, extensionName *string, version *string, flags *ExtensionQueryFlags, accountToken *string, accountTokenHeader *string) (*PublishedExtension, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if version != nil {
        queryParams.Add("version", *version)
    }
    if flags != nil {
        queryParams.Add("flags", string(*flags))
    }
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("e11ea35a-16fe-4b80-ab11-c4cab88a0966")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] REST endpoint to update an extension.
// ctx
// uploadStream (required): Stream to upload
// publisherName (required): Name of the publisher
// extensionName (required): Name of the extension
// bypassScopeCheck (optional): This parameter decides if the scope change check needs to be invoked or not
func (client Client) UpdateExtension(ctx context.Context, uploadStream *interface{}, publisherName *string, extensionName *string, bypassScopeCheck *bool) (*PublishedExtension, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if bypassScopeCheck != nil {
        queryParams.Add("bypassScopeCheck", strconv.FormatBool(*bypassScopeCheck))
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e11ea35a-16fe-4b80-ab11-c4cab88a0966")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// flags (required)
func (client Client) UpdateExtensionProperties(ctx context.Context, publisherName *string, extensionName *string, flags *PublishedExtensionFlags) (*PublishedExtension, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if flags == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", string(*flags))
    locationId, _ := uuid.Parse("e11ea35a-16fe-4b80-ab11-c4cab88a0966")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishedExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// hostType (required)
// hostName (required)
func (client Client) ShareExtensionWithHost(ctx context.Context, publisherName *string, extensionName *string, hostType *string, hostName *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if hostType == nil || *hostType == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hostType"} 
    }
    routeValues["hostType"] = *hostType
    if hostName == nil || *hostName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hostName"} 
    }
    routeValues["hostName"] = *hostName

    locationId, _ := uuid.Parse("328a3af8-d124-46e9-9483-01690cd415b9")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// hostType (required)
// hostName (required)
func (client Client) UnshareExtensionWithHost(ctx context.Context, publisherName *string, extensionName *string, hostType *string, hostName *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if hostType == nil || *hostType == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hostType"} 
    }
    routeValues["hostType"] = *hostType
    if hostName == nil || *hostName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hostName"} 
    }
    routeValues["hostName"] = *hostName

    locationId, _ := uuid.Parse("328a3af8-d124-46e9-9483-01690cd415b9")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// azureRestApiRequestModel (required)
func (client Client) ExtensionValidator(ctx context.Context, azureRestApiRequestModel *AzureRestApiRequestModel) error {
    if azureRestApiRequestModel == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "azureRestApiRequestModel"}
    }
    body, marshalErr := json.Marshal(*azureRestApiRequestModel)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("05e8a5e1-8c59-4c2c-8856-0ff087d1a844")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Send Notification
// ctx
// notificationData (required): Denoting the data needed to send notification
func (client Client) SendNotifications(ctx context.Context, notificationData *NotificationsData) error {
    if notificationData == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "notificationData"}
    }
    body, marshalErr := json.Marshal(*notificationData)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("eab39817-413c-4602-a49f-07ad00844980")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] This endpoint gets hit when you download a VSTS extension from the Web UI
// ctx
// publisherName (required)
// extensionName (required)
// version (required)
// accountToken (optional)
// acceptDefault (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetPackage(ctx context.Context, publisherName *string, extensionName *string, version *string, accountToken *string, acceptDefault *bool, accountTokenHeader *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version

    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    if acceptDefault != nil {
        queryParams.Add("acceptDefault", strconv.FormatBool(*acceptDefault))
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("7cb576f8-1cae-4c4b-b7b1-e4af5759e965")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (required)
// assetType (required)
// assetToken (optional)
// accountToken (optional)
// acceptDefault (optional)
// accountTokenHeader (optional): Header to pass the account token
func (client Client) GetAssetWithToken(ctx context.Context, publisherName *string, extensionName *string, version *string, assetType *string, assetToken *string, accountToken *string, acceptDefault *bool, accountTokenHeader *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version
    if assetType == nil || *assetType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "assetType"} 
    }
    routeValues["assetType"] = *assetType
    if assetToken != nil && *assetToken != "" {
        routeValues["assetToken"] = *assetToken
    }

    queryParams := url.Values{}
    if accountToken != nil {
        queryParams.Add("accountToken", *accountToken)
    }
    if acceptDefault != nil {
        queryParams.Add("acceptDefault", strconv.FormatBool(*acceptDefault))
    }
    additionalHeaders := make(map[string]string)
    if accountTokenHeader != nil {
        additionalHeaders["X-Market-AccountToken"] = *accountTokenHeader
    }
    locationId, _ := uuid.Parse("364415a1-0077-4a41-a7a0-06edd4497492")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete publisher asset like logo
// ctx
// publisherName (required): Internal name of the publisher
// assetType (optional): Type of asset. Default value is 'logo'.
func (client Client) DeletePublisherAsset(ctx context.Context, publisherName *string, assetType *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if assetType != nil {
        queryParams.Add("assetType", *assetType)
    }
    locationId, _ := uuid.Parse("21143299-34f9-4c62-8ca8-53da691192f9")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get publisher asset like logo as a stream
// ctx
// publisherName (required): Internal name of the publisher
// assetType (optional): Type of asset. Default value is 'logo'.
func (client Client) GetPublisherAsset(ctx context.Context, publisherName *string, assetType *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if assetType != nil {
        queryParams.Add("assetType", *assetType)
    }
    locationId, _ := uuid.Parse("21143299-34f9-4c62-8ca8-53da691192f9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update publisher asset like logo. It accepts asset file as an octet stream and file name is passed in header values.
// ctx
// uploadStream (required): Stream to upload
// publisherName (required): Internal name of the publisher
// assetType (optional): Type of asset. Default value is 'logo'.
// fileName (optional): Header to pass the filename of the uploaded data
func (client Client) UpdatePublisherAsset(ctx context.Context, uploadStream *interface{}, publisherName *string, assetType *string, fileName *string) (*map[string]string, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if assetType != nil {
        queryParams.Add("assetType", *assetType)
    }
    additionalHeaders := make(map[string]string)
    if fileName != nil {
        additionalHeaders["X-Market-UploadFileName"] = *fileName
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("21143299-34f9-4c62-8ca8-53da691192f9")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/octet-stream", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherQuery (required)
func (client Client) QueryPublishers(ctx context.Context, publisherQuery *PublisherQuery) (*PublisherQueryResult, error) {
    if publisherQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "publisherQuery"}
    }
    body, marshalErr := json.Marshal(*publisherQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2ad6ee0a-b53f-4034-9d1d-d009fda1212e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublisherQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisher (required)
func (client Client) CreatePublisher(ctx context.Context, publisher *Publisher) (*Publisher, error) {
    if publisher == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "publisher"}
    }
    body, marshalErr := json.Marshal(*publisher)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4ddec66a-e4f6-4f5d-999e-9e77710d7ff4")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Publisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
func (client Client) DeletePublisher(ctx context.Context, publisherName *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    locationId, _ := uuid.Parse("4ddec66a-e4f6-4f5d-999e-9e77710d7ff4")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// flags (optional)
func (client Client) GetPublisher(ctx context.Context, publisherName *string, flags *int) (*Publisher, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if flags != nil {
        queryParams.Add("flags", strconv.Itoa(*flags))
    }
    locationId, _ := uuid.Parse("4ddec66a-e4f6-4f5d-999e-9e77710d7ff4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Publisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisher (required)
// publisherName (required)
func (client Client) UpdatePublisher(ctx context.Context, publisher *Publisher, publisherName *string) (*Publisher, error) {
    if publisher == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "publisher"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    body, marshalErr := json.Marshal(*publisher)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4ddec66a-e4f6-4f5d-999e-9e77710d7ff4")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Publisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Endpoint to add/modify publisher membership. Currently Supports only addition/modification of 1 user at a time Works only for adding members of same tenant.
// ctx
// roleAssignments (required): List of user identifiers(email address) and role to be added. Currently only one entry is supported.
// publisherName (required): The name/id of publisher to which users have to be added
// limitToCallerIdentityDomain (optional): Should cross tenant addtions be allowed or not.
func (client Client) UpdatePublisherMembers(ctx context.Context, roleAssignments *[]PublisherUserRoleAssignmentRef, publisherName *string, limitToCallerIdentityDomain *bool) (*[]PublisherRoleAssignment, error) {
    if roleAssignments == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "roleAssignments"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName

    queryParams := url.Values{}
    if limitToCallerIdentityDomain != nil {
        queryParams.Add("limitToCallerIdentityDomain", strconv.FormatBool(*limitToCallerIdentityDomain))
    }
    body, marshalErr := json.Marshal(*roleAssignments)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4ddec66a-e4f6-4f5d-999e-9e77710d7ff4")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PublisherRoleAssignment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of questions with their responses associated with an extension.
// ctx
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// count (optional): Number of questions to retrieve (defaults to 10).
// page (optional): Page number from which set of questions are to be retrieved.
// afterDate (optional): If provided, results questions are returned which were posted after this date
func (client Client) GetQuestions(ctx context.Context, publisherName *string, extensionName *string, count *int, page *int, afterDate *time.Time) (*QuestionsResult, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if count != nil {
        queryParams.Add("count", strconv.Itoa(*count))
    }
    if page != nil {
        queryParams.Add("page", strconv.Itoa(*page))
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    locationId, _ := uuid.Parse("c010d03d-812c-4ade-ae07-c1862475eda5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue QuestionsResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Flags a concern with an existing question for an extension.
// ctx
// concern (required): User reported concern with a question for the extension.
// pubName (required): Name of the publisher who published the extension.
// extName (required): Name of the extension.
// questionId (required): Identifier of the question to be updated for the extension.
func (client Client) ReportQuestion(ctx context.Context, concern *Concern, pubName *string, extName *string, questionId *uint64) (*Concern, error) {
    if concern == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "concern"}
    }
    routeValues := make(map[string]string)
    if pubName == nil || *pubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pubName"} 
    }
    routeValues["pubName"] = *pubName
    if extName == nil || *extName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extName"} 
    }
    routeValues["extName"] = *extName
    if questionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)

    body, marshalErr := json.Marshal(*concern)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("784910cd-254a-494d-898b-0728549b2f10")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Concern
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a new question for an extension.
// ctx
// question (required): Question to be created for the extension.
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
func (client Client) CreateQuestion(ctx context.Context, question *Question, publisherName *string, extensionName *string) (*Question, error) {
    if question == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "question"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    body, marshalErr := json.Marshal(*question)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6d1d9741-eca8-4701-a3a5-235afc82dfa4")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Question
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes an existing question and all its associated responses for an extension. (soft delete)
// ctx
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// questionId (required): Identifier of the question to be deleted for the extension.
func (client Client) DeleteQuestion(ctx context.Context, publisherName *string, extensionName *string, questionId *uint64) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if questionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)

    locationId, _ := uuid.Parse("6d1d9741-eca8-4701-a3a5-235afc82dfa4")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates an existing question for an extension.
// ctx
// question (required): Updated question to be set for the extension.
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// questionId (required): Identifier of the question to be updated for the extension.
func (client Client) UpdateQuestion(ctx context.Context, question *Question, publisherName *string, extensionName *string, questionId *uint64) (*Question, error) {
    if question == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "question"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if questionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)

    body, marshalErr := json.Marshal(*question)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6d1d9741-eca8-4701-a3a5-235afc82dfa4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Question
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a new response for a given question for an extension.
// ctx
// response (required): Response to be created for the extension.
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// questionId (required): Identifier of the question for which response is to be created for the extension.
func (client Client) CreateResponse(ctx context.Context, response *Response, publisherName *string, extensionName *string, questionId *uint64) (*Response, error) {
    if response == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "response"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if questionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)

    body, marshalErr := json.Marshal(*response)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7f8ae5e0-46b0-438f-b2e8-13e8513517bd")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Response
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes a response for an extension. (soft delete)
// ctx
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// questionId (required): Identifies the question whose response is to be deleted.
// responseId (required): Identifies the response to be deleted.
func (client Client) DeleteResponse(ctx context.Context, publisherName *string, extensionName *string, questionId *uint64, responseId *uint64) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if questionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)
    if responseId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "responseId"} 
    }
    routeValues["responseId"] = strconv.FormatUint(*responseId, 10)

    locationId, _ := uuid.Parse("7f8ae5e0-46b0-438f-b2e8-13e8513517bd")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates an existing response for a given question for an extension.
// ctx
// response (required): Updated response to be set for the extension.
// publisherName (required): Name of the publisher who published the extension.
// extensionName (required): Name of the extension.
// questionId (required): Identifier of the question for which response is to be updated for the extension.
// responseId (required): Identifier of the response which has to be updated.
func (client Client) UpdateResponse(ctx context.Context, response *Response, publisherName *string, extensionName *string, questionId *uint64, responseId *uint64) (*Response, error) {
    if response == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "response"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if questionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "questionId"} 
    }
    routeValues["questionId"] = strconv.FormatUint(*questionId, 10)
    if responseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "responseId"} 
    }
    routeValues["responseId"] = strconv.FormatUint(*responseId, 10)

    body, marshalErr := json.Marshal(*response)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7f8ae5e0-46b0-438f-b2e8-13e8513517bd")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Response
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns extension reports
// ctx
// publisherName (required): Name of the publisher who published the extension
// extensionName (required): Name of the extension
// days (optional): Last n days report. If afterDate and days are specified, days will take priority
// count (optional): Number of events to be returned
// afterDate (optional): Use if you want to fetch events newer than the specified date
func (client Client) GetExtensionReports(ctx context.Context, publisherName *string, extensionName *string, days *int, count *int, afterDate *time.Time) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if days != nil {
        queryParams.Add("days", strconv.Itoa(*days))
    }
    if count != nil {
        queryParams.Add("count", strconv.Itoa(*count))
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    locationId, _ := uuid.Parse("79e0c74f-157f-437e-845f-74fbb4121d4c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of reviews associated with an extension
// ctx
// publisherName (required): Name of the publisher who published the extension
// extensionName (required): Name of the extension
// count (optional): Number of reviews to retrieve (defaults to 5)
// filterOptions (optional): FilterOptions to filter out empty reviews etcetera, defaults to none
// beforeDate (optional): Use if you want to fetch reviews older than the specified date, defaults to null
// afterDate (optional): Use if you want to fetch reviews newer than the specified date, defaults to null
func (client Client) GetReviews(ctx context.Context, publisherName *string, extensionName *string, count *int, filterOptions *ReviewFilterOptions, beforeDate *time.Time, afterDate *time.Time) (*ReviewsResult, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if count != nil {
        queryParams.Add("count", strconv.Itoa(*count))
    }
    if filterOptions != nil {
        queryParams.Add("filterOptions", string(*filterOptions))
    }
    if beforeDate != nil {
        queryParams.Add("beforeDate", (*beforeDate).String())
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    locationId, _ := uuid.Parse("5b3f819f-f247-42ad-8c00-dd9ab9ab246d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReviewsResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a summary of the reviews
// ctx
// pubName (required): Name of the publisher who published the extension
// extName (required): Name of the extension
// beforeDate (optional): Use if you want to fetch summary of reviews older than the specified date, defaults to null
// afterDate (optional): Use if you want to fetch summary of reviews newer than the specified date, defaults to null
func (client Client) GetReviewsSummary(ctx context.Context, pubName *string, extName *string, beforeDate *time.Time, afterDate *time.Time) (*ReviewSummary, error) {
    routeValues := make(map[string]string)
    if pubName == nil || *pubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pubName"} 
    }
    routeValues["pubName"] = *pubName
    if extName == nil || *extName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extName"} 
    }
    routeValues["extName"] = *extName

    queryParams := url.Values{}
    if beforeDate != nil {
        queryParams.Add("beforeDate", (*beforeDate).String())
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    locationId, _ := uuid.Parse("b7b44e21-209e-48f0-ae78-04727fc37d77")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReviewSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a new review for an extension
// ctx
// review (required): Review to be created for the extension
// pubName (required): Name of the publisher who published the extension
// extName (required): Name of the extension
func (client Client) CreateReview(ctx context.Context, review *Review, pubName *string, extName *string) (*Review, error) {
    if review == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "review"}
    }
    routeValues := make(map[string]string)
    if pubName == nil || *pubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pubName"} 
    }
    routeValues["pubName"] = *pubName
    if extName == nil || *extName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extName"} 
    }
    routeValues["extName"] = *extName

    body, marshalErr := json.Marshal(*review)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e6e85b9d-aa70-40e6-aa28-d0fbf40b91a3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Review
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes a review
// ctx
// pubName (required): Name of the pubilsher who published the extension
// extName (required): Name of the extension
// reviewId (required): Id of the review which needs to be updated
func (client Client) DeleteReview(ctx context.Context, pubName *string, extName *string, reviewId *uint64) error {
    routeValues := make(map[string]string)
    if pubName == nil || *pubName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pubName"} 
    }
    routeValues["pubName"] = *pubName
    if extName == nil || *extName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extName"} 
    }
    routeValues["extName"] = *extName
    if reviewId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "reviewId"} 
    }
    routeValues["reviewId"] = strconv.FormatUint(*reviewId, 10)

    locationId, _ := uuid.Parse("e6e85b9d-aa70-40e6-aa28-d0fbf40b91a3")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates or Flags a review
// ctx
// reviewPatch (required): ReviewPatch object which contains the changes to be applied to the review
// pubName (required): Name of the pubilsher who published the extension
// extName (required): Name of the extension
// reviewId (required): Id of the review which needs to be updated
func (client Client) UpdateReview(ctx context.Context, reviewPatch *ReviewPatch, pubName *string, extName *string, reviewId *uint64) (*ReviewPatch, error) {
    if reviewPatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reviewPatch"}
    }
    routeValues := make(map[string]string)
    if pubName == nil || *pubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pubName"} 
    }
    routeValues["pubName"] = *pubName
    if extName == nil || *extName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extName"} 
    }
    routeValues["extName"] = *extName
    if reviewId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reviewId"} 
    }
    routeValues["reviewId"] = strconv.FormatUint(*reviewId, 10)

    body, marshalErr := json.Marshal(*reviewPatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e6e85b9d-aa70-40e6-aa28-d0fbf40b91a3")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReviewPatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// category (required)
func (client Client) CreateCategory(ctx context.Context, category *ExtensionCategory) (*ExtensionCategory, error) {
    if category == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "category"}
    }
    body, marshalErr := json.Marshal(*category)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("476531a3-7024-4516-a76a-ed64d3008ad6")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionCategory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all setting entries for the given user/all-users scope
// ctx
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
// key (optional): Optional key under which to filter all the entries
func (client Client) GetGalleryUserSettings(ctx context.Context, userScope *string, key *string) (*map[string]interface{}, error) {
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope
    if key != nil && *key != "" {
        routeValues["key"] = *key
    }

    locationId, _ := uuid.Parse("9b75ece3-7960-401c-848b-148ac01ca350")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]interface{}
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set all setting entries for the given user/all-users scope
// ctx
// entries (required): A key-value pair of all settings that need to be set
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
func (client Client) SetGalleryUserSettings(ctx context.Context, entries *map[string]interface{}, userScope *string) error {
    if entries == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "entries"}
    }
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope

    body, marshalErr := json.Marshal(*entries)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("9b75ece3-7960-401c-848b-148ac01ca350")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// keyType (required)
// expireCurrentSeconds (optional)
func (client Client) GenerateKey(ctx context.Context, keyType *string, expireCurrentSeconds *int) error {
    routeValues := make(map[string]string)
    if keyType == nil || *keyType == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "keyType"} 
    }
    routeValues["keyType"] = *keyType

    queryParams := url.Values{}
    if expireCurrentSeconds != nil {
        queryParams.Add("expireCurrentSeconds", strconv.Itoa(*expireCurrentSeconds))
    }
    locationId, _ := uuid.Parse("92ed5cf4-c38b-465a-9059-2f2fb7c624b5")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// keyType (required)
func (client Client) GetSigningKey(ctx context.Context, keyType *string) (*string, error) {
    routeValues := make(map[string]string)
    if keyType == nil || *keyType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "keyType"} 
    }
    routeValues["keyType"] = *keyType

    locationId, _ := uuid.Parse("92ed5cf4-c38b-465a-9059-2f2fb7c624b5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// extensionStatisticsUpdate (required)
// publisherName (required)
// extensionName (required)
func (client Client) UpdateExtensionStatistics(ctx context.Context, extensionStatisticsUpdate *ExtensionStatisticUpdate, publisherName *string, extensionName *string) error {
    if extensionStatisticsUpdate == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "extensionStatisticsUpdate"}
    }
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    body, marshalErr := json.Marshal(*extensionStatisticsUpdate)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("a0ea3204-11e9-422d-a9ca-45851cc41400")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// days (optional)
// aggregate (optional)
// afterDate (optional)
func (client Client) GetExtensionDailyStats(ctx context.Context, publisherName *string, extensionName *string, days *int, aggregate *ExtensionStatsAggregateType, afterDate *time.Time) (*ExtensionDailyStats, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if days != nil {
        queryParams.Add("days", strconv.Itoa(*days))
    }
    if aggregate != nil {
        queryParams.Add("aggregate", string(*aggregate))
    }
    if afterDate != nil {
        queryParams.Add("afterDate", (*afterDate).String())
    }
    locationId, _ := uuid.Parse("ae06047e-51c5-4fb4-ab65-7be488544416")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDailyStats
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] This route/location id only supports HTTP POST anonymously, so that the page view daily stat can be incremented from Marketplace client. Trying to call GET on this route should result in an exception. Without this explicit implementation, calling GET on this public route invokes the above GET implementation GetExtensionDailyStats.
// ctx
// publisherName (required): Name of the publisher
// extensionName (required): Name of the extension
// version (required): Version of the extension
func (client Client) GetExtensionDailyStatsAnonymous(ctx context.Context, publisherName *string, extensionName *string, version *string) (*ExtensionDailyStats, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version

    locationId, _ := uuid.Parse("4fa7adb6-ca65-4075-a232-5f28323288ea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ExtensionDailyStats
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Increments a daily statistic associated with the extension
// ctx
// publisherName (required): Name of the publisher
// extensionName (required): Name of the extension
// version (required): Version of the extension
// statType (required): Type of stat to increment
func (client Client) IncrementExtensionDailyStat(ctx context.Context, publisherName *string, extensionName *string, version *string, statType *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version

    queryParams := url.Values{}
    if statType == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "statType"}
    }
    queryParams.Add("statType", *statType)
    locationId, _ := uuid.Parse("4fa7adb6-ca65-4075-a232-5f28323288ea")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// version (required)
func (client Client) GetVerificationLog(ctx context.Context, publisherName *string, extensionName *string, version *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName
    if version == nil || *version == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "version"} 
    }
    routeValues["version"] = *version

    locationId, _ := uuid.Parse("c5523abe-b843-437f-875b-5833064efe4d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

