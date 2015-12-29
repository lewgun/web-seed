package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Model the only model instance.
var M *model

type model struct {
	dsn     string
	chBatch chan *struct{}
	chWait  chan struct{}

	*xorm.Engine
}

func (m *model) init() error {

	var err error

	m.Engine, err = xorm.NewEngine("mysql", m.dsn)
	if err != nil {
		return err
	}

	m.Engine.ShowSQL = true
	m.Engine.SetMaxOpenConns(10)
	//m.Engine.ShowSQL = false
	return nil

}

func (m *model) run() {

	go func() {
		for range m.chBatch {
		}

		m.chWait <- struct{}{}
	}()
}

//Close 关闭功能
func (m *model) Close() {
	if m.chBatch == nil {
		return
	}

	close(m.chBatch)
	<-m.chWait

	m.chBatch = nil
}

//SharedInstance create the database's connection
func SharedInstance(dsn string, bufSize uint64) *model {

	// duplicate load.
	if M != nil {
		return M
	}

	M = &model{
		dsn:     dsn,
		chWait:  make(chan struct{}),
		chBatch: make(chan *struct{}, bufSize),
	}

	if err := M.init(); err != nil {
		return nil
	}

	//	M.run()

	return M
}
