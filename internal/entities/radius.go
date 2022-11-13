package entities

import (
	"time"
)

type Radacct struct {
	RadAcctId           int       `gorm:"column:radacctid;primaryKey;autoIncrement:true" json:"id"`
	AcctSessionId       string    `gorm:"column:acctsessionid;not null"`
	AcctUniqueId        string    `gorm:"column:acctuniqueid;not null;unique;index:radacct_active_session_idx,where:AcctStopTime is null"`
	UserName            string    `gorm:"column:username;index:radacct_start_user_idx" json:"username"`
	Realm               string    `gorm:"column:realm"`
	NASIPAddress        string    `gorm:"column:nasipaddress;index:radacct_bulk_close,where:AcctStopTime is null" json:"nasipaddress"`
	NASPortId           string    `gorm:"column:nasportid;" json:"nasportid"`
	NASPortType         string    `gorm:"column:nasporttype;"`
	AcctStartTime       time.Time `gorm:"column:acctstarttime;index:radacct_start_user_idx;index:radacct_bulk_close,where:AcctStopTime is null" json:"acctstarttime"`
	AcctUpdateTime      time.Time `gorm:"column:acctupdatetime;" json:"acctupdatetime"`
	AcctStopTime        time.Time `gorm:"column:acctstoptime" json:"acctstoptime"`
	AcctInterval        int64     `gorm:"column:acctinterval;type:bigint"`
	AcctSessionTime     int64     `gorm:"column:acctsessiontime;type:bigint"`
	AcctAuthentic       string    `gorm:"column:acctauthentic;"`
	ConnectInfo_start   string    `gorm:"column:connectinfo_start;"`
	ConnectInfo_stop    string    `gorm:"column:connectinfo_stop;"`
	AcctInputOctets     int64     `gorm:"column:acctinputoctets;type:bigint"`
	AcctOutputOctets    int64     `gorm:"column:acctoutputoctets;type:bigint"`
	CalledStationId     string    `gorm:"column:calledstationid;" json:"calledstationid"`
	CallingStationId    string    `gorm:"column:callingstationid;" json:"callingstationid"`
	AcctTerminateCause  string    `gorm:"column:acctterminatecause;"`
	ServiceType         string    `gorm:"column:servicetype;"`
	FramedProtocol      string    `gorm:"column:framedprotocol;"`
	FramedIPAddress     string    `gorm:"column:framedipaddress;"`
	FramedIPv6Address   string    `gorm:"column:framedipv6address;"`
	FramedIPv6Prefix    string    `gorm:"column:framedipv6prefix;"`
	FramedInterfaceId   string    `gorm:"column:framedinterfaceid;"`
	DelegatedIPv6Prefix string    `gorm:"column:delegatedipv6prefix;"`
	Class               string    `gorm:"column:class;index:radacct_calss_idx"`
}

func (Radacct) TableName() string {
	return "radacct"
}

type Radcheck struct {
	Id        uint32 `gorm:"column:id;primaryKey" json:"id"`
	UserName  string `gorm:"column:username;not null;default:'';index:radcheck_UserName" json:"username"`
	Attribute string `gorm:"not null;default:'';index:radcheck_UserName" json:"attribute"`
	Op        string `gorm:"column:op;not null;default:'==';size:2" json:"op"`
	Value     string `gorm:"not null;default:''" json:"value"`
}

func (Radcheck) TableName() string {
	return "radcheck"
}

type Radgroupcheck struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	GroupName string `gorm:"column:groupname;not null;default:'';index:radgroupcheck_GroupName"`
	Attribute string `gorm:"not null;default:'';index:radgroupcheck_GroupName"`
	Op        string `gorm:"column:op;not null;default:'==';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radgroupcheck) TableName() string {
	return "radgroupcheck"
}

type Radgroupreply struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	GroupName string `gorm:"column:groupname;not null;default:'';index:radgroupreply_GroupName"`
	Attribute string `gorm:"not null;default:'';index:radgroupreply_GroupName"`
	Op        string `gorm:"column:op;not null;default:'=';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radgroupreply) TableName() string {
	return "radgroupreply"
}

type Radreply struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	UserName  string `gorm:"column:username;not null;default:'';index:radreply_UserName"`
	Attribute string `gorm:"not null;default:'';index:radreply_UserName"`
	Op        string `gorm:"column:op;not null;default:'=';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radreply) TableName() string {
	return "radreply"
}

type Radusergroup struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	UserName  string `gorm:"column:username;not null;default:'';index:radusergroup_UserName"`
	GroupName string `gorm:"column:groupname;not null;default:''"`
	Priority  int    `gorm:"column:priority;not null;default:0"`
}

func (Radusergroup) TableName() string {
	return "radusergroup"
}

type Radpostauth struct {
	Id               int64     `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	UserName         string    `gorm:"column:username;not null;index:radpostauth_username_idx"`
	Pass             string    `gorm:"column:pass"`
	Reply            string    `gorm:"column:reply"`
	CalledStationId  string    `gorm:"column:calledstationid"`
	CallingStationId string    `gorm:"column:callingstationid"`
	Authdate         time.Time `gorm:"column:authdate;not null;default:now()"`
	Class            string    `gorm:"column:class;index:radpostauth_class_idx"`
}

func (Radpostauth) TableName() string {
	return "radpostauth"
}

type Nas struct {
	Id          int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	Nasname     string `gorm:"column:nasname;not null;index:nas_nasname"`
	Shortname   string `gorm:"column:shortname;not null"`
	Type        string `gorm:"column:type;not null;default:'other'"`
	Ports       int    `gorm:"column:ports"`
	Secret      string `gorm:"column:secret;not null"`
	Server      string `gorm:"column:server"`
	Community   string `gorm:"column:community"`
	Description string `gorm:"column:description"`
}

func (Nas) TableName() string {
	return "nas"
}

type Radverified struct {
	Id               int       `gorm:"column:id;primaryKey;unique;index;autoIncrement:true" json:"id"`
	UserName         string    `gorm:"column:username;not null;index:radknown_UserName" json:"username"`
	CallingStationId string    `gorm:"column:callingstationid;unique;index" json:"callingstationid"`
	AcctUpdateTime   time.Time `gorm:"column:acctupdatetime" json:"acctupdatetime"`
}

func (Radverified) TableName() string {
	return "radverified"
}

type Acct struct {
	RadAcctId        int       `gorm:"column:radacctid" json:"id"`
	UserName         string    `gorm:"column:username" json:"username"`
	NASIPAddress     string    `gorm:"column:nasipaddress" json:"nasipaddress"`
	NASPortId        string    `gorm:"column:nasportid" json:"nasportid"`
	AcctStartTime    time.Time `gorm:"column:acctstarttime" json:"acctstarttime"`
	AcctUpdateTime   time.Time `gorm:"column:acctupdatetime" json:"acctupdatetime"`
	AcctStopTime     time.Time `gorm:"column:acctstoptime" json:"acctstoptime"`
	CalledStationId  string    `gorm:"column:calledstationid" json:"calledstationid"`
	CallingStationId string    `gorm:"column:callingstationid" json:"callingstationid"`
	Valid            string    `gorm:"column:valid" json:"valid"`
	Verified         string    `gorm:"column:verified" json:"verified"`
}

func (Acct) TableName() string {
	return "accounting"
}
