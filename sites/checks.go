package sites

// NameResult represents the result of a username query
type NameResult struct {
	SiteName  string
	Available bool
	Alternate string
}

// IfPageNotFound tests a username is available if a request returns a 404
func IfPageNotFound(siteName, username string, fakeUA bool) *NameResult {

	name := NameResult{SiteName: siteName, Available: false, Alternate: ""}

	result, err := NameExists(siteName, username, fakeUA, 1)
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
func IfElementOnPage(siteName, username, xpath string) *NameResult {

	name := NameResult{SiteName: siteName, Available: false, Alternate: ""}

	result, err := InPage(siteName, username, xpath, 1)
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
func IfElementNotOnPage(siteName, username, xpath string) *NameResult {

	name := NameResult{SiteName: siteName, Available: false, Alternate: ""}

	result, err := NotInPage(siteName, username, xpath, 1)
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
