package client

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/zcong1993/mailer/common"
)

// MongoLogger is struct of mongo driven logger
type MongoLogger struct {
	ch chan<- *common.MailLog
	c  *mgo.Collection
}

func indexMongo(c *mgo.Collection) error {
	err := c.EnsureIndex(mgo.Index{
		Key:        []string{"mailmsg.id"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	})
	if err != nil {
		return err
	}
	err = c.EnsureIndex(mgo.Index{
		Key: []string{"mailmsg.to"},
	})
	if err != nil {
		return err
	}
	err = c.EnsureIndex(mgo.Index{
		Key: []string{"mailmsg.tag"},
	})
	if err != nil {
		return err
	}
	return nil
}

// NewMongoLogger construct a MongoLogger instance
func NewMongoLogger(url, db, table string, buffer int) *MongoLogger {
	ch := make(chan *common.MailLog, buffer)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	err = indexMongo(c)
	if err != nil {
		panic(err)
	}
	go func() {
		for msg := range ch {
			err := c.Insert(msg)
			if err != nil {
				fmt.Printf("mongo logger err %+v\n", err)
			}
		}
	}()
	return &MongoLogger{
		ch: ch,
		c:  c,
	}
}

// GetChannel impl logger interface
func (ml *MongoLogger) GetChannel() chan<- *common.MailLog {
	return ml.ch
}
