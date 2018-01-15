package models

// history table in database
type History struct{
	History_id       string
	History_date     string
	History_code     string
	History_by       string
	History_content  string
	History_notes    string
}

func ModelsUpdateHistory(id, date, code, by, content, notes string) error {
	sql_query := `INSERT INTO history (history_id, history_date, history_code, history_by, history_content, history_notes) VALUES
	(?, ?, ?, ?, ?, ?)
	`
	x, err := db.Queryx(sql_query, id, date, code, by, content, notes)
	defer x.Close()
	return err
}