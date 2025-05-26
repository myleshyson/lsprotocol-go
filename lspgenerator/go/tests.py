import os

from generator import model

from .utils import join


def generate_tests(spec: model.LSPModel):
	current_dir = os.path.dirname(os.path.abspath(__file__))
	stub_path = os.path.join(current_dir, "test.stub")

	with open(stub_path) as file:
		content = file.read()
		parts = content.split("// GENERATED_TESTS")
		cases = []
		for request in spec.requests:
			if not request.typeName:
				continue
			cases.append(
				join(
					[
						f'\t\tcase "{request.typeName}":',
						f"\t\tvalidateType[{request.typeName}](t, content, expectedResult, filepath)",
					],
				),
			)
			cases.append(
				join(
					[
						f'\t\tcase "{request.typeName.replace("Request", "Response")}":',
						f"\t\tvalidateType[{request.typeName.replace('Request', 'Response')}](t, content, expectedResult, filepath)",
					],
				),
			)

		for notification in spec.notifications:
			if not notification.typeName:
				continue
			cases.append(
				join(
					[
						f'\t\tcase "{notification.typeName}":',
						f"\t\tvalidateType[{notification.typeName}](t, content, expectedResult, filepath)",
					],
				),
			)

		return join([parts[0], join(cases), parts[1]])
