package cmd

import (
	"fmt"
	"os"

	"github.com/cikupin/ntfs-wr/pkg"
	"github.com/jedib0t/go-pretty/v6/table"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var listExternalDisk = &cobra.Command{
	Use:   "list",
	Short: "List external disk volumes",
	Long:  "List external disk volumes",
	PreRun: func(*cobra.Command, []string) {
		if !pkg.IsMacOS() {
			zlog.Fatal().Msg("this binary runs only for MacOS")
		}
	},
	Run: func(c *cobra.Command, args []string) {
		disks, err := pkg.ListExternalDisk()
		if err != nil {
			zlog.Fatal().Err(err).Msg("cannot list external disk")
		}

		fmt.Println()
		printlnDiskInTable(disks)
	},
}

func printlnDiskInTable(disks map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Volume", "Partition"})

	var rows []table.Row
	for k, v := range disks {
		rows = append(rows, table.Row{v, k})
	}

	t.AppendRows(rows)
	t.Render()
}
