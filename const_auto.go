package tox
/*
#include "tox/tox.h"
*/
import "C"

var _ERR_OPTIONS_NEWS = make(map[int]string)
func init(){_ERR_OPTIONS_NEWS[-1] = "TE-1: _ERR_OPTIONS_NEW"}
const ERR_OPTIONS_NEW_OK = int(C.TOX_ERR_OPTIONS_NEW_OK) // 0
func init(){_ERR_OPTIONS_NEWS[ERR_OPTIONS_NEW_OK] = "TE00: The function returned successfully."}
const ERR_OPTIONS_NEW_MALLOC = int(C.TOX_ERR_OPTIONS_NEW_MALLOC) // 1
func init(){_ERR_OPTIONS_NEWS[ERR_OPTIONS_NEW_MALLOC] = "TE01: The function failed to allocate enough memory for the options struct."}

var _ERR_NEWS = make(map[int]string)
func init(){_ERR_NEWS[-1] = "TE-1: _ERR_NEW"}
const ERR_NEW_OK = int(C.TOX_ERR_NEW_OK) // 0
func init(){_ERR_NEWS[ERR_NEW_OK] = "TE00: The function returned successfully."}
const ERR_NEW_NULL = int(C.TOX_ERR_NEW_NULL) // 1
func init(){_ERR_NEWS[ERR_NEW_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_NEW_MALLOC = int(C.TOX_ERR_NEW_MALLOC) // 2
func init(){_ERR_NEWS[ERR_NEW_MALLOC] = "TE02: The function was unable to allocate enough memory to store the internal structures for the Tox object."}
const ERR_NEW_PORT_ALLOC = int(C.TOX_ERR_NEW_PORT_ALLOC) // 3
func init(){_ERR_NEWS[ERR_NEW_PORT_ALLOC] = "TE03: The function was unable to bind to a port. This may mean that all ports have already been bound, e.g. by other Tox instances, or it may mean a permission error. You may be able to gather more information from errno."}
const ERR_NEW_PROXY_BAD_TYPE = int(C.TOX_ERR_NEW_PROXY_BAD_TYPE) // 4
func init(){_ERR_NEWS[ERR_NEW_PROXY_BAD_TYPE] = "TE04: proxy_type was invalid."}
const ERR_NEW_PROXY_BAD_HOST = int(C.TOX_ERR_NEW_PROXY_BAD_HOST) // 5
func init(){_ERR_NEWS[ERR_NEW_PROXY_BAD_HOST] = "TE05: proxy_type was valid but the proxy_host passed had an invalid format or was NULL."}
const ERR_NEW_PROXY_BAD_PORT = int(C.TOX_ERR_NEW_PROXY_BAD_PORT) // 6
func init(){_ERR_NEWS[ERR_NEW_PROXY_BAD_PORT] = "TE06: proxy_type was valid, but the proxy_port was invalid."}
const ERR_NEW_PROXY_NOT_FOUND = int(C.TOX_ERR_NEW_PROXY_NOT_FOUND) // 7
func init(){_ERR_NEWS[ERR_NEW_PROXY_NOT_FOUND] = "TE07: The proxy address passed could not be resolved."}
const ERR_NEW_LOAD_ENCRYPTED = int(C.TOX_ERR_NEW_LOAD_ENCRYPTED) // 8
func init(){_ERR_NEWS[ERR_NEW_LOAD_ENCRYPTED] = "TE08: The byte array to be loaded contained an encrypted save."}
const ERR_NEW_LOAD_BAD_FORMAT = int(C.TOX_ERR_NEW_LOAD_BAD_FORMAT) // 9
func init(){_ERR_NEWS[ERR_NEW_LOAD_BAD_FORMAT] = "TE09: The data format was invalid. This can happen when loading data that was saved by an older version of Tox, or when the data has been corrupted. When loading from badly formatted data, some data may have been loaded, and the rest is discarded. Passing an invalid length parameter also causes this error."}

var _ERR_BOOTSTRAPS = make(map[int]string)
func init(){_ERR_BOOTSTRAPS[-1] = "TE-1: _ERR_BOOTSTRAP"}
const ERR_BOOTSTRAP_OK = int(C.TOX_ERR_BOOTSTRAP_OK) // 0
func init(){_ERR_BOOTSTRAPS[ERR_BOOTSTRAP_OK] = "TE00: The function returned successfully."}
const ERR_BOOTSTRAP_NULL = int(C.TOX_ERR_BOOTSTRAP_NULL) // 1
func init(){_ERR_BOOTSTRAPS[ERR_BOOTSTRAP_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_BOOTSTRAP_BAD_HOST = int(C.TOX_ERR_BOOTSTRAP_BAD_HOST) // 2
func init(){_ERR_BOOTSTRAPS[ERR_BOOTSTRAP_BAD_HOST] = "TE02: The hostname could not be resolved to an IP address, or the IP address passed was invalid."}
const ERR_BOOTSTRAP_BAD_PORT = int(C.TOX_ERR_BOOTSTRAP_BAD_PORT) // 3
func init(){_ERR_BOOTSTRAPS[ERR_BOOTSTRAP_BAD_PORT] = "TE03: The port passed was invalid. The valid port range is (1, 65535)."}

var _ERR_SET_INFOS = make(map[int]string)
func init(){_ERR_SET_INFOS[-1] = "TE-1: _ERR_SET_INFO"}
const ERR_SET_INFO_OK = int(C.TOX_ERR_SET_INFO_OK) // 0
func init(){_ERR_SET_INFOS[ERR_SET_INFO_OK] = "TE00: The function returned successfully."}
const ERR_SET_INFO_NULL = int(C.TOX_ERR_SET_INFO_NULL) // 1
func init(){_ERR_SET_INFOS[ERR_SET_INFO_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_SET_INFO_TOO_LONG = int(C.TOX_ERR_SET_INFO_TOO_LONG) // 2
func init(){_ERR_SET_INFOS[ERR_SET_INFO_TOO_LONG] = "TE02: Information length exceeded maximum permissible size."}

var _ERR_FRIEND_ADDS = make(map[int]string)
func init(){_ERR_FRIEND_ADDS[-1] = "TE-1: _ERR_FRIEND_ADD"}
const ERR_FRIEND_ADD_OK = int(C.TOX_ERR_FRIEND_ADD_OK) // 0
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_ADD_NULL = int(C.TOX_ERR_FRIEND_ADD_NULL) // 1
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FRIEND_ADD_TOO_LONG = int(C.TOX_ERR_FRIEND_ADD_TOO_LONG) // 2
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_TOO_LONG] = "TE02: The length of the friend request message exceeded TOX_MAX_FRIEND_REQUEST_LENGTH."}
const ERR_FRIEND_ADD_NO_MESSAGE = int(C.TOX_ERR_FRIEND_ADD_NO_MESSAGE) // 3
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_NO_MESSAGE] = "TE03: The friend request message was empty. This, and the TOO_LONG code will never be returned from tox_friend_add_norequest."}
const ERR_FRIEND_ADD_OWN_KEY = int(C.TOX_ERR_FRIEND_ADD_OWN_KEY) // 4
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_OWN_KEY] = "TE04: The friend address belongs to the sending client."}
const ERR_FRIEND_ADD_ALREADY_SENT = int(C.TOX_ERR_FRIEND_ADD_ALREADY_SENT) // 5
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_ALREADY_SENT] = "TE05: A friend request has already been sent, or the address belongs to a friend that is already on the friend list."}
const ERR_FRIEND_ADD_BAD_CHECKSUM = int(C.TOX_ERR_FRIEND_ADD_BAD_CHECKSUM) // 6
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_BAD_CHECKSUM] = "TE06: The friend address checksum failed."}
const ERR_FRIEND_ADD_SET_NEW_NOSPAM = int(C.TOX_ERR_FRIEND_ADD_SET_NEW_NOSPAM) // 7
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_SET_NEW_NOSPAM] = "TE07: The friend was already there, but the nospam value was different."}
const ERR_FRIEND_ADD_MALLOC = int(C.TOX_ERR_FRIEND_ADD_MALLOC) // 8
func init(){_ERR_FRIEND_ADDS[ERR_FRIEND_ADD_MALLOC] = "TE08: A memory allocation failed when trying to increase the friend list size."}

