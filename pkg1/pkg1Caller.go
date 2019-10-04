package pkg1

import "pkg2"

func HelloInPkg1() {
	hello()
	pkg2.Hello()
}
