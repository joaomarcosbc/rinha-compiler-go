package main

type Loc struct {
	Start    int
	End      int
	Filename string
}

type Binary string

const (
	Concat Binary = "Concat"
	Add    Binary = "Add"
	Sub    Binary = "Sub"
	Mul    Binary = "Mul"
	Div    Binary = "Div"
	Rem    Binary = "Rem"
	Eq     Binary = "Eq"
	Neq    Binary = "Neq"
	Lt     Binary = "Lt"
	Gt     Binary = "Gt"
	Lte    Binary = "Lte"
	Gte    Binary = "Gte"
	And    Binary = "And"
	Or     Binary = "Or"
)

type Parameter struct {
	Text     string
	Location Loc
}

type Term interface{}

type BoolTerm struct {
	Kind     string
	Value    bool
	Location Loc
}

type IntTerm struct {
	Kind     string
	Value    int
	Location Loc
}

type StringTerm struct {
	Kind     string
	Value    string
	Location Loc
}

type VarTerm struct {
	Kind     string
	Text     string
	Location Loc
}

type FunctionTerm struct {
	Kind       string
	Parameters []Parameter
	Value      Term
	Location   Loc
}

type CallTerm struct {
	Kind      string
	Callee    Term
	Arguments []Term
	Location  Loc
}

type BinaryTerm struct {
	Kind     string
	Lhs      Term
	Op       Binary
	Rhs      Term
	Location Loc
}

type LetTerm struct {
	Kind     string
	Name     Parameter
	Value    Term
	Next     Term
	Location Loc
}

type IfTerm struct {
	Kind      string
	Condition Term
	Then      Term
	Otherwise Term
	Location  Loc
}

type PrintTerm struct {
	Kind     string
	Value    Term
	Location Loc
}

type Program struct {
	Name       string
	Expression Term
	Location   Loc
}
