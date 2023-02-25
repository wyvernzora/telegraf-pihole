package ftl

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wyvernzora/telegraf-pihole/test"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestReader_Read_ZeroReader(t *testing.T) {
	r := &reader{}
	var buf = make([]Query, 1)
	n, err := r.Read(buf)

	assert.NilError(t, err)
	assert.Equal(t, n, 0)
}

func TestReader_Read_ResultLargerThanBuffer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.ValidValues[0]...).
		AddRow(test.ValidValues[1]...)
	mock.ExpectQuery(".+").WithArgs(0, 1).WillReturnRows(rows)

	reader := NewReader(db, 0)
	var buf = make([]Query, 1)

	// Read first page
	n, err := reader.Read(buf)
	assert.NilError(t, err)
	assert.Equal(t, n, 1)
	assert.DeepEqual(t, buf[0], Query{
		Id:           1,
		Timestamp:    time.Unix(1676602467, 0),
		Type:         TypeA,
		Decision:     DecisionAllowed,
		Reason:       ReasonForwarded,
		Domain:       "example.com",
		Client:       "127.0.0.1",
		Forward:      "10.20.30.40#53",
		ReplyType:    ReplyTypeIP,
		ReplyTime:    5400 * time.Microsecond,
		DnsSecStatus: DnsSecUnknown,
	})
	assert.Equal(t, reader.Position(), int64(1))
}

func TestReader_Read_ResultSmallerThanBuffer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.ValidValues[0]...)
	mock.ExpectQuery(".+").WithArgs(0, 10).WillReturnRows(rows)

	reader := NewReader(db, 0)
	var buf = make([]Query, 10)

	// Read first page
	n, err := reader.Read(buf)
	assert.NilError(t, err)
	assert.Equal(t, n, 1)
	assert.Equal(t, reader.Position(), int64(1))
}

func TestReader_Read_GapInIds(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.LargeIdValues[0]...)
	mock.ExpectQuery(".+").
		WithArgs(0, 1).
		WillReturnRows(rows)

	reader := NewReader(db, 0)
	var buf = make([]Query, 1)

	// Read first page
	n, err := reader.Read(buf)
	assert.NilError(t, err)
	assert.Equal(t, n, 1)
	assert.Equal(t, reader.Position(), int64(42))
}

func TestReader_Read_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	mock.ExpectQuery(".+").
		WithArgs(0, 1).
		WillReturnError(errors.New("test error"))

	reader := NewReader(db, 0)
	var buf = make([]Query, 1)

	n, err := reader.Read(buf)
	assert.ErrorContains(t, err, "test error")
	assert.Equal(t, n, 0)
}

func TestReader_Read_ScanQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}

	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.InvalidValues[0]...)
	mock.ExpectQuery(".+").
		WithArgs(0, 1).
		WillReturnRows(rows)

	reader := NewReader(db, 0)
	var buf = make([]Query, 1)

	// Read first page
	n, err := reader.Read(buf)
	assert.ErrorContains(t, err, "unexpected value: 99")
	assert.Equal(t, n, 0)
}
