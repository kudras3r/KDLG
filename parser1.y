%{
package main
%}

%union{
Str string
Int int
Boo bool
Flt float64
}

%token BYTELIT INT16LIT INT32LIT INT64LIT FLOAT64LIT BOOLLIT STRLIT
%token IDENT BYTE INT16 INT32 INT64 FLOAT64 BOOL STR
%token MAIN FN RET IF ELSE FOR BRK NXT XOR CLASS ETERN
%token INCR DECR EQ NE GT GTE LT LTE AND OR
%token ASSIGN
%token SCOLON

%%

FieldDecl: Type VarDecls SCOLON ;
Type: BYTE | INT16 | INT32 | INT64 | FLOAT64 | BOOL | STR | Name ;
Name: IDENT | QualifiedName ;

QualifiedName: Name '.' IDENT ;
VarDecls: VarDecl | VarDecls ',' VarDecl ;
VarDecl: IDENT | VarDecl '[' ']' ;

FnDecl: FN IDENT '(' FormalParmLstOpt ')' Block ;
FormalParmLstOpt: FormalParm | FormalParmLst ',' FormalParm ;
FormalParmLst: FormalParm | FormalParmLst ',' FormalParm ;
FormalParm: Type VarDecl ;

Block: '{' BlockStmtsOpt '}' ;
BlockStmtsOpt: BlockStmts | ;
BlockStmts: BlockStmt | BlockStmts BlockStmt ;
BlockStmt: LocalVarDeclStmt | Stmt ;
LocalVarDeclStmt: LocalVarDecl ';' ;
LocalVarDecl: Type VarDecls ;

Stmt: Block | ';' | ExprStmt | BreakStmt | RetStmt
     | IfThenStmt | IfThenElseStmt | IfThenElseIfStmt 
     | ForStmt ;

ExprStmt: StmtExpr ';' ;
StmtExpr: Assignment | MethodCall | InstantiationExpr ;

IfThenStmt: IF '(' Expr ')' Block ;
IfThenElseStmt: IF '(' Expr ')' Block ELSE Block ;
IfThenElseIfStmt: IF '(' Expr ')' Block ElseIfSequence
     | IF '(' Expr ')' Block ElseIfSequence ELSE Block ;
ElseIfSequence: ElseIfStmt | ElseIfSequence ElseIfStmt ;
ElseIfStmt: ELSE IfThenStmt ;

ForStmt: FOR '(' ForInit ';' ExprOpt ';' ForUpdate ')' Block ;
ForInit: StmtExprLst | LocalVarDecl | ;
ExprOpt: Expr | ;
ForUpdate: StmtExprLst | ;
StmtExprLst: StmtExpr | StmtExprLst ',' StmtExpr ;

BreakStmt: BRK ';' ;
RetStmt: RET ExprOpt ';' ;

Primary:  Literal | '(' Expr ')' | FieldAccess | MethodCall ;
Literal: INT64LIT	| FLOAT64LIT | BOOLLIT | STRLIT ;

InstantiationExpr: Name '(' ArgLstOpt ')' ;
ArgLst: Expr | ArgLst ',' Expr ;
ArgLstOpt: ArgLst | ;
FieldAccess: Primary '.' IDENT ;

MethodCall: Name '(' ArgLstOpt ')'
	| Name '{' ArgLstOpt '}'
	| Primary '.' IDENT '(' ArgLstOpt ')'
	| Primary '.' IDENT '{' ArgLstOpt '}' ;

PostFixExpr: Primary | Name ;
UnaryExpr:  '-' UnaryExpr | '!' UnaryExpr | PostFixExpr ;
MulExpr: UnaryExpr | MulExpr '*' UnaryExpr
    | MulExpr '/' UnaryExpr | MulExpr '%' UnaryExpr ;
AddExpr: MulExpr | AddExpr '+' MulExpr | AddExpr '-' MulExpr ;
RelOp: LTE | GTE | '<' | '>' ;
RelExpr: AddExpr | RelExpr RelOp AddExpr ;

EqExpr: RelExpr | EqExpr EQ RelExpr | EqExpr NE RelExpr ;
CondAndExpr: EqExpr | CondAndExpr AND EqExpr ;
CondOrExpr: CondAndExpr | CondOrExpr OR CondAndExpr ;

Expr: CondOrExpr | Assignment ;
Assignment: LeftHandSide AssignOp Expr ;
LeftHandSide: Name | FieldAccess ;
AssignOp: ASSIGN | INCR | DECR ;


