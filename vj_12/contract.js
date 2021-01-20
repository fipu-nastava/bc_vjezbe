$( document ).ready(function() {
    connect();
});

function connect() {
  window.web3 = new Web3(Web3.givenProvider);
  window.web3.currentProvider.enable();
}

async function get_owner() {
  var owner = await token.methods.get_owner().call();
  logOut("get_owner() result: " + owner);
}

function balanceOf() {
  var address = $("#balance_address").val();

  getBalanceOf(address);
}

function getBalanceOf(address) {
  if (!web3.utils.isAddress(address)){
    logOut("Invalid address: " + address);
    return;
  }

  token.methods.balanceOf(address).call()
    .then(balance => {
      logOut("balanceOf(" + address + ") result: " + web3.utils.fromWei(balance) + " eth, " + balance + " wei");
    })
    .catch(err => logOut(JSON.stringify(err)));
}

async function getDonation() {
  var acc = await selectedAccount();
  token.methods.getDonation().send({from: acc})
    .then(receipt => logOut(JSON.stringify(receipt)))
    .catch(err => logOut(JSON.stringify(err)));

}


async function myBalance() {
  var acc = await selectedAccount();
  getBalanceOf(acc);
}


async function destroy() {
  var acc = await selectedAccount();
  token.methods.destroy().send({from: acc})
    .then(receipt => logOut(JSON.stringify(receipt)))
    .catch(err => logOut(JSON.stringify(err)));
}

function connect_contract() {
  var abi = [
            	{
            		"inputs": [],
            		"payable": false,
            		"stateMutability": "nonpayable",
            		"type": "constructor"
            	},
            	{
            		"payable": true,
            		"stateMutability": "payable",
            		"type": "fallback"
            	},
            	{
            		"constant": true,
            		"inputs": [
            			{
            				"internalType": "address",
            				"name": "_owner",
            				"type": "address"
            			}
            		],
            		"name": "balanceOf",
            		"outputs": [
            			{
            				"internalType": "uint256",
            				"name": "balance",
            				"type": "uint256"
            			}
            		],
            		"payable": false,
            		"stateMutability": "view",
            		"type": "function"
            	},
            	{
            		"constant": false,
            		"inputs": [],
            		"name": "destroy",
            		"outputs": [],
            		"payable": false,
            		"stateMutability": "nonpayable",
            		"type": "function"
            	},
            	{
            		"constant": false,
            		"inputs": [
            			{
            				"internalType": "address payable",
            				"name": "_recipient",
            				"type": "address"
            			}
            		],
            		"name": "destroyAndSend",
            		"outputs": [],
            		"payable": false,
            		"stateMutability": "nonpayable",
            		"type": "function"
            	},
            	{
            		"constant": false,
            		"inputs": [],
            		"name": "getDonation",
            		"outputs": [],
            		"payable": false,
            		"stateMutability": "nonpayable",
            		"type": "function"
            	},
            	{
            		"constant": true,
            		"inputs": [],
            		"name": "get_owner",
            		"outputs": [
            			{
            				"internalType": "address",
            				"name": "",
            				"type": "address"
            			}
            		],
            		"payable": false,
            		"stateMutability": "view",
            		"type": "function"
            	},
            	{
            		"constant": true,
            		"inputs": [],
            		"name": "totalSupply",
            		"outputs": [
            			{
            				"internalType": "uint256",
            				"name": "",
            				"type": "uint256"
            			}
            		],
            		"payable": false,
            		"stateMutability": "view",
            		"type": "function"
            	}
            ];

  var contract_address = $('#contract_address').val();

  window.token = new web3.eth.Contract(abi, contract_address);

  logOut("Token contract connected to " + token.options.address);
}

async function selectedAccount() {
  var acs = await web3.eth.getAccounts();
  return acs[0];
}


function logOut(val) {
  $("#out").val(val);
}
