package fsm

type transitionToBuilder struct {
	name             EventName
	action           ActionFunc
	transitionsSoFar map[StateKey]StateKey
	nextFrom         []StateKey
}

// To means the transition ends in the given state
func (t transitionToBuilder) To(to StateKey) EventBuilder {
	transitions := t.transitionsSoFar
	for _, from := range t.nextFrom {
		transitions[from] = to
	}
	return eventBuilder{t.name, t.action, transitions}
}

// ToNoChange means a transition ends in the same state it started in (just retriggers state cb)
func (t transitionToBuilder) ToNoChange() EventBuilder {
	transitions := t.transitionsSoFar
	for _, from := range t.nextFrom {
		transitions[from] = nil
	}
	return eventBuilder{t.name, t.action, transitions}
}

type eventBuilder struct {
	name             EventName
	action           ActionFunc
	transitionsSoFar map[StateKey]StateKey
}

// From begins describing a transition from a specific state
func (t eventBuilder) From(s StateKey) TransitionToBuilder {
	return transitionToBuilder{
		t.name,
		t.action,
		t.transitionsSoFar,
		[]StateKey{s},
	}
}

// FromAny begins describing a transition from any state
func (t eventBuilder) FromAny() TransitionToBuilder {
	return transitionToBuilder{
		t.name,
		t.action,
		t.transitionsSoFar,
		[]StateKey{nil},
	}
}

// FromMany begins describing a transition from many states
func (t eventBuilder) FromMany(sources ...StateKey) TransitionToBuilder {
	return transitionToBuilder{
		t.name,
		t.action,
		t.transitionsSoFar,
		sources,
	}
}

// Action describes actions taken on the state for this event
func (t eventBuilder) Action(action ActionFunc) EventBuilder {
	return eventBuilder{
		t.name,
		action,
		t.transitionsSoFar,
	}
}

// Event starts building a new event
func Event(name EventName) EventBuilder {
	return eventBuilder{name, nil, map[StateKey]StateKey{}}
}
