package config

import (
	"flag"
	"fmt"
	"time"
)

type Config struct {
	NightTariff    float64
	DayTariff      float64
	IdleTariffPerH float64
	IdleSpeedKM    float64
	SpeedLimit     float64
	PickupFee      float64
	MinFare        float64
	InputFile      string
	OutputFile     string
}

func New() *Config {
	conf := &Config{}

	flag.Float64Var(&conf.NightTariff, "nighttariff", 1.3, "night ride charging tariff per km")
	flag.Float64Var(&conf.DayTariff, "dayttariff", 0.74, "day ride charging tariff per km")
	flag.Float64Var(&conf.IdleTariffPerH, "idletariffperh", 11.9, "IDLE charging tariff per hour")
	flag.Float64Var(&conf.IdleSpeedKM, "idlespeedkm", 10, "speed to consider as IDLE")
	flag.Float64Var(&conf.SpeedLimit, "speedlimit", 100, "maximum valid speed limit")
	flag.Float64Var(&conf.PickupFee, "pickupfee", 1.3, "pickup fee")
	flag.Float64Var(&conf.MinFare, "minfare", 3.47, "minimum fare")
	flag.StringVar(&conf.InputFile, "inputfile", "testdata/paths.csv", "file to be processed")
	flag.StringVar(&conf.OutputFile, "outputfile", fmt.Sprintf("%v_result.csv", time.Now().UTC().Unix()), "file to be produced")

	flag.Parse()

	return conf
}
