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

const ImgToWebpWithOptionsCmd = CommonCmd + `-c:v libwebp -quality %d %s.webp`

// 图片带选项转webp格式,选项目前支持质量(0-100),ffmpeg默认75,这里默认90
func ImgToWebpWithOptions(filePath, dst string, quality int) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	if quality == 0 {
		quality = 90
	}
	return ffmpegCmd(fmt.Sprintf(ImgToWebpWithOptionsCmd, filePath, quality, dst))
}

const ImgTAvifCmd = CommonCmd + `-c:v libaom-av1 -crf 28 %s.avif`

// -cpu-used 4 -threads 8 会加速，但是图片大小会变大,质量变差

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
/*
distance
Set the target Butteraugli distance. This is a quality setting: lower distance yields higher quality, with distance=1.0 roughly comparable to libjpeg Quality 90 for photographic content. Setting distance=0.0 yields true lossless encoding. Valid values range between 0.0 and 15.0, and sane values rarely exceed 5.0. Setting distance=0.1 usually attains transparency for most input. The default is 1.0.

effort
Set the encoding effort used. Higher effort values produce more consistent quality and usually produces a better quality/bpp curve, at the cost of more CPU time required. Valid values range from 1 to 9, and the default is 7.

modular
Force the encoder to use Modular mode instead of choosing automatically. The default is to use VarDCT for lossy encoding and Modular for lossless. VarDCT is generally superior to Modular for lossy encoding but does not support lossless encoding.
*/
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
