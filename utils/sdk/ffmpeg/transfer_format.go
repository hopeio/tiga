package ffmpeg

import (
	"fmt"
	osi "github.com/hopeio/lemon/utils/os"
	"log"
)

const TransferFormatGPUCmd = ` -hwaccel qsv -i "%s" -c copy -y "%s"`

func TransferFormatGPU(filePath, dst string) error {
	command := fmt.Sprintf(execPath+TransferFormatGPUCmd, filePath, dst)
	log.Println(command)
	_, err := osi.Cmd(command)
	return err
}

const TransferFormatCmd = CommonCmd + ` -c copy -y "%s"`

func TransferFormat(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(TransferFormatCmd, filePath, dst))
}

const ConcatCmd = ` -f concat -safe 0  -i "%s" -c copy -y "%s"`

func Concat(filePath, dst string) error {
	return ffmpegCmd(fmt.Sprintf(ConcatCmd, filePath, dst))
}
