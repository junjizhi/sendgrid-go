package sendgrid

import "net/http"

type GlobalSuppressionService struct {
	sg *SGHTTPClient
}

type GlobalSuppression struct {
	RecipientEmail string `json:"recipient_email"`
}

func (gs *GlobalSuppressionService) Get(email string) (g GlobalSuppression, err error) {
	// GET: /asm/suppressions/global/:email_address
	var req *http.Request
	req, err = gs.sg.NewRequest("GET", "/asm/suppressions/global/"+email, nil)
	if err != nil {
		return
	}

	err = gs.sg.Do(req, &g)

	return g, err
}

func (gs *GlobalSuppressionService) Delete(email string) error {
	// DELETE: /asm/suppressions/global/:email_address
	var req *http.Request
	req, err := gs.sg.NewRequest("DELETE", "/asm/suppressions/global/"+email, nil)
	if err != nil {
		return err
	}

	return gs.sg.Do(req, nil)
}
