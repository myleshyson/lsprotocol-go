from generator import model

from .type_resolver import TypeResolver
from .utils import (
	join,
	lines_to_comments,
	method_to_camel_case,
)


def generate_requests(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	requests = sorted(spec.requests, key=lambda x: x.method)

	result = [
		"type RequestMethod string\n",
	]
	constants = "const (\n"
	constants += '\tUnknownRequestMethod RequestMethod = ""\n'

	for request in requests:
		camel_case_method = method_to_camel_case(request.method)
		constants += (
			f'\t{camel_case_method}Method RequestMethod = "{request.method}"\n'
		)
	constants += ")\n"
	result.append(constants)

	result.append("var RequestMethodMap = map[string]RequestMethod{")
	for request in requests:
		camel_case_method = method_to_camel_case(request.method)
		result.append(
			f'\t"{request.method}": {camel_case_method}Method,',
		)
	result.append("}\n")

	for request in requests:
		param_type = (
			type_resolver.resolve(request.params) if request.params else "any"
		)
		struct = [
			lines_to_comments(request.documentation),
			f"type {request.typeName} struct {{",
			'\tJsonRPC string `json:"jsonrpc"`',
			'\tID Or2[string, int32] `json:"id"`',
			'\tMethod RequestMethod `json:"method"`',
			f'\tParams {param_type} `json:"params"`',
		]

		struct.append("}")
		struct.append(f"func (t *{request.typeName}) isMessage() {{}}")
		struct.append(f"func (t *{request.typeName}) isRequest() {{}}")
		struct += [
			f"func (t *{request.typeName}) UnmarshalJSON(x []byte) error {{",
			"   var m map[string]any",
			"   if err := json.Unmarshal(x, &m); err != nil {",
			"       return err",
			"   }",
			'   if _, exists := m["method"]; !exists {',
			'       return fmt.Errorf("Missing required request field: method")',
			"   }",
			'   if _, exists := m["id"]; !exists {',
			'       return fmt.Errorf("Missing required request field: id")',
			"   }",
			'   if _, exists := m["jsonrpc"]; !exists {',
			'       return fmt.Errorf("Missing required request field: jsonrpc")',
			"   }",
			f"  type Alias {request.typeName}",
			"   var result Alias",
			"	decoder := json.NewDecoder(bytes.NewReader(x))",
			"	decoder.DisallowUnknownFields()",
			"   if err := decoder.Decode(&result); err != nil {",
			"       return err",
			"	}",
			f"	*t = {request.typeName}(result)",
			"   return nil",
			"}",
		]
		result.append(join(struct))

	for request in requests:
		if not request.result or not request.typeName:
			continue
		response_name = f"{request.typeName.replace('Request', 'Response')}"

		response_type = type_resolver.resolve(request.result, True)

		result_property = (
			f"\tResult {response_type} {type_resolver.json_mapping(response_type, 'result', True)}"
			if "nil" not in response_type
			else ""
		)

		struct = [
			f"type {response_name} struct {{",
			'	JsonRPC string `json:"jsonrpc"`',
			'	Id Or2[int32, string] `json:"id"`',
			result_property,
			'	Error *ResponseError `json:"error,omitzero"`',
			"}",
			f"func (t *{response_name}) UnmarshalJSON(x []byte) error {{",
			"   var m map[string]any",
			"   if err := json.Unmarshal(x, &m); err != nil {",
			"       return err",
			"   }",
			'   _, idExists := m["id"]',
			'   _, jsonrpcExists := m["jsonrpc"]',
			"   if !idExists || !jsonrpcExists {",
			'       return fmt.Errorf("response must have an id and jsonrpc field")',
			"   }",
			f"  type Alias {response_name}",
			"   var temp Alias",
			"   if err := json.Unmarshal(x, &temp); err != nil {",
			"       return err",
			"   }",
			f"  *t = {response_name}(temp)",
			"   return nil",
			"}",
		]
		struct.append(f"func (t *{response_name}) isMessage() {{}}")
		struct.append(f"func (t *{response_name}) isResponse() {{}}")
		result.append(join(struct))
	return result
