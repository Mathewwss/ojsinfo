// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "errors"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// Create user with exists on database
func New (identity string) (User, error) {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "user_id"
	query = query + " " + "FROM"
	query = query + " " + "users"
	query = query + " " + "WHERE"
	query = query + " " + "email = '" + identity + "'"
	query = query + " " + "OR username = '" + identity + "'"
	query = query + ";"

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return User{}, err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return User{}, err

	}

	// Start variables
	u := User{}

	// View results
	for res.Next() {
		// Get value
		err = res.Scan(&u.UID)

		// Check errors
		if err != nil {
			// Stop
			return User{}, err

		}
	}

	// Check user ID
	if u.UID == 0 {
		// Create erorr
		msg := "[ERROR] -> Not found user!"
		err = errors.New(msg)

		// Stop
		return User{}, err

	}

	// Finish
	return u, nil
}

// ------------------------------------------------------------------ //
