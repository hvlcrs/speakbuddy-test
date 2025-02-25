package models

// @Description Audio model information
type Audio struct {
	UserID   string `uri:"user_id" example:"test" binding:"required" gorm:"primaryKey"`
	PhraseID string `uri:"phrase_id" example:"test" binding:"required" gorm:"primaryKey"`
	Format   string `uri:"audio_format" example:".mp3"`
	Name     string
	Path     string
}
