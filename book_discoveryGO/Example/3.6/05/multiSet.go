//go:generate stringer -type=StringSet
package multiSet

type StringSet map[string]int

// create new MultiSet
func NewMultiSet() StringSet {
	return make(map[string]int, 0)
}

// Insert 함수는 집합에 val을 추가한다.
func (m StringSet) Insert(val string) {
	m[val]++
	return
}

// Erase 함수는 집합에서 val을 제거한다.
// 집합에 val이 없는 경우에는 아무 일도 일어나지 않는다.
func (m StringSet) Erase(val string) {
	if _, exists := m[val]; !exists {
		return
	}
	m[val]--
	if m[val] <= 0 {
		delete(m, val)
	}
	return
}

// Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func (m StringSet) Count(val string) int {
	return m[val]
}

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈 칸으로
// 구분하여 넣은 문자열을 반환한다.
//func (m StringSet) String() string {
//	vals := ""
//	for val, count := range m {
//		for i := 0; i < count; i++ {
//			vals += fmt.Sprintf(" %s", val)
//		}
//	}
//	return fmt.Sprintf("{%s }", vals)
//	return "{" + vals + " }"
//}
