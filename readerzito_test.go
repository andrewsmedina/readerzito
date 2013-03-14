package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct{}

var _ = gocheck.Suite(&S{})

func (s *S) TestSubscribe(c *gocheck.C) {
	session, err := mgo.Dial("127.0.0.1:27017")
	c.Assert(err, gocheck.IsNil)
	defer session.Close()
	collection = session.DB("test").C("readerzito")
	err = subscribe("andrews medina", "http://andrewsmedina.com/rss/")
	c.Assert(err, gocheck.IsNil)
	var f feed
	err = collection.Find(bson.M{"name": "andrews medina"}).One(&f)
	c.Assert(err, gocheck.IsNil)
}
