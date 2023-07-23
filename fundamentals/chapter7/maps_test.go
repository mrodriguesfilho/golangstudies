package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, got := dictionary.Search("t3st")
		AssertError(t, got, errNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add("test", "this is just a test")
		AssertError(t, err, nil)
		AssertDefinition(t, dictionary, key, value)
	})

	t.Run("adding an existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add("test", "this is just a test")
		AssertError(t, err, errWordExists)
		AssertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newValue := "this is just an update"
		err := dictionary.Update(key, newValue)
		AssertError(t, err, nil)
		AssertDefinition(t, dictionary, key, newValue)
	})

	t.Run("update a word that doesnt exist", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		newValue := "this is just an update"
		err := dictionary.Update(key, newValue)
		AssertError(t, err, errWordDoesntExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete an existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		dictionary.Delete(key)
		_, err := dictionary.Search(key)
		if err != errNotFound {
			t.Errorf("Expected %q to be deleted", key)
		}
	})
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func AssertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	AssertStrings(t, got, value)
}
