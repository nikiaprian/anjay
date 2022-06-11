package services

import "database/sql"

type SQLLite3 struct {
	*sql.DB
}