var _ERR_FRIEND_DELETES = make(map[int]string)
func init(){_ERR_FRIEND_DELETES[-1] = "TE-1: _ERR_FRIEND_DELETE"}
const ERR_FRIEND_DELETE_OK = int(C.TOX_ERR_FRIEND_DELETE_OK) // 0
func init(){_ERR_FRIEND_DELETES[ERR_FRIEND_DELETE_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_DELETE_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_DELETE_FRIEND_NOT_FOUND) // 1
func init(){_ERR_FRIEND_DELETES[ERR_FRIEND_DELETE_FRIEND_NOT_FOUND] = "TE01: There was no friend with the given friend number. No friends were deleted."}

var _ERR_FRIEND_BY_PUBLIC_KEYS = make(map[int]string)
func init(){_ERR_FRIEND_BY_PUBLIC_KEYS[-1] = "TE-1: _ERR_FRIEND_BY_PUBLIC_KEY"}
const ERR_FRIEND_BY_PUBLIC_KEY_OK = int(C.TOX_ERR_FRIEND_BY_PUBLIC_KEY_OK) // 0
func init(){_ERR_FRIEND_BY_PUBLIC_KEYS[ERR_FRIEND_BY_PUBLIC_KEY_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_BY_PUBLIC_KEY_NULL = int(C.TOX_ERR_FRIEND_BY_PUBLIC_KEY_NULL) // 1
func init(){_ERR_FRIEND_BY_PUBLIC_KEYS[ERR_FRIEND_BY_PUBLIC_KEY_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FRIEND_BY_PUBLIC_KEY_NOT_FOUND = int(C.TOX_ERR_FRIEND_BY_PUBLIC_KEY_NOT_FOUND) // 2
func init(){_ERR_FRIEND_BY_PUBLIC_KEYS[ERR_FRIEND_BY_PUBLIC_KEY_NOT_FOUND] = "TE02: No friend with the given Public Key exists on the friend list."}

var _ERR_FRIEND_GET_PUBLIC_KEYS = make(map[int]string)
func init(){_ERR_FRIEND_GET_PUBLIC_KEYS[-1] = "TE-1: _ERR_FRIEND_GET_PUBLIC_KEY"}
const ERR_FRIEND_GET_PUBLIC_KEY_OK = int(C.TOX_ERR_FRIEND_GET_PUBLIC_KEY_OK) // 0
func init(){_ERR_FRIEND_GET_PUBLIC_KEYS[ERR_FRIEND_GET_PUBLIC_KEY_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_GET_PUBLIC_KEY_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_GET_PUBLIC_KEY_FRIEND_NOT_FOUND) // 1
func init(){_ERR_FRIEND_GET_PUBLIC_KEYS[ERR_FRIEND_GET_PUBLIC_KEY_FRIEND_NOT_FOUND] = "TE01: No friend with the given number exists on the friend list."}

var _ERR_FRIEND_GET_LAST_ONLINES = make(map[int]string)
func init(){_ERR_FRIEND_GET_LAST_ONLINES[-1] = "TE-1: _ERR_FRIEND_GET_LAST_ONLINE"}
const ERR_FRIEND_GET_LAST_ONLINE_OK = int(C.TOX_ERR_FRIEND_GET_LAST_ONLINE_OK) // 0
func init(){_ERR_FRIEND_GET_LAST_ONLINES[ERR_FRIEND_GET_LAST_ONLINE_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_GET_LAST_ONLINE_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_GET_LAST_ONLINE_FRIEND_NOT_FOUND) // 1
func init(){_ERR_FRIEND_GET_LAST_ONLINES[ERR_FRIEND_GET_LAST_ONLINE_FRIEND_NOT_FOUND] = "TE01: No friend with the given number exists on the friend list."}

var _ERR_FRIEND_QUERYS = make(map[int]string)
func init(){_ERR_FRIEND_QUERYS[-1] = "TE-1: _ERR_FRIEND_QUERY"}
const ERR_FRIEND_QUERY_OK = int(C.TOX_ERR_FRIEND_QUERY_OK) // 0
func init(){_ERR_FRIEND_QUERYS[ERR_FRIEND_QUERY_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_QUERY_NULL = int(C.TOX_ERR_FRIEND_QUERY_NULL) // 1
func init(){_ERR_FRIEND_QUERYS[ERR_FRIEND_QUERY_NULL] = "TE01: The pointer parameter for storing the query result (name, message) was NULL. Unlike the `_self_` variants of these functions, which have no effect when a parameter is NULL, these functions return an error in that case."}
const ERR_FRIEND_QUERY_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_QUERY_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FRIEND_QUERYS[ERR_FRIEND_QUERY_FRIEND_NOT_FOUND] = "TE02: The friend_number did not designate a valid friend."}

var _ERR_SET_TYPINGS = make(map[int]string)
func init(){_ERR_SET_TYPINGS[-1] = "TE-1: _ERR_SET_TYPING"}
const ERR_SET_TYPING_OK = int(C.TOX_ERR_SET_TYPING_OK) // 0
func init(){_ERR_SET_TYPINGS[ERR_SET_TYPING_OK] = "TE00: The function returned successfully."}
const ERR_SET_TYPING_FRIEND_NOT_FOUND = int(C.TOX_ERR_SET_TYPING_FRIEND_NOT_FOUND) // 1
func init(){_ERR_SET_TYPINGS[ERR_SET_TYPING_FRIEND_NOT_FOUND] = "TE01: The friend number did not designate a valid friend."}

var _ERR_FRIEND_SEND_MESSAGES = make(map[int]string)
func init(){_ERR_FRIEND_SEND_MESSAGES[-1] = "TE-1: _ERR_FRIEND_SEND_MESSAGE"}
const ERR_FRIEND_SEND_MESSAGE_OK = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_OK) // 0
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_SEND_MESSAGE_NULL = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_NULL) // 1
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_FOUND] = "TE02: The friend number did not designate a valid friend."}
const ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_CONNECTED) // 3
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_FRIEND_NOT_CONNECTED] = "TE03: This client is currently not connected to the friend."}
const ERR_FRIEND_SEND_MESSAGE_SENDQ = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_SENDQ) // 4
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_SENDQ] = "TE04: An allocation error occurred while increasing the send queue size."}
const ERR_FRIEND_SEND_MESSAGE_TOO_LONG = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_TOO_LONG) // 5
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_TOO_LONG] = "TE05: Message length exceeded TOX_MAX_MESSAGE_LENGTH."}
const ERR_FRIEND_SEND_MESSAGE_EMPTY = int(C.TOX_ERR_FRIEND_SEND_MESSAGE_EMPTY) // 6
func init(){_ERR_FRIEND_SEND_MESSAGES[ERR_FRIEND_SEND_MESSAGE_EMPTY] = "TE06: Attempted to send a zero-length message."}

var _ERR_FILE_CONTROLS = make(map[int]string)
func init(){_ERR_FILE_CONTROLS[-1] = "TE-1: _ERR_FILE_CONTROL"}
const ERR_FILE_CONTROL_OK = int(C.TOX_ERR_FILE_CONTROL_OK) // 0
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_OK] = "TE00: The function returned successfully."}
const ERR_FILE_CONTROL_FRIEND_NOT_FOUND = int(C.TOX_ERR_FILE_CONTROL_FRIEND_NOT_FOUND) // 1
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_FRIEND_NOT_FOUND] = "TE01: The friend_number passed did not designate a valid friend."}
const ERR_FILE_CONTROL_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FILE_CONTROL_FRIEND_NOT_CONNECTED) // 2
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_FRIEND_NOT_CONNECTED] = "TE02: This client is currently not connected to the friend."}
const ERR_FILE_CONTROL_NOT_FOUND = int(C.TOX_ERR_FILE_CONTROL_NOT_FOUND) // 3
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_NOT_FOUND] = "TE03: No file transfer with the given file number was found for the given friend."}
const ERR_FILE_CONTROL_NOT_PAUSED = int(C.TOX_ERR_FILE_CONTROL_NOT_PAUSED) // 4
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_NOT_PAUSED] = "TE04: A RESUME control was sent, but the file transfer is running normally."}
const ERR_FILE_CONTROL_DENIED = int(C.TOX_ERR_FILE_CONTROL_DENIED) // 5
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_DENIED] = "TE05: A RESUME control was sent, but the file transfer was paused by the other party. Only the party that paused the transfer can resume it."}
const ERR_FILE_CONTROL_ALREADY_PAUSED = int(C.TOX_ERR_FILE_CONTROL_ALREADY_PAUSED) // 6
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_ALREADY_PAUSED] = "TE06: A PAUSE control was sent, but the file transfer was already paused."}
const ERR_FILE_CONTROL_SENDQ = int(C.TOX_ERR_FILE_CONTROL_SENDQ) // 7
func init(){_ERR_FILE_CONTROLS[ERR_FILE_CONTROL_SENDQ] = "TE07: Packet queue is full."}

