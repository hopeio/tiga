package client

import httpi "github.com/hopeio/tiga/utils/net/http"

type Header []string

func NewHeader() *Header {
	h := make(Header, 0, 6)
	return &h
}

func (h *Header) Add(k, v string) *Header {
	*h = append(*h, k, v)
	return h
}

func (h *Header) Set(k, v string) *Header {
	header := *h
	for i, s := range header {
		if s == k {
			header[i+1] = v
			return h
		}
	}
	return h.Add(k, v)
}

func (h Header) Clone() Header {
	newh := make(Header, len(h))
	copy(newh, h)
	return newh
}

func DefaultHeader() Header {
	return Header{
		httpi.HeaderAcceptLanguage, "zh-CN,zh;q=0.9;charset=utf-8",
		httpi.HeaderConnection, "keep-alive",
		httpi.HeaderUserAgent, UserAgentChrome117,
		//"Accept", "application/json,text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8", // 将会越来越少用，服务端一般固定格式
	}
}
