package mediainfo_driver

const (
	TYPE_GENERAL = "General"
	TYPE_VIDEO   = "Video"
	TYPE_AUDIO   = "Audio"

	WHICH     = "which"
	MEDIAINFO = "mediainfo"
	FFPROBE   = "ffprobe"

	ARG_VERSION = "--version"
	ARG_OUTPUT  = "--Output=JSON"

	ERR_DRIVER_NOT_INITIALIZED = "error: Driver not initialized, intitialize driver by calling Init(...)"
	ERR_DIR_PATH_PROVIDED      = "error: Provided path is a directory, expected executable file path"
	ERR_EXEC_PERMISSION        = "error: The file is not execuatable"
)
