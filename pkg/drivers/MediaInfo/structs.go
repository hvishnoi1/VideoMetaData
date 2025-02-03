package mediainfo_driver

type creatingLibrary struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type media struct {
	Ref   string  `json:"@ref"`
	Track []track `json:"track"`
}

// Extend this struct to add more fields
// Since this can grow very large, keep the fields sorted
type track struct {
	AudioCount            string `json:"AudioCount"`
	BitRate               string `json:"BitRate_Nominal"`
	CodecID               string `json:"CodecID"`
	CodecIDCompatible     string `json:"CodecID_Compatible"`
	Channels              string `json:"Channels"`
	Duration              string `json:"Duration"`
	FileExtension         string `json:"FileExtension"`
	FileModifiedDate      string `json:"File_Modified_Date"`
	FileModifiedDateLocal string `json:"File_Modified_Date_Local"`
	FileSize              string `json:"FileSize"`
	Format                string `json:"Format"`
	FrameCount            string `json:"FrameCount"`
	FrameRate             string `json:"Framerate"`
	Height                string `json:"Height"`
	ID                    string `json:"ID"`
	Language              string `json:"Language"`
	OverallBitRate        string `json:"OverallBitrate"`
	SampledHeight         string `json:"Sampled_Height"`
	SampledWidth          string `json:"Sampled_Width"`
	StoredHeight          string `json:"Stored_Height"`
	StoredWidth           string `json:"Stored_Width"`
	Type                  string `json:"@type"`
	VideoCount            string `json:"VideoCount"`
	Width                 string `json:"Width"`
}

type MediaInfoOutput struct {
	CreatingLibrary creatingLibrary `json:"creatingLibrary"`
	Media           media           `json:"media"`
}
