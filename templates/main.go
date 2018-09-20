// Copyright (c) 2018 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const mainSrc = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>{{ .Title }}</title>
	<style>
        body {
            font-family: Arial, "Helvetica Neue", Helvetica, sans-serif;
            -webkit-font-smoothing: antialiased;
            text-rendering: optimizeLegibility;
            line-height: 1.6;
            margin-bottom: 50px;
        }
        .content {
            max-width: 50em;
            margin: 0 auto;
            padding: 0 10px;
        }
        img {
            max-width: 100%;
            display: block;
            margin: 0 auto;
        }
        a {
            text-decoration: none;
        }
        a:link, .nav a, .nav a:visited {
            color: #07C;
        }
        a:hover, a:active, .nav a:visited:active, .nav a:visited:hover {
            color: #3af;
        }
        a:visited {
            color: #005999;
        }
        pre {
            overflow-x: auto;
            padding: 5px;
            background-color: #eee;
        }
        </style>
</head>
<body>
	<div class="content">
		<div class="nav">
			<a href="/">Home</a>
		</div>
		<hr>
		{{ .Content }}
	</div>
</body>
</html>`