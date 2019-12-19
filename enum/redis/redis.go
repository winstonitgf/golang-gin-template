package redis

type KeyList int

const (
	Configs KeyList = iota
	// LatestVideo
	// HotVideo
	// HotTag
)

func (d KeyList) String() string {
	// return [...]string{"Configs", "LatestVideo", "HotVideo", "HotTag"}[d]
	return [...]string{"Configs"}[d]
}


type WhiteIpList int

const (
	WhiteList WhiteIpList = iota

)

func (d WhiteIpList) StringWhiteList() string {

	return [...]string{"whiteList"}[d]
}