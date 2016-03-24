# go-template

Using godep for dependency management:
https://github.com/tools/godep

### Install godep
<pre><code>
$ go get github.com/tools/godep
</code></pre>

Run from *.../github.com/tools/godep*
<code><pre>
$ go install
</code></pre>

### Get a dependency
<pre><code>
$ go get github.com/.../somedependency
</code></pre>


### Save dependency
<pre><code>
$ godep save
</code></pre>

### Usage
<pre><code>
# Use this instead of the go toolchain
$ godep go build
$ godep go install
</code></pre>

### Links
A good intro on the issues with dependency management in Go:
http://jbeckwith.com/2015/05/29/dependency-management-go/
