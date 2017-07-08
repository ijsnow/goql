package language

// Location contains a range of UTF-8 character offsets and token references that
// identify the region of the source from which the AST derived.
type Location struct {
	/**
	 * The character offset at which this Node begins.
	 */
	Start int

	/**
	 * The character offset at which this Node ends.
	 */
	End int

	/**
	 * The Token at which this Node begins.
	 */
	StartToken Token

	/**
	 * The Token at which this Node ends.
	 */
	EndToken Token

	/**
	 * The Source document the AST represents.
	 */
	Source Source
}

// TokenKind represents the different kinds of tokens in a GraphQL document.
type TokenKind string

const (
	TokenSOF = iota
	TokenEOF
	TokenBang
	TokenDollar
	TokenOpenParen
	TokenCloseParen
	TokenEllipsis
	TokenColon
	TokenEqual
	TokenAt
	TokenOpenSquare
	TokenCloseSquare
	TokenOpenCurly
	TokenCloseCurly
	TokenBar
	TokenName
	TokenInt
	TokenFloat
	TokenString
	TokenComment
)

// Tokens holds all of the special tokens in the GraphQL language spec
var Tokens = map[string]int{
	"<SOF>":   TokenSOF,
	"<EOF>":   TokenEOF,
	"!":       TokenBang,
	"$":       TokenDollar,
	"(":       TokenOpenParen,
	")":       TokenCloseParen,
	"...":     TokenEllipsis,
	":":       TokenColon,
	"=":       TokenEqual,
	"@":       TokenAt,
	"[":       TokenOpenSquare,
	"]":       TokenCloseSquare,
	"{":       TokenOpenCurly,
	"}":       TokenCloseCurly,
	"|":       TokenBar,
	"Name":    TokenName,
	"Int":     TokenInt,
	"Float":   TokenFloat,
	"String":  TokenString,
	"Comment": TokenComment,
}

// TODO: make map of tokens that point to the TokenKind

// TODO: maybe make a func FindTokenKind that takes the token then finds the TokenKind in above map

// Token represents a range of characters represented by a lexical token
// within a Source.
type Token struct {

	/**
	 * The kind of Token.
	 */
	Kind TokenKind

	/**
	 * The character offset at which this Node begins.
	 */
	Start int

	/**
	 * The character offset at which this Node ends.
	 */
	End int

	/**
	 * The 1-indexed line number on which this Token appears.
	 */
	Line int

	/**
	 * The 1-indexed column number at which this Token begins.
	 */
	Column int

	/**
	 * For non-punctuation tokens, represents the interpreted value of the token.
	 */
	Value string

	/**
	 * Tokens exist as nodes in a double-linked-list amongst all tokens
	 * including ignored tokens. <SOF> is always the first node and <EOF>
	 * the last.
	 */
	Prev *Token
	Next *Token
}

// TODO: Decide what to do with the interfaces in this file

// ASTNode ...
type ASTNode interface{}

/**
 * ASTNodes is the list of all possible AST node types.
 */
// var ASTNodes = []ASTNode{
// NameNode,
// DocumentNode,
// OperationDefinitionNode,
// VariableDefinitionNode,
// VariableNode,
// SelectionSetNode,
// FieldNode,
// ArgumentNode,
// FragmentSpreadNode,
// InlineFragmentNode,
// FragmentDefinitionNode,
// IntValueNode,
// FloatValueNode,
// StringValueNode,
// BooleanValueNode,
// NullValueNode,
// EnumValueNode,
// ListValueNode,
// ObjectValueNode,
// ObjectFieldNode,
// DirectiveNode,
// NamedTypeNode,
// ListTypeNode,
// NonNullTypeNode,
// SchemaDefinitionNode,
// OperationTypeDefinitionNode,
// ScalarTypeDefinitionNode,
// ObjectTypeDefinitionNode,
// FieldDefinitionNode,
// InputValueDefinitionNode,
// InterfaceTypeDefinitionNode,
// UnionTypeDefinitionNode,
// EnumTypeDefinitionNode,
// EnumValueDefinitionNode,
// InputObjectTypeDefinitionNode,
// TypeExtensionDefinitionNode,
// DirectiveDefinitionNode,
// }

// NameNode ...
type NameNode struct {
	Loc   *Location
	Value string
}

// DocumentNode ...
type DocumentNode struct {
	Loc        *Location
	Definition []DefinitionNode
}

// TODO: Add methods that these need to implement

// DefinitionNode ...
type DefinitionNode interface{}

// export type DefinitionNode =
//   | OperationDefinitionNode
//   | FragmentDefinitionNode
//   | TypeSystemDefinitionNode; // experimental non-spec addition.

type OperationDefinitionNode struct {
	Loc                 *Location
	Operation           OperationTypeNode
	Name                NameNode
	VariableDefinitions *[]VariableDefinitionNode
	Directives          *[]DirectiveNode
	SelectionSet        SelectionSetNode
}

// OperationTypeNode ...
type OperationTypeNode string

const (
	OperationTypeQuery        = "query"
	OperationTypeMutation     = "mutation"
	OperationTypeSubscription = "subscription"
)

// VariableDefinitionNode ...
type VariableDefinitionNode struct {
	Loc          *Location
	Variable     VariableNode
	Type         TypeNode
	DefaultValue *ValueNode
}

// VariableNode ...
type VariableNode struct {
	Loc  *Location
	Name NameNode
}

// SelectionSetNode ...
type SelectionSetNode struct {
	Loc        *Location
	Selections []SelectionNode
}

// TODO: Add methods that these need to implement

// SelectionNode ...
type SelectionNode interface{}

