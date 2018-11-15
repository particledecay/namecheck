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
		github := sites.GitHub{}
		go github.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var fbCmd = &cobra.Command{
	Use:   "facebook",
	Short: "Query Facebook only",
	Long:  `Query Facebook for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		facebook := sites.Facebook{}
		go facebook.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var twitterCmd = &cobra.Command{
	Use:   "twitter",
	Short: "Query Twitter only",
	Long:  `Query Twitter for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		twitter := sites.Twitter{}
		go twitter.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var instaCmd = &cobra.Command{
	Use:   "instagram",
	Short: "Query Instagram only",
	Long:  `Query Instagram for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		insta := sites.Instagram{}
		go insta.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var twitchCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Query Twitch only",
	Long:  `Query Twitch for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		twitch := sites.Twitch{}
		go twitch.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var fortniteCmd = &cobra.Command{
	Use:   "fortnite",
	Short: "Query Fortnite only",
	Long:  `Query Fortnite for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		fortnite := sites.Fortnite{}
		go fortnite.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var gplusCmd = &cobra.Command{
	Use:   "gplus",
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

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Query Reddit only",
	Long:  `Query Reddit for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		reddit := sites.Reddit{}
		go reddit.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Query Docker only",
	Long:  `Query Docker for the given <username>.`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		docker := sites.Docker{}
		go docker.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var dotComCmd = &cobra.Command{
	Use:   "dotcom",
	Short: "Lookup .com domain",
	Long:  `Lookup availability of a .com domain`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		dotcom := sites.DomainDotCom{}
		go dotcom.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var dotOrgCmd = &cobra.Command{
	Use:   "dotorg",
	Short: "Lookup .org domain",
	Long:  `Lookup availability of a .org domain`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		dotorg := sites.DomainDotOrg{}
		go dotorg.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var dotNetCmd = &cobra.Command{
	Use:   "dotnet",
	Short: "Lookup .net domain",
	Long:  `Lookup availability of a .net domain`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		dotnet := sites.DomainDotNet{}
		go dotnet.Check(args[0], ch)
		printOutput(<-ch)
	},
}

var dotIoCmd = &cobra.Command{
	Use:   "dotio",
	Short: "Lookup .io domain",
	Long:  `Lookup availability of a .io domain`,
	Args:  usernameCheck,
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan *sites.NameResult)
		dotio := sites.DomainDotIo{}
		go dotio.Check(args[0], ch)
		printOutput(<-ch)
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
	rootCmd.AddCommand(redditCmd)
	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(dotComCmd)
	rootCmd.AddCommand(dotOrgCmd)
	rootCmd.AddCommand(dotNetCmd)
	rootCmd.AddCommand(dotIoCmd)
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

	checks := []sites.Site{
		sites.GitHub{},
		sites.Facebook{},
		sites.Twitter{},
		sites.Twitch{},
		sites.Fortnite{},
		sites.Instagram{},
		sites.GooglePlus{},
		sites.Reddit{},
		sites.Docker{},
		sites.DomainDotCom{},
		sites.DomainDotOrg{},
		sites.DomainDotNet{},
		sites.DomainDotIo{},
	}

	for i := range checks {
		go checks[i].Check(username, ch)
	}

	for i := 0; i < len(checks); i++ {
		name := <-ch
		printOutput(name)
	}
	close(ch)
}
