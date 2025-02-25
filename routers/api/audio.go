package api

import (
	"log"
	"path/filepath"
	"speakbuddy/models"
	"speakbuddy/pkg/configs"
	"speakbuddy/routers/response"
	"speakbuddy/service"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"
)

// UploadAudio godoc
// @Summary Upload audio file
// @Description Upload audio file
// @Schemes
// @Tags audio
// @Produce json
// @Success 200 {object} response.Audio
// @Param user_id path string true "User ID" example("test")
// @Param phrase_id path string true "Phrase ID" example("test")
// @Param audio formData file true "Audio file"
// @Accept multipart/form-data
// @Router /audio/user/{user_id}/phrase/{phrase_id} [post]
func UploadAudio(c *gin.Context) {
	audio := models.Audio{}
	res := response.Audio{}

	// Convert content length to MB and validate max file size
	if c.Request.ContentLength/(1<<20) > int64(configs.AppSetting.FileMaxSize) {
		c.JSON(400, gin.H{"msg": "Max file size exceeded"})
		return
	}

	file, err := c.FormFile("audio")
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&audio); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	destination := configs.AppSetting.FileSavePath + audio.UserID + "/" + audio.PhraseID + "/"
	audio.Path = destination + file.Filename

	// Save temporary file to local storage
	// On production, the file should be offloaded into cloud storage or persistent shared storage
	// This way the application can be scaled horizontally as long as the background worker has acccess to the storage
	if err := c.SaveUploadedFile(file, audio.Path); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// Retrieve temporary file metadata
	audio.Name = strings.TrimSuffix(filepath.Base(audio.Path), filepath.Ext(audio.Path))
	audio.Format = filepath.Ext(audio.Path)

	// Check supported file format
	if service.ValidateAudioFormat(audio.Format) != nil {
		// On production, this can be replaced by a cron to handle file cleanup or
		// setting up object life cycle in the cloud storage
		if service.CleanupLocalCache(audio.Path) != nil {
			log.Printf("[ERROR] Failed to remove file: %v", err)
			return
		}
		c.JSON(400, gin.H{"msg": "File format is not supported"})
		return
	}

	if audio.Format != configs.AppSetting.FileTargetExt {
		go service.TranscodeAudioAndCleanup(audio, configs.AppSetting.FileTargetExt)
	}

	// Save audio metadata to database
	audio.Format = configs.AppSetting.FileTargetExt
	audio.Path = destination + audio.Name + configs.AppSetting.FileTargetExt
	service.SaveAudio(audio)

	if err := copier.Copy(&res, &audio); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, res)
}

// GetAudio godoc
// @Summary Get transcoded audio file with specific format
// @Description Get transcoded audio file with specific format
// @Schemes
// @Tags audio
// @Success 200 {file} file
// @Param user_id path string true "User ID" example("test")
// @Param phrase_id path string true "Phrase ID" example("test")
// @Param audio_format path string true "Audio format" example(".mp3")
// @Router /audio/user/{user_id}/phrase/{phrase_id}/{audio_format} [get]
func GetAudio(c *gin.Context) {
	audio := models.Audio{}
	res := response.Audio{}

	if err := c.ShouldBindUri(&audio); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// Check supported file format
	if err := service.ValidateAudioFormat(audio.Format); err != nil {
		c.JSON(400, gin.H{"msg": "File format is not supported"})
		return
	}

	// Validate audio metadata from database
	result, err := service.GetAudioByUserAndPhraseID(audio.UserID, audio.PhraseID)
	if err != nil {
		c.JSON(404, gin.H{"msg": err.Error()})
		return
	}

	if err := copier.Copy(&res, &result); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	res.Format = audio.Format
	filePath := result.Path

	// Check if audio file format is different from requested format
	if audio.Format != result.Format {
		// Transcode audio file
		filePath, err = service.TranscodeAudio(result, audio.Format)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		defer service.CleanupLocalCache(filePath)
	}

	c.FileAttachment(filePath, result.Name)
}
