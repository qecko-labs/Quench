/*
 *   Copyright (c) 2026 qecko-labs
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package ch

import (
	"sync/atomic"
)

type ringSlot struct {
	sequence uint64
	val      any
}

type MPSC struct {
	cap   uint64
	mask  uint64
	head  uint64
	tail  atomic.Uint64
	slots []ringSlot
}

func NewMPSC(capPow2 int) *MPSC {
	cap := uint64(1)
	for cap < uint64(capPow2) {
		cap <<= 1
	}
	s := &MPSC{cap: cap, mask: cap - 1, slots: make([]ringSlot, cap)}
	for i := range s.slots {
		s.slots[i].sequence = uint64(i)
	}
	return s
}

func (q *MPSC) Enqueue(v any) bool {
	tail := q.tail.Add(1) - 1
	for {
		idx := tail & q.mask
		slot := &q.slots[idx]
		seq := atomic.LoadUint64(&slot.sequence)
		if seq == tail {
			slot.val = v
			atomic.StoreUint64(&slot.sequence, tail+1)
			return true
		}
	}
}

func (q *MPSC) Dequeue() (any, bool) {
	if q == nil || len(q.slots) == 0 {
		return nil, false
	}
	head := q.head
	idx := head & q.mask
	slot := &q.slots[idx]
	seq := atomic.LoadUint64(&slot.sequence)
	if seq != head+1 || slot.val == nil {
		return nil, false
	}
	v := slot.val
	slot.val = nil
	atomic.StoreUint64(&slot.sequence, head+q.cap)
	q.head = head + 1
	return v, true
}
