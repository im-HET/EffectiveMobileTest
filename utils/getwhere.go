package utils

func GetWhere(whereStrParam string, params ...Parameter) string {
	for _, p := range params {
		paramStr := p.String()
		if paramStr != "" {
			if whereStrParam != "" {
				whereStrParam += " and "
			}
			whereStrParam += paramStr
		}
	}
	if whereStrParam == "" {
		return whereStrParam
	}
	return "WHERE " + whereStrParam
}

func GetSet(params ...Parameter) string {
	resultStr := ""
	for _, p := range params {
		if len(p.values) > 0 && p.values[0] != "" {
			if resultStr != "" {
				resultStr += ", "
			}
			resultStr += p.Column + " " + p.Condition + " " + p.values[0]
		}
	}
	if resultStr == "" {
		return resultStr
	}
	return "SET " + resultStr
}
