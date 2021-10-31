package bv5

func SystemEvaluateModelCommons(m ModelCommons, session SessionContext) []string {
	if m.UserId_() == session.SessionId() {
		if m.IsActive_() {
			return []string{
				ActionOwner,
				ActionManage,
				ActionView,
			}
		} else {
			return []string{
				ActionView,
			}
		}
	}
	if m.CommonAccess_() == CommonAccessPublic {
		return []string{
			ActionPublicView,
		}
	}
	return []string{}
}
