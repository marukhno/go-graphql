package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/marukhno/go-graphql/graph/generated"
	"github.com/marukhno/go-graphql/graph/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	mockAcc := model.Account{
		ID: input.ID,
		Name: &model.Name{
			First: input.Name.First,
			Last:  input.Name.Last,
		},
		ShippingAddress: &model.Address{
			Country:  input.ShippingAddress.Country,
			Street:   input.ShippingAddress.Street,
			State:    input.ShippingAddress.State,
			Zip:      input.ShippingAddress.Zip,
			Building: input.ShippingAddress.Building,
		},
		CreditCard: &model.CreditCard{
			Number:         input.CreditCard.Number,
			Pin:            input.CreditCard.Pin,
			ExpirationDate: input.CreditCard.ExpirationDate,
		},
	}
	Accounts[input.ID] = &mockAcc
	return &mockAcc, nil
}

func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	if acc, ok := Accounts[id]; ok {
		return acc, nil
	} else {
		return nil, nil
	}
}

func (r *queryResolver) Accounts(ctx context.Context, limit *int) ([]*model.Account, error) {
	accArray := make([]*model.Account, 0, len(Accounts))
	for _,v := range Accounts {
		accArray = append(accArray, v)
	}
	l := *limit
	if l > len(accArray) {
		l = len(accArray)
	}

	return accArray[:l], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
