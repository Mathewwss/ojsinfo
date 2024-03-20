// ----------------------------- Package ---------------------------- //

package Submissions

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

// Get date start
func (s *Submission) GetAccess () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check publication
	err = s.GetPublicationInfo()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check publication status
	if s.Published == false {
		// Stop
		return nil

	}

	// Access type -> [assoc_type, file_type]
	access_type := [][]string{
		[]string{"Abstract", "1048585", "-1"},
		[]string{"HTML", "515", "1"},
		[]string{"PDF", "515", "2"},
		[]string{"Others", "515", "3"},
	}

	// Init value
	query := ""

	// View types
	for a := 0; a < len(access_type); a++ {
		// Temporary query
		tmp_query := fmt.Sprintf(`
			SELECT
				"%v" AS "Code",
				CASE
					WHEN
						SUM(metric) IS NULL
					THEN
						0
					ELSE
						SUM(metric)
				END AS "Count"
			FROM
				metrics
			WHERE
				submission_id = %v
				AND assoc_type = %v
		`, access_type[a][0], s.ID, access_type[a][1])

		// Same line
		Regex.OneLine(&tmp_query)

		// View file type
		if access_type[a][2] != "-1" {
			// Update values
			tmp_query = tmp_query + " "
			tmp_query = tmp_query + "AND file_type = "
			tmp_query = tmp_query + access_type[a][2]

		}

		// View loop
		if a != 0 {
			// Update value
			query = query + " "
			query = query + "UNION"
			query = query + " "
			query = query + tmp_query

		} else {
			// Update value
			query = query + " " + tmp_query
		}

	}

	// Finish query
	query = query + " "
	query = query + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start map
	s.Access = map[string]int64{
		"All": 0,
		"Abstract": 0,
		"HTML": 0,
		"PDF": 0,
		"Others": 0,
	}

	// Start values
	value := int64(0)
	code := ""

	// View results
	for res.Next() {
		// Get value
		err = res.Scan(&code, &value)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update map
		s.Access[code] = value
		s.Access["All"] = s.Access["All"] + value

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
