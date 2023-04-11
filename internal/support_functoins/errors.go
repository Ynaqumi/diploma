package support_functoins

func ErrorToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
