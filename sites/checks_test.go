package sites

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var stub404 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}))

var attempts int
var stub200 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if attempts > 0 {
		attempts = 0
		w.WriteHeader(http.StatusNotFound)
	} else {
		attempts++ // Clearly this is a terrible way to do this, but it works
		w.WriteHeader(http.StatusOK)
	}
}))

var stub500 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}))

var stubHTML = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	stuff := `
	<html>
		<body>
			<div id="user-not-found">not found</div>
			<div id="some-other-thing"></div>
		</body>
	</html>
	`
	w.Write([]byte(stuff))
}))

var renders int
var stubAltHTML = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var stuff string
	if renders > 0 {
		renders = 0
		stuff = `
		<html>
			<body>
				<div id="foobar">this username exists</div>
				<div>Some more test garbage</div>
				<div id="blahblahblah">this username does not exist</div>
			</body>
		</html>
		`
	} else {
		renders++ // I know, I know
		stuff = `
		<html>
			<body>
				<div id="user-found">this username exists</div>
				<div>Some other test garbage</div>
				<div id="user-not-found">this username does not exist</div>
			</body>
		</html>
		`
	}
	w.Write([]byte(stuff))
}))

var notRenders int
var stubNotAltHTML = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var stuff string
	if notRenders > 0 {
		notRenders = 0
		stuff = `
		<html>
			<body>
				<div id="foobar">this username exists</div>
				<div>Some more test garbage</div>
				<div id="blahblahblah">this username does not exist</div>
			</body>
		</html>
		`
	} else {
		notRenders++ // I know, I know
		stuff = `
		<html>
			<body>
				<div id="user-found">this username exists</div>
				<div>Some other test garbage</div>
				<div id="user-not-found">this username does not exist</div>
			</body>
		</html>
		`
	}
	w.Write([]byte(stuff))
}))

type MockSite struct {
	Site
	url       string
	userAgent string
}

func (m MockSite) Name() string {
	return "Mocked"
}

func (m MockSite) URL() string {
	return m.url
}

func (m MockSite) UserAgent() string {
	return m.userAgent
}

const TestUsername = "RippedMyPants"

// Should receive a 404 status and consider username available
func TestIfPageNotFound(t *testing.T) {
	noExist := MockSite{url: stub404.URL + "/?username=%s"}
	result := IfPageNotFound(noExist, TestUsername)
	assert.True(t, result.Available)
}

// Should receive a 200 and not consider username available
func TestIfPageFound(t *testing.T) {
	exists := MockSite{url: stub200.URL + "/?username=%s"}
	result := IfPageNotFound(exists, TestUsername)
	assert.False(t, result.Available)
	assert.Equal(t, "_"+TestUsername, result.Alternate)
}

func TestIfPageFoundWithUserAgent(t *testing.T) {
	exists := MockSite{url: stub200.URL + "/?username=%s", userAgent: "Googlebot/2.1 (+http://www.google.com/bot.html)"}
	result := IfPageNotFound(exists, TestUsername)
	assert.False(t, result.Available)
	assert.Equal(t, "_"+TestUsername, result.Alternate)
}

func TestBrokenIfPageFound(t *testing.T) {
	problem := MockSite{url: "jalsdkfjlaskdfj"}
	result := IfPageNotFound(problem, TestUsername)
	assert.True(t, result.Available)
	assert.Equal(t, "ERROR", result.Alternate)
}

func TestIfElementOnPage(t *testing.T) {
	userNotFound := MockSite{url: stubHTML.URL + "/?username=%s"}
	result := IfElementOnPage(userNotFound, TestUsername, "//div[@id='user-not-found']")
	assert.True(t, result.Available)

	badResult := IfElementOnPage(userNotFound, TestUsername, "//div[@id='you-will-never-find']")
	assert.False(t, badResult.Available)
	assert.Equal(t, "", badResult.Alternate)
}

func TestAlternateElementOnPage(t *testing.T) {
	userNotFound := MockSite{url: stubAltHTML.URL + "/?username=%s"}
	result := IfElementOnPage(userNotFound, TestUsername, "//div[@id='blahblahblah']")
	assert.False(t, result.Available)
	assert.Equal(t, "_"+TestUsername, result.Alternate)
}

func TestBrokenElementOnPage(t *testing.T) {
	problem := MockSite{url: "jalsdkfjlaskdfj"}
	result := IfElementOnPage(problem, TestUsername, "//div[1]")
	assert.True(t, result.Available)
	assert.Equal(t, "ERROR", result.Alternate)
}

func TestIfElementNotOnPage(t *testing.T) {
	notOnPage := MockSite{url: stubHTML.URL + "/username=%s"}
	result := IfElementNotOnPage(notOnPage, TestUsername, "//div[@id='user-exists']")
	assert.True(t, result.Available)
}

func TestIfElementNotOnPageButYesOnPageAnyway(t *testing.T) {
	yesOnPage := MockSite{url: stubHTML.URL + "/username=%s"}
	result := IfElementNotOnPage(yesOnPage, TestUsername, "//div[@id='user-not-found']")
	assert.False(t, result.Available)
}

func TestAlternateElementNotOnPage(t *testing.T) {
	userNotFound := MockSite{url: stubNotAltHTML.URL + "/username=%s"}
	result := IfElementNotOnPage(userNotFound, TestUsername, "//div[@id='user-found']")
	assert.False(t, result.Available)
	assert.Equal(t, "_"+TestUsername, result.Alternate)
}

func TestBrokenElementNotOnPage(t *testing.T) {
	problem := MockSite{url: "laksjdflkasjdl"}
	result := IfElementNotOnPage(problem, TestUsername, "//div[1]")
	assert.True(t, result.Available)
	assert.Equal(t, "ERROR", result.Alternate)
}
