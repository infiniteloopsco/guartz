package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	rnd "math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

//AllUniqueWord will return a bool value based on if the word contains
//unique characters or not.
func AllUniqueWord(s string) bool {
	strSlice := strings.Split(s, "")
	resultSl := make(map[string]int)

	for _, v := range strSlice {
		resultSl[v]++
	}

	for _, v := range resultSl {
		if v > 1 {
			return false
		}
	}
	return true
}

//ReverseStr returns the string reversed
func ReverseStr(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

//StrPermutation checks if two strings has the same elements in the same amount
func StrPermutation(a string, b string) bool {
	ma := make(map[string]int)
	mb := make(map[string]int)
	aSl := strings.Split(a, "")
	bSl := strings.Split(b, "")

	for _, v := range aSl {
		ma[v]++
	}
	for _, v := range bSl {
		mb[v]++
	}

	eq := reflect.DeepEqual(ma, mb)

	if eq {
		return true
	}

	return false

}

//ReplaceSpacesWSymb expects two strings, first arg being the string itself
//Second argument is the symbol to replace spaces
//E.G: s := "emir ozer" symb := "%20"  / Result will be "emir%20ozer"
func ReplaceSpacesWSymb(s string, symb string) (result string) {
	sp := strings.Fields(s)
	var tSp string
	for _, v := range sp {
		tSp += v + symb
	}
	return tSp[:len(tSp)-len(symb)]
}

//PseudoUUID generates a uuid like string
func PseudoUUID() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

//RandomString gets the desired length and returns a random string
func RandomString(l int) string {
	rnd.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

//RandInt returns a random integer between the min and max you set
func RandInt(min int, max int) int {
	rnd.Seed(time.Now().UnixNano())
	return min + rnd.Intn(max-min)
}

//DiffSlices takes two integer slices and returns the diff between them
func DiffSlices(X, Y []int) []int {
	m := make(map[int]int)

	for _, y := range Y {
		m[y]++
	}

	var ret []int
	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

//MD5 encription
func MD5(cad string) string {
	hash := sha1.New()
	hash.Write([]byte(cad))
	return hex.EncodeToString(hash.Sum(nil))
}

//FileExists verifies if a filepath exists
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
