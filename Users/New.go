// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
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
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return User{}, err

	}

	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "user_id"
	query = query + " " + "FROM"
	query = query + " " + "users"
	query = query + " " + "WHERE"
	query = query + " " + "email = '" + identity + "'"
	query = query + " " + "OR username = '" + identity + "'"
	query = query + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

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
