from __future__ import annotations

import itertools

from generator import model

from .type_resolver import TypeResolver
from .utils import (
	capitalize,
	join,
	lines_to_comments,
)


def generate_structs(
	spec: model.LSPModel,
	type_resolver: TypeResolver,
) -> list[str]:
	"""
	Since the LSP specification relies on inheritance for object definitions while Go lacks this feature, we flatten the hierarchy by
	including all properties from parent classes directly in each struct definition.
	"""
	results = []
	structs = sorted(spec.structures, key=lambda s: s.name)

	for struct in structs:
		properties: dict[str, model.Property] = {}
		_get_extended_properties(struct, properties, [struct.name], spec)

		# Sort properties by their dict key (property name)
		properties = {key: properties[key] for key in sorted(properties.keys())}

		# this is actually should just be an alias to LSPObject
		if len(properties) == 0:
			results.append(
				join(
					[
						lines_to_comments(struct.documentation),
						f"type {struct.name} LSPObject",
					]
				)
			)
			continue
		result = [
			lines_to_comments(struct.documentation),
			f"type {struct.name} struct {{",
		]

		for name, property in properties.items():
			formatted_name = capitalize(name)

			property_type = type_resolver.resolve(
				property.type,
				property.optional if property.optional else False,
			)

			is_pointer = type_resolver.is_pointer(
				property_type,
				property.optional,
				struct.name,
			)

			property_type = f"*{property_type}" if is_pointer else property_type

			json_mapping = type_resolver.json_mapping(
				property_type,
				property.name,
				property.optional,
			)

			property_string = (
				f"\t{formatted_name} {property_type} {json_mapping}"
			)

			result.append(lines_to_comments(property.documentation, 1))
			result.append(property_string)

		result.append("}")

		results.append(join(result))
	return results


def _get_extended_properties(
	struct: model.Structure,
	properties: dict[str, model.Property],
	namespace: list[str],
	spec: model.LSPModel,
):
	"""
	Helper function that recursivley extracts properties through a structs extended classses.
	Its bottom up, so the top level class will override any properties that extended classes
	also have.
	"""
	extends_list = struct.extends or []
	mixins_list = struct.mixins or []

	for t in itertools.chain(extends_list, mixins_list):
		if isinstance(t, model.ReferenceType):
			extended = _get_struct_from_name(t.name, spec)
			if extended:
				_get_extended_properties(
					extended,
					properties,
					namespace,
					spec,
				)

	# add properties from the current structure to the set
	for property in struct.properties:
		properties[property.name] = property


def _get_struct_from_name(
	name: str,
	spec: model.LSPModel,
) -> model.Structure | None:
	"""
	Gets a struct class from a string name of that struct.
	"""
	for some in itertools.chain(
		spec.structures,
		spec.typeAliases,
		spec.enumerations,
	):
		if some.name == name and isinstance(some, model.Structure):
			return some
	return None
