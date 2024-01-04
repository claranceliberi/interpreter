# interpreter
interpreter for monkey programming language

## Important Notes

- **Lexical Analysis**: The purpose of this phase is to read the input characters (source code) and produce as output a sequence of tokens that the parser uses for syntax analysis.
- **Syntax Analysis (Parsing)**: The purpose of this phase is to read the sequence of tokens produced by the lexical analyzer and produce as output a parse tree (or syntax tree AST) that the semantic analyzer uses for semantic analysis.

bellow is a picture that shows above phases:
![Compiler Phases](/assets/lexical_analysis.png)

- **REPL**: Read Eval Print Loop, is a simple interactive computer programming environment that takes single user inputs (i.e. single expressions), evaluates them, and returns the result to the user; a program written in a REPL environment is executed piecewise. (this is usually called console)
