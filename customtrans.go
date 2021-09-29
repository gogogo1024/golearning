// 在默认的http.Transport之上包了一层Transport并实现了RoundTrip()方法
// 做法有点类似于egg框架对context的包装
package main

import (
	"net/http"
)
type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper{
	if t.Transport == nil{
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request)(*http.Response, error){
	// do something
	// curl
	// process req.Header
	return t.Transport.RoundTrip(req)
}

// client可以理解为业务层，而不关心具体的Transport是如何实现的
// 类比于client是应用层，custom是传输层（在应用层的处理业务的基础上面补充了比如
// HTTP代理，gzip压缩，链接池管理，认证）
func (t *OurCustomTransport) Client() *http.Client{
	return &http.Client{Transport: t}
}
func main() {
	t :=&OurCustomTransport{}
	c :=t.Client()
	resp, err := c.Get("http://example.com")

}  