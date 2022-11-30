package transformer

import (
	"net"
	"reflect"
	"testing"
)

const ipv4AsInt32 = 3232235777
const ipv4AsString = "192.168.1.1"

func Test_Int2ip(t *testing.T) {
	type args struct {
		nn uint32
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{
			name: "test transform int32 to ipv4",
			args: args{
				nn: ipv4AsInt32,
			},
			want: net.ParseIP(ipv4AsString).To4(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2ip(tt.args.nn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int2ip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Ip2int(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "test transform ip to int32",
			args: args{
				ip: net.ParseIP(ipv4AsString).To4(),
			},
			want:    ipv4AsInt32,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ip2int(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("ip2int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ip2int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IpAsStringToInt(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "test transform ip as string to int",
			args: args{
				ip: ipv4AsString,
			},
			want:    ipv4AsInt32,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IpAsStringToInt("192.168.1.1")

			if (err != nil) != tt.wantErr {
				t.Errorf("ipAsStringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ipAsStringToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
