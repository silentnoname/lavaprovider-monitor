package main

import (
	"go.uber.org/zap"
	"lavaprovider-monitor/pkg/alert"
	"lavaprovider-monitor/pkg/check"
	"lavaprovider-monitor/pkg/config"
	"lavaprovider-monitor/pkg/log"
	"time"
)

func main() {
	log.InitLog()
	log.Log.Info("Start to monitor")
	log.Log.Info("Loading config")
	aconfig := config.GetAlertConfig()
	monitorconfig := config.GetLavaProviderMonitorConfig()
	log.Log.Info("Using grpc", zap.String("grpc", monitorconfig.LavaGrpc))
	log.Log.Info("Lava chainid", zap.String("chainid", monitorconfig.ChainID))
	log.Log.Info("Monitoring Lava provider address", zap.String("lavaprovideraddress", monitorconfig.LavaProviderAddress))
	log.Log.Info("Monitoring chains", zap.Strings("chains", monitorconfig.Chains))
	alert.CheckAlertConfig(aconfig)
	log.Log.Info("Will check and send alert every 10 minutes")
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		log.Log.Info("Start to check")
		// check
		alerts := check.CheckSpecificChainProvider(monitorconfig.LavaProviderAddress, monitorconfig.LavaGrpc, monitorconfig.ChainID, monitorconfig.Chains)
		log.Log.Info("Start to send alert")
		alerterr := alert.SendAlert(aconfig, alerts)
		if alerterr != nil {
			log.Log.Error("Failed to send alert", zap.Error(alerterr))
		}
	}
}
