package language

// Name

const KindName = "Name"

// Document

const (
	KindDocument            = "Document"
	KindOperationDefinition = "OperationDefinition"
	KindVariableDefinition  = "VariableDefinition"
	KindVariable            = "Variable"
	KindSelectionSet        = "SelectionSet"
	KindField               = "Field"
	KindArgument            = "Argument"
)

// Fragments

const (
	KindFragmentSpread     = "FragmentSpread"
	KindInlineFragment     = "InlineFragment"
	KindFragmentDefinition = "FragmentDefinition"
)

// Values

const (
	KindInt         = "IntValue"
	KindFloat       = "FloatValue"
	FloatString     = "StringValue"
	KindBoolean     = "BooleanValue"
	KindNull        = "NullValue"
	KindEnum        = "EnumValue"
	KindList        = "ListValue"
	KindObject      = "ObjectValue"
	KindObjectField = "ObjectField"
)

// Directives

const KindDirective = "Directive"

// Types

const (
	KindNamedType   = "NamedType"
	KindListType    = "ListType"
	KindNonNullType = "NonNullType"
)

// Type System Definitions

const (
	KindSchemaDefinition        = "SchemaDefinition"
	KindOperationTypeDefinition = "OperationTypeDefinition"
)

// Type Definitions

const (
	KindScalarTypeDefinition     = "ScalarTypeDefinition"
	KindObjectTypeDefinition     = "ObjectTypeDefinition"
	KindFieldDefinition          = "FieldDefinition"
	KindInputValueDefinition     = "InputValueDefinition"
	KindInterfaceTypeDefinition  = "InterfaceTypeDefinition"
	KindUnionTypeDefinition      = "UnionTypeDefinition"
	KindEnumTypeDefinition       = "EnumTypeDefinition"
	KindEnumValueDefinition      = "EnumValueDefinition"
	KindInputObjectTypeDefintion = "InputObjectTypeDefinition"
)

// Type Extensions

const KindTypeExtensionDefinition = "TypeExtensionDefinition"

// Directive Definitions

const KindDirectiveDefinition = "DirectiveDefinition"
