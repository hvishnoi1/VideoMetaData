package utils

import "strings"

const UNKNOWN_ASPECT_RATIO = "Unknown"

func AspectRatioFloatToString(width, height int) string {
	if height == 0 {
		return UNKNOWN_ASPECT_RATIO
	}

	switch float64(width) / float64(height) {
	case 1.:
		return "1:1"
	case 4. / 3.:
		return "4:3"
	case 3. / 4.:
		return "3:4"
	case 3. / 2.:
		return "3:2"
	case 2. / 3.:
		return "2:3"
	case 14. / 9.:
		return "14:9"
	case 16. / 9., 854. / 480., 854. / 470.:
		return "16:9"
	case 16. / 10.:
		return "16:10"
	case 9. / 16.:
		return "9:16"
	default:
		return UNKNOWN_ASPECT_RATIO
	}
}

// Converts common codec name to its formal name
func CommonCodecToFormal(codec string) string {
	codec = strings.ToLower(codec)
	switch codec {
	case "avc":
		{
			return "h.264"
		}
	case "hevc":
		{
			return "h.265"
		}
	case "mp4":
		{
			return "mpeg-4"
		}
	}

	return codec
}
