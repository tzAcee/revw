package RevwBL

import (
	"fmt"
	"revw/Logger"
	"time"
)

type Comment struct {
	ID    string
	text  string
	index uint
}

func newComment(text string, index uint) *Comment {
	return &Comment{HashString(text + fmt.Sprint(index) + time.Now().String()), text, index}
}

type Reader struct {
	ID       string
	comments map[string]Comment
}

func NewReader(id string) *Reader {
	return &Reader{id, make(map[string]Comment)}
}

func (r *Reader) AddComment(text string, index uint) *Comment {

	c := newComment(text, index)

	_, ok := r.comments[c.ID]

	if ok {
		Logger.GetLogger().Printf("A comment with ID '%v' already exists from reader with ID '%v'\n", c.ID, r.ID)
		return nil
	}

	r.comments[c.ID] = *c

	Logger.GetLogger().Printf("A new comment with ID '%v' was added from reader with ID '%v'\n", c.ID, r.ID)

	return c
}

func (r *Reader) DeleteComment(id string) {
	_, ok := r.comments[id]

	if ok {
		delete(r.comments, id)
	}
}

func (r *Reader) EditComment(id string, newText string) error {
	comment, ok := r.comments[id]

	if ok {
		comment.text = newText
	} else {
		return fmt.Errorf("comment with ID '%v' does not exist", id)
	}

	return nil
}
