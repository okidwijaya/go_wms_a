package models

import (
	"fmt"
	"log"
	"time"

	"github.com/okidwijaya/go_wms_a/config"
)

type DispatchingHeader struct {
	TrxOutPK      int       `json:"TrxOutPK"`
	TrxOutNo      string    `json:"TrxOutNo"`
	Whsidf        int       `json:"WhsIdf"`
	TrxOutDate    time.Time `json:"TrxOutDate"`
	TrxOutSuppIdf int       `json:"TrxOutSuppIdf"`
	TrxOutNotes   string    `json:"TrxOutNotes"`
}

type DispatchingDetail struct {
	TrxOutIDF         int `json:"trxOutIDF"`
	TrxOutDProductIdf int `json:"trxOutDProductIdf"`
	TrxOutDQtyDus     int `json:"trxOutDQtyDus"`
	TrxOutDQtyPcs     int `json:"trxOutDQtyPcs"`
}

func TransactionDispatchingHeader(header DispatchingHeader) (int64, error) {
	query := "INSERT INTO transaksi_pengeluaran_barang_header (TrxOutNo, Whsidf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes) VALUES (?, ?, ?, ?, ?)"
	result, err := config.DB.Exec(query, header.TrxOutNo, header.Whsidf, header.TrxOutDate, header.TrxOutSuppIdf, header.TrxOutNotes)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}
	return result.LastInsertId()
}

func TransactionDispatchingDetail(detail DispatchingDetail) error {
	query := "INSERT INTO transaksi_pengeluaran_barang_detail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, detail.TrxOutIDF, detail.TrxOutDProductIdf, detail.TrxOutDQtyDus, detail.TrxOutDQtyPcs)
	if err != nil {
		log.Printf("Failed to execute query: %s with values: %v, %v, %v, %v", query, detail.TrxOutIDF, detail.TrxOutDProductIdf, detail.TrxOutDQtyDus, detail.TrxOutDQtyPcs)
	}
	return err
}
