package ffmpeg

import (
	osi "github.com/hopeio/lemon/utils/os"
	"log"
)

// doc: https://ffmpeg.org/ffmpeg-codecs.html
// https://ffmpeg.org/download.html

const CommonCmd = `ffmpeg -i "%s" `

func ffmpegCmd(cmd string) error {
	log.Println(cmd)
	_, err := osi.Cmd(cmd)
	if err != nil {
		log.Println(err)
		return err
	}
	//log.Println(res)
	return nil
}
