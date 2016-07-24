package slimprotocol

import ()

func VERSION() string {
	return "Slim -- V0.3\n"
}

func BYE() string {
	return "bye"
}

func EXCEPTION(reason string) string {
	return "__EXCEPTION__:message:<<" + reason + ".>>"
}

func ABORT(reason string) string {
	return "__EXCEPTION__:ABORT_SLIM_TEST:message:<<" + reason + ".>>"
}
