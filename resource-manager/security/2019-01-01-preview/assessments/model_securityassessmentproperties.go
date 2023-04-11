package assessments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAssessmentProperties struct {
	AdditionalData  *map[string]string `json:"additionalData,omitempty"`
	DisplayName     *string            `json:"displayName,omitempty"`
	Links           *AssessmentLinks   `json:"links,omitempty"`
	ResourceDetails ResourceDetails    `json:"resourceDetails"`
	Status          AssessmentStatus   `json:"status"`
}
