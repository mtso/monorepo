package server

type CommandTable map[string]Command

type Command func(...string) (string, error)

var commands = CommandTable{
	"GET": func(key ...string) (string, error) {
		if len(key) < 1 {
			return "", ErrNoKey
		}
		value, ok := store.Get(key[0])
		if !ok {
			return "", ErrNil
		}
		return value, nil
	},
	"SET": func(args ...string) (string, error) {
		if len(args) < 1 {
			return "", ErrNoKey
		} else if len(args) < 2 {
			return "", ErrNoValue
		}
		key := args[0]
		value := args[1]
		store.Set(key, value)
		return "SET 1", nil
	},
}
