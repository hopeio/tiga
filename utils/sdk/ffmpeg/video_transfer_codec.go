package ffmpeg

import (
	"fmt"
)

const param = "-global_quality 20"

const H264ToH265ByIntelGPUCmd = `ffmpeg -hwaccel_output_format qsv -c:v h264_qsv -i %s -c:v hevc_qsv -preset veryslow -g 60 -gpu_copy 1 -c:a copy %s`

const cmd1 = `preset=veryslow,profile=main,look_ahead=1,global_quality=18`

func H264ToH265ByIntelGPU(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(H264ToH265ByIntelGPUCmd, filePath, dst))
}

// libaom-av1
const ToAv1Libaomav1Cmd = CommonCmd + "-c:v libaom-av1 -crf 30 -cpu-used 4 -y %s"

// cpu-used
// Set the quality/encoding speed tradeoff. Valid range is from 0 to 8, higher numbers indicating greater speed and lower quality. The default value is 1, which will be slow and high quality.
// 很慢,cpu-used调高质量差
func ToAV1ByLibaomav1(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(ToAv1Libaomav1Cmd, filePath, dst))
}

// libsvtav1
// librav1e

// libx265
const ToH265Cmd = CommonCmd + "-c:v libx265 -preset veryslow -crf 28 -y %s"

func ToH265ByXlib265(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(ToH265Cmd, filePath, dst))
}

// libvpx
