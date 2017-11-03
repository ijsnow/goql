package query

import (
	"fmt"

	"github.com/ijsnow/goql/internal/errors"
	"github.com/ijsnow/goql/internal/language"
)

type parser func(*Lexer) (language.ASTNode, error)

// Parse , Given a GraphQL source, parses it into a Document.
// Returns a GraphQLError if a syntax error is encountered.
// func Parse(src language.Source, options ...ParseOptions) (language.DocumentNode, error) {
// 	lexer := CreateLexer(src, options...)

// 	return parseDocument(lexer)
// }

/**
 * Given a string containing a GraphQL value (ex. `[42]`), parse the AST for
 * that value.
 * Throws GraphQLError if a syntax error is encountered.
 *
 * This is useful within tools that operate upon GraphQL Values directly and
 * in isolation of complete GraphQL documents.
 *
 * Consider providing the results to the utility function: valueFromAST().
 */
// func ParseValue(
// 	source language.Source,
// 	options ...ParseOptions,
// ) (*language.ASTNode, error) {
// 	lexer := CreateLexer(source, options...)

// 	_, err := expect(lexer, language.TokenSOF)
// 	if err != nil {
// 		return nil, err
// 	}

// 	value, err := parseValueLiteral(lexer, false)

// 	_, err = expect(lexer, language.TokenEOF)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return value, nil
// }

// /**
//  * Given a string containing a GraphQL Type (ex. `[Int!]`), parse the AST for
//  * that type.
//  * Throws GraphQLError if a syntax error is encountered.
//  *
//  * This is useful within tools that operate upon GraphQL Types directly and
//  * in isolation of complete GraphQL documents.
//  *
//  * Consider providing the results to the utility function: typeFromAST().
//  */
// export function parseType(
//   source: string | Source,
//   options?: ParseOptions
// ): TypeNode {
//   const sourceObj = typeof source === 'string' ? new Source(source) : source;
//   const lexer = createLexer(sourceObj, options || {});
//   expect(lexer, TokenKind.SOF);
//   const type = parseTypeReference(lexer);
//   expect(lexer, TokenKind.EOF);
//   return type;
// }

// /**
//  * Converts a name lex token into a name parse node.
//  */
// function parseName(lexer: Lexer<*>): NameNode {
//   const token = expect(lexer, TokenKind.NAME);
//   return {
//     kind: NAME,
//     value: ((token.value: any): string),
//     loc: loc(lexer, token)
//   };
// }

// Implements the parsing rules in the Document section.

/**
 * Document : Definition+
 */
// func parseDocument(lexer Lexer) (language.DocumentNode, error) {
// 	start := lexer.Token
// 	_, err := expect(lexer, language.TokenSOF)
// 	if err != nil {
// 		return nil, err
// 	}

// 	  definitions := []language.DefinitionNode

// 	  for {
// 		definitions = append(parseDefinition(lexer))

// 		if skip(lexer, TokenKind.EOF) {
// 			break
// 		}
// 	  }

// 	  return language.DocumentNode{
// 	    Definitions: definitions,
// 	    Loc: loc(lexer, start),
// 	  }
// }

// /**
//  * Definition :
//  *   - OperationDefinition
//  *   - FragmentDefinition
//  *   - TypeSystemDefinition
//  */
// function parseDefinition(lexer: Lexer<*>): DefinitionNode {
//   if (peek(lexer, TokenKind.BRACE_L)) {
//     return parseOperationDefinition(lexer);
//   }

//   if (peek(lexer, TokenKind.NAME)) {
//     switch (lexer.token.value) {
//       // Note: subscription is an experimental non-spec addition.
//       case 'query':
//       case 'mutation':
//       case 'subscription':
//         return parseOperationDefinition(lexer);

//       case 'fragment': return parseFragmentDefinition(lexer);

