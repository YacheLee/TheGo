package main

func main() {
	println(*f() == *f())

	v := 1
	incr(&v)
	reduce(&v)
	println(v)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func reduce(v *int) int {
	*v--
	return *v
}