var _ERR_FILE_SEEKS = make(map[int]string)
func init(){_ERR_FILE_SEEKS[-1] = "TE-1: _ERR_FILE_SEEK"}
const ERR_FILE_SEEK_OK = int(C.TOX_ERR_FILE_SEEK_OK) // 0
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_OK] = "TE00: The function returned successfully."}
const ERR_FILE_SEEK_FRIEND_NOT_FOUND = int(C.TOX_ERR_FILE_SEEK_FRIEND_NOT_FOUND) // 1
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_FRIEND_NOT_FOUND] = "TE01: The friend_number passed did not designate a valid friend."}
const ERR_FILE_SEEK_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FILE_SEEK_FRIEND_NOT_CONNECTED) // 2
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_FRIEND_NOT_CONNECTED] = "TE02: This client is currently not connected to the friend."}
const ERR_FILE_SEEK_NOT_FOUND = int(C.TOX_ERR_FILE_SEEK_NOT_FOUND) // 3
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_NOT_FOUND] = "TE03: No file transfer with the given file number was found for the given friend."}
const ERR_FILE_SEEK_DENIED = int(C.TOX_ERR_FILE_SEEK_DENIED) // 4
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_DENIED] = "TE04: File was not in a state where it could be seeked."}
const ERR_FILE_SEEK_INVALID_POSITION = int(C.TOX_ERR_FILE_SEEK_INVALID_POSITION) // 5
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_INVALID_POSITION] = "TE05: Seek position was invalid"}
const ERR_FILE_SEEK_SENDQ = int(C.TOX_ERR_FILE_SEEK_SENDQ) // 6
func init(){_ERR_FILE_SEEKS[ERR_FILE_SEEK_SENDQ] = "TE06: Packet queue is full."}

var _ERR_FILE_GETS = make(map[int]string)
func init(){_ERR_FILE_GETS[-1] = "TE-1: _ERR_FILE_GET"}
const ERR_FILE_GET_OK = int(C.TOX_ERR_FILE_GET_OK) // 0
func init(){_ERR_FILE_GETS[ERR_FILE_GET_OK] = "TE00: The function returned successfully."}
const ERR_FILE_GET_NULL = int(C.TOX_ERR_FILE_GET_NULL) // 1
func init(){_ERR_FILE_GETS[ERR_FILE_GET_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FILE_GET_FRIEND_NOT_FOUND = int(C.TOX_ERR_FILE_GET_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FILE_GETS[ERR_FILE_GET_FRIEND_NOT_FOUND] = "TE02: The friend_number passed did not designate a valid friend."}
const ERR_FILE_GET_NOT_FOUND = int(C.TOX_ERR_FILE_GET_NOT_FOUND) // 3
func init(){_ERR_FILE_GETS[ERR_FILE_GET_NOT_FOUND] = "TE03: No file transfer with the given file number was found for the given friend."}

var _ERR_FILE_SENDS = make(map[int]string)
func init(){_ERR_FILE_SENDS[-1] = "TE-1: _ERR_FILE_SEND"}
const ERR_FILE_SEND_OK = int(C.TOX_ERR_FILE_SEND_OK) // 0
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_OK] = "TE00: The function returned successfully."}
const ERR_FILE_SEND_NULL = int(C.TOX_ERR_FILE_SEND_NULL) // 1
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FILE_SEND_FRIEND_NOT_FOUND = int(C.TOX_ERR_FILE_SEND_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_FRIEND_NOT_FOUND] = "TE02: The friend_number passed did not designate a valid friend."}
const ERR_FILE_SEND_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FILE_SEND_FRIEND_NOT_CONNECTED) // 3
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_FRIEND_NOT_CONNECTED] = "TE03: This client is currently not connected to the friend."}
const ERR_FILE_SEND_NAME_TOO_LONG = int(C.TOX_ERR_FILE_SEND_NAME_TOO_LONG) // 4
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_NAME_TOO_LONG] = "TE04: Filename length exceeded TOX_MAX_FILENAME_LENGTH bytes."}
const ERR_FILE_SEND_TOO_MANY = int(C.TOX_ERR_FILE_SEND_TOO_MANY) // 5
func init(){_ERR_FILE_SENDS[ERR_FILE_SEND_TOO_MANY] = "TE05: Too many ongoing transfers. The maximum number of concurrent file transfers is 256 per friend per direction (sending and receiving)."}

var _ERR_FILE_SEND_CHUNKS = make(map[int]string)
func init(){_ERR_FILE_SEND_CHUNKS[-1] = "TE-1: _ERR_FILE_SEND_CHUNK"}
const ERR_FILE_SEND_CHUNK_OK = int(C.TOX_ERR_FILE_SEND_CHUNK_OK) // 0
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_OK] = "TE00: The function returned successfully."}
const ERR_FILE_SEND_CHUNK_NULL = int(C.TOX_ERR_FILE_SEND_CHUNK_NULL) // 1
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_NULL] = "TE01: The length parameter was non-zero, but data was NULL."}
const ERR_FILE_SEND_CHUNK_FRIEND_NOT_FOUND = int(C.TOX_ERR_FILE_SEND_CHUNK_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_FRIEND_NOT_FOUND] = "TE02: The friend_number passed did not designate a valid friend."}
const ERR_FILE_SEND_CHUNK_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FILE_SEND_CHUNK_FRIEND_NOT_CONNECTED) // 3
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_FRIEND_NOT_CONNECTED] = "TE03: This client is currently not connected to the friend."}
const ERR_FILE_SEND_CHUNK_NOT_FOUND = int(C.TOX_ERR_FILE_SEND_CHUNK_NOT_FOUND) // 4
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_NOT_FOUND] = "TE04: No file transfer with the given file number was found for the given friend."}
const ERR_FILE_SEND_CHUNK_NOT_TRANSFERRING = int(C.TOX_ERR_FILE_SEND_CHUNK_NOT_TRANSFERRING) // 5
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_NOT_TRANSFERRING] = "TE05: File transfer was found but isn't in a transferring state: (paused, done, broken, etc...) (happens only when not called from the request chunk callback)."}
const ERR_FILE_SEND_CHUNK_INVALID_LENGTH = int(C.TOX_ERR_FILE_SEND_CHUNK_INVALID_LENGTH) // 6
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_INVALID_LENGTH] = "TE06: Attempted to send more or less data than requested. The requested data size is adjusted according to maximum transmission unit and the expected end of the file. Trying to send less or more than requested will return this error."}
const ERR_FILE_SEND_CHUNK_SENDQ = int(C.TOX_ERR_FILE_SEND_CHUNK_SENDQ) // 7
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_SENDQ] = "TE07: Packet queue is full."}
const ERR_FILE_SEND_CHUNK_WRONG_POSITION = int(C.TOX_ERR_FILE_SEND_CHUNK_WRONG_POSITION) // 8
func init(){_ERR_FILE_SEND_CHUNKS[ERR_FILE_SEND_CHUNK_WRONG_POSITION] = "TE08: Position parameter was wrong."}

