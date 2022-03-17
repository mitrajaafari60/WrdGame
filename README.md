# Rest Api For 5-alphabet word Game based on Go. 


## Introduction

This Application creates a restful service with 2 paths/check , /start for 5 wordGame.

* ### *Description:*
    * This project is a demo and works as stand-alone backend-server and saves all data internally.

### Project Structure

* Project skeleton is as below:
    * the main part of the project and HTTP server is lunched with a defined path(/start, /check) in the base folder.
      postman collection and the binary file of the project in included here too.
        * controller
            * <sub>Methos/functions in  this folder controls the request formate and in cas of success service calls and the appropriate response is generated and sends back to the client </sub>
        * DataModel
            * <sub>In this folder, all used data model for restApi request and response and internal structures are defined, besides dictionary of 5 alphabet words also is included there.</sub> 
        * service
            * <sub>internal package that includes fundamental packages of the framework.</sub>

### how to lunch
  for easy use of the project for those who don't have a Go compiler in their system 
  they can just run simply (./worddDemo )  this command will lunch http server on port 8000
    - in case of need to change the listener port it can be defined as an input argument such as (./worddDemo -port=9000) 

  # urls
  BaseURL: /start
  Routes:
    -
      URL: start
      Description: assigns a random 5 alphabet word to the user-agent
      Method: GET
      Data: 
        http header: user-agent should be in the header of the request
        http body   : nothing
      Response :
        just standard rest response is respond in the response which means statusCode 200 for Ok cases, otherwise appropriate error status code for example 400 for empty User-Agent.
    -
      URL: /check
      Description: check the sent word in the request (which is json formate) and respond by checking the word which is assigned to that user-agent
      Method: POST
      Data : 
        http header: user-agent should be in the header of the request
        http body   : json struct sending word to be checked such as ({"word":"xxxxx"}) 
      Response :
        just standard rest response is respond in the response which means statusCode 200 for Ok cases, otherwise appropriate error status code for example 400 for empty User-Agent.
        Response :
          in case of Ok input (User-Agent in header and 5 alphabet word in body `{"word":"xxxxx"}`)
          -Body json formate struct such as
              {"resultColores":["","","","",""],"status":"error"} -->an unforseen error :( please try later
              {"resultColores":["Grey","Grey","Grey","Grey","Grey"],"status":"failed"}
              {"resultColores":["Green","Grey","Grey","Green","Grey"],"status":"failed"}
              {"resultColores":["Green","Green","Green","Green","Green"],"status":"success"}

          
          in case of NOT Ok input (User-Agent was NOT send or Null Body or alphabet word less that 5
            -http StatusCode 400     

### how to check        
  it is so simple! Don't worry just run and test the methods with the postman collection included in the base folder named as `WordGame.postman_collection.json`
  ## just  remember First call /start method and then /check method to  check the word!