package cli

import (
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Display Drive account information and API capabilities",
	Long:  "Retrieve and display information about the authenticated Drive account and supported API capabilities",
	RunE:  runAbout,
}

var aboutFields string

func init() {
	aboutCmd.Flags().StringVar(&aboutFields, "fields", "*", "Fields to retrieve")
	rootCmd.AddCommand(aboutCmd)
}

func runAbout(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	capabilities := map[string]interface{}{
		"version": "1.0.0",
		"api": map[string]interface{}{
			"supported_operations": []string{
				"files.list", "files.get", "files.upload", "files.download", "files.delete",
				"files.copy", "files.move", "files.trash", "files.restore", "files.revisions",
				"folders.create", "folders.list", "folders.delete", "folders.move",
				"permissions.list", "permissions.create", "permissions.update", "permissions.delete", "permissions.public",
				"drives.list", "drives.get",
				"auth.login", "auth.device", "auth.service-account", "auth.status", "auth.profiles", "auth.logout",
				"sheets.list", "sheets.get", "sheets.create", "sheets.batch-update",
				"sheets.values.get", "sheets.values.update", "sheets.values.append", "sheets.values.clear",
				"docs.list", "docs.get", "docs.read", "docs.create", "docs.update",
				"slides.list", "slides.get", "slides.read", "slides.create", "slides.update", "slides.replace",
				"admin.users.list", "admin.users.get", "admin.users.create", "admin.users.update",
				"admin.users.suspend", "admin.users.unsuspend", "admin.users.delete",
				"admin.groups.list", "admin.groups.get", "admin.groups.create", "admin.groups.update",
				"admin.groups.delete", "admin.groups.members.list", "admin.groups.members.add", "admin.groups.members.remove",
			},
			"supported_exports": []string{
				"pdf", "docx", "xlsx", "pptx", "txt", "html", "rtf", "csv",
			},
			"features": []string{
				"batch_operations", "path_resolution", "caching", "dry_run", "safety_checks",
				"shared_drives", "permissions", "revisions", "trash_management",
			},
		},
		"authentication": map[string]interface{}{
			"oauth2_flows": []string{"web", "device_code"},
			"scopes": []string{
				"https://www.googleapis.com/auth/drive",
				"https://www.googleapis.com/auth/drive.file",
				"https://www.googleapis.com/auth/drive.readonly",
				"https://www.googleapis.com/auth/drive.metadata.readonly",
				"https://www.googleapis.com/auth/spreadsheets",
				"https://www.googleapis.com/auth/spreadsheets.readonly",
				"https://www.googleapis.com/auth/documents",
				"https://www.googleapis.com/auth/documents.readonly",
				"https://www.googleapis.com/auth/presentations",
				"https://www.googleapis.com/auth/presentations.readonly",
				"https://www.googleapis.com/auth/admin.directory.user",
				"https://www.googleapis.com/auth/admin.directory.user.readonly",
				"https://www.googleapis.com/auth/admin.directory.group",
				"https://www.googleapis.com/auth/admin.directory.group.readonly",
			},
		},
		"output_formats": []string{"json", "table"},
		"configuration": map[string]interface{}{
			"config_file":     "~/.config/gdrive/config.json",
			"credentials_dir": "~/.config/gdrive/credentials",
			"cache_dir":       "~/.config/gdrive/cache",
		},
	}

	return out.WriteSuccess("about", capabilities)
}
