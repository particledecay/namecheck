# Namecheck
#### A username availability checker.

I decided one day I wanted to change my username on one site to something simpler, but quickly realized that I might not be able to get the same username on other sites that mattered to me. Coincidentally, I wanted to learn Go. This project gave me an excuse. Beware, a Go n00b wrote this. You were warned...

## Installation
```bash
go get -u github.com/ParticleDecay/namecheck
```

## Usage
#### Query all sites for a username
```bash
$ namecheck all PepperoniKing

Availability of 'PepperoniKing'
===================================

GitHub:        	Available
Facebook:      	Available
Instagram:     	Available
Twitter:       	Not Available (but '_PepperoniKing' is available)
Twitch:        	Not Available
Fortnite:      	Not Available
```
#### Query a single site
```bash
$ namecheck github CodeGolfer
GitHub:        	Not Available (but '_CodeGolfer' is available)
```

## Contributing
- Clone the project.
```bash
git clone git@github.com:ParticleDecay/namecheck.git
```
- This project uses [dep](https://github.com/golang/dep) for dependency management.
```bash
dep ensure
```
- This project uses [Cobra](https://github.com/spf13/cobra) for command-line interactions. To [add more sites](#adding-another-site), you'll need to also add another command for that site.

### Adding Another Site
It should be very straightforward to add another site. Let's add Google+ as an example (I might as well check it in, too):
1. Every site needs a `check` function in [check.go](cmd/check.go). It's up to you to find out how you would "check" that a username exists for a particular site. In Google+'s case, you can simply request the page at https://plus.google.com/+USERNAME, so we can use the `IfPageNotFound` function. This function just checks whether a username is available by looking for a 404 status code. You can find all the available checking functions (and add your own :) in [checks.go](sites/checks.go).

The `check` function should expect a username and a channel:
```go
func checkGooglePlus(username string, ch chan *sites.NameResult) {
    // Args are the site name, username, and whether a browser-like user agent needs to be sent.
	ch <- sites.IfPageNotFound("Google+", username, false)
}
```

2. The target URL needs to be added to the `URLS` map in [sites.go](sites/sites.go):
```go
var URLS = map[string]string{
    ...
	"Google+":   "https://plus.google.com/+%s",
}
```
The `"Google+"` string that we pass as the site name will be used to pull this URL. The `%s` in the URL string represents where the username will be inserted when the request is made.

3. Add the check function you wrote in step 1 to the `checkAll` function in [check.go](cmd/check.go), just above the loop:
```go
	// Fortnite
    go checkFortnite(username, ch)
    
    // Google+
    go checkGooglePlus(username, ch) // This is the new line you're adding

	for i := 0; i < len(sites.URLS); i++ {
		name := <-ch
		printOutput(name)
	}
	close(ch)
```

4. Preferably (though not required), add a specific command for the new site in [check.go](cmd/check.go) so Cobra can pick it up and allow it to be called individually (like `namecheck gplus USERNAME`):
```go
var gplusCmd = &cobra.Command{
	Use:   "gplus", // This will be the command name
	Short: "Query Google+ only",
	Long:  `Query Google+ for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkGooglePlus(args[0], ch) // This is the only line that needs to be specific to your site
		name := <-ch
		printOutput(name)
	},
}
```
... and add that command as a subcommand to the main command:
```go
func Execute() {
	rootCmd.AddCommand(chkCmd)
	rootCmd.AddCommand(ghCmd)
	rootCmd.AddCommand(fbCmd)
	rootCmd.AddCommand(twitterCmd)
	rootCmd.AddCommand(instaCmd)
	rootCmd.AddCommand(twitchCmd)
	rootCmd.AddCommand(fortniteCmd)
	rootCmd.AddCommand(gplusCmd) // This is the line you're adding
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

5. Build it and make sure it works!
```bash
$ go build
$ ./namecheck gplus ArsTechnica
Google+:       	Not Available
```