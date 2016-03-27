package checker

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/asciimoo/privacyscore/result"
	"github.com/asciimoo/privacyscore/utils"
)

const RESOURCE_LIMIT = 64
const TIMEOUT = 5
const USER_AGENT = "Mozilla/5.0 (compatible) PrivacyScore Checker v0.1.0"
const maxResponseBodySize = 1024 * 1024 * 5

var mutex = &sync.Mutex{}

type Checker interface {
	Check(*result.Result, *PageInfo)
}

type CheckJob struct {
	sync.RWMutex
	Result    *result.Result
	Resources map[string]*PageInfo
	Chan      chan bool
}

type PageInfo struct {
	ResponseBody   []byte
	ContentType    string
	StatusCode     int
	URL            *url.URL
	OriginalURL    *url.URL
	Cookies        []*http.Cookie
	Domain         string
	ResponseHeader *http.Header
}

var checkers []Checker = []Checker{
	&CookieChecker{},
	&HTMLChecker{},
	&HTTPSChecker{},
	&SecureHeaderChecker{},
	&CSSChecker{},
}

func Run(URL string) (*result.Result, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	c := newCheckJob(URL)
	finishedResources := 0
	c.CheckURL(URL)
	for finishedResources != len(c.Resources) && finishedResources < RESOURCE_LIMIT {
		select {
		case _ = <-c.Chan:
			finishedResources += 1
		}
	}
	if finishedResources == 0 {
		return c.Result, errors.New("Could not download host")
	}
	return c.Result, nil
}

func newCheckJob(URL string) *CheckJob {
	return &CheckJob{
		Result:    result.New(URL),
		Resources: make(map[string]*PageInfo),
		Chan:      make(chan bool, RESOURCE_LIMIT),
	}
}

func (c *CheckJob) CheckURL(URL string) {
	// URL already added
	if _, found := c.Resources[URL]; found {
		c.Chan <- true
		return
	}
	r, err := fetchURL(URL)
	if err != nil {
		c.Result.AddError(err)
		c.Chan <- false
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	u, _ := url.Parse(URL)
	p := &PageInfo{
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		r.Request.URL,
		u,
		r.Cookies(),
		utils.CropSubdomains(r.Request.URL.Host),
		&r.Header,
	}
	if err != nil {
		c.Result.AddError(err)
	}
	c.Lock()
	c.Resources[URL] = p
	c.Unlock()
	go func() {
		for _, ch := range checkers {
			ch.Check(c.Result, p)
		}
		c.Chan <- false
	}()
	return
}

func fetchURL(URL string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * time.Second),
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	response, err := client.Do(req)
	return response, err
}
