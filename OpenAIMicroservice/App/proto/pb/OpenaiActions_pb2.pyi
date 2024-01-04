from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class KeyRequest(_message.Message):
    __slots__ = ("school_id", "bearerToken")
    SCHOOL_ID_FIELD_NUMBER: _ClassVar[int]
    BEARERTOKEN_FIELD_NUMBER: _ClassVar[int]
    school_id: str
    bearerToken: str
    def __init__(self, school_id: _Optional[str] = ..., bearerToken: _Optional[str] = ...) -> None: ...

class KeyResponse(_message.Message):
    __slots__ = ("key", "error")
    KEY_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    key: str
    error: str
    def __init__(self, key: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...
