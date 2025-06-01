package repo

import (
	"context"
	"database/sql"

	"go-clean-template/internal/usecase"
)

type pegawaiRepo struct {
	db *sql.DB
}

func NewPegawaiRepository(db *sql.DB) usecase.PegawaiRepository {
	return &pegawaiRepo{db: db}
}

func (r *pegawaiRepo) FetchAll(ctx context.Context) ([]usecase.Pegawai, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, nama, email, jabatan FROM pegawai")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pegawaiList []usecase.Pegawai
	for rows.Next() {
		var p usecase.Pegawai
		if err := rows.Scan(&p.ID, &p.Nama, &p.Email, &p.Jabatan); err != nil {
			return nil, err
		}
		pegawaiList = append(pegawaiList, p)
	}

	return pegawaiList, nil
}
