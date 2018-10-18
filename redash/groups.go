package redash

import (
	"io/ioutil"
	"strconv"
)

type GetGroupInput struct {
	GroupId int
}

type GetGroupOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetGroup(input *GetGroupInput) *GetGroupOutput {
	path := "/api/groups/" + strconv.Itoa(input.GroupId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetGroupOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetGroupOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type GetGroupListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetGroupList() *GetGroupListOutput {
	path := "/api/groups"

	resp, err := c.Get(path)
	if err != nil {
		return &GetGroupListOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetGroupListOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
