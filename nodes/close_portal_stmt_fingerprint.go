// Auto-generated - DO NOT EDIT

package pg_query

func (node ClosePortalStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ClosePortalStmt")

	if node.Portalname != nil {
		ctx.WriteString(*node.Portalname)
	}
}