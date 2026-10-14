package order

import (
	"errors"
	"testing"
)

func TestNewOrder(t *testing.T) {
	tests := []struct {
		id      int
		wantErr error
	}{
		{1, nil},
		{0, ErrInvalidOrderID},
		{-1, ErrInvalidOrderID},
	}
	for _, tt := range tests {
		o, err := NewOrder(tt.id)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("NewOrder(%d) error = %v, want %v", tt.id, err, tt.wantErr)
		}
		if err == nil && (o == nil || o.ID != tt.id || o.Status != StatusNew) {
			t.Errorf("NewOrder(%d) = %+v", tt.id, o)
		}
		if err != nil && o != nil {
			t.Errorf("NewOrder(%d) returned order with error", tt.id)
		}
	}
}

func TestTotalEmpty(t *testing.T) {
	o, _ := NewOrder(1)
	if got := Total(*o); got != 0 {
		t.Errorf("Total() = %d, want 0", got)
	}
}

func TestTotalTwoItems(t *testing.T) {
	o, _ := NewOrder(1)
	if err := AddItem(o, Item{"Хлеб", 5000}); err != nil {
		t.Fatal(err)
	}
	if err := AddItem(o, Item{"Молоко", 9000}); err != nil {
		t.Fatal(err)
	}
	if len(o.Items) != 2 {
		t.Fatalf("len(Items) = %d, want 2", len(o.Items))
	}
	if got := Total(*o); got != 14000 {
		t.Errorf("Total() = %d, want 14000", got)
	}
}

func TestAddItemErrors(t *testing.T) {
	tests := []struct {
		name    string
		order   *Order
		item    Item
		wantErr error
	}{
		{"nil order", nil, Item{"Хлеб", 5000}, ErrNilOrder},
		{"empty name", &Order{ID: 1}, Item{"", 5000}, ErrEmptyItemName},
		{"negative price", &Order{ID: 1}, Item{"Хлеб", -1}, ErrNegativeItemPrice},
	}
	for _, tt := range tests {
		err := AddItem(tt.order, tt.item)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("%s: error = %v, want %v", tt.name, err, tt.wantErr)
		}
		if tt.order != nil && len(tt.order.Items) != 0 {
			t.Errorf("%s: invalid item was added", tt.name)
		}
	}
}

func TestCopyDoesNotChangeOriginal(t *testing.T) {
	o, _ := NewOrder(1)
	c := *o
	c.ID = 2
	c.Status = StatusPaid
	if o.ID != 1 || o.Status != StatusNew {
		t.Errorf("original changed: %+v", o)
	}
}

func TestStatusNames(t *testing.T) {
	for _, s := range []string{StatusNew, StatusPaid, StatusShipped, StatusDelivered, StatusCancelled} {
		if StatusNames[s] == "" {
			t.Errorf("no name for status %q", s)
		}
	}
}

func TestNilMapWritePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("writing to nil map should panic")
		}
	}()
	var m map[string]string
	m[StatusNew] = "Новый" //nolint:staticcheck // the panic is what this test checks
}

func TestMakeMapWrite(t *testing.T) {
	m := make(map[string]string)
	m[StatusNew] = "Новый"
	if m[StatusNew] != "Новый" {
		t.Errorf("m[%q] = %q", StatusNew, m[StatusNew])
	}
}
