package service

import (
	"log"
	"os"
	"slices"
	"speakbuddy/models"
	"speakbuddy/pkg/configs"
	"speakbuddy/pkg/db"

	"errors"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm/clause"
)

func ValidateAudioFormat(format string) error {
	if !slices.Contains(configs.AppSetting.FileAllowExts, format) {
		return errors.New("unsupported file format")
	}
	return nil
}

// Delete source file
// On production, this should be done in a separate queue-worker with proper locking mechanism
func CleanupLocalCache(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func SaveAudio(audio models.Audio) error {
	db.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "phrase_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"format", "name", "path"}),
	}).Create(&audio)
	return nil
}

func GetAudioByUserAndPhraseID(userID string, phraseID string) (models.Audio, error) {
	audio := models.Audio{}
	if err := db.Conn.Where("user_id = ? AND phrase_id = ?", userID, phraseID).First(&audio).Error; err != nil {
		log.Printf("[ERROR] Failed to get audio metadata: %s", err)
		return models.Audio{}, err
	}
	return audio, nil
}

// Transcode audio file
// On production, this should be done in a separate queue-worker with proper locking mechanism
func TranscodeAudio(audio models.Audio, format string) (string, error) {
	destination := configs.AppSetting.FileSavePath + audio.UserID + "/" + audio.PhraseID + "/" + audio.Name + format
	err := ffmpeg.Input(audio.Path).
		Output(destination).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		log.Printf("[ERROR] Failed to transcode audio: %s", err)
		return "", err
	}
	return destination, nil
}

// Transcode audio file and delete source file in background
// On production, this should be done in a separate queue-worker with proper locking mechanism
func TranscodeAudioAndCleanup(audio models.Audio, format string) (string, error) {
	TranscodeAudio(audio, format)
	CleanupLocalCache(audio.Path)

	return audio.Path, nil
}
