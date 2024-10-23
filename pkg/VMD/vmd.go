package vmd

import (
	driver "github.com/hvishnoi1/VideoMetaData/pkg/drivers"
	mediainfo_driver "github.com/hvishnoi1/VideoMetaData/pkg/drivers/MediaInfo"
)

func NewMediaInfoDriver() driver.Driver {
	return &mediainfo_driver.MediaInfoDriver{}
}

func NewFFProbeDriver() driver.Driver {
	// TODO add ffprobe driver
	return nil
}