var _ERR_CONFERENCE_NEWS = make(map[int]string)
func init(){_ERR_CONFERENCE_NEWS[-1] = "TE-1: _ERR_CONFERENCE_NEW"}
const ERR_CONFERENCE_NEW_OK = int(C.TOX_ERR_CONFERENCE_NEW_OK) // 0
func init(){_ERR_CONFERENCE_NEWS[ERR_CONFERENCE_NEW_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_NEW_INIT = int(C.TOX_ERR_CONFERENCE_NEW_INIT) // 1
func init(){_ERR_CONFERENCE_NEWS[ERR_CONFERENCE_NEW_INIT] = "TE01: The conference instance failed to initialize."}

var _ERR_CONFERENCE_DELETES = make(map[int]string)
func init(){_ERR_CONFERENCE_DELETES[-1] = "TE-1: _ERR_CONFERENCE_DELETE"}
const ERR_CONFERENCE_DELETE_OK = int(C.TOX_ERR_CONFERENCE_DELETE_OK) // 0
func init(){_ERR_CONFERENCE_DELETES[ERR_CONFERENCE_DELETE_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_DELETE_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_DELETE_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_DELETES[ERR_CONFERENCE_DELETE_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}

var _ERR_CONFERENCE_PEER_QUERYS = make(map[int]string)
func init(){_ERR_CONFERENCE_PEER_QUERYS[-1] = "TE-1: _ERR_CONFERENCE_PEER_QUERY"}
const ERR_CONFERENCE_PEER_QUERY_OK = int(C.TOX_ERR_CONFERENCE_PEER_QUERY_OK) // 0
func init(){_ERR_CONFERENCE_PEER_QUERYS[ERR_CONFERENCE_PEER_QUERY_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_PEER_QUERY_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_PEER_QUERY_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_PEER_QUERYS[ERR_CONFERENCE_PEER_QUERY_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}
const ERR_CONFERENCE_PEER_QUERY_PEER_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_PEER_QUERY_PEER_NOT_FOUND) // 2
func init(){_ERR_CONFERENCE_PEER_QUERYS[ERR_CONFERENCE_PEER_QUERY_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
const ERR_CONFERENCE_PEER_QUERY_NO_CONNECTION = int(C.TOX_ERR_CONFERENCE_PEER_QUERY_NO_CONNECTION) // 3
func init(){_ERR_CONFERENCE_PEER_QUERYS[ERR_CONFERENCE_PEER_QUERY_NO_CONNECTION] = "TE03: The client is not connected to the conference."}

var _ERR_CONFERENCE_INVITES = make(map[int]string)
func init(){_ERR_CONFERENCE_INVITES[-1] = "TE-1: _ERR_CONFERENCE_INVITE"}
const ERR_CONFERENCE_INVITE_OK = int(C.TOX_ERR_CONFERENCE_INVITE_OK) // 0
func init(){_ERR_CONFERENCE_INVITES[ERR_CONFERENCE_INVITE_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_INVITE_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_INVITE_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_INVITES[ERR_CONFERENCE_INVITE_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}
const ERR_CONFERENCE_INVITE_FAIL_SEND = int(C.TOX_ERR_CONFERENCE_INVITE_FAIL_SEND) // 2
func init(){_ERR_CONFERENCE_INVITES[ERR_CONFERENCE_INVITE_FAIL_SEND] = "TE02: The invite packet failed to send."}
const ERR_CONFERENCE_INVITE_NO_CONNECTION = int(C.TOX_ERR_CONFERENCE_INVITE_NO_CONNECTION) // 3
func init(){_ERR_CONFERENCE_INVITES[ERR_CONFERENCE_INVITE_NO_CONNECTION] = "TE03: The client is not connected to the conference."}

var _ERR_CONFERENCE_JOINS = make(map[int]string)
func init(){_ERR_CONFERENCE_JOINS[-1] = "TE-1: _ERR_CONFERENCE_JOIN"}
const ERR_CONFERENCE_JOIN_OK = int(C.TOX_ERR_CONFERENCE_JOIN_OK) // 0
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_JOIN_INVALID_LENGTH = int(C.TOX_ERR_CONFERENCE_JOIN_INVALID_LENGTH) // 1
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_INVALID_LENGTH] = "TE01: The cookie passed has an invalid length."}
const ERR_CONFERENCE_JOIN_WRONG_TYPE = int(C.TOX_ERR_CONFERENCE_JOIN_WRONG_TYPE) // 2
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_WRONG_TYPE] = "TE02: The conference is not the expected type. This indicates an invalid cookie."}
const ERR_CONFERENCE_JOIN_FRIEND_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_JOIN_FRIEND_NOT_FOUND) // 3
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_FRIEND_NOT_FOUND] = "TE03: The friend number passed does not designate a valid friend."}
const ERR_CONFERENCE_JOIN_DUPLICATE = int(C.TOX_ERR_CONFERENCE_JOIN_DUPLICATE) // 4
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_DUPLICATE] = "TE04: Client is already in this conference."}
const ERR_CONFERENCE_JOIN_INIT_FAIL = int(C.TOX_ERR_CONFERENCE_JOIN_INIT_FAIL) // 5
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_INIT_FAIL] = "TE05: Conference instance failed to initialize."}
const ERR_CONFERENCE_JOIN_FAIL_SEND = int(C.TOX_ERR_CONFERENCE_JOIN_FAIL_SEND) // 6
func init(){_ERR_CONFERENCE_JOINS[ERR_CONFERENCE_JOIN_FAIL_SEND] = "TE06: The join packet failed to send."}

var _ERR_CONFERENCE_SEND_MESSAGES = make(map[int]string)
func init(){_ERR_CONFERENCE_SEND_MESSAGES[-1] = "TE-1: _ERR_CONFERENCE_SEND_MESSAGE"}
const ERR_CONFERENCE_SEND_MESSAGE_OK = int(C.TOX_ERR_CONFERENCE_SEND_MESSAGE_OK) // 0
func init(){_ERR_CONFERENCE_SEND_MESSAGES[ERR_CONFERENCE_SEND_MESSAGE_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_SEND_MESSAGE_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_SEND_MESSAGE_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_SEND_MESSAGES[ERR_CONFERENCE_SEND_MESSAGE_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}
const ERR_CONFERENCE_SEND_MESSAGE_TOO_LONG = int(C.TOX_ERR_CONFERENCE_SEND_MESSAGE_TOO_LONG) // 2
func init(){_ERR_CONFERENCE_SEND_MESSAGES[ERR_CONFERENCE_SEND_MESSAGE_TOO_LONG] = "TE02: The message is too long."}
const ERR_CONFERENCE_SEND_MESSAGE_NO_CONNECTION = int(C.TOX_ERR_CONFERENCE_SEND_MESSAGE_NO_CONNECTION) // 3
func init(){_ERR_CONFERENCE_SEND_MESSAGES[ERR_CONFERENCE_SEND_MESSAGE_NO_CONNECTION] = "TE03: The client is not connected to the conference."}
const ERR_CONFERENCE_SEND_MESSAGE_FAIL_SEND = int(C.TOX_ERR_CONFERENCE_SEND_MESSAGE_FAIL_SEND) // 4
func init(){_ERR_CONFERENCE_SEND_MESSAGES[ERR_CONFERENCE_SEND_MESSAGE_FAIL_SEND] = "TE04: The message packet failed to send."}

var _ERR_CONFERENCE_TITLES = make(map[int]string)
func init(){_ERR_CONFERENCE_TITLES[-1] = "TE-1: _ERR_CONFERENCE_TITLE"}
const ERR_CONFERENCE_TITLE_OK = int(C.TOX_ERR_CONFERENCE_TITLE_OK) // 0
func init(){_ERR_CONFERENCE_TITLES[ERR_CONFERENCE_TITLE_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_TITLE_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_TITLE_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_TITLES[ERR_CONFERENCE_TITLE_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}
const ERR_CONFERENCE_TITLE_INVALID_LENGTH = int(C.TOX_ERR_CONFERENCE_TITLE_INVALID_LENGTH) // 2
func init(){_ERR_CONFERENCE_TITLES[ERR_CONFERENCE_TITLE_INVALID_LENGTH] = "TE02: The title is too long or empty."}
const ERR_CONFERENCE_TITLE_FAIL_SEND = int(C.TOX_ERR_CONFERENCE_TITLE_FAIL_SEND) // 3
func init(){_ERR_CONFERENCE_TITLES[ERR_CONFERENCE_TITLE_FAIL_SEND] = "TE03: The title packet failed to send."}

var _ERR_CONFERENCE_GET_TYPES = make(map[int]string)
func init(){_ERR_CONFERENCE_GET_TYPES[-1] = "TE-1: _ERR_CONFERENCE_GET_TYPE"}
const ERR_CONFERENCE_GET_TYPE_OK = int(C.TOX_ERR_CONFERENCE_GET_TYPE_OK) // 0
func init(){_ERR_CONFERENCE_GET_TYPES[ERR_CONFERENCE_GET_TYPE_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_GET_TYPE_CONFERENCE_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_GET_TYPE_CONFERENCE_NOT_FOUND) // 1
func init(){_ERR_CONFERENCE_GET_TYPES[ERR_CONFERENCE_GET_TYPE_CONFERENCE_NOT_FOUND] = "TE01: The conference number passed did not designate a valid conference."}

var _ERR_CONFERENCE_BY_IDS = make(map[int]string)
func init(){_ERR_CONFERENCE_BY_IDS[-1] = "TE-1: _ERR_CONFERENCE_BY_ID"}
const ERR_CONFERENCE_BY_ID_OK = int(C.TOX_ERR_CONFERENCE_BY_ID_OK) // 0
func init(){_ERR_CONFERENCE_BY_IDS[ERR_CONFERENCE_BY_ID_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_BY_ID_NULL = int(C.TOX_ERR_CONFERENCE_BY_ID_NULL) // 1
func init(){_ERR_CONFERENCE_BY_IDS[ERR_CONFERENCE_BY_ID_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_CONFERENCE_BY_ID_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_BY_ID_NOT_FOUND) // 2
func init(){_ERR_CONFERENCE_BY_IDS[ERR_CONFERENCE_BY_ID_NOT_FOUND] = "TE02: No conference with the given id exists on the conference list."}

var _ERR_CONFERENCE_BY_UIDS = make(map[int]string)
func init(){_ERR_CONFERENCE_BY_UIDS[-1] = "TE-1: _ERR_CONFERENCE_BY_UID"}
const ERR_CONFERENCE_BY_UID_OK = int(C.TOX_ERR_CONFERENCE_BY_UID_OK) // 0
func init(){_ERR_CONFERENCE_BY_UIDS[ERR_CONFERENCE_BY_UID_OK] = "TE00: The function returned successfully."}
const ERR_CONFERENCE_BY_UID_NULL = int(C.TOX_ERR_CONFERENCE_BY_UID_NULL) // 1
func init(){_ERR_CONFERENCE_BY_UIDS[ERR_CONFERENCE_BY_UID_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_CONFERENCE_BY_UID_NOT_FOUND = int(C.TOX_ERR_CONFERENCE_BY_UID_NOT_FOUND) // 2
func init(){_ERR_CONFERENCE_BY_UIDS[ERR_CONFERENCE_BY_UID_NOT_FOUND] = "TE02: No conference with the given uid exists on the conference list."}

var _ERR_FRIEND_CUSTOM_PACKETS = make(map[int]string)
func init(){_ERR_FRIEND_CUSTOM_PACKETS[-1] = "TE-1: _ERR_FRIEND_CUSTOM_PACKET"}
const ERR_FRIEND_CUSTOM_PACKET_OK = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_OK) // 0
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_OK] = "TE00: The function returned successfully."}
const ERR_FRIEND_CUSTOM_PACKET_NULL = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_NULL) // 1
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_FOUND = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_FOUND) // 2
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_FOUND] = "TE02: The friend number did not designate a valid friend."}
const ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_CONNECTED = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_CONNECTED) // 3
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_FRIEND_NOT_CONNECTED] = "TE03: This client is currently not connected to the friend."}
const ERR_FRIEND_CUSTOM_PACKET_INVALID = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_INVALID) // 4
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_INVALID] = "TE04: The first byte of data was not in the specified range for the packet type. This range is 200-254 for lossy, and 160-191 for lossless packets."}
const ERR_FRIEND_CUSTOM_PACKET_EMPTY = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_EMPTY) // 5
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_EMPTY] = "TE05: Attempted to send an empty packet."}
const ERR_FRIEND_CUSTOM_PACKET_TOO_LONG = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_TOO_LONG) // 6
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_TOO_LONG] = "TE06: Packet data length exceeded TOX_MAX_CUSTOM_PACKET_SIZE."}
const ERR_FRIEND_CUSTOM_PACKET_SENDQ = int(C.TOX_ERR_FRIEND_CUSTOM_PACKET_SENDQ) // 7
func init(){_ERR_FRIEND_CUSTOM_PACKETS[ERR_FRIEND_CUSTOM_PACKET_SENDQ] = "TE07: Packet queue is full."}

