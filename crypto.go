//Please dont recode my code
//Copyright© 2019 by Fajar Firdaus

package main

import (
"fmt"
"time"
"github.com/mgutz/ansi"
"crypto/md5"
"crypto/sha256"
"crypto/sha512"
"golang.org/x/crypto/md4"
"golang.org/x/crypto/ripemd160"
"io"
)

func main(){
	a1 := ansi.Color("\n\nMD5 : \n", "red")
	a2 := ansi.Color("\n\nSHA256 : \n", "red")
	a3 := ansi.Color("\n\nSHA512 : \n", "red")
	a4 := ansi.Color("\n\nMD4 : \n", "red")
	a5 := ansi.Color("\n\nRIPEMD160 : \n", "red")
	
	
	title := ansi.Color("Cryptografi By Fajar Firdaus", "green+b:white")
	x := ansi.Color("[===================]\n", "cyan")
	l := ansi.Color("Coder : Fajar Firdaus\n", "blue")
	s := ansi.Color("FB : Fajar Firdaus\n", "blue")
	rf := ansi.Color("IG : fajar_firdaus_7\n", "blue")
	p := ansi.Color("YT : iTech7732\n", "blue")
	y := ansi.Color("[===================]\n", "cyan")
	
	fmt.Print(title)
	fmt.Print("\n")
	fmt.Print(x)
	fmt.Print(l)
	fmt.Print(s)
	fmt.Print(rf)
	fmt.Print(p)
	fmt.Print(y)
	fmt.Print("\n")
	
	var in string
	i := ansi.Color("Input Text", "blue")
	inw := ansi.Color("[", "green")
	ins := ansi.Color("]", "green")
	fmt.Print(inw + i + ins + " " + " :")
	fmt.Scan(&in)
	
	var waktu = time.Now()
	a := ansi.Color("Time : ", "green")
	k := ansi.Color("[", "red")
	n := ansi.Color("]", "red")
    o := md5.New()
    io.WriteString(o, in)
	fmt.Print(k, a, waktu.Hour(), "-", waktu.Minute(), "-", waktu.Second(), n, " ", "Result : ")
	fmt.Print(a1)
	fmt.Printf("%x", o.Sum(nil))
	
	sha := sha256.New()
	io.WriteString(sha, in)
	fmt.Print(a2)
	fmt.Printf("%x", sha.Sum(nil))
	
	sha512 := sha512.New()
	io.WriteString(sha512, in)
	fmt.Print(a3)
	fmt.Printf("%x", sha512.Sum(nil))
	
	md4 := md4.New()
	io.WriteString(md4, in)
	fmt.Print(a4)
	fmt.Printf("%x", md4.Sum(nil))
	
	ripemd160 := ripemd160.New()
	io.WriteString(ripemd160, in)
	fmt.Print(a5)
	fmt.Printf("%x", ripemd160.Sum(nil))
	}
