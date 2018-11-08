# Deploying Contracts on Ethereum Public 
In this example, we deploy a basic contract on Ethereum Network. Will use Rinkeby which is test platform for deploying ethereum contracts. We will use
go-ethereum as ethereum client.

# Setup
* Golang v1.10 or higher
* If on windows, install mingw 64 bit (for gcc) https://sourceforge.net/projects/tdm-gcc/
* go get -d github.com/ethereum/go-ethereum && run go install ./..
* Install solidity compiler: https://github.com/ethereum/solidity/releases and update PATH variable
* Verify solc --version, verify geth version

# Create contract
* Create a folder named contracts 
* Create a file names rating.sol and define the contract format 
* Generate go code for this using abigen -sol rating.sol -pkg contracts -out rating.go 
* This will generate go code for our contract that we will use to interact with network 
* Before moving further, we should verify that contract working properly. Let's define a rating_test.go and write a unit test. 
  
# Deploy contract
go get gopkg.in/yaml.v2

  
  
  
  
