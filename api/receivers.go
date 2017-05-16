package main

type Scan struct {
	IcrID        string `json:"icrid"`
	JobID        string `json:"jobid"`
	CasketID     string `json:"casketid"`
	Width        string `json:"breadth"`
	Length       string `json:"length"`
	Height       string `json:"height"`
	BoxVol       string `json:"box_volume"`
	RealVol      string `json:"real_volume"`
	VmsStatus    string `json:"vms_status_code"`
	WeightStatus string `json:"weight_status_code"`
	Weight       string `json:"weight"`
	InputNo      string `json:"inductid"`
	ImageID      string `json:"imageid"`
	Uuid         string `json:"uid"`
	Barcode      string `json:"barcode"`
	ScanStatus   string `json:"scan_status_code"`
}

type ScanRcv struct {
	ChuteId string `json:"chuteid"`
}

type Event struct {
	Code string `json:"event_code"`
}

//Receive sort data
type Sort struct {
	JobId      string `json:"jobid"`
	Uuid       string `json:"uid"`
	SortStatus string `json:"sort_status_code"`
	ChuteId    string `json:"chuteid"`
}

type SorterSort struct {
	Sort
	SorterID string `json:"sorterid"`
}

type EcdsSort struct {
	Sort
	SorterID string `json:"ecdsid"`
}

type IcrSort struct {
	Sort
	SorterID string `json:"icrid"`
}

/*
//Hold chute-id specific data to send to main code
struct ChuteId{
	jobId [4]string
	casketId [2]string
	chuteId [2]string
}

struct InitConfig{
	subPacketType string
	elemCount [2]string
	elements [INIT_CONFIG_ELEMENTS_SIZE]string
}

//Common structure used by both ApiSnd and ApiRcv
struct Common{
	header [3]string
	reserved string
	version string
	rcv [2]string
	snd [2]string
	length [2]string
	timestamp [4]string
	Type [2]string
}

//Receive and hold any data sent for api
struct ApiRcv{
	Common common
	union{
		BarVms barVms
		Event event
		Sort sort
	}
}

//Hold and send data from api to server
struct ApiSnd{
	Common common
	union{
		ChuteId chuteId
		InitConfig initConfig
	}
}

*/
