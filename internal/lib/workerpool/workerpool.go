package workerpool

import (
	"log"
	"sync"
)

// Job интерфейс для задач
type Job interface {
	Process() error
}

// JobProcessor интерфейс для обработчиков задач
type JobProcessor interface {
	Process(job Job) error
}

// WorkerPool структура для Worker Pool
type WorkerPool struct {
	jobs    chan Job
	workers int
	wg      sync.WaitGroup
}

// NewWorkerPool создает новый WorkerPool
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan Job, 100),
		workers: workers,
	}
}

// Start запускает воркеры
func (p *WorkerPool) Start(processor JobProcessor) {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker(processor)
	}
}

// worker обрабатывает задачи
func (p *WorkerPool) worker(processor JobProcessor) {
	defer p.wg.Done()
	for job := range p.jobs {
		if err := processor.Process(job); err != nil {
			log.Printf("Failed to process job: %v", err)
		}
	}
}

// AddJob добавляет задачу в пул
func (p *WorkerPool) AddJob(job Job) {
	p.jobs <- job
}

// Stop останавливает пул
func (p *WorkerPool) Stop() {
	close(p.jobs)
	p.wg.Wait()
}
