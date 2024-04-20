package RevwBL

import (
	"fmt"
	"revw/Logger"
	"time"
)

type ReviewSession struct {
	ID               string // unique hash
	creationDateTime time.Time
	reviewText       *string
	readers          map[string]Reader
}

func NewReviewSession(text *string) *ReviewSession {
	creationTime := time.Now()

	reviewSession := ReviewSession{
		HashString(creationTime.String() + *text),
		creationTime,
		text,
		make(map[string]Reader),
	}

	Logger.GetLogger().Printf("Created a new review session with ID '%v'", reviewSession.ID)

	return &reviewSession
}

func (rs *ReviewSession) AddReader() *Reader {
	readerID := HashString(fmt.Sprint(len(rs.readers)) + time.Now().String())

	_, ok := rs.readers[readerID]

	if ok {
		Logger.GetLogger().Printf("A reader with ID '%v' already exists in session '%v'\n", readerID, rs.ID)
		return nil
	}

	reader := NewReader(readerID)
	rs.readers[readerID] = *reader

	Logger.GetLogger().Printf("A new reader with ID '%v' joined the session '%v'\n", readerID, rs.ID)

	return reader
}

func (rs *ReviewSession) GetReaderByID(id string) (*Reader, bool) {
	reader, ok := rs.readers[id]
	return &reader, ok
}
