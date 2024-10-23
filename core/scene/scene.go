package scene

import "fmt"

type Scene struct {
	Name string
}

func (s *Scene) ToString() {
	fmt.Printf("SCENE NAME: %s", s.Name)
}
