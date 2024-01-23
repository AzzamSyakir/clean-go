package migrate

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

func MigrateDB(db *sql.DB) error {
	// Memeriksa koneksi ke database
	err := db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal melakukan ping ke database")
	}
	// Panggil fungsi migrate untuk inisialisasi migrasi database
	if err := UserMigrate(db); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan migrasi user")
	}

	if err := TokenMigrate(db); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan migrasi token")
	}

	return err
}
