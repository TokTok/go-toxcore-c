package tox

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

// Callback wrapper declarations
void callbackGroupPeerNameWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, uint8_t*, size_t, void*);
void callbackGroupPeerStatusWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, Tox_User_Status, void*);
void callbackGroupTopicWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, uint8_t*, size_t, void*);
void callbackGroupPrivacyStateWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Privacy_State, void*);
void callbackGroupVoiceStateWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Voice_State, void*);
void callbackGroupTopicLockWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Topic_Lock, void*);
void callbackGroupPeerLimitWrapperForC(Tox*, Tox_Group_Number, uint32_t, void*);
void callbackGroupPasswordWrapperForC(Tox*, Tox_Group_Number, uint8_t*, size_t, void*);
void callbackGroupMessageWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, Tox_Message_Type, uint8_t*, size_t, Tox_Group_Message_Id, void*);
void callbackGroupPrivateMessageWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, Tox_Message_Type, uint8_t*, size_t, Tox_Group_Message_Id, void*);
void callbackGroupCustomPacketWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, uint8_t*, size_t, void*);
void callbackGroupCustomPrivatePacketWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, uint8_t*, size_t, void*);
void callbackGroupInviteWrapperForC(Tox*, uint32_t, uint8_t*, size_t, uint8_t*, size_t, void*);
void callbackGroupPeerJoinWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, void*);
void callbackGroupPeerExitWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, Tox_Group_Exit_Type, uint8_t*, size_t, void*);
void callbackGroupSelfJoinWrapperForC(Tox*, Tox_Group_Number, void*);
void callbackGroupJoinFailWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Join_Fail, void*);
void callbackGroupModerationWrapperForC(Tox*, Tox_Group_Number, Tox_Group_Peer_Number, Tox_Group_Mod_Event, uint8_t*, size_t, void*);

// fix no-use compile warning
static inline __attribute__((__unused__)) void fixnousegroupchat(void) {}
*/
import "C"

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
	"unsafe"
)

type size_t = C.size_t

// Group Chat
type cb_group_peer_name_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, name string, userData interface{})
type cb_group_peer_status_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, status int, userData interface{})
type cb_group_topic_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, topic string, userData interface{})
type cb_group_privacy_state_ftype func(this *Tox, groupNumber GroupNumber, privacyState GroupPrivacyState, userData interface{})
type cb_group_voice_state_ftype func(this *Tox, groupNumber GroupNumber, voiceState GroupVoiceState, userData interface{})
type cb_group_topic_lock_ftype func(this *Tox, groupNumber GroupNumber, topicLock GroupTopicLock, userData interface{})
type cb_group_peer_limit_ftype func(this *Tox, groupNumber GroupNumber, peerLimit uint32, userData interface{})
type cb_group_password_ftype func(this *Tox, groupNumber GroupNumber, password string, userData interface{})
type cb_group_message_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, message string, userData interface{})
type cb_group_private_message_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, message string, userData interface{})
type cb_group_custom_packet_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, data []byte, userData interface{})
type cb_group_custom_private_packet_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, data []byte, userData interface{})
type cb_group_invite_ftype func(this *Tox, groupNumber GroupNumber, friendNumber uint32, data string, userData interface{})
type cb_group_peer_join_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, userData interface{})
type cb_group_peer_exit_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, exitType GroupExitType, name string, userData interface{})
type cb_group_self_join_ftype func(this *Tox, groupNumber GroupNumber, userData interface{})
type cb_group_join_fail_ftype func(this *Tox, groupNumber GroupNumber, failType GroupJoinFail, userData interface{})
type cb_group_moderation_ftype func(this *Tox, groupNumber GroupNumber, peerNumber GroupPeerNumber, modEvent GroupModEvent, userName string, userData interface{})

func (this *Tox) GroupMaxTopicLength() uint32 {
	return uint32(C.tox_group_max_topic_length())
}

func (this *Tox) GroupMaxPartLength() uint32 {
	return uint32(C.tox_group_max_part_length())
}

func (this *Tox) GroupMaxMessageLength() uint32 {
	return uint32(C.tox_group_max_message_length())
}

func (this *Tox) GroupMaxCustomLossyPacketLength() uint32 {
	return uint32(C.tox_group_max_custom_lossy_packet_length())
}

func (this *Tox) GroupMaxCustomLosslessPacketLength() uint32 {
	return uint32(C.tox_group_max_custom_lossless_packet_length())
}

func (this *Tox) GroupMaxGroupNameLength() uint32 {
	return uint32(C.tox_group_max_group_name_length())
}

func (this *Tox) GroupMaxPasswordSize() uint32 {
	return uint32(C.tox_group_max_password_size())
}

func (this *Tox) GroupChatIdSize() uint32 {
	return uint32(C.tox_group_chat_id_size())
}

func (this *Tox) GroupPeerPublicKeySize() uint32 {
	return uint32(C.tox_group_peer_public_key_size())
}

func GroupPrivacyStateToString(value GroupPrivacyState) string {
	return C.GoString(C.tox_group_privacy_state_to_string(C.Tox_Group_Privacy_State(value)))
}

func GroupTopicLockToString(value GroupTopicLock) string {
	return C.GoString(C.tox_group_topic_lock_to_string(C.Tox_Group_Topic_Lock(value)))
}

func GroupVoiceStateToString(value GroupVoiceState) string {
	return C.GoString(C.tox_group_voice_state_to_string(C.Tox_Group_Voice_State(value)))
}

func GroupRoleToString(value GroupRole) string {
	return C.GoString(C.tox_group_role_to_string(C.Tox_Group_Role(value)))
}

func GroupExitTypeToString(value GroupExitType) string {
	return C.GoString(C.tox_group_exit_type_to_string(C.Tox_Group_Exit_Type(value)))
}

func GroupJoinFailToString(value GroupJoinFail) string {
	return C.GoString(C.tox_group_join_fail_to_string(C.Tox_Group_Join_Fail(value)))
}

