package sneaker

import (
	"bytes"
	"testing"
)

func TestGCMRoundTrip(t *testing.T) {
	key := make([]byte, 32)
	input := []byte("hello this is Stripe")

	enc, err := encrypt(key, input, []byte("yay"))
	if err != nil {
		t.Fatal(err)
	}

	dec, err := decrypt(key, enc, []byte("yay"))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(input, dec) {
		t.Errorf("Was %x, but expected %x", dec, input)
	}
}

func TestGCMRoundTripWithModifications(t *testing.T) {
	key := make([]byte, 32)
	input := []byte("hello this is Stripe")

	enc, err := encrypt(key, input, []byte("yay"))
	if err != nil {
		t.Fatal(err)
	}

	enc[5] ^= 1 // flip a bit

	dec, err := decrypt(key, enc, []byte("yay"))
	if err == nil {
		t.Fatalf("Was %x, but expected an error", dec)
	}
}

func TestGCMRoundTripWithBadData(t *testing.T) {
	key := make([]byte, 32)
	input := []byte("hello this is Stripe")

	enc, err := encrypt(key, input, []byte("yay"))
	if err != nil {
		t.Fatal(err)
	}

	dec, err := decrypt(key, enc, []byte("boo"))
	if err == nil {
		t.Fatalf("Was %x, but expected an error", dec)
	}
}
