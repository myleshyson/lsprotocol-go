from __future__ import annotations

import itertools
import pathlib
from typing import Any

from generator import model


# first, we want to extract types and store those somewhere. Anything thats a Union type will become an interface.
# If its scoped to a Struct, the interface will be namespaced to that struct.
# Each struct will get a Marshal and Unmarshall func if it has one of these union types
def generate_from_spec(
    spec: model.LSPModel,
    output_dir: str,
    test_dir: str,
) -> None:
    """Generate the code for the given spec."""
    type_generator = TypeCodeGenerator(spec)
    code: dict[str, str] = type_generator.generate()
    output_path = pathlib.Path(output_dir)
    test_path = pathlib.Path(test_dir)

    if not output_path.exists():
        output_path.mkdir(parents=True, exist_ok=True)

    if not test_path.exists():
        test_path.mkdir(parents=True, exist_ok=True)

    for file_name, content in code.items():
        pathlib.Path(output_dir, file_name).write_text(
            content,
            encoding="utf-8",
        )

    # enum_arg = ",".join([name for name in type_generator.enums])

    # subprocess.run(
    #     [
    #         "stringer",
    #         "-type",
    #         enum_arg,
    #         "-output",
    #         "../type_strings.go",
    #         "../types.go",
    #     ],
    #     check=False,
    # )


def is_numeric(value: Any) -> bool:  # noqa: ANN401
    """Check if a value is numeric.

    Args:
        value: The value to check.

    Returns:
        bool: True if the value is an int or float, False otherwise.
    """
    return isinstance(value, (int, float))


def capitalize(value: str) -> str:
    """
    Capitalize the first letter of a string without modifying the rest of the string.

    Args:
        value: The string to capitalize.

    Returns:
        str: The capitalized string.
    """
    return value[0].upper() + value[1:]


def method_to_camel_case(value: str) -> str:
    return "".join(capitalize(section) for section in value.split("/"))


def lines_to_comments(lines: str | None, pad: int = 0) -> str:
    if lines == None:
        return ""

    tabs = "\t" * pad
    return "\n".join(f"{tabs}// {line}" for line in lines.splitlines())


def is_first_char_uppercase(string):
    return len(string) > 0 and string[0].isupper()


class LspTypeError(Exception):
    def __init__(self, message: str) -> None:
        super().__init__(f"Unsupported type: {message}")


