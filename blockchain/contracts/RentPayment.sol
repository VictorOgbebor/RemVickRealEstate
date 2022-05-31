// SPDX-License-Identifier: BlockchainBic
pragma solidity 0.8.10;

address public renter;
address private landlord;

constructor(address _landlord) {
    landlord = _landlord;
}

modifier landlordOnly {
    
}

// subscription model
function leaseAgreement() returns (uint256) {
    
}

// deposit
function payRentInFull() public {

}

// withdraw
function withdrawFunds() public landlordOnly {
    
}

// SplitPayment
function InstallmentPayment() public {
    
}

function viewBalance()  returns (uint256) {
    
}

function viewLease()  returns () {
    
}

