// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
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
	query := fmt.Sprintf(`
		SELECT DISTINCT
			user_id, email, username, url, phone, country
		FROM
			users
		WHERE
			username = '%v'
			OR email = '%v'
		;
	`, identity, identity)

	// Same line
	Regex.OneLine(&query)

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
		err = res.Scan(
			&u.UID, &u.Email, &u.Username, &u.URL, &u.Phone, &u.Country,
		)

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
