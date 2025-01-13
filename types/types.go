package types

// 로그인 요청으로 LoginReq를 받을 것임
// body에 name을 받음, 필수요소이기 때문에 binding:"required" 조건을 줌
// binding:"required" 조건이 없으면 name의 값이 없어도 되는데, 조건이 달려있으면 필수적으로 값이 있어야 함.
type LoginReq struct {
	Name string `json:"name" binding:"required"`
}
