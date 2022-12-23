package cmd

import (
	"fmt"

	"github.com/cikupin/ntfs-wr/pkg"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var openVolume = &cobra.Command{
	Use:     "open",
	Short:   "Open mounted volume",
	Long:    "Open mounted volume",
	Example: "$ ntfs-wr open <volume_dir_name>",
	PreRun: func(*cobra.Command, []string) {
		if !pkg.IsMacOS() {
			zlog.Fatal().Msg("this binary runs only for MacOS")
		}
	},
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			zlog.Error().Msg("please supply volume dirname as an argument. exiting now...")
			return
		}

		dirName := args[0]
		if !pkg.IsVolumeExist(dirName) {
			zlog.Error().
				Str("volume_dir", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("volume directory is not exist")
			return
		}

		err := pkg.OpenVolumeDir(dirName)
		if err != nil {
			zlog.Error().Err(err).
				Str("volume_dir", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("cannot open volume dir using finder. you can manually open the directory on your computer and do copy-paste there.")
			return
		}
	},
}
