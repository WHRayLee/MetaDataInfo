package model

type ConditionInfo struct {
	leftColumnInfo  string
	sign            string
	rightColumnInfo string
}

func (c *ConditionInfo) SetLeftColumnInfo(leftColumnInfo string) {
	c.leftColumnInfo = leftColumnInfo
}

func (c *ConditionInfo) SetSign(sign string) {
	c.sign = sign
}

func (c *ConditionInfo) SetRightColumnInfo(rightColumnInfo string) {
	c.rightColumnInfo = rightColumnInfo
}

func (c ConditionInfo) LeftColumnInfo() string {
	return c.leftColumnInfo
}

func (c ConditionInfo) Sign() string {
	return c.sign
}

func (c ConditionInfo) RightColumnInfo() string {
	return c.rightColumnInfo
}

func NewConditionInfo() *ConditionInfo {
	return &ConditionInfo{}
}
