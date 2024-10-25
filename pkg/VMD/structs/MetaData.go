package structs

type VideoMetaData struct {
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	AspectRatio string `json:"aspectRatio,omitempty"`
	FileSize    int64  `json:"fileSize,omitempty"`

	VideoDuration  float32 `json:"videoDuration,omitempty"`
	VideoCodec     string  `json:"videoCodec,omitempty"`
	VideoBitRate   float32 `json:"videoBitRate,omitempty"`
	VideoFrameRate float32 `json:"videoFrameRate,omitempty"`

	AudioCodec     string  `json:"audioCodec,omitempty"`
	AudioBitRate   float32 `json:"audioBitRate,omitempty"`
	AudioFrameRate float32 `json:"audioFrameRate,omitempty"`
	AudioChannels  int     `json:"audioChannels,omitempty"`
	AudioDuration  float32 `json:"audioDuration,omitempty"`
}
