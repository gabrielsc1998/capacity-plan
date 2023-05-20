package internal

import "fmt"

const TOTAL_DAYS_PER_YEAR = 365
const TOTAL_SECONDS_PER_DAY = 100000

type Parameters struct {
	DAU               float64
	ReqPerUser        int
	SizePerReq        float64
	ReadFactor        float64
	WriteFactor       float64
	ReplicationFactor int
	result            Result
}

type Result struct {
	RequestsPerDay    float64
	RequestsPerSecond float64
	Bandwith          float64
	WritePerSecond    float64
	ReadPerSecond     float64
	Storage           float64
	StoragePerDay     float64
	StoragePerYear    float64
}

func (p *Parameters) requestsPerDay() {
	p.result.RequestsPerDay = p.DAU * float64(p.ReqPerUser)
}

func (p *Parameters) requestsPerSecond() {
	p.result.RequestsPerSecond = p.result.RequestsPerDay / TOTAL_SECONDS_PER_DAY
}

func (p *Parameters) bandwith() {
	p.result.Bandwith = p.result.RequestsPerSecond * p.SizePerReq
}

func (p *Parameters) writePerSecond() {
	p.result.WritePerSecond = p.result.RequestsPerSecond / (p.WriteFactor * 10)
}

func (p *Parameters) readPerSecond() {
	p.result.ReadPerSecond = p.result.RequestsPerSecond - p.result.WritePerSecond
}

func (p *Parameters) storage() {
	p.result.Storage = p.result.WritePerSecond * p.SizePerReq * float64(p.ReplicationFactor)
}

func (p *Parameters) storagePerDay() {
	p.result.StoragePerDay = p.result.Storage * TOTAL_SECONDS_PER_DAY
}

func (p *Parameters) storagePerYear() {
	p.result.StoragePerYear = p.result.StoragePerDay * TOTAL_DAYS_PER_YEAR
}

func (p *Parameters) Calculte() {
	p.requestsPerDay()
	p.requestsPerSecond()
	p.bandwith()
	p.writePerSecond()
	p.readPerSecond()
	p.storage()
	p.storagePerDay()
	p.storagePerYear()
}

func (p *Parameters) ShowResult() string {
	result := ""
	result += fmt.Sprint("- RequestsPerDay: ", p.result.RequestsPerDay, " rpd\n")
	result += fmt.Sprint("- RequestsPerSecond: ", p.result.RequestsPerSecond, " rps\n")
	result += fmt.Sprint("- Bandwith: ", p.result.Bandwith/M(1), " MB/s\n")
	result += fmt.Sprint("- WritePerSecond: ", p.result.WritePerSecond, " wps\n")
	result += fmt.Sprint("- ReadPerSecond: ", p.result.ReadPerSecond, " rps\n")
	result += fmt.Sprint("- Storage: ", p.result.Storage/M(1), " MB\n")
	result += fmt.Sprint("- StoragePerDay: ", p.result.StoragePerDay/G(1), " GB\n")
	result += fmt.Sprint("- StoragePerYear: ", p.result.StoragePerYear/T(1), " TB\n")
	return result
}
