package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// nonce := generateNonce()
	// fmt.Printf("nonce is %s \n", nonce)
	// checkNonce()
	generateNonce()
}

func checkNonce() bool {
	secret := "zhenzu"
	s := "zOWrTIbplg,1410495132,a72d07dd9cc66ddbe8034bc47b797ced4a43ce42"
	arr := strings.Split(s, ",")
	fmt.Println("arr is ", arr)
	salt := arr[0]
	maxTime, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	hash := arr[2]
	fmt.Println("salt is ", salt)
	fmt.Println("maxTime is ", maxTime)
	fmt.Println("hash is ", hash)

	back := fmt.Sprintf("%x", sha1.Sum([]byte(salt+secret+strconv.FormatInt(maxTime, 10))))
	fmt.Println("back is ", back)
	if back != hash {
		fmt.Println("invalide nonce")
		return false
	}
	if time.Now().Unix() > maxTime {
		fmt.Println("expired nonce")
		return false
	}
	fmt.Println("this is a valid nonce")
	return true
}

func generateNonce() string {
	secret := "zhenzu"
	salt := generateSalt()
	fmt.Printf("salt is %s \n", salt)
	timeStamp := time.Now().Unix()
	maxTime := timeStamp + 1
	fmt.Printf("maxTime is %d \n", maxTime)
	nonce := salt + "," + strconv.FormatInt(maxTime, 10) + "," + fmt.Sprintf("%x", sha1.Sum([]byte(salt+secret+strconv.FormatInt(maxTime, 10))))
	// nonce := salt + strconv.FormatInt(maxTime, 10)
	fmt.Println("nonce is :", nonce)
	return nonce
}

func generateSalt() string {
	//length of salt
	saltLength := 10
	rawChar := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	rawLength := utf8.RuneCountInString(rawChar)
	//salt
	salt := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for utf8.RuneCountInString(salt) < saltLength {
		salt += string(rawChar[r.Intn(rawLength)])
	}
	return salt
}
