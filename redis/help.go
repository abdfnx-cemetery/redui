package redis

type RedisHelp struct {
	Args    string
	Command string
	Desc    string
}

func RedisMatchedCommands(text string) []RedisHelp {
	var matched = make([]RedisHelp, 0)
	return matched
}
