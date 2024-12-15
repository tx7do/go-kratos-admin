package task

const (
	BackupTaskType = "backup"
)

type BackupTaskData struct {
	Name string `json:"name"`
}
