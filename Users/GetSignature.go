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

func (u *User) GetSignature () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Base query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			locale, setting_value
		FROM
			user_settings
		WHERE
			setting_name = 'signature'
			AND user_id = %v
		ORDER BY
			locale
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

	// Start variables
	u.Signature = map[string]string{}
	locale := ""
	value := ""

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &value)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update signature
		u.Signature[locale] = value

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
