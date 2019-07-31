package cmd

func getFilename(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "study.json"
}