func GroupModEventToString(value GroupModEvent) string {
	return C.GoString(C.tox_group_mod_event_to_string(C.Tox_Group_Mod_Event(value)))
}

func groupJoinErrorToString(err int) string {
	switch err {
	case 0: // TOX_ERR_GROUP_JOIN_OK
		return "success"
	case 1: // TOX_ERR_GROUP_JOIN_NULL
		return "NULL pointer"
	case 2: // TOX_ERR_GROUP_JOIN_CHAT_ID_INVALID
		return "chat id invalid"
	case 3: // TOX_ERR_GROUP_JOIN_BAD_PASSWORD
		return "bad password"
	case 4: // TOX_ERR_GROUP_JOIN_FAILED_DECRYPT
		return "failed to decrypt"
	case 5: // TOX_ERR_GROUP_JOIN_INVALID_HANDLE
		return "invalid handle"
	default:
		return fmt.Sprintf("unknown error %d", err)
	}
}

func (this *Tox) GroupNew(privacyState GroupPrivacyState, groupName string, name string) (GroupNumber, error) {
	this.lock()
	defer this.unlock()

	var _privacy_state = C.Tox_Group_Privacy_State(privacyState)
	
	var _group_name_cstr *C.char
	var _group_name_len C.size_t
	if len(groupName) > 0 {
		_group_name_cstr = C.CString(groupName)
		defer C.free(unsafe.Pointer(_group_name_cstr))
		_group_name_len = C.size_t(len(groupName))
	}
	
	var _name_cstr *C.char
	var _name_len C.size_t
	if len(name) > 0 {
		_name_cstr = C.CString(name)
		defer C.free(unsafe.Pointer(_name_cstr))
		_name_len = C.size_t(len(name))
	}

	var cerr C.Tox_Err_Group_New
	// C: tox_group_new(tox, privacy_state, group_name[], group_name_length, name[], name_length, error)
	r := C.tox_group_new(this.toxcore, _privacy_state,
		(*C.uint8_t)(unsafe.Pointer(_group_name_cstr)), _group_name_len,
		(*C.uint8_t)(unsafe.Pointer(_name_cstr)), _name_len, &cerr)
	if r == C.UINT32_MAX {
		return GroupNumber(r), toxerrf("group new failed: %d", cerr)
	}
	return GroupNumber(r), nil
}

func (this *Tox) GroupJoin(chatId string, name string, password string) (GroupNumber, error) {
	if chatId == "" || len(chatId) < 20 {
		return 0, errors.New("Invalid chatId:" + chatId)
	}

	data, err := hex.DecodeString(chatId)
	if err != nil {
		return 0, err
	}
	if data == nil || len(data) < 10 {
		return 0, errors.New("Invalid data: " + chatId)
	}

	this.lock()
	defer this.unlock()

	var _name_ptr *C.uint8_t
	var _name_len C.size_t
	if len(name) > 0 {
		_name_cstr := C.CString(name)
		defer C.free(unsafe.Pointer(_name_cstr))
		_name_ptr = (*C.uint8_t)(unsafe.Pointer(_name_cstr))
		_name_len = C.size_t(len(name))
	}

	var _password_ptr *C.uint8_t
	var _password_len C.size_t
	if len(password) > 0 {
		_password_cstr := C.CString(password)
		defer C.free(unsafe.Pointer(_password_cstr))
		_password_ptr = (*C.uint8_t)(unsafe.Pointer(_password_cstr))
		_password_len = C.size_t(len(password))
	}

	var cerr C.Tox_Err_Group_Join
	
	cData := C.CBytes(data)
	defer C.free(cData)
	
	log.Printf("GroupJoin: chatId=%s, data_len=%d, name=%s, password=%s", 
		chatId, len(data), name, password)
	
	r := C.tox_group_join(this.toxcore, (*C.uint8_t)(cData),
		_name_ptr, _name_len,
		_password_ptr, _password_len, &cerr)
	if r == C.UINT32_MAX {
		errStr := groupJoinErrorToString(int(cerr))
		return GroupNumber(r), toxerrf("group join failed: %s (code %d)", errStr, cerr)
	}
	return GroupNumber(r), nil
}

func (this *Tox) GroupIsConnected(groupNumber GroupNumber) (bool, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Is_Connected
	r := C.tox_group_is_connected(this.toxcore, _gn, &cerr)
	if cerr != 0 {
		return false, toxerrf("group is connected failed: %d", cerr)
	}
	return bool(r), nil
}

