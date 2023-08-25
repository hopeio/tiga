package libjxl

import (
	"fmt"
	osi "github.com/hopeio/lemon/utils/os"
	"strings"
)

// https://github.com/libjxl/libjxl/releases
// windows support: https://github.com/saschanaz/jxl-winthumb/releases administrator regsvr32 jxl_winthumb.dll
const ImgToJxlCmd = `cjxl %s %s.jxl`
const JxlImgToOtherCmd = `djxl %S %s`

func ImgToJxl(filePath, dst string) error {
	if strings.HasSuffix(dst, ".jxl") {
		dst = dst[:len(dst)-4]
	}
	_, err := osi.Cmd(fmt.Sprintf(ImgToJxlCmd, filePath, dst))
	return err
}
