package ffmpeg

import (
	"fmt"
	osi "github.com/hopeio/lemon/utils/os"
	"github.com/hopeio/lemon/utils/sdk/mp4box"
	"strings"
)

const ImgToWebpCmd = CommonCmd + `-c:v libwebp -lossless 1 -quality 100 -compression_level 6 %s.webp`

// 图片转webp格式
func ImgToWebp(filePath, dst string) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	return ffmpegCmd(fmt.Sprintf(ImgToWebpCmd, filePath, dst))
}

const ImgToWebpWithOptionsCmd = CommonCmd + `-c:v libwebp -quality %d -method 4 %s.webp`

// 图片带选项转webp格式,选项目前支持质量(0-100),推荐75
func ImgToWebpWithOptions(filePath, dst string, quality int) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	return ffmpegCmd(fmt.Sprintf(ImgToWebpWithOptionsCmd, filePath, quality, dst))
}

const ImgTAvifCmd = CommonCmd + `-c:v libaom-av1 -still-picture 1 %s.avif`

// More encoding options are available: -b 700k -tile-columns 600 -tile-rows 800 - example for the bitrate and tales.
func ImgToAvif(filePath, dst string) error {
	if strings.HasSuffix(dst, ".avif") {
		dst = dst[:len(dst)-5]
	}
	return ffmpegCmd(fmt.Sprintf(ImgTAvifCmd, filePath, dst))
}

const ImgToHeifCmd = CommonCmd + `-crf 12 -psy-rd 0.4 -aq-strength 0.4 -deblock 1:1 -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -preset veryslow -pix_fmt yuv420p101e -f hevc %s.hevc`
const ImgToHeifCmd2 = `ffmpeg -hide_banner -r 1 -i %s -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2,zscale=m=170m:r=pc" -pix_fmt yuv420p -frames 1 -c:v libx265 -preset veryslow -crf 20 -x265-params range=full:colorprim=smpte170m %s.hevc`
const ImgToHeifCmd3 = `ffmpeg -hide_banner -r 1 -i %s -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2,zscale=m=170m:r=pc" -pix_fmt yuv420p -frames 1 -c:v libx265 -preset veryslow -crf 20 -x265-params range=full:colorprim=smpte170m:aq-strength=1.2 -deblock -2:-2 %s.hevc
`

func ImgToHeic(filePath, dst string) error {
	if strings.HasSuffix(dst, ".heic") {
		dst = dst[:len(dst)-5]
	}
	_, err := osi.ContainQuotedCMD(fmt.Sprintf(ImgToHeifCmd, filePath, dst))
	if err != nil {
		return err
	}

	return mp4box.Heif(dst+".hevc", dst)
}

const ImgToJxlCmd = CommonCmd + `-c:v libjxl %s.jxl`

// 不可用,没有注明色彩空间的原因。需要显式写明 像素编码格式、色彩空间、转换色彩空间、目标色彩空间、色彩范围
// distance: Butteraugli distance, lower is better, 0.0 - lossless, 15.0 - minimum quality.
// effort: higher is better, 7 is the best quality, 1 - the worst.
func ImgToJxl(filePath, dst string) error {
	if strings.HasSuffix(dst, ".jxl") {
		dst = dst[:len(dst)-4]
	}

	return ffmpegCmd(fmt.Sprintf(ImgToJxlCmd, filePath, dst))
}

// TODO
func ImgToWebp2(filePath, dst string) error {
	return nil
}