func (this *Tox) GroupDisconnect(groupNumber GroupNumber) error {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Disconnect
	r := C.tox_group_disconnect(this.toxcore, _gn, &cerr)
	if r == false {
		return toxerrf("group disconnect failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupReconnect(groupNumber GroupNumber) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Reconnect
	r := C.tox_group_reconnect(this.toxcore, _gn, &cerr)
	if r == false {
		return toxerrf("group reconnect failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupLeave(groupNumber GroupNumber, partMessage string) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	
	var _part_message_ptr *C.uint8_t
	var _length C.size_t
	if len(partMessage) > 0 {
		_part_message_bytes := []byte(partMessage)
		_part_message_ptr = (*C.uint8_t)(&_part_message_bytes[0])
		_length = C.size_t(len(partMessage))
	}

	var cerr C.Tox_Err_Group_Leave
	r := C.tox_group_leave(this.toxcore, _gn, _part_message_ptr, _length, &cerr)
	if r == false {
		return toxerrf("group leave failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSelfSetName(groupNumber GroupNumber, name string) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _name = []byte(name)
	var _name_len = C.size_t(len(name))

	var cerr C.Tox_Err_Group_Self_Name_Set
	r := C.tox_group_self_set_name(this.toxcore, _gn, (*C.uint8_t)(&_name[0]), _name_len, &cerr)
	if r == false {
		return toxerrf("group self set name failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSelfGetNameSize(groupNumber GroupNumber) (size_t, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Self_Query
	r := C.tox_group_self_get_name_size(this.toxcore, _gn, &cerr)
	if r == C.SIZE_MAX {
		return 0, toxerrf("group self get name size failed: %d", cerr)
	}
	return size_t(r), nil
}

func (this *Tox) GroupSelfGetName(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	nameSize, err := this.GroupSelfGetNameSize(groupNumber)
	if err != nil {
		return "", err
	}
	if nameSize == 0 {
		return "", nil
	}

	nameBuf := make([]byte, nameSize)
	r := C.tox_group_self_get_name(this.toxcore, _gn, (*C.uint8_t)(&nameBuf[0]), nil)
	if r == false {
		return "", errors.New("group self get name failed")
	}
	return string(nameBuf), nil
}

func (this *Tox) GroupSelfSetStatus(groupNumber GroupNumber, status int) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _status = C.Tox_User_Status(status)

	var cerr C.Tox_Err_Group_Self_Status_Set
	r := C.tox_group_self_set_status(this.toxcore, _gn, _status, &cerr)
	if r == false {
		return toxerrf("group self set status failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSelfGetStatus(groupNumber GroupNumber) (int, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Self_Query
	r := C.tox_group_self_get_status(this.toxcore, _gn, &cerr)
	if int(r) == -1 {
		return int(r), toxerrf("group self get status failed: %d", cerr)
	}
	return int(r), nil
}

func (this *Tox) GroupSelfGetRole(groupNumber GroupNumber) (GroupRole, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Self_Query
	r := C.tox_group_self_get_role(this.toxcore, _gn, &cerr)
	if int(r) == -1 {
		return GroupRole(r), toxerrf("group self get role failed: %d", cerr)
	}
	return GroupRole(r), nil
}

func (this *Tox) GroupSelfGetPeerId(groupNumber GroupNumber) (GroupPeerNumber, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_Self_Query
	r := C.tox_group_self_get_peer_id(this.toxcore, _gn, &cerr)
	if r == C.UINT32_MAX {
		return GroupPeerNumber(r), toxerrf("group self get peer id failed: %d", cerr)
	}
	return GroupPeerNumber(r), nil
}

func (this *Tox) GroupSelfGetPublicKey(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pubkey [PUBLIC_KEY_SIZE]byte

	var cerr C.Tox_Err_Group_Self_Query
	r := C.tox_group_self_get_public_key(this.toxcore, _gn, (*C.uint8_t)(&_pubkey[0]), &cerr)
	if r == false {
		return "", toxerrf("group self get public key failed: %d", cerr)
	}
	pubkey := hex.EncodeToString(_pubkey[:])
	pubkey = strings.ToUpper(pubkey)
	return pubkey, nil
}

func (this *Tox) GroupPeerGetNameSize(groupNumber GroupNumber, peerNumber GroupPeerNumber) (size_t, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	var cerr C.Tox_Err_Group_Peer_Query
	r := C.tox_group_peer_get_name_size(this.toxcore, _gn, _pn, &cerr)
	if r == C.SIZE_MAX {
		return 0, toxerrf("group peer get name size failed: %d", cerr)
	}
	return size_t(r), nil
}

func (this *Tox) GroupPeerGetName(groupNumber GroupNumber, peerNumber GroupPeerNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	nameSize, err := this.GroupPeerGetNameSize(groupNumber, peerNumber)
	if err != nil {
		return "", err
	}
	if nameSize == 0 {
		return "", nil
	}

	nameBuf := make([]byte, nameSize)
	r := C.tox_group_peer_get_name(this.toxcore, _gn, _pn, (*C.uint8_t)(&nameBuf[0]), nil)
	if r == false {
		return "", errors.New("group peer get name failed")
	}
	return string(nameBuf), nil
}

func (this *Tox) GroupPeerGetStatus(groupNumber GroupNumber, peerNumber GroupPeerNumber) (int, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	var cerr C.Tox_Err_Group_Peer_Query
	r := C.tox_group_peer_get_status(this.toxcore, _gn, _pn, &cerr)
	if int(r) == -1 {
		return int(r), toxerrf("group peer get status failed: %d", cerr)
	}
	return int(r), nil
}

func (this *Tox) GroupPeerGetRole(groupNumber GroupNumber, peerNumber GroupPeerNumber) (GroupRole, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	var cerr C.Tox_Err_Group_Peer_Query
	r := C.tox_group_peer_get_role(this.toxcore, _gn, _pn, &cerr)
	if int(r) == -1 {
		return GroupRole(r), toxerrf("group peer get role failed: %d", cerr)
	}
	return GroupRole(r), nil
}

func (this *Tox) GroupPeerGetConnectionStatus(groupNumber GroupNumber, peerNumber GroupPeerNumber) (int, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	var cerr C.Tox_Err_Group_Peer_Query
	r := C.tox_group_peer_get_connection_status(this.toxcore, _gn, _pn, &cerr)
	if int(r) == -1 {
		return int(r), toxerrf("group peer get connection status failed: %d", cerr)
	}
	return int(r), nil
}

func (this *Tox) GroupPeerGetPublicKey(groupNumber GroupNumber, peerNumber GroupPeerNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)
	var _pubkey [PUBLIC_KEY_SIZE]byte

	var cerr C.Tox_Err_Group_Peer_Query
	r := C.tox_group_peer_get_public_key(this.toxcore, _gn, _pn, (*C.uint8_t)(&_pubkey[0]), &cerr)
	if r == false {
		return "", toxerrf("group peer get public key failed: %d", cerr)
	}
	pubkey := hex.EncodeToString(_pubkey[:])
	pubkey = strings.ToUpper(pubkey)
	return pubkey, nil
}

func (this *Tox) GroupGetNameSize(groupNumber GroupNumber) (size_t, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_name_size(this.toxcore, _gn, &cerr)
	if r == C.SIZE_MAX {
		return 0, toxerrf("group get name size failed: %d", cerr)
	}
	return size_t(r), nil
}

func (this *Tox) GroupGetName(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	nameSize, err := this.GroupGetNameSize(groupNumber)
	if err != nil {
		return "", err
	}
	if nameSize == 0 {
		return "", nil
	}

	nameBuf := make([]byte, nameSize)
	r := C.tox_group_get_name(this.toxcore, _gn, (*C.uint8_t)(&nameBuf[0]), nil)
	if r == false {
		return "", errors.New("group get name failed")
	}
	return string(nameBuf), nil
}

func (this *Tox) GroupGetTopicSize(groupNumber GroupNumber) (size_t, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_topic_size(this.toxcore, _gn, &cerr)
	if r == C.SIZE_MAX {
		return 0, toxerrf("group get topic size failed: %d", cerr)
	}
	return size_t(r), nil
}

func (this *Tox) GroupGetTopic(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	topicSize, err := this.GroupGetTopicSize(groupNumber)
	if err != nil {
		return "", err
	}
	if topicSize == 0 {
		return "", nil
	}

	topicBuf := make([]byte, topicSize)
	r := C.tox_group_get_topic(this.toxcore, _gn, (*C.uint8_t)(&topicBuf[0]), nil)
	if r == false {
		return "", errors.New("group get topic failed")
	}
	return string(topicBuf), nil
}

func (this *Tox) GroupSetTopic(groupNumber GroupNumber, topic string) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _topic = []byte(topic)
	var _topic_len = C.size_t(len(topic))

	var cerr C.Tox_Err_Group_Topic_Set
	r := C.tox_group_set_topic(this.toxcore, _gn, (*C.uint8_t)(&_topic[0]), _topic_len, &cerr)
	if r == false {
		return toxerrf("group set topic failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupGetChatId(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)
	idSize := C.tox_group_chat_id_size()
	idBuf := make([]byte, idSize)

	r := C.tox_group_get_chat_id(this.toxcore, _gn, (*C.uint8_t)(&idBuf[0]), nil)
	if r == false {
		return "", errors.New("group get chat id failed")
	}
	chatId := hex.EncodeToString(idBuf)
	chatId = strings.ToUpper(chatId)
	return chatId, nil
}

func (this *Tox) GroupByChatId(chatId string) (GroupNumber, error) {
	if chatId == "" || len(chatId) < 20 {
		return 0, errors.New("Invalid chatId:" + chatId)
	}

	data, err := hex.DecodeString(chatId)
	if err != nil {
		return 0, err
	}

	var _chat_id = (*C.uint8_t)(&data[0])

	r := C.tox_group_by_id(this.toxcore, _chat_id, nil)
	if r == C.UINT32_MAX {
		return 0, errors.New("group by chat id not found")
	}
	return GroupNumber(r), nil
}

func (this *Tox) GroupGetNumberGroups() uint32 {
	r := C.tox_group_get_number_groups(this.toxcore)
	return uint32(r)
}

func (this *Tox) GroupGetPrivacyState(groupNumber GroupNumber) (GroupPrivacyState, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_privacy_state(this.toxcore, _gn, &cerr)
	if int(r) == -1 {
		return GroupPrivacyState(r), toxerrf("group get privacy state failed: %d", cerr)
	}
	return GroupPrivacyState(r), nil
}

func (this *Tox) GroupSetPrivacyState(groupNumber GroupNumber, privacyState GroupPrivacyState) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _privacy_state = C.Tox_Group_Privacy_State(privacyState)

	var cerr C.Tox_Err_Group_Set_Privacy_State
	r := C.tox_group_set_privacy_state(this.toxcore, _gn, _privacy_state, &cerr)
	if r == false {
		return toxerrf("group set privacy state failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupGetVoiceState(groupNumber GroupNumber) (GroupVoiceState, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_voice_state(this.toxcore, _gn, &cerr)
	if int(r) == -1 {
		return GroupVoiceState(r), toxerrf("group get voice state failed: %d", cerr)
	}
	return GroupVoiceState(r), nil
}

func (this *Tox) GroupSetVoiceState(groupNumber GroupNumber, voiceState GroupVoiceState) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _voice_state = C.Tox_Group_Voice_State(voiceState)

	var cerr C.Tox_Err_Group_Set_Voice_State
	r := C.tox_group_set_voice_state(this.toxcore, _gn, _voice_state, &cerr)
	if r == false {
		return toxerrf("group set voice state failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupGetTopicLock(groupNumber GroupNumber) (GroupTopicLock, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_topic_lock(this.toxcore, _gn, &cerr)
	if int(r) == -1 {
		return GroupTopicLock(r), toxerrf("group get topic lock failed: %d", cerr)
	}
	return GroupTopicLock(r), nil
}

func (this *Tox) GroupSetTopicLock(groupNumber GroupNumber, topicLock GroupTopicLock) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _topic_lock = C.Tox_Group_Topic_Lock(topicLock)

	var cerr C.Tox_Err_Group_Set_Topic_Lock
	r := C.tox_group_set_topic_lock(this.toxcore, _gn, _topic_lock, &cerr)
	if r == false {
		return toxerrf("group set topic lock failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupGetPeerLimit(groupNumber GroupNumber) (uint16, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_peer_limit(this.toxcore, _gn, &cerr)
	if r == 0 && cerr != C.TOX_ERR_GROUP_STATE_QUERY_OK {
		return uint16(r), toxerrf("group get peer limit failed: %d", cerr)
	}
	return uint16(r), nil
}

func (this *Tox) GroupSetPeerLimit(groupNumber GroupNumber, peerLimit uint16) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _peer_limit = C.uint16_t(peerLimit)

	var cerr C.Tox_Err_Group_Set_Peer_Limit
	r := C.tox_group_set_peer_limit(this.toxcore, _gn, _peer_limit, &cerr)
	if r == false {
		return toxerrf("group set peer limit failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupGetPasswordSize(groupNumber GroupNumber) (size_t, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	var cerr C.Tox_Err_Group_State_Query
	r := C.tox_group_get_password_size(this.toxcore, _gn, &cerr)
	if r == C.SIZE_MAX {
		return 0, toxerrf("group get password size failed: %d", cerr)
	}
	return size_t(r), nil
}

func (this *Tox) GroupGetPassword(groupNumber GroupNumber) (string, error) {
	var _gn = C.Tox_Group_Number(groupNumber)

	passwordSize, err := this.GroupGetPasswordSize(groupNumber)
	if err != nil {
		return "", err
	}
	if passwordSize == 0 {
		return "", nil
	}

	passwordBuf := make([]byte, passwordSize)
	r := C.tox_group_get_password(this.toxcore, _gn, (*C.uint8_t)(&passwordBuf[0]), nil)
	if r == false {
		return "", errors.New("group get password failed")
	}
	return string(passwordBuf), nil
}

func (this *Tox) GroupSetPassword(groupNumber GroupNumber, password string) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _password = []byte(password)
	var _password_len = C.size_t(len(password))

	var cerr C.Tox_Err_Group_Set_Password
	r := C.tox_group_set_password(this.toxcore, _gn, (*C.uint8_t)(&_password[0]), _password_len, &cerr)
	if r == false {
		return toxerrf("group set password failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSendMessage(groupNumber GroupNumber, messageType int, message string) (GroupMessageId, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _message = []byte(message)
	var _length = C.size_t(len(message))

	switch messageType {
	case MESSAGE_TYPE_NORMAL:
	case MESSAGE_TYPE_ACTION:
	default:
		return 0, toxerrf("Invalid message type: %d", messageType)
	}

	var cerr C.Tox_Err_Group_Send_Message
	r := C.tox_group_send_message(this.toxcore, _gn, (C.Tox_Message_Type)(messageType),
		(*C.uint8_t)(&_message[0]), _length, &cerr)
	if cerr != C.TOX_ERR_GROUP_SEND_MESSAGE_OK {
		return GroupMessageId(r), toxerrf("group send message failed: %d", cerr)
	}
	return GroupMessageId(r), nil
}

func (this *Tox) GroupSendPrivateMessage(groupNumber GroupNumber, peerNumber GroupPeerNumber, messageType int, message string) (GroupMessageId, error) {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)
	var _message = []byte(message)
	var _length = C.size_t(len(message))

	switch messageType {
	case MESSAGE_TYPE_NORMAL:
	case MESSAGE_TYPE_ACTION:
	default:
		return 0, toxerrf("Invalid message type: %d", messageType)
	}

	var cerr C.Tox_Err_Group_Send_Private_Message
	r := C.tox_group_send_private_message(this.toxcore, _gn, _pn, (C.Tox_Message_Type)(messageType),
		(*C.uint8_t)(&_message[0]), _length, &cerr)
	if cerr != C.TOX_ERR_GROUP_SEND_PRIVATE_MESSAGE_OK {
		return GroupMessageId(r), toxerrf("group send private message failed: %d", cerr)
	}
	return GroupMessageId(r), nil
}

func (this *Tox) GroupSendCustomPacket(groupNumber GroupNumber, data []byte, lossless bool) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _data = (*C.uint8_t)(&data[0])
	var _length = C.size_t(len(data))
	var _lossless = C.bool(lossless)

	var cerr C.Tox_Err_Group_Send_Custom_Packet
	r := C.tox_group_send_custom_packet(this.toxcore, _gn, _lossless, _data, _length, &cerr)
	if r == false {
		return toxerrf("group send custom packet failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSendCustomPrivatePacket(groupNumber GroupNumber, peerNumber GroupPeerNumber, data []byte, lossless bool) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)
	var _data = (*C.uint8_t)(&data[0])
	var _length = C.size_t(len(data))
	var _lossless = C.bool(lossless)

	var cerr C.Tox_Err_Group_Send_Custom_Private_Packet
	r := C.tox_group_send_custom_private_packet(this.toxcore, _gn, _pn, _lossless, _data, _length, &cerr)
	if r == false {
		return toxerrf("group send custom private packet failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupInviteFriend(groupNumber GroupNumber, friendNumber uint32) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _fn = C.uint32_t(friendNumber)

	var cerr C.Tox_Err_Group_Invite_Friend
	r := C.tox_group_invite_friend(this.toxcore, _gn, _fn, &cerr)
	if r == false {
		return toxerrf("group invite friend failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupInviteAccept(inviteData string, friendNumber uint32, name string, password string) (GroupNumber, error) {
	if inviteData == "" || len(inviteData) < 20 {
		return 0, errors.New("Invalid inviteData:" + inviteData)
	}

	data, err := hex.DecodeString(inviteData)
	if err != nil {
		return 0, err
	}

	this.lock()
	defer this.unlock()

	var _fn = C.uint32_t(friendNumber)
	
	// invite_data 必须非 NULL
	_data := (*C.uint8_t)(&data[0])
	_length := C.size_t(len(data))
	
	// name 必须非 NULL 且长度 > 0
	nameBytes := []byte(name)
	_name := (*C.uint8_t)(&nameBytes[0])
	_name_len := C.size_t(len(nameBytes))
	
	var _password *C.uint8_t
	var _password_len C.size_t
	if len(password) > 0 {
		pwdBytes := []byte(password)
		_password = (*C.uint8_t)(&pwdBytes[0])
		_password_len = C.size_t(len(pwdBytes))
	}

	var cerr C.Tox_Err_Group_Invite_Accept
	r := C.tox_group_invite_accept(this.toxcore, _fn, _data, _length,
		_name, _name_len, _password, _password_len, &cerr)
	if r == C.UINT32_MAX {
		return GroupNumber(r), toxerrf("group invite accept failed: %d", cerr)
	}
	return GroupNumber(r), nil
}

func (this *Tox) GroupSetRole(groupNumber GroupNumber, peerNumber GroupPeerNumber, role GroupRole) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)
	var _role = C.Tox_Group_Role(role)

	var cerr C.Tox_Err_Group_Set_Role
	r := C.tox_group_set_role(this.toxcore, _gn, _pn, _role, &cerr)
	if r == false {
		return toxerrf("group set role failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupKickPeer(groupNumber GroupNumber, peerNumber GroupPeerNumber) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)

	var cerr C.Tox_Err_Group_Kick_Peer
	r := C.tox_group_kick_peer(this.toxcore, _gn, _pn, &cerr)
	if r == false {
		return toxerrf("group kick peer failed: %d", cerr)
	}
	return nil
}

func (this *Tox) GroupSetIgnore(groupNumber GroupNumber, peerNumber GroupPeerNumber, ignore bool) error {
	this.lock()
	defer this.unlock()

	var _gn = C.Tox_Group_Number(groupNumber)
	var _pn = C.Tox_Group_Peer_Number(peerNumber)
	var _ignore = C.bool(ignore)

	var cerr C.Tox_Err_Group_Set_Ignore
	r := C.tox_group_set_ignore(this.toxcore, _gn, _pn, _ignore, &cerr)
	if r == false {
		return toxerrf("group set ignore failed: %d", cerr)
	}
	return nil
}

//export callbackGroupPeerNameWrapperForC
func callbackGroupPeerNameWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_peer_names {
		cbfn := *(*cb_group_peer_name_ftype)(cbfni)
		name := C.GoStringN((*C.char)(unsafe.Pointer(a2)), C.int(a3))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), name, ud) })
	}
}

func (this *Tox) CallbackGroupPeerName(cbfn cb_group_peer_name_ftype, userData interface{}) {
	this.CallbackGroupPeerNameAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPeerNameAdd(cbfn cb_group_peer_name_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_peer_names[cbfnp]; ok {
		return
	}
	this.cb_group_peer_names[cbfnp] = userData

	C.tox_callback_group_peer_name(this.toxcore, (*C.tox_group_peer_name_cb)(C.callbackGroupPeerNameWrapperForC))
}

//export callbackGroupPeerStatusWrapperForC
func callbackGroupPeerStatusWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 C.Tox_User_Status, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_peer_statuses {
		cbfn := *(*cb_group_peer_status_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), int(a2), ud) })
	}
}

func (this *Tox) CallbackGroupPeerStatus(cbfn cb_group_peer_status_ftype, userData interface{}) {
	this.CallbackGroupPeerStatusAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPeerStatusAdd(cbfn cb_group_peer_status_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_peer_statuses[cbfnp]; ok {
		return
	}
	this.cb_group_peer_statuses[cbfnp] = userData

	C.tox_callback_group_peer_status(this.toxcore, (*C.tox_group_peer_status_cb)(C.callbackGroupPeerStatusWrapperForC))
}

//export callbackGroupTopicWrapperForC
func callbackGroupTopicWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_topics {
		cbfn := *(*cb_group_topic_ftype)(cbfni)
		topic := C.GoStringN((*C.char)(unsafe.Pointer(a2)), C.int(a3))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), topic, ud) })
	}
}

func (this *Tox) CallbackGroupTopic(cbfn cb_group_topic_ftype, userData interface{}) {
	this.CallbackGroupTopicAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupTopicAdd(cbfn cb_group_topic_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_topics[cbfnp]; ok {
		return
	}
	this.cb_group_topics[cbfnp] = userData

	C.tox_callback_group_topic(this.toxcore, (*C.tox_group_topic_cb)(C.callbackGroupTopicWrapperForC))
}

//export callbackGroupPrivacyStateWrapperForC
func callbackGroupPrivacyStateWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Privacy_State, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_privacy_states {
		cbfn := *(*cb_group_privacy_state_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPrivacyState(a1), ud) })
	}
}

func (this *Tox) CallbackGroupPrivacyState(cbfn cb_group_privacy_state_ftype, userData interface{}) {
	this.CallbackGroupPrivacyStateAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPrivacyStateAdd(cbfn cb_group_privacy_state_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_privacy_states[cbfnp]; ok {
		return
	}
	this.cb_group_privacy_states[cbfnp] = userData

	C.tox_callback_group_privacy_state(this.toxcore, (*C.tox_group_privacy_state_cb)(C.callbackGroupPrivacyStateWrapperForC))
}

//export callbackGroupVoiceStateWrapperForC
func callbackGroupVoiceStateWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Voice_State, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_voice_states {
		cbfn := *(*cb_group_voice_state_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupVoiceState(a1), ud) })
	}
}

func (this *Tox) CallbackGroupVoiceState(cbfn cb_group_voice_state_ftype, userData interface{}) {
	this.CallbackGroupVoiceStateAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupVoiceStateAdd(cbfn cb_group_voice_state_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_voice_states[cbfnp]; ok {
		return
	}
	this.cb_group_voice_states[cbfnp] = userData

	C.tox_callback_group_voice_state(this.toxcore, (*C.tox_group_voice_state_cb)(C.callbackGroupVoiceStateWrapperForC))
}

//export callbackGroupTopicLockWrapperForC
func callbackGroupTopicLockWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Topic_Lock, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_topic_locks {
		cbfn := *(*cb_group_topic_lock_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupTopicLock(a1), ud) })
	}
}

