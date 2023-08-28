package ffmpeg

import (
	"fmt"
	osi "github.com/hopeio/lemon/utils/os"
	"github.com/hopeio/lemon/utils/sdk/mp4box"
	"strings"
)

const ImgToWebpCmd = CommonCmd + `-c:v libwebp -lossless 1 -quality 100 -compression_level 6 %s.webp`
const GifToWebpCmd = CommonCmd + `-c:v libwebp -lossless 1 -quality 100 -compression_level 6 %s.webp`

// 图片转webp格式
func ImgToWebp(filePath, dst string) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	if strings.HasSuffix(dst, ".gif") {
		return ffmpegCmd(fmt.Sprintf(GifToWebpCmd, filePath, dst))
	}
	return ffmpegCmd(fmt.Sprintf(ImgToWebpCmd, filePath, dst))
}

const ImgToWebpWithOptionsCmd = CommonCmd + `-c:v libwebp -quality %d -method 4 %s.webp`

// 图片带选项转webp格式,选项目前支持质量(0-100),推荐75
func ImgToWebpWithOptions(filePath, dst string, quality int) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	if strings.HasSuffix(dst, ".gif") {
		// TODO
	}
	return ffmpegCmd(fmt.Sprintf(ImgToWebpWithOptionsCmd, filePath, quality, dst))
}

const ImgTAvifCmd = CommonCmd + `-c:v libaom-av1 -still-picture 1 %s.avif`

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

const ImgToJxlCmd = CommonCmd + `-c:v libjxl %s.jxl`

// 不可用,没有注明色彩空间的原因。需要显式写明 像素编码格式、色彩空间、转换色彩空间、目标色彩空间、色彩范围
func ImgToJxl(filePath, dst string) error {
	if strings.HasSuffix(dst, ".heif") {
		dst = dst[:len(dst)-5]
	}
	_, err := osi.ContainQuotedCMD(fmt.Sprintf(ImgToHeifCmd, filePath, dst))
	if err != nil {
		return err
	}

	return mp4box.Heif(dst+".hevc", dst)
}

// TODO
func ImgToWebp2(filePath, dst string) error {
	return nil
}
