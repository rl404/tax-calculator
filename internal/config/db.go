package config

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rl404/tax-calculator/internal/constant"
	"github.com/rl404/tax-calculator/internal/model"
	"github.com/rl404/tax-calculator/internal/utils"
)

// tableList is list of all required tables.
var tableList = []interface{}{
	model.Item{},
}

// InitDB to intiate db connection.
func (c *Config) InitDB() (db *gorm.DB, err error) {
	if c.Address == "" {
		return nil, constant.ErrRequiredDB
	}

	// Split address and port.
	split := strings.Split(c.Address, ":")
	if len(split) != 2 {
		return nil, constant.ErrInvalidDB
	}

	// Open db connection.
	conn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", split[0], split[1], c.DB, c.User, c.Password)
	db, err = gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	// Set base connection setting.
	db.DB().SetMaxIdleConns(DefaultMaxIdleConn)
	db.DB().SetMaxOpenConns(DefaultMaxOpenConn)
	db.DB().SetConnMaxLifetime(DefaultConnMaxLifeTime)
	db.SingularTable(true)
	db.LogMode(false)

	// Set default schema.
	err = db.Exec(fmt.Sprintf("SET search_path TO %s", c.Schema)).Error
	if err != nil {
		return db, err
	}

	gorm.DefaultTableNameHandler = func(dbVeiculosGorm *gorm.DB, defaultTableName string) string {
		if c.Schema == "" {
			c.Schema = "public"
		}
		return c.Schema + "." + defaultTableName
	}

	// Validate db structure.
	err = c.validateDB(db)
	if err != nil {
		return db, err
	}

	return db, nil
}

// validateDB to validate db structure.
func (c *Config) validateDB(db *gorm.DB) error {
	// Schema check.
	if !c.isSchemaExist(db) {
		err := c.createSchema(db)
		if err != nil {
			return err
		}
	}

	// Tables check.
	existingTables := c.getExistingTables(db)
	for _, model := range tableList {
		tableName := db.NewScope(model).TableName()
		if !utils.InArray(existingTables, tableName) {
			err := c.createTable(db, model)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// isSchemaExist to check if schema is exist.
func (c *Config) isSchemaExist(db *gorm.DB) (isExist bool) {
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_namespace WHERE nspname = ?)", c.Schema).Row().Scan(&isExist)
	return isExist
}

// createSchema to create new schema.
func (c *Config) createSchema(db *gorm.DB) error {
	query := fmt.Sprintf("CREATE SCHEMA \"%s\" AUTHORIZATION \"%s\"", c.Schema, c.User)
	return db.Exec(query).Error
}

// getExistingTables to get list of existing tables.
func (c *Config) getExistingTables(db *gorm.DB) (tables []string) {
	rows, _ := db.Raw("SELECT concat(table_schema, '.', table_name)  FROM information_schema.tables WHERE table_schema = ?", c.Schema).Rows()
	defer rows.Close()
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tables = append(tables, tableName)
	}
	return tables
}

// CreateTable to create table.
func (c *Config) createTable(db *gorm.DB, model interface{}) error {
	return db.CreateTable(model).Error
}