func (this *Tox) CallbackGroupTopicLock(cbfn cb_group_topic_lock_ftype, userData interface{}) {
	this.CallbackGroupTopicLockAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupTopicLockAdd(cbfn cb_group_topic_lock_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_topic_locks[cbfnp]; ok {
		return
	}
	this.cb_group_topic_locks[cbfnp] = userData

	C.tox_callback_group_topic_lock(this.toxcore, (*C.tox_group_topic_lock_cb)(C.callbackGroupTopicLockWrapperForC))
}

//export callbackGroupPeerLimitWrapperForC
func callbackGroupPeerLimitWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.uint32_t, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_peer_limits {
		cbfn := *(*cb_group_peer_limit_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), uint32(a1), ud) })
	}
}

func (this *Tox) CallbackGroupPeerLimit(cbfn cb_group_peer_limit_ftype, userData interface{}) {
	this.CallbackGroupPeerLimitAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPeerLimitAdd(cbfn cb_group_peer_limit_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_peer_limits[cbfnp]; ok {
		return
	}
	this.cb_group_peer_limits[cbfnp] = userData

	C.tox_callback_group_peer_limit(this.toxcore, (*C.tox_group_peer_limit_cb)(C.callbackGroupPeerLimitWrapperForC))
}

//export callbackGroupPasswordWrapperForC
func callbackGroupPasswordWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 *C.uint8_t, a2 C.size_t, a3 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_passwords {
		cbfn := *(*cb_group_password_ftype)(cbfni)
		password := C.GoStringN((*C.char)(unsafe.Pointer(a1)), C.int(a2))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), password, ud) })
	}
}

