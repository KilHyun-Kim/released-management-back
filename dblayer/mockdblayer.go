package dblayer

import (
	"encoding/json"
	"kilhyun-kim/released-back/models"
)

type MockDBLayer struct {
	err      error
	technics []models.Technic
}

func NewMockDBLayer(technics []models.Technic) *MockDBLayer {
	return &MockDBLayer{
		technics: technics,
	}
}

func NewMockDBLayerWithData() *MockDBLayer {
	TECHNICS := `[
		{
			"TechId":1,
			"Image": "리액트 이미지 URL",
			"ImageAlt": "리액트 이미지 ALT",
			"TechName": "React",


		},
		{
			"TechId": 2,
			"Image": "뷰 이미지 URL",
			"ImageAlt": "뷰 이미지 ALT",
			"TechName": "Vue",
		}		
	]`
	var technic []models.Technic
	json.Unmarshal([]byte(TECHNICS), &technic)
	return NewMockDBLayer(technic)
}

func (mock *MockDBLayer) GetMockTechnicData() []models.Technic {
	return mock.technics
}

func (mock *MockDBLayer) SetError(err error) {
	mock.err = err
}

// MockDBLayer 에 DBLayer 인터페이스의 메서드 구현
func (mock *MockDBLayer) GetAllTechs() ([]models.Technic, error) {
	// 에러 반환
	if mock.err != nil {
		return nil, mock.err
	}
	return mock.technics, nil
}
