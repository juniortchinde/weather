package main

type Store struct {
	stations map[string]Station
}

func NewStore() *Store {
	return &Store{stations: make(map[string]Station)}
}

func (s *Store) Put(st Station) {
	s.stations[st.Id] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (Station, bool) {
	st, ok := s.stations[id]
	return st, ok
}

func (s *Store) Delete(id string) bool {
	delete(s.stations, id)
	return s.Has(id)
}

func (s *Store) All() []Station {
	var stations []Station
	for _, st := range s.stations {
		stations = append(stations, st)
	}
	return stations
}
