package ffmpeg

import (
	"Douyin/config"
	"os/exec"
)

func Do(videoPath string, outputPath string) error {
	// 使用ffmpeg命令行工具截取一帧
	path := config.Conf.Ffmpeg
	cmd := exec.Command(path, "-i", videoPath, "-ss", "00:00:01.000", "-vframes", "1", outputPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
