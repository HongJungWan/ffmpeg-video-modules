package response

type VideoDownloadResponse struct {
	ID           int    `json:"id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"filePath"`
	DownloadLink string `json:"downloadLink"`
}
