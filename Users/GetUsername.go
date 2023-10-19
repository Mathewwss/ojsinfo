// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (u *User) GetUsername () error {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "username"
	query = query + " " + "FROM"
	query = query + " " + "users"
	query = query + " " + "WHERE"
	query = query + " " + "user_id = '" + fmt.Sprint(u.UID) + "'"
	query = query + ";"

	// Databae conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect db
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// View results
	for res.Next() {
		// Get value
		err = res.Scan(&u.Username)

		// Check errors
		if err != nil {
			// Stop
			return err

		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