//       // Note: the Type System IDL is an experimental non-spec addition.
//       case 'schema':
//       case 'scalar':
//       case 'type':
//       case 'interface':
//       case 'union':
//       case 'enum':
//       case 'input':
//       case 'extend':
//       case 'directive': return parseTypeSystemDefinition(lexer);
//     }
//   }

//   throw unexpected(lexer);
// }

// // Implements the parsing rules in the Operations section.

// /**
//  * OperationDefinition :
//  *  - SelectionSet
//  *  - OperationType Name? VariableDefinitions? Directives? SelectionSet
//  */
// function parseOperationDefinition(lexer: Lexer<*>): OperationDefinitionNode {
//   const start = lexer.token;
//   if (peek(lexer, TokenKind.BRACE_L)) {
//     return {
//       kind: OPERATION_DEFINITION,
//       operation: 'query',
//       name: null,
//       variableDefinitions: null,
//       directives: [],
//       selectionSet: parseSelectionSet(lexer),
//       loc: loc(lexer, start)
//     };
//   }
//   const operation = parseOperationType(lexer);
//   let name;
//   if (peek(lexer, TokenKind.NAME)) {
//     name = parseName(lexer);
//   }
//   return {
//     kind: OPERATION_DEFINITION,
//     operation,
//     name,
//     variableDefinitions: parseVariableDefinitions(lexer),
//     directives: parseDirectives(lexer),
//     selectionSet: parseSelectionSet(lexer),
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * OperationType : one of query mutation subscription
//  */
// function parseOperationType(lexer: Lexer<*>): OperationTypeNode {
//   const operationToken = expect(lexer, TokenKind.NAME);
//   switch (operationToken.value) {
//     case 'query': return 'query';
//     case 'mutation': return 'mutation';
//     // Note: subscription is an experimental non-spec addition.
//     case 'subscription': return 'subscription';
//   }

//   throw unexpected(lexer, operationToken);
// }

// /**
//  * VariableDefinitions : ( VariableDefinition+ )
//  */
// function parseVariableDefinitions(
//   lexer: Lexer<*>
// ): Array<VariableDefinitionNode> {
//   return peek(lexer, TokenKind.PAREN_L) ?
//     many(
//       lexer,
//       TokenKind.PAREN_L,
//       parseVariableDefinition,
//       TokenKind.PAREN_R
//     ) :
//     [];
// }

// /**
//  * VariableDefinition : Variable : Type DefaultValue?
//  */
// function parseVariableDefinition(lexer: Lexer<*>): VariableDefinitionNode {
//   const start = lexer.token;
//   return {
//     kind: VARIABLE_DEFINITION,
//     variable: parseVariable(lexer),
//     type: (expect(lexer, TokenKind.COLON), parseTypeReference(lexer)),
//     defaultValue:
//       skip(lexer, TokenKind.EQUALS) ? parseValueLiteral(lexer, true) : null,
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * Variable : $ Name
//  */
// function parseVariable(lexer: Lexer<*>): VariableNode {
//   const start = lexer.token;
//   expect(lexer, TokenKind.DOLLAR);
//   return {
//     kind: VARIABLE,
//     name: parseName(lexer),
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * SelectionSet : { Selection+ }
//  */
// function parseSelectionSet(lexer: Lexer<*>): SelectionSetNode {
//   const start = lexer.token;
//   return {
//     kind: SELECTION_SET,
//     selections:
//       many(lexer, TokenKind.BRACE_L, parseSelection, TokenKind.BRACE_R),
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * Selection :
//  *   - Field
//  *   - FragmentSpread
//  *   - InlineFragment
//  */
// function parseSelection(lexer: Lexer<*>): SelectionNode {
//   return peek(lexer, TokenKind.SPREAD) ?
//     parseFragment(lexer) :
//     parseField(lexer);
// }

// /**
//  * Field : Alias? Name Arguments? Directives? SelectionSet?
//  *
//  * Alias : Name :
//  */
// function parseField(lexer: Lexer<*>): FieldNode {
//   const start = lexer.token;

