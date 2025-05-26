import re
from typing import Any


def is_numeric(value: Any) -> bool:
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
	"""
	Helper function to take a comment block and split it into properly indented string
	"""
	if lines == None:
		return ""

	tabs = "\t" * pad
	return "\n".join(f"{tabs}// {line}" for line in lines.splitlines())


def is_first_char_uppercase(string):
	return len(string) > 0 and string[0].isupper()


def join(value: list[str], glue: str = "\n") -> str:
	return glue.join(value)


def split_by_uppercase(text):
	# Method 1: Split before each uppercase letter (except the first)
	return re.sub(r"(?<!^)(?=[A-Z])", "_", text).split("_")
