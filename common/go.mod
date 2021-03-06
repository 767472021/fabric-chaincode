module github.com/open-dovetail/fabric-chaincode/common

go 1.14

replace github.com/project-flogo/flow => github.com/yxuco/flow v1.1.1

replace github.com/project-flogo/core => github.com/yxuco/core v1.2.1

replace go.uber.org/multierr => go.uber.org/multierr v1.6.0

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200728190242-9b3ae92d8664
	github.com/hyperledger/fabric-protos-go v0.0.0-20190919234611-2a87503ac7c9
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852
	github.com/pkg/errors v0.9.1
	github.com/project-flogo/core v1.2.0
	github.com/project-flogo/flow v1.2.0
	github.com/stretchr/testify v1.6.1
)
