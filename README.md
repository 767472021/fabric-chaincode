# Flogo extension for Hyperledger Fabric chaincode

This Flogo extension is designed to allow developers to design and implement Hyperledger Fabric chaincode in the Flogo visual programming environment. This extension supports the following release versions:

- [Flogo Web UI](http://www.flogo.io/)
- [Hyperledger Fabric 2.2](https://www.hyperledger.org/projects/fabric)

The [Transaction Trigger](trigger/transaction) allows you to configure the chaincode transaction using normal input parameters and/or transient data.

It supports the following activities for storing and querying data on the distributed ledger and/or on private data collections.

- [Put](activity/put): Insert or update data on the distributed ledger or a private data collection, and optionally update associated compsite keys if they are specified.
- [Put All](activity/putall): Insert a list of records on the distributed ledger or a private collection, and optionally update composite keys of each record.
- [Get](activity/get): Retrieve the current state by a specified key from the distributed ledger or a private collection.
- [Get by Range](activity/getrange): Retrieve all states in a specified range of keys from the distributed ledger or a private collection. It supports resultset pagination for states from the distributed ledger.
- [Get by Composite Key](activity/getbycompositekey): Retrieve all states by a composite-key filter from the distributed ledger or a private collection. It supports resultset pagination for states from the distributed ledger.
- [Get History](activity/gethistory): Retrieve the state history of a specified key for data on the distributed ledger.
- [Query](activity/query): Retrieve all states by a Couchdb query statement from the distributed ledger or a private collection. It supports resultset pagination for states from the distributed ledger.
- [Delete](activity/delete): Mark the state as deleted for a specified key from the distributed ledger or a private collection, and deletes its composite keys. Optionally, it can delete only the state, or only a composite key.
- [Set Event](activity/setevent): Set a specified event and payload for a blockchain transaction.
- [Set Endorsement Policy](activity/endorsement): Set state-based endorsement policy by adding or deleting an endorsement organization, or by specifying a new endorsement policy.
- [Invoke Chaincode](activity/invokechaincode): Invoke a local chaincode, and returns response data from the called transaction.

With these extensions, Hyperledger Fabric chaincode can be designed and implemented with zero code.

## Getting Started

- Download and setup Golang from [here](https://golang.org/dl/).
- Clone this repo into an empty working directory: `git clone https://github.com/open-dovetail/fabric-chaincode.git`
- Setup develop environment by executing the script: `scripts/setup.sh`
- Build and run a sample Flogo model [marble](./samples/marble) as described in [README.md](./samples/marble/README.md)