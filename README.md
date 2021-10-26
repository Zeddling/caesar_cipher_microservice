#   Caesar Cipher Microservice
This is an API that encrypts data using the [caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher).

### Endpoints
1. POST - ```/decrypt``` <br>
&emsp; data - text to be decrypted<br>
&emsp; shift - a negative integer number representing the position shifting degree<br>

&emsp; <b>Response</b> - the encrypted text

2. POST - ```/encrypt``` <br>
&emsp; data - text to be encrypted<br>
&emsp; shift - a positive integer number representing the position shifting degree<br>

&emsp; <b>Response</b> - the encrypted text