#!/bin/bash

# Import variables from .env file
source .env

# Send the request to retrieve the winners
curl -X GET -H "Authorization: Bearer $TOKEN" http://localhost:8080/play/get-winners/$GAME_ID