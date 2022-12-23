package cmd

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/cikupin/ntfs-wr/pkg"
	"github.com/nexidian/gocliselect"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var mount = &cobra.Command{
	Use:   "mount",
	Short: "Mount external disk partition",
	Long:  "Mount external disk partition",
	PreRun: func(*cobra.Command, []string) {
		if !pkg.IsMacOS() {
			zlog.Fatal().Msg("this binary runs only for MacOS")
		}

		if !pkg.IsRoot() {
			zlog.Fatal().Msg("this binary runs only on root user")
		}
	},
	Run: func(c *cobra.Command, args []string) {
		disks, err := pkg.ListExternalDisk()
		if err != nil {
			zlog.Fatal().Err(err).Msg("cannot list external disk")
		}

		if len(disks) == 0 {
			zlog.Warn().Msgf("No external disk found using %s file system", pkg.FileSystem)
			return
		}

		fmt.Println()
		menu := gocliselect.NewMenu("Choose a disk partition")
		for k, v := range disks {
			item := fmt.Sprintf("Volume: %s, Partition: %s", v, k)
			menu.AddItem(item, k)
		}
		choice := menu.Display()

		if choice == "" {
			fmt.Println("\n")
			zlog.Info().Msg("No external disk selected. Exiting now...")
			return
		}

		dirName, err := pkg.MakeVolumeDir()
		if err != nil {
			zlog.Fatal().Err(err).Msg("cannot create temporary directory to mount partition")
		}

		fmt.Println()
		zlog.Info().
			Str("dir_name", dirName).
			Str("full_path", fmt.Sprintf("/Volumes/%s", dirName)).
			Msgf("folder was created on /Volumes to mount partition /dev/%s", choice)

		err = pkg.MountDisk(choice, dirName)
		if err != nil {
			zlog.Fatal().Err(err).
				Str("partition", fmt.Sprintf("/dev/%s", choice)).
				Str("full_path", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("cannot mount partition")
		}

		err = pkg.OpenVolumeDir(dirName)
		if err != nil {
			zlog.Error().Err(err).
				Str("volume_dir", fmt.Sprintf("/Volumes/%s", dirName)).
				Msg("cannot open volume dir using finder. you can manually open the directory on your computer and do copy-paste there.")
			return
		}

		Box := box.New(box.Config{
			Px:    2,
			Py:    1,
			Type:  "Classic",
			Color: "Cyan"})

		fmt.Println()
		Box.Print(
			"NTFS volume mounted succesfully!",
			fmt.Sprintf(`Partition mounted on /Volumes/%s. Please do copy-paste there.
Finder will open automatically. Or You can open the directory using the command below.

	$ ntfs-wr open /Volumes/%s

Use this command to unmount the partition:

	$ sudo ntfs-wr unmount /Volumes/%s
`, dirName, dirName, dirName),
		)
	},
}
