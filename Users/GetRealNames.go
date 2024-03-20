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

func (u *User) GetRealNames () error {
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
			s1.locale,
			CONCAT(s1.setting_value, " ", s2.setting_value)
		FROM
			(
				SELECT DISTINCT
					locale, setting_value, setting_name
				FROM
					user_settings
				WHERE
					setting_name = 'givenName'
					AND setting_value <> ''
					AND user_id = %v
			) AS s1
		INNER JOIN
			(
				SELECT DISTINCT
					locale, setting_value, setting_name
				FROM
					user_settings
				WHERE
					setting_name = 'familyName'
					AND setting_value <> ''
					AND user_id = %v
			) AS s2
		ON
			s1.locale = s2.locale
		;
	`, u.UID, u.UID)

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
	locale := ""
	name := ""
	u.RealNames = map[string]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &name)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update name
		u.RealNames[locale] = name

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
