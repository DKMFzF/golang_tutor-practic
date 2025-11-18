package name

import "testing"

func TestFullName(t *testing.T) {
	var userArr = []struct {
		TestName string
		Want     string
		User     *User
	}{
		{
			TestName: "#1",
			Want:     "Kirill Doroshev",
			User: &User{
				FirstName: "Kirill",
				LastName:  "Doroshev",
			},
		},
		{
			TestName: "#1",
			Want:     "Artem Artem",
			User: &User{
				FirstName: "Artem",
				LastName:  "Artem",
			},
		},
		{
			TestName: "#1",
			Want:     "Kirill ",
			User: &User{
				FirstName: "Kirill",
				LastName:  "",
			},
		},
		{
			TestName: "#1",
			Want:     " ",
			User: &User{
				FirstName: "",
				LastName:  "",
			},
		},
	}

	for _, v := range userArr {
		t.Run(v.TestName, func(t *testing.T) {
			if got := v.User.FullName(); got != v.Want {
				t.Errorf("%s: FullName() = %s; Want = %s", v.TestName, v.User.FullName(), v.Want)
			}
		})
	}
}
