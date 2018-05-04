// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
	graphql "github.com/sensu/sensu-go/graphql"
)

/*
JSONType ... JSON The JSON type describes any arbitrary JSON compatible data.

- Roughly equilevant to `union JSON = Int | Float | String | Boolean` however
  can also be a map or list of arbitrary length.
- Despite looking like an an object it's fields **cannot** be selected.
*/
var JSONType = graphql.NewType("JSON", graphql.ScalarKind)

// RegisterJSON registers JSON object type with given service.
func RegisterJSON(svc *graphql.Service, impl graphql.ScalarResolver) {
	svc.RegisterScalar(_ScalarTypeJSONDesc, impl)
}

// describe JSON's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ScalarTypeJSONDesc = graphql.ScalarDesc{Config: func() graphql1.ScalarConfig {
	return graphql1.ScalarConfig{
		Description: "The JSON type describes any arbitrary JSON compatible data.\n\n- Roughly equilevant to `union JSON = Int | Float | String | Boolean` however\n  can also be a map or list of arbitrary length.\n- Despite looking like an an object it's fields **cannot** be selected.",
		Name:        "JSON",
		ParseLiteral: func(_ ast.Value) interface{} {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ScalarResolver.")
		},
		ParseValue: func(_ interface{}) interface{} {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ScalarResolver.")
		},
		Serialize: func(_ interface{}) interface{} {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ScalarResolver.")
		},
	}
}}
