package dto

type ExecuteJobsRequest struct {
	JobIDs []int `json:"jobIds,omitempty"`
}
