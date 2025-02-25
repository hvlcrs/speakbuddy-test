package response

// @Description Audio API response information
type Audio struct {
	UserID   string `json:"user_id" binding:"required" uuid:"true"`
	PhraseID string `json:"phrase_id" binding:"required" uuid:"true"`
	Name     string `json:"file_name"`
	Format   string `json:"audio_format"`
}
