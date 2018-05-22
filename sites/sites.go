package sites

import "fmt"
import "net/http"
import "github.com/antchfx/htmlquery"

// URLS is a global map of sites and what URL to check
var URLS = map[string]string{
	"GitHub":    "https://github.com/%s",
	"Twitter":   "https://twitter.com/%s",
	"Facebook":  "https://www.facebook.com/%s",
	"Instagram": "https://www.instagram.com/%s/",
	"Fortnite":  "https://www.stormshield.one/pvp/stats/%s",
	"Twitch":    "https://api.twitch.tv/kraken/users/%s?client_id=bwllomi5jnmezyug7aprtkjqeb42hh",
	"Google+":   "https://plus.google.com/+%s",
}

// NameExists makes a simple check for an already existing username.
func NameExists(service, username string, fakeUserAgent bool, attempt int) (string, error) {
	url := fmt.Sprintf(URLS[service], username)
	var resp *http.Response
	var err error

	if fakeUserAgent == true {
		client := &http.Client{}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("User-Agent", "User-Agent: Mozilla/5.0 (Linux; Android 4.4; Nexus 5 Build/_BuildID_) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36")
		resp, err = client.Do(req)

		// Check for error
		if err != nil {
			return "", err
		}
	} else {
		resp, err = http.Get(url)

		// Check for error
		if err != nil {
			return "", err
		}
	}

	if resp.StatusCode == 404 { // Presumably, the username does not exist
		return username, nil
	}

	// Regular username might not be available, try something similar
	newUsername := fmt.Sprintf("_%s", username)
	if attempt > 3 { // Quit
		return "", nil
	} else {
		return NameExists(service, newUsername, fakeUserAgent, attempt+1)
	}
}

func NotInPage(service, username, xpath string, attempt int) (string, error) {
	url := fmt.Sprintf(URLS[service], username)
	doc, _ := htmlquery.LoadURL(url)

	missing := false
	for _, n := range htmlquery.Find(doc, xpath) {
		if n.Data == "" {
			missing = true
			break
		}
	}
	if missing == true {
		return username, nil
	}

	// Regular username might not be available, try something similar
	newUsername := fmt.Sprintf("_%s", username)
	if attempt > 3 { // Quit
		return "", nil
	} else {
		return NotInPage(service, newUsername, xpath, attempt+1)
	}
}

func InPage(service, username, xpath string, attempt int) (string, error) {
	url := fmt.Sprintf(URLS[service], username)
	doc, _ := htmlquery.LoadURL(url)

	found := false
	for _, n := range htmlquery.Find(doc, xpath) {
		if n.Data != "" {
			found = true
			break
		}
	}
	if found == true {
		return username, nil
	}

	// Regular username might not be available, try something similar
	newUsername := fmt.Sprintf("_%s", username)
	if attempt > 3 { // Quit
		return "", nil
	} else {
		return InPage(service, newUsername, xpath, attempt+1)
	}
}
