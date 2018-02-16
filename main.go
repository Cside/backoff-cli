package main

import (
	"github.com/jpillora/backoff"

	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.Int64Flag{Name: "min"},
		cli.Int64Flag{Name: "max"},
		cli.Int64Flag{Name: "factor"},
	}

	app.Action = func(c *cli.Context) error {
		min := c.Int64("min")
		max := c.Int64("max")
		factor := c.Int64("factor")

		if min == 0 || max == 0 || factor == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		b := &backoff.Backoff{
			Min:    time.Duration(min) * time.Second,
			Max:    time.Duration(max) * time.Second,
			Factor: float64(factor),
		}
		var total int64
		for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
			n := int64(b.Duration().Seconds())
			total += n
			fmt.Printf("%d\t%d\t%d\n", i, n, total)
		}
		return nil
	}

	app.Run(os.Args)
}
