package model

type JoinInfo struct {
	joinType          string
	joinTableInfo     TableBasicInfo
	joinConditionList []ConditionInfo
}

func (j *JoinInfo) JoinType() string {
	return j.joinType
}

func (j *JoinInfo) SetJoinType(joinType string) {
	j.joinType = joinType
}

func (j *JoinInfo) SetJoinConditionList(joinConditionList []ConditionInfo) {
	j.joinConditionList = joinConditionList
}

func (j *JoinInfo) SetJoinTableInfo(joinTableInfo TableBasicInfo) {
	j.joinTableInfo = joinTableInfo
}

func NewJoinInfo() *JoinInfo {
	return &JoinInfo{}
}