var _ERR_GET_PORTS = make(map[int]string)
func init(){_ERR_GET_PORTS[-1] = "TE-1: _ERR_GET_PORT"}
const ERR_GET_PORT_OK = int(C.TOX_ERR_GET_PORT_OK) // 0
func init(){_ERR_GET_PORTS[ERR_GET_PORT_OK] = "TE00: The function returned successfully."}
const ERR_GET_PORT_NOT_BOUND = int(C.TOX_ERR_GET_PORT_NOT_BOUND) // 1
func init(){_ERR_GET_PORTS[ERR_GET_PORT_NOT_BOUND] = "TE01: The instance was not bound to any port."}

// Group Chat 错误码
var _ERR_GROUP_NEWS = make(map[int]string)
func init(){_ERR_GROUP_NEWS[-1] = "TE-1: _ERR_GROUP_NEW"}
const ERR_GROUP_NEW_OK = int(C.TOX_ERR_GROUP_NEW_OK) // 0
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_NEW_TOO_LONG = int(C.TOX_ERR_GROUP_NEW_TOO_LONG) // 1
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_TOO_LONG] = "TE01: The group name exceeded the maximum allowed length."}
const ERR_GROUP_NEW_EMPTY = int(C.TOX_ERR_GROUP_NEW_EMPTY) // 2
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_EMPTY] = "TE02: The group name was empty."}
const ERR_GROUP_NEW_INIT = int(C.TOX_ERR_GROUP_NEW_INIT) // 3
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_INIT] = "TE03: The group instance failed to initialize."}
const ERR_GROUP_NEW_STATE = int(C.TOX_ERR_GROUP_NEW_STATE) // 4
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_STATE] = "TE04: The group state failed to initialize."}
const ERR_GROUP_NEW_ANNOUNCE = int(C.TOX_ERR_GROUP_NEW_ANNOUNCE) // 5
func init(){_ERR_GROUP_NEWS[ERR_GROUP_NEW_ANNOUNCE] = "TE05: The group announce failed to initialize."}

var _ERR_GROUP_JOINS = make(map[int]string)
func init(){_ERR_GROUP_JOINS[-1] = "TE-1: _ERR_GROUP_JOIN"}
const ERR_GROUP_JOIN_OK = int(C.TOX_ERR_GROUP_JOIN_OK) // 0
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_JOIN_INIT = int(C.TOX_ERR_GROUP_JOIN_INIT) // 1
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_INIT] = "TE01: The group instance failed to initialize."}
const ERR_GROUP_JOIN_BAD_CHAT_ID = int(C.TOX_ERR_GROUP_JOIN_BAD_CHAT_ID) // 2
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_BAD_CHAT_ID] = "TE02: The chat ID passed was invalid."}
const ERR_GROUP_JOIN_EMPTY = int(C.TOX_ERR_GROUP_JOIN_EMPTY) // 3
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_EMPTY] = "TE03: The chat ID was empty."}
const ERR_GROUP_JOIN_TOO_LONG = int(C.TOX_ERR_GROUP_JOIN_TOO_LONG) // 4
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_TOO_LONG] = "TE04: The password exceeded the maximum allowed length."}
const ERR_GROUP_JOIN_PASSWORD = int(C.TOX_ERR_GROUP_JOIN_PASSWORD) // 5
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_PASSWORD] = "TE05: The password did not match the group password."}
const ERR_GROUP_JOIN_CORE = int(C.TOX_ERR_GROUP_JOIN_CORE) // 6
func init(){_ERR_GROUP_JOINS[ERR_GROUP_JOIN_CORE] = "TE06: A core error occurred."}

