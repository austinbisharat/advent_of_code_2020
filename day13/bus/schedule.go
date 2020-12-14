package bus

type Schedule struct {
	ID  int
	Idx int
}

func (s Schedule) TimeUntilNextBus(curTime int) int {
	return s.ID - curTime%s.ID
}
