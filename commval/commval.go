package commval

const (
	DESC_HEALTHCHECK_OK			        string = "health check ok"
	DESC_UNSUPPORTED_ACTION			    string = "unsupported action"
	DESC_UNSUPPORTED_PATH	 		    	string = "unsupported url path"
	DESC_PARSE_FAILED  		        	string = "request parse failed"
	DESC_READ_BODY_FAILED				    string = "request body read failed"
	DESC_INVALID_CID				    	string = "invalid company id"
	DESC_INVALID_REGION			    	string = "invalid region"
	DESC_INVALID_RID					    string = "invalid request id"
	DESC_INVALID_MULTI					    string = "invalid multi"
	DESC_INVALID_DATA_TYPE					string = "invalid data type"
	DESC_INVALID_UID					    string = "invalid user id"
	DESC_INVALID_GID					    string = "invalid group id"
	DESC_INVALID_GNAME					    string = "invalid group name"
	DESC_INVALID_GTYPE					    string = "invalid group type"
	DESC_INVALID_DID					    string = "invalid department id"
	DESC_INVALID_UTYPE			            string = "invalid user type"
	DESC_EXTRA_PARAM			            string = "too many params"
	DESC_CONFIG_LOAD_FAILED			    string = "load config file failed"
	DESC_CONFIG_MAPPING_FAILED			    string = "mapping config content failed"
	DESC_INVALID_HEADER                    string = "invalid header"
	DESC_ADD_CID2BODY_FAILED				string = "add cid to body failed"
	DESC_HEADER_MISSING_AUTH               string = "invalid header, missing Authorization"
	DESC_HEADER_DECODE_FAILED              string = "invalid header, decode jwt token failed"
	DESC_HEADER_AUTH_FAILED                string = "invalid header, auth jwt token failed"
	DESC_HEADER_CUSTOMER_NIL               string = "invalid header, missing x-customer-id"
	DESC_CACHE_EXCEED_LENGTH	                    string = "policy exceed size limit"
	DESC_CACHE_FRAG_DISABLED	                    string = "frag mechanism disabled"
)

const (
	AdminUser string = "admin"
	RootUser string = "root_user"
	SuperUser string= "super_user"
	MaintainUserPrefix string = "muser@"
	FactoryUserPrefix string = "fuser@"
	DefaultPasswd string = "123456"
	FactoryUserNickName string = "???????????????"
	MaintainUserNickName string = "???????????????"
	FactoryUserAuthorityId string = "8883"
	MaintainUserAuthorityId string = "8884"

	AreaRootParentId int = -1

	ItemPeriodDay string = "?????????"
	ItemPeriodWeek string = "?????????"
	ItemPeriodMonth string = "?????????"
	ItemPeriodQuarter string = "????????????"
	ItemPeriodSemester string = "????????????"
)

const (
	SighInRedirectSuccess string = "http://smartsafety.njzhida.cn:8080/registersuccess.html"
	SighInRedirectOver string = "http://smartsafety.njzhida.cn:8080/registerover.html"
	SighInRedirectFailed string = "http://smartsafety.njzhida.cn:8080/registerfailed.html"
)

const (
	UserTypeAdmin = iota
	UserTypeFactoryUser
	UserTypeMaintainUser
)

const (
	AppUserRoleErr = iota
	AppUserRoleInspector
	AppUserRoleMaintain
	AppUserRoleFactory
)

const (
	TaskStatusNotStart = iota
	TaskStatusReportIssue
	TaskStatusAssignTask
	TaskStatusApproval
	TaskStatusEnd
	TaskStatusTimeOut
)

const (
	CronTaskTime int = 2 //?????????????????????????????????,??????2???
	TimeOutTaskCronTime int = 23
)

var TaskStatus map[int]string

func InitCommVal() {
	TaskStatus = make(map[int]string)
	TaskStatus[TaskStatusNotStart] = "?????????"
	TaskStatus[TaskStatusReportIssue] = "?????????????????????"
	TaskStatus[TaskStatusAssignTask] = "?????????????????????"
	TaskStatus[TaskStatusApproval] = "?????????????????????"
	TaskStatus[TaskStatusEnd] = "??????"
}

