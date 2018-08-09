# Namecheck [![CircleCI branch](https://img.shields.io/circleci/project/github/RedSparr0w/node-csgo-parser/master.svg?style=flat-square&logo=circleci)](https://circleci.com/gh/ParticleDecay/namecheck/tree/master)
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
Adding a new site should be pretty straightforward.

1. Create a new file for that site in the [sites](/sites) directory and implement a few methods:
```go
// googleplus.go
package sites

type GooglePlus struct {
	Site
}

func (g GooglePlus) Name() string {
	return "Google+"
}

func (g GooglePlus) URL() string {
	return "https://plus.google.com/+%s"
}

func (g GooglePlus) UserAgent() string {
	return ""
}

func (g GooglePlus) Check(username string, ch chan *NameResult) {
	// IfPageNotFound(site Site, username string, fakeUA bool)
	ch <- IfPageNotFound(g, username, false)
}
```

2. Add your new site to the `checkAll` function in [check.go](/cmd/check.go):
```go
func checkAll(username string) {
	...
	checks := []sites.Site{
		sites.GitHub{},
		sites.Facebook{},
		sites.Twitter{},
		sites.Twitch{},
		sites.Fortnite{},
		sites.Instagram{},
		sites.GooglePlus{}, // Add a line for your site
	}
	...
}
```

3. Add a separate function for your site (so it can be checked individually) and make sure it's added as a subcommand:
```go
var gplusCmd = &cobra.Command{
	Use:   "gplus", // This will be the command name
	Short: "Query Google+ only",
	Long:  `Query Google+ for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		gPlus := sites.GooglePlus{}
		go gPlus.Check(args[0], ch)
		printOutput(<-ch)
	},
}
...
func Execute() {
	...
	rootCmd.AddCommand(gplusCmd) // This adds your function as a subcommand
	...
}
```

4. Build it and make sure it works!
```bash
$ go build
$ ./namecheck gplus ArsTechnica
Google+:       	Not Available
```