from generator import model

from .type_resolver import TypeResolver
from .utils import capitalize, is_numeric, join, lines_to_comments


def generate_enums(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	results = []
	enums = sorted(spec.enumerations, key=lambda x: x.name)

	for enum in enums:
		if enum.documentation:
			results.append(lines_to_comments(enum.documentation))
		enum_type = type_resolver.resolve(enum.type)
		string = f"type {enum.name} {enum_type}\n"
		string += "const (\n"
		values = sorted(enum.values, key=lambda x: x.name)
		raw_values = []
		for value in values:
			raw_value = (
				f'"{value.value}"'
				if not is_numeric(value.value)
				else f"{value.value}"
			)
			raw_values.append(raw_value)
			string += f"\t{enum.name}{capitalize(value.name)} {enum.name} = {raw_value}\n"
		string += ")\n"
		if not enum.supportsCustomValues:
			string += join(
				[
					f"func (t {enum.name}) validate() error {{",
					"	switch t {",
					f"	case {join(list(set(raw_values)), ',')}:",
					"	return nil",
					"	}",
					f'	return fmt.Errorf("invalid {enum.name}: %v", t)',
					"}",
					"",
					f"func (t *{enum.name}) UnmarshalJSON(x []byte) error {{",
					f"	var test {enum_type}",
					"	if err := json.Unmarshal(x, &test); err != nil {",
					"		return err",
					"	}",
					f"	kind := {enum.name}(test)",
					"	if err := kind.validate(); err != nil {",
					"		return err",
					"	}",
					"	*t = kind",
					"	return nil",
					"}",
					"",
					f"func (t *{enum.name}) MarshalJSON() ([]byte, error) {{",
					"	if err := t.validate(); err != nil {",
					"		return nil, err",
					"	}",
					"	return json.Marshal(string(*t))",
					"}",
				],
			)
		results.append(string)

	return results
