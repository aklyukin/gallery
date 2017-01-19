//go:generate reform
package main

//reform:users
type User struct {
	ID       int32   `reform:"id,pk"`
	Email    string  `reform:"email"`
	Password string  `reform:"password"`
	RegDate  *string `reform:"reg_date"`
}

//reform:photos
type Photo struct {
	ID         int32  `reform:"id,pk"`
	PhotoName  string `reform:"photo_name"`
	FileName   string `reform:"filename"`
	CreateDate int    `reform:"create_date"`
}
