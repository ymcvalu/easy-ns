module easy-ns

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181025213731-e84da0312774
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181026203630-95b1ffbd15a5
)

require (
	github.com/sirupsen/logrus v1.1.1
	github.com/urfave/cli v1.20.0
	golang.org/x/crypto v0.0.0-20181025213731-e84da0312774 // indirect
	golang.org/x/sys v0.0.0-20181026203630-95b1ffbd15a5 // indirect
)
