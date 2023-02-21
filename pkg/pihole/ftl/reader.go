package ftl

import (
	"database/sql"
)

type Reader interface {
	Read([]Query) (int, error)
	Position() int64
}

func NewReader(db *sql.DB, position int64) Reader {
	return &reader{db, position}
}

type reader struct {
	sqldb    *sql.DB
	position Id
}

func (r *reader) Read(out []Query) (int, error) {
	if r.sqldb == nil {
		return 0, nil
	}

	const sqlQuery = `
		SELECT id, timestamp, type, status, domain, client, forward, reply_type, reply_time, dnssec 
		FROM queries
		WHERE id > $1
		ORDER BY id ASC
		LIMIT $2;
		`
	rows, err := r.sqldb.Query(sqlQuery, r.position, len(out))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var i int
	for i = 0; i < len(out) && rows.Next(); i++ {
		if err := ScanQuery(rows, &out[i]); err != nil {
			return i, err
		}
	}
	if i > 0 {
		r.position = out[i-1].Id
	}
	return i, rows.Err()
}

func (r *reader) Position() int64 {
	return r.position
}
