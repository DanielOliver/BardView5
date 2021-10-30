package db

import (
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/rs/zerolog/log"
	"server/bardlog"
)

//ignore_me go:generate go-bindata -pkg migrations .

const (
	MigrationType       = "Migration"
	KeyMigrationVersion = "migration_version"
	KeyMigrationDirty   = "migration_dirty"
)

func Migrate(connectionString string) {
	logger := log.With().Str(bardlog.KeyLogType, MigrationType).Logger()

	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})
	d, err := bindata.WithInstance(s)
	if err != nil {
		logger.Fatal().
			Err(err).Msg("Failed to bindata.")
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, connectionString)
	if err != nil {
		logger.Fatal().
			Err(err).Msg("Failed to open migrations.")
	}
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		logger.Fatal().
			Err(err).Msg("Failed to fetch current version.")
	}
	if dirty {
		logger.Fatal().
			Uint(KeyMigrationVersion, version).
			Bool(KeyMigrationDirty, dirty).
			Msg("Current Migration is dirty.")
	}
	logger.Info().
		Uint(KeyMigrationVersion, version).
		Bool(KeyMigrationDirty, dirty).
		Msg("Migrating bardview5")
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().
				Uint(KeyMigrationVersion, version).
				Bool(KeyMigrationDirty, dirty).
				Err(err).Msg("No change made.")
		}else {
			logger.Fatal().
				Uint(KeyMigrationVersion, version).
				Bool(KeyMigrationDirty, dirty).
				Err(err).Msg("Failed to migrate.")
		}
	}

	logger.Info().
		Bool(KeyMigrationDirty, dirty).
		Uint(KeyMigrationVersion, version).
		Msg("Migrated bardview5")

	version, dirty, err = m.Version()
	if err != nil {
		logger.Fatal().
			Err(err).Msg("Failed to fetch current version after migration.")
	}

	logger.Info().
		Bool(KeyMigrationDirty, dirty).
		Uint(KeyMigrationVersion, version).
		Msg("bardview5 database version")
}
