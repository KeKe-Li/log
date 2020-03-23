package log

import "time"

const TimeFormatLayout = "2006-01-02 15:04:05.000"

// 2006-01-02 15:04:05.000
func FormatTimeString(t time.Time) string {
	result := FormatTime(t)
	return string(result[:])
}

// 2006-01-02 15:04:05.000
func FormatTime(t time.Time) (result [23]byte) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()

	itoa(result[:4], year)
	result[4] = '-'
	itoa(result[5:7], int(month))
	result[7] = '-'
	itoa(result[8:10], day)
	result[10] = ' '
	itoa(result[11:13], hour)
	result[13] = ':'
	itoa(result[14:16], min)
	result[16] = ':'
	itoa(result[17:19], sec)
	result[19] = '.'
	itoa(result[20:], nsec/1e6)
	return
}

// itoa format integer to fixed-width decimal ASCII.
//
// requirement:
//  n must be greater than or equal to 0
//  buf must be large enough to accommodate n
func itoa(buf []byte, n int) {
	i := len(buf) - 1
	for n >= 10 {
		q := n / 10
		buf[i] = byte('0' + n - q*10)
		i--
		n = q
	}
	// n < 10
	buf[i] = byte('0' + n)
	i--
	for i >= 0 {
		buf[i] = '0'
		i--
	}
}
