package common

import (
	"testing"
	"reflect"
	"time"
	"errors"
	"log"
)

func TestKeySet(t *testing.T) {
	mp := map[string]interface{}{
		"one": struct{}{},
		"two": struct{}{},
	}

	actual := KeySet(mp)
	expected := []string{"one", "two"}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Arrays are not equal. Expected:%s, Actual:%s", expected, actual)
	}
}

func TestRetryAttempts(t *testing.T) {
	i := 0
	Retry(2, 1*time.Second, func() error {
		i++
		return errors.New("some error")
	})

	log.Println(i)
	if 2 != i {
		t.Errorf("Incorrect attempts count: %d", i)
	}
}