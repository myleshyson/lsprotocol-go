from .type_resolver import TypeResolver
from .utils import join


def generate_or_types(
	type_resolver: TypeResolver,
) -> list[str]:
	strings = []
	all_letters = list("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i in range(type_resolver.max_or_length - 1):
		strings.append(_or_type(i, all_letters, False))
		strings.append(_or_type(i, all_letters, True))
	return strings


def _or_type(i: int, all_letters: list[str], nullable: bool) -> str:
	or_length = i + 2
	generic_letters = list(all_letters[:or_length])
	generic_definition = join(
		[f"{char} any" for char in generic_letters],
		", ",
	)

	or_name = f"NullableOr{or_length}" if nullable else f"Or{or_length}"

	unmarshal_conditions = []

	for j, char in enumerate(generic_letters):
		unmarshal_conditions.append(
			join(
				[
					f"	var h{j} {char}",
					f"	decoder{j} := json.NewDecoder(bytes.NewReader(x))",
					f"	decoder{j}.DisallowUnknownFields()",
					f"	if err := decoder{j}.Decode(&h{j}); err == nil {{",
					f"		t.Value = h{j}",
					"		return nil",
					"	}",
				],
			),
		)

	if nullable:
		nullable_condition = join(
			[
				'	if string(x) == "null" {',
				"		t.Value = nil",
				"		return nil",
				"	}",
			],
		)
	else:
		nullable_condition = join(
			[
				'	if string(x) == "null" {',
				f'		return &UnmarshalError{{"{or_name}[{join(generic_letters, ",")}] cannot be null"}}',
				"	}",
			],
		)

	marshal_conditions = []
	for j, char in enumerate(generic_letters):
		marshal_conditions.append(
			join(
				[
					f"	case {char}:",
					"		return json.Marshal(x)",
				],
			),
		)
	marshal_nullable_condition = ""
	if nullable:
		marshal_nullable_condition = join(
			[
				"	if t.Value == nil {",
				'		return []byte("null"), nil',
				"	}",
			],
		)
	type_resolvers = []
	for i, char in enumerate(generic_letters):
		type_resolvers.append(
			join(
				[
					f"	var h{i} [0]{char}",
					f'	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h{i}).Elem()))',
				],
			),
		)
	result = [
		f"type {or_name}[{generic_definition}] struct {{",
		"	Value any",
		"}",
		f"func (t *{or_name}[{join(generic_letters, ',')}]) UnmarshalJSON(x []byte) error {{",
		nullable_condition,
		join(unmarshal_conditions),
		'	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))',
		"}",
		f"func (t {or_name}[{join(generic_letters, ',')}]) MarshalJSON() ([]byte, error) {{",
		marshal_nullable_condition,
		"	switch x := t.Value.(type) {",
		join(marshal_conditions),
		"	}\n",
		'	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())',
		"}",
		f"func (t {or_name}[{join(generic_letters, ',')}]) ConcreteTypes() string {{",
		"	var types []string",
		f"	{join(type_resolvers)}",
		'	return strings.Join(types, ",")',
		"}",
	]

	return join(result)
