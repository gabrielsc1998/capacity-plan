package main

import (
	internal "github.com/gabrielsc1998/capacity-plan/internal"
)

func main() {
	parameters := internal.Parameters{
		DAU:               internal.M(1),
		ReqPerUser:        2,
		SizePerReq:        10000,
		WriteFactor:       1,
		ReadFactor:        9,
		ReplicationFactor: 4,
	}

	parameters.Calculte()
	print(parameters.ShowResult())
}
