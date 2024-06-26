// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Canonical Ltd.

/*
 *  Metrics package is used to expose the metrics of the NRF service.
 */

package metrics

import (
	"net/http"

	"github.com/omec-project/nrf/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitMetrics initialises NRF metrics
func InitMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.InitLog.Errorf("Could not open metrics port: %v", err)
	}
}
