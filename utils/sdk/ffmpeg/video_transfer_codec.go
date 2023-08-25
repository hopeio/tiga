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

const ToAv1Cmd = CommonCmd + "-c:v libaom-av1 -crf 30 -row-mt 1 -tiles 2x2 -y %s"

func ToAV1ByLibaomav1(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(ToAv1Cmd, filePath, dst))
}

const ToH265Cmd = CommonCmd + "-c:v libx265 -preset veryslow -crf 28 -y %s"

func ToH265ByXlib265(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(ToH265Cmd, filePath, dst))
}
