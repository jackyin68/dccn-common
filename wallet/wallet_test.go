package wallet

import (
	//"fmt"
	"testing"
)

const to_sign_text = "123456"
const priv_key = "mB3bPsj9kUsAPmFw11Gqb6AYi7nQ8PFrNI+G62IRyYnstkObfl1KIeQi8pOfpNovC2iikxivhW9baCLStM2hyA=="
const expected_result = "78u5RkbyGvb8+0GvomXA2CUSimMLT83YEurkqv3eJiXSYh2wPTW24eksWDiBFbzHCyUxZkL6CuHCgFRXPPZABw=="

func TestGenerateKeys(t *testing.T) {
	a, b, c := GenerateKeys()
	if len(a) != 88 {
		t.Fail()
	}
	if len(b) != 44 {
		t.Fail()
	}
	if len(c) != 40 {
		t.Fail()
	}

}

func TestSha256Sign(t *testing.T) {
	sig, err := Sha256Sign(to_sign_text, priv_key)
        if err == nil {
		if sig != expected_result {
			t.Fail()
		}
        } else {
		t.Fail()
        }

}
