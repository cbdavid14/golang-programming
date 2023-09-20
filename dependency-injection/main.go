package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"golang-programming/dependency-injection/configuration"
	"golang-programming/dependency-injection/database"
	"golang-programming/dependency-injection/repository"
)

func main() {
	App := fx.New(
		fx.Provide(
			context.Background,
			configuration.New,
			database.CreateSqliteConnection,
			repository.New,
		),
		fx.Invoke(
			configureLifeCycleHooks,
		),
	)

	App.Run()
}

func configureLifeCycleHooks(lc fx.Lifecycle, db *sqlx.DB, repo *repository.Repository) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				fmt.Println("Starting database connection")

				// Inserting some data
				err := repo.SaveUser(ctx, "David", "Consa", 32)
				if err != nil {
					return err
				}
				users, err := repo.GetUsers(ctx)
				if err != nil {
					return err
				}
				fmt.Printf("Users: %#v\n", users)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Closing database connection")
				return db.Close()
			},
		},
	)
}
