package gojq

type env struct {
	funcDefs  map[string]map[int]*FuncDef
	variables map[string]*Pipe
}

func newEnv() *env {
	return &env{
		funcDefs:  make(map[string]map[int]*FuncDef),
		variables: make(map[string]*Pipe),
	}
}

func (env *env) addFuncDef(fd *FuncDef) {
	if _, ok := env.funcDefs[fd.Name]; !ok {
		env.funcDefs[fd.Name] = make(map[int]*FuncDef)
	}
	env.funcDefs[fd.Name][len(fd.Args)] = fd
}
