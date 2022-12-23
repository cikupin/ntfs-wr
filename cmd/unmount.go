package cmd

import (
	"fmt"

	"github.com/cikupin/ntfs-wr/pkg"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var unmount = &cobra.Command{
	Use:     "unmount",
	Short:   "Unmount external disk partition",
	Long:    "Unmount external disk partition",
	Example: "$ ntfs-wr unmount <volume_dir_name>",
	PreRun: func(*cobra.Command, []string) {
		if !pkg.IsMacOS() {
			zlog.Fatal().Msg("this binary runs only for MacOS")
		}

		if !pkg.IsRoot() {
			zlog.Fatal().Msg("this binary runs only on root user")
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

		err := pkg.UnmountDisk(dirName)
		if err != nil {
			zlog.Fatal().Err(err).
				Str("full_path", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("cannot unmount volume")
			return
		}

		err = pkg.RemoveVolumeDir(dirName)
		if err != nil {
			zlog.Warn().
				Str("full_path", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("volume unmounted succesfully, but failed to remove temporary volume dir. you can manually delete it.")
			return
		}

		zlog.Info().
			Str("full_path", fmt.Sprintf("/Volumes/%s", dirName)).
			Msg("volume unmounted succesfully")
	},
}
