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

func (j *Journal) GetGroups () error {
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
			t1.locale, t1.user_group_id,
			t1.setting_value
		FROM
			user_group_settings AS t1
		INNER JOIN
			user_groups AS t2
		ON
			t1.user_group_id = t2.user_group_id
		INNER JOIN
			journals AS t3
		ON
			t2.context_id = t3.journal_id
		WHERE
			t1.setting_name = 'name'
			AND t3.journal_id = %v
		ORDER BY
			t1.locale ASC, t1.user_group_id ASC
		;
	`, j.ID)

	// One line
	Regex.OneLine(&query)

	// Sql query

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	group := ""
	gid := -1
	locale := ""

	// Start map
	j.Groups = map[int]map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &gid, &group)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(j.Groups[gid]) == 0 {
			// start map
			j.Groups[gid] = map[string]string{}

		}

		// Update map
		j.Groups[gid][locale] = group

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
