/*
Copyright © 2024 David Lynch davite3@protonmail.com
*/
package cmd

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yoink",
	Short: "yoinks files from a directory to your working directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		remove, err := cmd.Flags().GetBool("delete")
		if err != nil {
			return err
		}
		for _, arg := range args {
			err = yoink(arg, remove)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func yoink(path string, remove bool) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir(){
		//TODO SUPPORT DIRECTORIES
		return fmt.Errorf("cannot yoink %v as directories are not supported yet", path)
	}

	_, file := filepath.Split(path)
	destination := filepath.Join(wd, file)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(destination, data, 0o644)
	if err != nil {
		return err
	}

	if remove {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("delete", "d", false, "Remove a file from the source directory")
}
