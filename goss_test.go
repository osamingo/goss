package goss

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestStruct struct {
	T int64
	P string
}

func (t *TestStruct) Target() int64 {
	return t.T
}

func (t *TestStruct) Priority() string {
	return t.P
}

func TestSortUtils(t *testing.T) {

	Convey("Declare SortedSlice By ASC", t, func() {

		s := &SortedSlice{}

		Convey("Add elements", func() {

			s.Add(&TestStruct{T: 100, P: "a"})
			s.Add(&TestStruct{T: 10, P: "b"})
			s.Add(&TestStruct{T: 10, P: "a"})
			s.Add(nil)

			Convey("Then SortedSlice is sorted by ASC", func() {

				ss := s.S
				So(len(ss), ShouldEqual, 3)

				var bts *TestStruct
				for _, i := range ss {
					ts := i.(*TestStruct)
					So(ts, ShouldHaveSameTypeAs, &TestStruct{})
					if bts != nil {
						So(ts.Target(), ShouldBeGreaterThanOrEqualTo, bts.Target())
						if ts.Target() == bts.Target() {
							So(ts.Priority(), ShouldBeGreaterThanOrEqualTo, bts.Priority())
						}
					}
					bts = ts
				}
			})
		})

	})

	Convey("Declare SortedSlice By DESC", t, func() {

		s := &SortedSlice{DESC: true}

		Convey("Add elements", func() {

			s.Add(&TestStruct{T: 10, P: "b"})
			s.Add(&TestStruct{T: 10, P: "a"})
			s.Add(&TestStruct{T: 100, P: "a"})
			s.Add(nil)

			Convey("Then SortedSlice is sorted by DESC", func() {

				ss := s.S
				So(len(ss), ShouldEqual, 3)

				var bts *TestStruct
				for _, i := range ss {
					ts := i.(*TestStruct)
					So(ts, ShouldHaveSameTypeAs, &TestStruct{})
					if bts != nil {
						So(ts.Target(), ShouldBeLessThanOrEqualTo, bts.Target())
						if ts.Target() == bts.Target() {
							So(ts.Priority(), ShouldBeGreaterThanOrEqualTo, bts.Priority())
						}
					}
					bts = ts
				}
			})
		})

	})

}
