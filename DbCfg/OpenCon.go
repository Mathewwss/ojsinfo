// ----------------------------- Package ---------------------------- //

package DbCfg

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import _ "github.com/go-sql-driver/mysql"
import "database/sql"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (c *DbCon) OpenCon () {
	// Open connection
	db, err := sql.Open(c.Driver, c.Settings)

	// Check errors
	if err != nil {
		// Update struct
		c.ConErr = err
		c.Opened = false

	} else {
		// Update struct
		c.ConErr = nil
		c.Con = db
		c.Opened = true

	}
}

// ------------------------------------------------------------------ //