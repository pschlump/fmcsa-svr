package metric

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../LICENSE.mit

import (
	"testing"
	//	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	//	m := NewMetrics()
func NewMetricsData(saveKey string, validKeys []MetricsTypeInfo, saveRateSeconds int, xgCfg *data.BaseConfigType, xdb map[string]bool, xlfp *os.File, xconn *pgxpool.Pool, xctx context.Context) (md *MetricsData) {
	//	assert.Equal(t, 2, m.q.SubmittedTasks())
	//	assert.Equal(t, 2, m.q.SuccessTasks())
}
