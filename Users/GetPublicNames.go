// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (u *User) GetPublicNames () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Base query
	query := fmt.Sprintf(`
		SELECT
			locale, setting_value
		FROM
			user_settings
		WHERE
			setting_name = 'preferredPublicName'
			AND user_id = %v
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

	// Start variables
	u.PublicNames = map[string]string{}
	locale := ""
	name := ""

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &name)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update names
		u.PublicNames[locale] = name

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
