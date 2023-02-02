package structcheck

import (
	"testing"

	"emperror.dev/errors"
	"github.com/stretchr/testify/require"
)

func TestCheckResourceInterfacesSetted(t *testing.T) {
	err := CheckResourceInterfacesSetted(nil)

	require.Error(t, err, "nada definido para verificação - irá retornar error")

	type iTest interface {
		TestFunc()
	}
	type iTest1 interface {
		TestFunc()
	}

	type testService struct {
		Test  iTest
		Test1 iTest1
		TTee  iTest
	}

	ttpqp := struct{ iTest }{}

	tt := testService{
		Test1: &ttpqp,
	}

	err = CheckResourceInterfacesSetted(&tt)

	require.Error(t, err, "recurso não definido - deve retornar error")

	require.Greater(t, len(errors.GetErrors(err)), 0, "neste cenário deve-se retornar diversos errors")

	err = CheckResourceInterfacesSetted(nil)

	require.Error(t, err, "Struct informada como nil para verificar do camos - irá retornar error")

	err = CheckResourceInterfacesSetted(struct{}{})

	require.Error(t, err, "struct má informada ocorrerá panic - mas a função deve recuperar e retornar como error")

	require.NotPanics(t, func() {
		err := CheckResourceInterfacesSetted(struct{}{})
		require.Error(t, err, "Como não ocorre panic deve recuperar e retornar como erro")
	}, "Essa função não pode ocorrer panics")

	tt = testService{
		Test:  &ttpqp,
		Test1: &ttpqp,
		TTee:  &ttpqp,
	}

	err = CheckResourceInterfacesSetted(&tt)

	require.Nil(t, err, "Todas as interfaces informadas não deve ocorrer error")
}

func TestCheckResourceStructsSetted(t *testing.T) {
	err := CheckResourceStructsSetted(nil)

	require.Error(t, err, "nada definido para verificação - irá retornar error")

	type StructTest struct {
		Field        string ""
		FieldPointer *int64
	}

	type testService struct {
		Test  *StructTest
		Test1 StructTest
		TTee  *StructTest
		TTee2 *StructTest
	}

	tt := testService{
		Test:  &StructTest{},
		Test1: StructTest{},
		TTee: &StructTest{
			Field:        "",
			FieldPointer: new(int64),
		},
		TTee2: nil,
	}

	err = CheckResourceStructsSetted(&tt)

	require.Error(t, err, "recurso não definido - deve retornar error")

	require.Greater(t, len(errors.GetErrors(err)), 0, "neste cenário deve-se retornar diversos errors")

	err = CheckResourceStructsSetted(nil)

	require.Error(t, err, "Struct nil informada - deverá ocorrer error")

	err = CheckResourceStructsSetted(struct{}{})

	require.Error(t, err, "struct má informada ocorrerá panic - mas a função deve recuperar e retornar como error")

	require.NotPanics(t, func() {
		err := CheckResourceStructsSetted(struct{}{})
		require.Error(t, err, "Como não ocorre panic deve recuperar e retornar como erro")
	}, "Essa função não pode ocorrer panics")

	n := int64(123)

	tt = testService{
		Test: &StructTest{
			Field:        "test",
			FieldPointer: &n,
		},
		Test1: StructTest{
			Field:        "statemented",
			FieldPointer: &n,
		},
		TTee: &StructTest{
			Field:        "another",
			FieldPointer: &n,
		},
		TTee2: nil,
	}

	err = CheckResourceStructsSetted(&tt)

	require.Error(t, err, "Teste uma struct declaradas como nil deve ocorrer error")

	tt = testService{
		Test: &StructTest{
			Field:        "test",
			FieldPointer: &n,
		},
		Test1: StructTest{
			Field:        "statemented",
			FieldPointer: &n,
		},
		TTee: &StructTest{
			Field:        "another",
			FieldPointer: &n,
		},
		TTee2: &StructTest{
			Field:        "sdfsdf",
			FieldPointer: &n,
		},
	}

	err = CheckResourceStructsSetted(&tt)

	require.Nil(t, err, "Structs declaradas não deve ocorrer error")
}

func TestCheckStructFilled(t *testing.T) {

	type iTest interface {
		TestFunc()
	}

	type testStrct struct {
		field string
	}

	type testService struct {
		Test      iTest
		Test1     iTest
		TTee      iTest
		StrctTest testStrct
	}

	tests := []struct {
		name      string
		resources testService
		wantErr   bool
	}{
		{
			name: "Struct vazia",
			resources: testService{
				Test:      nil,
				Test1:     nil,
				TTee:      nil,
				StrctTest: testStrct{},
			},
			wantErr: true,
		},
		{
			name: "Struct preenchida sem preencher as interfaces",
			resources: testService{
				Test:  nil,
				Test1: nil,
				TTee:  nil,
				StrctTest: testStrct{
					field: "Ad culpa sint quis anim est proident excepteur Lorem ipsum.",
				},
			},
			wantErr: true,
		},
		{
			name: "Struct preenchida completamente",
			resources: testService{
				Test:  &struct{ iTest }{},
				Test1: &struct{ iTest }{},
				TTee:  &struct{ iTest }{},
				StrctTest: testStrct{
					field: "Ad culpa sint quis anim est proident excepteur Lorem ipsum.",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := CheckStructFilled(&tt.resources)

			require.Equal(t, tt.wantErr, err != nil)

		})
	}
}
