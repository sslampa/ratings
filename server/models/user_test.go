package models

import (
	"log"
	"testing"
)

func init() {
	Initialize("ratings_app_test")
	Seed()
}

func TestAllUsers(t *testing.T) {
	expect := []User{
		{
			ID:       "1",
			Username: "sslampa",
		},
		{
			ID:       "2",
			Username: "tomanistor",
		},
		{
			ID:       "3",
			Username: "suzmas",
		},
	}

	actual, err := GetUsers()
	if err != nil {
		t.Error(err)
	}

	if len(actual) != len(expect) {
		t.Errorf("Expected %v to equal %v", len(actual), len(expect))
	}

	for i, a := range actual {
		userComp(t, expect[i], a)
	}

}

func TestUserID(t *testing.T) {
	expect := User{ID: "1", Username: "sslampa"}
	user, err := GetUser("id", "1")
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)

	_, err = GetUser("id", "3000")
	if err == nil {
		t.Errorf("Expected to find no user")
	}

	_, err = GetUser("something", "1")
	if err == nil {
		t.Errorf("Expected to find no user")
	}
}

func TestUserUsername(t *testing.T) {
	expect := User{ID: "1", Username: "sslampa"}
	user, err := GetUser("username", "sslampa")
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)
}

func TestUserPost(t *testing.T) {
	expect := User{ID: "4", Username: "cmfasulo"}
	user, err := PostUser("cmfasulo")
	if err != nil {
		t.Errorf("Expected query to return a user")
	}

	userComp(t, expect, user)

	_, err = PostUser("cmfasulo")
	if err == nil {
		t.Errorf("Expected username %v to not equal %v", user.Username, expect.Username)
	}
}

func TestUserDelete(t *testing.T) {
	expect, _ := PostUser("mrobock")

	err := DeleteUser("username", expect.Username)
	if err != nil {
		t.Errorf("Expected user to be found in db before deletion")
	}

	_, err = GetUser("username", expect.Username)
	if err == nil {
		t.Errorf("Expected user to not be found in db after deletion")
	}

	expect, _ = PostUser("mrobock")

	err = DeleteUser("id", expect.ID)
	if err != nil {
		t.Errorf("Expected user to be found in db before deletion")
	}

	_, err = GetUser("id", expect.ID)
	if err == nil {
		t.Errorf("Expected user to not be found in db after deletion")
	}

	expect, _ = PostUser("mrobock")

	err = DeleteUser("wrongInput", expect.Username)
	if err == nil {
		t.Errorf("Expected an error to be thrown for incorrect input value")
	}

	expect, _ = PostUser("mrobock")

	err = DeleteUser("username", "wrongInput")
	if err == nil {
		t.Errorf("Expected an error to be thrown for incorrect input value")
	}

	expect, _ = PostUser("mrobock")

	err = DeleteUser("id", "wrongInput")
	if err == nil {
		t.Errorf("Expected an error to be thrown for incorrect input value")
	}
}

func TestUserUpdateUsername(t *testing.T) {
	expect := User{ID: "3", Username: "sucrosm"}

	actual, err := UpdateUser("username", "suzmas", "sucrosm")
	if err != nil {
		t.Errorf("Expected no error")
	}

	userComp(t, expect, actual)
}

func TestUserUpdateID(t *testing.T) {
	expect := User{ID: "2", Username: "tomas"}

	actual, err := UpdateUser("id", "2", "tomas")
	if err != nil {
		t.Errorf("Expected no error")
	}

	userComp(t, expect, actual)
}

func TestUserUpdateIncorrectUsername(t *testing.T) {
	_, err := UpdateUser("username", "incorrectUser", "thisShouldNotBeHere")
	if err == nil {
		t.Errorf("Expected incorrectUser to not update")
	}

	_, err = UpdateUser("id", "0", "thisShouldNotBeHere")
	if err == nil {
		t.Errorf("Expected 0 to not update")
	}
}

func TestUserShowsCollection(t *testing.T) {
	es, err := GetShows("1")
	if err != nil {
		log.Fatal(err)
	}

	u, err := GetUser("id", "1")
	if err != nil {
		log.Fatal(err)
	}

	actual, err := GetUserShows(u.Username)
	if err != nil {
		log.Fatal(err)
	}

	if actual.ID != u.ID {
		t.Errorf("Expected %v to equal %v", actual.ID, u.ID)
	}

	if len(actual.Shows) != len(es) {
		t.Errorf("Expected %v to equal %v", len(actual.Shows), len(es))
	}
}

/*
** PRIVATE METHODS
**
**
 */
func userComp(t *testing.T, expected, actual User) {
	if expected.ID != actual.ID {
		t.Errorf("Expected id %v to equal %v", actual.ID, expected.ID)
	}

	if expected.Username != actual.Username {
		t.Errorf("Expected username %v to equal %v", actual.Username, expected.Username)
	}
}
