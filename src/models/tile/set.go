package tile

// largely inspired by https://www.bytesizego.com/blog/set-in-golang

type Set struct {
	elements map[Property]struct{}
}

func NewSet() Set {
	return Set{make(map[Property]struct{})}
}

func (s *Set) Add(p Property) {
	s.elements[p] = struct{}{}
}

func (s *Set) Remove(p Property) {
	delete(s.elements, p)
}

func (s *Set) Contains(p Property) bool {
	_, found := s.elements[p]
	return found
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []Property {
	properties := make([]Property, s.Size())
	i := 0
	for k := range s.elements {
		properties[i] = k
		i++
	}
	return properties
}

// TODO: add Union, Intersection, and Difference
