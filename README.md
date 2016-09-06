# capsule.io

A Go library &amp; tool for helping you manage private information within code. 

##Usage

Here is an example capsule file
```
key=value
oranges=a sweet orange fruit
foo=bar

```

***

```Go
package main

import (
	"github.com/Digitalblueeye/capsule.io"
	"fmt"
)

func main() {
	/* 

	Since I didn't specify a capsule, capsuleio.Get() will use the first .capsule file found in the local directory
	If I would like to specify a capsule, I would need to use

	capsuleio.Open("mycapsulefile.capsule")

	*/

	fmt.Println(capsuleio.Get("key")) //prints value


}

```

##Installation
```
go get github.com/Digitalblueeye/capsule.io
```

Thanks for using Capsule! Please let me know if you have any questions, issues or simply want to request a new feature!
