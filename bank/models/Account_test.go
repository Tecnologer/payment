package models

import "testing"

func TestAccount_Deposit(t *testing.T) {
	type fields struct {
		Client  *Client
		Number  string
		Balance float32
	}

	type args struct {
		amount float32
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         bool
		balanceExpected float32
	}{
		{
			name: "deposit_100",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			args: args{
				amount: 100,
			},
			balanceExpected: 1100,
		},
		{
			name: "deposit_minus_150",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			args: args{
				amount: -150,
			},
			balanceExpected: 1000,
			wantErr:         true,
		},
		{
			name: "deposit_100_negative_balance",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: -90,
			},
			args: args{
				amount: 100,
			},
			balanceExpected: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Client:  tt.fields.Client,
				Number:  tt.fields.Number,
				Balance: tt.fields.Balance,
			}

			err := a.Deposit(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}

			if a.Balance != tt.balanceExpected {
				t.Errorf("Deposit() = %v, want %v", a.Balance, tt.balanceExpected)
			}
		})
	}
}

func TestAccount_GetID(t *testing.T) {
	type fields struct {
		Client  *Client
		Number  string
		Balance float32
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get_id",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Client:  tt.fields.Client,
				Number:  tt.fields.Number,
				Balance: tt.fields.Balance,
			}

			if got := a.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	type fields struct {
		Client  *Client
		Number  string
		Balance float32
	}

	type args struct {
		amount float32
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         bool
		expectedBalance float32
	}{
		{
			name: "withdraw_100",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			args: args{
				amount: 100,
			},
			expectedBalance: 900,
		},
		{
			name: "withdraw_100_negative_amount",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			args: args{
				amount: -100,
			},
			expectedBalance: 1000,
			wantErr:         true,
		},
		{
			name: "withdraw_1100_insufficient_funds",
			fields: fields{
				Client:  &Client{Name: "Tecnologer"},
				Number:  "123456",
				Balance: 1000,
			},
			args: args{
				amount: 1100,
			},
			expectedBalance: 1000,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Client:  tt.fields.Client,
				Number:  tt.fields.Number,
				Balance: tt.fields.Balance,
			}
			if err := a.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}

			if a.Balance != tt.expectedBalance {
				t.Errorf("Withdraw() = %v, want %v", a.Balance, tt.expectedBalance)
			}
		})
	}
}
