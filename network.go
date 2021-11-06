package main

import (
	"fmt"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func run() {
	app := cli.NewApp()

	app.Name = "Network Lookup"
	app.Usage = "IP, CNAME, Name Servers"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}

				for _, ipHost := range ip {
					fmt.Println(ipHost)
				}

				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}

				fmt.Println(cname)

				return nil
			},
		},
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}

				for _, nsHost := range ns {
					fmt.Println(nsHost.Host)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	run()
}
