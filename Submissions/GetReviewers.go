// ----------------------------- Package ---------------------------- //

package Submissions

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

func (s *Submission) GetReviewers () error {
	// Base query
	query := "SELECT DISTINCT"
	query = query + " " + "t2.email, t1.round, t1.recommendation"
	query = query + " " + "FROM"
	query = query + " " + "review_assignments AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "users AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.reviewer_id = t2.user_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.date_confirmed IS NOT NULL"
	query = query + " " + "AND t1.date_completed IS NOT NULL"
	query = query + " " + "AND t1.declined = 0"
	query = query + " " + "AND t1.submission_id = '" + fmt.Sprint(s.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t1.round ASC,"
	query = query + " " + "t1.date_assigned ASC"
	query = query + " " + ";"

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

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

	// In OJS database decisions is int,
	// Compare decision number, on database, with decison names on
	// GUI system.
	possibles := map[int]string{
		1: "Aceitar",
		2: "Correções obrigatórias",
		3: "Submeter novamente para avaliação",
		4: "Submeter a outra revista",
		5: "Rejeitar",
		6: "Ver comentários",
	}

	// Start variables
	s.Reviewers = [][]string{}
	email := ""
	round := -1
	decision_id := -1

	// View results
	for res.Next() {

		// Get values
		err = res.Scan(&email, &round, &decision_id)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Reviewer info
		revs := []string{
			fmt.Sprint(round),
			email,
			possibles[decision_id],
		}

		// Update slice
		s.Reviewers = append(s.Reviewers, revs)

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
