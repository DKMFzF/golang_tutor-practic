package family

import (
	"maps"
	"reflect"
	"testing"
)

type TestPersonList struct {
	TestRole   Relationship
	TestPerson *Person
}

func TestAddNew(t *testing.T) {
	var familyTest = []struct {
		NameTest   string
		Want       *Family
		NewPersons []*TestPersonList
		WantErr    error
	}{
		{
			NameTest: "#1",
			Want: &Family{
				Members: map[Relationship]Person{
					Father: {
						FirstName: "Dima",
						LastName:  "Hueputov",
						Age:       28,
					},
					Mother: {
						FirstName: "Sobina",
						LastName:  "Hueputova",
						Age:       20,
					},
				},
			},
			NewPersons: []*TestPersonList{
				{
					TestRole: Mother,
					TestPerson: &Person{
						FirstName: "Sobina",
						LastName:  "Hueputova",
						Age:       28,
					},
				},
			},
			WantErr: ErrRelationshipAlreadyExists,
		},
		{
			NameTest: "#2",
			Want: &Family{
				Members: map[Relationship]Person{
					Father: {
						FirstName: "Kirill",
						LastName:  "Doroshev",
						Age:       20,
					},
				},
			},
			NewPersons: []*TestPersonList{
				{
					TestRole: Father,
					TestPerson: &Person{
						FirstName: "Dima",
						LastName:  "Hueputov",
						Age:       20,
					},
				},
			},
			WantErr: ErrRelationshipAlreadyExists,
		},
	}

	for _, v := range familyTest {
		t.Run(v.NameTest, func(t *testing.T) {
			// создаём семию для теста
			got := &Family{
				Members: map[Relationship]Person{},
			}

			// заполняем семью нашим начальным состоянием
			maps.Copy(got.Members, v.Want.Members)

			var gotErr error

			// добавляем новых персон и сохраняем ошибку
			for _, p := range v.NewPersons {
				gotErr = got.AddNew(p.TestRole, Person{
					FirstName: p.TestPerson.FirstName,
					LastName:  p.TestPerson.LastName,
					Age:       p.TestPerson.Age,
				})
			}

			if gotErr != v.WantErr {
				t.Fatalf("error = %v; want = %v", gotErr, v.WantErr)
			}

			if !reflect.DeepEqual(got, v.Want) {
				t.Fatalf("family mismatch:\n got =  %#v\n want = %#v", got, v.Want)
			}
		})
	}
}
