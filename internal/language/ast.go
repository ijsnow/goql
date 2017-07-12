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
	for _, node := range *n {
		if node.GetLoc() != nil {
			list = append(list, node)
		}
	}

	starts := make([]int, len(*n))
	for idx, s := range *n {
		starts[idx] = s.GetLoc().Start
	}
	return starts
}

type node struct {
	Loc *Location
}

// GetLoc gets the location of a node.
// It is primarily here to bind all of the node
// types to the ASTNode interface
func (n node) GetLoc() *Location {
	return n.Loc
}

// NameNode ...
type NameNode struct {
	node
	Value string
}

// DocumentNode ...
type DocumentNode struct {
	node
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
	node
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
	node
	Variable     VariableNode
	Type         TypeNode
	DefaultValue *ValueNode
}

// VariableNode ...
type VariableNode struct {
	node
	Name NameNode
}

// SelectionSetNode ...
type SelectionSetNode struct {
	node
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
	node
	Alias        *NameNode
	Name         NameNode
	Arguments    *[]ArgumentNode
	Directives   *[]DirectiveNode
	SelectionSet *SelectionSetNode
}

// ArgumentNode ...
type ArgumentNode struct {
	node
	Name  NameNode
	Value ValueNode
}

// Fragments

// FragmentSpreadNode ...
type FragmentSpreadNode struct {
	node
	Name       NameNode
	Directives *[]DirectiveNode
}

// InlineFragmentNode ...
type InlineFragmentNode struct {
	node
	TypeCondition *NamedTypeNode
	Directives    *[]DirectiveNode
	SelectionSet  *SelectionSetNode
}

// FragmentDefinitionNode ...
type FragmentDefinitionNode struct {
	node
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
	node
	Value string
}

// StringValueNode ...
type StringValueNode struct {
	node
	Value string
}

// BooleanValueNode ...
type BooleanValueNode struct {
	node
	Value bool
}

// NullValueNode ...
type NullValueNode struct {
	Loc *Location
}

// EnumValueNode ...
type EnumValueNode struct {
	node
	Value string
}

// ListValueNode ...
type ListValueNode struct {
	node
	Values []ValueNode
}

// ObjectValueNode ...
type ObjectValueNode struct {
	node
	Fields []ObjectFieldNode
}

// ObjectFieldNode ...
type ObjectFieldNode struct {
	node
	Name  NameNode
	Value ValueNode
}

// Directives

// DirectiveNode ...
type DirectiveNode struct {
	node
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
	node
	Name NameNode
}

// ListTypeNode ...
type ListTypeNode struct {
	node
	Type TypeNode
}

// NonNullTypeNode ...
type NonNullTypeNode struct {
	node
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
	node
	Directives     []DirectiveNode
	OperationTypes []OperationTypeDefinitionNode
}

// OperationTypeDefinitionNode ...
type OperationTypeDefinitionNode struct {
	node
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
	node
	Name       NameNode
	Directives *[]DirectiveNode
}

// ObjectTypeDefinitionNode ...
type ObjectTypeDefinitionNode struct {
	node
	Name       NameNode
	Interfaces *[]NamedTypeNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// FieldDefinitionNode ...
type FieldDefinitionNode struct {
	node
	Name       NameNode
	Arguments  []InputValueDefinitionNode
	Type       TypeNode
	Directives *[]DirectiveNode
}

// InputValueDefinitionNode ...
type InputValueDefinitionNode struct {
	node
	Name         NameNode
	Type         TypeNode
	DefaultValue *ValueNode
	Directives   *[]DirectiveNode
}

// InterfaceTypeDefinitionNode ...
type InterfaceTypeDefinitionNode struct {
	node
	Name       NameNode
	Directives *[]DirectiveNode
	Fields     []FieldDefinitionNode
}

// UnionTypeDefinitionNode ...
type UnionTypeDefinitionNode struct {
	node
	Name       NameNode
	Directives *[]DirectiveNode
	Types      []NamedTypeNode
}

// EnumTypeDefinitionNode ...
type EnumTypeDefinitionNode struct {
	node
	Name       NameNode
	Directives *[]DirectiveNode
	Values     []EnumValueDefinitionNode
}

// EnumValueDefinitionNode ...
type EnumValueDefinitionNode struct {
	node
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
	node
	Definition ObjectTypeDefinitionNode
}

// DirectiveDefinitionNode ...
type DirectiveDefinitionNode struct {
	node
	Name      NameNode
	Arguments *[]InputValueDefinitionNode
	Locations []NameNode
}
