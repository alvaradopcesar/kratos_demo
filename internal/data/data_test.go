package data

import (
	"reflect"
	"testing"
)

func Test_CustomerData(t *testing.T) {
	tests := []struct {
		name                   string
		customerSet            CustomerData
		wantErr                bool
		customerGet            CustomerData
		customerGetResponse    CustomerData
		customerGetAllResponse []CustomerData
	}{
		{
			name:                   "UserCase01",
			customerSet:            CustomerData{ID: "01", Name: "Cesar"},
			wantErr:                false,
			customerGet:            CustomerData{ID: "01"},
			customerGetResponse:    CustomerData{ID: "01", Name: "Cesar"},
			customerGetAllResponse: []CustomerData{{ID: "01", Name: "Cesar"}},
		},
		{
			name:                   "UserCase02",
			customerSet:            CustomerData{ID: "02", Name: "Andrew"},
			wantErr:                false,
			customerGet:            CustomerData{ID: "02"},
			customerGetResponse:    CustomerData{ID: "02", Name: "Andrew"},
			customerGetAllResponse: []CustomerData{{ID: "01", Name: "Cesar"}, {ID: "02", Name: "Andrew"}},
		},
	}
	r := &universalClient{
		DataInMemory: make(map[string]CustomerData),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := r.Set(tt.customerSet); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.Get(tt.customerGet.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.customerGetResponse {
				t.Errorf("Get() got = %v, want %v", got, tt.customerGetResponse)
			}

			gotAll, err := r.GetAllKeys()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAll, tt.customerGetAllResponse) {
				t.Errorf("GetAllKeys() got = %v, want %v", gotAll, tt.customerGetAllResponse)
			}
		})
	}
}

func Test_CustomerData_Exist(t *testing.T) {
	tests := []struct {
		name        string
		customerSet CustomerData
		wantErr     bool
		existKey    string
		existErr    bool
		existWant   bool
	}{
		{
			name:        "UserCase01",
			customerSet: CustomerData{ID: "01", Name: "Cesar"},
			wantErr:     false,
			existKey:    "01",
			existErr:    true,
			existWant:   true,
		},
		{
			name:        "UserCase02",
			customerSet: CustomerData{ID: "02", Name: "Andrew"},
			wantErr:     false,
			existKey:    "04",
			existErr:    false,
			existWant:   false,
		},
	}
	r := &universalClient{
		DataInMemory: make(map[string]CustomerData),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := r.Set(tt.customerSet); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := r.Exists(tt.existKey)
			if got != tt.existWant {
				t.Errorf("Exists() got = %v, want %v", got, tt.existWant)
			}

		})
	}
}

func Test_universalClient_Exists(t *testing.T) {
	type fields struct {
		DataInMemory map[string]CustomerData
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &universalClient{
				DataInMemory: tt.fields.DataInMemory,
			}
			got, err := r.Exists(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