func (this *Tox) CallbackGroupPassword(cbfn cb_group_password_ftype, userData interface{}) {
	this.CallbackGroupPasswordAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPasswordAdd(cbfn cb_group_password_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_passwords[cbfnp]; ok {
		return
	}
	this.cb_group_passwords[cbfnp] = userData

	C.tox_callback_group_password(this.toxcore, (*C.tox_group_password_cb)(C.callbackGroupPasswordWrapperForC))
}

//export callbackGroupMessageWrapperForC
func callbackGroupMessageWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 C.Tox_Message_Type, a3 *C.uint8_t, a4 C.size_t, a5 C.Tox_Group_Message_Id, a6 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	if int(a2) == MESSAGE_TYPE_NORMAL {
		for cbfni, ud := range this.cb_group_messages {
			cbfn := *(*cb_group_message_ftype)(cbfni)
			message := C.GoStringN((*C.char)(unsafe.Pointer(a3)), C.int(a4))
			this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), message, ud) })
		}
	} else {
		// Handle action messages if needed
	}
}

func (this *Tox) CallbackGroupMessage(cbfn cb_group_message_ftype, userData interface{}) {
	this.CallbackGroupMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupMessageAdd(cbfn cb_group_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_messages[cbfnp]; ok {
		return
	}
	this.cb_group_messages[cbfnp] = userData

	C.tox_callback_group_message(this.toxcore, (*C.tox_group_message_cb)(C.callbackGroupMessageWrapperForC))
}

//export callbackGroupPrivateMessageWrapperForC
func callbackGroupPrivateMessageWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 C.Tox_Message_Type, a3 *C.uint8_t, a4 C.size_t, a5 C.Tox_Group_Message_Id, a6 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_private_messages {
		cbfn := *(*cb_group_private_message_ftype)(cbfni)
		message := C.GoStringN((*C.char)(unsafe.Pointer(a3)), C.int(a4))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), message, ud) })
	}
}

