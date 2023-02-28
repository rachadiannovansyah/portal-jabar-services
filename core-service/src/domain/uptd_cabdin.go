package domain

import "context"

type UptdCabdin struct {
	ID         int64  `json:"id"`
	PrkName    string `json:"prk_name"`
	CbgName    string `json:"cbg_name"`
	CbgKotaKab string `json:"cbg_kota_kab"`
	CbgAlamat  string `json:"cbg_alamat"`
	CbgNoTlp   string `json:"cbg_no_tlp"`
	CbgJenis   string `json:"cbg_jenis"`
}

type UptdCabdinUsecase interface {
	Fetch(ctx context.Context) (res []UptdCabdin, err error)
}

type UptdCabdinRepository interface {
	Fetch(ctx context.Context) (res []UptdCabdin, err error)
}
