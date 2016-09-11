# capsule.io

A Go library &amp; tool for helping you manage private information within code. 
Notice: this library is meant to function as a development tool for those who post code in a remote or publicly viewed repo. Capsule.io is not intended to be used in final products or builds. 

Capsule.io was inspired by some functions in Laravel (PHP) and is extremely simple.

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
	"github.com/archit3cture/capsule.io"
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


##Documentation

#### func  Get

```go
func Get(key string) string
```
Gets the string value from the current capsule

#### func  Open

```go
func Open(capsule string)
```
Sets the source file for the key value store (aka the capsule)
