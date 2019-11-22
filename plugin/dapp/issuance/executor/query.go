// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/33cn/chain33/types"
	pty "github.com/33cn/plugin/plugin/dapp/issuance/types"
)

func (c *Issuance) Query_IssuanceInfoByID(req *pty.ReqIssuanceInfo) (types.Message, error) {
	issu,err := queryIssuanceByID(c.GetStateDB(), req.IssuanceId)
	if err != nil {
		clog.Error("Query_IssuanceInfoByID", "id", req.IssuanceId, "error", err)
		return nil, err
	}

	return &pty.RepIssuanceCurrentInfo{
		Status:             issu.Status,
		TotalBalance:       issu.TotalBalance,
		DebtCeiling:        issu.DebtCeiling,
		LiquidationRatio:   issu.LiquidationRatio,
		Balance:            issu.Balance,
		CollateralValue:    issu.CollateralValue,
		DebtValue:          issu.DebtValue,
		Period:             issu.Period,
		IssuId:             issu.IssuanceId,
		CreateTime:         issu.CreateTime,
	}, nil
}

func (c *Issuance) Query_IssuanceInfoByIDs(req *pty.ReqIssuanceInfos) (types.Message, error) {
	infos := &pty.RepIssuanceCurrentInfos{}
	for _, id := range req.IssuanceIds {
		issu,err := queryIssuanceByID(c.GetStateDB(), id)
		if err != nil {
			clog.Error("Query_IssuanceInfoByID", "id", id, "error", err)
			return nil, err
		}

		infos.Infos = append(infos.Infos, &pty.RepIssuanceCurrentInfo{
			Status:             issu.Status,
			TotalBalance:       issu.TotalBalance,
			DebtCeiling:        issu.DebtCeiling,
			LiquidationRatio:   issu.LiquidationRatio,
			Balance:            issu.Balance,
			CollateralValue:    issu.CollateralValue,
			DebtValue:          issu.DebtValue,
			Period:             issu.Period,
			IssuId:             issu.IssuanceId,
			CreateTime:         issu.CreateTime,
		})
	}

	return infos, nil
}

func (c *Issuance) Query_IssuanceByStatus(req *pty.ReqIssuanceByStatus) (types.Message, error) {
	ids := &pty.RepIssuanceIDs{}
	issuIDRecords, err := queryIssuanceByStatus(c.GetLocalDB(), req.Status, req.Index)
	if err != nil {
		clog.Error("Query_IssuanceByStatus", "get issuance error", err)
		return nil, err
	}
	ids.IDs = append(ids.IDs, issuIDRecords...)

	return ids, nil
}

func (c *Issuance) Query_IssuanceRecordByID(req *pty.ReqIssuanceDebtInfo) (types.Message, error) {
	ret := &pty.RepIssuanceDebtInfo{}
	issuRecord, err := queryIssuanceRecordByID(c.GetStateDB(), req.IssuanceId, req.DebtId)
	if err != nil {
		clog.Error("Query_IssuanceRecordByID", "get issuance record error", err)
		return nil, err
	}

	ret.Record = issuRecord
	return ret, nil
}

func (c *Issuance) Query_IssuanceRecordsByAddr(req *pty.ReqIssuanceRecordsByAddr) (types.Message, error) {
	ret := &pty.RepIssuanceRecords{}
	records, err := queryIssuanceRecordByAddr(c.GetStateDB(), c.GetLocalDB(), req.Addr, req.Index)
	if err != nil {
		clog.Error("Query_IssuanceDebtInfoByAddr", "get issuance record error", err)
		return nil, err
	}

	if req.Status == 0 {
		ret.Records = records
	} else {
		for _,record := range records {
			if record.Status == req.Status {
				ret.Records = append(ret.Records, record)
			}
		}
	}

	return ret, nil
}

func (c *Issuance) Query_IssuanceRecordsByStatus(req *pty.ReqIssuanceRecordsByStatus) (types.Message, error) {
	ret := &pty.RepIssuanceRecords{}
	records, err := queryIssuanceRecordsByStatus(c.GetStateDB(), c.GetLocalDB(), req.Status, req.Index)
	if err != nil {
		clog.Error("Query_IssuanceDebtInfoByStatus", "get issuance record error", err)
		return nil, err
	}

	ret.Records = append(ret.Records, records...)
	return ret, nil
}

func (c *Issuance) Query_IssuancePrice(req *pty.ReqIssuanceRecordsByStatus) (types.Message, error) {
	price, err := getLatestPrice(c.GetStateDB())
	if err != nil {
		clog.Error("Query_CollateralizePrice", "error", err)
		return nil, err
	}

	return &pty.RepIssuancePrice{Price:price}, nil
}