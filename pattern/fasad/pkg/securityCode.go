package pkg

type securityCode struct {
	code string
}

func newSecurityCode(code string) *securityCode {
	return &securityCode{
		code: code,
	}
}
func (s *securityCode) checkSecurityCode(code string) bool {
	if s.code == code {
		return true
	}
	return false
}
