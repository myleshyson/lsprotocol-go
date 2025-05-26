from generator import model


class LspTypeError(Exception):
	def __init__(self, message: str) -> None:
		super().__init__(f"Unsupported type: {message}")


class TypeResolver:
	def __init__(self, spec: model.LSPModel) -> None:
		self.spec = spec
		self.max_or_length = 2
		self.base_types = {
			"nil",
			"uint32",
			"int32",
			"float64",
			"string",
			"bool",
			"any",
		}

	def is_pointer(
		self,
		type_str: str,
		is_optional: bool | None = False,
		struct: str | None = None,
	) -> bool:
		is_optional = is_optional or False
		return type_str == struct or (
			type_str not in self.base_types
			and "[]" not in type_str
			and "map[" not in type_str
			and is_optional
		)

	def json_mapping(
		self,
		type: str,
		field: str,
		optional: bool | None = False,
	) -> str:
		optional = optional or False

		omit_type = ",omitempty" if type in self.base_types else ",omitzero"

		if not optional:
			omit_type = ""

		return f'`json:"{field}{omit_type}"`'

	def resolve(
		self,
		lsp_type: model.LSP_TYPE_SPEC
		| model.StringLiteralType
		| model.BaseType
		| model.OrType
		| model.ReferenceType
		| model.ArrayType
		| model.BaseMapKeyType
		| model.ReferenceMapKeyType
		| model.EnumValueType,
		is_optional: bool = False,
	) -> str:
		if isinstance(lsp_type, model.AndType):
			raise Exception("Dont know what to do with this yet")

		type_name = (
			lsp_type.name if hasattr(lsp_type, "name") and lsp_type.name else ""  # pyright: ignore
		)

		if isinstance(lsp_type, model.StringLiteralType):
			return "string"
		if isinstance(lsp_type, model.BaseType):
			return f"{self._resolve_base_type(lsp_type.name)}"
		if isinstance(lsp_type, model.OrType):
			return f"{self._resolve_or_type(lsp_type)}"
		if isinstance(lsp_type, model.ReferenceType):
			if type_name.lower() == "lspany":
				return "any"
			return f"{type_name}"
		if isinstance(lsp_type, model.ArrayType):
			slice_type = self.resolve(
				lsp_type.element,
				is_optional,
			)
			return f"[]{slice_type}"
		if isinstance(lsp_type, model.MapType):
			key_type = self.resolve(lsp_type.key, is_optional)
			value_type = self.resolve(
				lsp_type.value,
				is_optional,
			)
			return f"map[{key_type}]{value_type}"
		if isinstance(lsp_type, model.BaseMapKeyType):
			return self._resolve_base_type(type_name)
		if isinstance(lsp_type, model.ReferenceMapKeyType):
			return f"{type_name}" if type_name else ""
		if isinstance(lsp_type, model.LiteralType):
			return "LSPObject"
		if isinstance(lsp_type, model.TupleType):
			return f"Tuple[{', '.join(self.resolve(field, is_optional) for field in lsp_type.items)}]"
		if isinstance(lsp_type, model.EnumValueType):
			return self._resolve_base_type(lsp_type.name)
		raise LspTypeError(lsp_type)

	def _resolve_or_type(
		self,
		or_type: model.OrType,
	):
		types = [self.resolve(type) for type in or_type.items]
		length = len(types)

		if "nil" in types:
			types.remove("nil")

		# this is just a "nullable" field. e.g. [Position, nil] is really just *Position
		if len(types) == 1:
			return f"*{types[0]}"

		pointer = "*" if len(types) < length else ""

		self.max_or_length = max(self.max_or_length, len(types))

		if pointer:
			return f"NullableOr{len(types)}[{', '.join(types)}]"

		return f"Or{len(types)}[{', '.join(types)}]"

	def _resolve_base_type(self, type: str):
		match type.lower():
			case "null":
				return "nil"
			case "nil":
				return "nil"
			case "uinteger":
				return "uint32"
			case "integer":
				return "int32"
			case "boolean":
				return "bool"
			case "float":
				return "float64"
			case "double":
				return "float64"
			case "decimal":
				return "float64"
			case "lspany":
				return "any"
			case "lspobject":
				return "map[string]any"
			case _:
				return type
