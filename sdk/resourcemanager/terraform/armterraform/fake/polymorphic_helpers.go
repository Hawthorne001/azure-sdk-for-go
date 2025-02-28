// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/terraform/armterraform"
)

func unmarshalBaseExportModelClassification(rawMsg json.RawMessage) (armterraform.BaseExportModelClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armterraform.BaseExportModelClassification
	switch m["type"] {
	case string(armterraform.TypeExportQuery):
		b = &armterraform.ExportQuery{}
	case string(armterraform.TypeExportResource):
		b = &armterraform.ExportResource{}
	case string(armterraform.TypeExportResourceGroup):
		b = &armterraform.ExportResourceGroup{}
	default:
		b = &armterraform.BaseExportModel{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
