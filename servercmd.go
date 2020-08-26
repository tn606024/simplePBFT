package main

import (
	"github.com/urfave/cli/v2"
)

var (
	nodeIdFlag = &cli.IntFlag{
		Name:	"id",
		Usage:	"id",
		Required: true,
	}
	nodeSubCommand = &cli.Command{
		Name:		 "node",
		Usage: 		 "start pbft node",
		Description: "start pbft node",
		ArgsUsage: 	 "<id>",
		Flags: []cli.Flag{
			nodeIdFlag,
		},
		Action: func(c *cli.Context) error {
			nodeId := c.Int("id")
			server := NewServer(nodeId)
			server.Start()
			return nil
		},
	}
	clientSubCommand = &cli.Command{
		Name:		 "client",
		Usage: 		 "start pbft client",
		Description: "start pbft client",
		ArgsUsage: 	 "",
		Action: func(c *cli.Context) error {
			client := NewClient()
			client.Start()
			return nil
		},
	}
	PBFTCommand = &cli.Command{
		Name:	"pbft",
		Usage:	"pbft commands",
		ArgsUsage: "",
		Category: "pbft Commands",
		Description: "",
		Subcommands: []*cli.Command{
			nodeSubCommand,
			clientSubCommand,
		},
	}
)