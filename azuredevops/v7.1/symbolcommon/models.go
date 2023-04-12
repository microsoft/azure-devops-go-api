// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package symbolcommon

// [Flags] Defines the level of debug-related information inside the .pdb file. These values can be combined together (bitwise OR'ed) to create a customized level.
type DebugInformationLevel string

type debugInformationLevelValuesType struct {
	None               DebugInformationLevel
	Binary             DebugInformationLevel
	Publics            DebugInformationLevel
	TraceFormatPresent DebugInformationLevel
	TypeInfo           DebugInformationLevel
	LineNumbers        DebugInformationLevel
	GlobalSymbols      DebugInformationLevel
	Private            DebugInformationLevel
	SourceIndexed      DebugInformationLevel
}

var DebugInformationLevelValues = debugInformationLevelValuesType{
	// If set, the .pdb file contains no debug information.
	None: "none",
	// If set, the .pdb file contains debug information which is binary.
	Binary: "binary",
	// If set, the .pdb file contains public symbols.
	Publics: "publics",
	// If set, the .pdb file contains trace format.
	TraceFormatPresent: "traceFormatPresent",
	// If set, the .pdb file contains type information.
	TypeInfo: "typeInfo",
	// If set, the .pdb file contains line number information.
	LineNumbers: "lineNumbers",
	// If set, the .pdb file contains symbol information.
	GlobalSymbols: "globalSymbols",
	// If set, the .pdb file contains public symbols and has type, line number and symbol information.
	Private: "private",
	// If set, the .pdb file supports the source server.
	SourceIndexed: "sourceIndexed",
}
