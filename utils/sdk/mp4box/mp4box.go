package mp4box

import osi "github.com/hopeio/lemon/utils/os"

// https://www.videohelp.com/software/MP4Box
const Mp4BoxCmd = `MP4Box -add-image (%s.hevc:primary) -ab heic -new %s.heif`

func Heif(filePath, dst string) error {
	_, err := osi.Cmd(Mp4BoxCmd)
	return err
}
