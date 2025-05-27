from generator import model

from .type_resolver import TypeResolver
from .utils import join, lines_to_comments, method_to_camel_case


def generate_notifications(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	notifications = sorted(spec.notifications, key=lambda x: x.method)
	result = []
	constants = "const (\n"
	constants += '\tUnknownNotificationMethod MethodKind = ""\n'

	for notification in notifications:
		method = notification.method.replace("$", "Optional")
		camel_case_method = method_to_camel_case(method)
		constants += f'\t{camel_case_method}Method MethodKind = "{notification.method}"\n'
	constants += ")\n"
	result.append(constants)
	result.append(
		"var NotificationMethodMap = map[string]MethodKind{",
	)
	for notification in notifications:
		method = notification.method.replace("$", "Optional")
		camel_case_method = method_to_camel_case(method)
		result.append(
			f'\t"{notification.method}": {camel_case_method}Method,',
		)
	result.append("}\n")

	for notification in notifications:
		param_type = (
			type_resolver.resolve(notification.params)
			if notification.params
			else "any"
		)
		struct = [
			lines_to_comments(notification.documentation),
			f"type {notification.typeName} struct {{",
			'\tJsonRPC string `json:"jsonrpc"`',
			'\tMethod MethodKind `json:"method"`',
			f'\tParams {param_type} `json:"params"`',
		]
		struct.append("}")
		struct.append(f"func (t {notification.typeName}) isMessage() {{}}")
		struct.append(
			f"func (t {notification.typeName}) isNotification() {{}}",
		)
		struct.append(
			f"func (t {notification.typeName}) GetMethod() MethodKind {{ return t.Method }}",
		)
		struct.append(
			f"func (t {notification.typeName}) GetParams() {param_type} {{ return t.Params }}",
		)
		struct += [
			f"func (t *{notification.typeName}) UnmarshalJSON(x []byte) error {{",
			"   var m map[string]any",
			"   if err := json.Unmarshal(x, &m); err != nil {",
			"       return err",
			"   }",
			'   if _, exists := m["method"]; !exists {',
			'       return fmt.Errorf("Missing required request field: method")',
			"   }",
			'   if _, exists := m["jsonrpc"]; !exists {',
			'       return fmt.Errorf("Missing required request field: jsonrpc")',
			"   }",
			f"  type Alias {notification.typeName}",
			"   var result Alias",
			"	decoder := json.NewDecoder(bytes.NewReader(x))",
			"	decoder.DisallowUnknownFields()",
			"   if err := decoder.Decode(&result); err != nil {",
			"       return err",
			"	}",
			f"	*t = {notification.typeName}(result)",
			"   return nil",
			"}",
		]
		result.append(join(struct))
	result.append("\n")
	return result
