package domain

type UserVideoPermission struct {
	ID             int            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID         int            `gorm:"not null" json:"userId"`
	VideoID        int            `gorm:"not null" json:"videoId"`
	PermissionType PermissionType `gorm:"type:ENUM('owner','editor','viewer');not null" json:"permissionType"`
}

type PermissionType string

const (
	Owner  PermissionType = "owner"
	Editor PermissionType = "editor"
	Viewer PermissionType = "viewer"
)

func NewUserVideoPermission(userID, videoID int, permissionType PermissionType) *UserVideoPermission {
	return &UserVideoPermission{
		UserID:         userID,
		VideoID:        videoID,
		PermissionType: permissionType,
	}
}
