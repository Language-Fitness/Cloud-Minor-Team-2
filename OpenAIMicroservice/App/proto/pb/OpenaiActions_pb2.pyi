from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class KeyRequest(_message.Message):
    __slots__ = ("bearerToken",)
    BEARERTOKEN_FIELD_NUMBER: _ClassVar[int]
    bearerToken: str
    def __init__(self, bearerToken: _Optional[str] = ...) -> None: ...

class KeyResponse(_message.Message):
    __slots__ = ("key", "error")
    KEY_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    key: str
    error: str
    def __init__(self, key: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...
