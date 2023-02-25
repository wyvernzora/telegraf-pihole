package test

import (
	"database/sql/driver"
)

var ColumnNames = []string{
	"id",
	"timestamp",
	"type",
	"status",
	"domain",
	"client",
	"forward",
	"reply_type",
	"reply_time",
	"dnssec",
}

var Timestamp int64 = 1676602467

var Domains = []string{
	"example.com",
	"random.org",
}
var IPs = []string{
	"127.0.0.1",
	"192.168.1.1",
}
var Forwards = []string{
	"10.20.30.40#53",
}

var ValidValues = [][]driver.Value{
	{1, Timestamp, 1, 2, Domains[0], IPs[0], Forwards[0], 4, 0.0054, 0},
	{2, Timestamp, 6, 2, Domains[1], IPs[0], Forwards[0], 2, 0.0011, 0},
}
var LargeIdValues = [][]driver.Value{
	{42, Timestamp, 6, 2, Domains[1], IPs[0], Forwards[0], 2, 0.0011, 0},
}
var InvalidValues = [][]driver.Value{
	{1, Timestamp, 99, 2, Domains[1], IPs[0], Forwards[0], 2, 0.0011, 0},
}
