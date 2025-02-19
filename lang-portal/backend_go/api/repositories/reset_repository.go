package repositories

func ResetHistory() error {
	_, err := DB.Exec("DELETE FROM word_review_items")
	return err
}

func FullReset() error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM word_review_items")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM study_sessions")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM study_activities")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM words_groups")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM groups")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM words")
	if err != nil {
		return err
	}

	return nil
}
