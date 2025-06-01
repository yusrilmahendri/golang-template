package usecase

import "context"

type Pegawai struct {
	ID      int
	Nama    string
	Email   string
	Jabatan string
}

type PegawaiRepository interface {
	FetchAll(ctx context.Context) ([]Pegawai, error)
}

type PegawaiUsecase interface {
	GetAllPegawai(ctx context.Context) ([]Pegawai, error)
}

type pegawaiUsecase struct {
	repo PegawaiRepository
}

func NewPegawaiUsecase(r PegawaiRepository) PegawaiUsecase {
	return &pegawaiUsecase{repo: r}
}

func (u *pegawaiUsecase) GetAllPegawai(ctx context.Context) ([]Pegawai, error) {
	return u.repo.FetchAll(ctx)
}
