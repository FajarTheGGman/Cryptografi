echo "Installing paxlage...."
apt-get install golang -y
go get -u golang.org/x/crypto/ripemd160
go get -u golang.org/x/crypto/md4
go run crypto.go
