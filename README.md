
# Node Bootstrap Readme Docs <sup>[![Version Badge](:sunglasses:)](https://npmjs.org/package/readme-docs)</sup>

Generate a sexy docs page with Node, Connect or Express, Twitter Bootstrap, and your Github repo's Readme.md file.

Check out an example generated page at <https://getprove.com/docs> and its respective [Readme.md](https://github.com/getprove/prove-api/blob/master/Readme.md).

You can also view an example of this repo itself by visiting <https://getprove.github.io/node-bootstrap-readme-docs>.


## Index

* [How it Works](#how-it-works)
* [Format](#format)
* [Examples](#examples)
* [Changelog](#changelog)
* [Contributors](#contributors)
* [License](#license)


## How it Works

### Command Line

Easily generate a Twitter Bootstrap HTML file that you can use in your `gh-pages` branch.

```bash
# install the bin script
npm install -g readme-docs

# view help if needed
readme-docs --help

# example usage #1 (build Readme.md from current working directory)

## download a Readme.md file
wget https://raw.github.com/getprove/node-bootstrap-readme-docs/master/Readme.md

## generate the html (uses cwd's Readme.md file)
readme-docs
Compiled readme docs to ./build/

# example usage #2 (set custom page title)
readme-docs -t 'My Awesome API'
Compiled readme docs to ./build/

# example usage #3 (set custom path to file)
readme-docs -i /some/path/to/Readme.md
Compiled readme docs to ./build/

# example usage #4 (add a link to the Github repo)
readme-docs -g https://github.com/user/repo
Compiled readme docs to ./build/

# example usage #4 (add a custom css file)
readme-docs -c custom.css
Compiled readme docs to ./build/
```

### Route Middleware

Express and Connect expose `app.settings` as `settings` in your views.

This module populates `app.settings.readmeDocs` with an object of readme's.

Simply require this module, then implement it with route middleware:

```js
var express = require('express')
  , app = express()
  , readmeDocs = require('readme-docs')(app)

app.set('view engine', 'jade')

app.use(express.static(__dirname + '/public'))

app.get('/', readmeDocs.middleware('getprove', 'node-bootstrap-readme-docs'), function(req, res, next) {
  res.render('docs', {
      title: 'Node Bootstrap Readme Docs'
    , github: 'https://github.com/getprove/node-bootstrap-readme-docs'
  })
})

app.listen(3000)

console.log('Visit <http://localhost:3000> in your browser to view the example.')
```

You can use any templating language, here are a few example implementations:

**Jade**:

```jade
ul!= settings.readmeDocs.getprove['node-bootstrap-readme-docs'].nav
div!= settings.readmeDocs.getprove['node-bootstrap-readme-docs'].body
```

**EJS**:

```ejs
<ul>
  <%= settings.readmeDocs.getprove['node-bootstrap-readme-docs'].nav %>
</ul>
<div>
  <%= settings.readmeDocs.getprove['node-bootstrap-readme-docs'].body %>
</div>
```

These examples both output the following HTML when rendered:

```html
<ul>
  <li><a href="#how-it-works">How it Works</a></li>
  <li><a href="#format">Format</a></li>
  <li><a href="#examples">Examples</a></li>
  <li><a href="#contributors">Contributors</a></li>
  <li><a href="#license">License</a></li>
</ul>
<div>
  <h1>Node Bootstrap Readme Docs <sup><a href="https://npmjs.org/package/readme-docs"><img src="http://vb.teelaun.ch/getprove/node-bootstrap-readme-docs.svg" alt="Version Badge"></a></sup></h1>
  <p>Generate a sexy docs page with Node, Connect or Express, Twitter Bootstrap, and your Github repo&#39;s Readme.md file.</p>
  <p>Check out an example generated page at <a href="https://getprove.com/docs">https://getprove.com/docs</a> and its respective <a href="https://github.com/getprove/prove-api/blob/master/Readme.md">Readme.md</a>.</p>
  <hr><h2><a href="#how-it-works" name="how-it-works" id="how-it-works">How it Works</a></h2>
  <p>Express and Connect expose <code>app.settings</code> as <code>settings</code> in your views.</p>
  <p>This module populates <code>app.settings.readmeDocs</code> with an object of readme&#39;s.</p>
  <p>Simply require this module, then implement it with route middleware:</p>
  <!-- ... continued HTML with github flavored markdown syntax ... -->
</div>
```


## Format

Your `Readme.md` file should contain the following format:

```md

# Docs

Some description of your awesome API.

## Index

* [First Link](#first-link)
* [Second Link](#second-link)
* [Third Link](#third-link)

## First Link

Some stuff goes here for the first link.

## Second Link

Some stuff goes here for the second link.

## Third Link

Some stuff goes here for the third link.

```

## Examples

See the [examples](https://github.com/getprove/node-bootstrap-readme-docs/tree/master/examples) folder for an Express &amp; Jade example.


## Contributors

* Nick Baugh <niftylettuce@gmail.com>


## Changelog

* 0.1.0 - Added custom CSS file inclusion with `-c` option
* 0.0.9 - Removed anchor styles for span.path
* 0.0.8 - Added typeahead autocomplete for <h2> tags
* 0.0.7 - Added CSS for NetDNA
* 0.0.6 - Allowed h1 tags to get converted to anchors
* 0.0.5 - Fixed regex parsing of headers to match Github parsing (e.g. #how-it-works? becomes #how-it-works), fixed CSS responsive issue with #spacer
* 0.0.4 - Resolved navigation bug (thanks to @oren)
* 0.0.3 - Resolved node binary env in bin script (thanks to @oren)
* 0.0.2 - Added CLI
* 0.0.1 - Initial release


## License

The MIT License

Copyright (c) 2013- Prove <support@getprove.com> (https://getprove.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.


