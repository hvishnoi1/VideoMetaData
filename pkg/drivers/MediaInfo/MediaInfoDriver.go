package mediainfo_driver

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/hvishnoi1/VideoMetaData/pkg/VMD/structs"
	"github.com/hvishnoi1/VideoMetaData/pkg/utils"
)

type MediaInfoDriver struct {
	executablePath string
	args           []string
	driver.Driver
}

// Checks if the required binaries are installed.
// Takes a preferred binary path, if not present it will search in the default PATH. If none is found, returns an error.
func (m *MediaInfoDriver) Init(preferredBinaryPath string, callArgs *[]string) error {
	if preferredBinaryPath == "" {
		// Find binary in default PATH and check if it exists
		cmd := exec.Command(WHICH, MEDIAINFO)

		out, err := cmd.Output()

		if err != nil {
			return err
		}

		m.executablePath = strings.TrimSuffix(string(out), "\n")
	} else {
		// Check if preferred binary exists

		if fileStat, err := os.Stat(preferredBinaryPath); err != nil {
			return err
		} else if fileStat.IsDir() {
			return errors.New(ERR_DIR_PATH_PROVIDED)
		} else if fileStat.Mode()&0111 == 0 {
			return errors.New(ERR_EXEC_PERMISSION)
		}

		m.executablePath = preferredBinaryPath
	}

	if callArgs != nil {
		m.args = *callArgs
	}
	m.args = append(m.args, ARG_OUTPUT)

	return nil
}

// Get video meta data using the specified driver
func (m *MediaInfoDriver) GetVideoMetadata(file string) (*structs.VideoMetaData, error) {
	if m == nil || m.executablePath == "" {
		return nil, errors.New(ERR_DRIVER_NOT_INITIALIZED)
	}

	args := append(m.args, file)
	cmd := exec.Command(m.executablePath, args...)

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var mediaInfo MediaInfoOutput

	err = json.Unmarshal(out, &mediaInfo)

	if err != nil {
		return nil, err
	}

	videoMetaData := &structs.VideoMetaData{}

	mapMediaInfoJsonToVideoMetaData(mediaInfo, videoMetaData)

	return videoMetaData, nil
}

// Returns the driver version
func (m *MediaInfoDriver) Version() string {
	if m == nil || m.executablePath == "" {
		return errors.New(ERR_DRIVER_NOT_INITIALIZED).Error()
	}
	cmd := exec.Command(m.executablePath, ARG_VERSION)

	out, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return string(out)
}

func mapMediaInfoJsonToVideoMetaData(mediaInfoOutput MediaInfoOutput, videoMetadata *structs.VideoMetaData) {
	for _, track := range mediaInfoOutput.Media.Track {
		switch track.Type {
		case TYPE_GENERAL:
			videoMetadata.FileSize = utils.ParseFloat64(track.FileSize) / 1024.0 / 1024.0 // Converting Bytes to MebiBytes (MiB)
		case TYPE_VIDEO:
			videoMetadata.Height = utils.ParseInt(track.Height)
			videoMetadata.Width = utils.ParseInt(track.Width)
			videoMetadata.AspectRatio = utils.AspectRatioFloatToString(videoMetadata.Width, videoMetadata.Height)
			videoMetadata.VideoDuration = utils.ParseFloat64(track.Duration)
			videoMetadata.VideoCodec = utils.CommonCodecToFormal(track.Format)      // Converting common codec names to formal codec names
			videoMetadata.VideoBitRate = utils.ParseFloat64(track.BitRate) / 1000.0 // Converting bps to kbps
			videoMetadata.VideoFrameRate = utils.ParseFloat64(track.FrameRate)
		case TYPE_AUDIO:
			videoMetadata.AudioCodec = strings.ToLower(track.Format)
			videoMetadata.AudioBitRate = utils.ParseFloat64(track.BitRate) / 1000 // Converting bps to kbps
			videoMetadata.AudioFrameRate = utils.ParseFloat64(track.FrameRate)
			videoMetadata.AudioChannels = utils.ParseInt(track.Channels)
			videoMetadata.AudioDuration = utils.ParseFloat64(track.Duration)
		}
	}
}
