package sites

// Twitch (https://twitch.tv)
type Twitch struct {
	Site
}

// Name returns the name of the site to be displayed
func (t Twitch) Name() string {
	return "Twitch"
}

// URL is a path for the site with a template for the username
func (t Twitch) URL() string {
	return "https://api.twitch.tv/kraken/users/%s?client_id=bwllomi5jnmezyug7aprtkjqeb42hh"
}

// UserAgent returns an optional string to explicitly provide a User-Agent
func (t Twitch) UserAgent() string {
	return ""
}

// Check queries the Site for `username`
func (t Twitch) Check(username string, ch chan *NameResult) {
	ch <- IfPageNotFound(t, username)
}
