package wechatdomaincheck

import "testing"

func TestCheck(t *testing.T) {
	check, err := Check("https://czyt.tech")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(check)
}
