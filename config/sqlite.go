package config

import (
	"os" //verifica a existencia dos arquivos do db e cria se necessario.

	"github.com/itslucasmiranda/gojob/schemas" //contêm schemas de banco de dados que serãõ migrados automaticamente
	"gorm.io/driver/sqlite"                    //Driver SQLite para o GORM, permitindo comunicação com o db
	"gorm.io/gorm"                             //facilita a comunicação com o db
)

// cria o db caso não exista, abre a conexão e realiza as migrações
func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db" //define onde o banco de dados será armazenado usando a variável dbPath
	// verifica a existência do banco de dados
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Cria o banco de dados e o diretório
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Cria a conexão com o banco de dados
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migração do schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Retorna ao DB
	return db, nil
}
