// package trigger provides builders for hx-trigger attribute values.
package trigger

// A Trigger is a builder for hx-trigger attribute values.
type Trigger interface {
	String() string
	trigger()
}
