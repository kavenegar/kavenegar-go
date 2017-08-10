package kavenegar

//AccountService ...
type AccountService struct {
	client *Client
}

//NewAccountService ...
func NewAccountService(client *Client) *AccountService {
	m := &AccountService{client: client}
	return m
}

type AccountAPILogType string

const (
	Type_AccountAPILog_Justfaults AccountAPILogType = "justfaults"
	Type_AccountAPILog_Enabled    AccountAPILogType = "enabled"
	Type_AccountAPILog_Disabled   AccountAPILogType = "disabled"
)

var AccountAPILogMap = map[AccountAPILogType]string{
	Type_AccountAPILog_Justfaults: "justfaults",
	Type_AccountAPILog_Enabled:    "enabled",
	Type_AccountAPILog_Disabled:   "disabled",
}

func (t AccountAPILogType) String() string {
	return AccountAPILogMap[t]
}

//AccountDailyReport ...
type AccountDailyReportType string

const (
	Type_AccountDailyReport_Justfaults AccountDailyReportType = "justfaults"
	Type_AccountDailyReport_Enabled    AccountDailyReportType = "enabled"
	Type_AccountDailyReport_Disabled   AccountDailyReportType = "disabled"
)

var AccountDailyReportMap = map[AccountDailyReportType]string{
	Type_AccountDailyReport_Justfaults: "justfaults",
	Type_AccountDailyReport_Enabled:    "enabled",
	Type_AccountDailyReport_Disabled:   "disabled",
}

func (t AccountDailyReportType) String() string {
	return AccountDailyReportMap[t]
}

//AccountDebugMode ...
type AccountDebugModeType string

const (
	Type_AccountDebugMode_Enabled  AccountDebugModeType = "enabled"
	Type_AccountDebugMode_Disabled AccountDebugModeType = "disabled"
)

var AccountDebugModeMap = map[AccountDebugModeType]string{
	Type_AccountDebugMode_Enabled:  "enabled",
	Type_AccountDebugMode_Disabled: "disabled",
}

func (t AccountDebugModeType) String() string {
	return AccountDebugModeMap[t]
}

//AccountResendFailed ...
type AccountResendFailedType string

const (
	Type_AccountResendFailed_Enabled  AccountResendFailedType = "enabled"
	Type_AccountResendFailed_Disabled AccountResendFailedType = "disabled"
)

var AccountResendFailedMap = map[AccountResendFailedType]string{
	Type_AccountResendFailed_Enabled:  "enabled",
	Type_AccountResendFailed_Disabled: "disabled",
}

func (t AccountResendFailedType) String() string {
	return AccountResendFailedMap[t]
}
