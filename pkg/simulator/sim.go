package simulator

import (
    "container/heap"
)

type Event struct { At int64; F func() }

type Sim struct {
    Now int64
    pq  eventPQ
}

func (s *Sim) After(delta int64, f func()) { heap.Push(&s.pq, Event{At: s.Now+delta, F: f}) }
func (s *Sim) RunUntil(t int64) {
    for s.pq.Len()>0 && s.pq[0].At<=t {
        ev:=heap.Pop(&s.pq).(Event)
        s.Now=ev.At
        ev.F()
    }
}

type eventPQ []Event

func (pq eventPQ) Len() int { return len(pq) }
func (pq eventPQ) Less(i, j int) bool { return pq[i].At < pq[j].At }
func (pq eventPQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *eventPQ) Push(x interface{}) { *pq = append(*pq, x.(Event)) }
func (pq *eventPQ) Pop() interface{} {
    old := *pq; n := len(old); x := old[n-1]; *pq = old[0:n-1]; return x
}