// Package httpclient is a simple http client for personal use, derived from [this](https://github.com/asmcos/requests)
package httpclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

const (
	VERSION = "0.1"
)

type Request struct {
	req     *http.Request
	Header  *http.Header
	Client  *http.Client
	Debug   bool
	Cookies []*http.Cookie
}

type Response struct {
	Resp    *http.Response
	content []byte
	text    string
	req     *Request
}

type Header map[string]string
type Params map[string]string
type Auth []string
type Data map[string]string
type Files map[string]string

func BuildRequst() *Request {
	req := new(Request)
	req.req = &http.Request{
		Header: make(http.Header),
	}
	req.Debug = true
	req.Header = &req.req.Header
	req.req.Header.Set("User-Agent", "Go Requests "+VERSION)
	req.Client = &http.Client{}
	jar, _ := cookiejar.New(nil)
	req.Client.Jar = jar
	return req
}

func (req *Request) Get(originURL string, args ...interface{}) (resp *Response, err error) {
	req.req.Method = "GET"
	params := []map[string]string{}
	delete(req.req.Header, "Cookie")
	for _, arg := range args {
		switch arg.(type) {
		case Header:
			for k, v := range arg.(Header) {
				req.Header.Set(k, v)
			}
		case Params:
			params = append(params, arg.(Params))
		case Auth:
			auth := arg.(Auth)
			req.req.SetBasicAuth(auth[0], auth[1])
		}
	}
	log.Println(params)
	distURL, _ := buildURLParams(originURL, params...)
	URL, err := url.Parse(distURL)
	if err != nil {
		return nil, err
	}
	req.req.URL = URL
	req.ClientSetCookies()

	req.RequestDebug()

	res, err := req.Client.Do(req.req)
	if err != nil {
		return nil, err
	}
	resp = &Response{}
	resp.Resp = res
	resp.req = req

	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

func (req *Request) Post(origurl string, args ...interface{}) (resp *Response, err error) {
	req.req.Method = "POST"

	//set default
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set params ?a=b&b=c
	//set Header
	params := []map[string]string{}
	datas := []map[string]string{} // POST
	files := []map[string]string{} //post file

	//reset Cookies,
	//Client.Do can copy cookie from client.Jar to req.Header
	delete(req.req.Header, "Cookie")

	for _, arg := range args {
		switch a := arg.(type) {
		// arg is Header , set to request header
		case Header:

			for k, v := range a {
				req.Header.Set(k, v)
			}
			// arg is "GET" params
			// ?title=website&id=1860&from=login
		case Params:
			params = append(params, a)

		case Data: //Post form data,packaged in body.
			datas = append(datas, a)
		case Files:
			files = append(files, a)
		case Auth:
			// a{username,password}
			req.req.SetBasicAuth(a[0], a[1])
		}
	}

	disturl, _ := buildURLParams(origurl, params...)

	if len(files) > 0 {
		req.buildFilesAndForms(files, datas)

	} else {
		Forms := req.buildForms(datas...)
		req.setBodyBytes(Forms) // set forms to body
	}
	//prepare to Do
	URL, err := url.Parse(disturl)
	if err != nil {
		return nil, err
	}
	req.req.URL = URL

	req.ClientSetCookies()

	req.RequestDebug()

	res, err := req.Client.Do(req.req)

	// clear post param
	req.req.Body = nil
	req.req.GetBody = nil
	req.req.ContentLength = 0

	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp = &Response{}
	resp.Resp = res
	resp.req = req

	resp.Content()
	defer res.Body.Close()

	resp.ResponseDebug()
	return resp, nil
}

func (req *Request) ClientSetCookies() {
	if len(req.Cookies) > 0 {
		req.Client.Jar.SetCookies(req.req.URL, req.Cookies)
		req.ClearCookies()
	}
}

func (req *Request) ClearCookies() {
	req.Cookies = req.Cookies[:0]
}

func (resp *Response) Content() ([]byte, error) {
	if len(resp.content) > 1 {
		return resp.content, nil
	}
	body := resp.Resp.Body
	if resp.Resp.Header.Get("Content-Encoding") == "gzip" && resp.Resp.Header.Get("Accept-Encoding") != "" {
		reader, err := gzip.NewReader(body)
		if err != nil {
			return nil, err
		}
		body = reader
	}
	var err error
	resp.content, err = ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return resp.content, nil
}

func (resp *Response) Text() string {
	if resp.content == nil {
		resp.Content()
	}
	resp.text = string(resp.content)
	return resp.text
}

func (resp *Response) Json(v interface{}) error {
	if resp.content == nil {
		resp.Content()
	}
	return json.Unmarshal(resp.content, v)
}

func (req *Request) RequestDebug() {
	if !req.Debug {
		return
	}
	log.Println("=========== Debug ============")
	message, err := httputil.DumpRequestOut(req.req, false)
	if err != nil {
		return
	}
	log.Println(string(message))

	if len(req.Client.Jar.Cookies(req.req.URL)) > 0 {
		log.Println("Cookies: ")
		for _, cookie := range req.Client.Jar.Cookies(req.req.URL) {
			log.Println(cookie)
		}
	}
}
func (resp *Response) ResponseDebug() {

	if !resp.req.Debug {
		return
	}

	log.Println("=========== Debug ============")

	message, err := httputil.DumpResponse(resp.Resp, false)
	if err != nil {
		return
	}

	log.Println(string(message))

}

func buildURLParams(baseURL string, params ...map[string]string) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)

	if err != nil {
		return "", err
	}
	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}
	return addQueryParams(parsedURL, parsedQuery), nil
}

func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	if len(parsedQuery) > 0 {
		return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
	}
	return strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)
}

// upload file and form
// build to body format
func (req *Request) buildFilesAndForms(files []map[string]string, datas []map[string]string) {

	//handle file multipart

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for _, file := range files {
		for k, v := range file {
			part, err := w.CreateFormFile(k, v)
			if err != nil {
				log.Printf("Upload %s failed!", v)
				panic(err)
			}
			file := openFile(v)
			_, err = io.Copy(part, file)
			if err != nil {
				panic(err)
			}
		}
	}

	for _, data := range datas {
		for k, v := range data {
			w.WriteField(k, v)
		}
	}

	w.Close()
	// set file header example:
	// "Content-Type": "multipart/form-data; boundary=------------------------7d87eceb5520850c",
	req.req.Body = ioutil.NopCloser(bytes.NewReader(b.Bytes()))
	req.req.ContentLength = int64(b.Len())
	req.Header.Set("Content-Type", w.FormDataContentType())
}

// build post Form data
func (req *Request) buildForms(datas ...map[string]string) (Forms url.Values) {
	Forms = url.Values{}
	for _, data := range datas {
		for key, value := range data {
			Forms.Add(key, value)
		}
	}
	return Forms
}

// only set forms
func (req *Request) setBodyBytes(Forms url.Values) {

	// maybe
	data := Forms.Encode()
	req.req.Body = ioutil.NopCloser(strings.NewReader(data))
	req.req.ContentLength = int64(len(data))
}

// only set forms
func (req *Request) setBodyRawBytes(read io.ReadCloser) {
	req.req.Body = read
}

// openFile for post upload files
func openFile(filename string) *os.File {
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return r
}

func Get(url string, args ...interface{}) (resp *Response, err error) {
	req := BuildRequst()
	resp, err = req.Get(url, args...)
	return resp, err
}
