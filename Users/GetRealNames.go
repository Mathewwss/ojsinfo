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

func (u *User) GetRealNames () error {
	// Base query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "locale, setting_value"
	query = query + " " + "FROM"
	query = query + " " + "user_settings"
	query = query + " " + "WHERE"
	query = query + " " + "user_id = '" + fmt.Sprint(u.UID) + "'"

	// Database setitngs
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
	db, err := sql.Open(driver, con)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// name types
	setting_names := []string{
		"givenName",
		"familyName",
	}

	for a := 0; a < len(setting_names); a++ {
		// Check loop
		if a == 0 {
			// Start map
			u.RealNames = map[string]string{}

		}

		// Finaly query
		run := query + " " + "AND setting_name = '" + setting_names[a]
		run = run + "'"
		run = run + " " + "ORDER BY"
		run = run + " " + "locale"
		run = run + ";"

		// Run query
		res, err := db.Query(run)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Start variables
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

			// Empty name
			if name == "" {
				// Skip
				continue

			}

			// Check names
			if len(u.RealNames[locale]) == 0 {
				// First name
				u.RealNames[locale] = name

			} else {
				// Middles names
				u.RealNames[locale] = u.RealNames[locale] + " " + name

			}
		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
