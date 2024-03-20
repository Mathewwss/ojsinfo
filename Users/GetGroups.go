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

// Get login name
func (u *User) GetGroups () error {
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
			t1.user_group_id, t3.setting_value, t4.setting_value
		FROM
			user_user_groups AS t1
		INNER JOIN
			user_groups AS t2
		ON
			t1.user_group_id = t2.user_group_id
		INNER JOIN
			user_group_settings AS t3
		ON
			t1.user_group_id = t3.user_group_id
		INNER JOIN
			journal_settings AS t4
		ON
			t2.context_id = t4.journal_id
		WHERE
			t3.locale = 'pt_BR'
			AND t3.setting_name = 'name'
			AND t4.locale = 'pt_BR'
			AND t4.setting_name = 'name'
			AND t1.user_id = %v
		ORDER BY
			t4.setting_value ASC,
			t1.user_group_id ASC
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

	// Start variable
	journal := ""
	group := ""
	gid := -1

	// Start map
	u.Groups = map[string]map[int]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&gid, &group, &journal)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(u.Groups[journal]) == 0 {
			// start map
			u.Groups[journal] = map[int]string{}

		}

		// Update map
		u.Groups[journal][gid] = group

	}

	// Show user
	return nil
}

// ------------------------------------------------------------------ //
