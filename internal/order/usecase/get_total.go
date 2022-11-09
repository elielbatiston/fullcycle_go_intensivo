package usecase

import "github.com/batistondeoliveira/fullcycle_go_intensivo/internal/order/entity"

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type GetTotalOutputDTO struct {
	Total int
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