var _ERR_GROUP_IS_CONNECTEDS = make(map[int]string)
func init(){_ERR_GROUP_IS_CONNECTEDS[-1] = "TE-1: _ERR_GROUP_IS_CONNECTED"}
const ERR_GROUP_IS_CONNECTED_OK = int(C.TOX_ERR_GROUP_IS_CONNECTED_OK) // 0
func init(){_ERR_GROUP_IS_CONNECTEDS[ERR_GROUP_IS_CONNECTED_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_IS_CONNECTED_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_IS_CONNECTED_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_IS_CONNECTEDS[ERR_GROUP_IS_CONNECTED_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}

var _ERR_GROUP_DISCONNECTS = make(map[int]string)
func init(){_ERR_GROUP_DISCONNECTS[-1] = "TE-1: _ERR_GROUP_DISCONNECT"}
const ERR_GROUP_DISCONNECT_OK = int(C.TOX_ERR_GROUP_DISCONNECT_OK) // 0
func init(){_ERR_GROUP_DISCONNECTS[ERR_GROUP_DISCONNECT_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_DISCONNECT_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_DISCONNECT_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_DISCONNECTS[ERR_GROUP_DISCONNECT_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_DISCONNECT_ALREADY_DISCONNECTED = int(C.TOX_ERR_GROUP_DISCONNECT_ALREADY_DISCONNECTED) // 2
func init(){_ERR_GROUP_DISCONNECTS[ERR_GROUP_DISCONNECT_ALREADY_DISCONNECTED] = "TE02: The client is already disconnected from the group."}

var _ERR_GROUP_RECONNECTS = make(map[int]string)
func init(){_ERR_GROUP_RECONNECTS[-1] = "TE-1: _ERR_GROUP_RECONNECT"}
const ERR_GROUP_RECONNECT_OK = int(C.TOX_ERR_GROUP_RECONNECT_OK) // 0
func init(){_ERR_GROUP_RECONNECTS[ERR_GROUP_RECONNECT_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_RECONNECT_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_RECONNECT_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_RECONNECTS[ERR_GROUP_RECONNECT_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_RECONNECT_CORE = int(C.TOX_ERR_GROUP_RECONNECT_CORE) // 2
func init(){_ERR_GROUP_RECONNECTS[ERR_GROUP_RECONNECT_CORE] = "TE02: A core error occurred."}

var _ERR_GROUP_LEAVES = make(map[int]string)
func init(){_ERR_GROUP_LEAVES[-1] = "TE-1: _ERR_GROUP_LEAVE"}
const ERR_GROUP_LEAVE_OK = int(C.TOX_ERR_GROUP_LEAVE_OK) // 0
func init(){_ERR_GROUP_LEAVES[ERR_GROUP_LEAVE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_LEAVE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_LEAVE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_LEAVES[ERR_GROUP_LEAVE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_LEAVE_TOO_LONG = int(C.TOX_ERR_GROUP_LEAVE_TOO_LONG) // 2
func init(){_ERR_GROUP_LEAVES[ERR_GROUP_LEAVE_TOO_LONG] = "TE02: The parting message exceeded the maximum allowed length."}
const ERR_GROUP_LEAVE_FAIL_SEND = int(C.TOX_ERR_GROUP_LEAVE_FAIL_SEND) // 3
func init(){_ERR_GROUP_LEAVES[ERR_GROUP_LEAVE_FAIL_SEND] = "TE03: The parting packet failed to send."}

var _ERR_GROUP_SELF_QUERYS = make(map[int]string)
func init(){_ERR_GROUP_SELF_QUERYS[-1] = "TE-1: _ERR_GROUP_SELF_QUERY"}
const ERR_GROUP_SELF_QUERY_OK = int(C.TOX_ERR_GROUP_SELF_QUERY_OK) // 0
func init(){_ERR_GROUP_SELF_QUERYS[ERR_GROUP_SELF_QUERY_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SELF_QUERY_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SELF_QUERY_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SELF_QUERYS[ERR_GROUP_SELF_QUERY_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}

var _ERR_GROUP_SELF_NAME_SETS = make(map[int]string)
func init(){_ERR_GROUP_SELF_NAME_SETS[-1] = "TE-1: _ERR_GROUP_SELF_NAME_SET"}
const ERR_GROUP_SELF_NAME_SET_OK = int(C.TOX_ERR_GROUP_SELF_NAME_SET_OK) // 0
func init(){_ERR_GROUP_SELF_NAME_SETS[ERR_GROUP_SELF_NAME_SET_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SELF_NAME_SET_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SELF_NAME_SET_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SELF_NAME_SETS[ERR_GROUP_SELF_NAME_SET_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SELF_NAME_SET_TOO_LONG = int(C.TOX_ERR_GROUP_SELF_NAME_SET_TOO_LONG) // 2
func init(){_ERR_GROUP_SELF_NAME_SETS[ERR_GROUP_SELF_NAME_SET_TOO_LONG] = "TE02: The name exceeded the maximum allowed length."}
const ERR_GROUP_SELF_NAME_SET_INVALID = int(C.TOX_ERR_GROUP_SELF_NAME_SET_INVALID) // 3
func init(){_ERR_GROUP_SELF_NAME_SETS[ERR_GROUP_SELF_NAME_SET_INVALID] = "TE03: The name was empty."}
const ERR_GROUP_SELF_NAME_SET_FAIL_SEND = int(C.TOX_ERR_GROUP_SELF_NAME_SET_FAIL_SEND) // 4
func init(){_ERR_GROUP_SELF_NAME_SETS[ERR_GROUP_SELF_NAME_SET_FAIL_SEND] = "TE04: The name packet failed to send."}

var _ERR_GROUP_SELF_STATUS_SETS = make(map[int]string)
func init(){_ERR_GROUP_SELF_STATUS_SETS[-1] = "TE-1: _ERR_GROUP_SELF_STATUS_SET"}
const ERR_GROUP_SELF_STATUS_SET_OK = int(C.TOX_ERR_GROUP_SELF_STATUS_SET_OK) // 0
func init(){_ERR_GROUP_SELF_STATUS_SETS[ERR_GROUP_SELF_STATUS_SET_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SELF_STATUS_SET_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SELF_STATUS_SET_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SELF_STATUS_SETS[ERR_GROUP_SELF_STATUS_SET_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SELF_STATUS_SET_FAIL_SEND = int(C.TOX_ERR_GROUP_SELF_STATUS_SET_FAIL_SEND) // 2
func init(){_ERR_GROUP_SELF_STATUS_SETS[ERR_GROUP_SELF_STATUS_SET_FAIL_SEND] = "TE02: The status packet failed to send."}

var _ERR_GROUP_PEER_QUERYS = make(map[int]string)
func init(){_ERR_GROUP_PEER_QUERYS[-1] = "TE-1: _ERR_GROUP_PEER_QUERY"}
const ERR_GROUP_PEER_QUERY_OK = int(C.TOX_ERR_GROUP_PEER_QUERY_OK) // 0
func init(){_ERR_GROUP_PEER_QUERYS[ERR_GROUP_PEER_QUERY_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_PEER_QUERY_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_PEER_QUERY_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_PEER_QUERYS[ERR_GROUP_PEER_QUERY_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_PEER_QUERY_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_PEER_QUERY_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_PEER_QUERYS[ERR_GROUP_PEER_QUERY_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}

var _ERR_GROUP_STATE_QUERYS = make(map[int]string)
func init(){_ERR_GROUP_STATE_QUERYS[-1] = "TE-1: _ERR_GROUP_STATE_QUERY"}
const ERR_GROUP_STATE_QUERY_OK = int(C.TOX_ERR_GROUP_STATE_QUERY_OK) // 0
func init(){_ERR_GROUP_STATE_QUERYS[ERR_GROUP_STATE_QUERY_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_STATE_QUERY_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_STATE_QUERY_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_STATE_QUERYS[ERR_GROUP_STATE_QUERY_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}

var _ERR_GROUP_TOPIC_SETS = make(map[int]string)
func init(){_ERR_GROUP_TOPIC_SETS[-1] = "TE-1: _ERR_GROUP_TOPIC_SET"}
const ERR_GROUP_TOPIC_SET_OK = int(C.TOX_ERR_GROUP_TOPIC_SET_OK) // 0
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_TOPIC_SET_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_TOPIC_SET_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_TOPIC_SET_TOO_LONG = int(C.TOX_ERR_GROUP_TOPIC_SET_TOO_LONG) // 2
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_TOO_LONG] = "TE02: The topic exceeded the maximum allowed length."}
const ERR_GROUP_TOPIC_SET_PERMISSIONS = int(C.TOX_ERR_GROUP_TOPIC_SET_PERMISSIONS) // 3
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_PERMISSIONS] = "TE03: The client does not have permissions to set the topic."}
const ERR_GROUP_TOPIC_SET_FAIL_CREATE = int(C.TOX_ERR_GROUP_TOPIC_SET_FAIL_CREATE) // 4
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_FAIL_CREATE] = "TE04: The topic packet failed to create."}
const ERR_GROUP_TOPIC_SET_FAIL_SEND = int(C.TOX_ERR_GROUP_TOPIC_SET_FAIL_SEND) // 5
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_FAIL_SEND] = "TE05: The topic packet failed to send."}
const ERR_GROUP_TOPIC_SET_DISCONNECTED = int(C.TOX_ERR_GROUP_TOPIC_SET_DISCONNECTED) // 6
func init(){_ERR_GROUP_TOPIC_SETS[ERR_GROUP_TOPIC_SET_DISCONNECTED] = "TE06: The client is disconnected from the group."}

var _ERR_GROUP_BY_IDS = make(map[int]string)
func init(){_ERR_GROUP_BY_IDS[-1] = "TE-1: _ERR_GROUP_BY_ID"}
const ERR_GROUP_BY_ID_OK = int(C.TOX_ERR_GROUP_BY_ID_OK) // 0
func init(){_ERR_GROUP_BY_IDS[ERR_GROUP_BY_ID_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_BY_ID_NULL = int(C.TOX_ERR_GROUP_BY_ID_NULL) // 1
func init(){_ERR_GROUP_BY_IDS[ERR_GROUP_BY_ID_NULL] = "TE01: One of the arguments to the function was NULL when it was not expected."}
const ERR_GROUP_BY_ID_NOT_FOUND = int(C.TOX_ERR_GROUP_BY_ID_NOT_FOUND) // 2
func init(){_ERR_GROUP_BY_IDS[ERR_GROUP_BY_ID_NOT_FOUND] = "TE02: No group with the given chat ID exists on the group list."}

var _ERR_GROUP_SEND_MESSAGES = make(map[int]string)
func init(){_ERR_GROUP_SEND_MESSAGES[-1] = "TE-1: _ERR_GROUP_SEND_MESSAGE"}
const ERR_GROUP_SEND_MESSAGE_OK = int(C.TOX_ERR_GROUP_SEND_MESSAGE_OK) // 0
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SEND_MESSAGE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_MESSAGE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SEND_MESSAGE_TOO_LONG = int(C.TOX_ERR_GROUP_SEND_MESSAGE_TOO_LONG) // 2
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_TOO_LONG] = "TE02: The message exceeded the maximum allowed length."}
const ERR_GROUP_SEND_MESSAGE_EMPTY = int(C.TOX_ERR_GROUP_SEND_MESSAGE_EMPTY) // 3
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_EMPTY] = "TE03: The message was empty."}
const ERR_GROUP_SEND_MESSAGE_BAD_TYPE = int(C.TOX_ERR_GROUP_SEND_MESSAGE_BAD_TYPE) // 4
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_BAD_TYPE] = "TE04: The message type was invalid."}
const ERR_GROUP_SEND_MESSAGE_PERMISSIONS = int(C.TOX_ERR_GROUP_SEND_MESSAGE_PERMISSIONS) // 5
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_PERMISSIONS] = "TE05: The client does not have permissions to send messages."}
const ERR_GROUP_SEND_MESSAGE_FAIL_SEND = int(C.TOX_ERR_GROUP_SEND_MESSAGE_FAIL_SEND) // 6
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_FAIL_SEND] = "TE06: The message packet failed to send."}
const ERR_GROUP_SEND_MESSAGE_DISCONNECTED = int(C.TOX_ERR_GROUP_SEND_MESSAGE_DISCONNECTED) // 7
func init(){_ERR_GROUP_SEND_MESSAGES[ERR_GROUP_SEND_MESSAGE_DISCONNECTED] = "TE07: The client is disconnected from the group."}

var _ERR_GROUP_SEND_PRIVATE_MESSAGES = make(map[int]string)
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[-1] = "TE-1: _ERR_GROUP_SEND_PRIVATE_MESSAGE"}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_OK = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_OK) // 0
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_TOO_LONG = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_TOO_LONG) // 3
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_TOO_LONG] = "TE03: The message exceeded the maximum allowed length."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_EMPTY = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_EMPTY) // 4
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_EMPTY] = "TE04: The message was empty."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_BAD_TYPE = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_BAD_TYPE) // 5
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_BAD_TYPE] = "TE05: The message type was invalid."}
const ERR_GROUP_SEND_PRIVATE_MESSAGE_PERMISSIONS = int(C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_PERMISSIONS) // 6
func init(){_ERR_GROUP_SEND_PRIVATE_MESSAGES[ERR_GROUP_SEND_PRIVATE_MESSAGE_PERMISSIONS] = "TE06: The client does not have permissions to send private messages."}

