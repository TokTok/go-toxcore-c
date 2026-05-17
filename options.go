package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

typedef const char cchar_t;
extern void toxCallbackLog(Tox*, Tox_Log_Level, cchar_t*, uint32_t, cchar_t*, cchar_t*, void*);

*/
import "C"
import "unsafe"

const (
	SAVEDATA_TYPE_NONE       = int(C.TOX_SAVEDATA_TYPE_NONE)
	SAVEDATA_TYPE_TOX_SAVE   = int(C.TOX_SAVEDATA_TYPE_TOX_SAVE)
	SAVEDATA_TYPE_SECRET_KEY = int(C.TOX_SAVEDATA_TYPE_SECRET_KEY)
)

const (
	PROXY_TYPE_NONE   = int(C.TOX_PROXY_TYPE_NONE)
	PROXY_TYPE_HTTP   = int(C.TOX_PROXY_TYPE_HTTP)
	PROXY_TYPE_SOCKS5 = int(C.TOX_PROXY_TYPE_SOCKS5)
)

const (
	LOG_LEVEL_TRACE   = int(C.TOX_LOG_LEVEL_TRACE)
	LOG_LEVEL_DEBUG   = int(C.TOX_LOG_LEVEL_DEBUG)
	LOG_LEVEL_INFO    = int(C.TOX_LOG_LEVEL_INFO)
	LOG_LEVEL_WARNING = int(C.TOX_LOG_LEVEL_WARNING)
	LOG_LEVEL_ERROR   = int(C.TOX_LOG_LEVEL_ERROR)
)

type ToxOptions struct {
	Ipv6_enabled            bool
	Udp_enabled             bool
	Proxy_type              int32
	Proxy_host              string
	Proxy_port              uint16
	Savedata_type           int
	Savedata_data           []byte
	Tcp_port                uint16
	Local_discovery_enabled bool
	Start_port              uint16
	End_port                uint16
	Hole_punching_enabled   bool
	ThreadSafe              bool
	LogCallback             func(_ *Tox, level int, file string, line uint32, fname string, msg string)
}

func NewToxOptions() *ToxOptions {
	toxopts := C.tox_options_new(nil)
	defer C.tox_options_free(toxopts)

	opts := new(ToxOptions)
	opts.Ipv6_enabled = bool(C.tox_options_get_ipv6_enabled(toxopts))
	opts.Udp_enabled = bool(C.tox_options_get_udp_enabled(toxopts))
	opts.Proxy_type = int32(C.tox_options_get_proxy_type(toxopts))
	opts.Proxy_port = uint16(C.tox_options_get_proxy_port(toxopts))
	opts.Tcp_port = uint16(C.tox_options_get_tcp_port(toxopts))
	opts.Local_discovery_enabled = bool(C.tox_options_get_local_discovery_enabled(toxopts))
	opts.Start_port = uint16(C.tox_options_get_start_port(toxopts))
	opts.End_port = uint16(C.tox_options_get_end_port(toxopts))
	opts.Hole_punching_enabled = bool(C.tox_options_get_hole_punching_enabled(toxopts))

	return opts
}

func (this *ToxOptions) toCToxOptions() *C.struct_Tox_Options {
	toxopts := C.tox_options_new(nil)
	C.tox_options_set_experimental_owned_data(toxopts, true)
	C.tox_options_set_ipv6_enabled(toxopts, (C._Bool)(this.Ipv6_enabled))
	C.tox_options_set_udp_enabled(toxopts, (C._Bool)(this.Udp_enabled))

	if this.Savedata_data != nil {
		C.tox_options_set_savedata(toxopts, (*C.uint8_t)(&this.Savedata_data[0]), C.size_t(len(this.Savedata_data)))
		C.tox_options_set_savedata_type(toxopts, C.Tox_Savedata_Type(this.Savedata_type))
	}
	C.tox_options_set_tcp_port(toxopts, (C.uint16_t)(this.Tcp_port))

	C.tox_options_set_proxy_type(toxopts, C.Tox_Proxy_Type(this.Proxy_type))
	C.tox_options_set_proxy_port(toxopts, C.uint16_t(this.Proxy_port))
	if len(this.Proxy_host) > 0 {
		C.tox_options_set_proxy_host(toxopts, C.CString(this.Proxy_host))
	}

	C.tox_options_set_local_discovery_enabled(toxopts, C._Bool(this.Local_discovery_enabled))
	C.tox_options_set_start_port(toxopts, C.uint16_t(this.Start_port))
	C.tox_options_set_end_port(toxopts, C.uint16_t(this.End_port))
	C.tox_options_set_hole_punching_enabled(toxopts, C._Bool(this.Hole_punching_enabled))

	C.tox_options_set_log_callback(toxopts, (*C.tox_log_cb)((unsafe.Pointer)(C.toxCallbackLog)))

	return toxopts
}

//export toxCallbackLog
func toxCallbackLog(ctox *C.Tox, level C.Tox_Log_Level, file *C.cchar_t, line C.uint32_t, fname *C.cchar_t, msg *C.cchar_t, userdata unsafe.Pointer) {
	t := cbUserDatas.get(ctox)
	if t != nil && t.opts != nil && t.opts.LogCallback != nil {
		t.opts.LogCallback(t, int(level), C.GoString(file), uint32(line), C.GoString(fname), C.GoString(msg))
	}
}

type BootNode struct {
	Addr   string
	Port   int
	Pubkey string
}
