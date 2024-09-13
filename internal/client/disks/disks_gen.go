// Package disks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package disks

import (
	"time"
)

// Disk defines model for disk.
type Disk struct {
	Id        DiskId `json:"id"`
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	SizeGB    int    `json:"sizeGB"`
}

// DiskDetails defines model for diskDetails.
type DiskDetails struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        DiskId    `json:"id"`
	MountPath string    `json:"mountPath"`
	Name      string    `json:"name"`
	ServiceId *string   `json:"serviceId,omitempty"`
	SizeGB    int       `json:"sizeGB"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DiskId defines model for diskId.
type DiskId = string

// DiskPATCH defines model for diskPATCH.
type DiskPATCH struct {
	MountPath *string `json:"mountPath,omitempty"`
	Name      *string `json:"name,omitempty"`
	SizeGB    *int    `json:"sizeGB,omitempty"`
}

// DiskPOST defines model for diskPOST.
type DiskPOST struct {
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	ServiceId string `json:"serviceId"`
	SizeGB    int    `json:"sizeGB"`
}

// DiskIdQuery defines model for diskIdQuery.
type DiskIdQuery = []DiskId
