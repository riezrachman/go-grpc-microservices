package main

import (
	"os"
	"swing/server/pb"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func migrationStart() {

	if err := db_main.AutoMigrate(
		&pb.UserORM{},
		&pb.StoreORM{},
		&pb.ProductORM{},
	); err != nil {
		logrus.Fatalf("[main][func: migrationStart] Migration failed: %v", err)
		os.Exit(1)
	}

}

func runMigrationCmd() cli.Command {

	return cli.Command{
		Name:  "db-migrate",
		Usage: "run db migration",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {

			initDBMain()
			defer closeDBMain()

			logrus.Println("[main][func: migrationStart] Migration process begin...")

			migrationStart()

			logrus.Println("[main][func: migrationStart] Migration process finished...")

			return nil

		},
	}

}
