// @Author Ben.Zheng
// @DateTime 8/20/22 10:15 PM

package po

type IssueCore struct {
	BaseMetaNumericId
	BaseMetaCreatedAt
	SubjectId string `gorm:"column:subj_id;type:varchar(21);uniqueIndex:idx_issue_id;<-;"`
	Title     string `gorm:"column:title;type:nvarchar(32);<-;"`
}

type Issue[T IssueCore] struct {
	Core T                               `gorm:"embedded;"`
	Map  *SubjectIdMap[SubjectIdMapCore] `gorm:"foreignKey:SubjectId;"`
}

func (i Issue[T]) TableName() string {
	return "idk_issues"
}

func (i Issue[T]) GetCore() T {
	return i.Core
}

// IssueContentCore
// @field Question Describes what a bug/question is in short.
// @field Description Describes how to reproduce the bug/question step by step.
// @field Debug Describes the analysis process.
// @field Solution Describes a doable plan or an idea of how to solve the bug/question.
// @field Proof Describes a post verification for solution, describe the effect after bug fixing.
type IssueContentCore struct {
	AutoIncIdFullMode
	BaseVersionInfo
	IssueId             int    `gorm:"column:issue_id;type:int;index:idx_ref_issue_id;<-;"`
	Question            string `gorm:"column:question;type:text;<-;"`
	Description         string `gorm:"column:desc;type:text;<-;"`
	DescriptionDisabled bool   `gorm:"column:desc_disabled;type:boolean;<-;"`
	Debug               string `gorm:"column:debug;type:text;<-;"`
	DebugDisable        bool   `gorm:"column:debug_disabled;type:boolean;<-;"`
	Solution            string `gorm:"column:solution;type:text;<-;"`
	Proof               string `gorm:"column:proof;type:text;<-;"`
	ProofDisabled       bool   `gorm:"column:proof_disabled;type:boolean;<-;"`
}

type IssueContent[T IssueContentCore] struct {
	Core  T                 `gorm:"embedded;"`
	Issue *Issue[IssueCore] `gorm:"foreignKey:IssueId;"`
}

func (i IssueContent[T]) TableName() string {
	return "idk_issue_contents"
}

func (i IssueContent[T]) GetCore() T {
	return i.Core
}
