package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	scheme               = "https"
	LitResHost           = "litres.ru"
	wwwHost              = "www." + LitResHost
	partnersDownloadHost = "partnersdnld." + LitResHost

	coverPath        = "/pub/c/cover"
	myBooksFreshPath = "/pages/my_books_fresh"
	detailedDataPath = "/static/ds/" + DetailedDataFilename
	biblioBookPath   = "/pages/biblio_book/"

	coverPathTemplate = "/{id}.jpg"
	pagePathTemplate  = "/page-{num}"

	DetailedDataFilename = "detailed_data.xml.gz"
)

func MyBooksFreshUrl(page int) *url.URL {

	pp := strings.Replace(pagePathTemplate, "{num}", strconv.Itoa(page), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   myBooksFreshPath + pp,
	}
}

func HrefUrl(href string) *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   href,
	}
}

func CoverUrl(id int64) *url.URL {

	cp := strings.Replace(coverPathTemplate, "{id}", strconv.FormatInt(id, 10), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   partnersDownloadHost,
		Path:   coverPath + cp,
	}
}

//func DetailedDataUrl() *url.URL {
//	return &url.URL{
//		Scheme: scheme,
//		Host:   wwwHost,
//		Path:   detailedDataPath,
//	}
//}

//func BiblioBookUrl(id int64) *url.URL {
//	u := &url.URL{
//		Scheme: scheme,
//		Host:   wwwHost,
//		Path:   biblioBookPath,
//	}
//
//	q := u.Query()
//	q.Set("art", strconv.FormatInt(id, 10))
//
//	u.RawQuery = q.Encode()
//
//	return u
//}
