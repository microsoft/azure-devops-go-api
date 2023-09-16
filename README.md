[![Go](https://github.com/microsoft/azure-devops-go-api/workflows/Go/badge.svg)](https://github.com/microsoft/azure-devops-go-api/actions)
[![Build Status](https://dev.azure.com/mseng/vsts-cli/_apis/build/status/microsoft.azure-devops-go-api?branchName=dev)](https://dev.azure.com/mseng/vsts-cli/_build/latest?definitionId=9110&branchName=dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/microsoft/azure-devops-go-api)](https://goreportcard.com/report/github.com/microsoft/azure-devops-go-api)

# Azure DevOps Go API

This repository contains Go APIs for interacting with and managing Azure DevOps.

## Get started

### Authenticate Using Azure PAT

To use the API, establish a connection using a [personal access token](https://docs.microsoft.com/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops) and the URL to your Azure DevOps organization. Then get a client using the connection and make API calls.

```go
package main

import (
	"context"
	"log"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
)

func main() {
	organizationUrl := "https://dev.azure.com/myorg" // todo: replace value with your organization url
	personalAccessToken := "XXXXXXXXXXXXXXXXXXXXXXX" // todo: replace value with your PAT

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	// Get first page of the list of team projects for your organization
	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {
		// Log the page of team project names
		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			index++
		}

		// if continuationToken has a value, then there is at least one more page of projects to get
		if responseValue.ContinuationToken != "" {

			continuationToken, err := strconv.Atoi(responseValue.ContinuationToken)
			if err != nil {
				log.Fatal(err)
			}

			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}
}
```

### Authenticate Using Azure AD Access Tokens (OAuth)

When Authenticating using Azure AD Access Tokens, you can use either a default token credential or a connection chain token credential.
- [Read more here about using Service Principals and Managed identity to call Azure DevOps](https://learn.microsoft.com/en-us/azure/devops/integrate/get-started/authentication/service-principal-managed-identity?toc=%2Fazure%2Fdevops%2Fmarketplace-extensibility%2Ftoc.json&view=azure-devops)


#### Using a default token credential
```go

 connection := azuredevops.NewOAuthConnectionDefault(organizationUrl, azuredevops.TokenOptions{})
 ctx := context.Background()

 // Create a client to interact with the Core area
 coreClient, err := core.NewClient(ctx, connection)
 if err != nil {
 log.Fatal(err)
 }
 
 // The rest of the code below remains the same
```


#### Using a connection chain token credential

```go
package main

import (
	"context"
	"log"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
)

func main() {
	organizationUrl := "https://dev.azure.com/MyOrg" // todo: replace value with your organization url

	connection := azuredevops.NewOAuthConnectionChainToken(organizationUrl, azuredevops.TokenOptions{})


	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	// Get first page of the list of team projects for your organization
	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {
		// Log the page of team project names
		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			index++
		}

		// if continuationToken has a value, then there is at least one more page of projects to get
		if responseValue.ContinuationToken != "" {

			continuationToken, err := strconv.Atoi(responseValue.ContinuationToken)
			if err != nil {
				log.Fatal(err)
			}

			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}
}

```


### Authenticate using NewOauthGenericConnection

```go

// if you want to make the http call to get the token yourself, or you want to use a different library to get the token
connection := azuredevops.NewOAuthGenericConnection(organizationUrl, OAuthToken)
ctx := context.Background()

// Create a client to interact with the Core area
coreClient, err := core.NewClient(ctx, connection)
if err != nil {
log.Fatal(err)
}
// Rest of the code above remains the same

```

## API documentation

This Go library provides a thin wrapper around the Azure DevOps REST APIs. See the [Azure DevOps REST API reference](https://docs.microsoft.com/en-us/rest/api/azure/devops/?view=azure-devops-rest-5.1) for details on calling different APIs.

# Contributing

This project welcomes contributions and suggestions. Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
