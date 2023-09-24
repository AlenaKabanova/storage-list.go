package list

import "fmt"
import "math/rand"

type List struct {
	len       int64
	firstNode *node
	id int64
}

// NewList создает новый список
func NewList() *List {
	l := &List {
		len: 0,
		firstNode: nil,
		rand.Seed(time.Now().Unix()),
	    id := rand.Int63(),
	}
	return l
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	id = pkg.AI.GetID()
	n := node{
		index: id,
		data: data,
		nextNode: nil
	}

	if l.firstNode = nil {
		l.firstNode = &n
		l.len 
	}
	return index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	index = nil
	cn = fn
	for cn != nill {
		if cn.id = 3     //удаляется элемент с индексом 3
		if pn == nil {
			l.fn == cn.nn   //l относится к list, лист
			return 
		} 
		pn.nn = cn.nn 
		return
	}
	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	cn := l.firstNode
	pn := l.firstNode
	for cn != nil {
		if current.value == value {
			if cn == l.firstNode {
				l.firstNode = c.next
			}
		else {
			p.next = c.next
		}
		l.len = l.len - 1
	}
	pn = cn
	cn = c.next
}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	fn = nil
	cn = fn
	for cn != nill {
		for curvalue == value {
			l.fn == cn.nn 
			l.len = l.len -1
			return
		}
		pn.nn = cn.nn
		return
	}
	return
}


// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	cn = l.firstNode
	curindex = index
	for cn != nil {
		if curindex == index {
			return value
		}
		else {
			return ok
		}
	}
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	fn = nil
	cn = fn 
	firstindex = id
	curindex = firstindex
	for cn != nil { 
		if curvalue == value {
			return id
		}
	    else {
			return ok
		}

	}
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	fn = nil
	cn = fn 
	firstindex = id
	curindex = firstindex
	ids = id
	for cn != nil { 
		for curvalue == value {
			return ids
		}
	    else {
			return ok
		}
	}
	return
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	value = l.cn = nil && (l.len = 0)
	if value {
		return values
	}
	else {
		return ok
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil
	len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	cn = l.firstNode
	for cn != nill {
		fmt.Println (data, "  ", index)
		cn = l.nextNode
	}
}























