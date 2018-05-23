package sites

// NameResult represents the result of a username query
type NameResult struct {
	SiteName  string
	Available bool
	Alternate string
}

// IfPageNotFound tests a username is available if a request returns a 404
func IfPageNotFound(site Site, username string) *NameResult {

	name := NameResult{SiteName: site.Name(), Available: false, Alternate: ""}

	result, err := NameExists(site.URL(), username, site.UserAgent(), 1)
	if err != nil {
		name.Available = true
		name.Alternate = "ERROR"
		return &name
	}
	if result == username {
		name.Available = true
	} else if result != "" {
		name.Alternate = result
	}

	return &name
}

// IfElementOnPage tests a username is available if an element exists on a page
func IfElementOnPage(site Site, username, xpath string) *NameResult {

	name := NameResult{SiteName: site.Name(), Available: false, Alternate: ""}

	result, err := InPage(site.URL(), username, xpath, 1)
	if err != nil {
		name.Available = true
		name.Alternate = "ERROR"
		return &name
	}

	if result == username {
		name.Available = true
	} else if result != "" {
		name.Alternate = result
	}

	return &name
}

// IfElementNotOnPage tests a username is available if an element does not exist on a page
func IfElementNotOnPage(site Site, username, xpath string) *NameResult {

	name := NameResult{SiteName: site.Name(), Available: false, Alternate: ""}

	result, err := NotInPage(site.URL(), username, xpath, 1)
	if err != nil {
		name.Available = true
		name.Alternate = "ERROR"
		return &name
	}

	if result == username {
		name.Available = true
	} else if result != "" {
		name.Alternate = result
	}

	return &name
}
