package template_method

import "testing"

func TestExampleHttpDownloader(t *testing.T)  {
	var downloader  = NewHttpDownloader()
	downloader.Download("http://baidu.com/abc.zip")
}

func TestExampleFtpDownloader(t *testing.T)  {
	var downloader  = NewFtpDownloader()
	downloader.Download("ftp://baidu.com/abc.zip")
}