var _ERR_GROUP_SEND_CUSTOM_PACKETS = make(map[int]string)
func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[-1] = "TE-1: _ERR_GROUP_SEND_CUSTOM_PACKET"}
const ERR_GROUP_SEND_CUSTOM_PACKET_OK = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PACKET_OK) // 0
func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[ERR_GROUP_SEND_CUSTOM_PACKET_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SEND_CUSTOM_PACKET_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PACKET_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[ERR_GROUP_SEND_CUSTOM_PACKET_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SEND_CUSTOM_PACKET_TOO_LONG = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PACKET_TOO_LONG) // 2
func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[ERR_GROUP_SEND_CUSTOM_PACKET_TOO_LONG] = "TE02: The packet data exceeded the maximum allowed length."}
// ERR_GROUP_SEND_CUSTOM_PACKET_PERMISSIONS not in current tox.h
// const ERR_GROUP_SEND_CUSTOM_PACKET_PERMISSIONS = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PACKET_PERMISSIONS) // 3
// func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[ERR_GROUP_SEND_CUSTOM_PACKET_PERMISSIONS] = "TE03: The client does not have permissions to send custom packets."}
const ERR_GROUP_SEND_CUSTOM_PACKET_FAIL_SEND = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PACKET_FAIL_SEND) // 4
func init(){_ERR_GROUP_SEND_CUSTOM_PACKETS[ERR_GROUP_SEND_CUSTOM_PACKET_FAIL_SEND] = "TE04: The packet failed to send."}

var _ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS = make(map[int]string)
func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[-1] = "TE-1: _ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET"}
const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_OK = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_OK) // 0
func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_TOO_LONG = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_TOO_LONG) // 3
func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_TOO_LONG] = "TE03: The packet data exceeded the maximum allowed length."}
// ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_INVALID not in current tox.h
// const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_INVALID = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_INVALID) // 4
// func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_INVALID] = "TE04: The first byte of data was not in the specified range."}
// ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PERMISSIONS not in current tox.h
// const ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PERMISSIONS = int(C.TOX_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PERMISSIONS) // 5
// func init(){_ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKETS[ERR_GROUP_SEND_CUSTOM_PRIVATE_PACKET_PERMISSIONS] = "TE05: The client does not have permissions to send custom packets."}

