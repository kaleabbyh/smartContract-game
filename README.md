# game

# Implemented Features
* user registration
* user login
* create pool
* get all pools with their respective poollist
* get a single pool
* sort a poollis of a pool by their respective values
* delete a pool by ID
* get conctract-addresses
* deploy escrow
* deposit-eth
* deposit-token
* withdraw-token

# Download & Setup Instructions

* 1 - Clone project
* 2 - add .env file on main directory
* 3 - add values for envirnment variable on .env file
    # Replace with your db credentials
    * DB_URL=
    
    #### 
    * SECRET_KEY=
    * NEXTAUTH_SECRET=

    # Keystore password and public address
    * PASSWORD=
    * Admin=

    # For test scripts
    * ADMIN_TOKEN=
    * USER_TOKEN=


# Install if package is not installed
* 1 - cd main directory
* 2 - go get .

# Tests 
migration -> createPool -> createPlayerScores -> getWinners -> updatePool(game) -> deletePool(game)

* STEPS
* Add these variables in .env file
    * ADMIN_TOKEN="Bearer admin-token-here"
    * USER_TOKEN="Bearer user-token-here"
* Open two terminals on main directory
* On the first terminal: go run main.go
* On the second terminal: go run test/test.go

# Test GET-WINNERS:
* open test/get_winners.sh
* substitute your TOKEN and GAME_ID
* run chmod +x test/get_winners.sh
* run ./test/get_winners.sh