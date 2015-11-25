package sendgrid

import "net/http"

type BounceService struct {
	sg *SGHTTPClient
}

type Bounce struct {
	Email   string `json:"email"`
	Reason  string `json:"reason"`
	Status  string `json:"status"`
	Created int    `json:"created"`
}

type BounceListRequest struct {
	StartTime string `url:"start_time,omitempty"`
	EndTime   string `url:"end_time,omitempty"`
}

func (bs *BounceService) List(lr *BounceListRequest) (bl []Bounce, err error) {
	// GET: /suppression/bounces
	var req *http.Request
	req, err = bs.sg.NewRequest("GET", "/suppression/bounces", lr)
	if err != nil {
		return
	}

	err = bs.sg.Do(req, &bl)
	return bl, err
}

func (bs *BounceService) Get(email string) (b []Bounce, err error) {
	// GET: /suppression/bounces/:email
	var req *http.Request
	var bl []Bounce
	req, err = bs.sg.NewRequest("GET", "/suppression/bounces/"+email, nil)
	if err != nil {
		return
	}

	err = bs.sg.Do(req, &bl)

	return bl, err
}

func (bs *BounceService) Delete(email string) error {
	// DELETE: /suppression/bounces/:email
	var req *http.Request
	req, err := bs.sg.NewRequest("DELETE", "/suppression/bounces/"+email, nil)
	if err != nil {
		return err
	}

	return bs.sg.Do(req, nil)
}

// func (bs *BounceService) DeleteList(lr *EmailListRequest) error {
// 	// GET: /suppression/bounces
// 	var req *http.Request
// 	req, err := bs.sg.NewRequest("DELETE", "/suppression/bounces", lr)
// 	if err != nil {
// 		return err
// 	}

// 	return bs.sg.Do(req, nil)
// }
