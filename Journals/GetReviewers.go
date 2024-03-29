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

func (j *Journal) GetReviewers () error {
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
			s1.email
		FROM
			(
				SELECT DISTINCT
					t1.email, s1.date_completed
				FROM
					users AS t1
				INNER JOIN
					user_user_groups AS t2
				ON
					t1.user_id = t2.user_id
				INNER JOIN
					user_groups AS t3
				ON
					t2.user_group_id = t3.user_group_id
				LEFT JOIN
					(
						SELECT
							t1.reviewer_id, t1.date_completed
						FROM
							review_assignments AS t1
						INNER JOIN
							submissions AS t2
						ON
							t1.submission_id = t2.submission_id
						WHERE
							t2.context_id = %v
					) AS s1
				ON
					t1.user_id = s1.reviewer_id
				WHERE
					t3.role_id = '4096'
					AND t3.context_id = %v
				ORDER BY
					s1.date_completed DESC
			) AS s1
		;
	`, j.ID, j.ID)

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
	j.Reviewers = []string{}
	email := ""

	// View results
	for res.Next() {
		// Get values
		err := res.Scan(&email)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update slice
		j.Reviewers = append(j.Reviewers, email)

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
