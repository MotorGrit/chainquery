// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testInputs(t *testing.T) {
	t.Parallel()

	query := Inputs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testInputsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = input.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInputsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Inputs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInputsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InputSlice{input}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testInputsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := InputExists(tx, input.ID)
	if err != nil {
		t.Errorf("Unable to check if Input exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InputExistsG to return true, but got false.")
	}
}
func testInputsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	inputFound, err := FindInput(tx, input.ID)
	if err != nil {
		t.Error(err)
	}

	if inputFound == nil {
		t.Error("want a record, got nil")
	}
}
func testInputsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Inputs(tx).Bind(input); err != nil {
		t.Error(err)
	}
}

func testInputsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Inputs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInputsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	inputOne := &Input{}
	inputTwo := &Input{}
	if err = randomize.Struct(seed, inputOne, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}
	if err = randomize.Struct(seed, inputTwo, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = inputOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = inputTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Inputs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInputsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	inputOne := &Input{}
	inputTwo := &Input{}
	if err = randomize.Struct(seed, inputOne, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}
	if err = randomize.Struct(seed, inputTwo, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = inputOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = inputTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testInputsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInputsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx, inputColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInputToManySpentByInputOutputs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, outputDBTypes, false, outputColumnsWithDefault...)
	randomize.Struct(seed, &c, outputDBTypes, false, outputColumnsWithDefault...)

	b.SpentByInputID.Valid = true
	c.SpentByInputID.Valid = true
	b.SpentByInputID.Uint64 = a.ID
	c.SpentByInputID.Uint64 = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	output, err := a.SpentByInputOutputs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range output {
		if v.SpentByInputID.Uint64 == b.SpentByInputID.Uint64 {
			bFound = true
		}
		if v.SpentByInputID.Uint64 == c.SpentByInputID.Uint64 {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := InputSlice{&a}
	if err = a.L.LoadSpentByInputOutputs(tx, false, (*[]*Input)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SpentByInputOutputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SpentByInputOutputs = nil
	if err = a.L.LoadSpentByInputOutputs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SpentByInputOutputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", output)
	}
}

func testInputToManyAddOpSpentByInputOutputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c, d, e Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Output{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Output{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSpentByInputOutputs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SpentByInputID.Uint64 {
			t.Error("foreign key was wrong value", a.ID, first.SpentByInputID.Uint64)
		}
		if a.ID != second.SpentByInputID.Uint64 {
			t.Error("foreign key was wrong value", a.ID, second.SpentByInputID.Uint64)
		}

		if first.R.SpentByInput != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SpentByInput != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SpentByInputOutputs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SpentByInputOutputs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SpentByInputOutputs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testInputToManySetOpSpentByInputOutputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c, d, e Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Output{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetSpentByInputOutputs(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SpentByInputOutputs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSpentByInputOutputs(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SpentByInputOutputs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.SpentByInputID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.SpentByInputID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ID != d.SpentByInputID.Uint64 {
		t.Error("foreign key was wrong value", a.ID, d.SpentByInputID.Uint64)
	}
	if a.ID != e.SpentByInputID.Uint64 {
		t.Error("foreign key was wrong value", a.ID, e.SpentByInputID.Uint64)
	}

	if b.R.SpentByInput != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SpentByInput != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SpentByInput != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.SpentByInput != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SpentByInputOutputs[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SpentByInputOutputs[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testInputToManyRemoveOpSpentByInputOutputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c, d, e Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Output{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddSpentByInputOutputs(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SpentByInputOutputs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSpentByInputOutputs(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SpentByInputOutputs(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.SpentByInputID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.SpentByInputID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.SpentByInput != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SpentByInput != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SpentByInput != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.SpentByInput != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SpentByInputOutputs) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SpentByInputOutputs[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SpentByInputOutputs[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testInputToOneAddressUsingAddress(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Input
	var foreign Address

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	local.AddressID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AddressID.Uint64 = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Address(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := InputSlice{&local}
	if err = local.L.LoadAddress(tx, false, (*[]*Input)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Address = nil
	if err = local.L.LoadAddress(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInputToOneTransactionUsingTransaction(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Input
	var foreign Transaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TransactionID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Transaction(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := InputSlice{&local}
	if err = local.L.LoadTransaction(tx, false, (*[]*Input)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Transaction = nil
	if err = local.L.LoadTransaction(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInputToOneSetOpAddressUsingAddress(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Address{&b, &c} {
		err = a.SetAddress(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Address != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Inputs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AddressID.Uint64 != x.ID {
			t.Error("foreign key was wrong value", a.AddressID.Uint64)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AddressID.Uint64))
		reflect.Indirect(reflect.ValueOf(&a.AddressID.Uint64)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AddressID.Uint64 != x.ID {
			t.Error("foreign key was wrong value", a.AddressID.Uint64, x.ID)
		}
	}
}

func testInputToOneRemoveOpAddressUsingAddress(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetAddress(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveAddress(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Address(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Address != nil {
		t.Error("R struct entry should be nil")
	}

	if a.AddressID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Inputs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testInputToOneSetOpTransactionUsingTransaction(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Input
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Transaction{&b, &c} {
		err = a.SetTransaction(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Transaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Inputs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID, x.ID)
		}
	}
}
func testInputsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = input.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testInputsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InputSlice{input}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testInputsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Inputs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	inputDBTypes = map[string]string{`AddressID`: `bigint`, `Coinbase`: `varchar`, `Created`: `datetime`, `ID`: `bigint`, `IsCoinbase`: `tinyint`, `Modified`: `datetime`, `PrevoutHash`: `varchar`, `PrevoutN`: `int`, `PrevoutSpendUpdated`: `tinyint`, `ScriptSigHex`: `text`, `ScriptSigSSM`: `text`, `Sequence`: `int`, `TransactionHash`: `varchar`, `TransactionID`: `bigint`, `Value`: `decimal`}
	_            = bytes.MinRead
)

func testInputsUpdate(t *testing.T) {
	t.Parallel()

	if len(inputColumns) == len(inputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	if err = input.Update(tx); err != nil {
		t.Error(err)
	}
}

func testInputsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(inputColumns) == len(inputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	input := &Input{}
	if err = randomize.Struct(seed, input, inputDBTypes, true, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, input, inputDBTypes, true, inputPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(inputColumns, inputPrimaryKeyColumns) {
		fields = inputColumns
	} else {
		fields = strmangle.SetComplement(
			inputColumns,
			inputPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(input))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := InputSlice{input}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testInputsUpsert(t *testing.T) {
	t.Parallel()

	if len(inputColumns) == len(inputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	input := Input{}
	if err = randomize.Struct(seed, &input, inputDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = input.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Input: %s", err)
	}

	count, err := Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &input, inputDBTypes, false, inputPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	if err = input.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Input: %s", err)
	}

	count, err = Inputs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
