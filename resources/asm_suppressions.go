package sendgrid

import "net/http"

type ASMSuppressionService struct {
	sg *SGHTTPClient
}

func (as *ASMSuppressionService) Get(groupID string) (a []string, err error) {
	// GET: /asm/groups/:group_id
	var req *http.Request
	req, err = as.sg.NewRequest("GET", "/asm/groups/"+groupID+"/suppressions", nil)
	if err != nil {
		return
	}

	err = as.sg.Do(req, &a)

	return a, err
}

func (as *ASMSuppressionService) Delete(groupID string, email string) error {
	// DELETE: /asm/groups/:group_id/suppressions/:email_address
	var req *http.Request
	req, err := as.sg.NewRequest("DELETE", "/asm/groups/"+groupID+"/suppressions/"+email, nil)
	if err != nil {
		return err
	}

	return as.sg.Do(req, nil)
}
