package configs

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigInit(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("FILE_SAVE_PATH", "/tmp/files")
	os.Setenv("FILE_MAX_SIZE", "100")
	os.Setenv("FILE_ALLOW_EXTS", ".mp4,.mp3")
	os.Setenv("FILE_TARGET_EXT", ".mp3")
	os.Setenv("HTTP_PORT", "9090")
	os.Setenv("READ_TIMEOUT", "200s")
	os.Setenv("WRITE_TIMEOUT", "200s")

	// Initialize the configuration
	ConfigInit()

	// Test App settings
	assert.Equal(t, "/tmp/files", AppSetting.FileSavePath)
	assert.Equal(t, 100, AppSetting.FileMaxSize)
	assert.Equal(t, []string{".mp4", ".mp3"}, AppSetting.FileAllowExts)
	assert.Equal(t, ".mp3", AppSetting.FileTargetExt)

	// Test Server settings
	assert.Equal(t, "9090", ServerSetting.HttpPort)
	assert.Equal(t, 200*time.Second, ServerSetting.ReadTimeout)
	assert.Equal(t, 200*time.Second, ServerSetting.WriteTimeout)

	// Clean up environment variables
	os.Unsetenv("FILE_SAVE_PATH")
	os.Unsetenv("FILE_MAX_SIZE")
	os.Unsetenv("FILE_ALLOW_EXTS")
	os.Unsetenv("FILE_TARGET_EXT")
	os.Unsetenv("HTTP_PORT")
	os.Unsetenv("READ_TIMEOUT")
	os.Unsetenv("WRITE_TIMEOUT")
}

func TestConfigInitDefaults(t *testing.T) {
	// Initialize the configuration with default values
	ConfigInit()

	// Test App settings
	assert.Equal(t, "./files/upload/", AppSetting.FileSavePath)
	assert.Equal(t, 50, AppSetting.FileMaxSize)
	assert.Equal(t, []string{".mp4a", ".mp3", ".wav"}, AppSetting.FileAllowExts)
	assert.Equal(t, ".wav", AppSetting.FileTargetExt)

	// Test Server settings
	assert.Equal(t, "8081", ServerSetting.HttpPort)
	assert.Equal(t, 300*time.Second, ServerSetting.ReadTimeout)
	assert.Equal(t, 300*time.Second, ServerSetting.WriteTimeout)
}
