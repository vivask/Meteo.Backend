package entities

import "time"

type Radacct struct {
	RadAcctId           int       `gorm:"column:radacctid;primaryKey;unique;index;autoIncrement:true"`
	AcctSessionId       string    `gorm:"not null"`
	AcctUniqueId        string    `gorm:"not null;unique;index:radacct_active_session_idx,where:AcctStopTime is null"`
	UserName            string    `gorm:"index:radacct_start_user_idx"`
	Realm               string    `gorm:""`
	NASIPAddress        string    `gorm:"index:radacct_bulk_close,where:AcctStopTime is null"`
	NASPortId           string    `gorm:""`
	NASPortType         string    `gorm:""`
	AcctStartTime       time.Time `gorm:"index:radacct_start_user_idx;index:radacct_bulk_close,where:AcctStopTime is null"`
	AcctUpdateTime      time.Time `gorm:""`
	AcctStopTime        time.Time `gorm:"column:acctstoptime"`
	AcctInterval        int64     `gorm:"type:bigint"`
	AcctSessionTime     int64     `gorm:"type:bigint"`
	AcctAuthentic       string    `gorm:""`
	ConnectInfo_start   string    `gorm:""`
	ConnectInfo_stop    string    `gorm:""`
	AcctInputOctets     int64     `gorm:"type:bigint"`
	AcctOutputOctets    int64     `gorm:"type:bigint"`
	CalledStationId     string    `gorm:""`
	CallingStationId    string    `gorm:""`
	AcctTerminateCause  string    `gorm:""`
	ServiceType         string    `gorm:""`
	FramedProtocol      string    `gorm:""`
	FramedIPAddress     string    `gorm:""`
	FramedIPv6Address   string    `gorm:""`
	FramedIPv6Prefix    string    `gorm:""`
	FramedInterfaceId   string    `gorm:""`
	DelegatedIPv6Prefix string    `gorm:""`
	Class               string    `gorm:"index:radacct_calss_idx"`
}

func (Radacct) TableName() string {
	return "radacct"
}

type Radcheck struct {
	Id        int    `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserName  string `gorm:"not null;default:'';index:radcheck_UserName"`
	Attribute string `gorm:"not null;default:'';index:radcheck_UserName"`
	Op        string `gorm:"column:op;not null;default:'==';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radcheck) TableName() string {
	return "radcheck"
}

type Radgroupcheck struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	GroupName string `gorm:"not null;default:'';index:radgroupcheck_GroupName"`
	Attribute string `gorm:"not null;default:'';index:radgroupcheck_GroupName"`
	Op        string `gorm:"column:op;not null;default:'==';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radgroupcheck) TableName() string {
	return "radgroupcheck"
}

type Radgroupreply struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	GroupName string `gorm:"not null;default:'';index:radgroupreply_GroupName"`
	Attribute string `gorm:"not null;default:'';index:radgroupreply_GroupName"`
	Op        string `gorm:"column:op;not null;default:'=';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radgroupreply) TableName() string {
	return "radgroupreply"
}

type Radreply struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	UserName  string `gorm:"not null;default:'';index:radreply_UserName"`
	Attribute string `gorm:"not null;default:'';index:radreply_UserName"`
	Op        string `gorm:"column:op;not null;default:'=';size:2"`
	Value     string `gorm:"not null;default:''"`
}

func (Radreply) TableName() string {
	return "radreply"
}

type Radusergroup struct {
	Id        int    `gorm:"column:id;primaryKey;unique;index;autoIncrement:true"`
	UserName  string `gorm:"not null;default:'';index:radusergroup_UserName"`
	GroupName string `gorm:"not null;default:''"`
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
	CalledStationId  string    `gorm:""`
	CallingStationId string    `gorm:""`
	Authdate         time.Time `gorm:"column:authdate;not null;default:now()"`
	Class            string    `gorm:"index:radpostauth_class_idx"`
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
