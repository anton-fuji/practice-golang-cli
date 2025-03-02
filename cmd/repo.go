package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo [owner] [repo]",
	Short: "Get repository info from GitHub",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := args[0], args[1]
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		fmt.Println("Repository Name:", result["full_name"])
		fmt.Println("Description:", result["description"])
		fmt.Println("Stars:", result["stargazers_count"])
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)

}
