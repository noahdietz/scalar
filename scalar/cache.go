package scalar

import (
  "sort"
)

// Cache cache for active HPA created by Scalar
type Cache struct {
  Scalars map[string][]string
}

// Add adds an HPA reference keyed by the namespace and name
func (s *Cache) Add(name string, namespace string) {
  if _, exists := s.Scalars[namespace]; exists {
    s.Scalars[namespace] = append(s.Scalars[namespace], name)
  } else {
    s.Scalars[namespace] = []string{name}
  }
}

// Remove removes an HPA reference keyed by the namespace and name
func (s *Cache) Remove(name string, namespace string) {
  set, exists := s.Scalars[namespace]
  if !exists {
    return
  }

  ndx := sort.SearchStrings(set, name)
  s.Scalars[namespace] = append(set[:ndx], set[ndx+1:]...)
}

// Contains returns whether or not the cache contains a given HPA reference
func (s *Cache) Contains(name string, namespace string) bool {
  if set, exists := s.Scalars[namespace]; exists {
    for _, hpa := range set {
      if hpa == name {
        return true
      }
    }
  }

  return false
}

// InitCache initializes the scalar cache
func InitCache() (cache *Cache) {
  cache = &Cache{}
  cache.Scalars = make(map[string][]string)

  return
}