import io

class NamedBytesIO(io.BytesIO):
    def __init__(self, buffer, name):
        super().__init__(buffer)
        self.name = name

def convert_to_file_like_object(file_data, file_name):
    return NamedBytesIO(file_data, file_name)