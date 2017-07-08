// Package schema holds the types and functions for creating your graphql schema
//
package schema

// Schema is the type that will hold your graphql schema
type Schema struct {
	astNode SchemaDefinitionNode
	//   _queryType: GraphQLObjectType;
	//   _mutationType: ?GraphQLObjectType;
	//   _subscriptionType: ?GraphQLObjectType;
	//   _directives: Array<GraphQLDirective>;
	//   _typeMap: TypeMap;
	//   _implementations: { [interfaceName: string]: Array<GraphQLObjectType> };
	//   _possibleTypeMap: ?{
	//     [abstractName: string]: { [possibleName: string]: boolean }
	//   };
}
