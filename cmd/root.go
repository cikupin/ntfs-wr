package cmd

import (
	"github.com/cikupin/ntfs-wr/pkg"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listExternalDisk)
	RootCmd.AddCommand(mount)
}

var RootCmd = &cobra.Command{
	Use:   "ntfs-wr",
	Short: "Run ntfs-wr command",
	Long:  "MacOS CLI to mount and enable write on NTFS file system on external drive",
	PreRun: func(*cobra.Command, []string) {
		if !pkg.IsMacOS() {
			zlog.Fatal().Msg("this binary runs only for MacOS")
		}
	},
	Run: func(c *cobra.Command, args []string) {
		c.HelpFunc()(c, args)
	},
}
