from generator import model

from .type_resolver import TypeResolver
from .utils import join, lines_to_comments, method_to_camel_case


def generate_notifications(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	notifications = sorted(spec.notifications, key=lambda x: x.method)
	result = [
		"type NotificationMethod string\n",
	]
	constants = "const (\n"
	constants += '\tUnknownNotificationMethod = ""\n'

	for notification in notifications:
		method = notification.method.replace("$", "Optional")
		camel_case_method = method_to_camel_case(method)
		constants += f'\t{camel_case_method}Method NotificationMethod = "{notification.method}"\n'
	constants += ")\n"
	result.append(constants)
	result.append(
		"var NotificationMethodMap = map[string]NotificationMethod{",
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
			'\tMethod string `json:"method"`',
			f'\tParams {param_type} `json:"params"`',
		]
		struct.append("}")
		result.append(join(struct))
	result.append("\n")
	return result
