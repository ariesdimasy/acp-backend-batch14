package answer

type Person struct {
	Name       string
	NameEncode string
	Score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

/*
	r st u
	i jk l
	z ab c

*/

func (s *Person) Encode() string {
	var nameEncode = ""

	//var alphabet = "abcdefghijklmnopqrstuvwxyz" // 97 + 25 = 122
	// 65 + 25 = 90

	for i := 0; i < len(s.Name); i++ {
		var res = int(s.Name[i]) + 3
		if res > 122 {
			res = 97 + (res - 122 - 1)
		}

		nameEncode = nameEncode + string(res)
	}

	return nameEncode
}

/*
	u ts r
	l kj i
	z ab c

*/

func (s *Person) Decode() string {
	var nameEncode = ""

	for i := 0; i < len(s.Name); i++ {
		var res = int(s.Name[i]) - 3 // 99 - 3 = 96
		if res < 96 {
			res = 122 - ((97 - res) - 1)
			println(" res => ", res)
		}

		nameEncode = nameEncode + string(res)
	}

	return nameEncode
}
