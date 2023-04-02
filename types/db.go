package types

import "database/sql/driver"

type DB interface {
	Driver() driver.Driver
}
