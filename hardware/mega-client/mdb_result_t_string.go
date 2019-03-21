// Code generated by "stringer -type=Mdb_result_t -trimprefix=MDB_RESULT_"; DO NOT EDIT.

package mega

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MDB_RESULT_SUCCESS-1]
	_ = x[MDB_RESULT_BUSY-8]
	_ = x[MDB_RESULT_INVALID_CHK-9]
	_ = x[MDB_RESULT_NAK-10]
	_ = x[MDB_RESULT_TIMEOUT-11]
	_ = x[MDB_RESULT_INVALID_END-12]
	_ = x[MDB_RESULT_RECEIVE_OVERFLOW-13]
	_ = x[MDB_RESULT_SEND_OVERFLOW-14]
	_ = x[MDB_RESULT_CODE_ERROR-15]
	_ = x[MDB_RESULT_UART_READ_UNEXPECTED-16]
	_ = x[MDB_RESULT_UART_READ_ERROR-17]
	_ = x[MDB_RESULT_UART_READ_OVERFLOW-18]
	_ = x[MDB_RESULT_UART_READ_PARITY-19]
	_ = x[MDB_RESULT_UART_SEND_BUSY-20]
	_ = x[MDB_RESULT_UART_TXC_UNEXPECTED-21]
	_ = x[MDB_RESULT_TIMER_CODE_ERROR-24]
}

const (
	_Mdb_result_t_name_0 = "SUCCESS"
	_Mdb_result_t_name_1 = "BUSYINVALID_CHKNAKTIMEOUTINVALID_ENDRECEIVE_OVERFLOWSEND_OVERFLOWCODE_ERRORUART_READ_UNEXPECTEDUART_READ_ERRORUART_READ_OVERFLOWUART_READ_PARITYUART_SEND_BUSYUART_TXC_UNEXPECTED"
	_Mdb_result_t_name_2 = "TIMER_CODE_ERROR"
)

var (
	_Mdb_result_t_index_1 = [...]uint8{0, 4, 15, 18, 25, 36, 52, 65, 75, 95, 110, 128, 144, 158, 177}
)

func (i Mdb_result_t) String() string {
	switch {
	case i == 1:
		return _Mdb_result_t_name_0
	case 8 <= i && i <= 21:
		i -= 8
		return _Mdb_result_t_name_1[_Mdb_result_t_index_1[i]:_Mdb_result_t_index_1[i+1]]
	case i == 24:
		return _Mdb_result_t_name_2
	default:
		return "Mdb_result_t(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
