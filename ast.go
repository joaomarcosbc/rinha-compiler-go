package main

type Loc struct {
	Start    int
	End      int
	Filename string
}

type BinaryOp string

const (
	Concat BinaryOp = "Concat"
	Add    BinaryOp = "Add"
	Sub    BinaryOp = "Sub"
	Mul    BinaryOp = "Mul"
	Div    BinaryOp = "Div"
	Rem    BinaryOp = "Rem"
	Eq     BinaryOp = "Eq"
	Neq    BinaryOp = "Neq"
	Lt     BinaryOp = "Lt"
	Gt     BinaryOp = "Gt"
	Lte    BinaryOp = "Lte"
	Gte    BinaryOp = "Gte"
	And    BinaryOp = "And"
	Or     BinaryOp = "Or"
)

type Parameter struct {
	Text     string
	Location Loc
}

type Term interface{}

type Bool struct {
	Kind     string
	Value    bool
	Location Loc
}

type Int struct {
	Kind     string
	Value    int
	Location Loc
}

type String struct {
	Kind     string
	Value    string
	Location Loc
}

type Var struct {
	Kind     string
	Text     string
	Location Loc
}

type Function struct {
	Kind       string
	Parameters []Parameter
	Value      Term
	Location   Loc
}

type Call struct {
	Kind      string
	Callee    Term
	Arguments []Term
	Location  Loc
}

type Binary struct {
	Kind     string
	Lhs      Term
	Op       BinaryOp
	Rhs      Term
	Location Loc
}

type Let struct {
	Kind     string
	Name     Parameter
	Value    Term
	Next     Term
	Location Loc
}

type If struct {
	Kind      string
	Condition Term
	Then      Term
	Otherwise Term
	Location  Loc
}

type Print struct {
	Kind     string
	Value    Term
	Location Loc
}

type Program struct {
	Name       string
	Expression Term
	Location   Loc
}
