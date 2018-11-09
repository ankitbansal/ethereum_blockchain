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
* To deploy contract on Rinkeby network, we use metamask chrome extension. Create an account and point it to Rinkeby TestNet
* To get fund, Rinkeby requires you to post your account address via any social site. I tweeted my account address.
* Once done, you need to provide that post link to Rinkeby Authentication Faucet. Once done, you should see ethers in your account. You can delete tweet afterwards.
* Next thing is to have a node to deploy on ethereum. I used Infura that provide you nodes you can use rather than creating one. You need to create a free account on Infura to use it.
* This provide node address url which we will use to connect.  We also need private key and paraphrase for Metamask in order to connect.
* Once you have these, put all of these in a file and name it conf.yaml (refer conf.yaml.template for file format)
* Next we use yaml library to parse file: go get gopkg.in/yaml.v2
* That's pretty much deployment configuration we need. We can go ahead and deploy it now.

# Retrieve Message
* Once the deploy is successful, an address will be provided. You can check on the rinkeby url that contract has been mined. This usually takes couple of mins
* Now we can go ahead and retrieve the message
  
  
  
  
