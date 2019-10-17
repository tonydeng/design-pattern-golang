package template_method

import "fmt"

type Downloader interface {
	Download(url string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) Save() {
	fmt.Print("default save\n")
}

type HttpDownloader struct {
	*template
}

func NewHttpDownloader() Downloader {
	downloader := &HttpDownloader{}
	template := newTemplate(downloader)

	downloader.template = template
	return downloader
}

func (d *HttpDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HttpDownloader) save() {
	fmt.Print("http save\n")
}

type FtpDownloader struct {
	*template
}

func NewFtpDownloader() Downloader {
	downloader := &FtpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template

	return downloader
}

func (d *FtpDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

func (*FtpDownloader) save() {
	fmt.Print("ftp save\n")
}