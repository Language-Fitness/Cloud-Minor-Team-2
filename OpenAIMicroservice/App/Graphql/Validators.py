import base64
import re
from App.Utils.Exceptions.ValidationException import ValidationException
import magic


def validate_minimum_int(variable_name, value):
    if value < 1:
        raise ValidationException(f"Amount at {variable_name} must be at least 1")


def validate_string(variable_name, value):
    if value is None or value.strip() == '':
        raise ValidationException(f"{variable_name} must not be empty")


def validate_answer_options(answer_options):
    if not answer_options or not isinstance(answer_options, list):
        raise ValidationException("Answer options must be a non-empty list.")

    if any(option is None or option.strip() == '' for option in answer_options):
        raise ValidationException("Answer options must not contain empty or whitespace-only strings.")


def validate_file(file_data_base64):
    try:
        # Decode the base64 data to bytes
        file_data = base64.b64decode(file_data_base64)
    except Exception:
        raise ValidationException('Invalid Base64 data used in file upload.')

    # Use python-magic to identify the file type from its content
    mime_type = magic.from_buffer(file_data, mime=True)

    # List of acceptable MIME types for PDF and Word documents
    valid_mime_types = ['application/pdf', 'application/msword',
                        'application/vnd.openxmlformats-officedocument.wordprocessingml.document']

    if mime_type not in valid_mime_types:
        raise ValidationException("Invalid file format. Please upload a valid file.")


def validate_base64(token):
    if not token or len(token) % 4 != 0:
        return False

    if not re.match('^[A-Za-z0-9+/]+={0,2}$', token):
        return False
    return True
