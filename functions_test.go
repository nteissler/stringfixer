package stringfixer

import "testing"

func TestFirstCapital(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"nick", "Nick"},
		{"  pooP", "Poop"},
		{"123hi", "123hi"},
		{"", ""},
		{"NoPe", "Nope"},
	}
	for _, c := range cases {
		got := FirstCapital(c.in)
		if got != c.want {
			t.Errorf("FirstCapital(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDeleteExtension(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"nick.txt", "nick"},
		{"/path/to/file.txt", "file"},
		{"badfile.", "badfile."},
		{"another bad", "another bad"},
	}
	for _, c := range cases {
		got, _ := DeleteExtension(c.in)
		if got != c.want {
			t.Errorf("DeleteExtension(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
