package main

import (
	"image"
	"math/rand"

	cli "github.com/jawher/mow.cli"
)

func CommandChop(cmd *cli.Cmd) {
	cmd.Command("shuffle", "", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			rand.Shuffle(len(images), func(i, j int) {
				images[i], images[j] = images[j], images[i]
			})
		}
	})
	cmd.Command("drop-every", "", func(cmd *cli.Cmd) {
		n := cmd.IntArg("NTH", 2, "")
		cmd.Action = func() {
			var decimated []image.Image
			for i := range images {
				if i%*n == 0 {
					continue
				}
				decimated = append(decimated, images[i])
			}
			images = decimated
		}
	})
	cmd.Command("drop-first", "", func(cmd *cli.Cmd) {
		n := cmd.IntArg("N", len(images)/2, "default: n/2")
		cmd.Action = func() {
			if *n >= len(images) {
				*n = len(images) - 1
			}
			images = images[*n:]
		}
	})
	cmd.Command("drop-last", "", func(cmd *cli.Cmd) {
		n := cmd.IntArg("N", len(images)/2, "default: n/2")
		cmd.Action = func() {
			if *n > len(images) {
				*n = len(images)
			}
			images = images[:len(images)-*n]
		}
	})
	cmd.Command("reverse", "", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			for i := 0; i <= len(images)/2; i++ {
				j := len(images) - 1 - i
				images[i], images[j] = images[j], images[i]
			}
		}
	})
}
