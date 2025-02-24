%{
package main
import "fmt"

%}

%union {
    Int int
    Str string
}

%token INT64LIT
%token IDENT
%token FN
%token LPAREN RPAREN LBRACE RBRACE
%token SCOLON
%token ASSIGN
%token PLUS MIN MUL DIV
%token IF ELSE
%token EQ NE GT LT GTE LTE

%%

// Программа состоит из списка операторов
Program: StmtList ;

// Список операторов
StmtList: Stmt | StmtList Stmt | ;

// Оператор
Stmt: FnDecl
    | VarDecl
    | IfStmt
    | ExprStmt;

ExprStmt: Expr SCOLON;

// Объявление функции
FnDecl: FN IDENT LPAREN FormalParmLstOpt RPAREN LBRACE StmtList RBRACE;

// Список параметров (опционально)
FormalParmLstOpt: FormalParmLst | ;

// Список параметров
FormalParmLst: FormalParm | FormalParmLst ',' FormalParm;

// Параметр
FormalParm: IDENT;

// Объявление переменной
VarDecl: IDENT ASSIGN Expr SCOLON;

// Условный оператор
IfStmt: IF '(' Expr ')' LBRACE StmtList RBRACE ElseClause;

// Альтернативный оператор
ElseClause: ELSE LBRACE StmtList RBRACE | ;

// Выражение
Expr: Term
    | Expr PLUS Term
    | Expr MIN Term;

// Член
Term: Factor
    | Term MUL Factor
    | Term DIV Factor;

// Фактор
Factor: INT64LIT
      | IDENT
      | LPAREN Expr RPAREN;
