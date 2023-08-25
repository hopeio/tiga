package libheif

import (
	"fmt"
	osi "github.com/hopeio/lemon/utils/os"
	"github.com/hopeio/lemon/utils/sdk/mp4box"
	"strings"
)

// https://github.com/pphh77/libheif-Windowsbinary/releases

const ImgToHeifCmd = `heif-enc -q 90 -p x265:colorprim=smpte170m %s -o %s.heic`
const ImgToHeifCmd1 = `heif-enc -p x265:crf=20.5 -p x265:colorprim=smpte170m -p x265:rdoq-level=1 -p x265:aq-strength=1.2 -p x265:deblock=-2:-2 %s -o %s.heic
`

func ImgToHeif(filePath, dst string) error {
	if strings.HasSuffix(dst, ".heif") {
		dst = dst[:len(dst)-5]
	}
	_, err := osi.ContainQuotedCMD(fmt.Sprintf(ImgToHeifCmd, filePath, dst))
	if err != nil {
		return err
	}

	return mp4box.Heif(dst+".hevc", dst)
}
