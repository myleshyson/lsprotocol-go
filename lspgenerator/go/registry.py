from generator import model

from .utils import join


def generate_registry(spec: model.LSPModel) -> str:
	result = [
		"package protocol",
		"import (",
		'	"encoding/json"',
		")",
		"var MessageRegistry = map[string]func([]byte) (Message, error) {",
	]

	result.extend(
		[
			join(
				[
					f'	"{request.method}": func(data []byte) (Message, error) {{',
					f"		var message {request.typeName}",
					"		if err := json.Unmarshal(data, &message); err != nil {",
					"			return nil, err",
					"		}",
					"		return message, nil",
					"	},",
				],
			)
			for request in spec.requests
		],
	)
	result.extend(
		[
			join(
				[
					f'	"{notification.method}": func(data []byte) (Message, error) {{',
					f"		var message {notification.typeName}",
					"		if err := json.Unmarshal(data, &message); err != nil {",
					"			return nil, err",
					"		}",
					"		return message, nil",
					"	},",
				],
			)
			for notification in spec.notifications
		],
	)
	result.append("}")

	return join(result)
