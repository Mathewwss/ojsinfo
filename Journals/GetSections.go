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

func (j *Journal) GetSections () error {
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
			t1.locale, t1.section_id, t1.setting_value
		FROM
			section_settings AS t1
		INNER JOIN
			sections AS t2
		ON
			t1.section_id = t2.section_id
		WHERE
			t1.setting_name = 'title'
			AND t2.journal_id = %v
		ORDER BY
			t1.locale ASC, t1.section_id ASC
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
	section_name := ""
	section_id := -1
	locale := ""

	// Start map
	j.Sections = map[int]map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &section_id, &section_name)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(j.Sections[section_id]) == 0 {
			// start map
			j.Sections[section_id] = map[string]string{}

		}

		// Update map
		j.Sections[section_id][locale] = section_name

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
