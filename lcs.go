package main

/*
	rootVIII
	Longest and shortest common subsequences of 8-bit integers
*/

// Lcs stores the longest found slice sequences as well as the current slice.
type Lcs struct {
	foundSequences [][]uint8
	current        []uint8
	sortedList     []uint8
}

func (l *Lcs) storeByte(incoming byte) {
	l.current = append(l.current, incoming)
}

func (l *Lcs) reestablish(incoming byte) {
	l.current = append(l.current, incoming)
	l.foundSequences = append(l.foundSequences, l.current)
	l.current = nil
}

// GetLongest returns the largest list in foundSequences.
func (l Lcs) GetLongest() []uint8 {
	largest, found := 0, 0
	for index, slice := range l.foundSequences {
		length := len(slice)
		if length > largest {
			largest = length
			found = index
		}
	}
	return l.foundSequences[found]
}

// GetShortest returns the smallest list found in foundSequences.
func (l Lcs) GetShortest() []uint8 {
	smallest := int(^uint32(0))
	var found int
	for index, slice := range l.foundSequences {
		length := len(slice)
		if length < smallest {
			smallest = length
			found = index
		}
	}
	return l.foundSequences[found]
}

// GetAll returns all foundSequences.
func (l Lcs) GetAll() [][]uint8 {
	return l.foundSequences
}

// SortSequence sorts a valid uint8 slice: values 0 through 255.
func (l *Lcs) SortSequence(unsorted []uint8) {
	l.sortedList = MergeSort(unsorted)
	for i := 0; i < len(l.sortedList); i++ {
		if i == len(l.sortedList)-1 {
			if l.sortedList[i-1] == l.sortedList[i-1]+1 || l.sortedList[i-1] == l.sortedList[i] {
				l.storeByte(l.sortedList[i])
			} else {
				l.reestablish(l.sortedList[i])
			}
			break
		}
		if l.sortedList[i+1] == l.sortedList[i]+1 || l.sortedList[i] == l.sortedList[i+1] {
			l.storeByte(l.sortedList[i])
		} else {
			l.reestablish(l.sortedList[i])
		}
	}
	if len(l.current) > 0 {
		l.foundSequences = append(l.foundSequences, l.current)
	}
}
