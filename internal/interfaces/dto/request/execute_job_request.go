package request

type ExecuteJobsRequest struct {
	JobIDs []int `json:"jobIds,omitempty"`
}
