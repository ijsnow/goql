package errors

import (
	"testing"
)

// import { expect } from 'chai';
// import { describe, it } from 'mocha';

// import { parse, Source, GraphQLError, formatError } from '../../';
func TestConvertsNodesToPositionsAndLocations(t *testing.T) {
	// 	source := language.NewSource(`{
	//   field
	// }`)

	//     const ast = parse(source);
	//     const fieldNode = ast.definitions[0].selectionSet.selections[0];
	//     const e = new GraphQLError('msg', [ fieldNode ]);
	//     expect(e.nodes).to.deep.equal([ fieldNode ]);
	//     expect(e.source).to.equal(source);
	//     expect(e.positions).to.deep.equal([ 8 ]);
	//     expect(e.locations).to.deep.equal([ { line: 2, column: 7 } ]);
}

//   it('converts node with loc.start === 0 to positions and locations', () => {
//     const source = new Source(`{
//       field
//     }`);
//     const ast = parse(source);
//     const operationNode = ast.definitions[0];
//     const e = new GraphQLError('msg', [ operationNode ]);
//     expect(e.nodes).to.deep.equal([ operationNode ]);
//     expect(e.source).to.equal(source);
//     expect(e.positions).to.deep.equal([ 0 ]);
//     expect(e.locations).to.deep.equal([ { line: 1, column: 1 } ]);
//   });

//   it('converts source and positions to locations', () => {
//     const source = new Source(`{
//       field
//     }`);
//     const e = new GraphQLError('msg', null, source, [ 10 ]);
//     expect(e.nodes).to.equal(undefined);
//     expect(e.source).to.equal(source);
//     expect(e.positions).to.deep.equal([ 10 ]);
//     expect(e.locations).to.deep.equal([ { line: 2, column: 9 } ]);
//   });

//   it('serializes to include message', () => {
//     const e = new GraphQLError('msg');
//     expect(JSON.stringify(e)).to.equal('{"message":"msg"}');
//   });

//   it('serializes to include message and locations', () => {
//     const node = parse('{ field }').definitions[0].selectionSet.selections[0];
//     const e = new GraphQLError('msg', [ node ]);
//     expect(JSON.stringify(e)).to.equal(
//       '{"message":"msg","locations":[{"line":1,"column":3}]}'
//     );
//   });

//   it('serializes to include path', () => {
//     const e = new GraphQLError(
//       'msg',
//       null,
//       null,
//       null,
//       [ 'path', 3, 'to', 'field' ]
//     );
//     expect(e.path).to.deep.equal([ 'path', 3, 'to', 'field' ]);
//     expect(JSON.stringify(e)).to.equal(
//       '{"message":"msg","path":["path",3,"to","field"]}'
//     );
//   });

//   it('default error formatter includes path', () => {
//     const e = new GraphQLError(
//       'msg',
//       null,
//       null,
//       null,
//       [ 'path', 3, 'to', 'field' ]
//     );

//     expect(formatError(e)).to.deep.equal({
//       message: 'msg',
//       locations: undefined,
//       path: [ 'path', 3, 'to', 'field' ]
//     });
//   });
