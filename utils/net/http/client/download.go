package client

import (
	"errors"
	"fmt"
	fs2 "github.com/hopeio/lemon/utils/io/fs"
	"github.com/hopeio/lemon/utils/log"
	httpi "github.com/hopeio/lemon/utils/net/http"
	"io"
	"net/http"
	"time"
)

// TODO: Range StatusPartialContent 下载
type Download struct {
	Client  *http.Client
	Request *http.Request
	mode    uint8 // 模式，0-强制覆盖，1-不存在下载
}

func NewDownload(url string) (*Download, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 如果自己设置了接受编码，http库不会自动gzip解压，需要自己处理，不加Accept-Encoding和Range头会自动设置gzip
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set(httpi.HeaderAcceptLanguage, "zh-CN,zh;q=0.9;charset=utf-8")
	req.Header.Set(httpi.HeaderConnection, "keep-alive")
	req.Header.Set(httpi.HeaderUserAgent, UserAgent2)
	return &Download{
		Client:  defaultClient,
		Request: req,
	}, nil
}

func (d *Download) WithClient(c *http.Client) *Download {
	d.Client = c
	return d
}

func (d *Download) SetClient(set func(*http.Client)) *Download {
	set(d.Client)
	return d
}

func (d *Download) WithRequest(c *http.Request) *Download {
	d.Request = c
	return d
}

func (d *Download) SetRequest(set RequestOption) *Download {
	set(d.Request)
	return d
}

func (d *Download) WithOptions(opts ...RequestOption) *Download {
	for _, opt := range opts {
		opt(d.Request)
	}
	return d
}

func (d *Download) SetHeader(header Header) *Download {
	return d.WithOptions(SetHeader(header))
}

func (d *Download) AddHeader(header, value string) *Download {
	return d.WithOptions(SetHeader(Header{header, value}))
}

// 保留模式，如果文件已存在，不下载覆盖
func (d *Download) RetainMode() *Download {
	d.mode = 1
	return d
}

func (d *Download) GetReader() (io.ReadCloser, error) {
	if d.Client == nil || d.Request == nil {
		return nil, errors.New("client 或 request 为 nil")
	}
	var resp *http.Response
	var err error
	for i := 0; i < 3; i++ {
		if i > 0 {
			time.Sleep(time.Second)
		}
		resp, err = d.Client.Do(d.Request)
		if err != nil {
			log.Warn(err, "url:", d.Request.URL.Path)
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			resp.Body.Close()
			if resp.StatusCode == http.StatusNotFound {
				return nil, ErrNotFound
			}
			return nil, fmt.Errorf("返回错误,状态码:%d,url:%s", resp.StatusCode, d.Request.URL.Path)
		} else {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *Download) DownloadFile(filepath string) error {
	if d.mode == 1 && fs2.Exist(filepath) {
		return nil
	}
	reader, err := d.GetReader()
	if err != nil {
		return err
	}
	err = fs2.CreatFileFromReader(filepath, reader)
	err1 := reader.Close()
	if err1 != nil {
		log.Warn("Close Reader", err1)
	}
	return err
}

func GetFile(url string) (io.ReadCloser, error) {
	return GetFileWithReqOption(url, nil)
}

func GetFileWithReqOption(url string, opts ...RequestOption) (io.ReadCloser, error) {
	d, err := NewDownload(url)
	if err != nil {
		return nil, err
	}
	return d.WithOptions(opts...).GetReader()
}

func DownloadFile(filepath, url string) error {
	d, err := NewDownload(url)
	if err != nil {
		return err
	}
	return d.DownloadFile(filepath)
}

func GetImage(url string) (io.ReadCloser, error) {
	return GetFileWithReqOption(url, ImageOption)
}

func DownloadImage(filepath, url string) error {
	reader, err := GetFileWithReqOption(url, ImageOption)
	if err != nil {
		return err
	}
	return fs2.CreatFileFromReader(filepath, reader)
}

func ImageOption(req *http.Request) {
	req.Header.Set(httpi.HeaderAccept, "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
}
