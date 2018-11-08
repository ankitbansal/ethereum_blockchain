// https://hackernoon.com/a-step-by-step-guide-to-testing-and-deploying-ethereum-smart-contracts-in-go-9fc34b178d78
pragma solidity ^0.4.17;

contract Rating {

     struct RatingType {
    	uint value;
    	string providerName;
    	string customerName;
    }

    RatingType public rating;

    function Rating(uint value, string providerName, string customerName) public {
        rating = RatingType(value, providerName, customerName);
    }

}