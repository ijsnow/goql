package language

// Location contains a range of UTF-8 character offsets and token references that
// identify the region of the source from which the AST derived.
type Location struct {
	/**
	 * The character offset at which this Node begins.
	 */
	Start int `json:"start"`

	/**
	 * The character offset at which this Node ends.
	 */
	End int `json:"end"`

	/**
	 * The Token at which this Node begins.
	 */
	StartToken Token `json:"-"`

	/**
	 * The Token at which this Node ends.
	 */
	EndToken Token `json:"-"`

	/**
	 * The Source document the AST represents.
	 */
	Source Source `json:"-"`
}

// TODO: Decide what to do with the interfaces in this file

// ASTNode ...
type ASTNode interface {
	GetLoc() *Location
}

// NodeList is a list of ASTNodes with some helper func attached
type NodeList []ASTNode

// GetStarts filters a NodeList so that each item in the list
// has a Location and then builds a list of all the Start locations
func (n *NodeList) GetStarts() []int {
	var list NodeList
	for _, Node := range *n {
		if Node.GetLoc() != nil {
			list = append(list, Node)
		}
	}

	starts := make([]int, len(*n))
	for idx, s := range *n {
		starts[idx] = s.GetLoc().Start
	}
	return starts
}

type Node struct {
	Loc *Location
}

// GetLoc gets the location of a Node.
// It is primarily here to bind all of the Node
// types to the ASTNode interface
func (n Node) GetLoc() *Location {
	return n.Loc
}

// NameNode ...
type NameNode struct {
	Node
	Value string
}

// DocumentNode ...
type DocumentNode struct {
	Node
	Definitions []DefinitionNode
}

// TODO: Add methods that these need to implement

// DefinitionNode ...
type DefinitionNode interface{}

// export type DefinitionNode =
//   | OperationDefinitionNode
//   | FragmentDefinitionNode
//   | TypeSystemDefinitionNode; // experimental non-spec addition.

type OperationDefinitionNode struct {
	Node
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
	Node
	Variable     VariableNode
	Type         TypeNode
	DefaultValue *ValueNode
}

// VariableNode ...
type VariableNode struct {
	Node
	Name NameNode
}

// SelectionSetNode ...
type SelectionSetNode struct {
	Node
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
	Node
	Alias        *NameNode
	Name         NameNode
	Arguments    *[]ArgumentNode
	Directives   *[]DirectiveNode
	SelectionSet *SelectionSetNode
}

// ArgumentNode ...
type ArgumentNode struct {
	Node
	Name  NameNode
	Value ValueNode
}

// Fragments

// FragmentSpreadNode ...
type FragmentSpreadNode struct {
	Node
	Name       NameNode
	Directives *[]DirectiveNode
}

// InlineFragmentNode ...
type InlineFragmentNode struct {
	Node
	TypeCondition *NamedTypeNode
	Directives    *[]DirectiveNode
	SelectionSet  *SelectionSetNode
}

// FragmentDefinitionNode ...
type FragmentDefinitionNode struct {
	Node
	Name          NameNode
	TypeCondition NamedTypeNode
	Directives    *[]DirectiveNode
	SelectionSet  SelectionSetNode
}

// Values

// ValueNode ...
type ValueNode struct {
	Node
	Type  string
	Value interface{}
}

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
	Node
	Value string
}

// StringValueNode ...
type StringValueNode struct {
	Node
	Value string
}

// BooleanValueNode ...
type BooleanValueNode struct {
	Node
	Value bool
}

// NullValueNode ...
type NullValueNode struct {
	Loc *Location
}

// EnumValueNode ...
type EnumValueNode struct {
	Node
	Value string
}

// ListValueNode ...
type ListValueNode struct {
	Node
	Values []ASTNode
}

// ObjectValueNode ...
type ObjectValueNode struct {
	Node
	Fields []ObjectFieldNode
}

// ObjectFieldNode ...
type ObjectFieldNode struct {
	Node
	Name  NameNode
	Value ValueNode
}

// Directives

// DirectiveNode ...
type DirectiveNode struct {
	Node
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
	Node
	Name NameNode
}

// ListTypeNode ...
type ListTypeNode struct {
	Node
	Type TypeNode
}

// NonNullTypeNode ...
type NonNullTypeNode struct {
	Node
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
	Node
	Directives     []DirectiveNode
	OperationTypes []OperationTypeDefinitionNode
}

// OperationTypeDefinitionNode ...
type OperationTypeDefinitionNode struct {
	Node
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
	Node
	Name       NameNode
	Directives *[]DirectiveNode
}

// ObjectTypeDefinitionNode ...
type ObjectTypeDefinitionNode struct {
	Node
	Name       NameNode
	Interfaces *[]NamedTypeNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// FieldDefinitionNode ...
type FieldDefinitionNode struct {
	Node
	Name       NameNode
	Arguments  []InputValueDefinitionNode
	Type       TypeNode
	Directives *[]DirectiveNode
}

// InputValueDefinitionNode ...
type InputValueDefinitionNode struct {
	Node
	Name         NameNode
	Type         TypeNode
	DefaultValue *ValueNode
	Directives   *[]DirectiveNode
}

// InterfaceTypeDefinitionNode ...
type InterfaceTypeDefinitionNode struct {
	Node
	Name       NameNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// UnionTypeDefinitionNode ...
type UnionTypeDefinitionNode struct {
	Node
	Name       NameNode
	Directives *[]DirectiveNode
	Types      []NamedTypeNode
}

// EnumTypeDefinitionNode ...
type EnumTypeDefinitionNode struct {
	Node
	Name       NameNode
	Directives *[]DirectiveNode
	Values     []EnumValueDefinitionNode
}

// EnumValueDefinitionNode ...
type EnumValueDefinitionNode struct {
	Node
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
	Node
	Definition ObjectTypeDefinitionNode
}

// DirectiveDefinitionNode ...
type DirectiveDefinitionNode struct {
	Node
	Name      NameNode
	Arguments *[]InputValueDefinitionNode
	Locations []NameNode
}
