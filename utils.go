// *
// * Copyright 2013, Scott Cagno, All rights reserved
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * cached :: utils.go :: utilities, helpers, and globals
// *

package cached

// Request struct
type Request struct {
	Cmd string      `json:"cmd"`
	Key string      `json:"key"`
	Val interface{} `json:"val"`
}
