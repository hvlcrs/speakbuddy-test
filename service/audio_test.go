package service

import (
	"errors"
	"os"
	"testing"

	"speakbuddy/pkg/configs"

	"github.com/stretchr/testify/assert"
)

func TestValidateAudioFormat(t *testing.T) {
	configs.AppSetting.FileAllowExts = []string{".mp3", ".wav"}

	tests := []struct {
		format string
		err    error
	}{
		{".mp3", nil},
		{".wav", nil},
		{".flac", errors.New("unsupported file format")},
	}

	for _, test := range tests {
		err := ValidateAudioFormat(test.format)
		assert.Equal(t, test.err, err)
	}
}

func TestCleanupLocalCache(t *testing.T) {
	path := "testfile.txt"
	file, _ := os.Create(path)
	file.Close()

	err := CleanupLocalCache(path)
	assert.Nil(t, err)

	_, err = os.Stat(path)
	assert.True(t, os.IsNotExist(err))
}