var _ERR_GROUP_INVITE_FRIENDS = make(map[int]string)
func init(){_ERR_GROUP_INVITE_FRIENDS[-1] = "TE-1: _ERR_GROUP_INVITE_FRIEND"}
const ERR_GROUP_INVITE_FRIEND_OK = int(C.TOX_ERR_GROUP_INVITE_FRIEND_OK) // 0
func init(){_ERR_GROUP_INVITE_FRIENDS[ERR_GROUP_INVITE_FRIEND_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_INVITE_FRIEND_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_INVITE_FRIEND_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_INVITE_FRIENDS[ERR_GROUP_INVITE_FRIEND_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_INVITE_FRIEND_FRIEND_NOT_FOUND = int(C.TOX_ERR_GROUP_INVITE_FRIEND_FRIEND_NOT_FOUND) // 2
func init(){_ERR_GROUP_INVITE_FRIENDS[ERR_GROUP_INVITE_FRIEND_FRIEND_NOT_FOUND] = "TE02: The friend number passed did not designate a valid friend."}
// ERR_GROUP_INVITE_FRIEND_INVITE_EXISTS not in current tox.h
// const ERR_GROUP_INVITE_FRIEND_INVITE_EXISTS = int(C.TOX_ERR_GROUP_INVITE_FRIEND_INVITE_EXISTS) // 3
// func init(){_ERR_GROUP_INVITE_FRIENDS[ERR_GROUP_INVITE_FRIEND_INVITE_EXISTS] = "TE03: The friend has already been invited to the group."}
const ERR_GROUP_INVITE_FRIEND_FAIL_SEND = int(C.TOX_ERR_GROUP_INVITE_FRIEND_FAIL_SEND) // 4
func init(){_ERR_GROUP_INVITE_FRIENDS[ERR_GROUP_INVITE_FRIEND_FAIL_SEND] = "TE04: The invite packet failed to send."}

var _ERR_GROUP_INVITE_ACCEPTS = make(map[int]string)
func init(){_ERR_GROUP_INVITE_ACCEPTS[-1] = "TE-1: _ERR_GROUP_INVITE_ACCEPT"}
const ERR_GROUP_INVITE_ACCEPT_OK = int(C.TOX_ERR_GROUP_INVITE_ACCEPT_OK) // 0
func init(){_ERR_GROUP_INVITE_ACCEPTS[ERR_GROUP_INVITE_ACCEPT_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_INVITE_ACCEPT_BAD_INVITE = int(C.TOX_ERR_GROUP_INVITE_ACCEPT_BAD_INVITE) // 1
func init(){_ERR_GROUP_INVITE_ACCEPTS[ERR_GROUP_INVITE_ACCEPT_BAD_INVITE] = "TE01: The invite data was invalid."}

var _ERR_GROUP_SET_PASSWORDS = make(map[int]string)
func init(){_ERR_GROUP_SET_PASSWORDS[-1] = "TE-1: _ERR_GROUP_SET_PASSWORD"}
const ERR_GROUP_SET_PASSWORD_OK = int(C.TOX_ERR_GROUP_SET_PASSWORD_OK) // 0
func init(){_ERR_GROUP_SET_PASSWORDS[ERR_GROUP_SET_PASSWORD_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_PASSWORD_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_PASSWORD_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_PASSWORDS[ERR_GROUP_SET_PASSWORD_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_PASSWORD_TOO_LONG = int(C.TOX_ERR_GROUP_SET_PASSWORD_TOO_LONG) // 2
func init(){_ERR_GROUP_SET_PASSWORDS[ERR_GROUP_SET_PASSWORD_TOO_LONG] = "TE02: The password exceeded the maximum allowed length."}
const ERR_GROUP_SET_PASSWORD_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_PASSWORD_PERMISSIONS) // 3
func init(){_ERR_GROUP_SET_PASSWORDS[ERR_GROUP_SET_PASSWORD_PERMISSIONS] = "TE03: The client does not have permissions to set the password."}
const ERR_GROUP_SET_PASSWORD_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_PASSWORD_FAIL_SEND) // 4
func init(){_ERR_GROUP_SET_PASSWORDS[ERR_GROUP_SET_PASSWORD_FAIL_SEND] = "TE04: The password packet failed to send."}

var _ERR_GROUP_SET_TOPIC_LOCKS = make(map[int]string)
func init(){_ERR_GROUP_SET_TOPIC_LOCKS[-1] = "TE-1: _ERR_GROUP_SET_TOPIC_LOCK"}
const ERR_GROUP_SET_TOPIC_LOCK_OK = int(C.TOX_ERR_GROUP_SET_TOPIC_LOCK_OK) // 0
func init(){_ERR_GROUP_SET_TOPIC_LOCKS[ERR_GROUP_SET_TOPIC_LOCK_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_TOPIC_LOCK_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_TOPIC_LOCK_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_TOPIC_LOCKS[ERR_GROUP_SET_TOPIC_LOCK_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_TOPIC_LOCK_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_TOPIC_LOCK_PERMISSIONS) // 2
func init(){_ERR_GROUP_SET_TOPIC_LOCKS[ERR_GROUP_SET_TOPIC_LOCK_PERMISSIONS] = "TE02: The client does not have permissions to set the topic lock."}
const ERR_GROUP_SET_TOPIC_LOCK_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_TOPIC_LOCK_FAIL_SEND) // 3
func init(){_ERR_GROUP_SET_TOPIC_LOCKS[ERR_GROUP_SET_TOPIC_LOCK_FAIL_SEND] = "TE03: The topic lock packet failed to send."}

var _ERR_GROUP_SET_VOICE_STATES = make(map[int]string)
func init(){_ERR_GROUP_SET_VOICE_STATES[-1] = "TE-1: _ERR_GROUP_SET_VOICE_STATE"}
const ERR_GROUP_SET_VOICE_STATE_OK = int(C.TOX_ERR_GROUP_SET_VOICE_STATE_OK) // 0
func init(){_ERR_GROUP_SET_VOICE_STATES[ERR_GROUP_SET_VOICE_STATE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_VOICE_STATE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_VOICE_STATE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_VOICE_STATES[ERR_GROUP_SET_VOICE_STATE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_VOICE_STATE_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_VOICE_STATE_PERMISSIONS) // 2
func init(){_ERR_GROUP_SET_VOICE_STATES[ERR_GROUP_SET_VOICE_STATE_PERMISSIONS] = "TE02: The client does not have permissions to set the voice state."}
const ERR_GROUP_SET_VOICE_STATE_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_VOICE_STATE_FAIL_SEND) // 3
func init(){_ERR_GROUP_SET_VOICE_STATES[ERR_GROUP_SET_VOICE_STATE_FAIL_SEND] = "TE03: The voice state packet failed to send."}

var _ERR_GROUP_SET_PRIVACY_STATES = make(map[int]string)
func init(){_ERR_GROUP_SET_PRIVACY_STATES[-1] = "TE-1: _ERR_GROUP_SET_PRIVACY_STATE"}
const ERR_GROUP_SET_PRIVACY_STATE_OK = int(C.TOX_ERR_GROUP_SET_PRIVACY_STATE_OK) // 0
func init(){_ERR_GROUP_SET_PRIVACY_STATES[ERR_GROUP_SET_PRIVACY_STATE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_PRIVACY_STATE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_PRIVACY_STATE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_PRIVACY_STATES[ERR_GROUP_SET_PRIVACY_STATE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_PRIVACY_STATE_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_PRIVACY_STATE_PERMISSIONS) // 2
func init(){_ERR_GROUP_SET_PRIVACY_STATES[ERR_GROUP_SET_PRIVACY_STATE_PERMISSIONS] = "TE02: The client does not have permissions to set the privacy state."}
const ERR_GROUP_SET_PRIVACY_STATE_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_PRIVACY_STATE_FAIL_SEND) // 3
func init(){_ERR_GROUP_SET_PRIVACY_STATES[ERR_GROUP_SET_PRIVACY_STATE_FAIL_SEND] = "TE03: The privacy state packet failed to send."}

var _ERR_GROUP_SET_PEER_LIMITS = make(map[int]string)
func init(){_ERR_GROUP_SET_PEER_LIMITS[-1] = "TE-1: _ERR_GROUP_SET_PEER_LIMIT"}
const ERR_GROUP_SET_PEER_LIMIT_OK = int(C.TOX_ERR_GROUP_SET_PEER_LIMIT_OK) // 0
func init(){_ERR_GROUP_SET_PEER_LIMITS[ERR_GROUP_SET_PEER_LIMIT_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_PEER_LIMIT_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_PEER_LIMIT_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_PEER_LIMITS[ERR_GROUP_SET_PEER_LIMIT_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_PEER_LIMIT_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_PEER_LIMIT_PERMISSIONS) // 2
func init(){_ERR_GROUP_SET_PEER_LIMITS[ERR_GROUP_SET_PEER_LIMIT_PERMISSIONS] = "TE02: The client does not have permissions to set the peer limit."}
const ERR_GROUP_SET_PEER_LIMIT_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_PEER_LIMIT_FAIL_SEND) // 3
func init(){_ERR_GROUP_SET_PEER_LIMITS[ERR_GROUP_SET_PEER_LIMIT_FAIL_SEND] = "TE03: The peer limit packet failed to send."}

var _ERR_GROUP_SET_IGNORES = make(map[int]string)
func init(){_ERR_GROUP_SET_IGNORES[-1] = "TE-1: _ERR_GROUP_SET_IGNORE"}
const ERR_GROUP_SET_IGNORE_OK = int(C.TOX_ERR_GROUP_SET_IGNORE_OK) // 0
func init(){_ERR_GROUP_SET_IGNORES[ERR_GROUP_SET_IGNORE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_IGNORE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_IGNORE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_IGNORES[ERR_GROUP_SET_IGNORE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_IGNORE_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_IGNORE_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_SET_IGNORES[ERR_GROUP_SET_IGNORE_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
// ERR_GROUP_SET_IGNORE_FAIL_SEND not in current tox.h
// const ERR_GROUP_SET_IGNORE_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_IGNORE_FAIL_SEND) // 3
// func init(){_ERR_GROUP_SET_IGNORES[ERR_GROUP_SET_IGNORE_FAIL_SEND] = "TE03: The ignore packet failed to send."}

var _ERR_GROUP_SET_ROLES = make(map[int]string)
func init(){_ERR_GROUP_SET_ROLES[-1] = "TE-1: _ERR_GROUP_SET_ROLE"}
const ERR_GROUP_SET_ROLE_OK = int(C.TOX_ERR_GROUP_SET_ROLE_OK) // 0
func init(){_ERR_GROUP_SET_ROLES[ERR_GROUP_SET_ROLE_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_SET_ROLE_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_ROLE_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_SET_ROLES[ERR_GROUP_SET_ROLE_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_SET_ROLE_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_SET_ROLE_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_SET_ROLES[ERR_GROUP_SET_ROLE_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
const ERR_GROUP_SET_ROLE_PERMISSIONS = int(C.TOX_ERR_GROUP_SET_ROLE_PERMISSIONS) // 3
func init(){_ERR_GROUP_SET_ROLES[ERR_GROUP_SET_ROLE_PERMISSIONS] = "TE03: The client does not have permissions to set the role."}
// ERR_GROUP_SET_ROLE_FAIL_SEND not in current tox.h
// const ERR_GROUP_SET_ROLE_FAIL_SEND = int(C.TOX_ERR_GROUP_SET_ROLE_FAIL_SEND) // 4
// func init(){_ERR_GROUP_SET_ROLES[ERR_GROUP_SET_ROLE_FAIL_SEND] = "TE04: The role packet failed to send."}

var _ERR_GROUP_KICK_PEERS = make(map[int]string)
func init(){_ERR_GROUP_KICK_PEERS[-1] = "TE-1: _ERR_GROUP_KICK_PEER"}
const ERR_GROUP_KICK_PEER_OK = int(C.TOX_ERR_GROUP_KICK_PEER_OK) // 0
func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_OK] = "TE00: The function returned successfully."}
const ERR_GROUP_KICK_PEER_GROUP_NOT_FOUND = int(C.TOX_ERR_GROUP_KICK_PEER_GROUP_NOT_FOUND) // 1
func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_GROUP_NOT_FOUND] = "TE01: The group number passed did not designate a valid group."}
const ERR_GROUP_KICK_PEER_PEER_NOT_FOUND = int(C.TOX_ERR_GROUP_KICK_PEER_PEER_NOT_FOUND) // 2
func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_PEER_NOT_FOUND] = "TE02: The peer number passed did not designate a valid peer."}
const ERR_GROUP_KICK_PEER_PERMISSIONS = int(C.TOX_ERR_GROUP_KICK_PEER_PERMISSIONS) // 3
func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_PERMISSIONS] = "TE03: The client does not have permissions to kick peers."}
// ERR_GROUP_KICK_PEER_TOO_LONG not in current tox.h
// const ERR_GROUP_KICK_PEER_TOO_LONG = int(C.TOX_ERR_GROUP_KICK_PEER_TOO_LONG) // 4
// func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_TOO_LONG] = "TE04: The reason message exceeded the maximum allowed length."}
// ERR_GROUP_KICK_PEER_FAIL_SEND not in current tox.h
// const ERR_GROUP_KICK_PEER_FAIL_SEND = int(C.TOX_ERR_GROUP_KICK_PEER_FAIL_SEND) // 5
// func init(){_ERR_GROUP_KICK_PEERS[ERR_GROUP_KICK_PEER_FAIL_SEND] = "TE05: The kick packet failed to send."}
