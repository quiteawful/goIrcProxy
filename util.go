package main

func isWhitelisted(user string) bool {
	whitelist := [...]string{"doclol", "marduk"}

	for _, w := range whitelist {
		if w == user {
			return true
		}
	}
	return false
}
