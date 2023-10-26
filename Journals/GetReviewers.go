// ----------------------------- Package ---------------------------- //

package Journals

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

func (j *Journal) GetReviewers () error {
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

	// Database connection settings
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect db
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check connection
	err = db.Ping()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Run query
	res, err := db.Query(query)

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
