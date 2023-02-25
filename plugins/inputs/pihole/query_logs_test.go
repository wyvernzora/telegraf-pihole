package pihole

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/influxdata/telegraf/testutil"
	"github.com/wyvernzora/telegraf-pihole/test"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_gatherQueryLogs_ConfigFalse(t *testing.T) {
	pihole := New()
	pihole.GatherQueryLogs = false

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.NilError(t, err)
	assert.Equal(t, acc.HasMeasurement("query_log"), false)
	assert.Equal(t, acc.HasMeasurement("query_log_stats"), false)
}

func Test_gatherQueryLogs_OpenDbError(t *testing.T) {
	pihole := New()
	pihole.openDatabase = func(string) (*sql.DB, error) {
		return nil, errors.New("test error")
	}

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.Error(t, err, "test error")
}

func Test_gatherQueryLogs_ReaderError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(".*").WillReturnError(errors.New("test error"))

	pihole := New()
	pihole.openDatabase = func(string) (*sql.DB, error) {
		return db, nil
	}

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.Error(t, err, "test error")
}

func Test_gatherQueryLogs_ReadQueryLog(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.ValidValues[0]...)
	mock.ExpectQuery(".+").
		WithArgs(0, 5000).
		WillReturnRows(rows)

	pihole := New()
	pihole.Log = testutil.Logger{}
	pihole.GatherPluginMetrics = false
	pihole.openDatabase = func(string) (*sql.DB, error) {
		return db, nil
	}

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.NilError(t, err)
	assert.Equal(t, acc.HasMeasurement("query_log"), true)
	assert.Equal(t, acc.HasMeasurement("telegraf_pihole"), false)
}

func Test_gatherQueryLogs_PluginMetrics(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows(test.ColumnNames).
		AddRow(test.ValidValues[0]...)
	mock.ExpectQuery(".+").
		WithArgs(0, 5000).
		WillReturnRows(rows)

	pihole := New()
	pihole.Log = testutil.Logger{}
	pihole.openDatabase = func(string) (*sql.DB, error) {
		return db, nil
	}

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.NilError(t, err)
	assert.Equal(t, acc.HasMeasurement("query_log"), true)
	assert.Equal(t, acc.HasMeasurement("telegraf_pihole"), true)
}

func Test_gatherQueryLogs_PluginMetricsZeroResults(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows(test.ColumnNames)
	mock.ExpectQuery(".+").
		WithArgs(0, 5000).
		WillReturnRows(rows)

	pihole := New()
	pihole.Log = testutil.Logger{}
	pihole.openDatabase = func(string) (*sql.DB, error) {
		return db, nil
	}

	var acc = &testutil.Accumulator{}
	err := pihole.gatherQueryLogs(acc)

	assert.NilError(t, err)
	assert.Equal(t, acc.HasMeasurement("query_log"), false)
	assert.Equal(t, acc.HasMeasurement("telegraf_pihole"), true)
}
