// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
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
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "locale, setting_value, setting_name"
	query = query + " " + "FROM"
	query = query + " " + "user_settings"
	query = query + " " + "WHERE"
	query = query + " " + "("
	query = query + " " + "setting_name = 'givenName'"
	query = query + " " + "OR setting_name = 'familyName'"
	query = query + " " + ")"
	query = query + " " + "AND user_id = '" + fmt.Sprint(u.UID) + "'"
	query = query + " " + "AND setting_value <> ''"
	query = query + " " + "ORDER BY"
	query = query + " " + "locale ASC,"
	query = query + " " + "setting_name DESC"
	query = query + " " + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	locale := ""
	setting := ""
	name := ""
	u.RealNames = map[string]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &name, &setting)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		if setting == "givenName" {

			u.RealNames[locale] = name

		} else if setting == "familyName" {
			u.RealNames[locale] = u.RealNames[locale] + " " + name

		}

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
