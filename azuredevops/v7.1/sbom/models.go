// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package sbom


// This holds the configuration for the ManifestTool. The values in this file are populated from the command line, config file and default.
type Configuration struct {
    // Additional set of command-line arguments for Component Detector.
    AdditionalComponentDetectorArgs interface{} `json:"additionalComponentDetectorArgs,omitempty"`
    // The folder containing the build components and packages.
    BuildComponentPath interface{} `json:"buildComponentPath,omitempty"`
    // The root folder of the drop directory to validate or generate.
    BuildDropPath interface{} `json:"buildDropPath,omitempty"`
    // Full file name of a list file that contains all files to be validated.
    BuildListFile interface{} `json:"buildListFile,omitempty"`
    // The path of the signed catalog file used to validate the manifest.json.
    CatalogFilePath interface{} `json:"catalogFilePath,omitempty"`
    // The json file that contains the configuration for the DropValidator.
    ConfigFilePath interface{} `json:"configFilePath,omitempty"`
    // Comma separated list of docker image names or hashes to be scanned for packages, ex: ubuntu:16.04, 56bab49eef2ef07505f6a1b0d5bd3a601dfc3c76ad4460f24c91d6fa298369ab.
    DockerImagesToScan interface{} `json:"dockerImagesToScan,omitempty"`
    // Full file path to a file that contains list of external SBOMs to be included as External document reference.
    ExternalDocumentReferenceListFile interface{} `json:"externalDocumentReferenceListFile,omitempty"`
    // The Hash algorithm to use while verifying the hash value of a file.
    HashAlgorithm interface{} `json:"hashAlgorithm,omitempty"`
    // If set, will not fail validation on the files presented in Manifest but missing on the disk.
    IgnoreMissing interface{} `json:"ignoreMissing,omitempty"`
    // The root folder where the generated manifest (and other files like bsi.json) files will be placed. By default we will generate this folder in the same level as the build drop with the name '_manifest'
    ManifestDirPath interface{} `json:"manifestDirPath,omitempty"`
    // A list of name and version of the manifest that we are generating.
    ManifestInfo interface{} `json:"manifestInfo,omitempty"`
    // The action currently being performed by the manifest tool.
    ManifestToolAction *ManifestToolActions `json:"manifestToolAction,omitempty"`
    // The name of the package this SBOM represents.
    PackageName interface{} `json:"packageName,omitempty"`
    // The version of the package this SBOM represents.
    PackageVersion interface{} `json:"packageVersion,omitempty"`
    // The number of parallel threads to use for the workflows.
    Parallelism interface{} `json:"parallelism,omitempty"`
    // If you're downloading only a part of the drop using the '-r' or 'root' parameter in the drop client, specify the same string value here in order to skip validating paths that are not downloaded.
    RootPathFilter interface{} `json:"rootPathFilter,omitempty"`
    // If specified, we will store the generated telemetry for the execution of the SBOM tool at this path.
    TelemetryFilePath interface{} `json:"telemetryFilePath,omitempty"`
    // If set, will validate the manifest using the signed catalog file.
    ValidateSignature interface{} `json:"validateSignature,omitempty"`
}

// Encapsulates a configuration setting to provide metadata about the setting source and type.
type ConfigurationSetting struct {
    // The source where this setting came from.
    Source *string `json:"source,omitempty"`
    // The actual value of the setting.
    Value interface{} `json:"value,omitempty"`
}

// Used to provide the filename and hash of the SBOM file to be added to the catalog file.
type FileHash struct {
    // The filename of the SBOM.
    FileName *string `json:"fileName,omitempty"`
    // The string hash of the SBOM file.
    Hash *string `json:"hash,omitempty"`
    // The HashAlgorithmName used to generate the hash of the file.
    HashAlgorithmName *HashAlgorithmName `json:"hashAlgorithmName,omitempty"`
}

// Defines a manifest name and version.
type ManifestInfo struct {
    // The name of the manifest.
    Name *string `json:"name,omitempty"`
    // The version of the manifest.
    Version *string `json:"version,omitempty"`
}

// [Flags]
type ManifestToolActions string

type manifestToolActionsValuesType struct {
    None ManifestToolActions
    Validate ManifestToolActions
    Generate ManifestToolActions
    All ManifestToolActions
}

var ManifestToolActionsValues = manifestToolActionsValuesType{
    None: "none",
    Validate: "validate",
    Generate: "generate",
    All: "all",
}

// Enum that represents the result of a signing submission or status request.
type Result string

type resultValuesType struct {
    Success Result
    Failure Result
    InProgress Result
}

var ResultValues = resultValuesType{
    Success: "success",
    Failure: "failure",
    InProgress: "inProgress",
}

// Represents a SBOM file object and contains additional properties related to the file.
type SBOMFile struct {
    // The size of the SBOM file in bytes.
    FileSizeInBytes *uint64 `json:"fileSizeInBytes,omitempty"`
    // The path where the final generated SBOM is placed.
    SbomFilePath *string `json:"sbomFilePath,omitempty"`
    // The name and version of the format of the generated SBOM.
    SbomFormatName *ManifestInfo `json:"sbomFormatName,omitempty"`
}

// The telemetry that is logged to a file/console for the given SBOM execution.
type SBOMTelemetry struct {
    // All available bsi data from the task build execution which includes build and system environment variables like repository and build information.
    BsiData *map[string]string `json:"bsiData,omitempty"`
    // The source of the bsi data.
    BsiSource *string `json:"bsiSource,omitempty"`
    // The end to end results of the extension task.
    E2ETaskResult *string `json:"e2ETaskResult,omitempty"`
    // A list of ConfigurationSetting`1 representing each input parameter used in the validation.
    Parameters *Configuration `json:"parameters,omitempty"`
    // The result of the execution
    Result *string `json:"result,omitempty"`
    // A list of the SBOM formats and related file properties that was used in the generation/validation of the SBOM.
    SbomFormatsUsed *[]SBOMFile `json:"sbomFormatsUsed,omitempty"`
    // Any internal switches and their value that were used during the execution. A switch can be something that was provided through a configuraiton or an environment variable.
    Switches *map[string]interface{} `json:"switches,omitempty"`
    // Error messages that came from the extension task.
    TaskErrorMessage *string `json:"taskErrorMessage,omitempty"`
    // The unique id for this telemetry
    TelemetryId *string `json:"telemetryId,omitempty"`
    // The result of the tool as a numeric value.
    ToolExecutionResult *int `json:"toolExecutionResult,omitempty"`
}

// The base reponse object for all responses from the signing api.
type SignResponseBase struct {
    // The customer correlation id that is sent to ESRP for correlating the current request to ESRP.
    CustomerCorrelationId *string `json:"customerCorrelationId,omitempty"`
    // If this is an error response, it will have more information about the error.
    ErrorInfo *string `json:"errorInfo,omitempty"`
    // The result of the response.
    Result *Result `json:"result,omitempty"`
}

// The response returned by the sign status api.
type SignStatusResponse struct {
    // The customer correlation id that is sent to ESRP for correlating the current request to ESRP.
    CustomerCorrelationId *string `json:"customerCorrelationId,omitempty"`
    // If this is an error response, it will have more information about the error.
    ErrorInfo *string `json:"errorInfo,omitempty"`
    // The result of the response.
    Result *Result `json:"result,omitempty"`
    // The pre-signed download url used to download the signed catalog file.
    DownloadUrl *string `json:"downloadUrl,omitempty"`
}
