// ----------------------------- Package ---------------------------- //

package Journals

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

func (j *Journal) GetNames () error {
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
			locale, setting_value
		FROM
			journal_settings
		WHERE
			setting_name = 'name'
			AND journal_id = %v
		ORDER BY
			locale
		;
	`, j.ID)

	// Same line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	name := ""
	locale := ""
	j.Names = map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &name)
		if err != nil {
			// Stop
			return err

		}

		// Update map
		j.Names[locale] = name

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
