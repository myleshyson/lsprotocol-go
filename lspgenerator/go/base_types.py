import re

from generator import model

from .type_resolver import TypeResolver
from .utils import join, lines_to_comments


def generate_base_types(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	result = [
		"type DocumentUri string\n",
		"type URI string\n",
	]

	type_aliases = sorted(spec.typeAliases, key=lambda x: x.name)

	for alias in type_aliases:
		result.append(lines_to_comments(alias.documentation))

		match alias.name.lower():
			case "lspany":
				resolved_type = "any"
			case "lspobject":
				resolved_type = "map[string]any"
			case _:
				resolved_type = type_resolver.resolve(alias.type)

		result.append(
			f"type {alias.name} {resolved_type}\n",
		)

		# Aliases that point to an Or generic type, need to update their unmarshal type.
		if re.match(r"^Or\d\[", resolved_type):
			result.append(
				join(
					[
						f"func (t *{alias.name}) UnmarshalJSON(x []byte) error {{",
						f"	return (*{resolved_type})(t).UnmarshalJSON(x)",
						"}",
						f"func (t *{alias.name}) MarshalJSON() ([]byte, error) {{",
						f"	return (*{resolved_type})(t).MarshalJSON()",
						"}",
					],
				),
			)

	result.append(
		"\n".join(
			[
				"type UnmarshalError struct {",
				"   msg string",
				"}",
			],
		),
	)

	result.append(
		"\n".join(
			[
				"func (e UnmarshalError) Error() string {",
				"   return e.msg",
				"}",
			],
		),
	)
	result.append(
		"\n".join(
			[
				"type Tuple[T1, T2 any] struct {",
				"   V1 T1",
				"   V2 T2",
				"}",
			],
		),
	)
	result.append(
		"\n".join(
			[
				"func (t Tuple[T1, T2]) MarshalJSON() ([]byte, error) {",
				"   return json.Marshal([2]any{t.V1, t.V2})",
				"}",
			],
		),
	)

	result.append(
		"\n".join(
			[
				"func (t *Tuple[T1, T2]) UnmarshalJSON(data []byte) error {",
				"   var arr [2]json.RawMessage",
				"   if err := json.Unmarshal(data, &arr); err != nil {",
				"       return err",
				"   }",
				"   var t1 T1",
				"   var t2 T2",
				"   if err := json.Unmarshal(arr[0], &t1); err != nil {",
				"       return err",
				"   }",
				"   if err := json.Unmarshal(arr[1], &t2); err != nil {",
				"       return err",
				"   }",
				"   t.V1 = t1",
				"   t.V2 = t2",
				"   return nil",
				"}",
			],
		),
	)
	result.append(
		join(
			[
				"type ResponseError struct {",
				'\tCode int32 `json:"code"`',
				'\tMessage string `json:"message"`',
				'\tData any `json:"data,omitempty"`',
				"}",
			],
		),
	)
	result.append("type MethodKind string")
	result.append(
		join(
			[
				"type Message interface {",
				"	isMessage()",
				"}",
				"type IncomingMessage interface {",
				"	GetMethod() MethodKind",
				"	GetParams() any",
				"}",
				"type Request interface {",
				"	isRequest()",
				"}",
				"type Notification interface {",
				"	isNotification()",
				"}",
				"type Response interface {",
				"	isResponse()",
				"}",
			],
		),
	)
	result.append("\n")
	return result
