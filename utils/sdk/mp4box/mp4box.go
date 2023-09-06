package mp4box

import osi "github.com/hopeio/lemon/utils/os"

// https://www.videohelp.com/software/MP4Box
const Mp4BoxCmd = `MP4Box -add-image (%s.hevc:primary) -ab heic -new %s.heic`

func Heic(filePath, dst string) error {
	_, err := osi.Cmd(Mp4BoxCmd)
	return err
}
