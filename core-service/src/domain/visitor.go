package domain

import (
	"context"
)

type CounterVisitorResponse struct {
	Data CounterVisitor `json:"data"`
}

type ExternalCounterVisitor struct {
	Result CounterVisitor `json:"result"`
}

type CounterVisitor struct {
	Online30mnt          int `json:"online30mnt"`
	Visitor_today        int `json:"visitor_today"`
	Visit_today          int `json:"visit_today"`
	Visitor_yesterday    int `json:"visitor_yesterday"`
	Visit_yesterday      int `json:"visit_yesterday"`
	Visitor_all          int `json:"visitor_all"`
	Visit_all            int `json:"visit_all"`
	Growth_visitor       int `json:"growth_visitor"`
	Growth_visit         int `json:"growth_visit"`
	Last_update_pipeline int `json:"last_update_pipeline"`
}

type VisitorUsecase interface {
	GetCounterVisitor(ctx context.Context, path string) CounterVisitorResponse
}

type ExternalVisitorRepository interface {
	GetCounterVisitor() (ExternalCounterVisitor, error)
}
