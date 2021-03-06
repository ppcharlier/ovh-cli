package main

import (
	"github.com/codegangsta/cli"

	"github.com/toorop/govh"
	"github.com/toorop/govh/cloud"
)

// getCloudCmds return commands for cloud subsection
func getCloudCmds(client *govh.OVHClient) (cloudCmds []cli.Command) {
	cloud, err := cloud.New(client)
	if err != nil {
		return
	}
	cloudCmds = []cli.Command{
		{
			Name:        "passports",
			Usage:       "Return a list of cloud passports",
			Description: "Example: ovh cloud passports",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
			},
			Action: func(c *cli.Context) {
				passports, err := cloud.GetPassports()
				dieOnError(err)
				println(formatOutput(passports, c.Bool("json")))
				dieOk()
			},
		}, {
			Name:        "prices",
			Usage:       "Return a list of cloud prices",
			Description: "Example: ovh cloud prices",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
			},
			Action: func(c *cli.Context) {
				prices, err := cloud.GetPrices()
				dieOnError(err)
				println(formatOutput(prices, c.Bool("json")))
				dieOk()
			},
		}, {

			Name:        "projectids",
			Usage:       "Return a list of projects ID",
			Description: "Example: ovh cloud projectsid",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
			},
			Action: func(c *cli.Context) {
				ids, err := cloud.GetProjectIDs()
				dieOnError(err)
				println(formatOutput(ids, c.Bool("json")))
				dieOk()
			},
		}, {
			Name:        "project",
			Usage:       "Return project info",
			Description: "Example: ovh cloud project PROJECT_ID",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				project, err := cloud.GetProject(c.Args().First())
				dieOnError(err)
				println(formatOutput(project, c.Bool("json")))
				dieOk()
			},
		},
		/*
					   ovh cloud instance list PROJECTID
			           ovh cloud snapshot list PROJECTID
			           ovh cloud snapshot info PROJECTID SNAPSHOTID
					   ovh cloud snapshot create PROJECTID INSTANCEID
					   ovh cloud snapshot delete PROJECTID SNAPSHOTID
					   ovh cloud snapshot info PROJECTID SNAPSHOTID
					   ovh cloud instance backup PROJECTID INSTANCEID --hourly 24 --daily 7
		*/
		{
			Name:        "instance",
			Usage:       "Manage instances",
			Description: "desc",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Usage:       "List instances of project",
					Description: "Example: ovh cloud instance list PROJECTID",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						instances, err := cloud.GetInstances(c.Args().First())
						dieOnError(err)
						println(formatOutput(instances, c.Bool("json")))
						dieOk()
					},
				},
			},
		},
		{
			Name:        "snapshot",
			Usage:       "Manage snapshot (usage)",
			Description: "Manage snapshot",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Usage:       "Example: ovh cloud snapshot list PROJECTID",
					Description: "List snapshot of a project",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						instances, err := cloud.GetSnapshots(c.Args().First())
						dieOnError(err)
						println(formatOutput(instances, c.Bool("json")))
						dieOk()
					},
				},
			},
		},
	}
	return
}
