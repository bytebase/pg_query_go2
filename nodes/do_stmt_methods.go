// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DoStmt) MarshalJSON() ([]byte, error) {
	type DoStmtMarshalAlias DoStmt
	return json.Marshal(map[string]interface{}{
		"DOSTMT": (*DoStmtMarshalAlias)(&node),
	})
}

func (node *DoStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DoStmt) Deparse() string {
	panic("Not Implemented")
}