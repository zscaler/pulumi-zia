package provider

import (
	"reflect"
	"sort"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
)

// diffSkippingUnsetInputs compares two Args structs (old from state, new from user inputs)
// and builds a DetailedDiff map. For any optional field (pointer or slice) where the new
// input is nil, the diff is skipped — the user didn't set that field, so changes from the
// API should be ignored. This mirrors Terraform's Optional+Computed behaviour.
//
// Slice comparisons are order-insensitive (matching Terraform TypeSet semantics).
func diffSkippingUnsetInputs(oldArgs, newArgs any) (map[string]p.PropertyDiff, bool) {
	diff := map[string]p.PropertyDiff{}

	oldVal := reflect.ValueOf(oldArgs)
	newVal := reflect.ValueOf(newArgs)

	if oldVal.Kind() == reflect.Ptr {
		oldVal = oldVal.Elem()
	}
	if newVal.Kind() == reflect.Ptr {
		newVal = newVal.Elem()
	}

	typ := newVal.Type()
	diffFields(typ, oldVal, newVal, "", diff)

	return diff, len(diff) > 0
}

func diffFields(typ reflect.Type, oldVal, newVal reflect.Value, prefix string, diff map[string]p.PropertyDiff) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if !field.IsExported() {
			continue
		}

		pulumiTag := field.Tag.Get("pulumi")
		if pulumiTag == "" || pulumiTag == "-" {
			continue
		}
		tagParts := strings.Split(pulumiTag, ",")
		fieldName := tagParts[0]

		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		oldField := oldVal.Field(i)
		newField := newVal.Field(i)

		// For optional fields (pointers, slices, maps): if the user didn't set the
		// field (nil), skip the comparison entirely — treat it as "no opinion."
		switch newField.Kind() {
		case reflect.Ptr:
			if newField.IsNil() {
				continue
			}
		case reflect.Slice:
			if newField.IsNil() {
				continue
			}
		case reflect.Map:
			if newField.IsNil() {
				continue
			}
		}

		if !equalValues(oldField, newField) {
			kind := p.Update
			if isZeroOrNil(oldField) {
				kind = p.Add
			}
			diff[fieldName] = p.PropertyDiff{Kind: kind, InputDiff: true}
		}
	}
}

// equalValues compares two reflect.Values. For slices of strings or ints it uses
// set-equality (order-insensitive) to match Terraform's TypeSet semantics.
// For slices of structs it compares element-by-element, skipping nil fields in
// b (the new/user inputs) so that computed fields like nested IDs don't cause drift.
// For *bool pointers it normalises nil to false so that an unset optional bool
// is equivalent to an explicit false.
func equalValues(a, b reflect.Value) bool {
	if a.Kind() == reflect.Ptr && b.Kind() == reflect.Ptr &&
		a.Type().Elem().Kind() == reflect.Bool && b.Type().Elem().Kind() == reflect.Bool {
		aBool := !a.IsNil() && a.Elem().Bool()
		bBool := !b.IsNil() && b.Elem().Bool()
		return aBool == bBool
	}

	if a.Kind() == reflect.Slice && b.Kind() == reflect.Slice {
		if a.IsNil() && b.IsNil() {
			return true
		}
		if a.Len() != b.Len() {
			return false
		}
		elemKind := a.Type().Elem().Kind()
		if elemKind == reflect.String {
			return stringSliceSetEqual(a, b)
		}
		if elemKind == reflect.Int {
			return intSliceSetEqual(a, b)
		}
		if elemKind == reflect.Struct {
			return structSliceEqualSkippingNil(a, b)
		}
	}
	return reflect.DeepEqual(a.Interface(), b.Interface())
}

// structSliceEqualSkippingNil compares two slices of structs element-by-element.
// For each pair, nil pointer/slice fields in b (new inputs) are treated as "not
// provided" and skipped, matching the top-level diffFields behaviour.
func structSliceEqualSkippingNil(a, b reflect.Value) bool {
	for i := 0; i < a.Len(); i++ {
		if !structEqualSkippingNil(a.Index(i), b.Index(i)) {
			return false
		}
	}
	return true
}

func structEqualSkippingNil(a, b reflect.Value) bool {
	typ := b.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		pulumiTag := field.Tag.Get("pulumi")
		if pulumiTag == "" || pulumiTag == "-" {
			continue
		}
		bField := b.Field(i)
		aField := a.Field(i)

		switch bField.Kind() {
		case reflect.Ptr:
			if bField.IsNil() {
				continue
			}
		case reflect.Slice:
			if bField.IsNil() {
				continue
			}
		case reflect.Map:
			if bField.IsNil() {
				continue
			}
		}

		if !equalValues(aField, bField) {
			return false
		}
	}
	return true
}

func stringSliceSetEqual(a, b reflect.Value) bool {
	as := make([]string, a.Len())
	bs := make([]string, b.Len())
	for i := 0; i < a.Len(); i++ {
		as[i] = a.Index(i).String()
	}
	for i := 0; i < b.Len(); i++ {
		bs[i] = b.Index(i).String()
	}
	sort.Strings(as)
	sort.Strings(bs)
	for i := range as {
		if as[i] != bs[i] {
			return false
		}
	}
	return true
}

func intSliceSetEqual(a, b reflect.Value) bool {
	as := make([]int, a.Len())
	bs := make([]int, b.Len())
	for i := 0; i < a.Len(); i++ {
		as[i] = int(a.Index(i).Int())
	}
	for i := 0; i < b.Len(); i++ {
		bs[i] = int(b.Index(i).Int())
	}
	sort.Ints(as)
	sort.Ints(bs)
	for i := range as {
		if as[i] != bs[i] {
			return false
		}
	}
	return true
}

func isZeroOrNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		return v.IsNil()
	default:
		return v.IsZero()
	}
}
