package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	alphabet := []string{"a", "i", "u", "e", "o", "ka", "ki", "ku", "ke", "ko", "ga", "gi", "gu", "ge", "go", "sa", "shi", "su", "se", "so", "za", "ji", "su", "ze", "zo", "ta", "chi", "tsu", "te", "to", "da", "dji", "dzu", "de", "do", "na", "ni", "nu", "ne", "no", "ha", "hi", "fu", "he", "ho", "ba", "bi", "bu", "be", "bo", "pa", "pi", "pu", "po", "ma", "mi", "mu", "me", "mo", "ya", "yu", "yo", "ra", "ri", "ru", "re", "ro", "wa", "wi", "we", "wo", "n", "vu", "gya", "gyu", "gyo", "ja", "ju", "jo", "nya", "nyu", "nyo", "bya", "byu", "byo", "mya", "myu", "myo", "rya", "ryu", "ryo", "pya", "pyu", "pyo"}

	rand.Seed(time.Now().UnixNano())
	result := make([]string, 3)
	for index := 0; index < 3; index++ {
		randIndex := rand.Intn(len(alphabet))
		result[index] = alphabet[randIndex]
	}
	fmt.Println(result)
}
