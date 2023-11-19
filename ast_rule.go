type myrule struct {
	issue.MetaData
	gosec.CallList
}
func (r *myrule) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) { // HL
	if node := r.ContainsPkgCallExpr(n, c, false); node != nil {
		return c.NewIssue(n, r.ID(), r.What, r.Severity, r.Confidence), nil
	}
	return nil, nil
}
func NewMyRule(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	rule := &myrule{
		CallList: gosec.NewCallList(),
		MetaData: issue.MetaData{
			ID:         id,
			What:       "Potential HTTP request found",
			Severity:   issue.Medium,
			Confidence: issue.Medium,
		},
	}
	rule.AddAll("net/http", "Do", "Get", "Head", "Post", "PostForm", "RoundTrip")
	return rule, []ast.Node{(*ast.CallExpr)(nil)} // HL
}

func (r *myrule) ID() string {
	return r.MetaData.ID
}

