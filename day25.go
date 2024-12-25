package main

import (
	"fmt"
	"bufio"
	"os"
)

type Obj struct {
	lock bool
	count [5]int
	match int
}

func main() {
	fic,_ := os.Open("data")
	sca := bufio.NewScanner(fic)

	lo:=[]Obj{}

	for sca.Scan() {
		col_lock:=[5]int{-1,-1,-1,-1,-1}
		col_key:=[5]int{-1,-1,-1,-1,-1}
		finish_lock:=[5]bool{}
		finish_key:=[5]bool{}
		buf:=[7]string{}

		for i:=0; i< 7; i++ {
			text:=sca.Text()
			sca.Scan()
			buf[i]=text
		}

		for i:=0; i< 7; i++ {
			text:=buf[i]
			for j:=0; j < 5; j++ {
				if text[j]=='.' {
					finish_lock[j]=true
				}
				if finish_lock[j] { continue }

				col_lock[j]++
			}
		}

		for i:=6; i>=0; i-- {
			text:=buf[i]
			for j:=0; j < 5; j++ {
				if text[j]=='.' {
					finish_key[j]=true
				}
				if finish_key[j] { continue }

				col_key[j]++
			}
		}


		islock:=true
		for i:=0; i < 5; i++ {
			if col_lock[i] == -1 {
				islock=false
				break
			}
		}

		obj:=Obj{lock:islock,match:-1}
		if islock {
			obj.count = col_lock
		} else {
			obj.count = col_key
		}

		lo=append(lo, obj)
	}

	count:=0

	for c_l,ro_l:=range lo {
		if ro_l.lock==false {
			continue
		}

		for c_k,ro_k:=range lo {
			if ro_k.lock==true {
				continue
			}

			match:=true
			for i:=0; i < 5; i++ {
				if ro_l.count[i]+ro_k.count[i] > 5 {
					match=false
					break
				}
			}
			if match {
				ro_l.match=c_k
				ro_k.match=c_l
				count++
			}
		}
	}

	fmt.Println(count)
}