//   const nameOrAlias = parseName(lexer);
//   let alias;
//   let name;
//   if (skip(lexer, TokenKind.COLON)) {
//     alias = nameOrAlias;
//     name = parseName(lexer);
//   } else {
//     alias = null;
//     name = nameOrAlias;
//   }

//   return {
//     kind: FIELD,
//     alias,
//     name,
//     arguments: parseArguments(lexer),
//     directives: parseDirectives(lexer),
//     selectionSet:
//       peek(lexer, TokenKind.BRACE_L) ? parseSelectionSet(lexer) : null,
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * Arguments : ( Argument+ )
//  */
// function parseArguments(lexer: Lexer<*>): Array<ArgumentNode> {
//   return peek(lexer, TokenKind.PAREN_L) ?
//     many(lexer, TokenKind.PAREN_L, parseArgument, TokenKind.PAREN_R) :
//     [];
// }

// /**
//  * Argument : Name : Value
//  */
// function parseArgument(lexer: Lexer<*>): ArgumentNode {
//   const start = lexer.token;
//   return {
//     kind: ARGUMENT,
//     name: parseName(lexer),
//     value: (expect(lexer, TokenKind.COLON), parseValueLiteral(lexer, false)),
//     loc: loc(lexer, start)
//   };
// }

// // Implements the parsing rules in the Fragments section.

// /**
//  * Corresponds to both FragmentSpread and InlineFragment in the spec.
//  *
//  * FragmentSpread : ... FragmentName Directives?
//  *
//  * InlineFragment : ... TypeCondition? Directives? SelectionSet
//  */
// function parseFragment(
//   lexer: Lexer<*>
// ): FragmentSpreadNode | InlineFragmentNode {
//   const start = lexer.token;
//   expect(lexer, TokenKind.SPREAD);
//   if (peek(lexer, TokenKind.NAME) && lexer.token.value !== 'on') {
//     return {
//       kind: FRAGMENT_SPREAD,
//       name: parseFragmentName(lexer),
//       directives: parseDirectives(lexer),
//       loc: loc(lexer, start)
//     };
//   }
//   let typeCondition = null;
//   if (lexer.token.value === 'on') {
//     lexer.advance();
//     typeCondition = parseNamedType(lexer);
//   }
//   return {
//     kind: INLINE_FRAGMENT,
//     typeCondition,
//     directives: parseDirectives(lexer),
//     selectionSet: parseSelectionSet(lexer),
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * FragmentDefinition :
//  *   - fragment FragmentName on TypeCondition Directives? SelectionSet
//  *
//  * TypeCondition : NamedType
//  */
// function parseFragmentDefinition(lexer: Lexer<*>): FragmentDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'fragment');
//   return {
//     kind: FRAGMENT_DEFINITION,
//     name: parseFragmentName(lexer),
//     typeCondition: (expectKeyword(lexer, 'on'), parseNamedType(lexer)),
//     directives: parseDirectives(lexer),
//     selectionSet: parseSelectionSet(lexer),
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * FragmentName : Name but not `on`
//  */
// function parseFragmentName(lexer: Lexer<*>): NameNode {
//   if (lexer.token.value === 'on') {
//     throw unexpected(lexer);
//   }
//   return parseName(lexer);
// }

// // Implements the parsing rules in the Values section.

/**
 * Value[Const] :
 *   - [~Const] Variable
 *   - IntValue
 *   - FloatValue
 *   - StringValue
 *   - BooleanValue
 *   - NullValue
 *   - EnumValue
 *   - ListValue[?Const]
 *   - ObjectValue[?Const]
 *
 * BooleanValue : one of `true` `false`
 *
 * NullValue : `null`
 *
 * EnumValue : Name but not `true`, `false` or `null`
 */
