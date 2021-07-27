package helper

import "database/sql"

func NullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}

	return ns
}

func StringValue(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
