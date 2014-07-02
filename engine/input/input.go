package input

import (
	"time"
)

type KeyState struct {
	State, Begin, End bool
}

func (k *KeyState) Update(next bool) {
	k.Begin = !k.State && next
	k.End = k.State && !next
	k.State = next
}

type Input struct {
	Running bool

	Action KeyState

	North KeyState
	South KeyState
	East  KeyState
	West  KeyState

	TimeStamp time.Time
	DeltaTime time.Duration
}
