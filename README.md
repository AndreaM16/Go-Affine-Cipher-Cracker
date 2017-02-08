# Go-Affine-Cipher-Cracker ![](https://stuff.mit.edu/afs/sipb/project/golang/arch/go1.2.1-linux-amd64/favicon.ico)
Affine Cipher Cracker in Golang

Find k = (a, b) used to encrypt a given plaintext into a given ciphertext.

##What you need to run it
 - `go : ~1.7.5`
 - `npm: ~4.1.2`

##How to use

 - Once downloaded, go to project's folder and run `go build src/main.go`. A bin will be generated
 - Clone https://github.com/AndreaM16/Affine-Cipher-Simulator
 - Install http://enclosejs.com/ with `sudo npm install -g enclose`
 - Go into newly cloned project and create a bin from the js file running: `enclose ./affine affine-cipher.js`. A new bin called affine will be generated.
 - Move the newly created bin inside Go-Affine-Cipher-Cracker folder
 - Run: `./main plaintext ciphertext`

###Where
 - `plaintext` is a chosen plaintext. Single words ONLY supported for now.
 - `ciphertext` is its encrypted corres corresponding. Single words ONLY supported for now.

##How it works
Let k = (a, b)
  - If `a` and `m` are coprime
  - Ek(n) = an + b mod m
  - Dk(y) = a^-1(y-b) mod m

It's possible to perform a `Known-Plaintext-Attack` by brute-force, `a` can assume only 12 values that are coprime with `m` (26), and `b` can vary in 26 values. We have a total number of `12 * 26 = 312` keys.

##Example

`./main affine ihhwvc` returns `a=5` and `b=8` which are the correct keys. For instance, if you run `./affine encrypt 5 8 affine` you obtain `ihhwvc` and if you run `./affine decrypt 5 8 ihhwvc` you obtain `affine`.
