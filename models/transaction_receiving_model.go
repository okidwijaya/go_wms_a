package models

import (
	"fmt"
	"log"
	"time"

	"github.com/okidwijaya/go_wms_a/config"
)

type ReceiverHeader struct {
	TrxInPK      int       `json:"TrxInPK"`
	TrxInNo      string    `json:"TrxInNo"`
	WhsIdf       int       `json:"WhsIdf"`
	TrxInDate    time.Time `json:"TrxInDate"`
	TrxInSuppIdf int       `json:"TrxInSuppIdf"`
	TrxInNotes   string    `json:"TrxInNotes"`
}

type ReceiverDetail struct {
	// TrxInDPK         int `json:"trxInDPK"`
	TrxInIDF         int `json:"trxInIDF"`
	TrxInDProductIdf int `json:"trxInDProductIdf"`
	TrxInDQtyDus     int `json:"trxInDQtyDus"`
	TrxInDQtyPcs     int `json:"trxInDQtyPcs"`
}

func TransactionReceiveHeader(header ReceiverHeader) (int64, error) {
	query := "INSERT INTO transaksi_penerimaan_barang_header (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes) VALUES (?, ?, ?, ?, ?)"
	result, err := config.DB.Exec(query, header.TrxInNo, header.WhsIdf, header.TrxInDate, header.TrxInSuppIdf, header.TrxInNotes)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}
	return result.LastInsertId()
}

func TransactionReceiveDetail(detail ReceiverDetail) error {
	query := "INSERT INTO transaksi_penerimaan_barang_detail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, detail.TrxInIDF, detail.TrxInDProductIdf, detail.TrxInDQtyDus, detail.TrxInDQtyPcs)
	if err != nil {
		log.Printf("Failed to execute query: %s with values: %v, %v, %v, %v", query, detail.TrxInIDF, detail.TrxInDProductIdf, detail.TrxInDQtyDus, detail.TrxInDQtyPcs)
	}
	return err
}
