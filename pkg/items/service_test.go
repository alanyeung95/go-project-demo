package items_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/alanyeung95/GoProjectDemo/pkg/items"
	"github.com/alanyeung95/GoProjectDemo/pkg/mocks/mock_items"
)

func TestCreateItem(t *testing.T) {
	RegisterFailHandler(Fail)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_items.NewMockRepository(mockCtrl)
	expectedResponse := &items.Item{}
	mockRepo.EXPECT().Upsert(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedResponse, nil)
	service := items.NewService(mockRepo)

	resp, err := service.CreateItem(context.TODO(), nil, &items.Item{ID: "123"})
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp).ShouldNot(BeNil())

}

func TestGetItemByID(t *testing.T) {
	RegisterFailHandler(Fail)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_items.NewMockRepository(mockCtrl)
	expectedResponse := &items.Item{}
	mockRepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(expectedResponse, nil)
	service := items.NewService(mockRepo)

	resp, err := service.GetItemByID(context.TODO(), nil, "any_id")
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp).ShouldNot(BeNil())
}
