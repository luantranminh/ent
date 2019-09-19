// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/json/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldURL)))
		},
	)
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldURL)))
		},
	)
}

// RawIsNil applies the IsNil predicate on the "raw" field.
func RawIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldRaw)))
		},
	)
}

// RawNotNil applies the NotNil predicate on the "raw" field.
func RawNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldRaw)))
		},
	)
}

// DirsIsNil applies the IsNil predicate on the "dirs" field.
func DirsIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldDirs)))
		},
	)
}

// DirsNotNil applies the NotNil predicate on the "dirs" field.
func DirsNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldDirs)))
		},
	)
}

// IntsIsNil applies the IsNil predicate on the "ints" field.
func IntsIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldInts)))
		},
	)
}

// IntsNotNil applies the NotNil predicate on the "ints" field.
func IntsNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldInts)))
		},
	)
}

// FloatsIsNil applies the IsNil predicate on the "floats" field.
func FloatsIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldFloats)))
		},
	)
}

// FloatsNotNil applies the NotNil predicate on the "floats" field.
func FloatsNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldFloats)))
		},
	)
}

// StringsIsNil applies the IsNil predicate on the "strings" field.
func StringsIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldStrings)))
		},
	)
}

// StringsNotNil applies the NotNil predicate on the "strings" field.
func StringsNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldStrings)))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
