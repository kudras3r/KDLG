package main

import (
	"fmt"
)

func main() {
	input := "func asd(){x=5;}"
	lexer := New(input)

	// for {
	// 	token := lexer.Lex(&yySymType{})
	// 	if token == 0 {
	// 		fmt.Println(1)
	// 		break
	// 	}
	// 	fmt.Println(token)
	// }
	p := yyNewParser()

	result := p.Parse(lexer)
	fmt.Println("Parsed result:", result)
}

// FormalParmLst: FormalParm | FormalParmLst ',' FormalParm ;
// FormalParm: Type VarDecl ;

// Block: '{' BlockStmtsOpt '}' ;
// BlockStmtsOpt: BlockStmts | ;
// BlockStmts: BlockStmt | BlockStmts BlockStmt ;
// BlockStmt: LocalVarDeclStmt | Stmt ;
// LocalVarDeclStmt: LocalVarDecl SCOLON ;
// LocalVarDecl: Type VarDecls ;

// Stmt: Block | SCOLON | ExprStmt | BreakStmt | RetStmt
//      | IfThenStmt | IfThenElseStmt | IfThenElseIfStmt
//      | ForStmt ;

// ExprStmt: StmtExpr SCOLON ;
// StmtExpr: Assignment | MethodCall | InstantiationExpr ;

// IfThenStmt: IF '(' Expr ')' Block ;
// IfThenElseStmt: IF '(' Expr ')' Block ELSE Block ;
// IfThenElseIfStmt: IF '(' Expr ')' Block ElseIfSequence
//      | IF '(' Expr ')' Block ElseIfSequence ELSE Block ;
// ElseIfSequence: ElseIfStmt | ElseIfSequence ElseIfStmt ;
// ElseIfStmt: ELSE IfThenStmt ;

// ForStmt: FOR '(' ForInit SCOLON ExprOpt SCOLON ForUpdate ')' Block ;
// ForInit: StmtExprLst | LocalVarDecl | ;
// ExprOpt: Expr | ;
// ForUpdate: StmtExprLst | ;
// StmtExprLst: StmtExpr | StmtExprLst ',' StmtExpr ;

// BreakStmt: BRK SCOLON ;
// RetStmt: RET ExprOpt SCOLON ;

// Primary:  Literal | '(' Expr ')' | FieldAccess | MethodCall ;
// Literal: INT64LIT	| FLOAT64LIT | BOOLLIT | STRLIT ;

// InstantiationExpr: Name '(' ArgLstOpt ')' ;
// ArgLst: Expr | ArgLst ',' Expr ;
// ArgLstOpt: ArgLst | ;
// FieldAccess: Primary '.' IDENT ;

// MethodCall: Name '(' ArgLstOpt ')'
// 	| Name '{' ArgLstOpt '}'
// 	| Primary '.' IDENT '(' ArgLstOpt ')'
// 	| Primary '.' IDENT '{' ArgLstOpt '}' ;

// PostFixExpr: Primary | Name ;
// UnaryExpr:  '-' UnaryExpr | '!' UnaryExpr | PostFixExpr ;
// MulExpr: UnaryExpr | MulExpr '*' UnaryExpr
//     | MulExpr '/' UnaryExpr | MulExpr '%' UnaryExpr ;
// AddExpr: MulExpr | AddExpr '+' MulExpr | AddExpr '-' MulExpr ;
// RelOp: LTE | GTE | '<' | '>' ;
// RelExpr: AddExpr | RelExpr RelOp AddExpr ;

// EqExpr: RelExpr | EqExpr EQ RelExpr | EqExpr NE RelExpr ;
// CondAndExpr: EqExpr | CondAndExpr AND EqExpr ;
// CondOrExpr: CondAndExpr | CondOrExpr OR CondAndExpr ;

// Expr: CondOrExpr | Assignment ;
// Assignment: LeftHandSide AssignOp Expr ;
// LeftHandSide: Name | FieldAccess ;
// AssignOp: ASSIGN | INCR | DECR ;
