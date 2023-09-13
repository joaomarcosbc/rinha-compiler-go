package main

type Env struct {
	Objects map[string]Value
}

type Closure struct {
	Body       Term
	Parameters []string
	Env        Env
}

type Value struct {
	Kind  string
	Value interface{}
}