func (this *Tox) CallbackGroupPrivateMessage(cbfn cb_group_private_message_ftype, userData interface{}) {
	this.CallbackGroupPrivateMessageAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPrivateMessageAdd(cbfn cb_group_private_message_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_private_messages[cbfnp]; ok {
		return
	}
	this.cb_group_private_messages[cbfnp] = userData

	C.tox_callback_group_private_message(this.toxcore, (*C.tox_group_private_message_cb)(C.callbackGroupPrivateMessageWrapperForC))
}

//export callbackGroupCustomPacketWrapperForC
func callbackGroupCustomPacketWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_custom_packets {
		cbfn := *(*cb_group_custom_packet_ftype)(cbfni)
		data := C.GoBytes(unsafe.Pointer(a2), C.int(a3))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), data, ud) })
	}
}

func (this *Tox) CallbackGroupCustomPacket(cbfn cb_group_custom_packet_ftype, userData interface{}) {
	this.CallbackGroupCustomPacketAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupCustomPacketAdd(cbfn cb_group_custom_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_custom_packets[cbfnp]; ok {
		return
	}
	this.cb_group_custom_packets[cbfnp] = userData

	C.tox_callback_group_custom_packet(this.toxcore, (*C.tox_group_custom_packet_cb)(C.callbackGroupCustomPacketWrapperForC))
}

//export callbackGroupCustomPrivatePacketWrapperForC
func callbackGroupCustomPrivatePacketWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 *C.uint8_t, a3 C.size_t, a4 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_custom_private_packets {
		cbfn := *(*cb_group_custom_private_packet_ftype)(cbfni)
		data := C.GoBytes(unsafe.Pointer(a2), C.int(a3))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), data, ud) })
	}
}

