package sendgrid

import "net/http"

type ASMGroupService struct {
	sg *SGHTTPClient
}

type ASMGroup struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	LastEmailSentAt int    `json:"last_email_sent_at"`
	IsDefault       bool   `json:"is_default"`
	Unsubscribes    int    `json:"unsubscribes"`
}

func (as *ASMGroupService) List() (al []ASMGroup, err error) {
	// GET: /asm/groups
	var req *http.Request
	req, err = as.sg.NewRequest("GET", "/asm/groups", nil)
	if err != nil {
		return
	}

	err = as.sg.Do(req, &al)
	return al, err
}

func (as *ASMGroupService) Get(groupID string) (a ASMGroup, err error) {
	// GET: /asm/groups/:group_id
	var req *http.Request
	req, err = as.sg.NewRequest("GET", "/asm/groups/"+groupID, nil)
	if err != nil {
		return
	}

	err = as.sg.Do(req, &a)

	return a, err
}

func (as *ASMGroupService) Delete(groupID string) error {
	// DELETE: /asm/groups/:group_id
	var req *http.Request
	req, err := as.sg.NewRequest("DELETE", "/asm/groups/"+groupID, nil)
	if err != nil {
		return err
	}

	return as.sg.Do(req, nil)
}
