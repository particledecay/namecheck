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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Error: A username is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		checkAll(args[0])
	},
}

// Execute adds the subcommands and runs the main command.
func Execute() {
	rootCmd.AddCommand(chkCmd)
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

func printOutput(siteName string, available bool, alternative string) {
	if available == true && alternative == "ERROR" { // This is just an error
		errorLine := fmt.Sprintf("%s:\tERROR\n", siteName)
		fmt.Printf(errorLine)
		return
	}

	// Print the output
	output := "Available"
	if available == false {
		output = fmt.Sprintf("Not %s", output)

		if alternative != "" {
			output = fmt.Sprintf("%s (but '%s' is available)", output, alternative)
		}
	}
	siteNameWithColon := fmt.Sprintf("%s:", siteName)
	output = fmt.Sprintf("%-15s\t%s", siteNameWithColon, output)
	fmt.Println(output)
}

func checkAll(username string) {
	printHeader(username)

	// GitHub
	siteName, available, alternative := checkGitHub(username)
	printOutput(siteName, available, alternative)

	// Facebook
	siteName, available, alternative = checkFacebook(username)
	printOutput(siteName, available, alternative)

	// Twitter
	siteName, available, alternative = checkTwitter(username)
	printOutput(siteName, available, alternative)

	// Instagram
	siteName, available, alternative = checkInstagram(username)
	printOutput(siteName, available, alternative)

	// Twitch
	siteName, available, alternative = checkTwitch(username)
	printOutput(siteName, available, alternative)

	// Fortnite
	siteName, available, alternative = checkFortnite(username)
	printOutput(siteName, available, alternative)
}

func checkGitHub(username string) (string, bool, string) {
	return sites.IfPageNotFound("GitHub", username, false)
}

func checkFacebook(username string) (string, bool, string) {
	return sites.IfPageNotFound("Facebook", username, false)
}

func checkTwitter(username string) (string, bool, string) {
	return sites.IfPageNotFound("Twitter", username, false)
}

func checkInstagram(username string) (string, bool, string) {
	return sites.IfPageNotFound("Instagram", username, true)
}

func checkTwitch(username string) (string, bool, string) {
	return sites.IfPageNotFound("Twitch", username, false)
}

func checkFortnite(username string) (string, bool, string) {
	return sites.IfElementOnPage("Fortnite", username, "//div[contains(@class, 'profile__title')]//h2/following-sibling::div[1]/div[contains(text(), 'Not found')]")
}