func parseValueLiteral(lexer *Lexer, isConst bool) (language.ASTNode, error) {
	token := lexer.Token

	// switch token.Kind {
	// case language.TokenBracketLeft:
	// 	return parseList(lexer, isConst)
	// case language.TokenBraceLeft:
	// 	return parseObject(lexer, isConst)
	// case language.TokenInt:
	// 	lexer.Advance()

	// 	return language.IntValueNode{
	// 		Value: token.Value,
	// 		Loc:   loc(lexer, token),
	// 	}
	// case language.TokenFloat:
	// 	lexer.Advance()
	// 	return language.FloatValueNode{
	// 		Value: token.Value,
	// 		loc:   loc(lexer, token),
	// 	}
	//     case TokenKind.STRING:
	//       lexer.advance();
	//       return {
	//         kind: (STRING: 'StringValue'),
	//         value: ((token.value: any): string),
	//         loc: loc(lexer, token)
	//       };
	//     case TokenKind.NAME:
	//       if (token.value === 'true' || token.value === 'false') {
	//         lexer.advance();
	//         return {
	//           kind: (BOOLEAN: 'BooleanValue'),
	//           value: token.value === 'true',
	//           loc: loc(lexer, token)
	//         };
	//       } else if (token.value === 'null') {
	//         lexer.advance();
	//         return {
	//           kind: (NULL: 'NullValue'),
	//           loc: loc(lexer, token)
	//         };
	//       }
	//       lexer.advance();
	//       return {
	//         kind: (ENUM: 'EnumValue'),
	//         value: ((token.value: any): string),
	//         loc: loc(lexer, token)
	//       };
	//     case TokenKind.DOLLAR:
	//       if (!isConst) {
	//         return parseVariable(lexer);
	//       }
	//       break;
	//}

	return nil, unexpected(lexer, token)
}

func parseConstValue(lexer *Lexer) (language.ASTNode, error) {
	return parseValueLiteral(lexer, true)
}

func parseValueValue(lexer *Lexer) (language.ASTNode, error) {
	return parseValueLiteral(lexer, false)
}

/**
 * ListValue[Const] :
 *   - [ ]
 *   - [ Value[?Const]+ ]
 */
func parseList(lexer *Lexer, isConst bool) (language.ASTNode, error) {
	start := lexer.Token
	var item parser
	if isConst {
		item = parseConstValue
	} else {
		item = parseValueValue
	}

	vals, err := any(lexer, language.TokenBracketLeft, item, language.TokenBracketRight)
	if err != nil {
		return nil, err
	}

	return language.ListValueNode{language.Node{Loc: loc(lexer, start)}, vals}, nil
}

/**
 * ObjectValue[Const] :
 *   - { }
 *   - { ObjectField[?Const]+ }
 */
func parseObject(lexer *Lexer, isConst bool) (*language.ASTNode, error) {
	start := lexer.Token

	_, err := expect(lexer, language.TokenBraceLeft)
	if err != nil {
		return nil, err
	}

	fields := make([]language.ObjectFieldNode, 0)

	for !skip(lexer, language.TokenBraceRight) {
		f, err := parseObjectField(lexer, isConst)
		if err != nil {
			return nil, err
		}

		fields := append(fields, f)
	}

	return &language.ObjectValueNode{
		language.Node{Loc: loc(lexer, start)},
		fields,
	}, nil
}

/**
 * ObjectField[Const] : Name : Value[?Const]
 */
func parseObjectField(lexer *Lexer, isConst bool) (*language.ASTNode, error) {
	start := lexer.Token

	_, err := expect(lexer, language.TokenColon)
	if err != nil {
		return nil, err
	}

	name, err := parseName(lexer)
	if err != nil {
		return nil, err
	}

	val, err := parseValueLiteral(lexer, isConst)
	if err != nil {
		return nil, err
	}

	return &language.ObjectFieldNode{
		language.Node{Loc: loc(lexer, start)},
		name,
		val,
	}, nil
}

// // Implements the parsing rules in the Directives section.

