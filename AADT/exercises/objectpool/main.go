package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

type PoolObject interface {
	getID() string
}

type pool struct {
	idle   []PoolObject
	active []PoolObject
	cap    int
	mutex  *sync.Mutex
}

func initPool(poolObjects []PoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Can't create pool of 0 length")
	}

	if len(poolObjects) == 10 {
		return nil, fmt.Errorf("Can't create pool larger than 10 connections")
	}

	active := make([]PoolObject, 0)
	pool := &pool{
		idle:   poolObjects,
		active: active,
		cap:    len(poolObjects),
		mutex:  new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) register() (PoolObject, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object are free")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loaned object with id : %s\n", obj.getID())
	return obj, nil
}

func (p *pool) receive(target PoolObject) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Returned pool object with ID: %s\n", target.getID())
	return nil
}

func (p *pool) remove(target PoolObject) error {
	curActive := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[curActive-1], p.active[i] = p.active[i], p.active[curActive-1]
			p.active = p.active[:curActive-1]
			return nil
		}
	}
	return fmt.Errorf("Object doesn't belong to pool")
}

func main() {
	connections := make([]PoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("init pool error: %s", err)
	}
	conn1, err := pool.register()
	if err != nil {
		log.Fatalf("Pool register error: %s", err)
	}
	conn2, err := pool.register()
	if err != nil {
		log.Fatalf("Pool register error: %s", err)
	}
	pool.receive(conn1)
	pool.receive((conn2))
}
