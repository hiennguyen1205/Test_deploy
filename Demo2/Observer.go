package main

// Interface đối tượng nào muốn nhận được thông báo sẽ phải tuân thủ
type Observer interface {
	//Thông tin cần cập nhật ở đây là giá trị int
	Update(value int)
}
