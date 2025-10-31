package models

const (
	ImageResizeTask= "image:resize"
	ImageConvertToPNGTask="image:convertPNG"
)
type TaskImagePayload struct {
	JobID      int            `json:"job_id"`
	Image   Imageconfig `json:"image_config"`
	OuputPath  string         `json:"output_path"`
	Parameters map[string]any `json:"parameters"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
