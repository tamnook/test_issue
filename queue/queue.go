package queue

import (
	"sync"
	"time"
)

type Issue struct {
	Status    string
	N         int
	D         float64
	N1        float64
	I         float64
	TTL       float64
	Iteration int
	AddTime   time.Time
	StartTime *time.Time
	EndTime   *time.Time
	Result    *float64
}

type IssueQueue struct {
	queue  []Issue
	locker *sync.RWMutex
}

func InitQueue() *IssueQueue {
	return &IssueQueue{
		queue: make([]Issue, 0),
	}
}
func (q *IssueQueue) GetIssues() []Issue {
	ans := make([]Issue, 0)
	i := 1
	for _, v := range q.queue {
		if v.EndTime == nil {
			ans = append(ans, v)
			i++
			continue
		}
		if time.Now().After(v.EndTime.Add(time.Duration(int(v.TTL*1000)) * time.Millisecond)) {
			ans = append(ans, v)
			i++
			continue
		}
	}
	return ans
}
func (q *IssueQueue) AddIssue(i Issue) {
	q.locker.Lock()
	q.queue = append(q.queue, i)
	q.locker.Unlock()
}
