// Copyright (c) 2017 Femtowiki authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const indexSrc = `
{{ define "content" }}
<div id="section-tabs">
	<div id="section-search">
		<form method="GET" action="">
			<input type="text" name="query" placeholder="Search Femtowiki">
			<input class="btn btn-default" type="submit" value="Search">
		</form>
	</div>
	<div id="section-tabs-right">
		<span><a href="">Source</a></span>
		<span class="active"><a href="">Read</a></span>
	</div>
	<div id="section-tabs-left">
		<span class="active"><a href="">Main</a></span>
		<span><a href="">Discussion</a></span>
	</div>
</div>
<div id="meat">
{{ if .IsEditMode }}
<form action="" method="POST">
	<textarea rows="50">Content</textarea>
	<input type="submit" class="btn btn-default" value="Update">
</form>
{{ else }}
<h1>Pagination in SQL</h1>
<p>
On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish. In a free hour, when our power of choice is untrammeled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.
</p>

<div class="toc">
<ol>
<li><a href="">Truncate the results</a></li>
<li>
	<a href="">Legal</a>
	<ul>
		<li><a href="">Jane vs. Roe</a></li>
		<li><a href="">Jon vs. Jane</a></li>
		<li><a href="">Wade vs. Roe</a></li>
	</ul>
</li>
<li><a href="">Variations</a></li>
</ol>
</div>

<h2>Truncate the results</h2>
<p>
Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.
Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.
A small river named Duden flows by their place and supplies it with the necessary regelialia.
It is a paradisematic country, in which roasted parts of sentences fly into your mouth.
Even the all-powerful Pointing has no control about the blind texts it is an almost unorthographic life One day however a small line of blind text by the name of Lorem Ipsum decided to leave for the far World of Grammar.
The Big Oxmox advised her not to do so, because there were thousands of bad Commas, wild Question Marks and devious Semikoli, but the Little Blind Text didn’t listen.
She packed her seven versalia, put her initial into the belt and made herself on the way.
When she reached the first hills of the Italic Mountains, she had a last view back on the skyline of her hometown Bookmarksgrove, the headline of Alphabet Village and the subline of her own road, the Line Lane.
Pityful a rethoric question ran over her cheek, then she continued her way. On her way she met a copy.
The copy warned the Little Blind Text, that where it came from it would have been rewritten a thousand times and everything that was left from its origin would be the word "and" and the Little Blind Text should turn around and return to its own, safe country.
But nothing the copy said could convince her and so it didn’t take long until a few insidious Copy Writers ambushed her, made her drunk with Longe and Parole and dragged her into their agency, where they abused her for their projects again and again. And if she hasn’t been rewritten, then they are still using her.
Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.
Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.
A small river named Duden flows by their place and supplies it with the necessary regelialia.
It is a paradisematic country, in which roasted parts of sentences fly into your mouth.
Even the all-powerful Pointing has no control about the blind texts it is an almost unorthographic life One
</p>
<h2>Legal</h2>
<p>
Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.
Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.
A small river named Duden flows by their place and supplies it with the necessary regelialia.
It is a paradisematic country, in which roasted parts of sentences fly into your mouth.
Even the all-powerful Pointing has no control about the blind texts it is an almost unorthographic life One day however a small line of blind text by the name of Lorem Ipsum decided to leave for the far World of Grammar.
The Big Oxmox advised her not to do so, because there were thousands of bad Commas, wild Question Marks and devious Semikoli, but the Little Blind Text didn’t listen.
She packed her seven versalia, put her initial into the belt and made herself on the way.
When she reached the first hills of the Italic Mountains, she had a last view back on the skyline of her hometown Bookmarksgrove, the headline of Alphabet Village and the subline of her own road, the Line Lane.
Pityful a rethoric question ran over her cheek, then she continued her way. On her way she met a copy.
The copy warned the Little Blind Text, that where it came from it would have been rewritten a thousand times and everything that was left from its origin would be the word "and" and the Little Blind Text should turn around and return to its own, safe country.
But nothing the copy said could convince her and so it didn’t take long until a few insidious Copy Writers ambushed her, made her drunk with Longe and Parole and dragged her into their agency, where they abused her for their projects again and again. And if she hasn’t been rewritten, then they are still using her.
Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.
Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.
A small river named Duden flows by their place and supplies it with the necessary regelialia.
It is a paradisematic country, in which roasted parts of sentences fly into your mouth.
Even the all-powerful Pointing has no control about the blind texts it is an almost unorthographic life One
</p>
{{ end }}

</div>
{{ end }}`