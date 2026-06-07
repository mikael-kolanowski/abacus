package engine

type EvalResult struct {
	Value float64
}

type Engine struct {
}

func (e *Engine) Eval(source string) (EvalResult, error) {
	return EvalResult{Value: 0}, nil
}

func New() *Engine {
	return &Engine{}
}
