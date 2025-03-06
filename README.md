# go-openfda
A Open FDA API client to make using the FDA's data easy with go.

## Coverage
Only handles the devices endpoint. I had plans to add more but I changed my profession

## Needed
- authorized requests
- paginiation w/ search after
- rate limiting
- backoff mechanism
- nice printing of fda objects
- total query limit reaches
- graceful error  handling

# Devices Endpoints
-[x] 510k
-[x] Classification
-[x] Covid19Serology
-[x] Enforcement
-[x] Events
-[x] PMA
-[x] Recall
-[x] Registrations and Listing
-[x] UDI

## Usage

Make writing queries to the openFDA api easier with go. 

If you need more heavy duty use see:
https://open.fda.gov/apis/downloads/

## Issues 

## Author
Nick Janapol

## Contributing 
Welcome. Take this over, give it full coverage etc. 

## License
MIT License

Copyright (c) 2024 Nick Janapol

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