// /**
//  * Directives : Directive+
//  */
// function parseDirectives(lexer: Lexer<*>): Array<DirectiveNode> {
//   const directives = [];
//   while (peek(lexer, TokenKind.AT)) {
//     directives.push(parseDirective(lexer));
//   }
//   return directives;
// }

// /**
//  * Directive : @ Name Arguments?
//  */
// function parseDirective(lexer: Lexer<*>): DirectiveNode {
//   const start = lexer.token;
//   expect(lexer, TokenKind.AT);
//   return {
//     kind: DIRECTIVE,
//     name: parseName(lexer),
//     arguments: parseArguments(lexer),
//     loc: loc(lexer, start)
//   };
// }

// // Implements the parsing rules in the Types section.

// /**
//  * Type :
//  *   - NamedType
//  *   - ListType
//  *   - NonNullType
//  */
// export function parseTypeReference(lexer: Lexer<*>): TypeNode {
//   const start = lexer.token;
//   let type;
//   if (skip(lexer, TokenKind.BRACKET_L)) {
//     type = parseTypeReference(lexer);
//     expect(lexer, TokenKind.BRACKET_R);
//     type = ({
//       kind: LIST_TYPE,
//       type,
//       loc: loc(lexer, start)
//     }: ListTypeNode);
//   } else {
//     type = parseNamedType(lexer);
//   }
//   if (skip(lexer, TokenKind.BANG)) {
//     return ({
//       kind: NON_NULL_TYPE,
//       type,
//       loc: loc(lexer, start)
//     }: NonNullTypeNode);
//   }
//   return type;
// }

// /**
//  * NamedType : Name
//  */
// export function parseNamedType(lexer: Lexer<*>): NamedTypeNode {
//   const start = lexer.token;
//   return {
//     kind: NAMED_TYPE,
//     name: parseName(lexer),
//     loc: loc(lexer, start)
//   };
// }

// // Implements the parsing rules in the Type Definition section.

// /**
//  * TypeSystemDefinition :
//  *   - SchemaDefinition
//  *   - TypeDefinition
//  *   - TypeExtensionDefinition
//  *   - DirectiveDefinition
//  *
//  * TypeDefinition :
//  *   - ScalarTypeDefinition
//  *   - ObjectTypeDefinition
//  *   - InterfaceTypeDefinition
//  *   - UnionTypeDefinition
//  *   - EnumTypeDefinition
//  *   - InputObjectTypeDefinition
//  */
// function parseTypeSystemDefinition(lexer: Lexer<*>): TypeSystemDefinitionNode {
//   if (peek(lexer, TokenKind.NAME)) {
//     switch (lexer.token.value) {
//       case 'schema': return parseSchemaDefinition(lexer);
//       case 'scalar': return parseScalarTypeDefinition(lexer);
//       case 'type': return parseObjectTypeDefinition(lexer);
//       case 'interface': return parseInterfaceTypeDefinition(lexer);
//       case 'union': return parseUnionTypeDefinition(lexer);
//       case 'enum': return parseEnumTypeDefinition(lexer);
//       case 'input': return parseInputObjectTypeDefinition(lexer);
//       case 'extend': return parseTypeExtensionDefinition(lexer);
//       case 'directive': return parseDirectiveDefinition(lexer);
//     }
//   }

//   throw unexpected(lexer);
// }

// /**
//  * SchemaDefinition : schema Directives? { OperationTypeDefinition+ }
//  *
//  * OperationTypeDefinition : OperationType : NamedType
//  */
// function parseSchemaDefinition(lexer: Lexer<*>): SchemaDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'schema');
//   const directives = parseDirectives(lexer);
//   const operationTypes = many(
//     lexer,
//     TokenKind.BRACE_L,
//     parseOperationTypeDefinition,
//     TokenKind.BRACE_R
//   );
//   return {
//     kind: SCHEMA_DEFINITION,
//     directives,
//     operationTypes,
//     loc: loc(lexer, start),
//   };
// }

