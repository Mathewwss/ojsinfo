// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Users"
import "fmt"
import "github.com/Mathewwss/ojsinfo/Regex"
import "database/sql"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *Submission) GetReviewers () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Base query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			t2.email, t1.round, t1.recommendation, t1.unconsidered,
			t1.cancelled, t1.declined, t1.date_assigned,
			t1.date_notified, t1.date_confirmed, t1.date_completed,
			t1.date_acknowledged, t1.date_due, t1.date_response_due
		FROM
			review_assignments AS t1
		INNER JOIN
			users AS t2
		ON
			t1.reviewer_id = t2.user_id
		WHERE
			t1.submission_id = %v
		ORDER BY
			t1.round DESC,
			t1.date_assigned ASC
		;
	`, s.ID)

	// One line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	var email string
	var round int64
	var recommendation sql.NullInt64
	var unconsidered sql.NullInt64
	var cancelled int64
	var declined int64
	var date_assigned sql.NullString
	var date_notified sql.NullString
	var date_confirmed sql.NullString
	var date_completed sql.NullString
	var date_acknowledged sql.NullString
	var date_due sql.NullString
	var date_response_due sql.NullString
	first := true

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(
			&email, &round, &recommendation, &unconsidered,
			&cancelled, &declined, &date_assigned, &date_notified,
			&date_confirmed, &date_completed, &date_acknowledged,
			&date_due, &date_response_due,
		)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Check loop
		if first == true {
			// Update values
			first = false
			s.Reviews = make([][]Review, round)

		}

		// Get user
		user, err := Users.New(email)

		// Check errors
		if err != nil {

			return err

		}

		// Create review
		rev := Review{}
		rev.UID = user.UID
		rev.Round = round
		rev.Cancelled = cancelled
		rev.Declined = declined
		rev.Unconsidered = unconsidered.Int64
		rev.Recommendation = fmt.Sprint(recommendation.Int64)
		rev.DateAssigned = date_assigned.String
		rev.DateNotified = date_notified.String
		rev.DateConfirmed = date_confirmed.String
		rev.DateCompleted = date_completed.String
		rev.DateAcknowledged = date_acknowledged.String
		rev.DateDue = date_due.String
		rev.DateResponseDue = date_response_due.String

		// Update reviews
		s.Reviews[round - 1] = append(s.Reviews[round - 1], rev)

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
