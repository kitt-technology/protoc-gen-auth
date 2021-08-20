package auth

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"github.com/graphql-go/graphql/language/ast"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*graphql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*graphql.Field

var FieldBehaviour_enum = graphql.NewEnum(graphql.EnumConfig{
	Name: "FieldBehaviour",
	Values: graphql.EnumValueConfigMap{
		"ID": &graphql.EnumValueConfig{
			Value: FieldBehaviour(0),
		},
		"IDS": &graphql.EnumValueConfig{
			Value: FieldBehaviour(1),
		},
	},
})

var FieldBehaviour_type = graphql.NewScalar(graphql.ScalarConfig{
	Name: "FieldBehaviour",
	ParseValue: func(value interface{}) interface{} {
		return nil

	},
	Serialize: func(value interface{}) interface{} {
		return value.(FieldBehaviour).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})
