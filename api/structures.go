package api

type UserToken struct {
	Token string `json:"token"`
}

type SyncStatus struct {
	TempIdMapping        map[string]int    `json:"TempIdMapping"`
	GlobalSequenceNumber int               `json:"seq_no_global"`
	SequenceNumber       int               `json:"seq_no"`
	UserId               int               `json:"UserId"`
	SyncStatus           map[string]string `json:"SyncStatus"`
}

type TodoistError struct {
	Error string `json:"error"`
	Tag   string `json:"error_tag"`
	Code  int    `json:"error_code"`
	Extra string `json:"error_extra"`
}

type Data struct {
	User     *User     `json:"user"`
	UserId   *string   `json:"user_id"`
	Projects []Project `json:"Projects"`
	Tasks    []Task    `json:"Items"`
	Labels   []Label   `json:"Labels"`
}

type User struct {
	Id                         string   `json:"id"`
	APIToken                   string   `json:"api_token"`
	Email                      string   `json:"email"`
	Fullname                   string   `json:"full_name"`
	StartPage                  string   `json:"start_page"`
	Timezone                   string   `json:"timezone"`
	TimezoneOffset             []string `json:"tz_offset"`
	TimeFormat                 int      `json:"time_format"`
	DateFormat                 int      `json:"date_format"`
	SortOrder                  int      `json:"sort_order"`
	MobileNumber               string   `json:"mobile_number"`
	MobileHost                 string   `json:"mobile_host"`
	SubscriptionExpirationDate string   `json:"premium_until"`
}

type Filter struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Color     int    `json:"color"`
	Order     int    `json:"item_order"`
	Query     string `json:"query"`
	IsDeleted bool   `json:"is_deleted"`
	Id        string `json:"id"`
}

type Label struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Color int    `json:"color,omitempty"`
	Order int    `json:"item_order,omitempty"`
}

type Project struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Color     int    `json:"color"`
	Indent    int    `json:"indent"`
	Order     int    `json:"item_order"`
	Collapsed int    `json:"collapsed"`
}

type Task struct {
	Id             int    `json:"id,omitempty"`
	UserId         int    `json:"user_id,omitempty"`
	Content        string `json:"content"`
	ProjectId      int    `json:"project_id"`
	Date           string `json:"date_string,omitempty"`
	DateLang       string `json:"date_lang,omitempty"`
	DueDateUTC     string `json:"due_date_utc,omitempty"`
	InHistory      int    `json:"in_history,omitempty"`
	Collapsed      int    `json:"collapsed,omitempty"`
	Priority       int    `json:"priority,omitempty"`
	Indent         int    `json:"indent,omitempty"`
	Order          int    `json:"item_order,omitempty"`
	Children       []int  `json:"children,omitempty"`
	Labels         []int  `json:"labels,omitempty"`
	Assigned       int    `json:"assigned_by_uid,omitempty"`
	ResponsibleUID int    `json:"responsible_uid,omitempty"`
}

type Reminder struct {
	Id           string `json:"id"`
	ItemId       string `json:"item_id"`
	Service      string `json:"service"`
	Type         string `json:"type"`
	DueDateUTC   string `json:"due_date_utc"`
	Date         string `json:"date_string"`
	DateLang     string `json:"date_lang"`
	UserToNotify string `json:"notify_uid"`
}

type Note struct {
	Id        string `json:"id"`
	ItemId    string `json:"item_id"`
	Content   string `json:"content"`
	ProjectId string `json:"project_id"`
	File      File   `json:"file_attachment"`
}

type File struct {
	Filename    string `json:"file_name"`
	Size        int    `json:"file_size"`
	Type        string `json:"file_type"`
	URL         string `json:"file_url"`
	UploadState string `json:"upload_state"`
}

type ImageFile struct {
	File
	LargeThumbnail  []string `json:"tn_l"`
	MediumThumbnail []string `json:"tn_m"`
	SmallThumbnail  []string `json:"tn_s"`
}

type AudioFile struct {
	File
	Duration int `json:"file_duration"` // duration in seconds
}

type Projects []Project

func (projects Projects) Len() int {
	return len(projects)
}
func (projects Projects) Swap(i, j int) {
	projects[i], projects[j] = projects[j], projects[i]
}
func (projects Projects) Less(i, j int) bool {
	return projects[i].Order < projects[j].Order
}

type Tasks []Task

func (tasks Tasks) Len() int {
	return len(tasks)
}
func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}
func (tasks Tasks) Less(i, j int) bool {
	return tasks[i].ProjectId < tasks[j].ProjectId
}