// export type SelectionNode =
//   | FieldNode
//   | FragmentSpreadNode
//   | InlineFragmentNode;

// FieldNode ...
type FieldNode struct {
	Loc          *Location
	Alias        *NameNode
	Name         NameNode
	Arguments    *[]ArgumentNode
	Directives   *[]DirectiveNode
	SelectionSet *SelectionSetNode
}

// ArgumentNode ...
type ArgumentNode struct {
	Loc   *Location
	Name  NameNode
	Value ValueNode
}

// Fragments

// FragmentSpreadNode ...
type FragmentSpreadNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
}

// InlineFragmentNode ...
type InlineFragmentNode struct {
	Loc           *Location
	TypeCondition *NamedTypeNode
	Directives    *[]DirectiveNode
	SelectionSet  *SelectionSetNode
}

// FragmentDefinitionNode ...
type FragmentDefinitionNode struct {
	Loc           *Location
	Name          NameNode
	TypeCondition NamedTypeNode
	Directives    *[]DirectiveNode
	SelectionSet  SelectionSetNode
}

// Values

// ValueNode ...
type ValueNode interface{}

// export type ValueNode =
//   | VariableNode
//   | IntValueNode
//   | FloatValueNode
//   | StringValueNode
//   | BooleanValueNode
//   | NullValueNode
//   | EnumValueNode
//   | ListValueNode
//   | ObjectValueNode;

// IntValueNode ...
type IntValueNode struct {
	Loc   Location
	Value string
}

// FloatValueNode ...
type FloatValueNode struct {
	Loc   *Location
	Value string
}

// StringValueNode ...
type StringValueNode struct {
	Loc   *Location
	Value string
}

// BooleanValueNode ...
type BooleanValueNode struct {
	Loc   *Location
	Value bool
}

// NullValueNode ...
type NullValueNode struct {
	Loc *Location
}

// EnumValueNode ...
type EnumValueNode struct {
	Loc   *Location
	Value string
}

// ListValueNode ...
type ListValueNode struct {
	Loc    *Location
	Values []ValueNode
}

// ObjectValueNode ...
type ObjectValueNode struct {
	Loc    *Location
	Fields []ObjectFieldNode
}

// ObjectFieldNode ...
type ObjectFieldNode struct {
	Loc   *Location
	Name  NameNode
	Value ValueNode
}

// Directives

// DirectiveNode ...
type DirectiveNode struct {
	Loc       *Location
	Name      NameNode
	Arguments *[]ArgumentNode
}

// Type Reference

// TypeNode ...
type TypeNode interface{}

// export type TypeNode =
//   | NamedTypeNode
//   | ListTypeNode
//   | NonNullTypeNode;

// NamedTypeNode ...
type NamedTypeNode struct {
	Loc  *Location
	Name NameNode
}

// ListTypeNode ...
type ListTypeNode struct {
	Loc  *Location
	Type TypeNode
}

// NonNullTypeNode ...
type NonNullTypeNode struct {
	Loc  *Location
	Type TypeNode
}

// Type System Definition

// TypeSystemDefinitionNode ...
type TypeSystemDefinitionNode interface{}

// export type TypeSystemDefinitionNode =
//   | SchemaDefinitionNode
//   | TypeDefinitionNode
//   | TypeExtensionDefinitionNode
//   | DirectiveDefinitionNode;

// SchemaDefinitionNode ...
type SchemaDefinitionNode struct {
	Loc            *Location
	Directives     []DirectiveNode
	OperationTypes []OperationTypeDefinitionNode
}

// OperationTypeDefinitionNode ...
type OperationTypeDefinitionNode struct {
	Loc       *Location
	Operation OperationTypeNode
	Type      NamedTypeNode
}

// TypeDefinitionNode ...
type TypeDefinitionNode interface{}

// export type TypeDefinitionNode =
//   | ScalarTypeDefinitionNode
//   | ObjectTypeDefinitionNode
//   | InterfaceTypeDefinitionNode
//   | UnionTypeDefinitionNode
//   | EnumTypeDefinitionNode
//   | InputObjectTypeDefinitionNode;

// ScalarTypeDefinitionNode ...
type ScalarTypeDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
}

// ObjectTypeDefinitionNode ...
type ObjectTypeDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Interfaces *[]NamedTypeNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// FieldDefinitionNode ...
type FieldDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Arguments  []InputValueDefinitionNode
	Type       TypeNode
	Directives *[]DirectiveNode
}

// InputValueDefinitionNode ...
type InputValueDefinitionNode struct {
	Loc          *Location
	Name         NameNode
	Type         TypeNode
	DefaultValue *ValueNode
	Directives   *[]DirectiveNode
}

// InterfaceTypeDefinitionNode ...
type InterfaceTypeDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// UnionTypeDefinitionNode ...
type UnionTypeDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
	Types      []NamedTypeNode
}

// EnumTypeDefinitionNode ...
type EnumTypeDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
	Values     []EnumValueDefinitionNode
}

// EnumValueDefinitionNode ...
type EnumValueDefinitionNode struct {
	Loc        *Location
	Name       NameNode
	Directives *[]DirectiveNode
}

// InputObjectTypeDefinitionNode ...
type InputObjectTypeDefinitionNode struct {
	Loc        Location
	Name       NameNode
	Directives *[]DirectiveNode
	Fields     []InputValueDefinitionNode
}

// TypeExtensionDefinitionNode ...
type TypeExtensionDefinitionNode struct {
	Loc        *Location
	Definition ObjectTypeDefinitionNode
}

// DirectiveDefinitionNode ...
type DirectiveDefinitionNode struct {
	Loc       *Location
	Name      NameNode
	Arguments *[]InputValueDefinitionNode
	Locations []NameNode
}
