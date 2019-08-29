package utils

type CircuitBreaker struct {
	TimeOut int
	MaxRetries int
	CountRetries int
	State int
}

var(
	CB *CircuitBreaker
)

const (
	OPEN = 0
	HALFOPEN = 1
	CLOUSE = 2
)

func NewCircuitBreaker(timeOut int, maxRetries int,countRetries int,state int) {

	CB = &CircuitBreaker{
		TimeOut:      timeOut,
		MaxRetries:   maxRetries,
		CountRetries: countRetries,
		State: 		  state,
	}
}

func (c *CircuitBreaker) ResetState(){
	c.CountRetries = 0
}

func (c *CircuitBreaker) SetState(state int){
	c.State = state
}

func (c *CircuitBreaker) AddCountRetries(){
	c.CountRetries++
}