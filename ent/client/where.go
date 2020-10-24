// Code generated by entc, DO NOT EDIT.

package client

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/prmsrswt/edu-accounts/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ClientID applies equality check predicate on the "clientID" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ClientIDEQ applies the EQ predicate on the "clientID" field.
func ClientIDEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientIDNEQ applies the NEQ predicate on the "clientID" field.
func ClientIDNEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientID), v))
	})
}

// ClientIDIn applies the In predicate on the "clientID" field.
func ClientIDIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClientID), v...))
	})
}

// ClientIDNotIn applies the NotIn predicate on the "clientID" field.
func ClientIDNotIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClientID), v...))
	})
}

// ClientIDGT applies the GT predicate on the "clientID" field.
func ClientIDGT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientID), v))
	})
}

// ClientIDGTE applies the GTE predicate on the "clientID" field.
func ClientIDGTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientID), v))
	})
}

// ClientIDLT applies the LT predicate on the "clientID" field.
func ClientIDLT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientID), v))
	})
}

// ClientIDLTE applies the LTE predicate on the "clientID" field.
func ClientIDLTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientID), v))
	})
}

// ClientIDContains applies the Contains predicate on the "clientID" field.
func ClientIDContains(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientID), v))
	})
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "clientID" field.
func ClientIDHasPrefix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientID), v))
	})
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "clientID" field.
func ClientIDHasSuffix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientID), v))
	})
}

// ClientIDEqualFold applies the EqualFold predicate on the "clientID" field.
func ClientIDEqualFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientID), v))
	})
}

// ClientIDContainsFold applies the ContainsFold predicate on the "clientID" field.
func ClientIDContainsFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientID), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.Client {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Client(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecret), v))
	})
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecret), v))
	})
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecret), v))
	})
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecret), v))
	})
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecret), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Client) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Client) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Client) predicate.Client {
	return predicate.Client(func(s *sql.Selector) {
		p(s.Not())
	})
}