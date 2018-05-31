package sites

import "fmt"
import "net/http"
import "github.com/antchfx/htmlquery"

// Site is a site that can be queried for a username
type Site interface {
	Name() string
	URL() string
	UserAgent() string
	Check(username string, ch chan *NameResult)
}

// NameExists makes a simple check for an already existing username.
func NameExists(urlTemplate, username string, userAgent string, attempt int) (string, error) {
	url := fmt.Sprintf(urlTemplate, username)
	var resp *http.Response
	var err error

	if userAgent != "" {
		client := &http.Client{}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("User-Agent", userAgent)
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
	}

	return NameExists(urlTemplate, newUsername, userAgent, attempt+1)
}

// NotInPage decides a username is available if the given element is not found
func NotInPage(urlTemplate, username, xpath string, attempt int) (string, error) {
	url := fmt.Sprintf(urlTemplate, username)
	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		return "", err
	}

	missing := true
	for _, n := range htmlquery.Find(doc, xpath) {
		if n.Data != "" {
			missing = false
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
	}

	return NotInPage(urlTemplate, newUsername, xpath, attempt+1)
}

// InPage decides a username is available if the given element is found
func InPage(urlTemplate, username, xpath string, attempt int) (string, error) {
	url := fmt.Sprintf(urlTemplate, username)
	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		return "", err
	}

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
	}

	return InPage(urlTemplate, newUsername, xpath, attempt+1)
}
