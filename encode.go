package mtproto

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
	"time"
)

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	return ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
}

type EncodeBuf struct {
	buf []byte
}

func NewEncodeBuf(cap int) *EncodeBuf {
	d := &EncodeBuf{make([]byte, 0, cap)}
	return d
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buf = append(e.buf, res...)
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *EncodeBuf) VectorInt(v []int32) {
	x := make([]byte, 4+4+len(v)*4)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorLong(v []int64) {
	x := make([]byte, 4+4+len(v)*8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorString(v []string) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.String(v)
	}
}

func (e *EncodeBuf) Vector(v []TL) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.buf = append(e.buf, v.Encode()...)
	}
}

// TODO: Does only server send messages below?
func (e TL_msg_container) Encode() []byte { return nil }
func (e TL_resPQ) Encode() []byte {
	x := NewEncodeBuf(90)
	x.UInt(crc_resPQ)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.BigInt(e.Pq)
	x.VectorLong(e.Fingerprints)
	return x.buf
}

func (e TL_server_DH_params_ok) Encode() []byte {
	x := NewEncodeBuf(1000)
	x.UInt(crc_server_DH_params_ok)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.StringBytes(e.Encrypted_answer)
	return x.buf
}

func (e TL_server_DH_params_fail) Encode() []byte { return nil }

func (e TL_server_DH_inner_data) Encode() []byte {
	x := NewEncodeBuf(600)
	x.UInt(crc_server_DH_inner_data)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.Int(e.G)
	x.BigInt(e.Dh_prime)
	x.BigInt(e.G_a)
	x.Int(e.Server_time)
	return x.buf
}

func (e TL_dh_gen_ok) Encode() []byte {
	x := NewEncodeBuf(56)
	x.UInt(crc_dh_gen_ok)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.Bytes(e.New_nonce_hash1)
	return x.buf
}
func (e TL_rpc_result) Encode() []byte           { return nil }
func (e TL_rpc_error) Encode() []byte            { return nil }
func (e TL_new_session_created) Encode() []byte  { return nil }
func (e TL_bad_server_salt) Encode() []byte      { return nil }
func (e TL_bad_msg_notification) Encode() []byte { return nil }

func (e TL_req_pq) Encode() []byte {
	x := NewEncodeBuf(20)
	x.UInt(crc_req_pq)
	x.Bytes(e.Nonce)
	return x.buf
}

func (e TL_p_q_inner_data) Encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(crc_p_q_inner_data)
	x.BigInt(e.Pq)
	x.BigInt(e.P)
	x.BigInt(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.Bytes(e.New_nonce)
	return x.buf
}

func (e TL_req_DH_params) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_req_DH_params)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.BigInt(e.P)
	x.BigInt(e.Q)
	x.Long(int64(e.Fp))
	x.StringBytes(e.Encdata)
	return x.buf
}

func (e TL_client_DH_inner_data) Encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_client_DH_inner_data)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.Long(e.Retry)
	x.BigInt(e.G_b)
	return x.buf
}

func (e TL_set_client_DH_params) Encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(crc_set_client_DH_params)
	x.Bytes(e.Nonce)
	x.Bytes(e.Server_nonce)
	x.StringBytes(e.Encdata)
	return x.buf
}

func (e TL_ping) Encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(crc_ping)
	x.Long(e.Ping_id)
	return x.buf
}

func (e TL_pong) Encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(crc_pong)
	x.Long(e.Msg_id)
	x.Long(e.Ping_id)
	return x.buf
}

func (e TL_msgs_ack) Encode() []byte {
	x := NewEncodeBuf(64)
	x.UInt(crc_msgs_ack)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e TL_boolFalse) Encode() []byte {
	x := NewEncodeBuf(4)
	x.UInt(crc_boolFalse)
	return x.buf
}

func (e TL_boolTrue) Encode() []byte {
	x := NewEncodeBuf(4)
	x.UInt(crc_boolTrue)
	return x.buf
}

func (e TL_null) Encode() []byte {
	x := NewEncodeBuf(4)
	x.UInt(crc_null)
	return x.buf
}
