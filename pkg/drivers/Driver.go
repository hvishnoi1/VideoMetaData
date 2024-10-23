package driver

import "github.com/hvishnoi1/VideoMetaData/pkg/VMD/structs"

type Driver interface {
	// Initializes the driver. This is a suitable function to check if required binaries exist and their paths
	Init(string, *[]string) error

	// Fetches metadata for provided file path or URL
	GetVideoMetadata(string) (*structs.VideoMetaData, error)

	// Get version of the driver
	Version() string
}
