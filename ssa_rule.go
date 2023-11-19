// START OMIT
func newMyAnalyzer(id string, description string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     id,
		Doc:      description,
		Run:      runMyCheck,
		Requires: []*analysis.Analyzer{buildssa.Analyzer},
	}
}

func runMyCheck(pass *analysis.Pass) (interface{}, error) {
	ssaResult, err := getSSAResult(pass) // HL
	if err != nil {
		return nil, err
	}
	for _, mcall := range ssaResult.SSA.SrcFuncs { // HL
		for _, block := range mcall.DomPreorder() {
			for _, instr := range block.Instrs {
				switch instr := instr.(type) {
				case *ssa.Call: // HL
					if fn, ok := instr.Call.Value.(*ssa.Function); ok { // HL
						if fn.Name() == "net/http." { // HL
							// TODO
						}
					}
					// END OMIT
				}
			}
		}
	}
}
