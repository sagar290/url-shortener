package Structs

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// func GenerateToken(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }

// func GenerateToken(length int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	b := make([]byte, length)
// 	rand.Read(b)
// 	return fmt.Sprintf("%x", b)[:length]
// }
