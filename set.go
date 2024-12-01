package main

type Set map[string]bool

func (s Set) Add(val string) {
	exists := s.Exists(val)
	if !exists {
		s[val] = true
	}
}

func (s Set) Exists(val string) bool {
	_, exists := s[val]
	return exists
}

func (s Set) Delete(val string) {
	exists := s.Exists(val)
	if exists {
		delete(s, val)
	}
}

func (s Set) Keys() []string {
	keys := make([]string, 0, len(s))

	for k, _ := range s {
		keys = append(keys, k)
	}

	return keys
}