func (this *Tox) CallbackGroupCustomPrivatePacket(cbfn cb_group_custom_private_packet_ftype, userData interface{}) {
	this.CallbackGroupCustomPrivatePacketAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupCustomPrivatePacketAdd(cbfn cb_group_custom_private_packet_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_custom_private_packets[cbfnp]; ok {
		return
	}
	this.cb_group_custom_private_packets[cbfnp] = userData

	C.tox_callback_group_custom_private_packet(this.toxcore, (*C.tox_group_custom_private_packet_cb)(C.callbackGroupCustomPrivatePacketWrapperForC))
}

//export callbackGroupInviteWrapperForC
func callbackGroupInviteWrapperForC(m *C.Tox, friendNumber C.uint32_t, cookie *C.uint8_t, cookieLength C.size_t, groupName *C.uint8_t, groupNameLength C.size_t, userData unsafe.Pointer) {
    var this = cbUserDatas.get(m)
    if this == nil {
        return
    }
    
    // Process invite data (cookie)
    var cookieData []byte
    if cookie != nil && cookieLength > 0 && cookieLength < 10*1024*1024 {
        cookieData = C.GoBytes(unsafe.Pointer(cookie), C.int(cookieLength))
    }
    cookieStr := hex.EncodeToString(cookieData)
    cookieStr = strings.ToUpper(cookieStr)
    
    for cbfni, ud := range this.cb_group_invites {
        cbfn := *(*cb_group_invite_ftype)(cbfni)
        this.putcbevts(func() { cbfn(this, GroupNumber(0), uint32(friendNumber), cookieStr, ud) })
    }
}

func (this *Tox) CallbackGroupInvite(cbfn cb_group_invite_ftype, userData interface{}) {
    this.CallbackGroupInviteAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupInviteAdd(cbfn cb_group_invite_ftype, userData interface{}) {
    cbfnp := (unsafe.Pointer)(&cbfn)
    if _, ok := this.cb_group_invites[cbfnp]; ok {
        return
    }
    this.cb_group_invites[cbfnp] = userData
    
    C.tox_callback_group_invite(this.toxcore, (*C.tox_group_invite_cb)(C.callbackGroupInviteWrapperForC))
}

//export callbackGroupPeerJoinWrapperForC
func callbackGroupPeerJoinWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_peer_joins {
		cbfn := *(*cb_group_peer_join_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), ud) })
	}
}

