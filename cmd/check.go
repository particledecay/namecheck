package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ParticleDecay/namecheck/sites"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "namecheck",
	Short: "Namecheck is an availability checker for usernames",
	Long: `An availability checker for usernames. Currently queries
			Facebook, Twitter, Instagram, GitHub, Epic Games (Fortnite),
			and Twitch, with more support planned for the future.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Error: An action is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

var chkCmd = &cobra.Command{
	Use:   "all",
	Short: "Query all sites",
	Long:  `Query all sites for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		checkAll(args[0])
	},
}

var ghCmd = &cobra.Command{
	Use:   "github",
	Short: "Query GitHub only",
	Long:  `Query GitHub for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkGitHub(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

var fbCmd = &cobra.Command{
	Use:   "facebook",
	Short: "Query Facebook only",
	Long:  `Query Facebook for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkFacebook(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

var twitterCmd = &cobra.Command{
	Use:   "twitter",
	Short: "Query Twitter only",
	Long:  `Query Twitter for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkTwitter(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

var instaCmd = &cobra.Command{
	Use:   "instagram",
	Short: "Query Instagram only",
	Long:  `Query Instagram for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkInstagram(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

var twitchCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Query Twitch only",
	Long:  `Query Twitch for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkTwitch(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

var fortniteCmd = &cobra.Command{
	Use:   "fortnite",
	Short: "Query Fortnite only",
	Long:  `Query Fortnite for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		go checkFortnite(args[0], ch)
		name := <-ch
		printOutput(name)
	},
}

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

func usernameCheck(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Error: A username is required")
	}
	return nil
}

// Execute adds the subcommands and runs the main command
func Execute() {
	rootCmd.AddCommand(chkCmd)
	rootCmd.AddCommand(ghCmd)
	rootCmd.AddCommand(fbCmd)
	rootCmd.AddCommand(twitterCmd)
	rootCmd.AddCommand(instaCmd)
	rootCmd.AddCommand(twitchCmd)
	rootCmd.AddCommand(fortniteCmd)
	rootCmd.AddCommand(gplusCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printHeader(username string) {
	fmt.Println("")
	fmt.Println("Availability of '" + username + "'")
	fmt.Println("===================================")
	fmt.Println("")
}

func printOutput(name *sites.NameResult) {
	if name.Available == true && name.Alternate == "ERROR" { // This is just an error
		errorLine := fmt.Sprintf("%s:\tERROR\n", name.SiteName)
		fmt.Printf(errorLine)
		return
	}

	// Print the output
	output := "Available"
	if name.Available == false {
		output = fmt.Sprintf("Not %s", output)

		if name.Alternate != "" {
			output = fmt.Sprintf("%s (but '%s' is available)", output, name.Alternate)
		}
	}
	siteNameWithColon := fmt.Sprintf("%s:", name.SiteName)
	output = fmt.Sprintf("%-15s\t%s", siteNameWithColon, output)
	fmt.Println(output)
}

func checkAll(username string) {
	printHeader(username)

	ch := make(chan *sites.NameResult)

	// GitHub
	go checkGitHub(username, ch)

	// Facebook
	go checkFacebook(username, ch)

	// Twitter
	go checkTwitter(username, ch)

	// Instagram
	go checkInstagram(username, ch)

	// Twitch
	go checkTwitch(username, ch)

	// Fortnite
	go checkFortnite(username, ch)

	// Google+
	go checkGooglePlus(username, ch)

	for i := 0; i < len(sites.URLS); i++ {
		name := <-ch
		printOutput(name)
	}
	close(ch)
}

func checkGitHub(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("GitHub", username, false)
}

func checkFacebook(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("Facebook", username, false)
}

func checkTwitter(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("Twitter", username, false)
}

func checkInstagram(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("Instagram", username, true)
}

func checkTwitch(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("Twitch", username, false)
}

func checkFortnite(username string, ch chan *sites.NameResult) {
	ch <- sites.IfElementOnPage("Fortnite", username, "//div[contains(@class, 'profile__title')]//h2/following-sibling::div[1]/div[contains(text(), 'Not found')]")
}

func checkGooglePlus(username string, ch chan *sites.NameResult) {
	ch <- sites.IfPageNotFound("Google+", username, false)
}
