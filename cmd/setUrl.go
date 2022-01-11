package cmd

import (
	"github.com/hahwul/authz0/pkg/models"
	set "github.com/hahwul/authz0/pkg/set"
	"github.com/spf13/cobra"
)

var url string
var allowRole, denyRole []string

// setUrlCmd represents the setUrl command
var setUrlCmd = &cobra.Command{
	Use:   "setUrl",
	Short: "Append URL to Template",
	Run: func(cmd *cobra.Command, args []string) {
		url := models.URL{
			URL:       url,
			AllowRole: allowRole,
			DenyRole:  denyRole,
		}
		if len(args) >= 1 {
			set.SetURL(args[0], url)
		}
	},
}

func init() {
	rootCmd.AddCommand(setUrlCmd)

	setUrlCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "URL")
	setUrlCmd.PersistentFlags().StringSliceVar(&allowRole, "allowRole", []string{}, "Allow role names")
	setUrlCmd.PersistentFlags().StringSliceVar(&denyRole, "denyRole", []string{}, "Deny role names")
}
