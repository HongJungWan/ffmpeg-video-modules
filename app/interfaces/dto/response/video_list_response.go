package response

type VideoJobDetail struct {
	ID             int    `json:"id"`
	JobType        string `json:"jobType"`
	Parameters     string `json:"parameters"`
	Status         string `json:"status"`
	ResultFilePath string `json:"resultFilePath"`
}

type FinalVideoDetail struct {
	ID           int    `json:"id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"filePath"`
	DownloadLink string `json:"downloadLink"`
}

type VideoDetailResponse struct {
	ID               int               `json:"id"`
	Filename         string            `json:"filename"`
	FilePath         string            `json:"filePath"`
	Duration         int               `json:"duration"`
	Status           string            `json:"status"`
	CreatedAt        string            `json:"createdAt"`
	UpdatedAt        string            `json:"updatedAt"`
	TrimJobs         []VideoJobDetail  `json:"trimJobs,omitempty"`
	ConcatJobs       []VideoJobDetail  `json:"concatJobs,omitempty"`
	FinalVideoDetail *FinalVideoDetail `json:"finalVideoDetail,omitempty"`
}
