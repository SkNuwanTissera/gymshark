package server

import "github.com/SkNuwanTissera/gymshark/internal/validator"

func (s *Server) validateSizeOnValue(v *validator.Validator, size int) {
	v.Check(size > 0, "size", "size must be positive number")
}

func (s *Server) validateItemsOnValue(v *validator.Validator, size int) {
	v.Check(size > 0, "items", "items must be positive number")
}
