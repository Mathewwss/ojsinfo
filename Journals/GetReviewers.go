// ----------------------------- Package ---------------------------- //

package Journals

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

func (j *Journal) GetReviewers () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "s1.email"
	query = query + " " + "FROM"
	query = query + " " + "("
	query = query + " " + "SELECT DISTINCT"
	query = query + " " + "t1.email, s1.date_completed"
	query = query + " " + "FROM"
	query = query + " " + "users AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "user_user_groups AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.user_id = t2.user_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "user_groups AS t3"
	query = query + " " + "ON"
	query = query + " " + "t2.user_group_id = t3.user_group_id"
	query = query + " " + "LEFT JOIN"
	query = query + " " + "("
	query = query + " " + "SELECT"
	query = query + " " + "t1.reviewer_id, t1.date_completed"
	query = query + " " + "FROM"
	query = query + " " + "review_assignments AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "submissions AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.submission_id = t2.submission_id"
	query = query + " " + "WHERE"
	query = query + " " + "t2.context_id = '" + fmt.Sprint(j.ID) + "'"
	query = query + " " + ") AS s1"
	query = query + " " + "ON"
	query = query + " " + "t1.user_id = s1.reviewer_id"
	query = query + " " + "WHERE"
	query = query + " " + "t3.role_id = '4096'"
	query = query + " " + "AND t3.context_id = '" + fmt.Sprint(j.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "s1.date_completed DESC"
	query = query + " " + ") AS s1"
	query = query + " " + ";"

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