// function parseOperationTypeDefinition(
//   lexer: Lexer<*>
// ): OperationTypeDefinitionNode {
//   const start = lexer.token;
//   const operation = parseOperationType(lexer);
//   expect(lexer, TokenKind.COLON);
//   const type = parseNamedType(lexer);
//   return {
//     kind: OPERATION_TYPE_DEFINITION,
//     operation,
//     type,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * ScalarTypeDefinition : scalar Name Directives?
//  */
// function parseScalarTypeDefinition(lexer: Lexer<*>): ScalarTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'scalar');
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   return {
//     kind: SCALAR_TYPE_DEFINITION,
//     name,
//     directives,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * ObjectTypeDefinition :
//  *   - type Name ImplementsInterfaces? Directives? { FieldDefinition+ }
//  */
// function parseObjectTypeDefinition(lexer: Lexer<*>): ObjectTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'type');
//   const name = parseName(lexer);
//   const interfaces = parseImplementsInterfaces(lexer);
//   const directives = parseDirectives(lexer);
//   const fields = any(
//     lexer,
//     TokenKind.BRACE_L,
//     parseFieldDefinition,
//     TokenKind.BRACE_R
//   );
//   return {
//     kind: OBJECT_TYPE_DEFINITION,
//     name,
//     interfaces,
//     directives,
//     fields,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * ImplementsInterfaces : implements NamedType+
//  */
// function parseImplementsInterfaces(lexer: Lexer<*>): Array<NamedTypeNode> {
//   const types = [];
//   if (lexer.token.value === 'implements') {
//     lexer.advance();
//     do {
//       types.push(parseNamedType(lexer));
//     } while (peek(lexer, TokenKind.NAME));
//   }
//   return types;
// }

// /**
//  * FieldDefinition : Name ArgumentsDefinition? : Type Directives?
//  */
// function parseFieldDefinition(lexer: Lexer<*>): FieldDefinitionNode {
//   const start = lexer.token;
//   const name = parseName(lexer);
//   const args = parseArgumentDefs(lexer);
//   expect(lexer, TokenKind.COLON);
//   const type = parseTypeReference(lexer);
//   const directives = parseDirectives(lexer);
//   return {
//     kind: FIELD_DEFINITION,
//     name,
//     arguments: args,
//     type,
//     directives,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * ArgumentsDefinition : ( InputValueDefinition+ )
//  */
// function parseArgumentDefs(lexer: Lexer<*>): Array<InputValueDefinitionNode> {
//   if (!peek(lexer, TokenKind.PAREN_L)) {
//     return [];
//   }
//   return many(lexer, TokenKind.PAREN_L, parseInputValueDef, TokenKind.PAREN_R);
// }

// /**
//  * InputValueDefinition : Name : Type DefaultValue? Directives?
//  */
// function parseInputValueDef(lexer: Lexer<*>): InputValueDefinitionNode {
//   const start = lexer.token;
//   const name = parseName(lexer);
//   expect(lexer, TokenKind.COLON);
//   const type = parseTypeReference(lexer);
//   let defaultValue = null;
//   if (skip(lexer, TokenKind.EQUALS)) {
//     defaultValue = parseConstValue(lexer);
//   }
//   const directives = parseDirectives(lexer);
//   return {
//     kind: INPUT_VALUE_DEFINITION,
//     name,
//     type,
//     defaultValue,
//     directives,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * InterfaceTypeDefinition : interface Name Directives? { FieldDefinition+ }
//  */
// function parseInterfaceTypeDefinition(
//   lexer: Lexer<*>
// ): InterfaceTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'interface');
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   const fields = any(
//     lexer,
//     TokenKind.BRACE_L,
//     parseFieldDefinition,
//     TokenKind.BRACE_R
//   );
//   return {
//     kind: INTERFACE_TYPE_DEFINITION,
//     name,
//     directives,
//     fields,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * UnionTypeDefinition : union Name Directives? = UnionMembers
//  */
// function parseUnionTypeDefinition(lexer: Lexer<*>): UnionTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'union');
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   expect(lexer, TokenKind.EQUALS);
//   const types = parseUnionMembers(lexer);
//   return {
//     kind: UNION_TYPE_DEFINITION,
//     name,
//     directives,
//     types,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * UnionMembers :
//  *   - `|`? NamedType
//  *   - UnionMembers | NamedType
//  */
// function parseUnionMembers(lexer: Lexer<*>): Array<NamedTypeNode> {
//   // Optional leading pipe
//   skip(lexer, TokenKind.PIPE);
//   const members = [];
//   do {
//     members.push(parseNamedType(lexer));
//   } while (skip(lexer, TokenKind.PIPE));
//   return members;
// }

// /**
//  * EnumTypeDefinition : enum Name Directives? { EnumValueDefinition+ }
//  */
// function parseEnumTypeDefinition(lexer: Lexer<*>): EnumTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'enum');
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   const values = many(
//     lexer,
//     TokenKind.BRACE_L,
//     parseEnumValueDefinition,
//     TokenKind.BRACE_R
//   );
//   return {
//     kind: ENUM_TYPE_DEFINITION,
//     name,
//     directives,
//     values,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * EnumValueDefinition : EnumValue Directives?
//  *
//  * EnumValue : Name
//  */
// function parseEnumValueDefinition(lexer: Lexer<*>): EnumValueDefinitionNode {
//   const start = lexer.token;
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   return {
//     kind: ENUM_VALUE_DEFINITION,
//     name,
//     directives,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * InputObjectTypeDefinition : input Name Directives? { InputValueDefinition+ }
//  */
// function parseInputObjectTypeDefinition(
//   lexer: Lexer<*>
// ): InputObjectTypeDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'input');
//   const name = parseName(lexer);
//   const directives = parseDirectives(lexer);
//   const fields = any(
//     lexer,
//     TokenKind.BRACE_L,
//     parseInputValueDef,
//     TokenKind.BRACE_R
//   );
//   return {
//     kind: INPUT_OBJECT_TYPE_DEFINITION,
//     name,
//     directives,
//     fields,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * TypeExtensionDefinition : extend ObjectTypeDefinition
//  */
// function parseTypeExtensionDefinition(
//   lexer: Lexer<*>
// ): TypeExtensionDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'extend');
//   const definition = parseObjectTypeDefinition(lexer);
//   return {
//     kind: TYPE_EXTENSION_DEFINITION,
//     definition,
//     loc: loc(lexer, start),
//   };
// }

// /**
//  * DirectiveDefinition :
//  *   - directive @ Name ArgumentsDefinition? on DirectiveLocations
//  */
// function parseDirectiveDefinition(lexer: Lexer<*>): DirectiveDefinitionNode {
//   const start = lexer.token;
//   expectKeyword(lexer, 'directive');
//   expect(lexer, TokenKind.AT);
//   const name = parseName(lexer);
//   const args = parseArgumentDefs(lexer);
//   expectKeyword(lexer, 'on');
//   const locations = parseDirectiveLocations(lexer);
//   return {
//     kind: DIRECTIVE_DEFINITION,
//     name,
//     arguments: args,
//     locations,
//     loc: loc(lexer, start)
//   };
// }

// /**
//  * DirectiveLocations :
//  *   - `|`? Name
//  *   - DirectiveLocations | Name
//  */
// function parseDirectiveLocations(lexer: Lexer<*>): Array<NameNode> {
//   // Optional leading pipe
//   skip(lexer, TokenKind.PIPE);
//   const locations = [];
//   do {
//     locations.push(parseName(lexer));
//   } while (skip(lexer, TokenKind.PIPE));
//   return locations;
// }

// Core parsing utility functions

/**
 * Returns a location object, used to identify the place in
 * the source that created a given parsed object.
 */
func loc(lexer *Lexer, startToken *language.Token) *language.Location {
	if lexer.options.NoLocation {
		return nil
	}

	return newLoc(*startToken, *lexer.LastToken, lexer.Source)
}

func newLoc(startToken language.Token, endToken language.Token, source language.Source) *language.Location {
	return &language.Location{
		Start:      startToken.Start,
		End:        endToken.End,
		StartToken: startToken,
		EndToken:   endToken,
		Source:     source,
	}
}

/**
 * Determines if the next token is of a given kind
 */
func peek(lexer *Lexer, kind language.TokenKind) bool {
	return lexer.Token.Kind == kind
}

/**
 * If the next token is of the given kind, return true after advancing
 * the lexer. Otherwise, do not change the parser state and return false.
 */
func skip(lexer *Lexer, kind language.TokenKind) bool {
	match := lexer.Token.Kind == kind
	if match {
		lexer.Advance()
	}
	return match
}

/**
 * If the next token is of the given kind, return that token after advancing
 * the lexer. Otherwise, do not change the parser state and throw an error.
 */
func expect(lexer *Lexer, kind language.TokenKind) (*language.Token, error) {
	token := lexer.Token
	if token.Kind == kind {
		lexer.Advance()
		return token, nil
	}

	return nil, errors.NewSyntaxError(
		lexer.Source,
		token.Start,
		fmt.Sprintf("Expected %s, found %s", string(kind), getTokenDesc(*token)),
	)
}

/**
 * If the next token is a keyword with the given value, return that token after
 * advancing the lexer. Otherwise, do not change the parser state and return
 * false.
 */
func expectKeyword(lexer *Lexer, value string) (*language.Token, error) {
	token := lexer.Token

	if token.Kind == language.TokenName && token.Value == value {
		lexer.Advance()
		return token, nil
	}

	return nil, errors.NewSyntaxError(
		lexer.Source,
		token.Start,
		fmt.Sprintf("Expected \"%s\", found %s", value, getTokenDesc(*token)),
	)
}

/**
 * Helper function for creating an error when an unexpected lexed token
 * is encountered.
 */
func unexpected(lexer *Lexer, atToken *language.Token) error {
	var token *language.Token
	if atToken != nil {
		token = atToken
	} else {
		token = lexer.Token
	}

	return errors.NewSyntaxError(
		lexer.Source,
		token.Start,
		fmt.Sprintf("Unexpected %s", getTokenDesc(*token)),
	)
}

/**
 * Returns a possibly empty list of parse nodes, determined by
 * the parseFn. This list begins with a lex token of openKind
 * and ends with a lex token of closeKind. Advances the parser
 * to the next lex token after the closing token.
 */
func any(
	lexer *Lexer,
	openKind language.TokenKind,
	parseFn parser,
	closeKind language.TokenKind,
) ([]language.ASTNode, error) {
	_, err := expect(lexer, openKind)
	if err != nil {
		return nil, err
	}

	nodes := make([]language.ASTNode, 0)

	for !skip(lexer, closeKind) {
		n, err := parseFn(lexer)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}

	return nodes, nil
}

/**
 * Returns a non-empty list of parse nodes, determined by
 * the parseFn. This list begins with a lex token of openKind
 * and ends with a lex token of closeKind. Advances the parser
 * to the next lex token after the closing token.
 */
func many(
	lexer *Lexer,
	openKind language.TokenKind,
	parseFn parser,
	closeKind language.TokenKind,
) ([]language.ASTNode, error) {
	_, err := expect(lexer, openKind)
	if err != nil {
		return nil, err
	}

	n, err := parseFn(lexer)
	if err != nil {
		return nil, err
	}

	nodes := []language.ASTNode{n}
	for !skip(lexer, closeKind) {
		n, err := parseFn(lexer)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}

	return nodes, nil
}
