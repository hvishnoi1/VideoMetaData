package structs

type VideoMetaData struct {
	Width       int     `json:"width,omitempty"`
	Height      int     `json:"height,omitempty"`
	AspectRatio string  `json:"aspectRatio,omitempty"`
	FileSize    float64 `json:"fileSize,omitempty"` // MebiByte (MiB)

	VideoDuration  float64 `json:"videoDuration,omitempty"` // seconds (s)
	VideoCodec     string  `json:"videoCodec,omitempty"`
	VideoBitRate   float64 `json:"videoBitRate,omitempty"` // KiloBit per second (kbps)
	VideoFrameRate float64 `json:"videoFrameRate,omitempty"`

	AudioCodec     string  `json:"audioCodec,omitempty"`
	AudioBitRate   float64 `json:"audioBitRate,omitempty"` // KiloBit per second (kbps)
	AudioFrameRate float64 `json:"audioFrameRate,omitempty"`
	AudioChannels  int     `json:"audioChannels,omitempty"`
	AudioDuration  float64 `json:"audioDuration,omitempty"` // seconds (s)
}