class TypeCodeGenerator:
    def __init__(self, spec: model.LSPModel) -> None:
        self.spec = spec
        self.enums = set()
        self.constants = set()
        self.interfaces = {}
        self.maps = set()
        self.slices = set()
        self.interface_aliases = set()
        self.type_graph: dict[str, set[str]] = {}
        self.base_types = {
            "nil",
            "uint32",
            "int32",
            "float64",
            "string",
            "bool",
            "any",
        }

    def generate(self) -> dict[str, str]:
        header = self.generate_header()
        base_types = self.generate_base_types()
        structs = self.generate_structs()
        enums = self.generate_enums()
        requests = self.generate_requests()
        notifications = self.generate_notifications()
        or_types = self.generate_or_types()

        result = (
            header
            + base_types
            + structs
            + enums
            + requests
            + notifications
            + or_types
        )
        return {
            "lsprotocol/types.go": "\n".join(result),
        }

    def generate_notifications(self) -> list[str]:
        notifications = sorted(self.spec.notifications, key=lambda x: x.method)
        result = [
            "type NotificationMethod string\n",
        ]
        constants = "const (\n"
        constants += '\tUnknownNotificationMethod = ""\n'

        for notification in notifications:
            method = notification.method.replace("$", "Optional")
            camel_case_method = method_to_camel_case(method)
            self.constants.add(camel_case_method.lower())
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
            struct = [
                lines_to_comments(notification.documentation),
                f"type {notification.typeName} struct {{",
                '\tJsonRPC string `json:"jsonrpc"`',
                '\tMethod string `json:"method"`',
            ]
            if notification.params:
                struct.append(
                    f'\tParams {self._resolve_type(notification.params, "Notification")} `json:"params"`',
                )
            struct.append("}")
            result.append("\n".join(struct))
        result.append("\n")
        return result

    def generate_requests(self) -> list[str]:
        requests = sorted(self.spec.requests, key=lambda x: x.method)
        result = [
            "type RequestMethod string\n",
        ]
        constants = "const (\n"
        constants += '\tUnknownRequestMethod RequestMethod = ""\n'

        for request in requests:
            camel_case_method = method_to_camel_case(request.method)
            self.constants.add(camel_case_method.lower())
            constants += f'\t{camel_case_method}Method RequestMethod = "{request.method}"\n'
        constants += ")\n"
        result.append(constants)

        result.append("var RequestMethodMap = map[string]RequestMethod{")
        for request in requests:
            camel_case_method = method_to_camel_case(request.method)
            result.append(
                f'\t"{request.method}": {camel_case_method}Method,',
            )
        result.append("}\n")

        self.interfaces["Or_Request_id"] = ["int32", "string"]
        self.interfaces["Or_Response_id"] = ["int32", "string"]

        for request in requests:
            struct = [
                lines_to_comments(request.documentation),
                f"type {request.typeName} struct {{",
                '\tJsonRPC string `json:"jsonrpc"`',
                '\tID Or_Request_id `json:"id"`',
                '\tMethod RequestMethod `json:"method"`',
            ]
            if request.params:
                param_type = self._resolve_type(request.params, "Request")
                struct.append(
                    f'\tParams {param_type} `json:"params"`',
                )
            struct.append("}")
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
                "   if err := json.Unmarshal(x, &result); err != nil {",
                "       return err",
                "   }",
                f"   *t = {request.typeName}(result)",
                "   return nil",
                "}",
            ]
            result.append("\n".join(struct))

        for request in requests:
            if not request.result or not request.typeName:
                continue

            response_name = f"{request.typeName.replace('Request', 'Response')}"

            response_type = self._resolve_type(
                request.result,
                response_name,
            )

            omit_type = "omitempty"

            if is_first_char_uppercase(response_type):
                omit_type = "omitzero"

            if self._is_or_nullable(response_type):
                response_type = f"*{response_type}"

            struct = [
                f"type {response_name} struct {{",
                '   Id Or_Request_id `json:"id"`',
                f'  Result *{response_type} `json:"result,{omit_type}"`'
                if "nil" not in response_type
                else "",
                '   Error *ResponseError `json:"error,omitzero"`',
                "}",
            ]
            struct += [
                f"func (t *{response_name}) UnmarshalJSON(x []byte) error {{",
                "   var m map[string]any",
                "   if err := json.Unmarshal(x, &m); err != nil {",
                "       return err",
                "   }",
                '   _, idExists := m["id"]',
                '   _, jsonrpcExists := m["jsonrpc"]',
                "   if !idExists || !jsonrpcExists {",
                '       return fmt.Errorf("response must have an id and jsonrpc field.")',
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
            result.append("\n".join(struct))
            if response_type == "nil":
                continue
        result.append("\n")
        return result

    def generate_base_types(self) -> list[str]:
        result = [
            "type DocumentUri string\n",
            "type URI string\n",
        ]

        type_aliases = sorted(self.spec.typeAliases, key=lambda x: x.name)

        for alias in type_aliases:
            result.append(lines_to_comments(alias.documentation))

            match alias.name.lower():
                case "lspany":
                    resolved_type = "any"
                case "lspobject":
                    resolved_type = "map[string]any"
                case _:
                    resolved_type = self._resolve_type(alias.type, alias.name)

            result.append(
                f"type {alias.name} {resolved_type}\n",
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
                    "func (t Tuple[T1, T2]) UnmarshalJSON(data []byte) error {",
                    "   var arr [2]json.RawMessage",
                    "   if err := json.Unmarshal(data, &arr); err != nil {",
                    "       return err",
                    "   }",
                    "   if err := json.Unmarshal(arr[0], &t.V1); err != nil {",
                    "       return err",
                    "   }",
                    "   if err := json.Unmarshal(arr[1], &t.V2); err != nil {",
                    "       return err",
                    "   }",
                    "   return nil",
                    "}",
                ],
            ),
        )
        result.append(
            "\n".join(
                [
                    "type ResponseError struct {",
                    '\tCode int32 `json:"code"`',
                    '\tMessage string `json:"message"`',
                    '\tData any `json:"data,omitempty"`',
                    "}",
                ],
            ),
        )
        result.append("\n")
        return result

    def generate_header(self) -> list[str]:
        return [
            "package lsprotocol\n\n",
            "import (",
            '\t"encoding/json"',
            '\t"fmt"',
            '\t"bytes"',
            ")",
        ]

    def generate_enums(self) -> list[str]:
        results = []
        enums = sorted(self.spec.enumerations, key=lambda x: x.name)
        for enum in enums:
            if enum.documentation:
                results.append(lines_to_comments(enum.documentation))
            string = self._make_enum_type(enum)
            prefix = ""
            for value in enum.values:
                if value.name.lower() in self.constants:
                    prefix = enum.name + "_"
                    break

            string += "const (\n"
            values = sorted(enum.values, key=lambda x: x.name)
            for value in values:
                self.constants.add(value.name.lower())
                raw_value = value.value
                if not is_numeric(raw_value):
                    raw_value = f'"{raw_value}"'
                string += f"\t{prefix}{capitalize(value.name)} {enum.name} = {raw_value}\n"
            string += ")\n"

            results.append(string)

        return results

    def generate_structs(self) -> list[str]:
        """
        Golang doesn't have inheritance. Instead we'll just put all the properties in the params themselves
        """
        results = []
        structs = sorted(self.spec.structures, key=lambda s: s.name)

        for struct in structs:
            self.constants.add(struct.name.lower())

            properties: dict[str, model.Property] = {}
            self._get_extended_properties(struct, properties, [struct.name])

            # Sort properties by their dict key (property name)
            properties = {
                key: properties[key] for key in sorted(properties.keys())
            }

            result = [
                lines_to_comments(struct.documentation),
                f"type {struct.name} struct {{",
            ]

            for name, property in properties.items():
                formatted_name = capitalize(name)

                property_type = self._resolve_type(
                    property.type,
                    f"{struct.name}_{formatted_name}",
                    property.optional if property.optional else False,
                )

                if property_type == struct.name or (
                    property_type not in self.base_types
                    and "[]" not in property_type
                    and "map[" not in property_type
                    and property.optional
                ):
                    property_type = f"*{property_type}"

                omit_type = "omitempty"

                if is_first_char_uppercase(property_type):
                    omit_type = "omitzero"

                omit_empty = f",{omit_type}" if property.optional else ""

                property_string = f'\t{formatted_name} {property_type} `json:"{property.name}{omit_empty}"`'

                result.append(lines_to_comments(property.documentation, 1))
                result.append(property_string)

            result.append("}")

            results.append("\n".join(result))
        return results

    def generate_or_types(self) -> list[str]:
        results = []

        for interface_name, interface_types in self.interfaces.items():
            result = [
                f"// Or type for [{','.join(interface_types)}]\n"
                f"type {interface_name} struct {{",
                '\tValue any `json:"value"`',
                "}",
                self._create_or_type_marshaller(
                    interface_name,
                    interface_types,
                ),
                self._create_or_type_unmarshaller(
                    interface_name,
                    interface_types,
                ),
            ]
            results.append("\n".join(result))
        return results

    def _get_struct_from_name(
        self,
        name: str,
    ) -> model.Structure | None:
        for some in itertools.chain(
            self.spec.structures,
            self.spec.typeAliases,
            self.spec.enumerations,
        ):
            if some.name == name and isinstance(some, model.Structure):
                return some
        return None

    def _get_extended_properties(
        self,
        struct: model.Structure,
        properties: dict[str, model.Property],
        namespace: list[str],
    ):
        extends_list = struct.extends or []
        mixins_list = struct.mixins or []

        for t in itertools.chain(extends_list, mixins_list):
            if isinstance(t, model.ReferenceType):
                extended = self._get_struct_from_name(t.name)
                if extended:
                    self._get_extended_properties(
                        extended,
                        properties,
                        namespace,
                    )

        # add properties from the current structure to the set
        for property in struct.properties:
            properties[property.name] = property

    def _make_enum_type(self, enum: model.Enum) -> str:
        string = ""
        if enum.documentation:
            split_docs = enum.documentation.split("\n")
            split_docs = ["// " + line.rstrip("\n") for line in split_docs]
            string += "\n".join(split_docs) + "\n"

        enum_type = self._resolve_base_type(enum.type.name)

        enum_type_name = enum.name

        if enum_type != "string":
            self.enums.add(enum_type_name)
            string += f"//go:generate stringer -type={enum_type_name}\n"

        string += f"type {enum_type_name} {enum_type}\n"

        return string

    def _create_or_type_marshaller(self, or_type: str, types: list[str]) -> str:
        result = [
            f"func (t *{or_type}) MarshalJSON() ([]byte, error) {{",
            "   switch x := t.Value.(type) {",
            "   case nil:",
            '       return []byte("null"), nil',
        ]

        for i, name in enumerate(types):
            type = self._resolve_base_type(name)

            if type.lower() in ["nil", "null"]:
                continue

            branch = [
                f"    case {type}:",
                "        return json.Marshal(x)",
            ]
            result.append("\n".join(branch))

        names = ", ".join(
            [self._resolve_base_type(name) for name in types],
        )

        result.append(
            "\n".join(
                [
                    "   }\n"
                    f'  return nil, fmt.Errorf("unknown type %T is not one of {names}", t)',
                    "}",
                ],
            ),
        )
        return "\n".join(result)

    def _create_or_type_unmarshaller(
        self,
        or_type: str,
        types: list[str],
    ) -> str:
        result = [
            f"func (t *{or_type}) UnmarshalJSON(x []byte) error {{",
            "\tstringValue := string(x)",
            "\tdecoder := json.NewDecoder(bytes.NewReader(x))",
            "\tdecoder.DisallowUnknownFields()",
        ]

        names = ", ".join(
            [self._resolve_base_type(name) for name in types],
        )
        error_return = f'\treturn &UnmarshalError{{msg: "unmarshal failed to match one of {names} in {or_type} for value " + stringValue}}'

        if "nil" in self.interfaces[or_type]:
            result.append(
                "\n".join(
                    [
                        '   if string(x) == "null" {',
                        "       t.Value = nil",
                        "       return nil",
                        "   }",
                    ],
                ),
            )
        else:
            result.append(
                "\n".join(
                    [
                        '\tif string(x) == "null" {',
                        f"\t{error_return}",
                        "\t}",
                    ],
                ),
            )

        for i, name in enumerate(types):
            if name.lower() == "nil":
                continue
            branches = []
            if i != 0:
                branches += [
                    "\tdecoder = json.NewDecoder(bytes.NewReader(x))",
                    "\tdecoder.DisallowUnknownFields()",
                ]
            branches += [
                f"   var h{i} {self._resolve_base_type(name)}",
                f"   if err := decoder.Decode(&h{i}); err == nil {{",
                f"       t.Value = h{i}",
                "       return nil",
                "   }",
            ]
            result.append("\n".join(branches))

        result.append(error_return)
        result.append("}")

        return "\n".join(result)

    def _resolve_type(
        self,
        lsp_type: model.LSP_TYPE_SPEC
        | model.StringLiteralType
        | model.BaseType
        | model.OrType
        | model.ReferenceType
        | model.ArrayType
        | model.BaseMapKeyType
        | model.ReferenceMapKeyType,
        or_prefix: str | None = None,
        is_optional: bool = False,
    ):
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
            return f"{self._resolve_or_type(lsp_type, or_prefix)}"
        if isinstance(lsp_type, model.ReferenceType):
            if type_name.lower() == "lspany":
                return "any"
            return f"{type_name}"
        if isinstance(lsp_type, model.ArrayType):
            slice_type = self._resolve_type(
                lsp_type.element,
                or_prefix,
                is_optional,
            )
            if slice_type == "Or_None":
                print(type_name)
            return f"[]{slice_type}"
        if isinstance(lsp_type, model.MapType):
            key_type = self._resolve_type(lsp_type.key, or_prefix, is_optional)
            value_type = self._resolve_type(
                lsp_type.value,
                or_prefix,
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
            return f"Tuple[{', '.join(self._resolve_type(field, or_prefix, is_optional) for field in lsp_type.items)}]"
        raise LspTypeError(lsp_type)

    def _resolve_or_type(
        self,
        or_type: model.OrType,
        or_prefix: str | None = None,
    ):
        types = [self._resolve_type(type, or_prefix) for type in or_type.items]

        # this is "nullable". so make sure it's a pointer
        if "nil" in types and len(types) == 2:
            for type in types:
                if type != "nil" and "[]" not in type and "map" not in type:
                    return f"*{type}"

        interface_name = f"Or_{or_prefix}"

        self.type_graph[interface_name] = set(types)

        self.interfaces[interface_name] = types

        return interface_name

    def _is_or_nullable(self, type) -> bool:
        if not type.startswith("Or_"):
            return False
        if "[]" in type:
            return False
        for some in self.type_graph[type]:
            if some == "nil":
                return True
            if "Or_" in some and some in self.type_graph:
                return self._is_or_nullable(type)

        return False

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
