# Gist display gist information in the terminal

## Usage
### Display gist information
```shell
$ gist <id> --info
# output
Owner: dylan3g
Description: A very cool description
Gist id: ...
Created at: 2022-06-30T18:22:16Z
Updated at: 2022-06-30T18:22:17Z
Files:
	foo.go
```
### Display gist content
```shell
$ gist <id> --content
# output
func main() {
    print("hello")
}
```
### How can i get syntax highlight for the gist content?
You can use [bat](https://github.com/sharkdp/bat)
```shell
$ gist <id> --content | bat -l <language>
```
## Instalation
```shell
$ git clone https://github.com/dylan3g/gist.git
$ cd gist
$ go install . # or if you just want the binary, go build .
