from __future__ import annotations

import pathlib

from generator import model

from .base_types import generate_base_types
from .enums import generate_enums
from .notifications import generate_notifications
from .or_types import generate_or_types
from .requests import generate_requests
from .structs import generate_structs
from .tests import generate_tests
from .type_resolver import TypeResolver
from .utils import join


def generate_from_spec(
	spec: model.LSPModel,
	output_dir: str,
	test_dir: str,
) -> None:
	"""Generate the code for the given spec."""
	type_resolver = TypeResolver(spec)
	header = generate_header()
	base_types = generate_base_types(spec, type_resolver)
	structs = generate_structs(spec, type_resolver)
	enums = generate_enums(spec, type_resolver)
	requests = generate_requests(spec, type_resolver)
	notifications = generate_notifications(spec, type_resolver)
	or_types = generate_or_types(type_resolver)
	code = {
		"lsprotocol/types.go": join(
			header
			+ base_types
			+ structs
			+ requests
			+ notifications
			+ enums
			+ or_types,
		),
		"lsprotocol/types_test.go": generate_tests(spec),
	}
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


def generate_header() -> list[str]:
	return [
		"package protocol\n\n",
		"import (",
		'\t"encoding/json"',
		'\t"fmt"',
		'\t"bytes"',
		'\t"reflect"',
		'\t"strings"',
		")",
	]