func (this *Tox) CallbackGroupPeerJoin(cbfn cb_group_peer_join_ftype, userData interface{}) {
	this.CallbackGroupPeerJoinAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPeerJoinAdd(cbfn cb_group_peer_join_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_peer_joins[cbfnp]; ok {
		return
	}
	this.cb_group_peer_joins[cbfnp] = userData

	C.tox_callback_group_peer_join(this.toxcore, (*C.tox_group_peer_join_cb)(C.callbackGroupPeerJoinWrapperForC))
}

//export callbackGroupPeerExitWrapperForC
func callbackGroupPeerExitWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 C.Tox_Group_Exit_Type, a3 *C.uint8_t, a4 C.size_t, a5 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_peer_exits {
		cbfn := *(*cb_group_peer_exit_ftype)(cbfni)
		name := C.GoStringN((*C.char)(unsafe.Pointer(a3)), C.int(a4))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), GroupExitType(a2), name, ud) })
	}
}

func (this *Tox) CallbackGroupPeerExit(cbfn cb_group_peer_exit_ftype, userData interface{}) {
	this.CallbackGroupPeerExitAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupPeerExitAdd(cbfn cb_group_peer_exit_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_peer_exits[cbfnp]; ok {
		return
	}
	this.cb_group_peer_exits[cbfnp] = userData

	C.tox_callback_group_peer_exit(this.toxcore, (*C.tox_group_peer_exit_cb)(C.callbackGroupPeerExitWrapperForC))
}

//export callbackGroupSelfJoinWrapperForC
func callbackGroupSelfJoinWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_self_joins {
		cbfn := *(*cb_group_self_join_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), ud) })
	}
}

func (this *Tox) CallbackGroupSelfJoin(cbfn cb_group_self_join_ftype, userData interface{}) {
	this.CallbackGroupSelfJoinAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupSelfJoinAdd(cbfn cb_group_self_join_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_self_joins[cbfnp]; ok {
		return
	}
	this.cb_group_self_joins[cbfnp] = userData

	C.tox_callback_group_self_join(this.toxcore, (*C.tox_group_self_join_cb)(C.callbackGroupSelfJoinWrapperForC))
}

//export callbackGroupJoinFailWrapperForC
func callbackGroupJoinFailWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Join_Fail, a2 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_join_fails {
		cbfn := *(*cb_group_join_fail_ftype)(cbfni)
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupJoinFail(a1), ud) })
	}
}

func (this *Tox) CallbackGroupJoinFail(cbfn cb_group_join_fail_ftype, userData interface{}) {
	this.CallbackGroupJoinFailAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupJoinFailAdd(cbfn cb_group_join_fail_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_join_fails[cbfnp]; ok {
		return
	}
	this.cb_group_join_fails[cbfnp] = userData

	C.tox_callback_group_join_fail(this.toxcore, (*C.tox_group_join_fail_cb)(C.callbackGroupJoinFailWrapperForC))
}

//export callbackGroupModerationWrapperForC
func callbackGroupModerationWrapperForC(m *C.Tox, a0 C.Tox_Group_Number, a1 C.Tox_Group_Peer_Number, a2 C.Tox_Group_Mod_Event, a3 *C.uint8_t, a4 C.size_t, a5 unsafe.Pointer) {
	var this = cbUserDatas.get(m)
	for cbfni, ud := range this.cb_group_moderations {
		cbfn := *(*cb_group_moderation_ftype)(cbfni)
		userName := C.GoStringN((*C.char)(unsafe.Pointer(a3)), C.int(a4))
		this.putcbevts(func() { cbfn(this, GroupNumber(a0), GroupPeerNumber(a1), GroupModEvent(a2), userName, ud) })
	}
}

func (this *Tox) CallbackGroupModeration(cbfn cb_group_moderation_ftype, userData interface{}) {
    this.CallbackGroupModerationAdd(cbfn, userData)
}
func (this *Tox) CallbackGroupModerationAdd(cbfn cb_group_moderation_ftype, userData interface{}) {
	cbfnp := (unsafe.Pointer)(&cbfn)
	if _, ok := this.cb_group_moderations[cbfnp]; ok {
		return
	}
	this.cb_group_moderations[cbfnp] = userData

	C.tox_callback_group_moderation(this.toxcore, (*C.tox_group_moderation_cb)(C.callbackGroupModerationWrapperForC))
}
