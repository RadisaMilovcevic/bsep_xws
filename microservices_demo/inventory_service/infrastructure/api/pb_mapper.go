package api

import (
	pb "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/inventory_service"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/inventory_service/domain"
)

func mapProduct(product *domain.Product) *pb.Product {
	productPb := &pb.Product{
		Id:        product.ProductId,
		ColorCode: product.ColorCode,
		Quantity:  int64(product.Quantity),
	}
	return productPb
}
