package command

import (
	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/command/app"
	"github.com/wincentrtz/gobase/gobase/command/db"
	"github.com/wincentrtz/gobase/gobase/command/generate"
	instance "github.com/wincentrtz/gobase/gobase/infrastructures/db"
)

func GetCommands() cli.Commands {
	return cli.Commands{
		{
			Name:        "serve",
			Aliases:     []string{"srv", "serve"},
			Usage:       "Serve your app",
			Description: "It will serve your app so you app will run",
			Category:    "serve",
			//TODO: replace with action function since it will be deprecated soon
			Action: func(c *cli.Context) {
				app.Serve()
			},
		},
		{
			Name:         "database",
			Aliases:      []string{"d", "db"},
			Usage:        "Do a certain database operation",
			Category:     "database",
			OnUsageError: nil,
			Subcommands: cli.Commands{
				{
					Name:     "migrate",
					Aliases:  []string{"mg", "mig", "migrate"},
					Usage:    "Do a database migration",
					Category: "database",
					Action: func(c *cli.Context) {
						db.Migrate(instance.Postgres())
					},
				},
				{
					Name:     "seed",
					Aliases:  []string{"sd", "seed"},
					Usage:    "Do a database seeder",
					Category: "database",
					Action: func(c *cli.Context) {
						db.Seed(instance.Postgres())
					},
				},
				{
					Name:     "fresh",
					Aliases:  []string{"fr", "fresh"},
					Usage:    "Drop and create table",
					Category: "database",
					Action: func(c *cli.Context) {
						db.Fresh(instance.Postgres())
					},
				}, {
					Name:     "clear",
					Aliases:  []string{"clr", "clear"},
					Usage:    "Drop table from migration",
					Category: "database",
					Action: func(c *cli.Context) {
						db.Drop(instance.Postgres())
					},
				},
			},
		},
		{
			Name:     "generate",
			Aliases:  []string{"gn", "gen", "generate"},
			Usage:    "Generate boilerplate for code",
			Category: "generate",
			Subcommands: cli.Commands{
				{
					Name:     "domain",
					Aliases:  []string{"dom", "domain"},
					Usage:    "Generate Domain file",
					Category: "generate",
					Action: func(c *cli.Context) {
						generate.Domain(c)
					},
				},
				{
					Name:     "response",
					Aliases:  []string{"res", "response"},
					Usage:    "Generate Response file",
					Category: "generate",
					Action: func(c *cli.Context) {
						generate.Response(c)
					},
				},
				{
					Name:     "request",
					Aliases:  []string{"req", "request"},
					Usage:    "Generate Request file",
					Category: "generate",
					Action: func(c *cli.Context) {
						generate.Request(c)
					},
				},
				{
					Name:     "seeder",
					Aliases:  []string{"sdr", "seeder"},
					Usage:    "Generate Seeder file",
					Category: "generate",
					Action: func(c *cli.Context) {
						generate.Seeder(c)
					},
				},
			},
		},
	}
}
