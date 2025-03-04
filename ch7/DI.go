//dependency injection

package ch7

import "fmt"

func LogOutput(output string) {
	fmt.Println(output)
}

type DataStore struct {
	userData map[string]string
}

func (s DataStore) userNameForId(id string) (string, bool) {
	name, ok := s.userData[id]
	return name, ok
}

func factFunc() DataStore {
	return DataStore{
		userData: map[string]string{
			"1": "abdul",
			"2": "kiran",
			"3": "venkatesh",
		},
	}
}

// Your business logic needs some data to work with, so it requires a data store.
// You also want your business logic to log when it is invoked, so it depends on a logger.
// However, you don’t want to force it to depend on LogOutput or SimpleDataStore,
// because you might want to use a different logger or data store later.
type Store interface {
	userNameForId(id string) (string, bool)
}
type Log interface {
	logger(message string)
}

// make the LogOutput function meets the Log interface
type logAdapter func(message string)

func (lg logAdapter) logger(message string) {
	lg(message)
}
