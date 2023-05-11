package objPool

import (
	"errors"
	"time"
)

type ReusableObj struct {
	// 可以替换为数据库连接类、网络连接类等等
}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

// 创建对象池
func NewObjPool(numOfObj int) *ObjPool {
	m_objPool := ObjPool{}
	m_objPool.bufChan = make(chan *ReusableObj, numOfObj) // buffered channel
	for i := 0; i < numOfObj; i++ {
		m_objPool.bufChan <- &ReusableObj{}
	}
	return &m_objPool
}

// 从对象池中获取对象
// 如果想要一个对象池里面存放不同类型的对象，可以使用空接口 Interface{}
// 但在实际应用中并不推荐这样做，最好每个对象池只存放一种类型的对象
func (p *ObjPool) GetObjPool(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

// 向对象池中放回对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
