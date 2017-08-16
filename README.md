# capsule.io

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/rucuriousyet/capsule.io)

The little library for fetching application secrets!

Notice: this library is meant to function as a development tool for those who post code in a remote or publicly viewed repo. Capsule.io is not intended to be used in final products or builds. 

## Usage

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
	"github.com/rucuriousyet/capsule.io"
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

## Installation
```
go get github.com/rucuriousyet/capsule.io
```

Thanks for using Capsule! Please let me know if you have any questions, issues or simply want to request a new feature!
