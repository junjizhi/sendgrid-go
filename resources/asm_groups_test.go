package sendgrid

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestASMGroupGet(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("asm_group.get.json", t)
	defer data.Close()

	const groupID = "100"
	mux.HandleFunc("/asm/groups/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		checkURLSuffix(t, r, groupID)
		io.Copy(w, data)
	})

	a, err := client.ASMGroup.Get(groupID)
	if err != nil {
		t.Fatal(err)
	}
	if strconv.Itoa(a.ID) != groupID {
		t.Fatalf("Group ID = %v, want %v", a.ID, groupID)
	}

	testBadURL(t, func() error {
		_, err := client.ASMGroup.Get(groupID)
		return err
	})
}

func TestASMGroupList(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("asm_groups.list.json", t)
	defer data.Close()

	mux.HandleFunc("/asm/groups", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		io.Copy(w, data)
	})

	b, err := client.ASMGroup.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) < 2 {
		t.Fatalf("Should include 2 or more asm groups")
	}

	testBadURL(t, func() error {
		_, err := client.ASMGroup.List()
		return err
	})
}

func TestASMGroupDelete(t *testing.T) {
	setup()
	defer teardown()

	const groupID = "1"
	const noID = "9999999"
	mux.HandleFunc("/asm/groups/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "DELETE")
		split := strings.Split(r.URL.Path, "/")
		if split[3] == noID {
			http.Error(w, "invalid id", http.StatusNotFound)
		}
	})

	if err := client.ASMGroup.Delete(groupID); err != nil {
		t.Fatal(err)
	}

	if err := client.ASMGroup.Delete(noID); err == nil {
		t.Fatal("expected HTTP 404")
	}

	testBadURL(t, func() error {
		return client.ASMGroup.Delete(groupID)
	})
}
