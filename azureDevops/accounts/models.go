// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package accounts

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "time"
)

type Account struct {
    // Identifier for an Account
    AccountId *uuid.UUID `json:"accountId,omitempty"`
    // Name for an account
    AccountName *string `json:"accountName,omitempty"`
    // Owner of account
    AccountOwner *uuid.UUID `json:"accountOwner,omitempty"`
    // Current account status
    AccountStatus *AccountStatus `json:"accountStatus,omitempty"`
    // Type of account: Personal, Organization
    AccountType *AccountType `json:"accountType,omitempty"`
    // Uri for an account
    AccountUri *string `json:"accountUri,omitempty"`
    // Who created the account
    CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
    // Date account was created
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    HasMoved *bool `json:"hasMoved,omitempty"`
    // Identity of last person to update the account
    LastUpdatedBy *uuid.UUID `json:"lastUpdatedBy,omitempty"`
    // Date account was last updated
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Namespace for an account
    NamespaceId *uuid.UUID `json:"namespaceId,omitempty"`
    NewCollectionId *uuid.UUID `json:"newCollectionId,omitempty"`
    // Organization that created the account
    OrganizationName *string `json:"organizationName,omitempty"`
    // Extended properties
    Properties *interface{} `json:"properties,omitempty"`
    // Reason for current status
    StatusReason *string `json:"statusReason,omitempty"`
}

type AccountCreateInfoInternal struct {
    AccountName *string `json:"accountName,omitempty"`
    Creator *uuid.UUID `json:"creator,omitempty"`
    Organization *string `json:"organization,omitempty"`
    Preferences *AccountPreferencesInternal `json:"preferences,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    ServiceDefinitions *[]azureDevops.KeyValuePair `json:"serviceDefinitions,omitempty"`
}

type AccountPreferencesInternal struct {
    Culture *interface{} `json:"culture,omitempty"`
    Language *interface{} `json:"language,omitempty"`
    TimeZone *interface{} `json:"timeZone,omitempty"`
}

type AccountStatus string

type AccountType string

type AccountUserStatus string
