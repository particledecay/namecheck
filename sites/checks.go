package sites

func IfPageNotFound(siteName, username string, fakeUA bool) (string, bool, string) {

	available := false
	alternative := ""

	result, err := NameExists(siteName, username, fakeUA, 1)
	if err != nil {
		return siteName, true, "ERROR"
	}
	if result == username {
		available = true
	} else if result != "" {
		alternative = result
	}

	return siteName, available, alternative
}

func IfElementOnPage(siteName, username, xpath string) (string, bool, string) {

	available := false
	alternative := ""

	result, err := InPage(siteName, username, xpath, 1)
	if err != nil {
		return siteName, true, "ERROR"
	}

	if result == username {
		available = true
	} else if result != "" {
		alternative = result
	}

	return siteName, available, alternative
}

func IfElementNotOnPage(siteName, username, xpath string) (string, bool, string) {

	available := false
	alternative := ""

	result, err := NotInPage(siteName, username, xpath, 1)
	if err != nil {
		return siteName, true, "ERROR"
	}

	if result == username {
		available = true
	} else if result != "" {
		alternative = result
	}

	return siteName, available, alternative
}
