// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
import "fmt"
import "strings"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// Get login name
func (u *User) GetLanguages () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			locales
		FROM
			users
		WHERE
			user_id = %v
			AND locales <> ''
		;
	`, u.UID)

	// Same line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Temp
	value := ""

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&value)

		// Check errors
		if err != nil {
			// Stop
			return err

		}
	}

	// Update value
	u.Languages = strings.Split(value, ":")

	// Show user
	return nil
}

// ------------------------------------------------------------------ //
