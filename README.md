# Azure DevOps Go API
This repository contains Go APIs for interacting with and managing Azure DevOps.

## Get started
To use the API, establish a connection using a [personal access token](https://docs.microsoft.com/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops) and the URL to your Azure DevOps organization. Then get a client using the connection and make API calls.

```
package main

import (
    "context"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "github.com/microsoft/azure-devops-go-api/azureDevops/core"
    "log"
)

func main() {
    organizationUrl := "https://dev.azure.com/myorg"  // todo: replace value with your organization url
    personalAccessToken := "XXXXXXXXXXXXXXXXXXXXXXX"  // todo: replace value with your PAT

	// Create a connection to your organization
	connection := azureDevops.NewConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, *connection)
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all team projects
	teamProjectReferences, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the list of team project names for your organization
	for index, teamProjectReference := range *teamProjectReferences {
		log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
	}
}
```

## API documentation

This Go library provides a thin wrapper around the Azure DevOps REST APIs. See the [Azure DevOps REST API reference](https://docs.microsoft.com/en-us/rest/api/azure/devops/?view=azure-devops-rest-5.1) for details on calling different APIs.


# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
