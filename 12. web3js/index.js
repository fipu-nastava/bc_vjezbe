async function init() {
    $out = $("#out");
    $out.val("");

    var rpc = $("#web3").val();

    if (rpc) {
        window.web3 = new Web3(new Web3.providers.HttpProvider(rpc));
    }
    else {
        // Available wallet provider such as MetaMask
        window.web3 = new Web3(Web3.givenProvider);
    }
    await window.web3.currentProvider.enable();


    web3.eth.getAccounts()
    .then(accounts => {
      $("#out").val(["Successfully connected to " + rpc,
                     "Accounts present:", ""]
                    .concat(accounts).join("\n"));
    }).catch(e => {
      console.error(e);
      $("#out").val(e);
    });
}

function transaction() {
    var src = $("#from").val();
    var dst = $("#to").val();
    var amount = $("#amount").val();

    web3.eth.sendTransaction({
        from: src,
        to: dst,
        value: web3.utils.toWei(amount, "ether")
    }, handleResult)
}

function handleResult(err, txid) {
    if (!err) {
        alert("Transaction created with txid: " + txid);
        console.log(err, txid);
        return
    }
    else {
        $("#out").val(JSON.stringify(err));
    }
}

async function selectedAccount() {
  var acs = await web3.eth.getAccounts();
  return acs[0];
}

async function deploy() {
    var src = await selectedAccount();
    //console.log(src);
    // bytecode_data = "0x60806040526806f05b59d3b200000060015534801561001d57600080fd5b506001546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610205806100726000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806318160ddd146100515780634bf717291461006f57806370a082311461007957806383197ef0146100d1575b600080fd5b6100596100db565b6040518082815260200191505060405180910390f35b6100776100e1565b005b6100bb6004803603602081101561008f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061016f565b6040518082815260200191505060405180910390f35b6100d96101b7565b005b60015481565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141561016d5760146000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b3373ffffffffffffffffffffffffffffffffffffffff16fffea265627a7a723158202897a3168c8726f5772162f91ae59ded451819d8ef4ccc8180c6f20bf7220c2e64736f6c634300050c0032"
    var bytecode_data = $("#bytecode").val()

    web3.eth.sendTransaction({
        from: src,
        gas: 3000000,
        data: bytecode_data
    }, handleResult);

}

function clearAll() {
    var fields = ["out", "from", "to", "web3", "code", "txid", "amount"]

    for(var f of fields) {
        $("#"+f).val("");
    }
}

function getTxStatus() {
    var txid = $("#txid").val();

    web3.eth.getTransactionReceipt(txid, (err, receipt) => {
        console.log(receipt);
        $("#out").val(JSON.stringify(receipt))
    });
}
